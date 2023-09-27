package models

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Pelota struct {
	Posicion  pixel.Vec
	Velocidad pixel.Vec
	Radius    float64
	Sprite    *pixel.Sprite
}

func NuevaPelota(posicion pixel.Vec, sprite *pixel.Sprite) *Pelota {
	return &Pelota{
		Posicion:  posicion,
		Velocidad: pixel.V(3, -3), // Puedes ajustar la velocidad inicial según lo que desees
		Radius:    10,             // Este valor es arbitrario, ajústalo según el tamaño de tu sprite
		Sprite:    sprite,
	}
}

func (p *Pelota) Actualizar(win *pixelgl.Window, barra *Barra) {
	p.Posicion = p.Posicion.Add(p.Velocidad)

	// Rebotar en los lados de la pantalla
	if p.Posicion.X-p.Radius <= 0 || p.Posicion.X+p.Radius >= win.Bounds().W() {
		p.Velocidad.X = -p.Velocidad.X
	}

	// Rebotar en la parte superior de la pantalla
	if p.Posicion.Y+p.Radius >= win.Bounds().H() {
		p.Velocidad.Y = -p.Velocidad.Y
	}

	// Colisión con la barra
	// Nota: Esto es una colisión rectangular simple, puede requerir ajustes para ser más preciso.
	if p.Posicion.X-p.Radius < barra.Posicion.X+barra.Tamano.X &&
		p.Posicion.X+p.Radius > barra.Posicion.X &&
		p.Posicion.Y-p.Radius < barra.Posicion.Y+barra.Tamano.Y &&
		p.Posicion.Y+p.Radius > barra.Posicion.Y {
		p.Velocidad.Y = -p.Velocidad.Y
	}

	// En caso de que la pelota caiga por debajo de la ventana, podrías restablecer su posición o manejarlo como un "game over".
	if p.Posicion.Y-p.Radius < 0 {
		// p.Posicion = win.Bounds().Center() // Por ejemplo, reseteando la posición de la pelota
	}
}

func (p *Pelota) Dibujar(win *pixelgl.Window) {
	p.Sprite.Draw(win, pixel.IM.Scaled(pixel.ZV, 2*p.Radius/p.Sprite.Frame().W()).Moved(p.Posicion))
}
