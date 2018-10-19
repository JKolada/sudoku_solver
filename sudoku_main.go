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
    
    first_lvl_algorithms_counter := 0
    second_lvl_algortihms_counter := 0

  for {
    for {
      s.gotChanged = s.solveBasedOnMarkers()
      //print9x9(s.solution)
      //print9x9x9(s.solution, s.markerTable) //todo delete
      s.gotChanged = s.solveByUniqueCandidate() || s.gotChanged
      //print9x9(s.solution)
      //print9x9x9(s.solution, s.markerTable) //todo delete
      first_lvl_algorithms_counter++
      if !s.gotChanged {break}
    }

    if !s.checkIfFinishedAndCorrect() {
      fmt.Println("NOW BIG ONE\n")
      //print9x9x9(s.solution, s.markerTable) //todo delete
      second_lvl_algortihms_counter++
      s.gotChanged = s.solveBasingOnMarkersImplications()      
      if !s.gotChanged {break}
    } else {
      fmt.Printf("Finished after %d simple loops and %d 2nd level algorithms loops\n\n\n", first_lvl_algorithms_counter, second_lvl_algortihms_counter)
      break
    }
  }
  
  print9x9(s.solution)
}