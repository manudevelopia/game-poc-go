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

	frameOX          = 0
	framePlayerOY     = 68
	framePlayerWidth  = 71
	framePlayerHeight = 68
	frameNum          = 4
)

type Game struct {
	players []Character
	count   int
}

func (game *Game) Update() error {
	game.count++
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) && game.players[0].x < screenWidthLimit {
		game.players[0].moveRight()
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && game.players[0].x > 0 {
		game.players[0].moveLeft()
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowUp) && game.players[0].y > 0 {
		game.players[0].moveUp()
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowDown) && game.players[0].y < screenHeightLimit {
		game.players[0].moveDown()
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		game.players[0].changeFuryStatus()
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) && game.players[1].x < screenWidthLimit {
		game.players[1].moveRight()
	} else if ebiten.IsKeyPressed(ebiten.KeyA) && game.players[1].x > 0 {
		game.players[1].moveLeft()
	} else if ebiten.IsKeyPressed(ebiten.KeyW) && game.players[1].y > 0 {
		game.players[1].moveUp()
	} else if ebiten.IsKeyPressed(ebiten.KeyS) && game.players[1].y < screenHeightLimit {
		game.players[1].moveDown()
	}
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		game.players[1].changeFuryStatus()
	}
	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	if checkCollision(game.players[1], game.players[0]) {
		ebitenutil.DebugPrint(screen, "I gotcha you")
	}
	frameSpot := (game.count / 5) % frameNum
	for i := 0; i < len(game.players); i++ {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(game.players[i].x), float64(game.players[i].y))
		sx, sy := frameOX+frameSpot*framePlayerWidth, framePlayerOY*game.players[i].state
		screen.DrawImage(game.players[i].img.SubImage(image.Rect(sx, sy, sx+framePlayerWidth, sy+framePlayerHeight)).(*ebiten.Image), op)
	}
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
		[]Character{
			{0, screenHeightLimit, 3, WalkingUp, false, loadImage("bowser_tile.png"), loadImage("bowser_fury.png")},
			{0, screenHeightLimit, 3, WalkingUp, false, loadImage("peach_tile.png"), loadImage("bowser_fury.png")},
		}, 0}); err != nil {
		log.Fatal(err)
	}
}
