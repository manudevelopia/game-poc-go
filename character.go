package main

import "github.com/hajimehoshi/ebiten/v2"

type Character struct {
	x        int
	y        int
	speed    int
	fury     bool
	img      *ebiten.Image
	fury_img *ebiten.Image
}

func (character Character) moveRight() {
	character.x += character.speed
}

func (character Character) moveLeft() {
	character.x -= character.speed
}

func (character Character) moveUp() {
	character.y -= character.speed
}

func (character Character) moveDown() {
	character.y += character.speed
}

func (character Character) changeFuryStatus() {
	character.fury = !character.fury
}
