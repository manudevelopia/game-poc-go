package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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
}

func (game *Game) Update() error {
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
	if game.bowser.fury == true {
		screen.DrawImage(game.bowser.fury_img, op)
	} else {
		screen.DrawImage(game.bowser.img, op)
	}
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{
		Character{0, screenHeightLimit, 3, false,
			loadImage("bowser.png"), loadImage("bowser_fury.png")}}); err != nil {
		log.Fatal(err)
	}
}
