package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/vsomera/scratch-api/api"
	"github.com/vsomera/scratch-api/storage"
)

func main() {
	godotenv.Load()
	listenAddr := flag.String("listenaddr", ":8080", "port")
	flag.Parse()

	// start new sql connection
	store, err := storage.NewMySqlStore()
	if err != nil {
		log.Fatal(err)
	}

	// create an API se
	server := api.NewApiServer(*listenAddr, store)

	fmt.Println("server running on port: ", *listenAddr)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
