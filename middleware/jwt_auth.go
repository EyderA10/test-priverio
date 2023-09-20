package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// verify the token validation
func MiddlewareJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get the token by Authorization header
		tokenString := ctx.Request.Header.Get("Authorization")

		// verify Token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// verify the sign method is correct
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("sign method not valid")
			}

			// return the secret key that was used to get the token
			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		if err != nil {
			// invalid token returns status code 401 Unauthorized
			ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
			ctx.Abort() // stop the execution from request
			return
		}

		// get claims by token valid
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// save the claims into the context
			ctx.Set("roles", claims["roles"])
			ctx.Set("username", claims["username"])
			ctx.Next()
		} else {
			// invalid token returns status code 401 Unauthorized
			ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
			ctx.Abort()
		}
	}
}
