package main 

//import "fmt"

func (s *Sudoku) solveByNakedSubsets() bool {
   var sthChanged bool

   var markerCount [9][9]uint8
   var cellMarkers [9]bool

   // in int belows I'll be storing number of found cells that have the same cellMarkers
   var similarRowCellCounter, similarColumnCellCounter  int
   var rowSolvedCellCounter,  columnSolvedCellCounter   [9]int

   for a := range s.solution {
      for b := range s.solution[a] {
         if s.solution[a][b] == 0 {
            for c := range s.markerTable[a][b] {
               if s.markerTable[a][b][c] {
                  markerCount[a][b]++
               }
            }
         } else {
            rowSolvedCellCounter[a]++
            columnSolvedCellCounter[b]++
         }
      }
   }

   for a := range s.solution {
      for b := range s.solution[a] {
         if s.solution[a][b] == 0 {
            if markerCount[a][b] == 2 {
               cellMarkers = s.markerTable[a][b]
               for c := range s.solution[a] {

                  if c != b && s.markerTable[a][c] == cellMarkers {
                     similarRowCellCounter++
                  }

                  if similarRowCellCounter == 2 - 1  {
                     similarRowCellCounter = 0
                     if rowSolvedCellCounter[a] != 2 {
                        /*
                        fmt.Printf("row:  cell (%d:%d) is the same like  (%d:%d)\n", a+1, b+1, a+1, c+1)
                        fmt.Println(cellMarkers)
                        fmt.Println(s.markerTable[a][c]) */
                        for col := range s.solution[a] {
                           if s.solution[a][col] == 0 && s.markerTable[a][col] != cellMarkers {
                              for numberMarker := range s.markerTable[a][col] {
                                 if cellMarkers[numberMarker] {
                                    s.markerTable[a][col][numberMarker] = false
                                    sthChanged = true
   
                                 }
                              }
                           }
                        }
                        //print9x9x9(s.solution, s.markerTable)
                     }
                  }

                  if c != a && s.markerTable[c][b] == cellMarkers {
                     similarColumnCellCounter++
                  }

                  if similarColumnCellCounter == 2 - 1  {
                     similarColumnCellCounter = 0
                     if columnSolvedCellCounter[b] != 2 {
                        /*
                        fmt.Printf("column: cell (%d:%d) is the same like (%d:%d)\n", a+1, b+1, c+1, b+1)
                        fmt.Println(cellMarkers)
                        fmt.Println(s.markerTable[c][b]) */
                        for row := range s.solution[a] {
                           if s.solution[row][b] == 0 && s.markerTable[row][b] != cellMarkers {
                              for numberMarker := range s.markerTable[row][b] {
                                 if cellMarkers[numberMarker] {
                                    s.markerTable[row][b][numberMarker] = false
                                    sthChanged = true

                                 }
                              }
                           }
                        }
                        //print9x9x9(s.solution, s.markerTable)
                     }
                  }
               }
            }
         }
      }
   }
  return sthChanged
}

func (s *Sudoku) solveByLockedSubsets() bool {
  var ret bool
  return ret
}