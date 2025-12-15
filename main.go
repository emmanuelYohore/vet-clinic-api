package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/emmanuelYohore/vet-clinic-api/config"
	_ "github.com/emmanuelYohore/vet-clinic-api/docs"
	"github.com/emmanuelYohore/vet-clinic-api/pkg/authentification"
	"github.com/emmanuelYohore/vet-clinic-api/pkg/cat"
	"github.com/emmanuelYohore/vet-clinic-api/pkg/treatment"
	"github.com/emmanuelYohore/vet-clinic-api/pkg/user"
	"github.com/emmanuelYohore/vet-clinic-api/pkg/visit"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Routes(configuration *config.Config) *chi.Mux {
	router := chi.NewRouter()

	router.Mount("/login", authentification.Routes(configuration))

	router.Group(func(r chi.Router) {
		r.Use(authentification.AuthMiddleware("your_secret_key"))
		catRoutes := cat.Routes(configuration)
		r.Group(func(cr chi.Router) {
			cr.Use(authentification.RequireRole("admin", "user"))
			cr.Get("/api/v1/cats", catRoutes.ServeHTTP)
			cr.Get("/api/v1/cats/{id}", catRoutes.ServeHTTP)
			cr.Get("/api/v1/cats/{id}/history", catRoutes.ServeHTTP)
		})

		r.Group(func(cr chi.Router) {
			cr.Use(authentification.RequireRole("admin"))
			cr.Post("/api/v1/cats", catRoutes.ServeHTTP)
			cr.Put("/api/v1/cats/{id}", catRoutes.ServeHTTP)
			cr.Delete("/api/v1/cats/{id}", catRoutes.ServeHTTP)
		})

		visitRoutes := visit.Routes(configuration)
		r.Group(func(vr chi.Router) {
			vr.Use(authentification.RequireRole("admin", "user"))
			vr.Get("/api/v1/visits", visitRoutes.ServeHTTP)
			vr.Get("/api/v1/visits/{id}", visitRoutes.ServeHTTP)
		})

		r.Group(func(vr chi.Router) {
			vr.Use(authentification.RequireRole("admin"))
			vr.Post("/api/v1/visits", visitRoutes.ServeHTTP)
			vr.Put("/api/v1/visits/{id}", visitRoutes.ServeHTTP)
			vr.Delete("/api/v1/visits/{id}", visitRoutes.ServeHTTP)
		})

		treatmentRoutes := treatment.Routes(configuration)
		r.Group(func(tr chi.Router) {
			tr.Use(authentification.RequireRole("admin", "user"))
			tr.Get("/api/v1/treatments", treatmentRoutes.ServeHTTP)
			tr.Get("/api/v1/treatments/{id}", treatmentRoutes.ServeHTTP)
		})

		r.Group(func(tr chi.Router) {
			tr.Use(authentification.RequireRole("admin"))
			tr.Post("/api/v1/treatments", treatmentRoutes.ServeHTTP)
			tr.Put("/api/v1/treatments/{id}", treatmentRoutes.ServeHTTP)
			tr.Delete("/api/v1/treatments/{id}", treatmentRoutes.ServeHTTP)
		})

		r.Group(func(ur chi.Router) {
			ur.Use(authentification.RequireRole("admin"))
			ur.Mount("/api/v1/users", user.Routes(configuration))
		})

		r.Get("/protected", func(w http.ResponseWriter, req *http.Request) {
			userEmail := authentification.GetUserFromContext(req.Context())
			userRole := authentification.GetRoleFromContext(req.Context())
			w.Write([]byte(fmt.Sprintf("Welcome, %s! Your role: %s", userEmail, userRole)))
		})
	})

	router.Get("/swagger", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/index.html", http.StatusMovedPermanently)
	})

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))
	return router

}

func main() {
	configuration, err := config.New()
	if err != nil {
		log.Panicln("Configuration error:", err)
	}

	router := Routes(configuration)

	log.Println("Serving on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
