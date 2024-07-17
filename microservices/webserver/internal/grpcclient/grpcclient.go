package grpcclient

import (
	"log"
	"google.golang.org/grpc"
	
	// pbSquareOff "github.com/krishnakashyap0704/microservices/squareoff/generated"
    // pbTradeList "github.com/krishnakashyap0704/microservices/tradelist/generated"
    // pbUserLogin "github.com/krishnakashyap0704/microservices/userlogin/generated"

	
	pbSquareOff "github.com/krishnakashyap0704/microservices/squareoff/generated"
    pbTradeList "github.com/krishnakashyap0704/microservices/tradelist/generated"
    pbUserLogin "github.com/krishnakashyap0704/microservices/userlogin/generated"
	
)

func NewSquareOffClient() (pbSquareOff.SquareOffServiceClient, error) {
	conn, err := grpc.Dial("localhost:8089", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil, err
	}
	return pbSquareOff.NewSquareOffServiceClient(conn), nil
}

func NewTradeListClient() (pbTradeList.TradeListServiceClient, error) {
	conn, err := grpc.Dial("localhost:8089", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil, err
	}
	return pbTradeList.NewTradeListServiceClient(conn), nil
}

func NewUserLoginClient() (pbUserLogin.UserLoginServiceClient, error) {
	conn, err := grpc.Dial("localhost:8089", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil, err
	}
	return pbUserLogin.NewUserLoginServiceClient(conn), nil
}
