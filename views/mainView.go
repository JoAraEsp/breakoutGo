package views

import (
	"breakout/scenes"
	"image"
	_ "image/png"
	"log"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type FondoView struct {
	background *pixel.Sprite
}

func NewFondoView(backgroundImagePath string) (*FondoView, error) {
	background, err := cargarSprite(backgroundImagePath)
	if err != nil {
		return nil, err
	}

	return &FondoView{
		background: background,
	}, nil
}

func (fv *FondoView) Dibujar(win *pixelgl.Window) {
	win.Clear(pixel.RGB(0, 0, 0))
	fv.background.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
}

func cargarSprite(ruta string) (*pixel.Sprite, error) {
	pic, err := loadPicture(ruta)
	if err != nil {
		return nil, err
	}
	return pixel.NewSprite(pic, pic.Bounds()), nil
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

func Run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Breakout versión patito",
		Bounds: pixel.R(0, 0, 800, 600),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		log.Fatalf("Failed to create window: %v", err)
	}

	background, err := NewFondoView("assets/background.png")
	if err != nil {
		log.Fatalf("Failed to create background: %v", err)
	}

	gameScene, err := scenes.NewGameScene(win, "assets/bar.png", "assets/ball.png", "assets/block.png")
	if err != nil {
		log.Fatalf("Failed to create game scene: %v", err)
	}

	for !win.Closed() {
		background.Dibujar(win)  
		gameScene.ActualizarYDibujar()
		win.Update()
	}
}
