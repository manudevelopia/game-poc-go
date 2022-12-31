package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

func loadImage(filename string) *ebiten.Image {
	var err error
	img, _, err := ebitenutil.NewImageFromFile("img/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	return img
}
