package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/caarlos0/env/v6"
	"github.com/ethereum/go-ethereum/common"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type funcConfig struct {
	DBName string `env:"DB_NAME,required"`
}

func main() {
	if err := run(context.Background()); err != nil {
		log.Fatal().Err(err).Stack().Send()
	}
}

func run(ctx context.Context) error {
	var cfg funcConfig
	if err := env.Parse(&cfg); err != nil {
		return errors.Wrap(err, "config.LoadDefaultConfig")
	}

	db, err := initDB(cfg.DBName)
	if err != nil {
		return err
	}

	h := handler{
		db: db,
	}

	lambda.Start(h.handle)
	return nil
}

func initDB(dbName string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dbName)
	if err != nil {
		return nil, err
	}

	// DB configuration
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

type handler struct {
	db *sql.DB
}

func (h *handler) handle(
	ctx context.Context, req *events.APIGatewayProxyRequest,
) (*events.APIGatewayProxyResponse, error) {
	votesReq, err := parseRequest(req)
	if err != nil {
		return &events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest}, nil
	}

	res, err := h.votesForVoter(ctx, votesReq)
	if err != nil {
		log.Error().Err(err).Msg("error")
		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
		}, nil
	}

	return res, nil
}

var errBadRequest = errors.New("bad request")

func parseRequest(req *events.APIGatewayProxyRequest) (*VotesRequest, error) {
	votesReq := &VotesRequest{
		Symbol: req.QueryStringParameters["symbol"],
		Voter:  common.HexToAddress(req.PathParameters["voterAddress"]),
	}
	if votesReq.Voter == (common.Address{}) {
		return nil, errBadRequest
	}

	if startEpochIDRaw := req.QueryStringParameters["startEpochID"]; startEpochIDRaw != "" {
		startEpochID, err := parseBigInt(startEpochIDRaw)
		if err != nil {
			return nil, err
		}

		votesReq.StartEpochID = startEpochID
	}

	if endEpochIDRaw := req.QueryStringParameters["endEpochID"]; endEpochIDRaw != "" {
		endEpochID, err := parseBigInt(endEpochIDRaw)
		if err != nil {
			return nil, err
		}

		votesReq.EndEpochID = endEpochID
	}

	if limitRaw := req.QueryStringParameters["limit"]; limitRaw != "" {
		limit, err := strconv.ParseInt(limitRaw, 10, 32)
		if err != nil {
			return nil, err
		}

		if limit <= 0 || limit > 1000 {
			return nil, errBadRequest
		}

		votesReq.Limit = int(limit)
	}

	if offsetRaw := req.QueryStringParameters["offset"]; offsetRaw != "" {
		offset, err := strconv.ParseInt(offsetRaw, 10, 32)
		if err != nil {
			return nil, err
		}

		votesReq.Offset = int(offset)
	}

	return votesReq, nil
}

func parseBigInt(raw string) (*big.Int, error) {
	res, ok := new(big.Int).SetString(raw, 10)
	if !ok {
		return nil, errBadRequest
	}

	return res, nil
}

type VotesRequest struct {
	StartEpochID *big.Int
	EndEpochID   *big.Int
	Symbol       string
	Limit        int
	Offset       int
	Voter        common.Address
}

func (h *handler) votesForVoter(
	ctx context.Context, votesReq *VotesRequest,
) (*events.APIGatewayProxyResponse, error) {
	q := buildQuery(votesReq)
	log.Info().Str("query", q).Msg("built query")

	rows, err := h.db.QueryContext(ctx, q, votesReq.Voter.String())
	if err != nil {
		return nil, err
	}

	var results []PriceRevealed
	for rows.Next() {
		var pf PriceRevealed
		err := rows.Scan(&pf.EpochID, &pf.Voter, &pf.Price, &pf.Timestamp, &pf.VotePowerNat, &pf.VotePowerAsset, &pf.Symbol)
		if err != nil {
			return nil, err
		}

		results = append(results, pf)
	}

	resultsData, err := json.Marshal(results)
	if err != nil {
		return nil, err
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(resultsData),
	}, nil
}

type PriceRevealed struct {
	EpochID        string `json:"epochID"`
	Voter          string `json:"voter"`
	Price          string `json:"price"`
	Timestamp      string `json:"Timestamp"`
	VotePowerNat   string `json:"votePowerNat"`
	VotePowerAsset string `json:"votePowerAsset"`
	Symbol         string `json:"symbol"`
}

func buildQuery(votesReq *VotesRequest) string {
	var b strings.Builder
	b.WriteString("SELECT EpochID, Voter, Price, Timestamp, VotePowerNat, VotePowerAsset, Symbol FROM PriceRevealed WHERE Voter=?")

	if votesReq.Symbol != "" {
		b.WriteString(fmt.Sprintf(" AND Symbol='%s'", votesReq.Symbol))
	}

	if votesReq.StartEpochID != nil {
		b.WriteString(fmt.Sprintf(" AND EpochID>=%s", votesReq.StartEpochID))
	}

	if votesReq.EndEpochID != nil {
		b.WriteString(fmt.Sprintf(" AND EpochID<%s", votesReq.EndEpochID))
	}

	b.WriteString(" ORDER BY EpochID ASC")

	if votesReq.Limit != 0 {
		b.WriteString(fmt.Sprintf(" LIMIT %d", votesReq.Limit))
	}

	if votesReq.Offset != 0 {
		b.WriteString(fmt.Sprintf(" OFFSET %d", votesReq.Offset))
	}

	return b.String()
}
