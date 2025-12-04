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
		r.Mount("/api/v1/cats", cat.Routes(configuration))
		r.Mount("/api/v1/visits", visit.Routes(configuration))
		r.Mount("/api/v1/treatments", treatment.Routes(configuration))
		r.Mount("/api/v1/users", user.Routes(configuration))
		r.Get("/protected", func(w http.ResponseWriter, req *http.Request) {
			user := authentification.GetUserFromContext(req.Context())
			w.Write([]byte(fmt.Sprintf("Welcome, %s!", user)))
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
