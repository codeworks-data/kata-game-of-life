package main

import (
	"fmt"
	"kata-life-game/src/game"
	"time"
)

func main() {
	board := game.GetNewBoard(10, 10)

	board.SetAlive(2, 3)
	board.SetAlive(4, 5)
	board.SetAlive(3, 3)
	board.SetAlive(1, 1)
	board.SetAlive(2, 1)
	board.SetAlive(1, 2)
	board.SetAlive(2, 4)
	board.SetAlive(4, 6)
	board.SetAlive(3, 4)
	board.SetAlive(1, 2)
	board.SetAlive(2, 2)
	board.SetAlive(9, 8)
	board.SetAlive(8, 7)
	board.SetAlive(6, 5)
	board.SetAlive(7, 7)
	board.SetAlive(9, 8)
	board.SetAlive(8, 9)
	board.SetAlive(9, 8)
	board.SetAlive(5, 2)
	board.SetAlive(5, 4)
	board.SetAlive(5, 6)
	board.SetAlive(5, 4)
	board.SetAlive(5, 2)

	fmt.Printf("%v", board)

	for i := 0; i < 20; i++ {
		board = board.ComputeNextGen()
		fmt.Printf("%v", board)
		time.Sleep(3 * time.Second)
	}

}