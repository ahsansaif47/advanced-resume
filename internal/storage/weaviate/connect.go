package weaviate

import (
	"context"
	"log"
	"sync"

	"github.com/weaviate/weaviate-go-client/v5/weaviate"
)

type Database struct {
	Client *weaviate.Client
}

var dbInstance *Database
var once sync.Once

func GetDatabaseConnection() *Database {
	once.Do(func() {
		dbInstance = &Database{
			Client: ConnectWeaviate(),
		}
	})
	return dbInstance
}

func ConnectWeaviate() *weaviate.Client {
	cfg := weaviate.Config{
		Host:   "localhost:7070",
		Scheme: "http",
	}

	client, err := weaviate.NewClient(cfg)
	if err != nil {
		log.Fatalf("error creating weaviate client: %v", err)
	}

	live, err := client.Misc().LiveChecker().Do(context.Background())
	if err != nil {
		log.Fatalf("error checking live status of weaviate: %v", err)
	}

	if live {
		log.Println("Weviate is Live :")
	}

	return client
}
