package main

import (
	"context"
	"fmt"
	"github.com/InjectiveLabs/sdk-go/client/core"
	exchangeclient "github.com/InjectiveLabs/sdk-go/client/exchange"

	"github.com/InjectiveLabs/sdk-go/client"

	chainclient "github.com/InjectiveLabs/sdk-go/client/chain"
	"github.com/InjectiveLabs/sdk-go/client/common"
)

func main() {
	network := common.LoadNetwork("mainnet", "lb")

	clientCtx, err := chainclient.NewClientContext(
		network.ChainId,
		"",
		nil,
	)
	if err != nil {
		panic(err)
	}

	exchangeClient, err := exchangeclient.NewExchangeClient(network)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	marketsAssistant, err := core.NewMarketsAssistantUsingExchangeClient(ctx, exchangeClient)
	if err != nil {
		panic(err)
	}

	chainClient, err := chainclient.NewChainClientWithMarketsAssistant(
		clientCtx,
		network,
		marketsAssistant,
		common.OptionGasPrices(client.DefaultGasPriceWithDenom),
	)

	if err != nil {
		panic(err)
	}

	failEventCh := make(chan map[string]uint, 10000)
	go chainClient.StreamEventOrderFail("inj1rwv4zn3jptsqs7l8lpa3uvzhs57y8duemete9e", failEventCh)
	for {
		e := <-failEventCh
		fmt.Println(e)
	}
}
