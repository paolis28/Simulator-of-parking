package models

import (
	"github.com/oakmound/oak/v4/alg/floatgeom"
)

// ParkingSpotDirection representa una dirección específica y un punto asociado para maniobrar en un estacionamiento.
type ParkingSpotDirection struct {
	Direction string // Dirección del movimiento: "left", "right", "up", "down".
	Point     float64
}

// NewParkingSpotDirection crea y devuelve una nueva instancia de ParkingSpotDirection.
func NewParkingSpotDirection(direction string, point float64) *ParkingSpotDirection {
	return &ParkingSpotDirection{
		Direction: direction,
		Point:     point,
	}
}

// ParkingSpot representa un lugar de estacionamiento.
type ParkingSpot struct {
	Area                 *floatgeom.Rect2        // Área que delimita el lugar de estacionamiento.
	DirectionsForParking []*ParkingSpotDirection // Direcciones para estacionar en este lugar.
	DirectionsForLeaving []*ParkingSpotDirection // Direcciones para salir de este lugar.
	Number               int
	IsAvailable          bool
}

// NewParkingSpot crea y devuelve un nuevo lugar de estacionamiento con los parámetros especificados.
func NewParkingSpot(x, y, x2, y2 float64, row, number int) *ParkingSpot {
	directionsForParking := GetDirectionsForParking(x, y, row)
	directionsForLeaving := GetDirectionsForLeaving()
	area := floatgeom.NewRect2(x, y, x2, y2)

	return &ParkingSpot{
		Area:                 &area,
		DirectionsForParking: directionsForParking,
		DirectionsForLeaving: directionsForLeaving,
		Number:               number,
		IsAvailable:          true,
	}
}

// GetDirectionsForParking genera y devuelve las direcciones necesarias para estacionar en función de la fila.
func GetDirectionsForParking(x, y float64, row int) []*ParkingSpotDirection {
	var directions []*ParkingSpotDirection

	// Define la dirección principal basada en la fila.
	switch row {
	case 1:
		directions = append(directions, NewParkingSpotDirection("down", 45))
	case 2:
		directions = append(directions, NewParkingSpotDirection("down", 135))
	case 3:
		directions = append(directions, NewParkingSpotDirection("down", 225))
	case 4:
		directions = append(directions, NewParkingSpotDirection("down", 315))
	}

	// Añade movimientos adicionales hacia la derecha y abajo.
	directions = append(directions, NewParkingSpotDirection("right", x+5))
	directions = append(directions, NewParkingSpotDirection("down", y+5))

	return directions
}

// GetDirectionsForLeaving genera y devuelve las direcciones necesarias para salir del lugar de estacionamiento.
func GetDirectionsForLeaving() []*ParkingSpotDirection {
	var directions []*ParkingSpotDirection

	// Añade las direcciones para salir del estacionamiento.
	directions = append(directions, NewParkingSpotDirection("right", 600))
	directions = append(directions, NewParkingSpotDirection("up", 15))
	directions = append(directions, NewParkingSpotDirection("left", 355))

	return directions
}

// Métodos de acceso y modificación:

// GetArea devuelve el área del lugar de estacionamiento.
func (p *ParkingSpot) GetArea() *floatgeom.Rect2 {
	return p.Area
}

// GetNumber devuelve el número de identificación del lugar de estacionamiento.
func (p *ParkingSpot) GetNumber() int {
	return p.Number
}

// GetDirectionsForParking devuelve las direcciones necesarias para estacionar en el lugar.
func (p *ParkingSpot) GetDirectionsForParking() []*ParkingSpotDirection {
	return p.DirectionsForParking
}

// GetDirectionsForLeaving devuelve las direcciones necesarias para salir del lugar.
func (p *ParkingSpot) GetDirectionsForLeaving() []*ParkingSpotDirection {
	return p.DirectionsForLeaving
}

// GetIsAvailable verifica si el lugar de estacionamiento está disponible.
func (p *ParkingSpot) GetIsAvailable() bool {
	return p.IsAvailable
}

// SetIsAvailable establece el estado de disponibilidad del lugar de estacionamiento.
func (p *ParkingSpot) SetIsAvailable(isAvailable bool) {
	p.IsAvailable = isAvailable
}

// Métodos adicionales para obtener las coordenadas del área:

// GetX devuelve la coordenada mínima X del lugar de estacionamiento.
func (p *ParkingSpot) GetX() float64 {
	return p.Area.Min.X()
}

// GetY devuelve la coordenada mínima Y del lugar de estacionamiento.
func (p *ParkingSpot) GetY() float64 {
	return p.Area.Min.Y()
}

// GetX2 devuelve la coordenada máxima X del lugar de estacionamiento.
func (p *ParkingSpot) GetX2() float64 {
	return p.Area.Max.X()
}

// GetY2 devuelve la coordenada máxima Y del lugar de estacionamiento.
func (p *ParkingSpot) GetY2() float64 {
	return p.Area.Max.Y()
}
