package models

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Barra struct {
	Posicion  pixel.Vec
	Tamano    pixel.Vec
	Velocidad float64
	Sprite    *pixel.Sprite
}

func NuevaBarra(posicion pixel.Vec, sprite *pixel.Sprite) *Barra {
	return &Barra{
		Posicion:  posicion,
		Tamano:    pixel.V(100, 20),
		Velocidad: 5.0,
		Sprite:    sprite,
	}
}

func (b *Barra) Actualizar(win *pixelgl.Window) {
	if win.Pressed(pixelgl.KeyLeft) {
		b.Posicion.X -= b.Velocidad
	}
	if win.Pressed(pixelgl.KeyRight) {
		b.Posicion.X += b.Velocidad
	}

	if win.MouseInsideWindow() {
		b.Posicion.X = win.MousePosition().X
	}

	if b.Posicion.X < 0 {
		b.Posicion.X = 0
	}
	if b.Posicion.X > win.Bounds().W()-b.Tamano.X {
		b.Posicion.X = win.Bounds().W() - b.Tamano.X
	}
}

func (b *Barra) Dibujar(win *pixelgl.Window) {

	b.Sprite.Draw(win, pixel.IM.Moved(b.Posicion.Add(b.Tamano.Scaled(0.5))))
}
