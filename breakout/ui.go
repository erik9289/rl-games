package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Define some constant for fontsizes
const (
	FS_SCORE     = 24
	FS_HIGHSCORE = 8
	FS_GAMEOVER  = 24
	FS_RESTART   = 12
	CenterPosY   = 260 // Same as default paddle Y position
)

// draw_ui draw the UI components like:
//   - the score, high score and number of lives left
//   - 'Game Over' notification
func drawUI() {
	// Display num_lives (as heart icons) in the upper left corner
	for i := range numLives {
		rl.DrawTextureEx(lives_img, rl.Vector2{X: float32(5 + (i * 14)), Y: 5}, 0, 0.4, rl.White)
	}

	// Display the score in the center of the screen
	score_text := fmt.Sprintf("%03d", score)
	rl.DrawText(score_text, centerText(score_text, FS_SCORE, screenSize), 5, FS_SCORE, rl.White)

	// Display highscore in upper right corner
	highscore_text := fmt.Sprintf("High: %03d", highscore)
	highscore_text_width := rl.MeasureText(highscore_text, FS_HIGHSCORE)
	rl.DrawText(highscore_text, screenSize-highscore_text_width-5, 5, FS_HIGHSCORE, rl.White)

	// Display 'Game Over'
	if gameOver {
		game_over_text := fmt.Sprint("Game Over")
		rl.DrawText(
			game_over_text,
			centerText(game_over_text, FS_GAMEOVER, screenSize),
			CenterPosY-60,
			FS_GAMEOVER,
			rl.Red,
		)
		game_over_restart_text := fmt.Sprint("SPACE to restart")
		rl.DrawText(
			game_over_restart_text,
			centerText(game_over_restart_text, FS_RESTART, screenSize),
			CenterPosY-30,
			FS_RESTART,
			rl.White,
		)
	}
}

func centerText(text string, font_size, screen_size int) int32 {
	text_width := rl.MeasureText(text, int32(font_size))
	return int32(screen_size/2) - text_width/2
}
