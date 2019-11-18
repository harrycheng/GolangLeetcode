package mysolution

import "fmt"

/*
Element at grid[i][j] becomes at grid[i][j + 1].
Element at grid[i][m - 1] becomes at grid[i + 1][0].
Element at grid[n - 1][m - 1] becomes at grid[0][0].
*/
func shiftGrid(grid [][]int, k int) [][]int {
	for i:= 0; i< k; i++{
		grid = shiftGridOne(grid)
	}
	return  grid
}

func shiftGridOne(grid [][]int) [][]int {
	current := 0
	saved := 0
	flag := false

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if flag {
				current = saved
			} else {
				current = grid[i][j]
			}

			m, n := i, j+1
			if i < len(grid)-1 && j == len(grid[i])-1 {
				m = i + 1
				n = 0
			}

			if i == len(grid)-1 && j == len(grid[i])-1 {
				m = 0
				n = 0
			}

			saved = grid[m][n]
			grid[m][n] = current
			flag = true
		}
	}
	return grid
}

func TestshiftGrid() {
	grid := [][]int{{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	print2D(shiftGrid(grid, 1))
}

func print2D(grid [][]int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Printf("%d ", grid[i][j])
		}
		fmt.Print("\n")
	}
}
