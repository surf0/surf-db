package config


import (
	"server/ent"
	"context"
	"os"
	"log"
)


var (
	client *ent.Client
)

func GetClient() (*ent.Client) {
	return client
}


func SetClient(newClient *ent.Client) {
	client = newClient
}

func NewClient() (*ent.Client, error) {
	// Create an ent.Client.
	client, err := ent.Open("mysql", os.Getenv("DSN"))

    if err != nil {
        log.Fatalf("failed opening connection to mysql: %v", err)
    }

    // Run the auto migration tool.
    if err := client.Schema.Create(context.Background()); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }


	return client, err
}