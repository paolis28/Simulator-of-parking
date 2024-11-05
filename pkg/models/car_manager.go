package models

import "sync"

// CarManager gestiona una colección de autos.
type CarManager struct {
	mu   sync.Mutex
	Cars []*Car
}

// NewCarManager crea una nueva instancia de CarManager.
func NewCarManager() *CarManager {
	return &CarManager{
		Cars: []*Car{},
	}
}

// AddCar añade un auto al gestor.
func (cm *CarManager) AddCar(car *Car) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.Cars = append(cm.Cars, car)
}

// RemoveCar elimina un auto del gestor.
func (cm *CarManager) RemoveCar(car *Car) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	for i, c := range cm.Cars {
		if c == car {
			cm.Cars = append(cm.Cars[:i], cm.Cars[i+1:]...)
			break
		}
	}
}

// GetCars devuelve la lista de autos.
func (cm *CarManager) GetCars() []*Car {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	carsCopy := make([]*Car, len(cm.Cars))
	copy(carsCopy, cm.Cars)
	return carsCopy
}
