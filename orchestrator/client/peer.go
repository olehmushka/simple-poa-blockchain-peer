package client

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	rm "simple-poa-blockchain-peer/lib/requestMetadata"
	pb "simple-poa-blockchain-peer/orchestrator/modules/peer"
	"simple-poa-blockchain-peer/orchestrator/state"
)

func getClientConn(host, port string) (pb.PeerServiceClient, *grpc.ClientConn, error) {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
	}
	conn, err := grpc.Dial(host+port, opts...)
	if err != nil {
		return nil, nil, err
	}

	return pb.NewPeerServiceClient(conn), conn, nil
}

func getMetadata() context.Context {
	md := rm.GetOwnMetadata(state.GetState().GetFingerprint())
	return metadata.NewOutgoingContext(context.Background(), md)
}

func GetAllPeers(host, port string) ([]*pb.Peer, error) {
	client, conn, err := getClientConn(host, port)
	if err != nil {
		return []*pb.Peer{}, err
	}
	defer conn.Close()

	response, err := client.GetAll(getMetadata(), &pb.Empty{})
	if err != nil {
		return []*pb.Peer{}, err
	}
	return response.Peers, nil
}

func AddPeer(host, port, peerId string) (*pb.Empty, error) {
	client, conn, err := getClientConn(host, port)
	if err != nil {
		return &pb.Empty{}, err
	}
	defer conn.Close()

	response, err := client.Add(getMetadata(), &pb.PeerID{PeerID: peerId})
	if err != nil {
		return &pb.Empty{}, err
	}
	return response, nil
}

func RemovePeer(host, port, peerId string) (*pb.Empty, error) {
	client, conn, err := getClientConn(host, port)
	if err != nil {
		return &pb.Empty{}, err
	}
	defer conn.Close()

	response, err := client.Remove(getMetadata(), &pb.PeerID{PeerID: peerId})
	if err != nil {
		return &pb.Empty{}, err
	}
	return response, nil
}

func PeerSelfAdd(host, port string) (*pb.Empty, error) {
	client, conn, err := getClientConn(host, port)
	if err != nil {
		return &pb.Empty{}, err
	}
	defer conn.Close()

	response, err := client.SelfAdd(getMetadata(), &pb.Empty{})
	if err != nil {
		return &pb.Empty{}, err
	}
	return response, nil
}

func PeerSelfRemove(host, port string) (*pb.Empty, error) {
	client, conn, err := getClientConn(host, port)
	if err != nil {
		return &pb.Empty{}, err
	}
	defer conn.Close()

	response, err := client.SelfRemove(getMetadata(), &pb.Empty{})
	if err != nil {
		return &pb.Empty{}, err
	}
	return response, nil
}
