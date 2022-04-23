package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	p1   float32
	p2   float32
	ball struct{ x, y float32 }

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

	game.p1sprite = ebiten.NewImage(5, 30)
	game.p1sprite.Fill(color.RGBA{255, 255, 255, 255})
	game.p2sprite = ebiten.NewImage(5, 30)
	game.p2sprite.Fill(color.RGBA{255, 255, 255, 255})
	game.ballsprite = ebiten.NewImage(5, 5)
	game.ballsprite.Fill(color.RGBA{255, 255, 255, 255})
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{200, 30, 30, 100})
	img := ebiten.NewImage(10, 200)
	img.Fill(color.RGBA{255, 255, 255, 255})
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(155, 0)
	screen.DrawImage(img, op)

	op.GeoM.Translate(-155, 50)
	screen.DrawImage(g.p1sprite, op)

	op.GeoM.Translate(315, 50)
	screen.DrawImage(g.p2sprite, op)

	op.GeoM.Reset()
	op.GeoM.Translate(50, 100)
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
