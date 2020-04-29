package server

import (
	"context"
	"google.golang.org/grpc"

	pb "simple-poa-blockchain-peer/orchestrator/modules/peer"
)

type PeerServer struct {
	pb.UnimplementedPeerServiceServer
}

func (s *PeerServer) GetAll(c context.Context, request *pb.Empty) (*pb.Peers, error) {
	return &pb.Peers{Peers: make([]*pb.Peer, 0)}, nil
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

func RegisterPeerServer(s *grpc.Server) {
  pb.RegisterPeerServiceServer(s, &PeerServer{})
}
