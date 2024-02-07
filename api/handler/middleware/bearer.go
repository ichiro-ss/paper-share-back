package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func CheckToken(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Vary", "Authorization")

		// get Authorization Header(Bearer ~)
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			// could set an anonymous user
			http.Error(w, "empty auth header", http.StatusUnauthorized)

		}

		// ["Bearer", "<Authorization>"]
		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 {
			http.Error(w, "invalid auth header", http.StatusUnauthorized)
		}

		// use Bearer auth
		if headerParts[0] != "Bearer" {
			http.Error(w, "unauthorized - no bearer", http.StatusUnauthorized)
		}

		// get ~(JWT token)
		tokenString := headerParts[1]
		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("ACCESS_SECRET_KEY"), nil
		})
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Printf("user_id: %v\n", int64(claims["user_id"].(float64)))
			h.ServeHTTP(w, r)
		} else {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
	}
	return http.HandlerFunc(fn)
}
