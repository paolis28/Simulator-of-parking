package car_view

import (
	"parking/src/models/car"
	"github.com/oakmound/oak/v4/render"
	"github.com/oakmound/oak/v4/render/mod"
	"github.com/oakmound/oak/v4/scene"
)

type CarView struct {
	Car     *car.Car
	Sprite  *render.Switch
	Context *scene.Context
}

func NewCarView(car *car.Car, ctx *scene.Context) *CarView {
	sprite, err := render.LoadSprite(car.ModelPath)
	if err != nil {
		return nil
	}

	// Crear versiones modificadas del sprite para cada dirección
	upSprite := sprite
	downSprite := sprite.Copy().Modify(mod.FlipY)
	leftSprite := sprite.Copy().Modify(mod.Rotate(90))
	rightSprite := sprite.Copy().Modify(mod.Rotate(-90))

	// Crear el render.Switch con las versiones para cada dirección
	spriteSwitch := render.NewSwitch("up", map[string]render.Modifiable{
		"up":    upSprite,
		"down":  downSprite,
		"left":  leftSprite,
		"right": rightSprite,
	})

	x, y := car.GetPosition()
	spriteSwitch.SetPos(x, y)
	render.Draw(spriteSwitch, 3) // Dibujar el sprite en la capa 3

	carView := &CarView{
		Car:     car,
		Sprite:  spriteSwitch,
		Context: ctx,
	}

	car.RegisterObserver(carView)
	return carView
}

// Update actualiza la posición y dirección del sprite del auto.
func (cv *CarView) Update(data interface{}) {
	car := data.(*car.Car)
	x, y := car.GetPosition()
	cv.Sprite.SetPos(x, y)

	direction := car.GetDirectionName()
	cv.Sprite.Set(direction)
}
