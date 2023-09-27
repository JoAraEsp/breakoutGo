package main

import (
	"breakout/scenes"
	"github.com/faiface/pixel/pixelgl"
)

func run() {
	gameScene, err := scenes.NewGame()
	if err != nil {
		panic(err)
	}
	gameScene.EjecutarJuego()
}

func main() {
	pixelgl.Run(run)
}
