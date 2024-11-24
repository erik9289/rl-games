package main

const NumLevels = 3

type Level [NumBlocksX][NumBlocksY]Block
type LevelCfg [NumBlocksY][NumBlocksX * 3]uint8

var levels [NumLevels]*Level
var levelCfgs = [NumLevels]LevelCfg{level1, level2, level3}

func initLevels() {
	for n := range NumLevels {
		levels[n] = configureLevel(levelCfgs[n])
	}
}

func configureLevel(lvlCfg LevelCfg) *Level {
	lvl := &Level{}
	for y := range NumBlocksY {
		for idx := 0; idx < NumBlocksX*3; idx += 3 {
			lvl[idx/3][y].color = BlockColor(lvlCfg[y][idx])
			lvl[idx/3][y].shields = int(lvlCfg[y][idx+1])
			lvl[idx/3][y].visible = (lvlCfg[y][idx+2] != 0)
		}
	}
	return lvl
}

func isLevelCleared(levelNr int) bool {
	for x := range NumBlocksX {
		for y := range NumBlocksY {
			if levels[levelNr][x][y].visible {
				return false
			}
		}
	}
	return true
}

var (
	level1 = LevelCfg{
		{3, 1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1},
		{3, 1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1},
		{2, 1, 1, 2, 1, 1, 2, 1, 1, 2, 1, 1, 2, 1, 1, 2, 1, 1, 2, 1, 1, 2, 1, 1, 2, 1, 1, 2, 1, 1},
		{2, 1, 1, 2, 1, 1, 2, 1, 1, 2, 1, 1, 2, 1, 1, 2, 1, 1, 2, 1, 1, 2, 1, 1, 2, 1, 1, 2, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{0, 1, 1, 0, 1, 1, 0, 1, 1, 0, 1, 1, 4, 1, 1, 5, 1, 1, 6, 1, 1, 0, 1, 1, 0, 1, 1, 0, 1, 1},
		{0, 1, 1, 0, 1, 1, 0, 1, 1, 0, 1, 1, 4, 1, 1, 5, 1, 1, 6, 1, 1, 0, 1, 1, 0, 1, 1, 0, 1, 1},
	}

	level2 = LevelCfg{
		{3, 1, 1, 3, 1, 1, 2, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 0, 1, 1, 3, 1, 1, 3, 1, 1},
		{3, 1, 1, 3, 1, 1, 2, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 0, 1, 1, 3, 1, 1, 3, 1, 1},
		{3, 1, 1, 3, 1, 1, 2, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 0, 1, 1, 3, 1, 1, 3, 1, 1},
		{3, 1, 1, 3, 1, 1, 2, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 0, 1, 1, 3, 1, 1, 3, 1, 1},
		{3, 1, 1, 3, 1, 1, 2, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 0, 1, 1, 3, 1, 1, 3, 1, 1},
		{3, 1, 1, 3, 1, 1, 2, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 0, 1, 1, 3, 1, 1, 3, 1, 1},
		{3, 1, 1, 3, 1, 1, 2, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 0, 1, 1, 3, 1, 1, 3, 1, 1},
		{3, 1, 1, 3, 1, 1, 2, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 0, 1, 1, 3, 1, 1, 3, 1, 1},
	}

	level3 = LevelCfg{
		{3, 1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1},
		{3, 1, 1, 2, 1, 1, 2, 1, 1, 2, 1, 1, 2, 1, 1, 2, 1, 1, 2, 1, 1, 2, 1, 1, 2, 1, 1, 3, 1, 1},
		{3, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 3, 1, 1},
		{3, 1, 1, 2, 1, 1, 1, 1, 1, 0, 1, 1, 0, 1, 1, 0, 1, 1, 0, 1, 1, 1, 1, 1, 2, 1, 1, 3, 1, 1},
		{3, 1, 1, 2, 1, 1, 1, 1, 1, 0, 1, 1, 0, 1, 1, 0, 1, 1, 0, 1, 1, 1, 1, 1, 2, 1, 1, 3, 1, 1},
		{3, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 3, 1, 1},
		{3, 1, 1, 2, 1, 1, 2, 1, 1, 2, 1, 1, 2, 1, 1, 2, 1, 1, 2, 1, 1, 2, 1, 1, 2, 1, 1, 3, 1, 1},
		{3, 2, 1, 3, 2, 1, 3, 2, 1, 3, 2, 1, 3, 2, 1, 3, 2, 1, 3, 2, 1, 3, 2, 1, 3, 2, 1, 3, 2, 1},
	}
)
