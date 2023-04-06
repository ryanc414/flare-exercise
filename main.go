package main

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"math/big"
	"sync"
	"time"

	"github.com/alexflint/go-arg"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/ryanc414/flare-exercise/ftso"
	ftsomanager "github.com/ryanc414/flare-exercise/ftsoManager"
	pricesubmitter "github.com/ryanc414/flare-exercise/priceSubmitter"
	"golang.org/x/sync/errgroup"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("godotenv: %s", err)
	}

	if err := run(context.Background()); err != nil {
		log.Fatal().Err(err).Send()
	}
}

type CliArgs struct {
	DBName                string         `arg:"--db-name,required,env:DB_NAME"`
	PriceSubmitterAddress common.Address `arg:"--price-submitter-addr,required,env:PRICE_SUBMITTER_ADDRESS"`
	RPC                   string         `arg:"--rpc,required,env:RPC_URL"`
}

func run(ctx context.Context) error {
	var cfg CliArgs
	arg.MustParse(&cfg)

	ethClient, err := ethclient.DialContext(ctx, cfg.RPC)
	if err != nil {
		return err
	}

	ftsoClient, err := ftso.NewFtso(common.Address{}, ethClient)
	if err != nil {
		return err
	}

	ftsoABI, err := ftso.FtsoMetaData.GetAbi()
	if err != nil {
		return err
	}

	db, err := initDB(cfg.DBName)
	if err != nil {
		return err
	}

	defer db.Close()

	ftsoAddrs, err := getFTSOAddrs(ctx, ethClient, cfg.PriceSubmitterAddress)
	if err != nil {
		return err
	}

	symbols, err := buildSymbolsMap(ctx, ethClient, ftsoAddrs)
	if err != nil {
		return err
	}

	p := eventPoller{
		cfg:       cfg,
		db:        db,
		eth:       ethClient,
		ftso:      ftsoClient,
		ftsoABI:   ftsoABI,
		ftsoAddrs: ftsoAddrs,
		symbols:   symbols,
	}

	return p.pollEventsLoop(ctx)
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

func getFTSOAddrs(
	ctx context.Context, ethClient *ethclient.Client, priceSubmitterAddr common.Address,
) ([]common.Address, error) {
	priceSubmitter, err := pricesubmitter.NewPriceSubmitter(priceSubmitterAddr, ethClient)
	if err != nil {
		return nil, err
	}

	mngrAddr, err := priceSubmitter.GetFtsoManager(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	log.Print("Manager Address:", mngrAddr)

	ftsoManager, err := ftsomanager.NewFtsoManager(mngrAddr, ethClient)
	if err != nil {
		return nil, err
	}

	ftsoAddrs, err := ftsoManager.GetFtsos(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	log.Print("FTSO Addresses:", len(ftsoAddrs), ftsoAddrs)
	return ftsoAddrs, nil
}

func buildSymbolsMap(
	ctx context.Context, ethClient *ethclient.Client, ftsoAddrs []common.Address,
) (map[common.Address]string, error) {
	eg, ctx := errgroup.WithContext(ctx)
	symbols := make(map[common.Address]string)
	var mu sync.Mutex
	opts := &bind.CallOpts{Context: ctx}

	for i := range ftsoAddrs {
		addr := ftsoAddrs[i]

		eg.Go(func() error {
			ftsoClient, err := ftso.NewFtso(addr, ethClient)
			if err != nil {
				return err
			}

			sym, err := ftsoClient.Symbol(opts)
			if err != nil {
				return err
			}

			mu.Lock()
			symbols[addr] = sym
			mu.Unlock()

			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	return symbols, nil
}

type eventPoller struct {
	cfg       CliArgs
	db        *sql.DB
	eth       *ethclient.Client
	ftso      *ftso.Ftso
	ftsoABI   *abi.ABI
	ftsoAddrs []common.Address
	symbols   map[common.Address]string
}

var pollDuration = 5 * time.Second

func (p *eventPoller) pollEventsLoop(ctx context.Context) error {
	for {
		if err := p.pollEvents(ctx); err != nil {
			log.Print("error:", err)
		}

		log.Print("sleeping")
		time.Sleep(pollDuration)
	}
}

var lookbackBlocks = big.NewInt(15)

func (p *eventPoller) pollEvents(ctx context.Context) error {
	currentBlock, err := p.getCurrentBlockNumber(ctx)
	if err != nil {
		return err
	}

	log.Print("currentBlock:", currentBlock)

	fromBlock := new(big.Int).Sub(currentBlock, lookbackBlocks)

	q := ethereum.FilterQuery{
		FromBlock: fromBlock,
		Addresses: p.ftsoAddrs,
		Topics: [][]common.Hash{{
			p.ftsoABI.Events["PriceRevealed"].ID,
			p.ftsoABI.Events["PriceFinalized"].ID,
		}},
	}

	logs, err := p.eth.FilterLogs(ctx, q)
	if err != nil {
		return err
	}

	log.Info().Msgf("got %d logs", len(logs))

	events, err := p.parseEvents(logs)
	if err != nil {
		return err
	}

	log.Info().Int("PR", len(events.PriceRevealed)).Int("PF", len(events.PriceFinalized)).Msg("parsed events")

	return p.storeEvents(ctx, events)
}

func (p *eventPoller) getCurrentBlockNumber(ctx context.Context) (*big.Int, error) {
	header, err := p.eth.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}

	return header.Number, nil
}

type Events struct {
	PriceRevealed  []*ftso.FtsoPriceRevealed
	PriceFinalized []*ftso.FtsoPriceFinalized
}

func (p *eventPoller) parseEvents(logs []types.Log) (*Events, error) {
	var events Events

	for i := range logs {
		lg := &logs[i]

		switch lg.Topics[0] {
		case p.ftsoABI.Events["PriceRevealed"].ID:
			event, err := p.ftso.ParsePriceRevealed(*lg)
			if err != nil {
				return nil, err
			}

			events.PriceRevealed = append(events.PriceRevealed, event)

		case p.ftsoABI.Events["PriceFinalized"].ID:
			event, err := p.ftso.ParsePriceFinalized(*lg)
			if err != nil {
				return nil, err
			}

			events.PriceFinalized = append(events.PriceFinalized, event)

		default:
			return nil, errors.Errorf("unexpected event type: %s", lg.Topics[0])
		}
	}

	return &events, nil
}

func (p *eventPoller) storeEvents(ctx context.Context, events *Events) error {
	for _, event := range events.PriceRevealed {
		if err := p.storePriceRevealed(ctx, event); err != nil {
			return err
		}
	}

	for _, event := range events.PriceFinalized {
		if err := p.storePriceFinalized(ctx, event); err != nil {
			return err
		}
	}

	return nil
}

func (p *eventPoller) storePriceRevealed(ctx context.Context, event *ftso.FtsoPriceRevealed) error {
	sym := p.symbols[event.Raw.Address]
	if sym == "" {
		return errors.Errorf("symbol not found for FTSO %s", event.Raw.Address)
	}

	pk := getPRKey(event.EpochId, event.Voter, sym)

	res, err := p.db.Exec(
		"INSERT INTO PriceRevealed (PK, EpochID, Voter, Price, Timestamp, VotePowerNat, VotePowerAsset, Symbol) VALUES (?, ?, ?, ?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE PK=PK",
		pk,
		event.EpochId.Int64(),
		event.Voter.Hex(),
		event.Price.String(),
		event.Timestamp.String(),
		event.VotePowerNat.String(),
		event.VotePowerAsset.String(),
		sym,
	)
	if err != nil {
		return err
	}

	_, err = res.LastInsertId()
	return err
}

var i32Max = new(big.Int).Exp(big.NewInt(2), big.NewInt(31), nil)

func getPRKey(epochID *big.Int, voter common.Address, sym string) int32 {
	h := sha256.New()
	h.Write(epochID.Bytes())
	h.Write(voter[:])
	h.Write([]byte(sym))

	bs := h.Sum(nil)
	bsNum := new(big.Int).SetBytes(bs)
	bs32 := new(big.Int).Mod(bsNum, i32Max)

	return int32(bs32.Int64())
}

func (p *eventPoller) storePriceFinalized(ctx context.Context, event *ftso.FtsoPriceFinalized) error {
	sym := p.symbols[event.Raw.Address]
	if sym == "" {
		return errors.Errorf("symbol not found for FTSO %s", event.Raw.Address)
	}

	pk := getPFKey(event.EpochId, sym)

	res, err := p.db.Exec(
		"INSERT INTO PriceFinalized (PK, EpochID, Price, RewardedFTSO, LowRewardPrice, HighRewardPrice, FinalizationType, Timestamp, Symbol) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE PK=PK",
		pk,
		event.EpochId.Int64(),
		event.Price.String(),
		event.RewardedFtso,
		event.LowRewardPrice.String(),
		event.HighRewardPrice.String(),
		event.FinalizationType,
		event.Timestamp.String(),
		sym,
	)
	if err != nil {
		return err
	}

	_, err = res.LastInsertId()
	return err
}

func getPFKey(epochID *big.Int, sym string) int32 {
	h := sha256.New()
	h.Write(epochID.Bytes())
	h.Write([]byte(sym))

	bs := h.Sum(nil)
	bsNum := new(big.Int).SetBytes(bs)
	bs32 := new(big.Int).Mod(bsNum, i32Max)

	return int32(bs32.Int64())
}
