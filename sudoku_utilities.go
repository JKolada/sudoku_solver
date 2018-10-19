package main

import (
  "fmt"
  "strconv"
)

/*  UTILITIES FUNCTIONS 
    Mostly function used for console debbuing: printing data
 */

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

func print9x9x9(sudoku [9][9]uint8, markerTable [9][9][9]bool) {
  print9x9x9Line()
  for a := range markerTable {
    fmt.Print("|")
    for b := range markerTable[a] {
      fmt.Print(" ")
      if sudoku[a][b] != 0 {
        fmt.Printf(">>> %d <<< ", sudoku[a][b])
      } else {
        for c := range markerTable[a][b] {
          if markerTable[a][b][c] {
            fmt.Printf("%d",c+1)
          } else {
            fmt.Printf(" ")
          }
        }
        fmt.Print(" ")
      }
      if b%3 == 2 {fmt.Print("|")}
    }
  fmt.Printf("\n")
  if a%3 == 2 {print9x9x9Line()}  
  }
}

func getSudokuNumberToPrint(a uint8) string {
  if a == 0 {
    return "_ "
  }  else {
    return strconv.Itoa(int(a)) + " "
  }
}

func printSudokuLine() {
  fmt.Println("  -------------------------")
}

func print9x9x9Line() {
  fmt.Print("|---------------------------------------------------")
  fmt.Println("--------------------------------------------------|")
}

func fillFalse(tofill *[9][9]bool) {
  for a := range tofill {
    for b := range tofill[a] {
      tofill[a][b] = false
    }
  }
}