package main

import (
	"simple-poa-blockchain-peer/config"
	"simple-poa-blockchain-peer/orchestrator"
)

func main() {
  config.InitConfig()
	o := orchestrator.NewOrchestrator(config.Keys.PeerPort)
	o.Run()
}
