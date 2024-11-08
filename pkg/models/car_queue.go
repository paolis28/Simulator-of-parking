package models

import (
	"sync"
)

// CarQueue representa una cola de autos.
type CarQueue struct {
	mu   sync.Mutex
	cars []*Car // Slice de autos en la cola.
}

// NewCarQueue crea una nueva instancia de CarQueue.
func NewCarQueue() *CarQueue {
	return &CarQueue{
		cars: []*Car{},
	}
}

// Enqueue añade un auto a la cola.
func (cq *CarQueue) Enqueue(car *Car) {
	cq.mu.Lock()
	defer cq.mu.Unlock()
	cq.cars = append(cq.cars, car)
}

// Dequeue elimina y devuelve el primer auto de la cola.
func (cq *CarQueue) Dequeue() *Car {
	cq.mu.Lock()
	defer cq.mu.Unlock()
	if len(cq.cars) == 0 {
		return nil
	}
	car := cq.cars[0]
	cq.cars = cq.cars[1:]
	return car
}

// GetPositionInQueue devuelve la posición de un auto en la cola.
func (cq *CarQueue) GetPositionInQueue(car *Car) int {
	cq.mu.Lock()
	defer cq.mu.Unlock()
	for i, c := range cq.cars {
		if c == car {
			return i
		}
	}
	return -1
}

// GetCarAhead devuelve el auto que está delante en la cola.
func (cq *CarQueue) GetCarAhead(car *Car) *Car {
	position := cq.GetPositionInQueue(car)
	if position > 0 {
		return cq.cars[position-1]
	}
	return nil
}

// RemoveCar elimina un auto de la cola.
func (cq *CarQueue) RemoveCar(car *Car) {
	cq.mu.Lock()
	defer cq.mu.Unlock()
	for i, c := range cq.cars {
		if c == car {
			cq.cars = append(cq.cars[:i], cq.cars[i+1:]...)
			break
		}
	}
}
