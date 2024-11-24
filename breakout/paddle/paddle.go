package paddle

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	// Default or constant values
	paddleWidth  = 50
	paddleHeight = 6
	paddlePosY   = 260 // Y-position on the screen
	paddleSpeed  = 200 // Speed in pixels per second
)

type Paddle struct {
	X        float32
	Y        float32
	Width    float32
	Height   float32
	Speed    float32
	Velocity float32
	Color    rl.Color
	rect     rl.Rectangle
}

func NewPaddle() *Paddle {
	p := &Paddle{}
	p.Y = paddlePosY
	p.Width = paddleWidth
	p.Height = paddleHeight
	p.Speed = paddleSpeed
	p.Color = rl.Color{50, 150, 90, 255}
	return p
}

func (p *Paddle) UpdatePosition(dt float32, screenSize float32) {
	p.X += p.Velocity * dt
	p.X = rl.Clamp(p.X, 0, screenSize-p.Width)
	p.rect = rl.Rectangle{X: p.X, Y: p.Y, Width: p.Width, Height: p.Height}
	p.Velocity = 0 // reset
}

func (p *Paddle) Draw() {
	rl.DrawRectangleRec(p.rect, p.Color)
}

// CheckCollision checks if a ball collides with the paddle, it returns the collisionNormal vector
// The returned collisionNormal vector is zero (rl.Vector2Zero) if no collision
func (p *Paddle) CheckCollision(ballPos, previousBallPos rl.Vector2, ballRadius float32) rl.Vector2 {
	// Check for collision between ball and paddle
	if rl.CheckCollisionCircleRec(ballPos, ballRadius, p.rect) {
		var collisionNormal rl.Vector2
		if previousBallPos.Y < p.rect.Y+p.rect.Height {
			collisionNormal = rl.Vector2Add(collisionNormal, rl.Vector2{X: 0, Y: -1})
			ballPos.Y = p.rect.Y - ballRadius
		}
		// In case the ball hits the bottom of the paddle
		if previousBallPos.Y > p.rect.Y+p.rect.Height {
			collisionNormal = rl.Vector2Add(collisionNormal, rl.Vector2{X: 0, Y: 1})
			ballPos.Y = p.rect.Y + p.rect.Height + ballRadius
		}
		// From the left of the paddle
		if previousBallPos.X < p.rect.X {
			collisionNormal = rl.Vector2Add(collisionNormal, rl.Vector2{X: -1, Y: 0})
		}
		// From the right of the paddle
		if previousBallPos.X > p.rect.X+p.rect.Width {
			collisionNormal = rl.Vector2Add(collisionNormal, rl.Vector2{X: 0, Y: -1})
		}
		return collisionNormal
	}
	return rl.Vector2Zero()
}
