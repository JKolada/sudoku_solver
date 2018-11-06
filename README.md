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

Sudoku level | Puzzle number | Backtracking by row [s] | Backtracking by block [s] | Deduction algorithms [s]
------------ | ------------- | ----------------------- | ------------------------- | ------------------------
Very Easy  | 1     |     0,0000046    |    0,0000056    |    0,0000199
Medium     | 2     |     0,0003938    |    0,0006071    |    0,0000323
Hard       | 3     |     0,0001613    |    0,0002212    |    0,0000745
Hard       | 4     |     0,0006751    |    0,0008075    |    0,0001196
Hard       | 5     |     0,0002767    |    0,0003110    |    0,0000934
Hard       | 6     |     0,0003598    |    0,0002155    |    0,0001183
Hard       | 7     |     0,0059266    |    0,0092348    |    0,0000777
God        | 8     |     0,0878996    |    0,0480385    |    **X**
God        | 9     |     0,0033681    |    0,0012883    |    **X**
God        | 10    |     0,0021168    |    0,0016321    |    **X**
Hardest    | 11    |     0,0080305    |    0,0192789    |    **X**

--------------------------------------------

#### Averages on different puzzle levels

Sudoku level | Backtracking by row [s] | Backtracking by block [s] | Deduction algorithms [s]
------------ | ----------------------- | ------------------------- | ------------------------
Very Easy  | 0,0000046    |    0,0000056  |    0,0000199
Medium     | 0,0003938    |    0,0006071  |    0,0000323
Hard       | 0,0014799    |    0,0021580  |    0,0000967
God        | 0,0253537    |    0,0175594  |    **X**
**All levels**| **0,0099284**| **0,0074219** |    **X**

-----------------------------------------

#### Benchmark execution
**To run benchmarks on your machine, just execute the line below:**
`\go_sudoku>go test -bench=.`
