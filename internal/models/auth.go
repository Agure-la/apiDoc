package models

type AuthScheme struct {
	Type        AuthType `json:"type"`
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
}

type AuthType string

const (
	AuthAPIKey AuthType = "apiKey"
	AuthOAuth2 AuthType = "oauth2"
	AuthJWT    AuthType = "jwt"
	AuthNone   AuthType = "none"
)
