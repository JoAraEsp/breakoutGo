package main

import (
	"breakout/models"
	"breakout/scenes"
	"image"
	_ "image/png"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func run() {

	cfg := pixelgl.WindowConfig{
		Title:  "Breakout versión patito",
		Bounds: pixel.R(0, 0, 800, 600),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	spriteBarra := cargarSprite("assets/bar.png")
	spritePelota := cargarSprite("assets/ball.png")
	spriteBloque := cargarSprite("assets/block.png")

	barra := models.NuevaBarra(pixel.V(win.Bounds().Center().X, 50), spriteBarra)
	pelota := models.NuevaPelota(win.Bounds().Center(), spritePelota)
	bloques := crearBloques(spriteBloque)

	fondo, err := scenes.NewFondoScene("assets/background.png")
	if err != nil {
		panic(err)
	}

	for !win.Closed() {
		win.Clear(pixel.RGB(0, 0, 0))

		fondo.Dibujar(win)

		barra.Actualizar(win)
		barra.Dibujar(win)

		pelota.Actualizar(win, barra)
		pelota.Dibujar(win)

		for i := 0; i < len(bloques); i++ {
			bloque := bloques[i]
			if bloque.ColisionaConPelota(pelota) {

				pelota.Velocidad.Y = -pelota.Velocidad.Y
				bloques = append(bloques[:i], bloques[i+1:]...)
				i--
			} else {
				bloque.Dibujar(win)
			}
		}

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}

func cargarSprite(ruta string) *pixel.Sprite {
	pic, err := loadPicture(ruta)
	if err != nil {
		panic(err)
	}
	return pixel.NewSprite(pic, pic.Bounds())
}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func crearBloques(sprite *pixel.Sprite) []*models.Bloque {
	var bloques []*models.Bloque
	for x := 50.0; x < 800; x += 60 {
		for y := 500.0; y < 600; y += 30 {
			bloque := models.NuevoBloque(pixel.V(x, y), sprite)
			bloques = append(bloques, bloque)
		}
	}
	return bloques
}
