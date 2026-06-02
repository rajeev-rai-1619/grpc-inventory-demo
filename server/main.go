package main

import (
	"context"
	"log"
	"net"

	inventorypb "grpc-inventory-demo/proto"
	"google.golang.org/grpc"
)

type inventoryServer struct {
	inventorypb.UnimplementedInventoryServiceServer
	stock map[string]int32
}

func newInventoryServer() *inventoryServer {
	return &inventoryServer{
		stock: map[string]int32{
			"SKU-12345": 12,
			"SKU-99999": 0,
		},
	}
}

func (s *inventoryServer) CheckStock(_ context.Context, req *inventorypb.StockRequest) (*inventorypb.StockResponse, error) {
	currentStock := s.stock[req.GetItemId()]
	available := currentStock >= req.GetRequestedQty()

	log.Printf("CheckStock item=%s requested=%d available=%t remaining=%d",
		req.GetItemId(),
		req.GetRequestedQty(),
		available,
		currentStock,
	)

	return &inventorypb.StockResponse{
		Available:    available,
		RemainingQty: currentStock,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	inventorypb.RegisterInventoryServiceServer(grpcServer, newInventoryServer())

	log.Println("Inventory gRPC server listening on :50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve gRPC server: %v", err)
	}
}
