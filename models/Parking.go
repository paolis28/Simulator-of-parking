package models

import "sync"

// Parking representa el estacionamiento, incluyendo los lugares de estacionamiento y la cola de autos
type Parking struct {
	mu            sync.Mutex     // Mutex para sincronizar el acceso a los lugares de estacionamiento
	spots         []*ParkingSpot // Lista de lugares de estacionamiento
	queueCars     *CarQueue      // Cola de autos esperando para estacionarse
	availableCond *sync.Cond     // Condición para esperar y señalar la disponibilidad de lugares
}

// NewParking crea una nueva instancia de Parking con una lista de lugares de estacionamiento
func NewParking(spots []*ParkingSpot) *Parking {
	queue := NewCarQueue() // Inicializa una nueva cola de autos

	p := &Parking{
		spots:     spots, // Asigna los lugares de estacionamiento
		queueCars: queue, // Asigna la cola de autos
	}
	// Inicializa la condición de disponibilidad usando el mutex del estacionamiento
	p.availableCond = sync.NewCond(&p.mu)

	return p
}

// GetSpots devuelve la lista de lugares de estacionamiento
func (p *Parking) GetSpots() []*ParkingSpot {
	return p.spots
}

// GetParkingSpotAvailable busca un lugar de estacionamiento disponible
func (p *Parking) GetParkingSpotAvailable() *ParkingSpot {
	p.mu.Lock()         // Bloquea el mutex para acceso seguro
	defer p.mu.Unlock() // Asegura el desbloqueo al finalizar la función

	for {
		// Recorre la lista de lugares de estacionamiento
		for _, spot := range p.spots {
			if spot.GetIsAvailable() { // Si encuentra un lugar disponible
				spot.SetIsAvailable(false) // Lo marca como no disponible
				return spot                // Devuelve el lugar
			}
		}
		// Si no hay lugares disponibles, espera hasta que se libere uno
		p.availableCond.Wait()
	}
}

// ReleaseParkingSpot libera un lugar de estacionamiento, señalizando que está disponible
func (p *Parking) ReleaseParkingSpot(spot *ParkingSpot) {
	p.mu.Lock()         // Bloquea el mutex para acceso seguro
	defer p.mu.Unlock() // Asegura el desbloqueo al finalizar la función

	spot.SetIsAvailable(true) // Marca el lugar como disponible
	p.availableCond.Signal()  // Señaliza que hay un lugar disponible
}

// GetQueueCars devuelve la cola de autos del estacionamiento
func (p *Parking) GetQueueCars() *CarQueue {
	return p.queueCars
}
