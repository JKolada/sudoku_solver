package sudoku_solver
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
func (s *Sudoku) Resolve() {
    var gotChanged bool

    s.initializeMarkerTable()
    s.correctMarkerTable()    
    
    first_lvl_algorithms_counter := 0
    second_lvl_algortihms_counter := 0

  for {
    for {
      gotChanged = s.solveBasingOnMarkers()
      gotChanged = s.solveByUniqueCandidate() || gotChanged
      gotChanged = s.solveByNakedAndLockedSubsets(2) || gotChanged
      gotChanged = s.solveByHiddenSingles() || gotChanged
      gotChanged = s.solveByPointingPairs() || gotChanged

      first_lvl_algorithms_counter++
      if !gotChanged {break}
    }
    
    print9x9(s.solution)
    print9x9x9(s.solution, s.markerTable) //todo delete

    if !s.checkIfFinishedAndCorrect() {
      fmt.Printf(">>>>>>>>>>>> Started using 2nd level algorithms <<<<<<<<<<<<\n\n")
      print9x9x9(s.solution, s.markerTable) //todo delete
      second_lvl_algortihms_counter++
      gotChanged = s.solveByPointingBlockSubsets()   
      //gotChanged = s.solveByNakedAndLockedSubsets(3) || gotChanged

      if !gotChanged {
        if s.checkIfSudokuIsCorrect() {
          fmt.Printf("GAVE UP after %d simple loops and %d, 2nd level algorithms loops\n\n\n", first_lvl_algorithms_counter, second_lvl_algortihms_counter)
        } else {
          fmt.Println("There is a logical problem with sudoku solving. It could be poorly designed")
        }
        break
      }
    } else {
      fmt.Printf("SUDOKU COMPLETED\nIt needed >%d< basic solving algorithm loops\n" ,first_lvl_algorithms_counter)
      /*
      if second_lvl_algortihms_counter != 0 {
        fmt.Printf("...and >%d< more complex algorithm loops\n\n", second_lvl_algortihms_counter)
      } else {
        if first_lvl_algorithms_counter < 4 {
          fmt.Printf("... it was.. VERY EASY\n")
        } else if first_lvl_algorithms_counter < 10 {
          fmt.Printf("The level of puzzles was: MEDIUM\n")
        }
      } */
      break
    }
  }

  print9x9(s.solution)
}