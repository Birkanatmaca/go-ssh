package main

import (
	"context"
	"fmt"
	"log"

	"github.com/birkan-is/go-ssh/internal/config"
	"github.com/birkan-is/go-ssh/internal/sshclient"
)

func main() {
	cfg := config.Parse()

	if err := cfg.Validate(); err != nil {
		log.Fatal(err)
	}

	client, err := sshclient.Dial(context.Background(), cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	fmt.Printf("ssh connected")
}
