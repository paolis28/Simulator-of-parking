package views

import (
	"estacionamiento/pkg/models"
	"image/color"

	"github.com/oakmound/oak/v4/alg/floatgeom"
	"github.com/oakmound/oak/v4/entities"
	"github.com/oakmound/oak/v4/scene"
)

// ParkingView representa la vista del estacionamiento.
type ParkingView struct {
	Parking *models.Parking
	Context *scene.Context
}

// NewParkingView crea una nueva instancia de ParkingView.
func NewParkingView(parking *models.Parking, ctx *scene.Context) *ParkingView {
	pv := &ParkingView{
		Parking: parking,
		Context: ctx,
	}
	pv.setupScene()
	return pv
}

// setupScene configura visualmente el estacionamiento.
func (pv *ParkingView) setupScene() {
	ctx := pv.Context

	// Dibuja el área de fondo del estacionamiento.
	parkingArea := floatgeom.NewRect2(0, 0, 900, 450)
	entities.New(ctx,
		entities.WithRect(parkingArea),
		entities.WithColor(color.RGBA{50, 50, 50, 255}),
		entities.WithDrawLayers([]int{0}),
	)

	// Carretera superior.
	entities.New(ctx,
		entities.WithRect(floatgeom.NewRect2(0, 0, 340, 50)),
		entities.WithColor(color.RGBA{30, 30, 30, 255}),
		entities.WithDrawLayers([]int{0}),
	)

	// Carretera inferior.
	entities.New(ctx,
		entities.WithRect(floatgeom.NewRect2(0, 400, 900, 450)),
		entities.WithColor(color.RGBA{30, 30, 30, 255}),
		entities.WithDrawLayers([]int{0}),
	)

	// Carretera izquierda (más ancha).
	entities.New(ctx,
		entities.WithRect(floatgeom.NewRect2(0, 0, 100, 450)),
		entities.WithColor(color.RGBA{30, 30, 30, 255}),
		entities.WithDrawLayers([]int{0}),
	)

	// Carretera derecha.
	entities.New(ctx,
		entities.WithRect(floatgeom.NewRect2(800, 0, 900, 450)),
		entities.WithColor(color.RGBA{30, 30, 30, 255}),
		entities.WithDrawLayers([]int{0}),
	)

	// Líneas segmentadas en las carreteras superior e inferior.
	for x := 60; x < 320; x += 40 {
		entities.New(ctx,
			entities.WithRect(floatgeom.NewRect2(float64(x), 24, float64(x+20), 29)),
			entities.WithColor(color.RGBA{255, 255, 0, 255}),
			entities.WithDrawLayers([]int{1}),
		)
		entities.New(ctx,
			entities.WithRect(floatgeom.NewRect2(float64(x), 425, float64(x+20), 430)),
			entities.WithColor(color.RGBA{255, 255, 255, 255}),
			entities.WithDrawLayers([]int{1}),
		)
	}

	// Líneas segmentadas en las carreteras izquierda y derecha.
	for y := 60; y < 390; y += 40 {
		entities.New(ctx,
			entities.WithRect(floatgeom.NewRect2(45, float64(y), 50, float64(y+20))),
			entities.WithColor(color.RGBA{255, 255, 0, 255}),
			entities.WithDrawLayers([]int{1}),
		)
		entities.New(ctx,
			entities.WithRect(floatgeom.NewRect2(850, float64(y), 855, float64(y+20))),
			entities.WithColor(color.RGBA{255, 255, 255, 255}),
			entities.WithDrawLayers([]int{1}),
		)
	}

	// Líneas rojas y segmentadas en el estacionamiento.
	entities.New(ctx,
		entities.WithRect(floatgeom.NewRect2(340, 5, 630, 10)),
		entities.WithColor(color.RGBA{255, 0, 0, 255}),
		entities.WithDrawLayers([]int{0}),
	)
	for x := 340; x < 610; x += 20 {
		entities.New(ctx,
			entities.WithRect(floatgeom.NewRect2(float64(x), 24, float64(x+15), 29)),
			entities.WithColor(color.RGBA{255, 255, 0, 255}),
			entities.WithDrawLayers([]int{0}),
		)
	}

	entities.New(ctx,
		entities.WithRect(floatgeom.NewRect2(340, 400, 630, 405)),
		entities.WithColor(color.RGBA{255, 0, 0, 255}),
		entities.WithDrawLayers([]int{0}),
	)
	entities.New(ctx,
		entities.WithRect(floatgeom.NewRect2(340, 70, 345, 400)),
		entities.WithColor(color.RGBA{255, 0, 0, 255}),
		entities.WithDrawLayers([]int{0}),
	)
	entities.New(ctx,
		entities.WithRect(floatgeom.NewRect2(625, 10, 630, 400)),
		entities.WithColor(color.RGBA{255, 0, 0, 255}),
		entities.WithDrawLayers([]int{0}),
	)

	// Líneas segmentadas en los bordes de cada lugar de estacionamiento.
	for _, spot := range pv.Parking.Spots {
		// Obtener las coordenadas del área del lugar de estacionamiento.
		minX := spot.Area.Min.X()
		minY := spot.Area.Min.Y()
		maxX := spot.Area.Max.X()
		maxY := spot.Area.Max.Y()

		// Líneas verticales segmentadas.
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
		// Líneas horizontales segmentadas.
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
