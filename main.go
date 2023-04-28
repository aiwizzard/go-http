package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/aiwizzard/go-http/api"
)

func main() {
    listAddr := flag.String("listenaddr", ":3000", "the server address")
    flag.Parse()

    server := api.NewServer(*listAddr)
    fmt.Println("server running on port:",*listAddr)
    log.Fatal(server.Start())
}
