package helpers

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
)

// Helper function to generate JWT token
func GenerateJWTToken(userID string, userRoles []string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "guaz-webhooks",
		"sub": userID,
		"https://hasura.io/jwt/claims": map[string]interface{}{
			"x-hasura-default-role":  userRoles[0],
			"x-hasura-allowed-roles": userRoles,
			"x-hasura-user-id":       userID,
		},
	})
	return token.SignedString([]byte(os.Getenv("JWT_PRIVATE_KEY")))
}
