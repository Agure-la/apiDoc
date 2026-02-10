package spec

import "time"

// APIDocument represents a single API across versions.
type APIDocument struct {
	Name        string            `json:"name"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Versions    []APIVersion      `json:"versions"`
	Metadata    map[string]string `json:"metadata,omitempty"`
}

// APIVersion represents a specific API version.
type APIVersion struct {
	Version     string          `json:"version"`
	Deprecated  bool            `json:"deprecated"`
	ReleaseDate *time.Time      `json:"releaseDate,omitempty"`
	Info        VersionInfo     `json:"info"`

	Endpoints   []Endpoint      `json:"endpoints"`
	Schemas     []Schema        `json:"schemas"`
	Auth        []AuthScheme    `json:"auth"`
	Errors      []APIError      `json:"errors,omitempty"`
}

// VersionInfo contains human-facing metadata.
type VersionInfo struct {
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
}
