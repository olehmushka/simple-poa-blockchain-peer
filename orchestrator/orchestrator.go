package orchestrator

import (
  "github.com/rs/zerolog"
  "os"
  "simple-poa-blockchain-peer/lib/loghook"
  "simple-poa-blockchain-peer/orchestrator/server"
)

type Orchestrator struct {
	Server *server.Server
	Log    zerolog.Logger
}

func NewOrchestrator(srvPort string) *Orchestrator {
	log := zerolog.New(os.Stderr).With().Timestamp().Logger().Hook(loghook.Caller).Hook(loghook.Severity)
	return &Orchestrator{
		Server: server.NewServer(log, srvPort),
		Log:    log,
	}
}

func (o *Orchestrator) Run() {
	o.Server.Run()
}
