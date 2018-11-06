DESCRIPTION AND REMARKS:
The application aims to solve quickly any sudoku puzzle.

At first, it uses all implemented deduction algorithms with no blind guessing. If the puzzle is too complicated (but still valid), the backtracking algorithm is used. It is a "brute force" algorithm that iterates through all cells in sudoku, row by row, trying to match any number as potential solution. If it sees that it is impossible to fill another cell, it comes back to the previous one and tries an another possible cell value.

Remarks:
- Of course the main plan for the application is to solve any puzzle efficiently with no brute force. There still are a lot of complex deduction algorithms waiting to be implemented.
- Benchmark results can change, especially for deduction algorithms, because there was no efficiency correction implemented.

---------------------------------------------------
HOW TO USE IT, AS A CONSOLE APP AND A GOLANG MODULE

How to use an app as a console app:

1. Build the application with "go build" command in the folder. 

2. Execute the line below to see the demonstration of solving a sudoku puzzle.
\go_sudoku>go_sudoku.exe

3. To receive a solution for a specific sudoku puzzle, write all numbers of the sudoku in one line as an argument for the application execution. Use blank space or zero number for not-filled cells. Lines below execute the application for the same puzzle. 

\go_sudoku>go_sudoku.exe "       3  1   26  5  7    9  29  1   7       4   3  5  9   48      6   23  5   7 "
\go_sudoku>go_sudoku.exe 000000030010002600500700009002900100070000000400030050090004800000060002300500070

How to use an app as a golang module:

import "sudoku_solver" // use with the right path of the sudoku_solver folder

// and in code
			a := sudoku_solver.NewSudoku(parsedInput) 
			// -> initialize the main struct with [9][9]uint8 table where any blank space to fill in sudoku is 0
			if a != nil {
			  a.Resolve()
			  // -> the constructor method returns sudoku only if input sudoku is correct
			  // then it is possible to execute the sudoku solver algorithms sequence
			}

---------------------------------------------------
PLANS FOR FUTURE

Plan for the updates:
- Finishing the X-Wing algorithm
- Implementation of the Swordfish algorithm
- Implementation of the X-Cycles algorithm
- Implementation of other deduction algorithms
- Implementation of sudoku solving history saving possibility


-----------------------------------------------------------
BENCHMARK RESULTS

ThinkPad T540p
i7-4710MQ CPU @ 2,50GHz
16,0 GB RAM
The most updated results are in excel file in main repository folder.

Sudoku lvl  puzzle     Backtracking     Backtracking     Deduction
            number     by row [ns]      by block [ns]    algorithms
Very Easy   1          0,0000046        0,0000056        0,0000199
Medium      2          0,0003938        0,0006071        0,0000323
Hard        3          0,0001613        0,0002212        0,0000745
Hard        4          0,0006751        0,0008075        0,0001196
Hard        5          0,0002767        0,0003110        0,0000934
Hard        6          0,0003598        0,0002155        0,0001183
Hard        7          0,0059266        0,0092348        0,0000777
God         8          0,0878996        0,0480385        x
God         9          0,0033681        0,0012883        x
God         10         0,0021168        0,0016321        x
Hardest     11         0,0080305        0,0192789        x


Averages on level
-------------------------------------------------------
Sudoku lvl  Backtracking     Backtracking     Deduction
            by row [ns]      by block [ns]    algorithms
Very Easy   0,0000046        0,0000056        0,0000199
Medium      0,0003938        0,0006071        0,0000323
Hard        0,0014799        0,0021580        0,0000967
God         0,0253537        0,0175594        x
ALL LEVELS  0,0099284        0,0074219        x (cant solve the God, so it does not matter)


---------------------------------------------------
BENCHMARK EXECUTION
To run benchmarks on your machine, just execute the line below:
\go_sudoku>go test -bench=.

