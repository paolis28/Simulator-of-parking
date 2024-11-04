package models

import (
	"github.com/oakmound/oak/v4/alg/floatgeom"
)

// ParkingSpot representa un lugar de estacionamiento con sus propiedades y direcciones para maniobras
type ParkingSpot struct {
	area                 *floatgeom.Rect2        // Área que delimita el lugar de estacionamiento
	directionsForParking *[]ParkingSpotDirection // Direcciones para estacionar en este lugar
	directionsForLeaving *[]ParkingSpotDirection // Direcciones para salir de este lugar
	number               int                     // Número de identificación del lugar de estacionamiento
	isAvailable          bool                    // Estado de disponibilidad del lugar
}

// NewParkingSpot crea y devuelve un nuevo lugar de estacionamiento con los parámetros especificados
func NewParkingSpot(x, y, x2, y2 float64, row, number int) *ParkingSpot {
	directionsForParking := getDirectionForParking(x, y, row) // Obtiene las direcciones necesarias para estacionar
	directionsForLeaving := getDirectionsForLeaving()         // Obtiene las direcciones necesarias para salir
	area := floatgeom.NewRect2(x, y, x2, y2)                  // Define el área del lugar de estacionamiento

	return &ParkingSpot{
		area:                 &area,                // Asigna el área al lugar de estacionamiento
		directionsForParking: directionsForParking, // Asigna las direcciones para estacionar
		directionsForLeaving: directionsForLeaving, // Asigna las direcciones para salir
		number:               number,               // Asigna el número de identificación
		isAvailable:          true,                 // Inicialmente, el lugar está disponible
	}
}

// getDirectionForParking genera y devuelve las direcciones necesarias para estacionar en función de la fila
func getDirectionForParking(x, y float64, row int) *[]ParkingSpotDirection {
	var directions []ParkingSpotDirection

	// Define la dirección principal basada en la fila
	if row == 1 {
		directions = append(directions, *newParkingSpotDirection("down", 45))
	} else if row == 2 {
		directions = append(directions, *newParkingSpotDirection("down", 135))
	} else if row == 3 {
		directions = append(directions, *newParkingSpotDirection("down", 225))
	} else if row == 4 {
		directions = append(directions, *newParkingSpotDirection("down", 315))
	}

	// Añade movimientos adicionales hacia la derecha y abajo
	directions = append(directions, *newParkingSpotDirection("right", x+5))
	directions = append(directions, *newParkingSpotDirection("down", y+5))

	return &directions
}

// getDirectionsForLeaving genera y devuelve las direcciones necesarias para salir del lugar de estacionamiento
func getDirectionsForLeaving() *[]ParkingSpotDirection {
	var directions []ParkingSpotDirection

	// Añade las direcciones para salir del estacionamiento
	directions = append(directions, *newParkingSpotDirection("right", 600))
	directions = append(directions, *newParkingSpotDirection("up", 15))
	directions = append(directions, *newParkingSpotDirection("left", 355))

	return &directions
}

// GetArea devuelve el área del lugar de estacionamiento
func (p *ParkingSpot) GetArea() *floatgeom.Rect2 {
	return p.area
}

// GetNumber devuelve el número de identificación del lugar de estacionamiento
func (p *ParkingSpot) GetNumber() int {
	return p.number
}

// GetDirectionsForParking devuelve las direcciones necesarias para estacionar en el lugar
func (p *ParkingSpot) GetDirectionsForParking() *[]ParkingSpotDirection {
	return p.directionsForParking
}

// GetDirectionsForLeaving devuelve las direcciones necesarias para salir del lugar
func (p *ParkingSpot) GetDirectionsForLeaving() *[]ParkingSpotDirection {
	return p.directionsForLeaving
}

// GetIsAvailable verifica si el lugar de estacionamiento está disponible
func (p *ParkingSpot) GetIsAvailable() bool {
	return p.isAvailable
}

// SetIsAvailable establece el estado de disponibilidad del lugar de estacionamiento
func (p *ParkingSpot) SetIsAvailable(isAvailable bool) {
	p.isAvailable = isAvailable
}
