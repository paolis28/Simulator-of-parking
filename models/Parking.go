package models

import "sync"

type Parking struct {
	mu            sync.Mutex
	spots         []*ParkingSpot
	queueCars     *CarQueue
	availableCond *sync.Cond
}

func NewParking(spots []*ParkingSpot) *Parking {
	queue := NewCarQueue()

	p := &Parking{
		spots:     spots,
		queueCars: queue,
	}
	p.availableCond = sync.NewCond(&p.mu)

	return p
}

func (p *Parking) GetSpots() []*ParkingSpot {
	return p.spots
}

func (p *Parking) GetParkingSpotAvailable() *ParkingSpot {
	p.mu.Lock()
	defer p.mu.Unlock()

	for {
		for _, spot := range p.spots {
			if spot.GetIsAvailable() {
				spot.SetIsAvailable(false)
				return spot
			}
		}
		p.availableCond.Wait()
	}
}
