package state

import (
	"github.com/gofrs/uuid"
	"simple-poa-blockchain-peer/config"
	"simple-poa-blockchain-peer/lib/contracts"
)

type State struct {
	ID   string
	Host string
	Port string
}

func (s *State) GetFingerprint() *contracts.PeerFingerprint {
	return &contracts.PeerFingerprint{
		ID:   s.ID,
		Host: s.Host,
		Port: s.Port,
	}
}

var instance *State

func GetState() *State {
  if instance == nil {
    instance = createState()
  }
  return instance
}

func createState() *State {
	id, _ := uuid.NewV4()
	return &State{
		ID:   id.String(),
		Host: config.Keys.PeerHost,
		Port: config.Keys.PeerPort,
	}
}
