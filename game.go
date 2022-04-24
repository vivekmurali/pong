package main

func (g *Game) countScore() {
	if g.ball.x >= 315 {
		g.score.p1 += 1
		if g.score.p1 >= 5 {
			g.winner = 1
			g.running = false
			return
		}
	}

	if g.ball.x <= 0 {
		g.score.p2 += 1
		if g.score.p2 >= 5 {
			g.winner = 2
			g.running = false
			return
		}
	}

}
