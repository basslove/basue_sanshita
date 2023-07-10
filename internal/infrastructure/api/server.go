package api

import (
	"context"
	"fmt"
	"github.com/basue_sanshita/internal/infrastructure/api/handler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
)

// Unary RPC = 1 : 1
// Server streaming RPC = 1 : n >> file download
// Client streaming RPC = n : 1  >> file upload
// Bidirectional streaming RPC >> like websocket(chatとか)
func Run(ctx context.Context) error {
	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	s := grpc.NewServer()

	h := handler.NewHandler()
	h.Register(s)

	reflection.Register(s)

	go func() {
		log.Printf("start gRPC server port: %v", port)
		s.Serve(listener)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()

	return nil
}
