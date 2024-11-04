package main

import (
	"estacionamiento/scenes"
	"fmt"

	"github.com/oakmound/oak/v4"
)

func main() {
	fmt.Println("Estacionamiento concurrente.")

	// Crea una nueva escena de estacionamiento
	parkingScene := scenes.NewParkingScene()

	// Inicia la escena del estacionamiento
	parkingScene.Start()

	// Configura e inicializa el motor Oak
	_ = oak.Init("parkingScene", func(c oak.Config) (oak.Config, error) {
		c.BatchLoad = true                // Habilita la carga por lotes de recursos
		c.Assets.ImagePath = "assets/img" // Ruta de las imagenes
		// c.Assets.AudioPath = "assets/audio" // Audio, pero no se usa en este proyecto

		// Dimensiones de la ventana
		c.Screen.Width = 900  // Ancho de la ventana
		c.Screen.Height = 450 // Altura de la ventana

		return c, nil // Devuelve la configuración actualizada
	})
}
