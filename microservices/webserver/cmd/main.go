package main

import (
	"log"
	"net/http"

	"github.com/krishnakashyap0704/microservices/webserver/internal/handlers"
)

func main() {
	

	
	http.HandleFunc("/squareoff", handlers.SquareOffHandler)
	http.HandleFunc("/tradelist", handlers.TradeListHandler)
	http.HandleFunc("/login", handlers.LoginHandler)

	log.Println("Starting web server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("could not start server: %s", err)
	}
}


