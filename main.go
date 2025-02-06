package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

const screenWidth, screenHeight = 1024, 720

// Game struct
type Game struct{}

// Update handles game logic (empty for now)
func (g *Game) Update() error {
	return nil
}

// Draw renders the frame
func (g *Game) Draw(screen *ebiten.Image) {
	// Set text position
	x, y := 50, 150

	// Draw text using the built-in font
	text.Draw(screen, "Hello, Godlike!! sasamba", basicfont.Face7x13, x, y, color.White)
}

// Layout sets the game window size
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Basic Ebiten Text")

	// Run the game
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
