package utils

func IsAlive(grid [][]rune, h, w int, livingChar rune, edge bool) bool {
	livingCellsCount := 0
	lenW := len(grid[0])
	lenH := len(grid)
	if edge {
		if grid[(h+1)%lenH][w] == livingChar {
			livingCellsCount++
		}

		if grid[h][(w+1)%lenW] == livingChar {
			livingCellsCount++
		}

		if grid[h][(w-1+lenW)%lenW] == livingChar {
			livingCellsCount++
		}

		if grid[(h-1+lenH)%lenH][w] == livingChar {
			livingCellsCount++
		}

		if grid[(h-1+lenH)%lenH][(w-1+lenW)%lenW] == livingChar {
			livingCellsCount++
		}

		if grid[(h-1+lenH)%lenH][(w+1)%lenW] == livingChar {
			livingCellsCount++
		}

		if grid[(h+1)%lenH][(w-1+lenW)%lenW] == livingChar {
			livingCellsCount++
		}

		if grid[(h+1)%lenH][(w+1)%lenW] == livingChar {
			livingCellsCount++
		}
	} else {
		direction := [][]int{
			{1, 1}, {0, 1}, {1, 0}, {-1, 1}, {1, -1}, {-1, 0}, {0, -1}, {-1, -1},
		}

		for i := 0; i < 8; i++ {
			curH := h + direction[i][0]
			curW := w + direction[i][1]

			if curH >= 0 && curH < lenH && curW >= 0 && curW < lenW && grid[curH][curW] == livingChar {
				livingCellsCount++
			}
		}

	}

	if grid[h][w] == livingChar && (livingCellsCount == 2 || livingCellsCount == 3) {
		return true
	} else if grid[h][w] != livingChar && livingCellsCount == 3 {
		return true
	}

	return false
}
