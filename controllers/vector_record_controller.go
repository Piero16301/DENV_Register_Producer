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

var validateVectorRecord = validator.New()

func AddVectorRecordToQueue() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")

		// Validar que el body está en formato JSON
		var vectorRecord models.VectorRecord
		if err := json.NewDecoder(request.Body).Decode(&vectorRecord); err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			response := responses.VectorRecordResponse{
				Status:  http.StatusBadRequest,
				Message: "El cuerpo de la solicitud debe estar en formato JSON",
				Data:    err.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		// Se valida que el body tenga los campos requeridos
		if validationErr := validateVectorRecord.Struct(vectorRecord); validationErr != nil {
			writer.WriteHeader(http.StatusBadRequest)
			response := responses.VectorRecordResponse{
				Status:  http.StatusBadRequest,
				Message: "No se han enviado todos los campos requeridos",
				Data:    validationErr.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		// Codificar registro de vector a JSON
		vectorRecordJson, err := json.Marshal(vectorRecord)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			response := responses.VectorRecordResponse{
				Status:  http.StatusInternalServerError,
				Message: "No se ha podido codificar el registro de vector a JSON",
				Data:    err.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		// Enviar registro de vector a la cola de Kafka
		message := &sarama.ProducerMessage{
			Topic: "register-vector-record",
			Value: sarama.StringEncoder(vectorRecordJson),
		}
		partition, offset, err := configs.Producer.SendMessage(message)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			response := responses.VectorRecordResponse{
				Status:  http.StatusInternalServerError,
				Message: "No se ha podido enviar el registro de vector a la cola de Kafka",
				Data:    err.Error(),
			}
			_ = json.NewEncoder(writer).Encode(response)
			return
		}

		// Responder con el mensaje de éxito
		writer.WriteHeader(http.StatusCreated)
		response := responses.VectorRecordResponse{
			Status:  http.StatusCreated,
			Message: "Se ha registrado el registro de vector",
			Data:    fmt.Sprintf("Registro de vector almacenado en: register-vector-record/%d/%d", partition, offset),
		}
		_ = json.NewEncoder(writer).Encode(response)
	}
}
