package models

import (
	"math"
	"math/rand"
	"sync"
	"time"
)

// Car representa un auto dentro de la simulación.
type Car struct {
	mu        sync.Mutex
	X         float64
	Y         float64
	DX        float64 // Desplazamiento en X
	DY        float64 // Desplazamiento en Y
	ModelPath string
	observers []Observer
}

// NewCar crea una nueva instancia de Car.
func NewCar() *Car {
	modelPaths := []string{
		"assets/img/brownCar.png",
		"assets/img/greenCar.png",
		"assets/img/orangeCar.png",
		"assets/img/whiteCar.png",
	}
	rand.Seed(time.Now().UnixNano())
	modelPath := modelPaths[rand.Intn(len(modelPaths))]

	return &Car{
		X:         300, // Posición X inicial
		Y:         400, // Posición Y inicial
		DX:        0,
		DY:        -1, // Movimiento inicial hacia arriba
		ModelPath: modelPath,
		observers: []Observer{},
	}
}

// RegisterObserver agrega un observador al auto.
func (c *Car) RegisterObserver(o Observer) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.observers = append(c.observers, o)
}

// RemoveObserver elimina un observador del auto.
func (c *Car) RemoveObserver(o Observer) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for i, observer := range c.observers {
		if observer == o {
			c.observers = append(c.observers[:i], c.observers[i+1:]...)
			break
		}
	}
}

// NotifyObservers notifica a todos los observadores sobre un cambio.
func (c *Car) NotifyObservers() {
	c.mu.Lock()
	observers := make([]Observer, len(c.observers))
	copy(observers, c.observers)
	c.mu.Unlock()

	for _, observer := range observers {
		observer.Update(c)
	}
}

// Move actualiza la posición y dirección del auto y notifica a los observadores.
func (c *Car) Move(dx, dy float64) {
	c.mu.Lock()
	c.X += dx
	c.Y += dy
	c.DX = dx
	c.DY = dy
	c.mu.Unlock()
	c.NotifyObservers()
}

// GetPosition devuelve la posición actual del auto.
func (c *Car) GetPosition() (float64, float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.X, c.Y
}

// GetDirection devuelve la última dirección de movimiento del auto.
func (c *Car) GetDirection() (float64, float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.DX, c.DY
}

// GetDirectionName devuelve la dirección en formato "up", "down", "left", "right".
func (c *Car) GetDirectionName() string {
	c.mu.Lock()
	dx := c.DX
	dy := c.DY
	c.mu.Unlock()

	if dx == 0 && dy == 0 {
		return "up" // Dirección por defecto
	}

	angle := math.Atan2(dy, dx) * (180 / math.Pi)

	if angle >= -45 && angle <= 45 {
		return "right"
	} else if angle > 45 && angle < 135 {
		return "down"
	} else if angle >= 135 || angle <= -135 {
		return "left"
	} else {
		return "up"
	}
}

// SetX establece la posición X del auto.
func (c *Car) SetX(x float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.X = x
	c.NotifyObservers()
}

// SetY establece la posición Y del auto.
func (c *Car) SetY(y float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Y = y
	c.NotifyObservers()
}
