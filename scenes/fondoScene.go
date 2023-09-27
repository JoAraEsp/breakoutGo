package scenes

import (
	"image"
	_ "image/png"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type FondoScene struct {
	background *pixel.Sprite
}

func NewFondoScene(backgroundImagePath string) (*FondoScene, error) {

	background, err := cargarSprite(backgroundImagePath)
	if err != nil {
		return nil, err
	}

	return &FondoScene{
		background: background,
	}, nil
}

func (fs *FondoScene) Dibujar(win *pixelgl.Window) {

	fs.background.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
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
