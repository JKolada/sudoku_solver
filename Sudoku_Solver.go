package main

import "fmt"

type Sudoku struct {
   inputTable [9][9] int
   possibilityTable [9][9][9] bool

   solution [9][9] int

   gotChanged bool
}


//func (s Sudoku) printInput


// function printing sudoku table on a console
func print9x9(sudoku [][]int) {
  printSudokuLine()
  for a := range sudoku {
  	fmt.Print("| ")
    for b := range sudoku[a] {
      fmt.Print(sudoku[a][b], " ")
      if b%3 == 2 {fmt.Print("| ")}
    }
    fmt.Println()
    if a%3 == 2 {printSudokuLine()}
  }
}

func printSudokuLine() {
  fmt.Println("-------------------------")
}
