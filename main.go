package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	p1   float64
	p2   float64
	ball struct{ x, y, vx, vy float64 }

	running bool

	winner uint8

	p1sprite   *ebiten.Image
	p2sprite   *ebiten.Image
	ballsprite *ebiten.Image
}

var game *Game

func init() {
	game = &Game{}
	game.running = false
	game.p1 = 85
	game.p2 = 67.5
	game.ball.x = 157.5
	game.ball.y = 97.5

	game.ball.vx = 3
	game.ball.vy = 0

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

	if !g.running {
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			g.running = true
		}
		return nil
	}
	g.ball.x += g.ball.vx
	g.ball.y += g.ball.vy

	if g.ball.x >= 315 || g.ball.x <= 0 {
		g.running = false
		g.ball.x = 157.5
		g.ball.y = 97.5
		g.ball.vy = 0
		g.p1 = 85
		g.p2 = 95

		//TODO set winner
	}

	if g.ball.x >= 310 && (g.ball.y <= g.p2+30 && g.ball.y >= g.p2-5) {

		g.ball.vx *= -1

		relInter := g.p2 + 15 - (g.ball.y + 2.5)
		g.ball.vy = (relInter / 17.5) * -1.5

		if g.ball.vy > 1.5 {
			g.ball.vy = 1.5
		}
		if g.ball.vy < -1.5 {
			g.ball.vy = -1.5
		}
	}

	if g.ball.x <= 5 && (g.ball.y <= g.p1+30 && g.ball.y >= g.p1-5) {
		g.ball.vx *= -1

		relInter := g.p1 + 15 - (g.ball.y + 2.5)
		g.ball.vy = (relInter / 17.5) * -1.5

		if g.ball.vy > 1.5 {
			g.ball.vy = 1.5
		}
		if g.ball.vy < -1.5 {
			g.ball.vy = -1.5
		}
	}

	if g.ball.y <= 0 || g.ball.y >= 195 {
		g.ball.vy *= -1
	}

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

	// if ebiten.IsKeyPressed(ebiten.KeySpace) {
	// 	g.running = false
	// }

	if t := inpututil.KeyPressDuration(ebiten.KeySpace); t > 2 && t < 10 {
		g.running = false
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
	ebitenutil.DebugPrint(screen, fmt.Sprintf("ball.vy = %v; ball.y = %v;  p2.y = %v", g.ball.vy, g.ball.y, g.p2))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 200
}

func main() {
	ebiten.SetWindowSize(640, 400)
	ebiten.SetWindowTitle("PONG")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
