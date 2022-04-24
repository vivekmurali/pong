package main

func (g *Game) reset() {
	g.running = false
	g.p1 = 85
	g.p2 = 0
	g.ball.x = 157.5
	g.ball.y = 97.5

	g.ball.vx = 3
	g.ball.vy = 0

	g.score.p1 = 0
	g.score.p2 = 0

	g.winner = 0
}

func (g *Game) resetWithoutScore() {
	g.running = false
	g.p1 = 85
	g.p2 = 0
	g.ball.x = 157.5
	g.ball.y = 97.5

	g.ball.vx = 3
	g.ball.vy = 0
}
