package weaviate

import (
	"context"

	"github.com/weaviate/weaviate-go-client/v5/weaviate"
	"github.com/weaviate/weaviate/entities/models"
)

type IWeaviateRepository interface {
	AddNewResume(className string, props map[string]any) (string, error)
	BatchAddResume(className string, items []map[string]any) error
	VectorSearch(className, query string) (any, error)
}

type WeaviateRepository struct {
	ctx    context.Context
	Client *weaviate.Client
}

func NewWeviateRepository(ctx context.Context, client *weaviate.Client) IWeaviateRepository {
	return &WeaviateRepository{
		ctx:    ctx,
		Client: client,
	}
}

func (r *WeaviateRepository) AddNewResume(className string, props map[string]any) (string, error) {
	model, err := r.Client.Data().Creator().WithClassName(className).WithProperties(props).Do(r.ctx)
	if err != nil {
		return "", err
	}

	return model.Object.ID.String(), nil
}

func (r *WeaviateRepository) BatchAddResume(className string, items []map[string]any) error {

	batch := r.Client.Batch().ObjectsBatcher()
	for _, item := range items {
		batch = batch.WithObjects(&models.Object{
			Class:      className,
			Properties: item,
		})
	}

	_, err := batch.Do(r.ctx)
	return err
}

func (r *WeaviateRepository) VectorSearch(className, query string) (any, error) {
	response, err := r.Client.GraphQL().Get().
		WithClassName("").
		WithNearText(r.Client.GraphQL().NearTextArgBuilder().WithConcepts([]string{query})).
		WithLimit(2).
		Do(context.Background())

	return response, err
}
