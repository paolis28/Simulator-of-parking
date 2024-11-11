package park

import (
	"math/rand"
	"parking/src/controllers/carContro"
	"parking/src/models/list"
	"parking/src/models/car"
	"parking/src/models/parkingModel"
	"parking/src/views/car_view"
	"parking/src/views/parking_view"
	"time"
	"github.com/oakmound/oak/v4/scene"
)

type ParkingController struct {
	Parking    *parkingModel.Parking
	View       *parking_view.ParkingView
	DoorChan   chan struct{}
	PathChan   chan struct{}
	CarManager *list.CarManager
}

func NewParkingController(parking *parkingModel.Parking) *ParkingController {
	doorChan := make(chan struct{}, 1)
	doorChan <- struct{}{} 

	pathChan := make(chan struct{}, 1)
	pathChan <- struct{}{}

	return &ParkingController{
		Parking:    parking,
		DoorChan:   doorChan,
		PathChan:   pathChan,
		CarManager: list.NewCarManager(),
	}
}

func (pc *ParkingController) StartCarGeneration(ctx *scene.Context) {
	go func() {
		for {
			car := car.NewCar()
			CarController := carContro.NewCarController(car, pc.Parking, pc.CarManager, pc.DoorChan, pc.PathChan)
			carView := car_view.NewCarView(car, ctx)
			CarController.CarView = carView
			go CarController.Start()
			time.Sleep(time.Second * time.Duration(rand.Intn(2)+1))
		}
	}()
}
