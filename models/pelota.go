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
		Velocidad: pixel.V(3, -3), 
		Radius:    10,             
		Sprite:    sprite,
	}
}

func (p *Pelota) Actualizar(win *pixelgl.Window, barra *Barra) {
	p.Posicion = p.Posicion.Add(p.Velocidad)

	
	if p.Posicion.X-p.Radius <= 0 || p.Posicion.X+p.Radius >= win.Bounds().W() {
		p.Velocidad.X = -p.Velocidad.X
	}

	
	if p.Posicion.Y+p.Radius >= win.Bounds().H() {
		p.Velocidad.Y = -p.Velocidad.Y
	}

	
	if p.Posicion.X-p.Radius < barra.Posicion.X+barra.Tamano.X &&
		p.Posicion.X+p.Radius > barra.Posicion.X &&
		p.Posicion.Y-p.Radius < barra.Posicion.Y+barra.Tamano.Y &&
		p.Posicion.Y+p.Radius > barra.Posicion.Y {
		p.Velocidad.Y = -p.Velocidad.Y
	}

	
	if p.Posicion.Y-p.Radius < 0 {
		
	}
}

func (p *Pelota) Dibujar(win *pixelgl.Window) {
	p.Sprite.Draw(win, pixel.IM.Scaled(pixel.ZV, 2*p.Radius/p.Sprite.Frame().W()).Moved(p.Posicion))
}
