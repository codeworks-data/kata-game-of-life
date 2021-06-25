package game

import "fmt"

type Board [][]int

type Game struct {
	Board Board
}

type Cell struct {
	X int
	Y int
}

const (
	DEAD = iota
	ALIVE
)

func GetNewBoard(height int, width int) Board {
	board := make([][]int, width)
	for i := 0; i < width; i++ {
		board[i] = make([]int, height)
		for j := 0; j < height; j++ {
			board[i][j] = DEAD
		}
	}
	return board
}

func (board Board) ComputeNextGen() Board {
	nextGen := GetNewBoard(len(board), len(board[0]))

	for i := range board {
		for j := range board[i]{
			neighboursNb := board.getNeighboursNb(i, j)
			if board[i][j] == DEAD {
				if neighboursNb == 3 {
					nextGen[i][j] = ALIVE
				}
			} else {
				if neighboursNb == 2 || neighboursNb == 3 {
					nextGen[i][j] = ALIVE
				}
			}
		}
	}

	return nextGen
}

func (board Board) getNeighboursNb(x int, y int) int {
	cellsToCheck := []Cell{
		{x-1, y-1},
		{x-1, y},
		{x-1, y+1},
		{x, y-1},
		{x, y+1},
		{x+1, y+1},
		{x+1, y},
		{x+1, y-1},
	}

	result := 0
	for _, cell := range cellsToCheck {
		if cell.IsInsideBoard(board) && board[cell.X][cell.Y] == ALIVE {
			result += 1
		}
	}

	return result
}

func (cell *Cell) IsInsideBoard(board Board) bool {
	return cell.X >= 0 && cell.X < len(board) && cell.Y >= 0 && cell.Y < len(board[0])
}

func (board Board) String() string {
	result := ""
	for i := range board {
		result += fmt.Sprintf("%v\n", board[i])
	}
	return result + "\n\n\n"
}

func (board Board) SetAlive(x int, y int) {
	board[x][y] = ALIVE
}