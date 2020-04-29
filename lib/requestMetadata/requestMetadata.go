package requestMetadata

import (
	"google.golang.org/grpc/metadata"
	"simple-poa-blockchain-peer/lib/contracts"
)

func GetOwnMetadata(c *contracts.PeerFingerprint) metadata.MD {
	md := make(map[string]string)
	md[contracts.MdID] = c.ID
	md[contracts.MdHost] = c.Host
	md[contracts.MdPort] = c.Port
	return metadata.New(md)
}
