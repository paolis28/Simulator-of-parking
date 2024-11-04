package main

import (
	"estacionamiento/scenes"
	"fmt"

	"github.com/oakmound/oak/v4"
)

func main() {
	fmt.Println("Estacionamiento concurrente.")
	parkingScene := scenes.NewParkingScene()

	parkingScene.Start()

	_ = oak.Init("parkingScene", func(c oak.Config) (oak.Config, error) {
		c.BatchLoad = true
		c.Assets.ImagePath = "assets/img"
		//c.Assets.AudioPath = "assets/audio"

		c.Screen.Width = 900
		c.Screen.Height = 450
		return c, nil
	})
}
