package scenes

import (
	"parking/src/models/parkingModel"
	"parking/src/models/parking_place"
	"parking/src/controllers/park"
	"parking/src/views/parking_view"
	"github.com/oakmound/oak/v4/scene"
)

func NewParkingScene() *scene.Scene {
	return &scene.Scene{
		Start: func(ctx *scene.Context) {
			spots := []*parking_place.ParkingPlace{
			// ---------- Fila 1 ----------
				//*** columna 1 ***
				parking_place.NewParkingSpot(
					380, 70, 460, 150, 1, 1),
				//*** columna 2 ***
				parking_place.NewParkingSpot(470, 70, 550, 150, 1, 2),
				//*** columna 3 ***
				parking_place.NewParkingSpot(560, 70, 640, 150, 1, 3),
				//*** columna 4 *** 
				parking_place.NewParkingSpot(650, 70, 730, 150, 1, 4),
				//*** columna 5 ***
				parking_place.NewParkingSpot(740, 70, 820, 150, 1, 5),

			// ---------- Fila 2 ----------
				//*** columna 1 ***
				parking_place.NewParkingSpot(380, 180, 460, 260, 2, 6),
				//*** columna 2 ***
				parking_place.NewParkingSpot(470, 180, 550, 260, 2, 7),
				//*** columna 3 ***
				parking_place.NewParkingSpot(560, 180, 640, 260, 2, 8),
				//*** columna 4 ***
				parking_place.NewParkingSpot(650, 180, 730, 260, 2, 9),
				//*** columna 5 ***
				parking_place.NewParkingSpot(740, 180, 820, 260, 2, 10),

			// ---------- Fila 3 ----------
				//*** columna 1 ***  
				parking_place.NewParkingSpot(380, 290, 460, 370, 3, 11),
				//*** columna 2 ***
				parking_place.NewParkingSpot(470, 290, 550, 370, 3, 12),
				//*** columna 3 ***
				parking_place.NewParkingSpot(560, 290, 640, 370, 3, 13),
				//*** columna 4 *** 
				parking_place.NewParkingSpot(650, 290, 730, 370, 3, 14),
				//*** columna 5 ***
				parking_place.NewParkingSpot(740, 290, 820, 370, 3, 15),

			// ---------- Fila 4 ----------
				//*** columna 1 ***
				parking_place.NewParkingSpot(380, 400, 460, 480, 4, 16),
				//*** columna 2 ***
				parking_place.NewParkingSpot(470, 400, 550, 480, 4, 17),
				//*** columna 3 ***
				parking_place.NewParkingSpot(560, 400, 640, 480, 4, 18),
				//*** columna 4 ***
				parking_place.NewParkingSpot(650, 400, 730, 480, 4, 19),
				//*** columna 5 *** 
				parking_place.NewParkingSpot(740, 400, 820, 480, 4, 20),
			}

			parking := parkingModel.NewParking(spots)
			parkingController := park.NewParkingController(parking)
			parkingView := parking_view.NewParkingView(parking, ctx)
			parkingController.View = parkingView
			parkingController.StartCarGeneration(ctx)
		},
	}
}

