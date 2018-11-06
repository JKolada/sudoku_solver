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
    fmt.Println("\nNeed to use backtracking algorithm\n")
    s.SolveByRowBacktracking()
  }
  //++++++++++++++++++++++//
  s.sumUp()
}

func(s *Sudoku) ResolveByDeduction() {
  fmt.Printf(" Sudoku received to solve:\n")
  print9x9(s.inputTable)
  //++++++++++++++++++++++//
  s.ResolveWithoutPrinting()
  //++++++++++++++++++++++//
  s.sumUp()
}

func (s *Sudoku) ResolveByBrute_Row() {
  fmt.Printf(" Sudoku received to solve:\n")
  print9x9(s.inputTable)
  //++++++++++++++++++++++//
  s.SolveByRowBacktracking()  
  //++++++++++++++++++++++//
  s.sumUp()
}

func (s *Sudoku) ResolveByBrute_Block() {
  fmt.Printf(" Sudoku received to solve:\n")
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

  for {
    for {
      for {
        /* BASIC ALGORITHMS (least complex)
          solveBasingOnMarkers - solution found if there is only one number marked for the cell
          solveByUniqueCandidate - solution found if the cell is only one with specific number marked for the block
          solveByHiddenSingles - solution found if the cell is only one with specific number marked for the row and the column
         */
        gotChanged = s.solveBasingOnMarkers()
        gotChanged = s.solveByUniqueCandidate() || gotChanged
        gotChanged = s.solveByHiddenSingles() || gotChanged

        if !gotChanged {break}
      }

      s.isSolved = s.checkIfFinishedAndCorrect()
      if s.isSolved {break}

      /* AVERAGE ALGORITHMS (it can take some time)
        solveByNakedAndLockedSubsets
        solveByPointingPairs
        solveByHiddenPairs

        In this case, if some of algorithms listed below has changed any marker,
        we are coming back to the basic algorithms loop, immediately, without executing the rest of them
      */
      gotChanged = s.solveByNakedAndLockedSubsets(2)
      if !gotChanged {gotChanged = s.solveByPointingPairs()}
      if !gotChanged {gotChanged = s.solveByHiddenPairs()}

      if !gotChanged {break}
    }    

    //print9x9(s.solution)
    //print9x9x9(s.solution, s.markerTable) //todo delete

    s.isSolved = s.checkIfFinishedAndCorrect()
    if s.isSolved {break}

    /* COMPLEX ALGORITHMS
      solveByPointingBlockSubsets
      solveByXWing
      solveBySwordfish
      solveByNakedAndLockedSubsets
    */
    gotChanged = s.solveByPointingBlockSubsets()
    //s.solveByXWing()
    //s.solveBySwordfish()
    s.checkIfSudokuIsCorrect()
    //gotChanged = s.solveByNakedAndLockedSubsets(3) || gotChanged

    if !gotChanged {
      s.isCorrect = s.checkIfSudokuIsCorrect()
      s.isSolved = s.checkIfFinishedAndCorrect()
      break        
    }
  }
}

func(s *Sudoku) sumUp() {

  if s.isSolved {
    fmt.Printf("\n Sudoku solved:\n")    
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