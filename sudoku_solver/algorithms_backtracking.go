package sudoku_solver 

//import "fmt"

func (s *Sudoku) solveByRowBacktracking() {
  var prev_a, prev_b int

  for a := 0; a < 9; a++ {
  	for b := 0; b < 9; b++ {
  		if s.inputTable[a][b] == 0 {
        if !s.tryNextSolution(a, b) {
          a2 := a
          b2 := b
          for ; a2 > -1; a2-- {
            for ; b2 > -1; b2-- {
              if (a2 != a || b2 != b) &&
                  s.inputTable[a2][b2] == 0 {
                prev_a = a2
                prev_b = b2
                //fmt.Printf("previous cell (a,b)=(%d,%d)\n",a2+1, b2+1)
                a2 = -1 
                b2 = -1
                break
              }           
            }
            b2 = 8
          }

          a_down := prev_a
          b_down := prev_b
          for ; a_down >= 0; a_down-- {
            for ; b_down >= 0; b_down-- {
              if s.inputTable[a_down][b_down] == 0 {                
                //fmt.Printf("Going down with (a,b)=(%d,%d)\n",a_down+1, b_down+1)
                if s.tryNextSolution(a_down, b_down) {
                  a = a_down
                  b = b_down
                  a_down = -1
                  b_down = -1
                  break
                }
              } 
            }
            b_down = 8
          }
    		}
      }
      //print9x9(s.solution)
  	}
  }

  s.isSolved = s.checkIfFinishedAndCorrect()
  if !s.isSolved {
    s.isCorrect = s.checkIfSudokuIsCorrect()
  }
}

func (s *Sudoku) solveByBlockBacktracking() {
  var prev_a, prev_b int

  var  gettingToBlockBefore bool

  for a_min := 0; a_min < 9; a_min += 3 {
    a_max := a_min + 3
    for b_min := 0; b_min < 9; b_min += 3 {
      b_max := b_min + 3
      for a := a_min; a < a_max; a++ {
        for b := b_min; b < b_max; b++ {

          if s.inputTable[a][b] == 0 {
            if !s.tryNextSolution(a, b) {

              ad_max := (a/3)*3 + 2
              bd_max := (b/3)*3 + 2

              a_down := a 
              b_down := b

              gettingToBlockBefore = false

              for ; ad_max > 0; ad_max -= 3 {
                ad_min := ad_max - 3
                for ; bd_max > 0; bd_max -= 3 {
                  bd_min := bd_max - 3

                  if gettingToBlockBefore {                    
                    a_down = ad_max
                    b_down = bd_max
                  }
                  for ; a_down > ad_min; a_down-- {
                    for ; b_down > bd_min; b_down-- {
                      if (a_down != a || b_down != b) &&
                          s.inputTable[a_down][b_down] == 0 {
                        prev_a = a_down
                        prev_b = b_down
                        //fmt.Printf("previous cell (a,b)=(%d,%d)\n",prev_a+1, prev_b+1)
                        a_down = -1 
                        b_down = -1
                        ad_max = 0 
                        bd_max = 0
                        break
                      }  
                    }
                    b_down = bd_max
                  }
                  gettingToBlockBefore = true
                }
                bd_max = 8
              }

              ad_max = (prev_a/3)*3 + 2
              bd_max = (prev_b/3)*3 + 2

              a_down = prev_a
              b_down = prev_b
               
              gettingToBlockBefore = false
              for ; ad_max > 0; ad_max -= 3 {
                ad_min := ad_max - 3
                for ; bd_max > 0; bd_max -= 3 {
                  bd_min := bd_max - 3
                  //fmt.Printf("gettingToBlockBefore\na min = %d, a max = %d, b min = %d, b max = %d\n",ad_min, ad_max, bd_min, bd_max )
                  if gettingToBlockBefore {                    
                    a_down = ad_max
                    b_down = bd_max
                  }                  
                  for ; a_down > ad_min; a_down-- {
                    for ; b_down > bd_min; b_down-- {

                      //fmt.Printf("Going down with (a,b)=(%d,%d)\n",a_down+1, b_down+1)
                      if s.inputTable[a_down][b_down] == 0 {
                        if s.tryNextSolution(a_down, b_down) {
                          //setting proper cell and block loop iterations
                          a = a_down
                          b = b_down
                          a_min = (a/3)*3
                          a_max = a_min + 3                          
                          b_min = (b/3)*3
                          b_max = b_min + 3
                          // below, the way of exiting last 4 for loops
                          a_down = -1 
                          b_down = -1
                          ad_max = -1
                          bd_max = -1
                          break
                        }
                      } 
                    }
                    b_down = bd_max
                  }
                  //fmt.Printf("getting to the block before na min = %d, a max = %d, b min = %d, b max = %d\n",ad_min, ad_max, bd_min, bd_max )                
                  gettingToBlockBefore = true
                }
                bd_max = 8
                //fmt.Printf("getting to the block higher\n")                
              }
            }
          }

        //print9x9(s.solution)
        }
      }
    }
  }

  s.isSolved = s.checkIfFinishedAndCorrect()
  if !s.isSolved {
    s.isCorrect = s.checkIfSudokuIsCorrect()
  }
}

func(s *Sudoku) tryNextSolution(a_idx, b_idx int) bool {
  //fmt.Printf("Trying to fill cell (a,b)=(%d,%d)\n", a_idx+1, b_idx+1)
  for marker := s.solution[a_idx][b_idx] + 1; marker < 10; marker++ {          
    if s.canIfillIt(marker, a_idx, b_idx) {
      s.solution[a_idx][b_idx] = marker
      return true
    }
  }
  s.solution[a_idx][b_idx] = 0
  //fmt.Printf("NOTHING FOR (%d,%d)\n", a_idx+1, b_idx+1)
  return false
}

func(s *Sudoku) canIfillIt(marker uint8, a_idx, b_idx int) bool{
  // checking row and column
	for idx := 0; idx < 9; idx++ {
		if idx != b_idx && s.solution[a_idx][idx] == marker {
      //fmt.Printf("%d marker NOT for %d col\n",marker,idx+1)
			return false
		}
		if idx != a_idx && s.solution[idx][b_idx] == marker {

      //fmt.Printf("%d marker NOT for %d row\n",marker,idx+1 )
			return false
		}
	}

  // checking block
  a_min := a_idx/3 * 3
  a_max := a_min + 2
  b_min := b_idx/3 * 3
  b_max := b_min + 2

  for a := a_min; a <= a_max; a++ {
  	for b := b_min; b <= b_max; b++ {
  		if (a != a_idx || b != b_idx) && s.solution[a][b] == marker {
        //fmt.Printf("%d marker NOT for (%d,%d) block\n",marker,a_idx+1,b_idx+1)
        return false
      }
  	}
  }
  //fmt.Printf("%d marker for (a,b)=(%d,%d)\n",marker,a_idx+1,b_idx+1)
  return true
}