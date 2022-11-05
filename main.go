package main

import (
	"fmt"
	"gin-example/init/config"

	"gin-example/internal/routes"
	//"gin-example/configs"
	_ "gin-example/init"
	"net/http"
	"time"
)

func main() {
	HostPort := fmt.Sprintf(":%d", config.Cfg.Server.HostPort)
	initRouter := routes.InitRoutes()
	server := &http.Server{
		Addr:         HostPort,
		Handler:      initRouter,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	_ = server.ListenAndServe()
}
