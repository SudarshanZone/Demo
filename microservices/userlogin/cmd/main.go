package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/krishnakashyap0704/microservices/userlogin/internal/database"
	pb "github.com/krishnakashyap0704/microservices/userlogin/generated"
)

type server struct {
	pb.UnimplementedUserLoginServiceServer
	db *sql.DB
}

func (s *server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var hashedPassword string
	err := s.db.QueryRow("SELECT password FROM users WHERE username=$1", req.GetUsername()).Scan(&hashedPassword)
	if err != nil {
		return &pb.LoginResponse{Success: false, Message: "Invalid username or password"}, nil
	}

	// Compare the provided password with the hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.GetPassword())); err != nil {
		return &pb.LoginResponse{Success: false, Message: "Invalid username or password"}, nil
	}

	// Log the login event
	_, err = s.db.Exec("INSERT INTO userlog (username, login_time) VALUES ($1, NOW())", req.GetUsername())
	if err != nil {
		return nil, fmt.Errorf("failed to log login event: %v", err)
	}

	return &pb.LoginResponse{Success: true, Message: "Login successful"}, nil
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
	userLogService := &server{db: db} // Pass the database connection to the server struct
	pb.RegisterUserLoginServiceServer(s, userLogService)
	reflection.Register(s)

	log.Println("gRPC server is running on port 8089")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
