package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/aiwizzard/go-http/api"
	"github.com/aiwizzard/go-http/storage"
)

func main() {
    listAddr := flag.String("listenaddr", ":3000", "the server address")
    flag.Parse()

    store := storage.NewMemoryStorage()

    server := api.NewServer(*listAddr, store)
    fmt.Println("server running on port:",*listAddr)
    
    log.Fatal(server.Start())
}
