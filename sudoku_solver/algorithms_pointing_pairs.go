package sudoku_solver 

//import "fmt"

func (s *Sudoku) solveByPointingPairs() bool {
  var ret bool

  //Counters if there exists any same marker in the corresponding row/column
  //but with exclusion of block for which cell we are currently iterating (basing on a,b)
  var rowMarkerCounter, columnMarkerCounter [9]int

  //fmt.Printf("starting pointing pairs algorithm\n\n")
  //print9x9x9(s.solution, s.markerTable)

  for a := range s.solution {
  	a_min := (a/3)*3       // minimal index of a block row
  	a_max := (a/3)*3 + 2   // maximal index of a block row
  	
  	for b := range s.solution[a] {
      b_min := (b/3)*3     // minimal index of a block column
      b_max := (b/3)*3 + 2 // maximal index of a block column

  		if s.solution[a][b] == 0 {
  			for c := range s.markerTable[a][b] {
  				if s.markerTable[a][b][c] {

 						for b2 := range s.solution[a] {
  					  if b2 > b_max || b2 < b_min {
  					  //counting how many 'c' markers are in the row/column
  					  	if s.markerTable[a][b2][c] {  					    			
  					  	  rowMarkerCounter[c]++
  					  	}
  					  }
  					}

  					// iterating in row/column outside the block, where (a,b) cell is
  					for a2 := range s.solution[a] {
  						if a2 > a_max || a2 < a_min {
								if s.markerTable[a2][b][c] {  					    			
  					    columnMarkerCounter[c]++
  					    }
  					  }
  					}
  				}
  			}

  			// Updating all markers in the block basing on marker counters.
  			// If there is no marker found in corresponding to cell (a,b) row/column,
  			// the marker is deleted from the whole block where the cell is,
  			// but with exclusion of the row/column itself
		  	for marker := range rowMarkerCounter {
  				if s.markerTable[a][b][marker] {
			    	if rowMarkerCounter[marker] == 0 {
			    		//fmt.Printf("Found marker to change, number %d, based on row and cell (%d,%d)\n", marker+1, a+1, b+1)
			    		for a2 := a_min; a2 <= a_max; a2++ {
					  		for b2 := b_min; b2 <= b_max; b2++ {
					  			if a2 != a {
					  				s.markerTable[a2][b2][marker] = false
					  			}
					  		}
			    		}
  						//print9x9x9(s.solution, s.markerTable)
			  		}
			    	if columnMarkerCounter[marker] == 0 {  		  		
			    		//fmt.Printf("Found marker to change, number %d, based on column and cell (%d,%d)\n", marker+1, a+1, b+1)
			    		for a2 := a_min; a2 <= a_max; a2++ {
					  		for b2 := b_min; b2 <= b_max; b2++ {
					  			if b2 != b {
					  				s.markerTable[a2][b2][marker] = false
					  			}
					  		}
			    		}
  						//print9x9x9(s.solution, s.markerTable)
			  		}
			  	}
		  	}
  			fillZeroes9(&rowMarkerCounter)
  			fillZeroes9(&columnMarkerCounter)
  		}
  	}
  }

  return ret
}