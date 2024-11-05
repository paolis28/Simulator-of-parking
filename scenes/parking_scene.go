package scenes

import (
	"estacionamiento/pkg/controllers"
	"estacionamiento/pkg/models"
	"estacionamiento/pkg/views"
	"fmt"
	"math/rand"
	"time"

	"github.com/oakmound/oak/v4/scene"
)

func NewParkingScene() *scene.Scene {
	return &scene.Scene{
		Start: func(ctx *scene.Context) {
			spots := []*models.ParkingSpot{
				// Fila 1
				models.NewParkingSpot(380, 70, 410, 100, 1, 1),
				models.NewParkingSpot(425, 70, 455, 100, 1, 2),
				models.NewParkingSpot(470, 70, 500, 100, 1, 3),
				models.NewParkingSpot(515, 70, 545, 100, 1, 4),
				models.NewParkingSpot(560, 70, 590, 100, 1, 5),

				// Fila 2
				models.NewParkingSpot(380, 160, 410, 190, 2, 6),
				models.NewParkingSpot(425, 160, 455, 190, 2, 7),
				models.NewParkingSpot(470, 160, 500, 190, 2, 8),
				models.NewParkingSpot(515, 160, 545, 190, 2, 9),
				models.NewParkingSpot(560, 160, 590, 190, 2, 10),

				// Fila 3
				models.NewParkingSpot(380, 250, 410, 280, 3, 11),
				models.NewParkingSpot(425, 250, 455, 280, 3, 12),
				models.NewParkingSpot(470, 250, 500, 280, 3, 13),
				models.NewParkingSpot(515, 250, 545, 280, 3, 14),
				models.NewParkingSpot(560, 250, 590, 280, 3, 15),

				// Fila 4
				models.NewParkingSpot(380, 340, 410, 370, 4, 16),
				models.NewParkingSpot(425, 340, 455, 370, 4, 17),
				models.NewParkingSpot(470, 340, 500, 370, 4, 18),
				models.NewParkingSpot(515, 340, 545, 370, 4, 19),
				models.NewParkingSpot(560, 340, 590, 370, 4, 20),
			}

			parking := models.NewParking(spots)
			parkingController := controllers.NewParkingController(parking)
			parkingView := views.NewParkingView(parking, ctx)
			parkingController.View = parkingView

			doorChan := make(chan struct{}, 1)
			doorChan <- struct{}{} // Inicializar el semáforo de la puerta

			pathChan := make(chan struct{}, 1)
			pathChan <- struct{}{} // Inicializar el semáforo del camino compartido

			carManager := models.NewCarManager()

			go func() {
				for {
					fmt.Println("Generando un nuevo auto") // Mensaje de depuración
					car := models.NewCar()
					carController := controllers.NewCarController(car, parking, carManager, doorChan, pathChan)
					carView := views.NewCarView(car, ctx)
					carController.CarView = carView
					go carController.Start()
					time.Sleep(time.Second * time.Duration(rand.Intn(2)+1))
				}
			}()
		},
	}
}
