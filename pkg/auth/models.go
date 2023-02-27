package auth

type JwtClaimNeeded struct {
	Email             string `json:"email"`
	Name              string `json:"name"`
	PreferredUsername string `json:"preferred_username"`
}

type OAuth2ConfigAuthType string

const Default OAuth2ConfigAuthType = "default"
const AzureAD OAuth2ConfigAuthType = "azure_ad"

type OAuth2Config struct {
	Type     OAuth2ConfigAuthType `json:"type" binding:"required"`
	Scope    string               `json:"scope" binding:"required"`
	Issuer   string               `json:"issuer" binding:"required"`
	ClientID string               `json:"client_id" binding:"required"`
}
