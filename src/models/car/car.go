package car

import (
	"math"
	"math/rand"
	"sync"
	"time"
	"parking/src/models/observer"
)

type Car struct {
	mu        sync.Mutex
	X         float64
	Y         float64
	DX        float64 
	DY        float64 
	ModelPath string
	observers []observer.Observer
}

func NewCar() *Car {
	modelPaths := []string{
		"assets/img/carVan.png",
		"assets/img/carWhite.png",
		"assets/img/carPolice.png",
		"assets/img/carBlack.png",
	}
	rand.Seed(time.Now().UnixNano())
	modelPath := modelPaths[rand.Intn(len(modelPaths))]

	return &Car{
		X:         300, 
		Y:         400, 
		DX:        0,
		DY:        -1, 
		ModelPath: modelPath,
		observers: []observer.Observer{},
	}
}

// RegisterObserver agrega un observador al auto.
func (c *Car) RegisterObserver(o observer.Observer) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.observers = append(c.observers, o)
}

// RemoveObserver elimina un observador del auto.
func (c *Car) RemoveObserver(o observer.Observer) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for i, observer := range c.observers {
		if observer == o {
			c.observers = append(c.observers[:i], c.observers[i+1:]...)
			break
		}
	}
}

//notifica a todos los observadores sobre un cambio.
func (c *Car) NotifyObservers() {
	c.mu.Lock()
	observers := make([]observer.Observer, len(c.observers))
	copy(observers, c.observers)
	c.mu.Unlock()

	for _, observer := range observers {
		observer.Update(c)
	}
}

//notifica a los observadores.
func (c *Car) Move(dx, dy float64) {
	c.mu.Lock()
	c.X += dx
	c.Y += dy
	c.mu.Unlock()
	c.NotifyObservers()
}

func (c *Car) SetDirection(dx, dy float64) {
	c.mu.Lock()
	c.DX = dx
	c.DY = dy
	c.mu.Unlock()
	c.NotifyObservers()
}

func (c *Car) GetPosition() (float64, float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.X, c.Y
}

func (c *Car) GetDirection() (float64, float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.DX, c.DY
}

func (c *Car) GetDirectionName() string {
	c.mu.Lock()
	dx := c.DX
	dy := c.DY
	c.mu.Unlock()

	if dx == 0 && dy == 0 {
		return "up" 
	}

	angle := math.Atan2(dy, dx) * (180 / math.Pi)

	if angle >= -45 && angle <= 45 {
		return "right"
	} else if angle > 45 && angle < 135 {
		return "up" 
	} else if angle >= 135 || angle <= -135 {
		return "left"
	} else {
		return "down"
	}
}

//posición X 
func (c *Car) SetX(x float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.X = x
	c.NotifyObservers()
}

//posición Y 
func (c *Car) SetY(y float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Y = y
	c.NotifyObservers()
}
