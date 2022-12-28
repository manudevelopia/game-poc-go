package main

import "github.com/hajimehoshi/ebiten/v2"

type Character struct {
	x       int
	y       int
	speed   int
	state   int
	fury    bool
	img     *ebiten.Image
	furyImg *ebiten.Image
}

const (
	WalkingDown = iota
	WalkingLeft
	WalkingRight
	WalkingUp
)

func (character *Character) moveRight() {
	character.state = WalkingRight
	character.x += character.speed
}

func (character *Character) moveLeft() {
	character.state = WalkingLeft
	character.x -= character.speed
}

func (character *Character) moveUp() {
	character.state = WalkingUp
	character.y -= character.speed
}

func (character *Character) moveDown() {
	character.state = WalkingDown
	character.y += character.speed
}

func (character *Character) changeFuryStatus() {
	character.fury = !character.fury
}
