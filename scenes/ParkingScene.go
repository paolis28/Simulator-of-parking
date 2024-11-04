package scenes

import (
	"estacionamiento/models"
	"image/color"
	"math/rand"
	"sync"
	"time"

	"github.com/oakmound/oak/v4"
	"github.com/oakmound/oak/v4/alg/floatgeom"
	"github.com/oakmound/oak/v4/entities"
	"github.com/oakmound/oak/v4/event"
	"github.com/oakmound/oak/v4/scene"
)

var (
	// Definición de los lugares de estacionamiento
	spots = []*models.ParkingSpot{
		models.NewParkingSpot(380, 70, 410, 100, 1, 1),
		models.NewParkingSpot(425, 70, 455, 100, 1, 2),
		models.NewParkingSpot(470, 70, 500, 100, 1, 3),
		models.NewParkingSpot(515, 70, 545, 100, 1, 4),
		models.NewParkingSpot(560, 70, 590, 100, 1, 5),

		models.NewParkingSpot(380, 160, 410, 190, 2, 6),
		models.NewParkingSpot(425, 160, 455, 190, 2, 7),
		models.NewParkingSpot(470, 160, 500, 190, 2, 8),
		models.NewParkingSpot(515, 160, 545, 190, 2, 9),
		models.NewParkingSpot(560, 160, 590, 190, 2, 10),

		models.NewParkingSpot(380, 250, 410, 280, 3, 11),
		models.NewParkingSpot(425, 250, 455, 280, 3, 12),
		models.NewParkingSpot(470, 250, 500, 280, 3, 13),
		models.NewParkingSpot(515, 250, 545, 280, 3, 14),
		models.NewParkingSpot(560, 250, 590, 280, 3, 15),

		models.NewParkingSpot(380, 340, 410, 370, 4, 16),
		models.NewParkingSpot(425, 340, 455, 370, 4, 17),
		models.NewParkingSpot(470, 340, 500, 370, 4, 18),
		models.NewParkingSpot(515, 340, 545, 370, 4, 19),
		models.NewParkingSpot(560, 340, 590, 370, 4, 20),
	}
	parking    = models.NewParking(spots) // Inicializa el estacionamiento con los lugares definidos
	doorMutex  sync.Mutex                 // Mutex para sincronizar el acceso a la puerta
	carManager = models.NewCarManager()   // Gestor de autos
)

// ParkingScene representa la escena del estacionamiento
type ParkingScene struct {
}

// NewParkingScene crea una nueva instancia de ParkingScene
func NewParkingScene() *ParkingScene {
	return &ParkingScene{}
}

// Start inicia la escena del estacionamiento
func (ps *ParkingScene) Start() {
	isFirstTime := true

	_ = oak.AddScene("parkingScene", scene.Scene{
		Start: func(ctx *scene.Context) {
			_ = ctx.Window.SetBorderless(true) // Configura la ventana sin bordes
			setUpScene(ctx)                    // Configura la escena del estacionamiento

			event.GlobalBind(ctx, event.Enter, func(enterPayload event.EnterPayload) event.Response {
				if !isFirstTime {
					return 0
				}

				isFirstTime = false

				// Inicia un ciclo de generación de autos
				for {
					go carCycle(ctx)
					time.Sleep(time.Millisecond * time.Duration(getRandomNumber(1000, 2000)))
				}
			})
		},
	})
}

// setUpScene configura visualmente la escena del estacionamiento
func setUpScene(ctx *scene.Context) {
	// Dibuja el área de fondo del estacionamiento
	parkingArea := floatgeom.NewRect2(0, 0, 1000, 1000)
	entities.New(ctx, entities.WithRect(parkingArea), entities.WithColor(color.RGBA{50, 50, 50, 255}), entities.WithDrawLayers([]int{0}))

	// Carretera superior
	entities.New(ctx, entities.WithRect(floatgeom.NewRect2(0, 0, 340, 50)), entities.WithColor(color.RGBA{50, 50, 50, 255}), entities.WithDrawLayers([]int{0}))

	// Carretera inferior
	entities.New(ctx, entities.WithRect(floatgeom.NewRect2(0, 950, 1000, 1000)), entities.WithColor(color.RGBA{50, 50, 50, 255}), entities.WithDrawLayers([]int{0}))

	// Carretera izquierda (más ancha)
	entities.New(ctx, entities.WithRect(floatgeom.NewRect2(0, 0, 100, 1000)), entities.WithColor(color.RGBA{50, 50, 50, 255}), entities.WithDrawLayers([]int{0}))

	// Carretera derecha
	entities.New(ctx, entities.WithRect(floatgeom.NewRect2(950, 0, 1000, 1000)), entities.WithColor(color.RGBA{50, 50, 50, 255}), entities.WithDrawLayers([]int{0}))

	// Líneas segmentadas en las carreteras superior e inferior
	for x := 60; x < 320; x += 40 {
		entities.New(ctx, entities.WithRect(floatgeom.NewRect2(float64(x), 24, float64(x+20), 29)), entities.WithColor(color.RGBA{255, 255, 0, 255}), entities.WithDrawLayers([]int{1}))
		entities.New(ctx, entities.WithRect(floatgeom.NewRect2(float64(x), 975, float64(x+20), 980)), entities.WithColor(color.RGBA{255, 255, 255, 255}), entities.WithDrawLayers([]int{1}))
	}

	// Líneas segmentadas en las carreteras izquierda y derecha
	for y := 60; y < 940; y += 40 {
		entities.New(ctx, entities.WithRect(floatgeom.NewRect2(45, float64(y), 50, float64(y+20))), entities.WithColor(color.RGBA{255, 255, 0, 255}), entities.WithDrawLayers([]int{1}))
		entities.New(ctx, entities.WithRect(floatgeom.NewRect2(775, float64(y), 980, float64(y+20))), entities.WithColor(color.RGBA{255, 255, 255, 255}), entities.WithDrawLayers([]int{1}))
	}

	// Líneas rojas y segmentadas en el estacionamiento
	entities.New(ctx, entities.WithRect(floatgeom.NewRect2(340, 5, 630, 10)), entities.WithColor(color.RGBA{255, 0, 0, 255}), entities.WithDrawLayers([]int{0}))
	for x := 340; x < 610; x += 20 {
		entities.New(ctx, entities.WithRect(floatgeom.NewRect2(float64(x), 24, float64(x+15), 29)), entities.WithColor(color.RGBA{255, 255, 0, 255}), entities.WithDrawLayers([]int{0}))
	}

	entities.New(ctx, entities.WithRect(floatgeom.NewRect2(340, 400, 630, 405)), entities.WithColor(color.RGBA{255, 0, 0, 255}), entities.WithDrawLayers([]int{0}))
	entities.New(ctx, entities.WithRect(floatgeom.NewRect2(340, 70, 345, 400)), entities.WithColor(color.RGBA{255, 0, 0, 255}), entities.WithDrawLayers([]int{0}))
	entities.New(ctx, entities.WithRect(floatgeom.NewRect2(625, 10, 630, 400)), entities.WithColor(color.RGBA{255, 0, 0, 255}), entities.WithDrawLayers([]int{0}))

	// Líneas segmentadas en los bordes de cada lugar de estacionamiento
	for _, spot := range spots {
		for y := int(spot.GetArea().Min.Y()) + 2; y < int(spot.GetArea().Max.Y())-2; y += 10 {
			entities.New(ctx, entities.WithRect(floatgeom.NewRect2(spot.GetArea().Min.X(), float64(y), spot.GetArea().Min.X()+2, float64(y+5))), entities.WithColor(color.RGBA{255, 255, 255, 255}))
			entities.New(ctx, entities.WithRect(floatgeom.NewRect2(spot.GetArea().Max.X()-2, float64(y), spot.GetArea().Max.X(), float64(y+5))), entities.WithColor(color.RGBA{255, 255, 255, 255}))
		}

		for x := int(spot.GetArea().Min.X()); x < int(spot.GetArea().Max.X()); x += 20 {
			entities.New(ctx, entities.WithRect(floatgeom.NewRect2(float64(x), spot.GetArea().Min.Y(), float64(x+15), spot.GetArea().Min.Y()+2)), entities.WithColor(color.RGBA{255, 255, 255, 255}))
			entities.New(ctx, entities.WithRect(floatgeom.NewRect2(float64(x), spot.GetArea().Max.Y()-2, float64(x+15), spot.GetArea().Max.Y())), entities.WithColor(color.RGBA{255, 255, 255, 255}))
		}
	}
}

// carCycle maneja el ciclo de vida de un auto en la escena
func carCycle(ctx *scene.Context) {
	car := models.NewCar(ctx) // Crea un nuevo auto

	carManager.AddCar(car)  // Añade el auto al gestor
	car.Enqueue(carManager) // Encola el auto

	spotAvailable := parking.GetParkingSpotAvailable() // Obtiene un lugar de estacionamiento disponible

	doorMutex.Lock()         // Bloquea el acceso a la puerta
	car.JoinDoor(carManager) // El auto se mueve hacia la puerta
	doorMutex.Unlock()       // Desbloquea la puerta

	car.Park(spotAvailable, carManager) // El auto estaciona en el lugar disponible

	time.Sleep(time.Millisecond * time.Duration(getRandomNumber(30000, 50000))) // Tiempo de estacionamiento

	car.LeaveSpot(carManager)                 // El auto sale del lugar de estacionamiento
	parking.ReleaseParkingSpot(spotAvailable) // Libera el lugar de estacionamiento

	car.Leave(spotAvailable, carManager) // El auto sale del estacionamiento

	doorMutex.Lock()         // Bloquea la puerta para evitar colisiones
	car.ExitDoor(carManager) // El auto sale por la puerta
	doorMutex.Unlock()       // Desbloquea la puerta

	car.GoAway(carManager)    // El auto se aleja del área
	car.Remove()              // Elimina el auto de la escena
	carManager.RemoveCar(car) // Elimina el auto del gestor
}

// getRandomNumber genera un número aleatorio entre min y max
func getRandomNumber(min, max int) float64 {
	source := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(source)
	return float64(generator.Intn(max-min+1) + min)
}
