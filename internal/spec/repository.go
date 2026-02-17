package spec

import (
	"sync"

	"github.com/agure-la/api-docs/internal/errors"
	"github.com/agure-la/api-docs/internal/models"
)

// Repository handles data storage and retrieval
type Repository struct {
	cache   map[string]*models.APIDocument
	cacheMu sync.RWMutex
}

func NewRepository() *Repository {
	return &Repository{
		cache: make(map[string]*models.APIDocument),
	}
}

func (r *Repository) Save(doc *models.APIDocument) error {
	r.cacheMu.Lock()
	defer r.cacheMu.Unlock()

	r.cache[doc.Name] = doc
	return nil
}

func (r *Repository) FindByName(name string) (*models.APIDocument, error) {
	r.cacheMu.RLock()
	defer r.cacheMu.RUnlock()

	doc, exists := r.cache[name]
	if !exists {
		return nil, errors.NotFound("API")
	}

	return doc, nil
}

func (r *Repository) FindAll() []models.APIDocument {
	r.cacheMu.RLock()
	defer r.cacheMu.RUnlock()

	apis := make([]models.APIDocument, 0, len(r.cache))
	for _, doc := range r.cache {
		apis = append(apis, *doc)
	}

	return apis
}

func (r *Repository) Delete(name string) error {
	r.cacheMu.Lock()
	defer r.cacheMu.Unlock()

	if _, exists := r.cache[name]; !exists {
		return errors.NotFound("API")
	}

	delete(r.cache, name)
	return nil
}

func (r *Repository) Exists(name string) bool {
	r.cacheMu.RLock()
	defer r.cacheMu.RUnlock()

	_, exists := r.cache[name]
	return exists
}
