package models

import (
	"sync"

	"github.com/oakmound/oak/v4/alg/floatgeom"
	"github.com/oakmound/oak/v4/entities"
	"github.com/oakmound/oak/v4/render"
	"github.com/oakmound/oak/v4/render/mod"
	"github.com/oakmound/oak/v4/scene"
)

const (
	EntranceSpotX = 355.00 // punto de entrada estacionamiento
	Speed         = 10     // velocidad del auto
)

type Car struct {
	mu     sync.Mutex
	area   floatgeom.Rect2
	entity *entities.Entity
}

func NewCar(ctx *scene.Context) *Car {
	area := floatgeom.NewRect2(300, 490, 320, 510)

	modelToImage := map[int]string{
		1: "assets/img/brownCar.png",
		2: "assets/img/greenCar.png",
		3: "assets/img/orangeCar.png",
		4: "assets/img/whiteCar.png",
	}
	carModelNum := int(getRandomNumber(1, 4))
	carModel, _ := modelToImage[carModelNum]

	sprite, _ := render.LoadSprite(carModel)

	newSwitch := render.NewSwitch("up", map[string]render.Modifiable{
		"up":    sprite,
		"down":  sprite.Copy().Modify(mod.FlipY),
		"left":  sprite.Copy().Modify(mod.Rotate(90)),
		"right": sprite.Copy().Modify(mod.Rotate(-90)),
	})

	entity := entities.New(ctx, entities.WithRect(area), entities.WithRenderable(newSwitch), entities.WithDrawLayers([]int{1, 2}))

	return &Car{
		area:   area,
		entity: entity,
	}
}
