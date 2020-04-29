package server

import (
  "github.com/rs/zerolog"
  "google.golang.org/grpc"

  "net"
)

type Server struct {
	Log        zerolog.Logger
	port       string
}

func NewServer(log zerolog.Logger, port string) *Server {
	return &Server{
		Log: log,
		port: port,
	}
}

func (s *Server) Run() {
	lis, err := net.Listen("tcp", s.port)
	if err != nil {
		s.Log.Error().Err(err).Msgf("Failed to listen %s port", s.port)
	}
	server := grpc.NewServer()
	RegisterPeerServer(server, s.Log)

  s.Log.Info().Msgf("Server is running on %s port...", s.port)
	if err := server.Serve(lis); err != nil {
    s.Log.Error().Err(err).Msg("Failed to server gRPC server")
  }
}
