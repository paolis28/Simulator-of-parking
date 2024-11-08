package controllers

import (
	"estacionamiento/pkg/models"
	"estacionamiento/pkg/views"
	"fmt"
	"math/rand"
	"time"
)

// CarController maneja la lógica de los autos.
type CarController struct {
	Car        *models.Car
	Parking    *models.Parking
	CarView    *views.CarView
	CarManager *models.CarManager
	DoorChan   chan struct{}
	PathChan   chan struct{}
}

// Instancia de CarController.
func NewCarController(car *models.Car, parking *models.Parking, carManager *models.CarManager, doorChan chan struct{}, pathChan chan struct{}) *CarController {
	return &CarController{
		Car:        car,
		Parking:    parking,
		CarManager: carManager,
		DoorChan:   doorChan,
		PathChan:   pathChan,
	}
}

// Start inicia el ciclo de vida del auto.
func (cc *CarController) Start() {
	fmt.Println("Iniciando el ciclo de vida del auto") // Mensaje de depuración
	cc.CarManager.AddCar(cc.Car)

	cc.Enqueue()

	spot := cc.Parking.GetAvailableSpot()

	cc.Park(spot)

	time.Sleep(time.Second * time.Duration(rand.Intn(15)+20))

	cc.LeaveSpot()
	cc.Parking.ReleaseSpot(spot)

	cc.Leave(spot)

	// Los autos que salen adquieren el PathChan para tener prioridad
	<-cc.PathChan
	cc.ExitDoor()
	cc.PathChan <- struct{}{}

	cc.GoAway()
	cc.CarManager.RemoveCar(cc.Car)
}

// Implementación de los métodos de movimiento del auto.
func (cc *CarController) Enqueue() {
	fmt.Println("Auto encolado, iniciando movimiento hacia la cola")
	cc.Parking.QueueCars.Enqueue(cc.Car)

	minY := 45.0    // Posición Y objetivo al frente de la cola
	spacing := 50.0 // Distancia mínima entre autos

	for cc.Car.Y > minY {
		// Obtener el auto delante en la cola
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
			cc.Car.SetDirection(0, -1) // Establecer dirección hacia arriba
			cc.Car.Move(0, -1)
		} else {
			cc.Car.SetDirection(0, -1) // Mantener dirección hacia arriba
		}
		time.Sleep(10 * time.Millisecond)
	}

	// Adquirir acceso a la puerta
	<-cc.DoorChan
	defer func() { cc.DoorChan <- struct{}{} }()

	// Adquirir acceso al camino compartido
	<-cc.PathChan
	defer func() { cc.PathChan <- struct{}{} }()

	// Pasar por la puerta con detección de colisiones
	cc.JoinDoor()

	// Ahora que hemos pasado la puerta, removemos el auto de la cola
	cc.Parking.QueueCars.RemoveCar(cc.Car)
}

func (cc *CarController) JoinDoor() {
	fmt.Println("Auto acercándose a la puerta de entrada") // Mensaje de depuración
	minDistance := 50.0                                    // Distancia mínima entre autos
	for cc.Car.X < 355 {
		canMove := true
		// Verificar colisión
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
			cc.Car.SetDirection(1, 0) // Establecer dirección hacia la derecha
			cc.Car.Move(1, 0)
		} else {
			cc.Car.SetDirection(1, 0) // Mantener dirección hacia la derecha
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func (cc *CarController) Park(spot *models.ParkingSpot) {
	fmt.Println("Auto estacionando") // Mensaje de depuración
	for _, direction := range spot.GetDirectionsForParking() {
		cc.move(direction)
	}
}

func (cc *CarController) LeaveSpot() {
	fmt.Println("Auto saliendo del lugar de estacionamiento") // Mensaje de depuración
	cc.Car.SetDirection(0, -1)                                // Establecer dirección hacia arriba
	cc.Car.Move(0, -30)
}

func (cc *CarController) Leave(spot *models.ParkingSpot) {
	fmt.Println("Auto saliendo del estacionamiento") // Mensaje de depuración
	for _, direction := range spot.GetDirectionsForLeaving() {
		cc.move(direction)
	}
}

func (cc *CarController) ExitDoor() {
	fmt.Println("Auto saliendo por la puerta") // Mensaje de depuración
	minDistance := 50.0                        // Distancia mínima entre autos
	for cc.Car.X > 300 {
		canMove := true
		// Verificar colisión
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
			cc.Car.SetDirection(-1, 0) // Establecer dirección hacia la izquierda
			cc.Car.Move(-1, 0)
		} else {
			cc.Car.SetDirection(-1, 0) // Mantener dirección hacia la izquierda
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func (cc *CarController) GoAway() {
	fmt.Println("Auto alejándose") // Mensaje de depuración
	cc.Car.SetDirection(-1, 0)     // Establecer dirección hacia la izquierda
	for cc.Car.X > -20 {
		cc.Car.Move(-1, 0)
		time.Sleep(5 * time.Millisecond)
	}
	// Removemos el auto del CarManager al final de GoAway
	cc.CarManager.RemoveCar(cc.Car)
}

// Método auxiliar para mover el auto en una dirección específica con detección de colisiones.
func (cc *CarController) move(direction *models.ParkingSpotDirection) {
	minDistance := 50.0 // Distancia mínima entre autos
	for {
		var canMove bool = true
		var dx, dy float64

		switch direction.Direction {
		case "left":
			if cc.Car.X <= direction.Point {
				return
			}
			dx, dy = -1, 0
			// Verificar colisión
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
			// Verificar colisión
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
			// Verificar colisión
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
			// Verificar colisión
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
			cc.Car.SetDirection(dx, dy) // Establecer la dirección actual
			cc.Car.Move(dx, dy)
		} else {
			// Mantener la dirección actual
			cc.Car.SetDirection(dx, dy)
		}
		time.Sleep(10 * time.Millisecond)
	}
}
