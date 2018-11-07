# Sudoku solver 1.0

## Description and remarks:
#### The application aims to solve quickly any sudoku puzzle.

   Sudoku Solver tries to solve puzzles by deduction algorithms, with no blind guessing. If the puzzle complexity is very high, the backtracking algorithm is used.

   `Backtracking algorithm is a *brute force* algorithm that iterates through all cells in sudoku, row by row, trying to match any number as a potential solution. If it sees that it is impossible to fill another cell, the cell solution is cleared and algorithm goes back to the previous cell and tries an another possible cell value.`

#### Remarks:
The main plan for the application is to solve any puzzle efficiently with no brute force. There still are a lot of complex deduction algorithms waiting to be implemented.

Benchmark results can change, especially for deduction algorithms, because there was no efficiency correction implemented.

## Usage instruction

#### How to use an app as a console app:

1. If you want to build your own version of the app, you are probably familiar with *golang* app building, but if not, then you need to install proper version of *go* from [here](https://golang.org/dl/), and then, build an application with `go build -o sudoku_solver.exe` command executed in the fetched repository folder. 

2. Execute the *.exe* file to see a demonstration of solving a sudoku puzzle.

3. To receive a solution for a specific sudoku puzzle, write all numbers of the sudoku, row by row, in one line as an argument for the application execution. Use blank space or zero number for not-filled cells. Lines below execute the application for the same puzzle. 




```
sudoku_solver.exe "       3  1   26  5  7    9  29  1   7       4   3  5  9   48      6   23  5   7 "
sudoku_solver.exe 000000030010002600500700009002900100070000000400030050090004800000060002300500070
```

##### The commands resolve the following sudoku puzzle:

Column no. | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9
------------- | - | - | - | - | - | - | - | - | -
**row no. 1** | _ | _ | _ | _ | _ | _ | _ | 3 | _
**row no. 2** | _ | 1 | _ | _ | _ | 2 | 6 | _ | _
**row no. 3** | 5 | _ | _ | 7 | _ | _ | _ | _ | 9
**row no. 4** | _ | _ | 2 | 9 | _ | _ | 1 | _ | _
**row no. 5** | _ | 7 | _ | _ | _ | _ | _ | _ | _
**row no. 6** | 4 | _ | _ | _ | 3 | _ | _ | 5 | _
**row no. 7** | _ | 9 | _ | _ | _ | 4 | 8 | _ | _
**row no. 8** | _ | _ | _ | _ | 6 | _ | _ | _ | 2
**row no. 9** | 3 | _ | _ | 5 | _ | _ | _ | 7 | _



#### How to use an app as a golang module:

```golang
import "sudoku_solver" // use with the right path of the sudoku_solver folder
```

```golang
...
  a := sudoku_solver.NewSudoku(parsedInput) 
  // -> initialize the main struct with [9][9]uint8 table where any blank space to fill in sudoku is 0
  if a != nil {
    a.Resolve()
    // -> the constructor method returns sudoku only if input sudoku is correct
    // then it is possible to execute the sudoku solver algorithms sequence
  }
...
```

### Algorithms implemented and related go functions:
#### Basic algorithms, `.\sudoku_solver\algorithms_basic.go` file content
* Unique candidate - *solveByUniqueCandidate()* method
* Hidden singles - *solveByHiddenSingles()* method

**The file also consists of functions:**
* *solveBasingOnMarkers()* method - looks for the cells that have only one marker checked
* *checkIfFinishedAndCorrect()* - counts the sum of every row and column and checks if it equals 45 - useful in quick check if filled sudoku solution is correct
* *checkIfSudokuIsCorrect()* - checks if there is no the same value in row/column/block for every cell -
it is useful in sudoku validity check
* *fillSolutionCell()* - simply fills the cell with result, executes *correctMarkersBasedOnCellSolution()*
* *initializeMarkerTable()* - it initializes the markers for input sudoku, deletes all markers for already filled cells 
* *correctMarkersBasedOnCellSolution()* - deletes markers for corresponding row/column/block for given solution cell
* *correctMarkerTable()* - executes *correctMarkersBasedOnCellSolution()* for every cell, part of the sudoku initialization before solving algorithm sequence execution

#### Average complexity algorithms
* Naked subsets & locked subsets - `.\sudoku_solver\algorithms_naked_and_locked_subsets.go`
* Pointing pairs - `.\sudoku_solver\algorithms_pointing_pairs.go`
* Hidden pairs - `.\sudoku_solver\algorithms_hidden_pairs.go`

#### High complexity algorithms
* Pointing block subsets - `.\sudoku_solver\algorithms_pointing_block_subsets.go`
 
## Future plans:

#### Plan for the updates:
* Finishing the X-Wing algorithm
* Implementation of the Swordfish algorithm
* Implementation of the X-Cycles algorithm
* Implementation of other deduction algorithms
* Implementation of sudoku solving history saving possibility

### Benchmark results

##### ThinkPad T540p `i7-4710MQ CPU@2,50GHz; 16,0 GB RAM`

The most updated results are in excel file in main repository folder.

Sudoku level | Puzzle no. | Backtracking by row [µs] | Backtracking by block [µs] | Deduction algorithms [µs]
------------ | ------------: | ----------------------- | ------------------------- | ------------------------
Very Easy  | 1     |     7,6        |    8,8 		  |    12,2
Medium     | 2     |     379,2      |    572,0 		|    19,1
Hard       | 3     |     160,6      |    210,9 		|    39,5
Hard       | 4     |     634,6      |    759,6 		|    114,2
Hard       | 5     |     275,6      |    294,2 		|    71,4
Hard       | 6     |     362,8      |    217,0 		|    123,2
Hard       | 7     |     6 321,5    |    8 805,0 	|    108,3
God        | 8     |     87 149,4   |    46 973,1 |    **X**
God        | 9     |     3 398,1    |    1 359,2 	|    **X**
God        | 10    |     2 186,7    |    1 624,1 	|    **X**
Hardest    | 11    |     8 355,3    |    19 648,7 |    **X**

--------------------------------------------

#### Averages on different puzzle levels

Sudoku level | Backtracking by row [µs] | Backtracking by block [µs] | Deduction algorithms [µs]
------------ | ----------------------- | ------------------------- | ------------------------
Very Easy     | 7,6 |    8,8  |    12,2
Medium        | 379,2 |    572,0  |    19,1
Hard          | 1551,0 |    2057,3  |    91,3
God           | 25272,4 |    17401,3  |    **X**
**All levels**| **9930,1**| **7315,7** |    **X**

-----------------------------------------

#### Benchmark execution
**To run benchmarks on your machine, just execute the line below:**
`\go_sudoku>go test -bench=.`
