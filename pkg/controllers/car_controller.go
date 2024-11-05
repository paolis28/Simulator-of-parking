package controllers

import (
	"estacionamiento/pkg/models"
	"estacionamiento/pkg/views"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// CarController maneja la lógica de los autos.
type CarController struct {
	Car        *models.Car
	Parking    *models.Parking
	CarView    *views.CarView
	CarManager *models.CarManager
	DoorMutex  *sync.Mutex
	PathMutex  *sync.Mutex // Nuevo mutex para el camino compartido
}

// NewCarController crea una nueva instancia de CarController.
func NewCarController(car *models.Car, parking *models.Parking, carManager *models.CarManager, doorMutex *sync.Mutex, pathMutex *sync.Mutex) *CarController {
	return &CarController{
		Car:        car,
		Parking:    parking,
		CarManager: carManager,
		DoorMutex:  doorMutex,
		PathMutex:  pathMutex,
	}
}

// Start inicia el ciclo de vida del auto.
func (cc *CarController) Start() {
	fmt.Println("Iniciando el ciclo de vida del auto") // Mensaje de depuración
	cc.CarManager.AddCar(cc.Car)

	cc.Enqueue()

	spot := cc.Parking.GetAvailableSpot()

	cc.Park(spot)

	time.Sleep(time.Second * time.Duration(rand.Intn(10)+15))

	cc.LeaveSpot()
	cc.Parking.ReleaseSpot(spot)

	cc.Leave(spot)

	// Los autos que salen adquieren el PathMutex para tener prioridad
	cc.PathMutex.Lock()
	cc.ExitDoor()
	cc.PathMutex.Unlock()

	cc.GoAway()
	cc.CarManager.RemoveCar(cc.Car)
}

// Implementación de los métodos de movimiento del auto.
func (cc *CarController) Enqueue() {
	fmt.Println("Auto encolado, iniciando movimiento hacia la cola")
	cc.Parking.QueueCars.Enqueue(cc.Car)
	// No removemos el auto de la cola aquí

	minY := 45.0    // Posición Y objetivo al frente de la cola
	spacing := 20.0 // Distancia mínima entre autos

	for cc.Car.Y > minY {
		// Obtener el auto delante en la cola
		carAhead := cc.Parking.QueueCars.GetCarAhead(cc.Car)
		canMove := true

		if carAhead != nil {
			// Obtener la posición del auto delante
			_, aheadY := carAhead.GetPosition()
			ccY := cc.Car.Y

			if ccY-aheadY < spacing {
				canMove = false
			}
		}

		if canMove {
			cc.Car.Move(0, -1)
		}
		time.Sleep(10 * time.Millisecond)
	}

	// Esperar por la puerta
	cc.DoorMutex.Lock()
	defer cc.DoorMutex.Unlock()

	// Esperar a que el camino compartido esté libre
	cc.PathMutex.Lock()
	defer cc.PathMutex.Unlock()

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
			cc.Car.Move(1, 0)
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
			cc.Car.Move(-1, 0)
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func (cc *CarController) GoAway() {
	fmt.Println("Auto alejándose") // Mensaje de depuración
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
			cc.Car.Move(dx, dy)
		}
		time.Sleep(10 * time.Millisecond)
	}
}
