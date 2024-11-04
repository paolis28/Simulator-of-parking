package models

import (
	"math/rand"
	"sync"
	"time"

	"github.com/oakmound/oak/v4/alg/floatgeom"
	"github.com/oakmound/oak/v4/entities"
	"github.com/oakmound/oak/v4/render"
	"github.com/oakmound/oak/v4/render/mod"
	"github.com/oakmound/oak/v4/scene"
)

const (
	entranceSpotX = 355.00 // Punto de entrada al estacionamiento en el eje X
	speed         = 5      // Velocidad del auto en milisegundos por paso de movimiento
)

// Car representa un auto dentro de la simulación
type Car struct {
	mu     sync.Mutex       // Mutex para manejar concurrencia segura
	area   floatgeom.Rect2  // Área ocupada por el auto
	entity *entities.Entity // Entidad gráfica que representa el auto
}

// NewCar crea una nueva instancia de Car
func NewCar(ctx *scene.Context) *Car {
	area := floatgeom.NewRect2(300, 490, 320, 510) // Área inicial del auto

	// Mapear un número aleatorio a un modelo de auto
	modelToImage := map[int]string{
		1: "assets/img/brownCar.png",
		2: "assets/img/greenCar.png",
		3: "assets/img/orangeCar.png",
		4: "assets/img/whiteCar.png",
	}
	carModelNum := int(getRandomNumber(1, 4)) // Generar un número aleatorio para seleccionar el modelo
	carModel := modelToImage[carModelNum]     // Obtener la imagen del modelo seleccionado

	sprite, _ := render.LoadSprite(carModel) // Cargar el sprite del auto

	// Crear un render switch para manejar diferentes direcciones del auto
	newSwitch := render.NewSwitch("up", map[string]render.Modifiable{
		"up":    sprite,                                // Imagen del auto mirando hacia arriba
		"down":  sprite.Copy().Modify(mod.FlipY),       // Imagen invertida verticalmente
		"left":  sprite.Copy().Modify(mod.Rotate(90)),  // Rotar 90 grados para izquierda
		"right": sprite.Copy().Modify(mod.Rotate(-90)), // Rotar -90 grados para derecha
	})

	// Crear la entidad gráfica del auto
	entity := entities.New(ctx, entities.WithRect(area), entities.WithRenderable(newSwitch), entities.WithDrawLayers([]int{1, 2}))

	return &Car{
		area:   area,
		entity: entity,
	}
}

// moveTowards mueve el auto hacia un punto específico en una dirección dada
func (c *Car) moveTowards(direction string, point float64, manager *CarManager) {
	if direction == "left" {
		for c.X() > point {
			if !c.isCollision("left", manager.GetCars()) {
				c.entity.Renderable.(*render.Switch).Set("left") // Cambiar la dirección del sprite
				c.ShiftX(-1)                                     // Mover el auto a la izquierda
				time.Sleep(speed * time.Millisecond)             // Pausa para simular la velocidad del movimiento
			}
		}
	} else if direction == "right" {
		for c.X() < point {
			if !c.isCollision("right", manager.GetCars()) {
				c.entity.Renderable.(*render.Switch).Set("right")
				c.ShiftX(1)
				time.Sleep(speed * time.Millisecond)
			}
		}
	} else if direction == "up" {
		for c.Y() > point {
			if !c.isCollision("up", manager.GetCars()) {
				c.entity.Renderable.(*render.Switch).Set("up")
				c.ShiftY(-1)
				time.Sleep(speed * time.Millisecond)
			}
		}
	} else if direction == "down" {
		for c.Y() < point {
			if !c.isCollision("down", manager.GetCars()) {
				c.entity.Renderable.(*render.Switch).Set("down")
				c.ShiftY(1)
				time.Sleep(speed * time.Millisecond)
			}
		}
	}
}

// Enqueue mueve el auto hacia la cola de entrada
func (c *Car) Enqueue(manager *CarManager) {
	c.moveTowards("up", 45, manager) // Mover hacia arriba hasta el punto 45
}

// JoinDoor mueve el auto hacia la puerta de entrada
func (c *Car) JoinDoor(manager *CarManager) {
	c.moveTowards("right", entranceSpotX, manager) // Mover hacia la derecha hasta la posición de entrada
}

// ExitDoor mueve el auto hacia la puerta de salida
func (c *Car) ExitDoor(manager *CarManager) {
	c.moveTowards("left", 300, manager) // Mover hacia la izquierda para salir
}

// Park estaciona el auto en el espacio designado
func (c *Car) Park(spot *ParkingSpot, manager *CarManager) {
	for index := 0; index < len(*spot.GetDirectionsForParking()); index++ {
		directions := *spot.GetDirectionsForParking()
		if directions[index].Direction == "right" {
			c.moveTowards("right", directions[index].Point, manager)
		} else if directions[index].Direction == "down" {
			c.moveTowards("down", directions[index].Point, manager)
		}
	}
}

// Leave retira el auto del espacio de estacionamiento
func (c *Car) Leave(spot *ParkingSpot, manager *CarManager) {
	for index := 0; index < len(*spot.GetDirectionsForLeaving()); index++ {
		directions := *spot.GetDirectionsForLeaving()
		if directions[index].Direction == "left" {
			c.moveTowards("left", directions[index].Point, manager)
		} else if directions[index].Direction == "right" {
			c.moveTowards("right", directions[index].Point, manager)
		} else if directions[index].Direction == "up" {
			c.moveTowards("up", directions[index].Point, manager)
		} else if directions[index].Direction == "down" {
			c.moveTowards("down", directions[index].Point, manager)
		}
	}
}

// LeaveSpot mueve el auto fuera del espacio de estacionamiento
func (c *Car) LeaveSpot(manager *CarManager) {
	spotY := c.Y()
	c.moveTowards("up", spotY-30, manager) // Mover hacia arriba para salir del spot
}

// GoAway mueve el auto fuera de la pantalla
func (c *Car) GoAway(manager *CarManager) {
	c.moveTowards("left", -20, manager) // Mover hacia la izquierda fuera de la vista
}

// ShiftY mueve el auto verticalmente
func (c *Car) ShiftY(dy float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entity.ShiftY(dy)
}

// ShiftX mueve el auto horizontalmente
func (c *Car) ShiftX(dx float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entity.ShiftX(dx)
}

// X obtiene la posición X actual del auto
func (c *Car) X() float64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.entity.X()
}

// Y obtiene la posición Y actual del auto
func (c *Car) Y() float64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.entity.Y()
}

// Remove destruye la entidad gráfica del auto
func (c *Car) Remove() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entity.Destroy()
}

// isCollision verifica si el auto colisiona con otro auto en una dirección específica
func (c *Car) isCollision(direction string, cars []*Car) bool {
	minDistance := 30.0 // Distancia mínima para considerar una colisión
	for _, car := range cars {
		if direction == "left" && c.X() > car.X() && c.X()-car.X() < minDistance && c.Y() == car.Y() {
			return true
		} else if direction == "right" && c.X() < car.X() && car.X()-c.X() < minDistance && c.Y() == car.Y() {
			return true
		} else if direction == "up" && c.Y() > car.Y() && c.Y()-car.Y() < minDistance && c.X() == car.X() {
			return true
		} else if direction == "down" && c.Y() < car.Y() && car.Y()-c.Y() < minDistance && c.X() == car.X() {
			return true
		}
	}
	return false
}

// getRandomNumber genera un número aleatorio entre min y max
func getRandomNumber(min, max int) float64 {
	source := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(source)
	return float64(generator.Intn(max-min+1) + min)
}
