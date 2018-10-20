package main




import "fmt"

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


func (s *Sudoku) checkIfSudokuIsCorrect() bool {
  var row_counter [9]int
  var column_counter [9]int
  var block_counter [9]int

  for a := range s.solution {
    for b := range s.solution[a] {
      if s.solution[a][b] != 0 {
        row_counter[s.solution[a][b] - 1]++
      }
      if s.solution[b][a] != 0 {
        column_counter[s.solution[b][a] - 1]++
      }
    }

    for c := range row_counter {
      if column_counter[c] > 1 || row_counter[c] > 1 {
        if column_counter[c] > 1 {
          fmt.Printf("There is too many numbers %d in column no %d\n\n", c+1, a+1)
        } else {
          fmt.Printf("There is too many numbers %d in row no %d\n\n", c+1, a+1)
        }
        return false
      }
      column_counter[c] = 0
      row_counter[c] = 0
    }
  }

  var a_min, b_min, a_max, b_max, a, b uint8

  for a_min = 0; a_min < 9; a_min += 3 {
    a_max = a_min + 2
    for b_min = 0; b_min < 9; b_min += 3 {
      b_max = b_min + 2
      for a = a_min; a <= a_max; a++ {
        for b = b_min; b <= b_max; b++ {
          // interested in only not filled in cells
          if s.solution[a][b] != 0 {            
            block_counter[s.solution[a][b] - 1]++
          }
        }
      }

      for c := range block_counter {
        if block_counter[c] > 1 {
          fmt.Printf("There is too many numbers %d in block (%d:%d)\n\n", c+1, (a_min+1)/3+1, (b_min+1)/3+1)
          return false
        } 
        block_counter[c] = 0
      }
    }
  }
  return true
}



func (s *Sudoku) fillSolutionCell(a, b, solution uint8) {
  s.solution[a][b] = solution
  fillFalse9(&s.markerTable[a][b])
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

  //the easiest way to iterate through blocks and have stored their index boundaries
  for a_min = 0; a_min < 9; a_min += 3 {
    a_max = a_min + 2

    for b_min = 0; b_min < 9; b_min += 3 {
      b_max = b_min + 2
      
      fillZeroes9(&blockSolution)
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