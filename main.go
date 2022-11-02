package main

import (
	"fmt"
	"gin-example/configs"
	_ "gin-example/init"
	"gin-example/internal/routes"
	"net/http"
	"time"
)

func main() {
	HostPort := fmt.Sprintf(":%d", configs.Cfg.Server.HostPort)
	initRouter := routes.InitRoutes()
	server := &http.Server{
		Addr:         HostPort,
		Handler:      initRouter,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	_ = server.ListenAndServe()
}
