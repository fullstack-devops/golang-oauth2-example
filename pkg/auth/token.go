package auth

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	ErrTokenNotPresentInHeader error = errors.New("token in header 'Authorization' not found")
)

func ValidateToken(c *gin.Context) (claims *JwtClaimNeeded, err error) {
	tokenString, err := getTokenFromRequest(c)
	if err != nil {
		return nil, err
	}

	// Parse and verify ID Token payload.
	idToken, err := verifier.Verify(ctx, tokenString)
	if err != nil {
		return nil, err
	}

	if err := idToken.Claims(&claims); err != nil {
		return nil, err
	}

	return
}

func getTokenFromRequest(c *gin.Context) (string, error) {
	bearerToken := c.Request.Header.Get("Authorization")

	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1], nil
	}
	return "", ErrTokenNotPresentInHeader
}
