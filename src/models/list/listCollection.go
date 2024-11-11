package list

import (
	"parking/src/models/car"
	"sync"
)

// estructura para la lista del objeto de car de los carros
type CarList struct {
	enqueueCh chan *car.Car         
	dequeueCh chan chan *car.Car  
	removeCh  chan *car.Car        
	doneCh    chan struct{}       
	cars      []*car.Car
	mu   	  sync.Mutex
}

func NewCarQueue() *CarList {
	cq := &CarList{
		enqueueCh: make(chan *car.Car),
		dequeueCh: make(chan chan *car.Car),
		removeCh:  make(chan *car.Car),
		doneCh:    make(chan struct{}), 
		cars:      []*car.Car{},
	}

	// Goroutine para el manejo de la lista
	go func() {
		for {
			select {
			case car := <-cq.enqueueCh:
				cq.cars = append(cq.cars, car)
			case ch := <-cq.dequeueCh:
				if len(cq.cars) > 0 {
					ch <- cq.cars[0]
					cq.cars = cq.cars[1:]
				} else {
					ch <- nil
				}
			case car := <-cq.removeCh:
				for i, c := range cq.cars {
					if c == car {
						cq.cars = append(cq.cars[:i], cq.cars[i+1:]...)
						break
					}
				}
			case <-cq.doneCh: 
				return
			}
		}
	}()

	return cq
}

func (cq *CarList) Enqueue(car *car.Car) {cq.enqueueCh <- car}
func (cq *CarList) Dequeue() *car.Car {
	ch := make(chan *car.Car)
	cq.dequeueCh <- ch
	return <-ch
}
func (cq *CarList) RemoveCar(car *car.Car) {cq.removeCh <- car}
func (cq *CarList) Shutdown() {close(cq.doneCh)}
func (cq *CarList) GetPositionInQueue(car *car.Car) int {
	cq.mu.Lock()
	defer cq.mu.Unlock()
	for i, c := range cq.cars {
		if c == car {
			return i
		}
	}
	return -1
}

func (cq *CarList) GetCarAhead(car *car.Car) *car.Car {
	position := cq.GetPositionInQueue(car)
	if position > 0 {
		return cq.cars[position-1]
	}
	return nil
}


// CarManager gestiona una colección de autos.
type CarManager struct {
	mu   sync.Mutex
	Cars []*car.Car
}

// NewCarManager crea una nueva instancia de CarManager.
func NewCarManager() *CarManager {
	return &CarManager{
		Cars: []*car.Car{},
	}
}

// AddCar añade un auto al gestor.
func (cm *CarManager) AddCar(car *car.Car) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.Cars = append(cm.Cars, car)
}

// RemoveCar elimina un auto del gestor.
func (cm *CarManager) RemoveCar(car *car.Car) {
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
func (cm *CarManager) GetCars() []*car.Car {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	carsCopy := make([]*car.Car, len(cm.Cars))
	copy(carsCopy, cm.Cars)
	return carsCopy
}