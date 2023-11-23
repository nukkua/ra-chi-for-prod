package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/nukkua/ra-chi/internal/mvc/database"
	"github.com/nukkua/ra-chi/internal/mvc/handlers"
	"github.com/nukkua/ra-chi/internal/mvc/middlewares"
)

func SetupRouter() *chi.Mux {
	db := database.SetupDatabase()

	r := chi.NewRouter()

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // o []string{"*"} para desarrollo
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{},
		AllowCredentials: true,
		MaxAge:           300,
	})

	// Usar el middleware CORS en todas las rutas
	r.Use(corsMiddleware.Handler)

	// Otros middlewares y configuraciones
	r.Use(middleware.Logger)

	// Tus rutas
	r.Post("/register", handlers.CreateUser(db))
	r.Post("/login", handlers.LoginUser(db))

	r.Group(func(r chi.Router) {
		r.Use(middlewares.JwtAuthentication)
		r.Get("/users", handlers.GetUsers(db))
	})

	return r
}
