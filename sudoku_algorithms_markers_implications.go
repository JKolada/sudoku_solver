package main 

import "fmt"

/* The most heavy algorithm, as far
   It looks for potential number
   that can be wrote down in one block only in one column or row.
   
   X - already solved cells
   1, 2 - potential solutions, not yet excluded (markers)
   We assume that there are others markers in positions of number 1 and 2

  exemplar block
   |X 2 X|     |     |
   |2 X X|     |     |
   |1 1 1|Y Y Y|Y Y Y|
   |

   In this case we can't still be sure in which cell
   there will be number 1 in The block,
   but we know that in whole row there can't be this 1 elsewhere outside block,
   so we are exluding it from row in markers table, from Y positions.
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


                  s.markerTable[b][d][a] = false
                  print9x9x9(s.solution, s.markerTable)

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
                                    
                  s.markerTable[d][b][a] = false
                  print9x9x9(s.solution, s.markerTable)
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