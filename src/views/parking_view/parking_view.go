package parking_view

import (
	"image/color"
	"parking/src/models/parkingModel"
	"github.com/oakmound/oak/v4/alg/floatgeom"
	"github.com/oakmound/oak/v4/entities"
	"github.com/oakmound/oak/v4/render"
	"github.com/oakmound/oak/v4/scene"
)

type ParkingView struct {
	Parking *parkingModel.Parking
	Context *scene.Context
}

func NewParkingView(parking *parkingModel.Parking, ctx *scene.Context) *ParkingView {
	pv := &ParkingView{
		Parking: parking,
		Context: ctx,
	}
	pv.setupScene()
	return pv
}

func (pv *ParkingView) setupScene() {
	ctx := pv.Context

	// Define el área de estacionamiento
	parkingArea := floatgeom.NewRect2(86, 86, 86, 1)

	// Carga la imagen de fondo para el área de estacionamiento
	backgroundImage, err := render.LoadSprite("assets/parking-fondo.jpg")
	if err != nil {
		panic("Error al cargar la imagen de fondo del estacionamiento: " + err.Error())
	}

	// Crea el entity del área de estacionamiento con la imagen como textura
	entities.New(ctx,
		entities.WithRect(parkingArea),
		entities.WithRenderable(backgroundImage),
		entities.WithDrawLayers([]int{0}),
	)

	roadBackground, err := render.LoadSprite("assets/road.jpg")
	if err != nil {
		panic("Error al cargar la imagen de fondo de la carretera: " + err.Error())
	}
	
	entities.New(ctx,
		entities.WithRect(floatgeom.NewRect2(400, 50, 230, 50)),
		entities.WithRenderable(roadBackground),
		entities.WithDrawLayers([]int{0}),
	)
	// Crea las líneas de los cajones de estacionamiento
	for _, spot := range pv.Parking.Spots {
		minX := spot.Area.Min.X()
		minY := spot.Area.Min.Y()
		maxX := spot.Area.Max.X()
		maxY := spot.Area.Max.Y()

		// Líneas verticales de los bordes de cada cajón
		for y := int(minY) + 2; y < int(maxY)-2; y += 10 {
			entities.New(ctx,
				entities.WithRect(floatgeom.NewRect2(minX, float64(y), minX+2, float64(y+5))),
				entities.WithColor(color.RGBA{255, 255, 255, 255}),
				entities.WithDrawLayers([]int{1}),
			)
			entities.New(ctx,
				entities.WithRect(floatgeom.NewRect2(maxX-2, float64(y), maxX, float64(y+5))),
				entities.WithColor(color.RGBA{255, 255, 255, 255}),
				entities.WithDrawLayers([]int{1}),
			)
		}

		// Líneas horizontales de los bordes de cada cajón
		for x := int(minX); x < int(maxX); x += 20 {
			entities.New(ctx,
				entities.WithRect(floatgeom.NewRect2(float64(x), minY, float64(x+15), minY+2)),
				entities.WithColor(color.RGBA{255, 255, 255, 255}),
				entities.WithDrawLayers([]int{1}),
			)
			entities.New(ctx,
				entities.WithRect(floatgeom.NewRect2(float64(x), maxY-2, float64(x+15), maxY)),
				entities.WithColor(color.RGBA{255, 255, 255, 255}),
				entities.WithDrawLayers([]int{1}),
			)
		}
	}
}
