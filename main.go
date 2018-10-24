package main

import "./sudoku_solver"

/* File with 'main' function, only executing test input data 

  available test inputs:

	sudokuInput_easy

	sudokuInput_medium

	sudokuInput_hard1
	sudokuInput_hard2 !!!
	sudokuInput_hard3
	sudokuInput_hard4
	sudokuInput_hard5

	sudokuInput_GOD1
	sudokuInput_GOD2
	sudokuInput_GOD3

	sudokuInput_GOD4
	sudokuInput_GOD5

already spent (~30h)
TO DO:
- GOD sudoku solving algorithms (swordfish) (~10h)
- efficiency tests and correction of algorithms (~10h)
- bakctracking algorithm / intelligent backtracking algorithm
- saving history of sudoku solving  (~5h)
- solving sudoku given on input of exe file (~3h)
*/


func main() {

	/*
	a := sudoku_solver.NewSudoku(sudokuInput_easy)
	if a != nil {
	  //a.ResolveByDeduction()  
	  a.ResolveByBrute_Block()
	} */
	

  b := sudoku_solver.NewSudoku(sudokuInput_medium) //sudokuInput_GOD2)
	if b != nil {
	  //a.ResolveByDeduction()  
	  b.ResolveByDeduction()
	}

	b = sudoku_solver.NewSudoku(sudokuInput_GOD3) //sudokuInput_GOD2)
	if b != nil {
	  b.ResolveByBrute_Block()
	}
}

