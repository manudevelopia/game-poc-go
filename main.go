package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	_ "image/png"
	"log"
)

const (
	screenWidth       = 640
	screenHeight      = 480
	screenWidthLimit  = 620
	screenHeightLimit = 440
)

func loadImage(filename string) *ebiten.Image {
	var err error
	img, _, err := ebitenutil.NewImageFromFile("img/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	return img
}

type Game struct {
	bowser Character
	count  int
}

const (
	frameOX     = 0
	frameOY     = 68
	frameWidth  = 71
	frameHeight = 68
	frameNum    = 4
)

func (game *Game) Update() error {
	game.count++
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) && game.bowser.x < screenWidthLimit {
		game.bowser.moveRight()
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && game.bowser.x > 0 {
		game.bowser.moveLeft()
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowUp) && game.bowser.y > 0 {
		game.bowser.moveUp()
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowDown) && game.bowser.y < screenHeightLimit {
		game.bowser.moveDown()
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		game.bowser.changeFuryStatus()
	}
	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(game.bowser.x), float64(game.bowser.y))

	i := (game.count / 5) % frameNum
	sx, sy := frameOX+i*frameWidth, frameOY*game.bowser.state
	screen.DrawImage(game.bowser.img.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{
		Character{0, screenHeightLimit, 3, WalkingUp, false,
			loadImage("bowser_tile.png"),
			loadImage("bowser_fury.png")}, 0}); err != nil {
		log.Fatal(err)
	}
}
