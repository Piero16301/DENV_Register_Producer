package models

import (
	"time"
)

type VectorRecord struct {
	Address   Address   `json:"address" validate:"required"`
	Comment   string    `json:"comment" validate:"required"`
	Datetime  time.Time `json:"datetime" validate:"required"`
	Latitude  float64   `json:"latitude" validate:"required"`
	Longitude float64   `json:"longitude" validate:"required"`
	PhotoUrl  string    `json:"photoUrl" validate:"required"`
}
