package middleware

import (
	"net/http"

	"github.com/rmsubekti/sporagium/dto"

	"github.com/gorilla/context"
)

func JwtAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		principal := Principal{
			Token: r.Header.Get("Authorization"),
		}
		if err := principal.Parse(); err != nil {
			dto.JSON(w, http.StatusUnauthorized, err.Error())
			return
		}
		context.Set(r, "principal", principal)

		next.ServeHTTP(w, r)
	})
}
