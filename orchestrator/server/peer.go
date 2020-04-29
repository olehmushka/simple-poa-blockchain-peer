package server

import (
	"context"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"

	pb "simple-poa-blockchain-peer/orchestrator/modules/peer"
)

type PeerServer struct {
	log zerolog.Logger
}

func (s *PeerServer) GetAll(c context.Context, request *pb.Empty) (response *pb.Peers, err error) {
	response.Peers = make([]*pb.Peer, 0)
	return response, nil
}

func (s *PeerServer) Add(c context.Context, request *pb.PeerID) (response *pb.Empty, err error) {
	return response, nil
}

func (s *PeerServer) Remove(c context.Context, request *pb.PeerID) (response *pb.Empty, err error) {
	return response, nil
}

func (s *PeerServer) SelfAdd(c context.Context, request *pb.Empty) (response *pb.Empty, err error) {
	return response, nil
}

func (s *PeerServer) SelfRemove(c context.Context, request *pb.Empty) (response *pb.Empty, err error) {
	return response, nil
}

func RegisterPeerServer(s *grpc.Server, log zerolog.Logger) {
  pb.RegisterPeerServiceServer(s, &PeerServer{log})
}
