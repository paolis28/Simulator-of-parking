package parking_place

import (
	"github.com/oakmound/oak/v4/alg/floatgeom"
)

type ParkingSpotDirection struct {
	Direction string 
	Point     float64
}

func NewParkingSpotDirection(direction string, point float64) *ParkingSpotDirection {
	return &ParkingSpotDirection{
		Direction: direction,
		Point:     point,
	}
}

type ParkingPlace struct {
	Area                 *floatgeom.Rect2        
	DirectionsForParking []*ParkingSpotDirection
	DirectionsForLeaving []*ParkingSpotDirection 
	Number               int
	IsAvailable          bool
}

func NewParkingSpot(x, y, x2, y2 float64, row, number int) *ParkingPlace {
	directionsForParking := GetDirectionsForParking(x, y, row)
	directionsForLeaving := GetDirectionsForLeaving()
	area := floatgeom.NewRect2(x, y, x2, y2)

	return &ParkingPlace{
		Area:                 &area,
		DirectionsForParking: directionsForParking,
		DirectionsForLeaving: directionsForLeaving,
		Number:               number,
		IsAvailable:          true,
	}
}

func GetDirectionsForParking(x, y float64, row int) []*ParkingSpotDirection {
	var directions []*ParkingSpotDirection
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
	directions = append(directions, NewParkingSpotDirection("right", x+5))
	directions = append(directions, NewParkingSpotDirection("down", y+5))

	return directions
}

func GetDirectionsForLeaving() []*ParkingSpotDirection {
	var directions []*ParkingSpotDirection

	directions = append(directions, NewParkingSpotDirection("right", 600))
	directions = append(directions, NewParkingSpotDirection("up", 15))
	directions = append(directions, NewParkingSpotDirection("left", 355))

	// directions = append(directions, NewParkingSpotDirection("right", 600))
	// directions = append(directions, NewParkingSpotDirection("up", 15))
	// directions = append(directions, NewParkingSpotDirection("left", 270))
	// directions = append(directions, NewParkingSpotDirection("down", 600))
	return directions
}


func (p *ParkingPlace) GetArea() *floatgeom.Rect2 {return p.Area}
func (p *ParkingPlace) GetNumber() int {return p.Number}
func (p *ParkingPlace) GetDirectionsForParking() []*ParkingSpotDirection {return p.DirectionsForParking}
func (p *ParkingPlace) GetDirectionsForLeaving() []*ParkingSpotDirection {return p.DirectionsForLeaving}
func (p *ParkingPlace) GetIsAvailable() bool {return p.IsAvailable}
func (p *ParkingPlace) SetIsAvailable(isAvailable bool) {p.IsAvailable = isAvailable}
func (p *ParkingPlace) GetX() float64 {return p.Area.Min.X()}
func (p *ParkingPlace) GetY() float64 {return p.Area.Min.Y()}
func (p *ParkingPlace) GetX2() float64 {return p.Area.Max.X()}
func (p *ParkingPlace) GetY2() float64 {return p.Area.Max.Y()}
