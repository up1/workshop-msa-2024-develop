package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hashicorp/vault-client-go"
)

func initVaultClient() *vault.Client {
	// prepare a client with the given base address
	client, err := vault.New(
		vault.WithAddress("http://0.0.0.0:8200"),
		vault.WithRequestTimeout(10*time.Second),
	)

	if err != nil {
		log.Fatal(err)
	}

	return client
}

func main() {
	readDataWithToken()
}

func readDataWithToken() {
	cl := initVaultClient()
	err := cl.SetToken(os.Getenv("USER_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := cl.Read(context.Background(), "cubbyhole/demo/secret")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("username: %v\n", resp.Data["username"])
	fmt.Printf("password: %v\n", resp.Data["password"])
}
