package controllers

import (
	"estacionamiento/pkg/models"
	"estacionamiento/pkg/views"
	"sync"
)

// ParkingController maneja la lógica del estacionamiento.
type ParkingController struct {
	Parking   *models.Parking
	View      *views.ParkingView
	DoorMutex sync.Mutex
}

// NewParkingController crea una nueva instancia de ParkingController.
func NewParkingController(parking *models.Parking) *ParkingController {
	return &ParkingController{
		Parking: parking,
	}
}
