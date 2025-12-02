package controllers

import "github.com/ahsansaif47/advanced-resume/internal/storage/weaviate"

type IWeaviateService interface {
	AddNewResume(resumeData map[string]any) (string, error)
	BatchAddResume(batchResume []map[string]any) error
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

func (s *WeaviateService) AddNewResume(resumeData map[string]any) (string, error) {
	return s.repo.AddNewResume("resume", resumeData)
}

func (s *WeaviateService) BatchAddResume(batchResume []map[string]any) error {
	return s.repo.BatchAddResume("resume", batchResume)
}

func (s *WeaviateService) VectorSearch(query string) (any, error) {
	return s.repo.VectorSearch("resume", query)
}
