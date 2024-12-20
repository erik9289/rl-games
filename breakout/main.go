package main

import (
	"breakout/paddle"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	Title        = "Breakout!"
	screenWidth  = 640
	screenHeight = 640
	screenSize   = 320
	FPS          = 60

	maxLives = 3

	ballSpeed  = 260
	ballRadius = 4
	ballStartY = 160
)

var (
	ballPos rl.Vector2
	ballDir rl.Vector2

	started   bool
	gameOver  bool
	score     int
	highscore int
	numLives  = maxLives

	levelCurrent int
	levelCnt     int

	gameOverSnd  rl.Sound
	hitPaddleSnd rl.Sound
	hitBlockSnd  rl.Sound

	livesImg rl.Texture2D

	// Colors
	bgColor   = rl.Color{150, 190, 220, 255}
	ballColor = rl.Color{200, 90, 20, 255}
)

var dt float32

func main() {

	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.InitWindow(screenWidth, screenHeight, Title)
	defer rl.CloseWindow()
	// rl.DisableBackfaceCulling()

	// Load assets
	rl.InitAudioDevice()
	defer rl.CloseAudioDevice()
	hitPaddleSnd = rl.LoadSound("assets/hit_paddle.wav")
	gameOverSnd = rl.LoadSound("assets/game_over.wav")
	hitBlockSnd = rl.LoadSound("assets/hit_block.wav")
	livesImg = rl.LoadTexture("assets/heart_32.png")

	rl.SetTargetFPS(FPS)

	initLevels()

	paddle := paddle.NewPaddle()

	restart(true, paddle)

	for !rl.WindowShouldClose() {
		// Update
		//---------------------------------------------------------------------

		// Keep dt zero ('wait') until SPACEBAR is pressed to start the game
		if !started {
			ballPos = rl.Vector2{
				X: screenSize/2 + float32(math.Cos(rl.GetTime()))*screenSize/2.5,
				Y: ballStartY,
			}

			if rl.IsKeyPressed(rl.KeySpace) {
				// Point the ball (vector) to the middle of the paddle
				paddle_middle := rl.Vector2{X: paddle.X + paddle.Width/2, Y: paddle.Y}
				ball_to_paddle := rl.Vector2Subtract(paddle_middle, ballPos)
				ballDir = rl.Vector2Normalize(ball_to_paddle) // Normalize the direction vector to 1
				started = true
			}
		} else if gameOver {
			restart(true, paddle)
			if rl.IsKeyPressed(rl.KeySpace) {
				started = true
			}
		}

		dt = rl.GetFrameTime()

		previous_ball_pos := ballPos
		ballPos = rl.Vector2Add(ballPos, rl.Vector2Scale(ballDir, ballSpeed*dt))

		// Check right wall and bounce
		if ballPos.X+ballRadius > screenSize {
			ballPos.X = screenSize - ballRadius
			ballDir = reflect(ballDir, rl.Vector2{X: -1, Y: 0})
		}
		// Check left wall and bounce
		if ballPos.X-ballRadius < 0 {
			ballPos.X = 0 + ballRadius
			ballDir = reflect(ballDir, rl.Vector2{X: 1, Y: 0})
		}
		// Check top wall and bounce
		if ballPos.Y-ballRadius < 0 {
			ballPos.Y = ballRadius
			ballDir = reflect(ballDir, rl.Vector2{X: 0, Y: 1})
		}
		// Check bottom, this means game over/restart
		if ballPos.Y+ballRadius*6 > screenSize {
			numLives -= 1
			if numLives == 0 {
				rl.PlaySound(gameOverSnd)
				gameOver = true
			}
			restart(false, paddle)
		}

		// Handle key press
		if rl.IsKeyDown(rl.KeyLeft) {
			paddle.Velocity -= paddle.Speed
		}
		if rl.IsKeyDown(rl.KeyRight) {
			paddle.Velocity += paddle.Speed
		}
		paddle.UpdatePosition(dt, screenSize)

		// Check for collision between ball and paddle
		collisionNormal := paddle.CheckCollision(ballPos, previous_ball_pos, ballRadius)
		// Apply the accumulated collision_normal and calculate the reflection
		if rl.Vector2Length(collisionNormal) != 0 {
			ballDir = reflect(ballDir, collisionNormal)
			rl.PlaySound(hitPaddleSnd)
		}

		// Check for collision with a block
		checkBlockCollision(previous_ball_pos)

		// Draw
		//---------------------------------------------------------------------
		camera := rl.Camera2D{Zoom: float32(rl.GetScreenHeight() / screenSize)}

		rl.BeginDrawing()
		rl.ClearBackground(bgColor)
		rl.BeginMode2D(camera)

		paddle.Draw()
		rl.DrawCircleV(ballPos, ballRadius, ballColor)
		DrawBlocks()
		drawUI()

		rl.EndMode2D()
		rl.EndDrawing()
	}
}

func restart(reset bool, paddle *paddle.Paddle) {
	paddle.X = screenSize/2 - paddle.Width/2
	ballPos = rl.Vector2{X: screenSize / 2, Y: ballStartY}
	started = false

	// Reset the blocks if no lives left or at the start
	if reset {
		numLives = maxLives
		score = 0
		gameOver = false
		levelCurrent = 0
		levelCnt = 0

		initLevels()
	}
}

func reflect(dir, normal rl.Vector2) rl.Vector2 {
	newDirection := rl.Vector2Reflect(dir, rl.Vector2Normalize(normal))
	return rl.Vector2Normalize(newDirection)
}
