package models

import (

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Barra struct {
	Posicion pixel.Vec
	Tamano   pixel.Vec
	Velocidad float64
	Sprite   *pixel.Sprite // O algún tipo de representación gráfica
}

func NuevaBarra(posicion pixel.Vec, sprite *pixel.Sprite) *Barra {
	return &Barra{
		Posicion:  posicion,
		Tamano:    pixel.V(100, 20),
		Velocidad: 5.0,
		Sprite:    sprite, // Asignamos el sprite aquí
	}
}

func (b *Barra) Actualizar(win *pixelgl.Window) {
	if win.Pressed(pixelgl.KeyLeft) {
		b.Posicion.X -= b.Velocidad
	}
	if win.Pressed(pixelgl.KeyRight) {
		b.Posicion.X += b.Velocidad
	}

	// Actualización basada en el movimiento del ratón
	if win.MouseInsideWindow() {
		b.Posicion.X = win.MousePosition().X
	}

	// Asegurar que la barra no salga de la ventana (ajustar según necesidad)
	if b.Posicion.X < 0 {
		b.Posicion.X = 0
	}
	if b.Posicion.X > win.Bounds().W() - b.Tamano.X {
		b.Posicion.X = win.Bounds().W() - b.Tamano.X
	}
}

func (b *Barra) Dibujar(win *pixelgl.Window) {
	// Aquí va la lógica para dibujar la barra en la ventana.
	// Por ejemplo, si estás usando un sprite para la barra, lo dibujarías aquí.
	b.Sprite.Draw(win, pixel.IM.Moved(b.Posicion.Add(b.Tamano.Scaled(0.5)))) // Centra el sprite en la posición de la barra
}
