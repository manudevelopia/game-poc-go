package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/png"
	"log"
)

var bowser_img *ebiten.Image
var bowser_fury_img *ebiten.Image

const (
	screenWidth       = 640
	screenHeight      = 480
	screenWidthLimit  = 600
	screenHeightLimit = 440
)

func init() {
	var err error
	bowser_img, _, err = ebitenutil.NewImageFromFile("img/bowser.png")
	bowser_fury_img, _, err = ebitenutil.NewImageFromFile("img/bowser_fury.png")
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct {
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) && bowser_x < screenWidthLimit {
		bowser_x += bowser_speed
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && bowser_x > 0 {
		bowser_x -= bowser_speed
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowUp) && bowser_y > 0 {
		bowser_y -= bowser_speed
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowDown) && bowser_y < screenHeightLimit {
		bowser_y += bowser_speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		bowser_fury = !bowser_fury
	}
	return nil
}

var bowser_x = 0
var bowser_y = 0
var bowser_speed = 3
var bowser_fury = false

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(bowser_x), float64(bowser_y))

	if bowser_fury == true {
		screen.DrawImage(bowser_fury_img, op)
	} else {
		screen.DrawImage(bowser_img, op)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
