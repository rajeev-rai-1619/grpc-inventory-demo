package main

import (
	"context"
	"log"
	"time"

	inventorypb "grpc-inventory-demo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := inventorypb.NewInventoryServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	resp, err := client.CheckStock(ctx, &inventorypb.StockRequest{
		ItemId:       "SKU-12345",
		RequestedQty: 2,
	})
	if err != nil {
		log.Fatalf("CheckStock RPC failed: %v", err)
	}

	if resp.GetAvailable() {
		log.Printf("Item available. Remaining quantity: %d", resp.GetRemainingQty())
		return
	}

	log.Printf("Item not available. Current quantity: %d", resp.GetRemainingQty())
}
