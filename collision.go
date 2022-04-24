package main

func (g *Game) collision() {
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
}
