package routes

import (
	"DENV_Register_Producer/controllers"
	"github.com/go-chi/chi/v5"
)

type VectorRecordResource struct{}

// Routes Rutas para registro de vectores (vector record)
func (vrr VectorRecordResource) Routes() chi.Router {
	router := chi.NewRouter()

	// Enviar registro de vector a la cola de Kafka
	router.Method("POST", "/", controllers.AddVectorRecordToQueue())

	return router
}
