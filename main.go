package main

import (
	"github.com/AwespireTech/InterfaceForCare-Indexer/config"
	"github.com/AwespireTech/InterfaceForCare-Indexer/database"
	"github.com/AwespireTech/InterfaceForCare-Indexer/tezos"
)

func main() {
	database.Init("***REMOVED***")
	tezos.Init(config.RPCURL)
	res, err := tezos.GetRiverList(config.FACTORY_BIGMAP)
	if err != nil {
		panic(err)
	}
	for _, river := range res {
		tezos.PrintContractStorage(river)
	}

}
