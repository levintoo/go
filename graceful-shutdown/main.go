package main

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var wg sync.WaitGroup

func main() {
	srv := http.Server{
		Addr:    ":8080",
		Handler: routes(),
	}

	go func() {
		fmt.Println("Listening on port 8080")
		if err := srv.ListenAndServe(); err != nil {
			fmt.Printf("Stopped listening: %v\n", err)
		}
	}()

	shutdown, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()
	<-shutdown.Done()

	fmt.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("Error shutting down server: %v\n", err)
	}
	fmt.Println("Server shut down")

	fmt.Println("Waiting for goroutines work to complete...")
	wg.Wait()
	fmt.Println("Background goroutines complete")
}
