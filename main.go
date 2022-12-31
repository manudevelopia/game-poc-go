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

	frameOX           = 0
	frameBowserOY     = 68
	frameBowserWidth  = 71
	frameBowserHeight = 68
	framePeachOY      = 68
	framePeachWidth   = 71
	framePeachHeight  = 68
	frameNum          = 4
)

type Game struct {
	bowser Character
	peach  Character
	count  int
}

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
	if ebiten.IsKeyPressed(ebiten.KeyD) && game.peach.x < screenWidthLimit {
		game.peach.moveRight()
	} else if ebiten.IsKeyPressed(ebiten.KeyA) && game.peach.x > 0 {
		game.peach.moveLeft()
	} else if ebiten.IsKeyPressed(ebiten.KeyW) && game.peach.y > 0 {
		game.peach.moveUp()
	} else if ebiten.IsKeyPressed(ebiten.KeyS) && game.peach.y < screenHeightLimit {
		game.peach.moveDown()
	}
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		game.peach.changeFuryStatus()
	}
	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	if checkCollision(game.peach, game.bowser) {
		ebitenutil.DebugPrint(screen, "I gotcha you")
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(game.bowser.x), float64(game.bowser.y))
	op2 := &ebiten.DrawImageOptions{}
	op2.GeoM.Translate(float64(game.peach.x), float64(game.peach.y))

	i := (game.count / 5) % frameNum
	sx, sy := frameOX+i*frameBowserWidth, frameBowserOY*game.bowser.state
	screen.DrawImage(game.bowser.img.SubImage(image.Rect(sx, sy, sx+frameBowserWidth, sy+frameBowserHeight)).(*ebiten.Image), op)
	sx2, sy2 := frameOX+i*framePeachWidth, framePeachOY*game.peach.state
	screen.DrawImage(game.peach.img.SubImage(image.Rect(sx2, sy2, sx2+framePeachWidth, sy2+framePeachHeight)).(*ebiten.Image), op2)
}

func checkCollision(one Character, two Character) bool {
	collisionX := one.x+71 >= two.x && two.x+71 >= one.x
	collisionY := one.y+68 >= two.y && two.y+68 >= one.y
	return collisionX && collisionY
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
			loadImage("bowser_fury.png")},
		Character{0, screenHeightLimit, 3, WalkingUp, false,
			loadImage("peach_tile.png"),
			loadImage("bowser_fury.png")}, 0}); err != nil {
		log.Fatal(err)
	}
}
