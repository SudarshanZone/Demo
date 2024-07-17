package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/krishnakashyap0704/microservices/tradelist/internal/database" // Import the database package
	pb "github.com/krishnakashyap0704/microservices/tradelist/generated"
)

type server struct {
	pb.UnimplementedTradeListServiceServer
	db *sql.DB
}

// TradeListOrder implements tradelist.TradeListServiceServer
func (s *server) TradeListOrder(ctx context.Context, req *pb.TradeListRequest) (*pb.TradeListResponse, error) {
	if s.db == nil {
		return nil, fmt.Errorf("database connection is nil")
	}

	query := `SELECT from_date, to_date, exchange_code, product_type, action, stock_code 
	          FROM tradelist 
	          WHERE exchange_code = $1 
	          AND ($2::text IS NULL OR product_type = $2) 
	          AND ($3::text IS NULL OR action = $3) 
	          AND ($4::text IS NULL OR stock_code = $4)`

	log.Printf("Executing query: %s\nWith params: %v, %v, %v, %v\n", query, req.ExchangeCode, req.ProductType, req.Action, req.StockCode)

	rows, err := s.db.QueryContext(ctx, query, req.ExchangeCode, req.ProductType, req.Action, req.StockCode)
	if err != nil {
		return nil, fmt.Errorf("error querying database: %v", err)
	}
	defer rows.Close()

	var trades []*pb.TradeListResponse_Trade
	for rows.Next() {
		var trade pb.TradeListResponse_Trade
		err := rows.Scan(
			&trade.FromDate,
			&trade.ToDate,
			&trade.ExchangeCode,
			&trade.ProductType,
			&trade.Action,
			&trade.StockCode,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		log.Printf("Scanned trade: %+v\n", trade)
		trades = append(trades, &trade)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	log.Printf("Returning %d trades\n", len(trades))

	return &pb.TradeListResponse{TradeList: trades}, nil
}

func main() {
	// Initialize the database connection
	database.Init()
	db := database.GetDB()
	defer db.Close()

	// Check if the db is nil
	if db == nil {
		log.Fatalf("database connection is nil")
	} else {
		log.Println("Database connection is successfully initialized")
	}

	// Test the database connection
	var version string
	err := db.QueryRow("SELECT version()").Scan(&version)
	if err != nil {
		log.Fatalf("Error querying database: %v", err)
	}
	log.Println("Database version:", version)

	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	tradeListService := &server{db: db} // Pass the database connection to the server struct
	pb.RegisterTradeListServiceServer(s, tradeListService)
	reflection.Register(s)

	log.Println("gRPC server is running on port 8089")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
