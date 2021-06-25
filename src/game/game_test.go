package game

import "testing"

func TestBoard_SetAlive_should_set_cell_to_alive(t *testing.T) {
	// GIVEN
	board := GetNewBoard(5, 5)
	x := 2
	y := 2

	// WHEN
	board.SetAlive(x, y)

	// THEN
	if board[x][y] != ALIVE {
		t.Error("expected alive cell, got: ", board[x][y])
	}
}

func TestBoard_SetAlive_should_not_change_other_cells(t *testing.T) {
	// GIVEN
	board := GetNewBoard(5, 5)
	x := 2
	y := 2

	// WHEN
	board.SetAlive(x, y)

	// THEN
	if board[x+1][y+1] != DEAD {
		t.Error("expected DEAD cell, got: ", board[x][y])
	}
}

func TestCell_IsInsideBoard_should_return_true_if_inside(t *testing.T) {
	// GIVEN
	board := GetNewBoard(5, 5)
	cell := Cell{3, 3}

	// WHEN
	result := cell.IsInsideBoard(board)

	// THEN
	if !result {
		t.Error("expected true")
	}
}

func TestCell_IsInsideBoard_should_return_false_if_not_inside(t *testing.T) {
	// GIVEN
	board := GetNewBoard(5, 5)
	cell := Cell{3, 3}

	// WHEN
	result := cell.IsInsideBoard(board)

	// THEN
	if !result {
		t.Error("expected true")
	}
}