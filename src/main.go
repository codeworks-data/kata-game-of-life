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

	board2 := game.GetNewBoard(10, 10)
	board2.SetAlive(2, 3)
	board2.SetAlive(3, 3)
	board2.SetAlive(4, 3)
	board2.SetAlive(4, 4)
	board2.SetAlive(5, 3)
	board2.SetAlive(5, 4)
	board2.SetAlive(6, 1)
	board2.SetAlive(6, 2)
	board2.SetAlive(6, 3)
	board2.SetAlive(6, 4)
	board2.SetAlive(6, 5)
	board2.SetAlive(6, 6)

	boards := []game.Board{
		board,
		board2,
	}

	channel0 := make(chan game.Board)
	channel1 := make(chan game.Board)

	for i := 0; i < 15; i++ {
		
		go boards[0].ComputeNextGen(channel0)		
		go boards[1].ComputeNextGen(channel1)

		nbBoardsComputed := 0
		for nbBoardsComputed < 2 {
			select {
			case board := <-channel0:
				boards[0] = board
				fmt.Println("Board 0:")
				fmt.Printf("%v", boards[0])
			case board := <-channel1:
				boards[1] = board
				fmt.Println("Board 1:")
				fmt.Printf("%v", boards[1])
			}

			fmt.Print("\n\n")
			nbBoardsComputed++
		}

		fmt.Println("------------------------------")
		time.Sleep(3 * time.Second)
	}

}