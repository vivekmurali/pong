package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/font"
)

type Game struct {
	p1   float64
	p2   float64
	ball struct{ x, y, vx, vy float64 }

	running bool

	score struct{ p1, p2 int }

	winner uint8

	p1sprite   *ebiten.Image
	p2sprite   *ebiten.Image
	ballsprite *ebiten.Image

	face font.Face
}

var game *Game

func init() {
	game = &Game{}
	game.reset()
	game.font()

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

	if g.winner > 0 {
		g.running = false
		if inpututil.IsKeyJustReleased(ebiten.KeySpace) {
			g.reset()
			g.running = true
		}
		return nil
	}
	if !g.running {
		if inpututil.IsKeyJustReleased(ebiten.KeySpace) {
			g.running = true
		}
		return nil
	}
	g.ball.x += g.ball.vx
	g.ball.y += g.ball.vy

	if g.ball.x >= 315 || g.ball.x <= 0 {
		g.countScore()
		g.resetWithoutScore()

		//TODO set winner

	}

	g.collision()

	g.control()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.alwaysDraw(screen)
	if g.winner >= 1 {
		g.drawWinner(screen)
	}
	// op := &ebiten.DrawImageOptions{}
	// rect := text.BoundString(g.face, fmt.Sprintf("P%v WINS!", g.winner))
	// op.GeoM.Translate(float64(160-(rect.Bounds().Dx())), 100)
	// text.DrawWithOptions(screen, fmt.Sprintf("P%v WINS!", g.winner), g.face, op)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("p1: %v; p2: %v", g.score.p1, g.score.p2))
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
