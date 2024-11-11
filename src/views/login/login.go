package login

import (
	"parking/interface" // Asegúrate de importar el paquete correcto de escenas
	"github.com/oakmound/oak/v4/event"
	"github.com/oakmound/oak/v4/render"
	"github.com/oakmound/oak/v4/scene"
	"github.com/oakmound/oak/v4"
)

var username string

func NewLoginScene() *scene.Scene {
	// Creamos un objeto renderizado para el texto que será actualizado
	instruction := render.NewText("Ingrese su nombre:", 100, 100)
	inputDisplay := render.NewText("", 100, 130)

	// Devolvemos la escena con la configuración
	return &scene.Scene{
		Start: func(ctx *scene.Context) {
			// Renderiza el texto de bienvenida e instrucciones
			ctx.DrawStack.Draw(instruction, 0)
			ctx.DrawStack.Draw(inputDisplay, 0)

			// Bind para manejar la entrada de teclado
			event.Bind(ctx, event.KeyDown, func(ke event.Key) int {
				if ke == event.Enter {
					if username != "" {
						// Cambia a la escena de parkingScene después de ingresar el nombre
						scene.SetScene(ctx, "parkingScene")
					}
				} else if ke == event.Backspace && len(username) > 0 {
					username = username[:len(username)-1]
				} else if len(ke.String()) == 1 && len(username) < 15 {
					username += ke.String()
				}
				// Actualiza el texto que muestra el nombre ingresado
				inputDisplay.Text = username
				return 0
			})
		},
		End: func() {
			// Limpia el nombre al terminar la escena
			username = ""
		},
	}
}
