package main

import (
	"context"
	"fmt"
	"gin-example/database"
	"gin-example/pkg/config"
	"gin-example/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	config.ConfigSetup("config/settings.yml")
	database.InitDb()
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf("%s:%s", config.ApplicationConfig.Host, config.ApplicationConfig.Port),
		Handler:        router,
		ReadTimeout:    config.ApplicationConfig.ReadTimeout,
		WriteTimeout:   config.ApplicationConfig.WriterTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	defer database.CloseDB()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
