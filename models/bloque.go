package models

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Bloque struct {
	Posicion pixel.Vec
	Tamano   pixel.Vec
	Sprite   *pixel.Sprite
}

func NuevoBloque(posicion pixel.Vec, sprite *pixel.Sprite) *Bloque {
	return &Bloque{
		Posicion: posicion,
		Tamano:   pixel.V(50, 20),  // Estos valores son ejemplos, ajusta según tu sprite.
		Sprite:   sprite,
	}
}

func (b *Bloque) Dibujar(win *pixelgl.Window) {
	b.Sprite.Draw(win, pixel.IM.Moved(b.Posicion.Add(b.Tamano.Scaled(0.5))))
}

func (b *Bloque) ColisionaConPelota(p *Pelota) bool {
	return p.Posicion.X-p.Radius < b.Posicion.X+b.Tamano.X &&
		p.Posicion.X+p.Radius > b.Posicion.X &&
		p.Posicion.Y-p.Radius < b.Posicion.Y+b.Tamano.Y &&
		p.Posicion.Y+p.Radius > b.Posicion.Y
}
