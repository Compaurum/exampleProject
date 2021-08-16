package main

import (
	"example/pkg/api"
	memory_storage "example/pkg/storage/memory-storage"
	"log"
	"net/http"
	"time"
)

func main() {
	storage := memory_storage.MemoryStorage{}

	server := http.Server{
		Addr:              ":8000",
		Handler:           api.NewHandler(storage),
		ReadHeaderTimeout: time.Second * 30,
		ReadTimeout:       time.Second * 60,
		WriteTimeout:      time.Second * 60,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
