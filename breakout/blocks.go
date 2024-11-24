package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	NumBlocksX  = 10
	NumBlocksY  = 8
	BlockWidth  = 28
	BlockHeight = 10
)

type Block struct {
	color   BlockColor
	shields int
	visible bool
}

type BlockColor int

const (
	Yellow BlockColor = iota
	Green
	Orange
	Red
	Gold
	Silver
	Bronze
)

var blockColorValues = []rl.Color{
	{253, 249, 150, 255}, // Yellow
	{180, 245, 190, 255}, // Green
	{170, 120, 250, 255}, // Orange
	{250, 90, 85, 255},   // Red
	{253, 208, 23, 255},  // Gold
	{192, 192, 192, 255}, // Silver
	{205, 127, 50, 255},  // Bronze
}

var blockColorScore = []int{
	2,  // Yellow
	4,  // Green
	6,  // Orange
	8,  // Red
	20, // Gold
	15, // Silver
	10, // Bronze
}

func calcBlockRect(x, y int) rl.Rectangle {
	return rl.Rectangle{
		X:      float32(20 + x*BlockWidth),
		Y:      float32(40 + y*BlockHeight),
		Width:  BlockWidth - 1,
		Height: BlockHeight - 1,
	}
}

func blockExists(x, y int) bool {
	if x < 0 || y < 0 || x >= NumBlocksX || y >= NumBlocksY {
		return false
	}
	return levels[level_current][x][y].visible
}

// Check for collisions with blocks
func checkBlockCollision(previousBallPos rl.Vector2) {
block_x_loop:
	for x := 0; x < NumBlocksX; x++ {
		for y := 0; y < NumBlocksY; y++ {
			if !levels[level_current][x][y].visible {
				continue
			}
			blockRect := calcBlockRect(x, y)
			if rl.CheckCollisionCircleRec(ball_pos, ballRadius, blockRect) {
				var collisionNormal = rl.Vector2{}
				// Ball is above block
				if previousBallPos.Y < blockRect.Y {
					collisionNormal = rl.Vector2Add(collisionNormal, rl.Vector2{X: 0, Y: -1})
				}
				// Ball is below the blocks
				if previousBallPos.Y > blockRect.Y+blockRect.Height {
					collisionNormal = rl.Vector2Add(collisionNormal, rl.Vector2{X: 0, Y: 1})
				}
				// Ball is on the left side of a blocks
				if previousBallPos.X < blockRect.X {
					collisionNormal = rl.Vector2Add(collisionNormal, rl.Vector2{X: -1, Y: 0})
				}
				// Ball is on the right sidde of a blocks
				if previousBallPos.X > blockRect.X+blockRect.Width {
					collisionNormal = rl.Vector2Add(collisionNormal, rl.Vector2{X: 1, Y: 0})
				}

				// Check if there where blocks left or right of current blocks by checking
				// the collsion_normal. This prevents 'horizontal' reflections when hitting a corner
				if blockExists(x+int(collisionNormal.X), y) {
					collisionNormal.X = 0
				}
				// Also for above and beneath
				if blockExists(x, y+int(collisionNormal.Y)) {
					collisionNormal.Y = 0
				}

				// Apply the accumulated collision_normal and calculate the reflection
				if rl.Vector2Length(collisionNormal) != 0 {
					ball_dir = reflect(ball_dir, collisionNormal)
				}

				// Now lower the shield or destroy the block!
				rl.SetSoundPitch(hitBlockSnd, rl.Vector2Length(collisionNormal)*0.8)
				rl.PlaySound(hitBlockSnd)
				levels[level_current][x][y].shields -= 1
				if levels[level_current][x][y].shields < 1 {
					levels[level_current][x][y].visible = false
				}

				// Update the score based on block row_colors
				blockColor := levels[level_current][x][y].color
				score += blockColorScore[blockColor]
				if score > highscore {
					highscore = score
				}

				// Check if all blocks have been cleared, then go to next Level
				if isLevelCleared(level_current) {
					fmt.Printf("* cleared level! level_current = %d\n", level_current)
					level_current = (level_current + 1) % NumLevels
					if level_current == 0 {
						// We cycled throug all available levels, reset the levels before we continue

						initLevels()
					}
					level_cnt += 1
					fmt.Printf(
						"updated level_current: level_current = %d, level_cnt=%d, NUM_LEVELS=%d\n",
						level_current,
						level_cnt,
						NumLevels,
					)
				}
				break block_x_loop // Breaking outer loop, preventing multiple collsions per frame
			}
		}
	}

}

func DrawBlocks() {
	for x := 0; x < NumBlocksX; x++ {
		for y := 0; y < NumBlocksY; y++ {
			if !levels[level_current][x][y].visible {
				continue // Skip blocks that are hit
			}
			block_rect := calcBlockRect(x, y)

			// rl.DrawRectangleRec(block_rect, block_color_values[row_colors[y]])
			rl.DrawRectangleRec(block_rect, blockColorValues[levels[level_current][x][y].color])
			top_left := rl.Vector2{X: block_rect.X, Y: block_rect.Y}
			top_right := rl.Vector2{X: block_rect.X + block_rect.Width, Y: block_rect.Y}
			bottom_left := rl.Vector2{X: block_rect.X, Y: block_rect.Y + block_rect.Height}
			bottom_right := rl.Vector2{
				X: block_rect.X + block_rect.Width,
				Y: block_rect.Y + block_rect.Height,
			}
			rl.DrawLineEx(top_left, top_right, 1, rl.Color{255, 255, 150, 100})
			rl.DrawLineEx(top_left, bottom_left, 1, rl.Color{255, 255, 150, 100})
			rl.DrawLineEx(bottom_left, bottom_right, 1, rl.Color{0, 0, 50, 100})
			rl.DrawLineEx(top_right, bottom_right, 1, rl.Color{0, 0, 50, 100})
		}
	}
}
