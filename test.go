package main

// TEST SUDOKU
var sudokuInput_easy = [9][9]uint8{
	{1,0,2, 0,7,3, 6,4,0},
	{5,0,8, 4,9,0, 2,0,1},
	{0,0,6, 0,8,2, 0,7,5},

	{3,0,7, 6,2,0, 1,0,4},
	{2,6,4, 0,1,8, 7,5,3},
	{0,0,0, 0,0,4, 0,2,0},

	{0,4,0, 0,5,7, 3,6,9},
	{7,0,3, 0,6,0, 0,0,0},
	{0,2,0, 0,4,1, 0,8,7},
}

var sudokuInput_medium = [9][9]uint8{
	{0,0,0, 0,2,0, 0,0,4},
	{0,7,0, 0,0,0, 0,0,0},
	{0,1,0, 5,0,0, 9,7,8},

	{0,0,8, 0,5,3, 0,0,0},
	{1,0,0, 2,0,8, 5,0,0},
	{0,0,4, 0,9,7, 0,0,0},

	{0,9,0, 8,0,0, 1,5,3},
	{0,3,0, 0,0,0, 0,0,0},
	{0,0,0, 0,6,0, 0,0,7},
}

var sudokuInput_hard = [9][9]uint8{
	{0,0,0, 0,0,0, 0,0,1},
	{0,0,4, 0,2,0, 0,3,0},
	{7,0,0, 0,0,9, 5,0,0},

	{2,0,0, 1,0,0, 0,9,0},
	{0,3,0, 0,0,0, 0,0,0},
	{0,0,8, 0,0,7, 0,0,6},

	{9,0,0, 0,0,0, 2,0,0},
	{0,6,0, 0,3,0, 0,8,0},
	{0,0,1, 0,0,5, 0,0,7},
}

var sudokuInput_GOD = [9][9]uint8{
	{0,0,0, 0,0,0, 0,0,1},
	{0,0,4, 0,2,0, 0,3,0},
	{7,0,0, 0,0,9, 5,0,0},

	{2,0,0, 1,0,0, 0,9,0},
	{0,3,0, 0,0,0, 0,0,0},
	{0,0,8, 0,0,7, 0,0,6},

	{9,0,0, 0,0,0, 2,0,0},
	{0,6,0, 0,3,0, 0,8,0},
	{0,0,1, 0,0,5, 0,0,7},
}

func main() {
    a := NewSudoku(sudokuInput_medium)
    a.resolve()    
}
