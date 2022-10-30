package main

import (
	"fmt"
	"gin-example/conf"
	"gin-example/routes"
	"net/http"
	"time"
)

func main() {
	HostPort := fmt.Sprintf(":%d", conf.Cfg.Server.HostPort)
	initRouter := routes.InitRoutes()
	server := &http.Server{
		Addr:         HostPort,
		Handler:      initRouter,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	_ = server.ListenAndServe()
	//if err != nil {
	//	log.Printf("Server err: %v", err)
	//}

}
