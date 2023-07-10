package main

import (
	"context"
	"github.com/basue_sanshita/internal/infrastructure/api"
	"math/rand"
	"time"
)

// Unary RPC = 1 : 1
// Server streaming RPC = 1 : n >> file download
// Client streaming RPC = n : 1  >> file upload
// Bidirectional streaming RPC >> like websocket(chatとか)
func main() {
	rand.Seed(time.Now().UnixNano())

	ctx := context.Background()
	err := api.Run(ctx)
	if err != nil {
		panic(err)
	}
}
