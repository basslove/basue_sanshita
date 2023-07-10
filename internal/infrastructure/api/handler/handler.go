package handler

import (
	pb "github.com/basue_sanshita/internal/infrastructure/api/grpc_gen"
	"google.golang.org/grpc"
)

type Handler struct {
	PancakeBakerHandler *PancakeBakerHandler
}

func NewHandler() *Handler {
	h := Handler{}
	h.PancakeBakerHandler = NewPancakeBakerHandler()

	return &h
}

func (h *Handler) Register(s *grpc.Server) {
	// $ grpcurl -plaintext localhost:8080 list
	pb.RegisterPancakeBakerServiceServer(s, h.PancakeBakerHandler)
}
