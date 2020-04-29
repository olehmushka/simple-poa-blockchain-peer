package config

import (
  baseconfig "simple-poa-blockchain-peer/lib/baseConfig"
)

type settingsKeys struct {
  PeerPort string `name:"PEER_PORT"`
}

var Keys = settingsKeys{}

func InitConfig() {
  baseconfig.ReadConfigValues(&Keys)
}
