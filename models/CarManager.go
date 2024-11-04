package models

import "sync"

// CarManager gestiona una colección de autos en el estacionamiento
type CarManager struct {
	Mutex sync.Mutex // Mutex para proteger el acceso concurrente a la lista de autos
	Cars  []*Car     // Lista de autos gestionados
}

// NewCarManager crea y devuelve una nueva instancia de CarManager
func NewCarManager() *CarManager {
	return &CarManager{
		Cars: make([]*Car, 0), // Inicializa la lista de autos vacía
	}
}

// AddCar añade un auto a la lista gestionada por CarManager
func (cm *CarManager) AddCar(car *Car) {
	cm.Mutex.Lock()                // Adquiere el lock para acceso seguro
	defer cm.Mutex.Unlock()        // Libera el lock al final del método
	cm.Cars = append(cm.Cars, car) // Añade el auto a la lista
}

// RemoveCar elimina un auto específico de la lista gestionada por CarManager
func (cm *CarManager) RemoveCar(car *Car) {
	cm.Mutex.Lock()
	defer cm.Mutex.Unlock()
	for i, c := range cm.Cars {
		if c == car {
			cm.Cars = append(cm.Cars[:i], cm.Cars[i+1:]...) // Elimina el auto encontrado
			break                                           // Termina el bucle una vez que el auto es encontrado y eliminado
		}
	}
}

// GetCars devuelve la lista de autos gestionados
func (cm *CarManager) GetCars() []*Car {
	cm.Mutex.Lock()
	defer cm.Mutex.Unlock()
	return cm.Cars // Devuelve una copia de la lista de autos
}
