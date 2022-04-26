package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gomono"
	"golang.org/x/image/font/opentype"
)

func (g *Game) alwaysDraw(screen *ebiten.Image) {
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

	op.GeoM.Reset()
	rect := text.BoundString(g.face, fmt.Sprintf("%v", g.score.p1))
	op.GeoM.Translate(float64(160-rect.Bounds().Dx()-10), 15)
	text.DrawWithOptions(screen, fmt.Sprintf("%v", g.score.p1), g.face, op)

	op.GeoM.Reset()
	rect = text.BoundString(g.face, fmt.Sprintf("%v", g.score.p2))
	op.GeoM.Translate(float64(160+10), 15)
	text.DrawWithOptions(screen, fmt.Sprintf("%v", g.score.p2), g.face, op)

}

func (g *Game) drawWinner(screen *ebiten.Image) {
	rect := text.BoundString(g.face, fmt.Sprintf("P%v WINS!", g.winner))
	// ebitenutil.DrawRect(screen, 150, 100, 50, 50, color.Black)
	// ebitenutil.DrawRect(screen, float64(rect.Min.X), float64(rect.Min.Y), float64(rect.Dx()), float64(rect.Dy()), color.Black)
	text.Draw(screen, fmt.Sprintf("P%v WINS!", g.winner), g.face, 169-(rect.Bounds().Dx()/2), 100, color.RGBA{255, 128, 0, 255})
}

func (g *Game) font() {
	f, err := opentype.Parse(gomono.TTF)
	if err != nil {
		log.Fatalf("Parse: %v", err)
	}
	g.face, err = opentype.NewFace(f, &opentype.FaceOptions{
		Size:    15,
		DPI:     74,
		Hinting: font.HintingNone,
	})
	if err != nil {
		log.Fatalf("NewFace: %v", err)
	}
}
