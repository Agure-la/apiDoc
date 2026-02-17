package spec

import (
	"fmt"
	"sync"

	"github.com/agure-la/api-docs/internal/config"
	"github.com/agure-la/api-docs/internal/models"
	"github.com/agure-la/api-docs/internal/spec/loader"
	"github.com/agure-la/api-docs/internal/spec/parser"
)

type Service struct {
	config   *config.Config
	loader   *loader.FileSystemLoader
	parser   *parser.Parser
	cache    map[string]*models.APIDocument
	cacheMu  sync.RWMutex
}

func NewService(cfg *config.Config) *Service {
	return &Service{
		config: cfg,
		loader: loader.NewFileSystemLoader(),
		parser: parser.New(),
		cache:  make(map[string]*models.APIDocument),
	}
}

func (s *Service) LoadAll() error {
	s.cacheMu.Lock()
	defer s.cacheMu.Unlock()

	s.cache = make(map[string]*models.APIDocument)

	for _, source := range s.config.Specs.Sources {
		doc, err := s.loadSpec(source.Name, source.Path, source.Version)
		if err != nil {
			return fmt.Errorf("failed to load spec %s: %w", source.Name, err)
		}
		s.cache[source.Name] = doc
	}

	return nil
}

func (s *Service) GetAPIs() []models.APIDocument {
	s.cacheMu.RLock()
	defer s.cacheMu.RUnlock()

	apis := make([]models.APIDocument, 0, len(s.cache))
	for _, doc := range s.cache {
		apis = append(apis, *doc)
	}

	return apis
}

func (s *Service) GetAPI(name string) (*models.APIDocument, error) {
	s.cacheMu.RLock()
	defer s.cacheMu.RUnlock()

	doc, exists := s.cache[name]
	if !exists {
		return nil, fmt.Errorf("API not found: %s", name)
	}

	return doc, nil
}

func (s *Service) GetAPIVersion(name, version string) (*models.APIVersion, error) {
	doc, err := s.GetAPI(name)
	if err != nil {
		return nil, err
	}

	for _, v := range doc.Versions {
		if v.Version == version {
			return &v, nil
		}
	}

	return nil, fmt.Errorf("version %s not found for API %s", version, name)
}

func (s *Service) GetAPIVersions(name string) ([]models.APIVersion, error) {
	doc, err := s.GetAPI(name)
	if err != nil {
		return nil, err
	}

	return doc.Versions, nil
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
