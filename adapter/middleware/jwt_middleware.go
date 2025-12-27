package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/chahar4/aura/core/services"
	"github.com/golang-jwt/jwt/v5"
)

func JwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		header := w.Header().Get("Authorization")
		trimHeader := strings.TrimPrefix(header, "Bearer ")

		secretKey := os.Getenv("SECRET_KEY")
		token, err := jwt.ParseWithClaims(trimHeader, &services.CustomeClaim{}, func(t *jwt.Token) (any, error) {
			return []byte(secretKey), nil
		})
		if err != nil || !token.Valid{
			//TODO:should return an error
			return 
		}
		claims , ok := token.Claims.(*services.CustomeClaim)
		if !ok{
			//TODO:should return an error
			return 
		}
		ctx := context.WithValue(r.Context(), "userID",claims.ID)
		ctx = context.WithValue(r.Context(), "username", claims.Username)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
