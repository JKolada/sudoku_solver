package main

import (
  "fmt"
)

type Sudoku struct {
   inputTable [9][9] uint8
   markerTable [9][9][9] bool
   solution [9][9] uint8
   gotChanged bool
}

/* constructor, printing input table*/
func NewSudoku(inputTable [9][9]uint8) *Sudoku {
    s := new(Sudoku)
    s.inputTable = inputTable
    s.solution = s.inputTable
    print9x9(s.solution)
    return s
}

/* main method, executing sudoku solving*/
func (s *Sudoku) resolve() {
    s.initializeMarkerTable()
    s.correctMarkerTable()    
    
    a := 0
    for {
      s.gotChanged = s.solveBasedOnMarkers()
      print9x9(s.solution)
      print9x9x9(s.solution, s.markerTable) //todo delete
      s.gotChanged = s.solveByUniqueCandidate()
      a++
      if !s.gotChanged {break}
    }

    if !(s.checkIfFinishedAndCorrect()) {

      fmt.Printf("simple solution is not enough.. %d\n\n",a)

      print9x9x9(s.solution, s.markerTable) //todo delete

      s.solveBasingOnPotentialityImplications()
    } else {
      fmt.Println("CORRECTION FLAG: ", s.checkIfFinishedAndCorrect())
    }

    print9x9(s.solution)
}