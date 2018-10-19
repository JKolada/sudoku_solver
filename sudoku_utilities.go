package main

import (
	"fmt"
	"strconv"
)

/*  UTILITIES FUNCTIONS  */

// function printing sudoku table on a console
func print9x9(sudoku [9][9]uint8) {
  fmt.Println("    1 2 3   4 5 6   7 8 9")	
  printSudokuLine()
  for a := range sudoku {
  	fmt.Printf("%d | ", a+1)
    for b := range sudoku[a] {
      fmt.Print(getSudokuNumberToPrint(sudoku[a][b]))
      if b%3 == 2 {fmt.Print("| ")}
    }
    fmt.Println()
    if a%3 == 2 {printSudokuLine()}
  }
}

// function prints all not excluded cell values
func print9x9x9(sudoku [9][9]uint8, potentialityTable [9][9][9]bool) {
   for a := range potentialityTable {
   	 for b := range potentialityTable[a] {
   	 	if sudoku[a][b] == 0 {
	   	 	fmt.Printf("a = %d , b = %d, potentialities = ", a+1, b+1)
	   	 	for c := range potentialityTable[a][b] {
	           if potentialityTable[a][b][c] {
	           	  fmt.Print((c+1), ", ")
	           }
	   	 	}
	   	 	fmt.Println()
   	    }
   	 }
   }
}

func getSudokuNumberToPrint(a uint8) string {
	if a == 0 {
		return "_ "
	}	else {
		return strconv.Itoa(int(a)) + " "
	}
}

func printSudokuLine() {
  fmt.Println("  -------------------------")
}