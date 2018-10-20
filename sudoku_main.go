package main
/* File containing Sudoku structure,
   its constructor method and 
   method 'resolve' responsible for executing sequences of solving algorithms
*/

import (
  "fmt"
)

type Sudoku struct {
   /* Table with input Sudoku puzzle, that will remain unchanged */
   inputTable     [9][9] uint8
   /* Most import temporary data that consists all cell possible solutions */
   markerTable [9][9][9] bool
   /* Copy of inputTable where will be cell solutions wrote down*/
   solution       [9][9] uint8

}

/* constructor, printing input table*/
func NewSudoku(inputTable [9][9]uint8) *Sudoku {
    s := new(Sudoku)
    s.inputTable = inputTable
    s.solution = s.inputTable
    print9x9(s.solution)
    if !s.checkIfSudokuIsCorrect() {
      fmt.Println("The input is incorrect\n\n")
      return nil
    }
    return s
}

/* main method, executing sudoku solving*/
func (s *Sudoku) resolve() {
    var gotChanged bool

    s.initializeMarkerTable()
    s.correctMarkerTable()    
    
    first_lvl_algorithms_counter := 0
    second_lvl_algortihms_counter := 0

asd :=0

  for {
    for {
      gotChanged = s.solveBasedOnMarkers()
      gotChanged = s.solveByUniqueCandidate() || gotChanged

      if !gotChanged {
        print9x9(s.solution)
        print9x9x9(s.solution, s.markerTable)
        asd++
        fmt.Printf("Podejście no%d, kołko łatwych algorytmów = %d\n\n", asd, first_lvl_algorithms_counter)
        gotChanged = s.solveByNakedSubsets()
        print9x9(s.solution)
        print9x9x9(s.solution, s.markerTable)
      } 

      first_lvl_algorithms_counter++
      if !gotChanged {break}
    }

    if !s.checkIfFinishedAndCorrect() {
      fmt.Println("NOW BIG ONE\n")
      print9x9x9(s.solution, s.markerTable) //todo delete
      second_lvl_algortihms_counter++
      gotChanged = s.solveBasingOnMarkersImplications()   
      if !gotChanged {
        if s.checkIfSudokuIsCorrect() {
          fmt.Printf("GAVE UP after %d simple loops and %d, 2nd level algorithms loops\n\n\n", first_lvl_algorithms_counter, second_lvl_algortihms_counter)
        } else {
          fmt.Println("There is a logical problem with sudoku solving. It could be poorly designed")
        }
        break
      }
    } else {
      fmt.Printf("Finished after %d simple loops and %d, 2nd level algorithms loops\n\n\n", first_lvl_algorithms_counter, second_lvl_algortihms_counter)
      break
    }
  }
  
  print9x9x9(s.solution, s.markerTable) //todo delete
  print9x9(s.solution)
}