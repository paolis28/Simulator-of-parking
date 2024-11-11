package carContro

import (
	"parking/src/views/car_view"
	"parking/src/models/car"
	"parking/src/models/parkingModel"
	"parking/src/models/parking_place"
	"parking/src/models/list"
	"math/rand"
	"time"
)

type CarController struct {
	Car        *car.Car
	Parking    *parkingModel.Parking
	CarView    *car_view.CarView
	CarManager *list.CarManager
	DoorChan   chan struct{}
	PathChan   chan struct{}
}

func NewCarController(car *car.Car, parking *parkingModel.Parking, carManager *list.CarManager, doorChan chan struct{}, pathChan chan struct{}) *CarController {
	return &CarController{
		Car:        car,
		Parking:    parking,
		CarManager: carManager,
		DoorChan:   doorChan,
		PathChan:   pathChan,
	}
}

func (cc *CarController) Start() {
	cc.CarManager.AddCar(cc.Car)
	cc.Enqueue()
	spot := cc.Parking.GetAvailableSpot()
	cc.Park(spot)
	time.Sleep(time.Second * time.Duration(rand.Intn(15)+20))
	cc.LeaveSpot()
	cc.Parking.ReleaseSpot(spot)
	cc.Leave(spot)

	<-cc.PathChan
	cc.ExitDoor()
	cc.PathChan <- struct{}{}

	cc.GoAway()
	cc.CarManager.RemoveCar(cc.Car)
}

func (cc *CarController) Enqueue() {
	cc.Parking.QueueCars.Enqueue(cc.Car)
	minY := 45.0    
	spacing := 50.0 

	for cc.Car.Y > minY {
		carAhead := cc.Parking.QueueCars.GetCarAhead(cc.Car)
		canMove := true
		if carAhead != nil {
			_, aheadY := carAhead.GetPosition()
			ccY := cc.Car.Y
			if ccY-aheadY < spacing {
				canMove = false
			}
		}
		if canMove {
			cc.Car.SetDirection(0, -1)
			cc.Car.Move(0, -1)
		} else {
			cc.Car.SetDirection(0, -1) 
		}
		time.Sleep(10 * time.Millisecond)
	}

	<-cc.DoorChan
	defer func() { cc.DoorChan <- struct{}{} }()
	<-cc.PathChan
	defer func() { cc.PathChan <- struct{}{} }()
	cc.JoinDoor()
	cc.Parking.QueueCars.RemoveCar(cc.Car)
}

func (cc *CarController) JoinDoor() {
	minDistance := 50.0                                    
	for cc.Car.X < 355 {
		canMove := true
		for _, otherCar := range cc.CarManager.GetCars() {
			if otherCar != cc.Car {
				otherX, otherY := otherCar.GetPosition()
				if cc.Car.Y == otherY && cc.Car.X < otherX && otherX-cc.Car.X < minDistance {
					canMove = false
					break
				}
			}
		}
		if canMove {
			cc.Car.SetDirection(1, 0) 
			cc.Car.Move(1, 0)
		} else {
			cc.Car.SetDirection(1, 0) 
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func (cc *CarController) Park(spot *parking_place.ParkingPlace) {
	for _, direction := range spot.GetDirectionsForParking() {
		cc.move(direction)
	}
}

func (cc *CarController) LeaveSpot() {
	cc.Car.SetDirection(0, -1)                                
	cc.Car.Move(0, -30)
}

func (cc *CarController) Leave(spot *parking_place.ParkingPlace) {
	for _, direction := range spot.GetDirectionsForLeaving() {
		cc.move(direction)
	}
}

func (cc *CarController) ExitDoor() {
	minDistance := 50.0                        
	for cc.Car.X > 300 {
		canMove := true
		for _, otherCar := range cc.CarManager.GetCars() {
			if otherCar != cc.Car {
				otherX, otherY := otherCar.GetPosition()
				if cc.Car.Y == otherY && cc.Car.X > otherX && cc.Car.X-otherX < minDistance {
					canMove = false
					break
				}
			}
		}
		if canMove {
			cc.Car.SetDirection(-1, 0) 
			cc.Car.Move(-1, 0)
		} else {
			cc.Car.SetDirection(-1, 0) 
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func (cc *CarController) GoAway() {
	cc.Car.SetDirection(-1, 0)     
	for cc.Car.X > -20 {
		cc.Car.Move(-1, 0)
		time.Sleep(5 * time.Millisecond)
	}
	cc.CarManager.RemoveCar(cc.Car)
}

func (cc *CarController) move(direction *parking_place.ParkingSpotDirection) {
	minDistance := 50.0
	for {
		var canMove bool = true
		var dx, dy float64

		switch direction.Direction {
		case "left":
			if cc.Car.X <= direction.Point {
				return
			}
			dx, dy = -1, 0
			for _, otherCar := range cc.CarManager.GetCars() {
				if otherCar != cc.Car {
					otherX, otherY := otherCar.GetPosition()
					if cc.Car.Y == otherY && cc.Car.X > otherX && cc.Car.X-otherX < minDistance {
						canMove = false
						break
					}
				}
			}
		case "right":
			if cc.Car.X >= direction.Point {
				return
			}
			dx, dy = 1, 0
			for _, otherCar := range cc.CarManager.GetCars() {
				if otherCar != cc.Car {
					otherX, otherY := otherCar.GetPosition()
					if cc.Car.Y == otherY && cc.Car.X < otherX && otherX-cc.Car.X < minDistance {
						canMove = false
						break
					}
				}
			}
		case "up":
			if cc.Car.Y <= direction.Point {
				return
			}
			dx, dy = 0, -1
			for _, otherCar := range cc.CarManager.GetCars() {
				if otherCar != cc.Car {
					otherX, otherY := otherCar.GetPosition()
					if cc.Car.X == otherX && cc.Car.Y > otherY && cc.Car.Y-otherY < minDistance {
						canMove = false
						break
					}
				}
			}
		case "down":
			if cc.Car.Y >= direction.Point {
				return
			}
			dx, dy = 0, 1
			for _, otherCar := range cc.CarManager.GetCars() {
				if otherCar != cc.Car {
					otherX, otherY := otherCar.GetPosition()
					if cc.Car.X == otherX && cc.Car.Y < otherY && otherY-cc.Car.Y < minDistance {
						canMove = false
						break
					}
				}
			}
		}

		if canMove {
			cc.Car.SetDirection(dx, dy)
			cc.Car.Move(dx, dy)
		} else {
			cc.Car.SetDirection(dx, dy)
		}
		time.Sleep(10 * time.Millisecond)
	}
}
