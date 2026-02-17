package spec

import (
	"fmt"

	"github.com/agure-la/api-docs/internal/config"
	"github.com/agure-la/api-docs/internal/errors"
	"github.com/agure-la/api-docs/internal/models"
	"github.com/agure-la/api-docs/internal/spec/loader"
	"github.com/agure-la/api-docs/internal/spec/parser"
)

// Service handles business logic for API documentation
type Service struct {
	repository *Repository
	loader     *loader.FileSystemLoader
	parser     *parser.Parser
	config     *config.Config
}

func NewService(cfg *config.Config) *Service {
	return &Service{
		repository: NewRepository(),
		loader:     loader.NewFileSystemLoader(),
		parser:     parser.New(),
		config:     cfg,
	}
}

func (s *Service) LoadAll() error {
	for _, source := range s.config.Specs.Sources {
		doc, err := s.loadSpec(source.Name, source.Path, source.Version)
		if err != nil {
			return fmt.Errorf("failed to load spec %s: %w", source.Name, err)
		}
		if err := s.repository.Save(doc); err != nil {
			return fmt.Errorf("failed to save spec %s: %w", source.Name, err)
		}
	}
	return nil
}

func (s *Service) GetAPIs() []models.APIDocument {
	return s.repository.FindAll()
}

func (s *Service) GetAPI(name string) (*models.APIDocument, error) {
	return s.repository.FindByName(name)
}

func (s *Service) GetAPIVersion(name, version string) (*models.APIVersion, error) {
	doc, err := s.repository.FindByName(name)
	if err != nil {
		return nil, err
	}

	for _, v := range doc.Versions {
		if v.Version == version {
			return &v, nil
		}
	}

	return nil, errors.NotFound(fmt.Sprintf("version %s for API %s", version, name))
}

func (s *Service) GetAPIVersions(name string) ([]models.APIVersion, error) {
	doc, err := s.repository.FindByName(name)
	if err != nil {
		return nil, err
	}

	return doc.Versions, nil
}

func (s *Service) CreateAPI(req *models.CreateAPIRequest) (*models.APIDocument, error) {
	if s.repository.Exists(req.Name) {
		return nil, errors.Conflict("API", fmt.Sprintf("API with name '%s' already exists", req.Name))
	}

	doc := &models.APIDocument{
		Name:        req.Name,
		Title:       req.Title,
		Description: req.Description,
		Versions:    []models.APIVersion{},
		Metadata:    req.Metadata,
	}

	if err := s.repository.Save(doc); err != nil {
		return nil, errors.InternalError(fmt.Sprintf("failed to save API: %v", err))
	}

	return doc, nil
}

func (s *Service) UpdateAPI(name string, req *models.UpdateAPIRequest) (*models.APIDocument, error) {
	doc, err := s.repository.FindByName(name)
	if err != nil {
		return nil, err
	}

	if req.Title != "" {
		doc.Title = req.Title
	}
	if req.Description != "" {
		doc.Description = req.Description
	}
	if req.Metadata != nil {
		doc.Metadata = req.Metadata
	}

	if err := s.repository.Save(doc); err != nil {
		return nil, errors.InternalError(fmt.Sprintf("failed to update API: %v", err))
	}

	return doc, nil
}

func (s *Service) DeleteAPI(name string) error {
	return s.repository.Delete(name)
}

func (s *Service) loadSpec(name, path, version string) (*models.APIDocument, error) {
	openapiDoc, err := s.loader.Load(path)
	if err != nil {
		return nil, err
	}

	apiVersion, err := s.parser.Parse(name, version, openapiDoc)
	if err != nil {
		return nil, err
	}

	return &models.APIDocument{
		Name:        name,
		Title:       apiVersion.Info.Title,
		Description: apiVersion.Info.Description,
		Versions:    []models.APIVersion{*apiVersion},
		Metadata:    make(map[string]string),
	}, nil
}
