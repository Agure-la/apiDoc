package loader

import (
	"fmt"
	"os"

	"github.com/getkin/kin-openapi/openapi3"
)

// FileSystemLoader loads OpenAPI specs from disk.
type FileSystemLoader struct {
	loader *openapi3.Loader
}

func NewFileSystemLoader() *FileSystemLoader {
	return &FileSystemLoader{
		loader: &openapi3.Loader{
			IsExternalRefsAllowed: true,
		},
	}
}

func (l *FileSystemLoader) Load(path string) (*openapi3.T, error) {
	if _, err := os.Stat(path); err != nil {
		return nil, fmt.Errorf("spec not found: %w", err)
	}

	spec, err := l.loader.LoadFromFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to load spec: %w", err)
	}

	if err := spec.Validate(l.loader.Context); err != nil {
		return nil, fmt.Errorf("invalid OpenAPI spec: %w", err)
	}

	return spec, nil
}
