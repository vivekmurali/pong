package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	p1   float64
	p2   float64
	ball struct{ x, y float64 }

	p1sprite   *ebiten.Image
	p2sprite   *ebiten.Image
	ballsprite *ebiten.Image
}

var game *Game

func init() {
	game = &Game{}
	game.p1 = 160
	game.p2 = 160
	game.ball.x = 160
	game.ball.y = 100

	// p1 paddle
	game.p1sprite = ebiten.NewImage(5, 30)
	game.p1sprite.Fill(color.RGBA{255, 255, 255, 255})

	// p2 paddle
	game.p2sprite = ebiten.NewImage(5, 30)
	game.p2sprite.Fill(color.RGBA{255, 255, 255, 255})

	// ball
	game.ballsprite = ebiten.NewImage(5, 5)
	game.ballsprite.Fill(color.RGBA{255, 255, 0, 255})
}

func (g *Game) Update() error {

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

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{200, 30, 30, 100})
	img := ebiten.NewImage(2, 200)
	img.Fill(color.RGBA{255, 255, 255, 255})
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(159, 0)
	screen.DrawImage(img, op)

	op.GeoM.Reset()
	op.GeoM.Translate(0, g.p1)
	screen.DrawImage(g.p1sprite, op)

	op.GeoM.Reset()
	op.GeoM.Translate(315, g.p2)
	screen.DrawImage(g.p2sprite, op)

	op.GeoM.Reset()
	op.GeoM.Translate(g.ball.x, g.ball.y)
	screen.DrawImage(g.ballsprite, op)
	// ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 200
}

func main() {
	ebiten.SetWindowSize(640, 400)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
