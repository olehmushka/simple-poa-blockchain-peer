package config

import (
  baseconfig "simple-poa-blockchain-peer/lib/baseConfig"
)

type settingsKeys struct {
	PeerHost string `name:"PEER_HOST" default:"0.0.0.0"`
  PeerPort string `name:"PEER_PORT"`
}

var Keys = settingsKeys{}

func InitConfig() {
  baseconfig.ReadConfigValues(&Keys)
}
