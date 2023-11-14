package middlewares


import (	
	"context"
	"net/http"
	"strings"
	"github.com/golang-jwt/jwt/v4"
	"github.com/nukkua/ra-chi/internal/app/handlers"
)

func JwtAuthentication (next http.Handler) http.Handler{
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){

		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == ""{
			http.Error(w, "Acceso denegado, No se proporciono el token", http.StatusUnauthorized)
			return 
		}

		splitToken := strings.Split(tokenHeader, "Bearer ")

		if len(splitToken) != 2{
			http.Error(w, "Error en el token de autenticacion", http.StatusUnauthorized)
			return 
			
		}
		tokenString:= splitToken[1]
		claims:= &handlers.Claims{}
		
		token , err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return handlers.JwtKey, nil
		})

		if err != nil {
			http.Error(w, "Error en el token de autenticacion"+ err.Error(), http.StatusUnauthorized)
			return 
		}

		if !token.Valid {
			http.Error(w, "Error en el token de autenticacion: token invalido", http.StatusUnauthorized)
		}
		
		ctx:= context.WithValue(r.Context(),"user", claims)
		r = r.WithContext(ctx)

		next.ServeHTTP(w,r)


	})
}
