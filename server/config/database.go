package config


import (
	"server/ent"
	"context"
	"fmt"
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

	fmt.Println(os.Getenv("DB_USERNAME"))
	dbconf := fmt.Sprintf("%s:%s@tcp(%s)/%s?tls=true",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		 os.Getenv("DB_DATABASE"))

	client, err := ent.Open("mysql", dbconf)
    if err != nil {
        log.Fatalf("failed opening connection to mysql: %v", err)
    }

    // Run the auto migration tool.
    if err := client.Schema.Create(context.Background()); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }


	return client, err
}