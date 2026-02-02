package main

import (
	"context"
	"geo_quest/internal/api"
	"geo_quest/internal/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	generator := &service.MockQuestGenerator{}
	router := api.NewRouter(generator)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// канал для сигналов
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		log.Println("Server started on :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	<-stop // ждём Ctrl+C

	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Println("Server exited properly")
}
