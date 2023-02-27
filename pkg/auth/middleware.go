package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

var (
	ctx          = context.Background()
	oauth2Config *oauth2.Config
	verifier     *oidc.IDTokenVerifier
)

func init() {
	config := getOAuthConfig()

	provider, err := oidc.NewProvider(ctx, config.Issuer)
	if err != nil {
		log.Fatalln(err)
	}

	oauth2Config = &oauth2.Config{
		ClientID: config.ClientID,
		// RedirectURL:  redirectURL,

		// Discovery returns the OAuth2 endpoints.
		Endpoint: provider.Endpoint(),

		// "openid" is a required scope for OpenID Connect flows.
		Scopes: strings.Split(config.Scope, " "),
	}

	verifier = provider.Verifier(&oidc.Config{ClientID: config.ClientID})
}

func getOAuthConfig() OAuth2Config {
	defaultScopes := "openid offline_access profile email"
	draftConfig := OAuth2Config{
		Type:     OAuth2ConfigAuthType(os.Getenv("OAUTH2_TYPE")),
		Issuer:   os.Getenv("OAUTH2_ISSUER"),
		ClientID: os.Getenv("OAUTH2_CLIENT_ID"),
	}

	switch draftConfig.Type {
	case Default:
		draftConfig.Scope = defaultScopes
	case AzureAD:
		draftConfig.Scope = fmt.Sprintf("api://%s/default %s", draftConfig.ClientID, defaultScopes)
	default:
		log.Fatalln("no valid OAUTH_TYPE")
	}

	return draftConfig
}

func AuthEndpoints(r *gin.Engine) *gin.RouterGroup {
	oauth2 := r.Group("oauth2")
	oauth2.GET("/config", getConfig)

	return oauth2
}

// @Summary Get the oauth2 configuration
// @Schemes
// @Description Returns the oauth2 configuration
// @Tags oauth2
// @Produce json
// @Success 200 {object} OAuth2Config
// @Router /oauth2/config [get]
func getConfig(c *gin.Context) {
	c.JSON(http.StatusOK, getOAuthConfig())
}

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := ValidateToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"Unauthorized": "Authentication required"})
			log.Println(err)
			c.Abort()
			return
		}
		c.Next()
	}
}
