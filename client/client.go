package client

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func ClientCreate() *ethclient.Client {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	return client
}
