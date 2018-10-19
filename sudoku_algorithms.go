package main

import (
  "fmt"
  "errors"
)

const ROW_COLUMN_SUM = 45
const ALL_SUDOKU_CELL_SUM = 405

func (s *Sudoku) checkIfFinishedAndCorrect() bool {
  var overall_check uint16
  for a := range s.solution {
    var row_check, column_check uint8
      for b := range s.solution[a] {
        row_check += s.solution[a][b]
        column_check += s.solution[b][a]
        overall_check += uint16(s.solution[a][b])
      }
      if row_check != ROW_COLUMN_SUM || column_check != ROW_COLUMN_SUM {
         return false
      }
  }
  if overall_check != ALL_SUDOKU_CELL_SUM {return false}
  return true
}

func (s *Sudoku) fillSolutionCell(a, b, solution uint8) {
  s.solution[a][b] = solution
  s.correctMarkersBasedOnCellSolution(a, b)
}

func (s *Sudoku) initializeMarkerTable() {
  for a := range s.markerTable {
    for b := range s.markerTable[a] {
      var temp bool
      if s.solution[a][b] != 0 {
        temp = false
      } else {
        temp = true
      }
      for c := range s.markerTable[a][b] {
         s.markerTable[a][b][c] = temp
      }
    }    
  }
}

func (s *Sudoku) correctMarkersBasedOnCellSolution(a, b uint8) {
  c_min := (a/3) * 3
  c_max := c_min + 3
  d_min := (b/3) * 3
  d_max := d_min + 3 

  if s.solution[a][b] == 0 {
    errors.New("the solution for the cell is not filled in")
  }

  //block markers correction
  for c := c_min; c < c_max; c++ {
    for d := d_min; d < d_max; d++ {
      s.markerTable[c][d][s.solution[a][b] - 1] = false
    }
  }

  //column markers correction
  for c := 0; c < 9; c++ {        
    s.markerTable[c][b][s.solution[a][b] - 1] = false
  }

  //row markers correction
  for d := 0; d < 9; d++ {        
    s.markerTable[a][d][s.solution[a][b] - 1] = false
  }
}

func (s *Sudoku) correctMarkerTable() {
  for a := range s.solution {
    for b := range s.solution[a] {
      if s.solution[a][b] != 0 {  
      /* it takes every filled cell and exludes the number
         as a potential solution
         from corresponding row, column and block 
         */      
        s.correctMarkersBasedOnCellSolution(uint8(a), uint8(b))
      }
    }
  }
}

/* The next step in usage of
   sudoku solving strategy ->
   Sole Candidate 
*/
func (s *Sudoku) solveBasedOnMarkers() bool {  
  someSolutionFound := false
  for a := range s.markerTable {
    for b := range s.markerTable[a] {
      if s.solution[a][b] == 0 {
        var solutionFound uint8 = 0
        for c := range s.markerTable[a][b] {
          if s.markerTable[a][b][c] == true {
            if solutionFound != 0 {
              solutionFound = 0
              break
            } 
            solutionFound = uint8(c + 1)
          }  
        }
        if solutionFound != 0 {
          //fmt.Printf("based on markers: new solution found, for a = %d, b = %d, and it is %d\n", a+1, b+1, solutionFound)
          s.fillSolutionCell(uint8(a), uint8(b), solutionFound)
          someSolutionFound = true
        } 
      }
    }
  }
  return someSolutionFound
}

func (s *Sudoku) solveByUniqueCandidate() bool {
  var blockSolution [9]int
  var ret bool

  var a_min, b_min, a_max, b_max, a, b uint8

  for a_min = 0; a_min < 9; a_min += 3 {
    a_max = a_min + 2

    for b_min = 0; b_min < 9; b_min += 3 {
      b_max = b_min + 2
      
      for l := range blockSolution {
        blockSolution[l] = 0
      }


      for a = a_min; a <= a_max; a++ {
        for b = b_min; b <= b_max; b++ {
          // interested in only not filled in cells
          if s.solution[a][b] == 0 {
            // ..that still have some potential solutions:
             for c := range s.markerTable[a][b] {
              if s.markerTable[a][b][c] {
                blockSolution[c]++
                //fmt.Printf("a = %d, b = %d, c = %d, truth = %d", a+1, b+1, c+1, blockSolution[c])
              }
            }
          }
        }
      }  

      for sol_idx := range blockSolution {
        if blockSolution[sol_idx] == 1 {
          for a = a_min; a <= a_max; a++ {
            for b = b_min; b <= b_max; b++ {
              // interested in only not filled in cells
              if s.solution[a][b] == 0 {
                // ..that still have some potential solutions:
                if s.markerTable[a][b][sol_idx] {
                  ret = true
                  //fmt.Printf("unique candidate: new solution found, for a = %d, b = %d and it is %d\n", a+1, b+1, sol_idx+1)
                  s.fillSolutionCell(a, b, uint8(sol_idx + 1))                    
                }
              }
            }
          }
        }
      }
    }
  }
  return ret
}



/* Probably the most heavy algorithm.
   It looks for potential number
   that can be wrote down in one block only in one column or row.
   
   X - already filled in
   1, 2 - potential solutions, not yet excluded

  exemplar block
   |X 2 X|
   |2 X X|
   |1 1 1|

   In this case we can't still be sure in which cell
   there will be number 1, but we know that in whole row there
   can't be this 1 elsewhere, so we are exluding in from whole row potentiality table.
*/


// FUUUUUUCK, HAVE TO RENAME IT, STILL DID NOT FIND HOW THE ALGORITHM IS CALLED
   // MAYBE I'VE INVENTED IT 3:-)
func (s *Sudoku) solveBasingOnMarkersImplications() bool {
  somethingChanged := false
  var row_potentiality [9][9]bool
  var column_potentiality [9][9]bool

  rowTruthCounter := [9]int{0,0,0, 0,0,0, 0,0,0}
  columnTruthCounter := [9]int{0,0,0, 0,0,0, 0,0,0}
 
  // Loop for every block with keeping the coordinates 
  // It would be more efficient to do this in one big nested loop algorithm, like below, than executing it for every block separately
  for a_min := 0; a_min < 9; a_min += 3 {
    a_max := a_min + 2

    for b_min := 0; b_min < 9; b_min += 3 {
      b_max := b_min + 2

      fillFalse(&row_potentiality)
      fillFalse(&column_potentiality)

      for a := a_min; a <= a_max; a++ {
        for b := b_min; b <= b_max; b++ {
          // interested in only not filled in cells
          if s.solution[a][b] == 0 {
            // ..that still have some potential solutions:
             for c := range s.markerTable[a][b] {
              if s.markerTable[a][b][c] {
                 row_potentiality[a][c] = true
                 column_potentiality[b][c] = true
              }
             }
          }
        }
      }  

      for a := range rowTruthCounter {
        rowTruthCounter[a] = 0
      }

      for a := a_min; a <= a_max; a++ {
        for c := range row_potentiality[a] {
          if row_potentiality[a][c] {
            rowTruthCounter[c]++
          }
        }
      }

     /*
      for a := range rowTruthCounter {
         fmt.Printf("block (a:b) = (%d:%d), a = %d, counter  = %d\n", (a_min+1)/3, (b_min+1)/3, a+1, rowTruthCounter[a])
      } */

      for a := range rowTruthCounter {
        // if found only one row with number potentiality
        if rowTruthCounter[a] == 1 {
          // looking for that row
          for b := range row_potentiality {
            if row_potentiality[b][a] {
              
              // and doing row possibility correction
              for d := 0; d < 9; d++ {
                if (d < b_min || d > b_max) && s.markerTable[b][d][a] {
                  fmt.Printf("FOUND ROW THAT INFLUENCED CORRECTIONS = %d, number = %d\n", b+1, a+1)

                  fmt.Printf("b_min = %d, b_max = %d, a = %d, b = %d\n", b_min+1, b_max+1, b+1, d+1)

                  print9x9x9(s.solution, s.markerTable)

                  s.markerTable[b][d][a] = false
                  somethingChanged = true
                }
              }
            }
          }
        }
      }

      for a := range columnTruthCounter {
        columnTruthCounter[a] = 0
      }

      for b := b_min; b <= b_max; b++ {
        for c := range column_potentiality[b] {
          if column_potentiality[b][c] {
            columnTruthCounter[c]++
          }
        }
      }

      for a := range columnTruthCounter {
        // if found only one column with number potentiality
        if columnTruthCounter[a] == 1 {
          // looking for that column
          for b := range column_potentiality {
            if column_potentiality[b][a] {
              // and doing column possibility correction
              for d := 0; d < 9; d++ {        
                if (d < a_min || d > a_max) && s.markerTable[d][b][a] {
                  fmt.Printf("FOUND COLUMN THAT INFLUENCED CORRECTIONS = %d, number = %d\n", b+1, a+1)
                  fmt.Printf("a_min = %d, a_max = %d, a = %d, b = %d", a_min+1, a_max+1, d+1, b+1)
                  

                  print9x9x9(s.solution, s.markerTable)
                  s.markerTable[d][b][a] = false
                  somethingChanged = true
                }
              }
            }
          }
        }
      }
    }
  }
  return somethingChanged
}
     
//rows truth counter


