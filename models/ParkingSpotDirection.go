package models

// ParkingSpotDirection representa una dirección específica y un punto asociado para maniobrar en un estacionamiento
type ParkingSpotDirection struct {
	Direction string  // Dirección del movimiento, puede ser "left", "right", "up", o "down"
	Point     float64 // Punto específico hacia el cual se debe mover
}

// newParkingSpotDirection crea y devuelve una nueva instancia de ParkingSpotDirection
func newParkingSpotDirection(direction string, point float64) *ParkingSpotDirection {
	return &ParkingSpotDirection{
		Direction: direction, // Asigna la dirección de movimiento
		Point:     point,     // Asigna el punto específico para el movimiento
	}
}
