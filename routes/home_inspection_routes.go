package routes

import (
	"DENV_Register_Producer/controllers"
	"github.com/go-chi/chi/v5"
)

type HomeInspectionResource struct{}

// Routes Rutas para inspección de viviendas (home inspection)
func (hir HomeInspectionResource) Routes() chi.Router {
	router := chi.NewRouter()

	// Enviar inspección de vivienda a la cola de Kafka
	router.Method("POST", "/", controllers.AddHomeInspectionToQueue())

	return router
}
