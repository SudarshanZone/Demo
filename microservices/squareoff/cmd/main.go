package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/krishnakashyap0704/microservices/squareoff/generated"
)

type server struct {
	pb.UnimplementedSquareOffServiceServer
	db *sql.DB
}

func (s *server) SquareOffOrder(ctx context.Context, req *pb.SquareOffRequest) (*pb.SquareOffResponse, error) {
	query := `
		SELECT trade_id, stock_code, quantity, action 
		FROM tradelist 
		WHERE product_type = 'intraday' 
		  AND status = 'pending'
	`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tradeID int
		var stockCode string
		var quantity int
		var action string
		if err := rows.Scan(&tradeID, &stockCode, &quantity, &action); err != nil {
			return nil, err
		}

		var apiResponse *pb.SquareOffResponse

		if action == "buy" {
			// Square off open long positions by selling them
			apiResponse = callExternalAPIToSellStock(stockCode, quantity)
		} else if action == "sell" {
			// Square off short positions by buying them back
			apiResponse = callExternalAPIToBuyStock(stockCode, quantity)
		}

		if apiResponse.Status != "success" {
			return nil, fmt.Errorf("failed to square off stock: %v", apiResponse.Message)
		}

		// Update the trade status to 'completed'
		_, err = s.db.Exec("UPDATE tradelist SET status = 'completed' WHERE trade_id = $1", tradeID)
		if err != nil {
			return nil, err
		}
	}

	return &pb.SquareOffResponse{
		Status:  "success",
		Message: "Orders squared off successfully",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterSquareOffServiceServer(s, &server{})
	reflection.Register(s)

	log.Println("gRPC server is running on port 8089")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// Dummy function to represent external API call for selling stock
func callExternalAPIToSellStock(stockCode string, quantity int) *pb.SquareOffResponse {
	// This function should be replaced with actual implementation
	return &pb.SquareOffResponse{
		Status:  "success",
		Message: fmt.Sprintf("Sold %d of %s", quantity, stockCode),
	}
}

// Dummy function to represent external API call for buying stock
func callExternalAPIToBuyStock(stockCode string, quantity int) *pb.SquareOffResponse {
	// This function should be replaced with actual implementation
	return &pb.SquareOffResponse{
		Status:  "success",
		Message: fmt.Sprintf("Bought %d of %s", quantity, stockCode),
	}
}
