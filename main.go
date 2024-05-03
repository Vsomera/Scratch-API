package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/vsomera/scratch-api/api"
)

func main() {
	listenAddr := flag.String("listenaddr", ":8080", "port")
	flag.Parse()

	server := api.NewApiServer(*listenAddr)
	fmt.Println("server running on port: ", *listenAddr)
	log.Fatal(server.Start())

}
