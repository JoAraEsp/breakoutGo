package scenes

import (
	"breakout/models"
	"image"
	_ "image/png"
	"os"
	"sync"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type GameScene struct {
	win     *pixelgl.Window
	barra   *models.Barra
	pelota  *models.Pelota
	bloques []*models.Bloque
}

func (gs *GameScene) EjecutarJuego() {
	for !gs.win.Closed() {
		gs.ActualizarYDibujar()
		gs.win.Update()
	}
}

func NewGameScene(win *pixelgl.Window, barPath, ballPath, blockPath string) (*GameScene, error) {
	barraSprite, err := cargarSprite(barPath)
	if err != nil {
		return nil, err
	}

	pelotaSprite, err := cargarSprite(ballPath)
	if err != nil {
		return nil, err
	}

	bloqueSprite, err := cargarSprite(blockPath)
	if err != nil {
		return nil, err
	}

	barra := models.NuevaBarra(pixel.V(win.Bounds().Center().X, 50), barraSprite)
	pelota := models.NuevaPelota(win.Bounds().Center(), pelotaSprite)
	bloques := crearBloques(bloqueSprite)

	return &GameScene{
		win:     win,
		barra:   barra,
		pelota:  pelota,
		bloques: bloques,
	}, nil
}

func (gs *GameScene) ActualizarYDibujar() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		gs.barra.Actualizar(gs.win)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		gs.pelota.Actualizar(gs.win, gs.barra)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < len(gs.bloques); i++ {
			bloque := gs.bloques[i]
			if bloque.ColisionaConPelota(gs.pelota) {
				gs.pelota.Velocidad.Y = -gs.pelota.Velocidad.Y
				gs.bloques = append(gs.bloques[:i], gs.bloques[i+1:]...)
				i--
			}
		}
	}()

	wg.Wait()

	gs.barra.Dibujar(gs.win)
	gs.pelota.Dibujar(gs.win)
	for _, bloque := range gs.bloques {
		bloque.Dibujar(gs.win)
	}
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

func crearBloques(sprite *pixel.Sprite) []*models.Bloque {
	var bloques []*models.Bloque
	for x := 40.0; x < 800; x += 60 {
		for y := 500.0; y < 600; y += 30 {
			bloque := models.NuevoBloque(pixel.V(x, y), sprite)
			bloques = append(bloques, bloque)
		}
	}
	return bloques
}
