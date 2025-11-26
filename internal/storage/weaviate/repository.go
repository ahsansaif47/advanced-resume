package weaviate

import (
	"context"

	"github.com/weaviate/weaviate-go-client/v5/weaviate"
)

type IWeaviateRepository interface {
	InsertData(ctx context.Context, className string, props map[string]any)
	BatchInsert()
	VectorSearch()
}

type WeaviateRepository struct {
	Client *weaviate.Client
}

func NewWeviateRepository(client *weaviate.Client) IWeaviateRepository {
	return &WeaviateRepository{
		Client: client,
	}
}

func (r *WeaviateRepository) InsertData(ctx context.Context, className string, props map[string]any) {

}

func (r *WeaviateRepository) BatchInsert() {

}

func (r *WeaviateRepository) VectorSearch() {

}
