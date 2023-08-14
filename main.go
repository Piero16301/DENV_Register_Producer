package main

import (
	"DENV_Register_Producer/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	// Enrutador de endpoints
	router := chi.NewRouter()

	// Middleware para CORS
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// Rutas para agregar de inspecciones de viviendas y registros de vectores
	router.Mount("/home-inspections", routes.HomeInspectionResource{}.Routes())
	router.Mount("/vector-records", routes.VectorRecordResource{}.Routes())

	// Habilitar CORS para método POST
	corsRouter := cors.New(cors.Options{
		AllowedMethods: []string{"POST"},
	}).Handler(router)

	// Iniciar servidor en el puerto 8080
	log.Fatal(http.ListenAndServe(":8080", corsRouter))
}
