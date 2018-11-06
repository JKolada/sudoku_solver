package main

import (
  "./sudoku_solver"
  "os"
  "fmt"
  "unicode"
)

/* File with 'main' function, only executing test input data 

  available test inputs:

	sudokuInput_easy

	sudokuInput_medium

	sudokuInput_hard1
	sudokuInput_hard2
	sudokuInput_hard3
	sudokuInput_hard4
	sudokuInput_hard5

	sudokuInput_GOD1
	sudokuInput_GOD2
	sudokuInput_GOD3

	sudokuInput_HARDEST
*/


func main() {
	if len(os.Args) == 2 {
	  inputCorrect, parsedInput := getInput(os.Args[1])
	  if inputCorrect {
			a := sudoku_solver.NewSudoku(parsedInput)
			if a != nil {
			  a.Resolve()
			}
		}
	} else {
		doIntro()
	}
}

func getInput(arg string) (bool, [9][9]uint8) {
	var result [9][9]uint8
  var counter int

	for pos := range arg {
		if unicode.IsDigit(rune(arg[pos])) {
			result[counter/9][counter%9] = uint8(arg[pos] - '0')
			counter++
		} else if arg[pos] == ' ' {
			result[counter/9][counter%9] = uint8(0)
			counter++
		}
	}

	if counter != 81 {	
		return false, result
	} else {
		return true, result
	}
}

func doIntro() {
  fmt.Println("To receive a solution for a specific sudoku puzzle,")
	fmt.Println("write all numbers of the sudoku in one line as an argument in the application execution.")
	fmt.Println("Use blank space or zero number for not-filled cells.\n")
	fmt.Println("Lines below execute the application for the same puzzle.")
	fmt.Println("sudoku_solver.exe \"       3  1   26  5  7    9  29  1   7       4   3  5  9   48      6   23  5   7 \"")
  fmt.Println("sudoku_solver.exe 000000030010002600500700009002900100070000000400030050090004800000060002300500070\n")
	fmt.Println("Demonstration of solving a hardcoded sudoku that was given as input:\n")
	a := sudoku_solver.NewSudoku(sudokuInput_easy)
	if a != nil {
	  a.ResolveByDeduction()
	}
}