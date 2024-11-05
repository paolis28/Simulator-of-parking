package main

import (
	"estacionamiento/scenes"

	"github.com/oakmound/oak/v4"
)

func main() {
	scene := scenes.NewParkingScene()
	oak.AddScene("parkingScene", *scene)
	err := oak.Init("parkingScene", func(c oak.Config) (oak.Config, error) {
		c.Screen.Width = 900
		c.Screen.Height = 400
		c.Assets.ImagePath = "assets/img"
		return c, nil
	})
	if err != nil {
		panic(err)
	}
}
