package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/krishnakashyap0704/microservices/webserver/internal/grpcclient"

	// pbSquareOff "github.com/krishnakashyap0704/microservices/squareoff/internal/generated"
	// pbTradeList "github.com/krishnakashyap0704/microservices/tradelist/internal/generated"
	// pbUserLogin "github.com/krishnakashyap0704/microservices/userlogin/internal/generated"

	pbSquareOff "github.com/krishnakashyap0704/microservices/squareoff/generated"
	pbTradeList "github.com/krishnakashyap0704/microservices/tradelist/generated"
	pbUserLogin "github.com/krishnakashyap0704/microservices/userlogin/generated"
	
)




func SquareOffHandler(w http.ResponseWriter, r *http.Request) {
	client, err := grpcclient.NewSquareOffClient()
	if err != nil {
		http.Error(w, "could not create client", http.StatusInternalServerError)
		return
	}

	req := &pbSquareOff.SquareOffRequest{}
	// Parse request and fill req

	resp, err := client.SquareOffOrder(context.Background(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func TradeListHandler(w http.ResponseWriter, r *http.Request) {
	client, err := grpcclient.NewTradeListClient()
	if err != nil {
		http.Error(w, "could not create client", http.StatusInternalServerError)
		return
	}

	req := &pbTradeList.TradeListRequest{}
	// Parse request and fill req

	resp, err := client.TradeListOrder(context.Background(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	client, err := grpcclient.NewUserLoginClient()
	if err != nil {
		http.Error(w, "could not create client", http.StatusInternalServerError)
		return
	}

	req := &pbUserLogin.LoginRequest{}
	// Parse request and fill req

	resp, err := client.Login(context.Background(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}