package controllers

import "github.com/ahsansaif47/advanced-resume/internal/storage/weaviate"

type IWeaviateService interface {
	InsertData(resumeData map[string]any) (string, error)
	BatchInsert(batchResume []map[string]any) error
	VectorSearch(query string) (any, error)
}

type WeaviateService struct {
	repo weaviate.IWeaviateRepository
}

func NewWeaviateService(repo weaviate.IWeaviateRepository) IWeaviateService {
	return &WeaviateService{
		repo: repo,
	}
}

func (s *WeaviateService) InsertData(resumeData map[string]any) (string, error) {
	return s.repo.InsertData("resume", resumeData)
}

func (s *WeaviateService) BatchInsert(batchResume []map[string]any) error {
	return s.repo.BatchInsert("resume", batchResume)

}

func (s *WeaviateService) VectorSearch(query string) (any, error) {
	return s.repo.VectorSearch("resume", query)
}
