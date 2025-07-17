package main

import (
	"fmt"
	"net/http"
	"time"
)

func getHome() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			backgroundWork()
		}()
		writer.Write([]byte("home!"))
	})
}

func backgroundWork() {
	fmt.Println("background work started!")
	time.Sleep(3 * time.Second)
	fmt.Println("background work complete!")
}
