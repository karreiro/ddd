package client

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	proto "ddd/dmn-client/proto"
)

const (
	address = "localhost:50051"
)

var (
	conn       *grpc.ClientConn
	gameClient proto.GameEngineClient
)

// Evaluate calls the DMN engine to execute the game engine
func Evaluate(input *proto.GameInput) *proto.GameOutput {
	return process(input)
}

func init() {
	conn = dial()
	gameClient = proto.NewGameEngineClient(conn)
}

func dial() *grpc.ClientConn {
	clientConn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Error during the gRPC dial...")
	}
	return clientConn
}

func process(input *proto.GameInput) *proto.GameOutput {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	result, err := gameClient.Process(ctx, input)

	defer cancel()

	if err != nil {
		log.Fatal("Error during the DMN call...")
	}
	return result
}
