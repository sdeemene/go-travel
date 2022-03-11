package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/stdeemene/go-travel/config"
	"github.com/stdeemene/go-travel/middleware"
	"github.com/stdeemene/go-travel/routers"
	"github.com/stdeemene/go-travel/security"
)

func main() {
	config := config.GetConfiguration()
	router := routers.GetRouters()

	jwtService := security.TokenConfig{Config: &config}
	fmt.Println("jwtService ", jwtService)

	router.Use(middleware.Cors, middleware.LoggingUri, jwtService.ProtectApi)

	srv := &http.Server{
		Handler:      middleware.Headers(router),
		Addr:         config.Server.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Printf("Application is running on Port %v....\n", config.Server.Port)
	log.Fatal(srv.ListenAndServe())

}
