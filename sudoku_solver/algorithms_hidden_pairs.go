package sudoku_solver 

// import "fmt"

func (s *Sudoku) solveByHiddenPairs() bool {
  // We are looking for two same markers that are placed in both of two cells
  // and nowhere else in one row or one column or one block:
  // then erasing all other markers from both cells 

  var markerChanged bool
  var ret bool

  var rowMarkerCounter, columnMarkerCounter [9]int
  var row2MarkerCounter, column2MarkerCounter int

  var block2MarkerCounter int
  var blockMarkerCounter [3][3][9]int

  // counter of solved cells in all rows/columns/blocks
  var rowUnfilledCounter, columnUnfilledCounter [9]int
  var blockUnfilledCounter [3][3]int

  // helper variables to hold temporary coordinates of first cell in row/column that meets requirements 
  var firstCellIndex int
  var firstCellMarker int

  // helper variables to hold temporary of first cell coordinates in block that meets requirements 
  var firstBlockCellColIndex, firstBlockCellRowIndex int


/*			Algorithm explained by example
col
no			Some row situation:
1.			  4 5   7   <- markers
2.			x           <- filled some solution
3.			3 4   6   9 -> actually 3 and 9 in this cell...
4.			x
5.			3 4     7 9 -> ...and these 3 & 9 are the hidden pair
6.			  4 5
7.			x
8.			  4   6 7
9.			  4   6 7
		  	___________
	  		2 6 2 3 4 2 <- sum of markers

			the rest of markers except 3 & 9 in columns 3. and 5. of the row are to delete

- finding first cell with filled 2 markers that exists only in 2 places in a row
   * found cell in column 3.
- finding another cell with filled same markers
   * found cell in column 5.
*/



  for a := range s.solution {
  	
  	for b := range s.solution[a] {
  		
  		if s.solution[a][b] == 0 {
        rowUnfilledCounter[a]++
        columnUnfilledCounter[b]++
        blockUnfilledCounter[a/3][b/3]++

        for c := range s.markerTable[a][b] {  
        	if s.markerTable[a][b][c] {
	          blockMarkerCounter[a/3][b/3][c]++
        	}
        }
  		}
  	}
  }

  /* THE PART RELATED TO ROW AND COLUMN */

  for a := range s.solution {
  	fillZeroes9(&rowMarkerCounter)
  	fillZeroes9(&columnMarkerCounter)

 	  for b := range s.solution[a] {
    	for c := range s.markerTable[a][b] {
    	  if s.markerTable[a][b][c] {
    			rowMarkerCounter[c]++
    		}
    		if s.markerTable[b][a][c] {
    			columnMarkerCounter[c]++
    		}    		
    	}
 	  }

		row2MarkerCounter = 0 
		column2MarkerCounter = 0
 	  for r := range rowMarkerCounter {
 	  	if rowMarkerCounter[r] == 2 {
 	  		row2MarkerCounter++
 	  	}
      if columnMarkerCounter[r] == 2 {
 	  		column2MarkerCounter++
 	  	}
 	  }

		if row2MarkerCounter > 1 &&
		   rowUnfilledCounter[a] > 2 {

			for b := range s.solution[a] {
				firstCellIndex = -1 
				firstCellMarker = -1

				for r := range rowMarkerCounter {

					if rowMarkerCounter[r] == 2 {

						if s.markerTable[a][b][r]   &&
						   firstCellMarker != -1    &&
						   b == firstCellIndex {

					 	  for b2 := range s.solution[a] {
					 	  	if s.markerTable[a][b2][firstCellMarker] &&
					 	  	   s.markerTable[a][b2][r] &&
					 	  	   b2 != b {
						      markerChanged = true
						      ret = true
						      /*
					 	  	  print9x9x9(s.solution, s.markerTable)
	    						fmt.Printf("a = %d, row markers = %v\n unfilled counters = %v\n", a+1, rowMarkerCounter, rowUnfilledCounter)
						      fmt.Printf("WOHOO a=%d, b1=%d, b2=%d, r1=%d, r2=%d\n", a+1, b+1, b2+1, r+1, firstCellMarker+1)
						      */
						      for r2 := range s.markerTable[a][b] {
						      	if r2 != r && r2 != firstCellMarker {
						      		s.markerTable[a][b][r2] = false
						      		s.markerTable[a][b2][r2] = false
						      	}
						      }
					 	  	  //print9x9x9(s.solution, s.markerTable)
					 	  	  break
						 	  }
					 	  }

						} else if rowMarkerCounter[r] == 2 &&
										 	firstCellMarker == -1    &&
										  s.markerTable[a][b][r] {
		          firstCellMarker = r
		          firstCellIndex  = b
						}

					}
			  }
			}
	  }
				      
		if column2MarkerCounter > 1 &&
		   columnUnfilledCounter[a] > 2 {

			for b := range s.solution[a] {
				firstCellMarker = -1
				firstCellIndex = -1 

				for r := range columnMarkerCounter {

					if columnMarkerCounter[r] == 2 {

						if s.markerTable[b][a][r] &&
						    firstCellMarker != -1 &&
						    b == firstCellIndex {

					 	  for b2 := range s.solution[a] {
					 	  	if s.markerTable[b2][a][firstCellMarker] &&
					 	  	   s.markerTable[b2][a][r] &&
					 	  	   b2 != b {
					 	  	  markerChanged = true
					 	  	  ret = true
					 	  	  /*
					 	  	  print9x9x9(s.solution, s.markerTable)
	    						fmt.Printf("a = %d, column markers = %v\n unfilled counters = %v\n", a+1, columnMarkerCounter, columnUnfilledCounter)
	    						fmt.Printf("WOHOO b=%d, a1=%d, a2=%d, r1=%d, r2=%d\n", a+1, b+1, b2+1, r+1, firstCellMarker+1)
						      */
						      for r2 := range s.markerTable[a][b] {
						      	if r2 != r && r2 != firstCellMarker {
						      		s.markerTable[b][a][r2] = false
						      		s.markerTable[b2][a][r2] = false
						      	}
						      }
					 	  	  //print9x9x9(s.solution, s.markerTable)
					 	  	  break
						 	  }
					 	  }
						} else if firstCellMarker == -1 &&
										  s.markerTable[b][a][r] {
		          firstCellMarker = r
		          firstCellIndex  = b
		          //fmt.Printf("first cell (a,b) = (%d,%d), r1 = %d\n", b+1,a+1,r+1)
						}

					}
			  }
			}
	  }
	  if markerChanged {
 			ret = s.solveBasingOnMarkers() || ret
 		  markerChanged = false
 		}
 	}


  /* THE PART RELATED TO BLOCK */
 	for a := range blockMarkerCounter {
 		for b := range blockMarkerCounter[a] {

 			for c := range blockMarkerCounter[a/3][b/3] {
 				if blockMarkerCounter[a/3][b/3][c] == 2 {
 					block2MarkerCounter++
 				} 				
 			}

 			if block2MarkerCounter > 1 &&
 				 blockUnfilledCounter[a][b] > 2 {
 				block2MarkerCounter = 0

 				col_min := a*3 
 				col_max := col_min + 2
 				row_min := b*3
 				row_max := row_min + 2


				firstCellMarker = -1
				firstBlockCellRowIndex = -1 
				firstBlockCellColIndex = -1

 				for col := col_min; col <= col_max; col++ {
				 	for row := row_min; row <= row_max; row++ {

  					for c := range blockMarkerCounter[a/3][b/3] {
  						if blockMarkerCounter[a/3][b/3][c] == 2 {
  							
  							if s.markerTable[row][col][c] &&
  								 firstCellMarker != -1 &&
  								 firstBlockCellColIndex == col &&
  								 firstBlockCellRowIndex == row {

					 				for col2 := col_min; col2 <= col_max; col2++ {
									 	for row2 := row_min; row2 <= row_max; row2++ {
							 
								 	  	if s.markerTable[row2][col2][firstCellMarker] &&
								 	  	   s.markerTable[row2][col2][c] &&
								 	  	   (firstBlockCellRowIndex != row2 || firstBlockCellColIndex != col2) {
								 	  	  markerChanged = true
								 	  	  ret = true
								 	  	  
								 	  	  //print9x9x9(s.solution, s.markerTable)
				    						//fmt.Printf("1:(a,b)=(%d,%d) 2:(a,b)=(%d,%d), r1 = %d, r2 = %d\n", row+1, col+1, row2+1,col2+1, firstCellMarker+1, c+1)
									      
									      for r := range s.markerTable[a][b] {
									      	if r != firstCellMarker &&
									      	   r != c {
									      		s.markerTable[row2][col2][r] = false
									      		s.markerTable[firstBlockCellRowIndex][firstBlockCellColIndex][r] = false
									      	}
									      }
								 	  	  //print9x9x9(s.solution, s.markerTable)
								 	  	  break
									 	  }
								 	  }
								 	} 								 
  							} else if firstCellMarker == -1 &&
  							          s.markerTable[row][col][c] {
  								firstCellMarker = c
	  							firstBlockCellColIndex = col
	  							firstBlockCellRowIndex = row
  							}
  						}
  					}
				 	}
 				}
 			}
 			if markerChanged {
 				ret = s.solveBasingOnMarkers() || ret
 				markerChanged = false
 			}
 		}
 	}

 	return ret
}