package handler

import (
	"context"
	pb "github.com/basue_sanshita/internal/infrastructure/api/grpc_gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"math/rand"
	"sync"
	"time"
)

type PancakeBakerHandler struct {
	report *report
	pb.UnimplementedPancakeBakerServiceServer
}

type report struct {
	sync.Mutex
	data map[pb.Pancake_Menu]int
}

func NewPancakeBakerHandler() *PancakeBakerHandler {
	return &PancakeBakerHandler{
		report: &report{
			data: make(map[pb.Pancake_Menu]int),
		},
	}
}

func (h *PancakeBakerHandler) Bake(ctx context.Context, req *pb.BakeRequest) (*pb.BakeResponse, error) {
	if req.Menu == pb.Pancake_UNKNOWN || req.Menu > pb.Pancake_SPICY_CURRY {
		return nil, status.Errorf(codes.InvalidArgument, "パンケーキを選んでください")
	}

	now := time.Now()
	h.report.Lock()
	h.report.data[req.Menu] = h.report.data[req.Menu] + 1
	h.report.Unlock()

	return &pb.BakeResponse{Pancake: &pb.Pancake{Menu: req.Menu,
		ChefName:       "hoge",
		TechnicalScore: rand.Float32(),
		CreateTime: &timestamppb.Timestamp{
			Seconds: now.Unix(),
			Nanos:   int32(now.Nanosecond()),
		},
	}}, nil
}

func (h *PancakeBakerHandler) Report(ctx context.Context, req *pb.ReportRequest) (*pb.ReportResponse, error) {
	counts := make([]*pb.Report_BakeCount, 0)

	h.report.Lock()
	for k, v := range h.report.data {
		counts = append(counts, &pb.Report_BakeCount{Menu: k, Count: int32(v)})
	}
	h.report.Unlock()

	return &pb.ReportResponse{Report: &pb.Report{BakeCounts: counts}}, nil
}
