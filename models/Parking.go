package models

import "sync"

type Parking struct {
	mu            sync.Mutex
	spots         []*ParkingSpot
	queueCars     *CarQueue
	availableCond *sync.Cond
}
