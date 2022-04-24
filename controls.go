package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) control() {
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		if g.p1 >= 2.5 {
			g.p1 -= 2.5
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyZ) {
		if g.p1 <= 167.5 {
			g.p1 += 2.5
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyK) {
		if g.p2 >= 2.5 {
			g.p2 -= 2.5
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyM) {
		if g.p2 <= 167.5 {
			g.p2 += 2.5
		}
	}

	//TODO: fix pause and play
	// if ebiten.IsKeyPressed(ebiten.KeySpace) {
	// 	// if t := inpututil.KeyPressDuration(ebiten.KeySpace); t > 5 && t < 10 {
	// 	// 	g.running = false
	// 	// }

	// 	if !repeatingKeyPressed(ebiten.KeySpace) {
	// 		g.running = false
	// 	}
	// }
	if inpututil.IsKeyJustReleased(ebiten.KeySpace) {
		g.running = false
	}

}

// repeatingKeyPressed return true when key is pressed considering the repeat state.
func repeatingKeyPressed(key ebiten.Key) bool {
	const (
		delay    = 30
		interval = 3
	)
	d := inpututil.KeyPressDuration(key)
	if d == 1 {
		return true
	}
	if d >= delay && (d-delay)%interval == 0 {
		return true
	}
	return false
}
