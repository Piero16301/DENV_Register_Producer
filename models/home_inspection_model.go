package models

import (
	"time"
)

type HomeInspection struct {
	Address           Address        `json:"address" validate:"required"`
	Comment           string         `json:"comment" validate:"required"`
	Datetime          time.Time      `json:"datetime" validate:"required"`
	Dni               string         `json:"dni" validate:"required"`
	Latitude          float64        `json:"latitude" validate:"required"`
	Longitude         float64        `json:"longitude" validate:"required"`
	PhotoUrl          string         `json:"photoUrl" validate:"required"`
	NumberInhabitants int32          `json:"numberInhabitants" validate:"required"`
	TypeContainer     TypeContainer  `json:"typeContainer" validate:"required"`
	HomeCondition     HomeCondition  `json:"homeCondition" validate:"required"`
	TotalContainer    TotalContainer `json:"totalContainer" validate:"required"`
	AegyptiFocus      AegyptiFocus   `json:"aegyptiFocus" validate:"required"`
	Larvicide         float32        `json:"larvicide" validate:"required"`
}

type Container struct {
	I int32 `json:"i" validate:"required"`
	P int32 `json:"p" validate:"required"`
	T int32 `json:"t" validate:"required"`
}

type TypeContainer struct {
	ElevatedTank   Container `json:"elevatedTank" validate:"required"`
	LowTank        Container `json:"lowTank" validate:"required"`
	CylinderBarrel Container `json:"cylinderBarrel" validate:"required"`
	BucketTub      Container `json:"bucketTub" validate:"required"`
	Tire           Container `json:"tire" validate:"required"`
	Flower         Container `json:"flower" validate:"required"`
	Useless        Container `json:"useless" validate:"required"`
	Others         Container `json:"others" validate:"required"`
}

type HomeCondition struct {
	InspectedHome     int32 `json:"inspectedHome" validate:"required"`
	ReluctantDwelling int32 `json:"reluctantDwelling" validate:"required"`
	ClosedHouse       int32 `json:"closedHouse" validate:"required"`
	UninhabitedHouse  int32 `json:"uninhabitedHouse" validate:"required"`
	HousingSpotlights int32 `json:"housingSpotlights" validate:"required"`
	TreatedHousing    int32 `json:"treatedHousing" validate:"required"`
}

type TotalContainer struct {
	InspectedContainers  int32 `json:"inspectedContainers" validate:"required"`
	ContainersSpotlights int32 `json:"containersSpotlights" validate:"required"`
	TreatedContainers    int32 `json:"treatedContainers" validate:"required"`
	DestroyedContainers  int32 `json:"destroyedContainers" validate:"required"`
}

type AegyptiFocus struct {
	Larvae int `json:"larvae" validate:"required"`
	Pupae  int `json:"pupae" validate:"required"`
	Adult  int `json:"adult" validate:"required"`
}
