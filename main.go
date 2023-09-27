package main

import (
	"breakout/views"
	"github.com/faiface/pixel/pixelgl"
)

func run() {
	views.Run()
}

func main() {
	pixelgl.Run(run)
}
