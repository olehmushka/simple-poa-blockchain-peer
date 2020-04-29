package contracts

type PeerFingerprint struct {
	ID   string
	Host string
	Port string
}

const (
	MdID   = "ID"
	MdHost = "HOST"
	MdPort = "PORT"
)
