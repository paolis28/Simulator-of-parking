package parkingModel

import (
	"parking/src/models/parking_place"
	"parking/src/models/list"
)

// Parking representa el estacionamiento.
type Parking struct {
	Spots          []*parking_place.ParkingPlace
	QueueCars      *list.CarList
	AvailableSpots chan *parking_place.ParkingPlace
}

// NewParking crea una nueva instancia de Parking.
func NewParking(spots []*parking_place.ParkingPlace) *Parking {
	availableSpots := make(chan *parking_place.ParkingPlace, len(spots))
	for _, spot := range spots {
		availableSpots <- spot
	}

	return &Parking{
		Spots:          spots,
		QueueCars:      list.NewCarQueue(),
		AvailableSpots: availableSpots,
	}
}

// GetAvailableSpot obtiene un lugar disponible.
func (p *Parking) GetAvailableSpot() *parking_place.ParkingPlace {
	return <-p.AvailableSpots
}

// ReleaseSpot libera un lugar ocupado.
func (p *Parking) ReleaseSpot(spot *parking_place.ParkingPlace) {
	p.AvailableSpots <- spot
}
