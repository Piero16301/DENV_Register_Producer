package controllers

import (
	"DENV_Register_Producer/configs"
	"DENV_Register_Producer/models"
	"DENV_Register_Producer/responses"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var validateHomeInspection = validator.New()

func AddHomeInspectionToQueue() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")

		// Validar que el body está en formato JSON
		var homeInspection models.HomeInspection
		if err := json.NewDecoder(request.Body).Decode(&homeInspection); err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			response := responses.HomeInspectionResponse{
				Status:  http.StatusBadRequest,
				Message: "El cuerpo de la solicitud no está en formato JSON",
				Data:    err.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		// Se valida que se envíen todos los campos requeridos
		if validationErr := validateHomeInspection.Struct(&homeInspection); validationErr != nil {
			writer.WriteHeader(http.StatusBadRequest)
			response := responses.HomeInspectionResponse{
				Status:  http.StatusBadRequest,
				Message: "No se han enviado todos los campos requeridos",
				Data:    validationErr.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		// Codificar inspección de vivienda a JSON
		homeInspectionJson, err := json.Marshal(homeInspection)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			response := responses.HomeInspectionResponse{
				Status:  http.StatusInternalServerError,
				Message: "No se ha podido codificar la inspección de vivienda a JSON",
				Data:    err.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		// Enviar inspección de vivienda a la cola de Kafka
		message := &sarama.ProducerMessage{
			Topic: "register-home-inspection",
			Value: sarama.StringEncoder(homeInspectionJson),
		}
		partition, offset, err := configs.Producer.SendMessage(message)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			response := responses.HomeInspectionResponse{
				Status:  http.StatusInternalServerError,
				Message: "No se ha podido enviar la inspección de vivienda a la cola de Kafka",
				Data:    err.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		// Responder con el mensaje de éxito
		writer.WriteHeader(http.StatusCreated)
		response := responses.HomeInspectionResponse{
			Status:  http.StatusCreated,
			Message: "Se ha registrado la inspección de vivienda",
			Data:    fmt.Sprintf("Inspección de vivienda almacenada en: register-home-inspection/%d/%d", partition, offset),
		}
		_ = json.NewEncoder(writer).Encode(response)
	}
}
