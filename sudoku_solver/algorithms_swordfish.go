package sudoku_solver 

import "fmt"

func (s *Sudoku) solveByXWing() bool {
 var ret bool

 var rowMarkerCounter, columnMarkerCounter [9]int

 // going through the markers
 // counting them per row and column
  for r := 0; r < 9; r++ {
    fillZeroes9(&rowMarkerCounter)    
    fillZeroes9(&columnMarkerCounter)

 	  for a := range s.solution {
 		  for b := range s.solution[a] {
 		  	if s.markerTable[a][b][r] {
    			rowMarkerCounter[a]++
    			columnMarkerCounter[b]++
    		}
 		  }
 	  }

	  firstCol := -1
	  for a := range rowMarkerCounter {
	  	if rowMarkerCounter[a] == 2 {
	 	  	for b := range columnMarkerCounter {
	 	  	  if columnMarkerCounter[b] > 1 {
	 	  			if s.markerTable[a][b][r] {
	 	  				if firstCol != -1 {

	 	  					for a2 := range rowMarkerCounter {
	 	  						if rowMarkerCounter[a2] == 2 &&
	 	  						   a2 != a && firstCol != b &&
	 	  						   s.markerTable[a2][firstCol][r] &&
	 	  						   s.markerTable[a2][b][r] &&
	 	  						   s.markerTable[a][firstCol][r] &&
	 	  						   s.markerTable[a][b][r] {
	 	  						  fmt.Printf("FOUND IT for row r=%d, a1=%d, a2=%d, b1=%d, b2=%d\n",r+1,a+1,a2+1,firstCol+1, b+1)
	 	  						  print9x9x9(s.solution, s.markerTable)
	 	  						}
	 	  					}
	 	  				} else {
	 	  					firstCol = b
	 	  				}
	 	  			}
	 	  		}
	 	  	}
  	  }
  	}

	  firstRow := -1
	  for b := range columnMarkerCounter {
	  	if columnMarkerCounter[b] == 2 {
	 	  	for a := range rowMarkerCounter {
	 	  	  if rowMarkerCounter[a] > 1 {
	 	  			if s.markerTable[a][b][r] {
	 	  				if firstRow != -1 {
	 	  					for b2 := range columnMarkerCounter {
	 	  						if columnMarkerCounter[b2] == 2 &&
	 	  						   b2 != b && firstRow != a &&
	 	  						   s.markerTable[firstRow][b2][r] &&
	 	  						   s.markerTable[firstRow][b][r] &&
	 	  						   s.markerTable[a][b2][r] &&
	 	  						   s.markerTable[a][b][r] {
	 	  						  fmt.Printf("FOUND IT for column r=%d, a1=%d, a2=%d, b1=%d, b2=%d\n",r+1,firstRow+1,a+1,b+1, b2+1)
	 	  						  print9x9x9(s.solution, s.markerTable)
	 	  						}
	 	  					}
	 	  				} else {
	 	  					firstRow = a
	 	  				}
	 	  			}
	 	  		}
	 	  	}
  	  }
  	}
  }

  return ret
}


/*

func (s *Sudoku) solveBySwordfish() bool {
 var ret bool

 var rowMarkerCounter, columnMarkerCounter [9]int
 var columnsToCheck, rowsToCheck [3]int

 // going through the markers
 // counting them per row and column
  for r := 0; r < 9; r++ {
    fillZeroes9(&rowMarkerCounter)    
    fillZeroes9(&columnMarkerCounter)

 	  for a := range s.solution {
 		  for b := range s.solution[a] {
 		  	if s.markerTable[a][b][r] {
    			rowMarkerCounter[a]++
    			columnMarkerCounter[b]++
    		}
 		  }
 	  }

 	  //Swordfish
	  columnsToCheck = [3]int{-1, -1, -1}
	  rowsToCheck = [3]int{-1, -1, -1}
	  for a := range rowMarkerCounter {
	  	// filling the indexes of rows that would be checkd
	  	if rowMarkerCounter[a] == 2 || rowMarkerCounter[a] == 3 {
	  	  counter := 0
	  		for b := 0; b < 9; b++ {
	  			if s.markerTable[a][b][r] {
	  		  	columnsToCheck[counter] = b
	  				counter++
	  				if counter == rowMarkerCounter[a] {
	  					break;
	  				}
	  			}
	  		}	  		
  	  }

	  	for b := 0; b < 9; b++ {
	    	if rowMarkerCounter {
	  	  	columnsToCheck[counter] = b
	  	  	counter++
	  			if counter == rowMarkerCounter[a] {
	  				break;
	  			}
	  		}
	  	}	  






  	}

	  firstRow := -1
	  for b := range columnMarkerCounter {
	  	if columnMarkerCounter[b] == 2 {
	 	  	for a := range rowMarkerCounter {
	 	  	  if rowMarkerCounter[a] > 1 {
	 	  			if s.markerTable[a][b][r] {
	 	  				if firstRow != -1 {
	 	  					for b2 := range columnMarkerCounter {
	 	  						if columnMarkerCounter[b2] == 2 &&
	 	  						   b2 != b && firstRow != a &&
	 	  						   s.markerTable[firstRow][b2][r] &&
	 	  						   s.markerTable[firstRow][b][r] &&
	 	  						   s.markerTable[a][b2][r] &&
	 	  						   s.markerTable[a][b][r] {
	 	  						  fmt.Printf("FOUND IT for column r=%d, a1=%d, a2=%d, b1=%d, b2=%d\n",r+1,firstRow+1,a+1,b+1, b2+1)
	 	  						  print9x9x9(s.solution, s.markerTable)
	 	  						}
	 	  					}
	 	  				} else {
	 	  					firstRow = a
	 	  				}
	 	  			}
	 	  		}
	 	  	}
  	  }
  	}
  }

  return ret
}

*/