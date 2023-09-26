package middleware

import (
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// verify auth of user by token
		userRole := getUserRoleFromToken(ctx) // get user role
		// verify if user have the role correct
		if !hasRequiredRole(userRole, roles) {
			ctx.IndentedJSON(http.StatusForbidden, gin.H{"error": "Unauthorized access"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

// get user by token
func getUserRoleFromToken(ctx *gin.Context) string {
	// get token by authorization header
	tokenString := ctx.Request.Header.Get("Authorization")
	// parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		return "ROLE_USER"
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// get the roles from the user by claims from token
		if roles, ok := claims["roles"].([]interface{}); ok {
			roleString := ""
			for i, role := range roles {
				if i > 0 {
					roleString += ","
				}
				roleString += role.(string)
			}
			return roleString
		}
	}
	return "ROLE_USER"
}

func hasRequiredRole(userRole string, requiredRoles []string) bool {
	for _, role := range requiredRoles {

		if role == userRole {
			return true
		}
	}
	return false
}
