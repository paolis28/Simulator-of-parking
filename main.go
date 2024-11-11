package main

import (
	"parking/interface"
	"github.com/oakmound/oak/v4"
)

func main() {
	scene := scenes.NewParkingScene()
	oak.AddScene("parkingScene", *scene)
	err := oak.Init("parkingScene", func(window oak.Config) (oak.Config, error) {
		window.Screen.Width = 900
		window.Screen.Height = 600
		window.Assets.ImagePath = "assets/img"
		return window, nil
	})
	if err != nil {
		panic(err)
	}
}

