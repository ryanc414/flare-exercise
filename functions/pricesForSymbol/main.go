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
	pricesReq, err := parseRequest(req)
	if err != nil {
		return &events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest}, nil
	}

	res, err := h.pricesForSymbol(ctx, pricesReq)
	if err != nil {
		log.Error().Err(err).Msg("error")
		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
		}, nil
	}

	return res, nil
}

var errBadRequest = errors.New("bad request")

func parseRequest(req *events.APIGatewayProxyRequest) (*PricesRequest, error) {
	pricesReq := &PricesRequest{
		Symbol: req.QueryStringParameters["symbol"],
	}
	if pricesReq.Symbol == "" {
		return nil, errBadRequest
	}

	if startEpochIDRaw := req.QueryStringParameters["startEpochID"]; startEpochIDRaw != "" {
		startEpochID, err := parseBigInt(startEpochIDRaw)
		if err != nil {
			return nil, err
		}

		pricesReq.StartEpochID = startEpochID
	}

	if endEpochIDRaw := req.QueryStringParameters["endEpochID"]; endEpochIDRaw != "" {
		endEpochID, err := parseBigInt(endEpochIDRaw)
		if err != nil {
			return nil, err
		}

		pricesReq.EndEpochID = endEpochID
	}

	if limitRaw := req.QueryStringParameters["limit"]; limitRaw != "" {
		limit, err := strconv.ParseInt(limitRaw, 10, 32)
		if err != nil {
			return nil, err
		}

		if limit <= 0 || limit > 1000 {
			return nil, errBadRequest
		}

		pricesReq.Limit = int(limit)
	}

	if offsetRaw := req.QueryStringParameters["offset"]; offsetRaw != "" {
		offset, err := strconv.ParseInt(offsetRaw, 10, 32)
		if err != nil {
			return nil, err
		}

		pricesReq.Offset = int(offset)
	}

	return pricesReq, nil
}

func parseBigInt(raw string) (*big.Int, error) {
	res, ok := new(big.Int).SetString(raw, 10)
	if !ok {
		return nil, errBadRequest
	}

	return res, nil
}

type PricesRequest struct {
	StartEpochID *big.Int
	EndEpochID   *big.Int
	Symbol       string
	Limit        int
	Offset       int
}

func (h *handler) pricesForSymbol(
	ctx context.Context, pricesReq *PricesRequest,
) (*events.APIGatewayProxyResponse, error) {
	q := buildQuery(pricesReq)
	log.Info().Str("query", q).Msg("built query")

	rows, err := h.db.QueryContext(ctx, q, pricesReq.Symbol)
	if err != nil {
		return nil, err
	}

	var results []PriceFinalized
	for rows.Next() {
		var pf PriceFinalized
		err := rows.Scan(&pf.EpochID, &pf.Price, &pf.RewardedFTSO, &pf.LowRewardPrice, &pf.HighRewardPrice, &pf.FinalizationType, &pf.Timestamp, &pf.Symbol)
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

type PriceFinalized struct {
	EpochID          string `json:"epochID"`
	Price            string `json:"price"`
	RewardedFTSO     bool   `json:"rewardedFTSO"`
	LowRewardPrice   string `json:"lowRewardPrice"`
	HighRewardPrice  string `json:"highRewardPrice"`
	FinalizationType uint8  `json:"FinalizationType"`
	Timestamp        string `json:"Timestamp"`
	Symbol           string `json:"symbol"`
}

func buildQuery(pricesReq *PricesRequest) string {
	var b strings.Builder
	b.WriteString("SELECT EpochID, Price, RewardedFTSO, LowRewardPrice, HighRewardPrice, FinalizationType, Timestamp, Symbol FROM PriceFinalized WHERE Symbol=?")

	if pricesReq.StartEpochID != nil {
		b.WriteString(fmt.Sprintf(" AND EpochID>=%s", pricesReq.StartEpochID))
	}

	if pricesReq.EndEpochID != nil {
		b.WriteString(fmt.Sprintf(" AND EpochID<%s", pricesReq.EndEpochID))
	}

	b.WriteString(" ORDER BY EpochID ASC")

	if pricesReq.Limit != 0 {
		b.WriteString(fmt.Sprintf(" LIMIT %d", pricesReq.Limit))
	}

	if pricesReq.Offset != 0 {
		b.WriteString(fmt.Sprintf(" OFFSET %d", pricesReq.Offset))
	}

	return b.String()
}
