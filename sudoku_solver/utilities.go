package sudoku_solver
/*  UTILITIES FUNCTIONS 
    Mostly function used for console debbuing: printing data
 */

import (
  "fmt"
  "strconv"
)

// function printing sudoku table on a console
func print9x9(sudoku [9][9]uint8) {
  fmt.Println(" b=> 1 2 3   4 5 6   7 8 9")
  fmt.Println(" a +-------+-------+-------+")
  for a := range sudoku {
    fmt.Printf(" %d | ", a+1)
    for b := range sudoku[a] {
      fmt.Print(getSudokuNumberToPrint(sudoku[a][b]))
      if b%3 == 2 {fmt.Print("| ")}
    }
    fmt.Println()
    if a%3 == 2 {fmt.Println("   +-------+-------+-------+")}
  }
}

func print9x9x9(sudoku [9][9]uint8, markerTable [9][9][9]bool) {
  fmt.Printf("    ")
  for a := 1; a < 10; a++ {
    fmt.Printf("  \\\\\\ %d /// ", a)
    if a%3 == 2 {fmt.Printf("|")}
  }
  fmt.Println()
  print9x9x9Line()
  for a := range markerTable {
    fmt.Printf(" %d |", a+1)
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

func print9x9x9Line() {
  fmt.Print("   |---------------------------------+-----------------")
  fmt.Println("----------------+---------------------------------|")
}

//////// FILLERS ////////

func fillFalse9x9(tofill *[9][9]bool) {
  for a := range tofill {
    for b := range tofill[a] {
      tofill[a][b] = false
    }
  }
}

func fillFalse9(tofill *[9]bool) {
  for a := range tofill {
    tofill[a] = false
  }
}

func fillZeroes9(tofill *[9]int) {
  for a := range tofill {
    tofill[a] = 0
  }
}

func fillZeroes9x9(tofill *[9][9]int) {
  for a := range tofill {
    for b := range tofill[a] {
      tofill[a][b] = 0
    }
  }
}

//////// END OF FILLERS ////////
