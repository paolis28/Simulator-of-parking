package models

// Parking representa el estacionamiento.
type Parking struct {
	Spots          []*ParkingSpot
	QueueCars      *CarQueue
	AvailableSpots chan *ParkingSpot
}

// NewParking crea una nueva instancia de Parking.
func NewParking(spots []*ParkingSpot) *Parking {
	availableSpots := make(chan *ParkingSpot, len(spots))
	for _, spot := range spots {
		availableSpots <- spot
	}

	return &Parking{
		Spots:          spots,
		QueueCars:      NewCarQueue(),
		AvailableSpots: availableSpots,
	}
}

// GetAvailableSpot obtiene un lugar disponible.
func (p *Parking) GetAvailableSpot() *ParkingSpot {
	return <-p.AvailableSpots
}

// ReleaseSpot libera un lugar ocupado.
func (p *Parking) ReleaseSpot(spot *ParkingSpot) {
	p.AvailableSpots <- spot
}
