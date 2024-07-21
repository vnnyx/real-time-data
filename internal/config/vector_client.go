package config

import (
	"context"
	"flag"
	"time"

	"github.com/rs/zerolog"
	"github.com/vnnyx/real-time-data/pb/vector"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewVectorClientConnection(cfg *Config, log *zerolog.Logger) (*grpc.ClientConn, vector.VectorClient, error) {
	addr := flag.String("addr", cfg.VectorHost, "address of the vector service")

	flag.Parse()

	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error().Err(err).Msg("Failed to create gRPC connection")
		return nil, nil, err
	}

	client := vector.NewVectorClient(conn)
	return conn, client, nil
}

func NewVectorClientContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 30*time.Second)
}
