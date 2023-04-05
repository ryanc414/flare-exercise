package main

import (
	"context"
	"log"
	"math/big"
	"time"

	"github.com/alexflint/go-arg"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/ryanc414/flare-exercise/ftso"
	ftsomanager "github.com/ryanc414/flare-exercise/ftsoManager"
	pricesubmitter "github.com/ryanc414/flare-exercise/priceSubmitter"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("godotenv: %s", err)
	}

	if err := run(context.Background()); err != nil {
		log.Fatal(err)
	}
}

type CliArgs struct {
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

	r := runner{
		cfg:     cfg,
		eth:     ethClient,
		ftso:    ftsoClient,
		ftsoABI: ftsoABI,
	}

	return r.run(ctx)
}

type runner struct {
	cfg     CliArgs
	eth     *ethclient.Client
	ftso    *ftso.Ftso
	ftsoABI *abi.ABI
}

func (r *runner) run(ctx context.Context) error {
	ftsoAddrs, err := r.getFTSOAddrs(ctx)
	if err != nil {
		return err
	}

	return r.pollEventsLoop(ctx, ftsoAddrs)
}

func (r *runner) getFTSOAddrs(ctx context.Context) ([]common.Address, error) {
	priceSubmitter, err := pricesubmitter.NewPriceSubmitter(r.cfg.PriceSubmitterAddress, r.eth)
	if err != nil {
		return nil, err
	}

	mngrAddr, err := priceSubmitter.GetFtsoManager(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	log.Print("Manager Address:", mngrAddr)

	ftsoManager, err := ftsomanager.NewFtsoManager(mngrAddr, r.eth)
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

var pollDuration = 5 * time.Second

func (r *runner) pollEventsLoop(
	ctx context.Context, ftsoAddrs []common.Address,
) error {
	for {
		if err := r.pollEvents(ctx, ftsoAddrs); err != nil {
			log.Print("error:", err)
		}

		log.Print("sleeping")
		time.Sleep(pollDuration)
	}
}

var lookbackBlocks = big.NewInt(30)

func (r *runner) pollEvents(
	ctx context.Context, ftsoAddrs []common.Address,
) error {
	currentBlock, err := r.getCurrentBlockNumber(ctx)
	if err != nil {
		return err
	}

	log.Print("currentBlock:", currentBlock)

	fromBlock := new(big.Int).Sub(currentBlock, lookbackBlocks)

	q := ethereum.FilterQuery{
		FromBlock: fromBlock,
		Addresses: ftsoAddrs,
		Topics: [][]common.Hash{{
			r.ftsoABI.Events["PriceRevealed"].ID,
			r.ftsoABI.Events["PriceFinalized"].ID,
		}},
	}

	logs, err := r.eth.FilterLogs(ctx, q)
	if err != nil {
		return err
	}

	log.Printf("got %d logs", len(logs))

	for i := range logs {
		if err := r.processLog(ctx, &logs[i]); err != nil {
			return err
		}
	}

	return nil
}

func (r *runner) getCurrentBlockNumber(ctx context.Context) (*big.Int, error) {
	header, err := r.eth.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}

	return header.Number, nil
}

func (r *runner) processLog(ctx context.Context, lg *types.Log) error {
	switch lg.Topics[0] {
	case r.ftsoABI.Events["PriceRevealed"].ID:
		event, err := r.ftso.ParsePriceRevealed(*lg)
		if err != nil {
			return err
		}

		_ = event

	case r.ftsoABI.Events["PriceFinalized"].ID:
		event, err := r.ftso.ParsePriceFinalized(*lg)
		if err != nil {
			return err
		}

		_ = event

	default:
		return errors.Errorf("unexpected event type: %s", lg.Topics[0])
	}

	return nil
}
