package main

import (
	"fmt"
	"errors"
)

type Sudoku struct {
   inputTable [9][9] uint8

   potentialityTable [9][9][9] bool

   solution [9][9] uint8
   gotChanged bool
}

func NewSudoku(inputTable [9][9]uint8) *Sudoku {
    s := new(Sudoku)
    s.inputTable = inputTable
    s.solution = s.inputTable
    print9x9(s.solution)
    return s
}

func (s *Sudoku) resolve() {
    s.initializePotentialityTable()
    s.correctPotentialityTable()    
    //print9x9x9(s.solution, s.potentialityTable)
    
    a := 0
    for {
      s.gotChanged = s.solveBasedOnPotentialities()
      a++
      if !s.gotChanged {break}
    }

    if !(s.checkIfFinishedAndCorrect()) {


      //print9x9x9(s.solution, s.potentialityTable) //todo delete

      s.correctPotentialityImplications()
    } else {
      fmt.Println("NUMBER OF POTENTIALITY LOOPS ", a)
      fmt.Println("CORRECTION FLAG: ", s.checkIfFinishedAndCorrect())
    }

    print9x9(s.solution)
}

func (s *Sudoku) checkIfFinishedAndCorrect() bool {
	var overall_check uint16
	for a := range s.solution {
	  var row_check, column_check uint8
      for b := range s.solution[a] {
        row_check += s.solution[a][b]
        column_check += s.solution[b][a]
        overall_check += uint16(s.solution[a][b])
        //fmt.Printf("a = %d, b = %d, row_check = %d, column_check = %d\n", a, b, row_check, column_check)
      }
      if row_check != 45 || column_check != 45 {
       	return false
      }
	}
	if overall_check != 405 {return false}
	return true
}

func (s *Sudoku) solve(a, b, solution uint8) {
  s.solution[a][b] = solution
  s.correctPotentialityBasedOnCell(a, b)
}

func (s *Sudoku) initializePotentialityTable() {
  for a := range s.potentialityTable {
    for b := range s.potentialityTable[a] {
      var temp bool
      if s.solution[a][b] != 0 {
      	temp = false
      } else {
      	temp = true
      }
  	  for c := range s.potentialityTable[a][b] {
  	     s.potentialityTable[a][b][c] = temp
      }
    }  	
  }
}

func (s *Sudoku) correctPotentialityBasedOnCell(a, b uint8) {
  c_min := (a/3) * 3
  c_max := c_min + 3
  d_min := (b/3) * 3
  d_max := d_min + 3 

  if s.solution[a][b] == 0 {
  	errors.New("the solution for the cell is not filled in")
  }

  /* DEBUGGING
  fmt.Println("a = ", a)
  fmt.Println("b = ", b)
  fmt.Println("a.b = ", s.solution[a][b])
  fmt.Println("c_min = ", c_min)
  fmt.Println("c_max = ", c_max)
  fmt.Println("d_min = ", d_min)
  fmt.Println("d_max = ", d_max) */

  //block possibility correction
  for c := c_min; c < c_max; c++ {
    for d := d_min; d < d_max; d++ {
      s.potentialityTable[c][d][s.solution[a][b] - 1] = false
    }
  }

  //column possibility correction
  for c := 0; c < 9; c++ {        
    s.potentialityTable[c][b][s.solution[a][b] - 1] = false
  }

  //row possibility correction
  for d := 0; d < 9; d++ {        
    s.potentialityTable[a][d][s.solution[a][b] - 1] = false
  }
}

func (s *Sudoku) correctPotentialityTable() {
  for a := range s.solution {
    for b := range s.solution[a] {
      if s.solution[a][b] != 0 {  
      /* it takes every filled cell and exludes the number
         as a potential solution
         from corresponding row, column and block */    	
      	s.correctPotentialityBasedOnCell(uint8(a), uint8(b))
      }
    }
  }
}

func (s *Sudoku) solveBasedOnPotentialities() bool {  
  someSolutionFound := false
  for a := range s.potentialityTable {
    for b := range s.potentialityTable[a] {
      if s.solution[a][b] == 0 {
	    var solutionFound uint8 = 0
	    for c := range s.potentialityTable[a][b] {
	      if s.potentialityTable[a][b][c] == true {
	      	if solutionFound != 0 {
              solutionFound = 0
	      	  break
	      	} 
	      	solutionFound = uint8(c + 1)
	      }	
	    }
	    if solutionFound != 0 {
          fmt.Printf("new solution found, for a = %d, b = %d, and it is a %d\n", a+1, b+1, solutionFound)
          s.solve(uint8(a), uint8(b), solutionFound)
          someSolutionFound = true
	    } 
      }
    }
  }
  return someSolutionFound
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

func (s *Sudoku) correctPotentialityImplications() bool {
   var rows_potentiality [9][9]bool
   var columns_potentiality [9][9]bool
 
   // Loop for every block with keeping the coordinates 
   // It would be more efficient to do this in one big nested loop algorithm, like below, than executing it for every block separately
   for a_min := 0; a_min < 2; a_min += 3 {
       a_max := a_min + 2

       for b_min := 0; b_min < 2; b_min += 3 {
           b_max := b_min + 2

			for a := a_min; a <= a_max; a++ {
				for b := b_min; b <= b_max; b++ {
		           if s.solution[a][b] == 0 {
		           	  for c := range s.potentialityTable[a][b] {
		           	  	if s.potentialityTable[a][b][c] {
		                	rows_potentiality[a][c] = true
		                	columns_potentiality[b][c] = true
		                }
		           	  }
		           }
		           /*
		           if s.solution[b][a] == 0 {
		           	  for c := range s.potentialityTable[b][a] {
		           	  	if s.potentialityTable[b][a][c] {
		                	columns_potentiality[b][c] = true
		                }
		           	  }
		           }
		           */

				}
			}




       }

   }
      
/*

	for a := range s.solution {
		for b := range s.solution[a] {
           if s.solution[a][b] == 0 {
           	  for c := range s.potentialityTable[a][b] {
           	  	if s.potentialityTable[a][b][c] {
                	rows_potentiality[a][c] = true
                }
           	  }
           }
           if s.solution[b][a] == 0 {
           	  for c := range s.potentialityTable[b][a] {
           	  	if s.potentialityTable[b][a][c] {
                	columns_potentiality[b][c] = true
                }
           	  }
           } 	

		}
	}
*/
//rows_potentiality debug
	fmt.Println("rows_potentiality debug")
    for a := range rows_potentiality {
    	for c := range rows_potentiality[a] {
    		fmt.Print(rows_potentiality[a][c], " ")
    	}
    	fmt.Println()
    }
//column_potentiality debug
fmt.Printf("\ncolumn_potentiality debug\n")
    for b := range columns_potentiality {
    	for c := range columns_potentiality[b] {
    		fmt.Print(columns_potentiality[b][c], " ")
    	}
    	fmt.Println()
    }

    return false
}