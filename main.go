package main

import (
	"context"
	"log"

	"github.com/alexflint/go-arg"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
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

	priceSubmitter, err := pricesubmitter.NewPriceSubmitter(cfg.PriceSubmitterAddress, ethClient)
	if err != nil {
		return err
	}

	mngrAddr, err := priceSubmitter.GetFtsoManager(&bind.CallOpts{Context: ctx})
	if err != nil {
		return err
	}

	log.Print("Manager Address:", mngrAddr)

	ftsoManager, err := ftsomanager.NewFtsoManager(mngrAddr, ethClient)
	if err != nil {
		return err
	}

	ftsoAddrs, err := ftsoManager.GetFtsos(&bind.CallOpts{Context: ctx})
	if err != nil {
		return err
	}

	log.Print("FTSO Addresses:", len(ftsoAddrs), ftsoAddrs)

	return nil
}
