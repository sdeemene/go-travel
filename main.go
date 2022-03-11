package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/stdeemene/go-travel/config"
	"github.com/stdeemene/go-travel/routers"
)

func main() {
	config := config.GetConfiguration()

	routes := routers.GetRouters()
	srv := &http.Server{
		Handler:      routes,
		Addr:         config.Server.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Printf("Starting server on port %v....", config.Server.Port)
	log.Fatal(srv.ListenAndServe())
	fmt.Printf("Listening on port %v....", config.Server.Port)
}
