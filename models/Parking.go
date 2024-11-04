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
