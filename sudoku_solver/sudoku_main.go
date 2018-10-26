package sudoku_solver
/* File containing Sudoku structure,
   its constructor method and 
   methods responsible for executing sequences of solving algorithms
*/

import "fmt"

type Sudoku struct {
   /* Table with input Sudoku puzzle, that will remain unchanged */
   inputTable     [9][9] uint8
   /* Most import temporary data that consists all cell possible solutions */
   markerTable [9][9][9] bool
   /* Copy of inputTable where will be cell solutions wrote down*/
   solution       [9][9] uint8

   isCorrect bool
   isSolved bool
}

/* constructor */
func NewSudoku(inputTable [9][9]uint8) *Sudoku {
    s := new(Sudoku)
    s.inputTable = inputTable
    s.solution = s.inputTable

    s.initializeMarkerTable()
    s.correctMarkerTable()

    if !s.checkIfSudokuIsCorrect() {
      fmt.Println("The input is incorrect\n\n")
      return nil
    }

    return s
}

func(s *Sudoku) Resolve() {
  fmt.Printf("Sudoku received to solve:\n")
  print9x9(s.inputTable)
  //++++++++++++++++++++++//
  s.ResolveWithoutPrinting()
  if !s.isSolved {
    fmt.Println("\nNeed to use backtracking algorithm...")
    s.SolveByRowBacktracking()
  }
  //++++++++++++++++++++++//
  s.sumUp()
}

func(s *Sudoku) ResolveByDeduction() {
  fmt.Printf("Sudoku received to solve:\n")
  print9x9(s.inputTable)
  //++++++++++++++++++++++//
  s.ResolveWithoutPrinting()
  //++++++++++++++++++++++//
  s.sumUp()
}

func (s *Sudoku) ResolveByBrute_Row() {
  fmt.Printf("Sudoku received to solve:\n")
  print9x9(s.inputTable)
  //++++++++++++++++++++++//
  s.SolveByRowBacktracking()  
  //++++++++++++++++++++++//
  s.sumUp()
}

func (s *Sudoku) ResolveByBrute_Block() {
  fmt.Printf("Sudoku received to solve:\n")
  print9x9(s.inputTable)
  //++++++++++++++++++++++//
  s.SolveByBlockBacktrackingVER2()
  //s.SolveByBlockBacktrackingVER2()
  //++++++++++++++++++++++//
  s.sumUp()
}


/* main method, executing sudoku solving*/
func (s *Sudoku) ResolveWithoutPrinting() {
  var gotChanged bool
  s.solveByXWing()

  for {
    for {
      // BASIC ALGORITHMS
      gotChanged = s.solveBasingOnMarkers()
      gotChanged = s.solveByUniqueCandidate()
      gotChanged = s.solveByNakedAndLockedSubsets(2) || gotChanged
      gotChanged = s.solveByHiddenSingles() || gotChanged
      s.solveByPointingPairs()
      gotChanged = s.solveByHiddenPairs() || gotChanged
      if !gotChanged {break}
    }    

    //print9x9(s.solution)
    //print9x9x9(s.solution, s.markerTable) //todo delete

    s.isSolved = s.checkIfFinishedAndCorrect()
    if !s.isSolved {


      gotChanged = s.solveByPointingBlockSubsets()
      s.solveByXWing()
      //s.solveBySwordfish()
      s.checkIfSudokuIsCorrect()
      //gotChanged = s.solveByNakedAndLockedSubsets(3) || gotChanged

      if !gotChanged {
        s.isCorrect = s.checkIfSudokuIsCorrect()
        break        
      }
    } else {
      break
    }
  }
}

func(s *Sudoku) sumUp() {

  if s.isSolved {
    fmt.Printf("\nSudoku solved:\n")    
    print9x9(s.solution)
  } else if s.isCorrect {
    fmt.Printf("GAVE UP. AIN'T NOBODY CAN SOLVE THIS!\n")    
    print9x9(s.solution)
    print9x9x9(s.solution, s.markerTable)
  } else {
    fmt.Println("There is a logical problem with sudoku solving. It could be poorly designed") 
    print9x9(s.solution)   
    print9x9x9(s.solution, s.markerTable)
  }

}