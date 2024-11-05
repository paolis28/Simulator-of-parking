package models

import "sync"

// Parking representa el estacionamiento.
type Parking struct {
	mu            sync.Mutex
	Spots         []*ParkingSpot
	QueueCars     *CarQueue
	availableCond *sync.Cond
}

// NewParking crea una nueva instancia de Parking.
func NewParking(spots []*ParkingSpot) *Parking {
	p := &Parking{
		Spots:     spots,
		QueueCars: NewCarQueue(),
	}
	p.availableCond = sync.NewCond(&p.mu)
	return p
}

// GetAvailableSpot obtiene un lugar disponible.
func (p *Parking) GetAvailableSpot() *ParkingSpot {
	p.mu.Lock()
	defer p.mu.Unlock()
	for {
		for _, spot := range p.Spots {
			if spot.IsAvailable {
				spot.IsAvailable = false
				return spot
			}
		}
		p.availableCond.Wait()
	}
}

// ReleaseSpot libera un lugar ocupado.
func (p *Parking) ReleaseSpot(spot *ParkingSpot) {
	p.mu.Lock()
	defer p.mu.Unlock()
	spot.IsAvailable = true
	p.availableCond.Signal()
}
