package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/vsomera/scratch-api/api"
	"github.com/vsomera/scratch-api/storage"
)

func main() {
	listenAddr := flag.String("listenaddr", ":8080", "port")
	flag.Parse()

	store, err := storage.NewMySqlStore()
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewApiServer(*listenAddr, store)

	fmt.Println("server running on port: ", *listenAddr)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
