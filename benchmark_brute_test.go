package main 

import (
	"testing"
	"./sudoku_solver"
)

/////////////////////////////////////////////////////////

func Benchmark_Sudoku_VeryEasy_Brute_Row(b *testing.B) {
    for i := 0; i < b.N; i++ {
      a := sudoku_solver.NewSudoku(sudokuInput_veryeasy)
	    if a != nil {
	      a.SolveByRowBacktracking()  
	    }
    }
}

func Benchmark_Sudoku_VeryEasy_Brute_Block(b *testing.B) {
    for i := 0; i < b.N; i++ {
      a := sudoku_solver.NewSudoku(sudokuInput_veryeasy)
	    if a != nil {
	      a.SolveByBlockBacktrackingVER2()  
	    }
    }
}

/////////////////////////////////////////////////////////

func Benchmark_Sudoku_Medium_Brute_Row(b *testing.B) {
    for i := 0; i < b.N; i++ {
      a := sudoku_solver.NewSudoku(sudokuInput_medium)
	    if a != nil {
	      a.SolveByRowBacktracking()  
	    }
    }
}

func Benchmark_Sudoku_Medium_Brute_Block(b *testing.B) {
    for i := 0; i < b.N; i++ {
      a := sudoku_solver.NewSudoku(sudokuInput_medium)
	    if a != nil {
	      a.SolveByBlockBacktrackingVER2()  
	    }
    }
}

/////////////////////////////////////////////////////////

func Benchmark_Sudoku_Hard1_Brute_Row(b *testing.B) {
    for i := 0; i < b.N; i++ {
      a := sudoku_solver.NewSudoku(sudokuInput_hard1)
	    if a != nil {
	      a.SolveByRowBacktracking()  
	    }
    }
}

func Benchmark_Sudoku_Hard1_Brute_Block(b *testing.B) {
    for i := 0; i < b.N; i++ {
      a := sudoku_solver.NewSudoku(sudokuInput_hard1)
	    if a != nil {
	      a.SolveByBlockBacktrackingVER2()  
	    }
    }
}

/////////////////////////////////////////////////////////

func Benchmark_Sudoku_Hard2_Brute_Row(b *testing.B) {
    for i := 0; i < b.N; i++ {
      a := sudoku_solver.NewSudoku(sudokuInput_hard2)
	    if a != nil {
	      a.SolveByRowBacktracking()  
	    }
    }
}

func Benchmark_Sudoku_Hard2_Brute_Block(b *testing.B) {
    for i := 0; i < b.N; i++ {
      a := sudoku_solver.NewSudoku(sudokuInput_hard2)
	    if a != nil {
	      a.SolveByBlockBacktrackingVER2()  
	    }
    }
}

/////////////////////////////////////////////////////////

func Benchmark_Sudoku_Hard3_Brute_Row(b *testing.B) {
    for i := 0; i < b.N; i++ {
      a := sudoku_solver.NewSudoku(sudokuInput_hard3)
	    if a != nil {
	      a.SolveByRowBacktracking()  
	    }
    }
}

func Benchmark_Sudoku_Hard3_Brute_Block(b *testing.B) {
    for i := 0; i < b.N; i++ {
      a := sudoku_solver.NewSudoku(sudokuInput_hard3)
	    if a != nil {
	      a.SolveByBlockBacktrackingVER2()  
	    }
    }
}

/////////////////////////////////////////////////////////

func Benchmark_Sudoku_Hard4_Brute_Row(b *testing.B) {
    for i := 0; i < b.N; i++ {
      a := sudoku_solver.NewSudoku(sudokuInput_hard4)
	    if a != nil {
	      a.SolveByRowBacktracking()  
	    }
    }
}

func Benchmark_Sudoku_Hard4_Brute_Block(b *testing.B) {
    for i := 0; i < b.N; i++ {
      a := sudoku_solver.NewSudoku(sudokuInput_hard4)
	    if a != nil {
	      a.SolveByBlockBacktrackingVER2()  
	    }
    }
}

/////////////////////////////////////////////////////////

func Benchmark_Sudoku_Hard5_Brute_Row(b *testing.B) {
    for i := 0; i < b.N; i++ {
      a := sudoku_solver.NewSudoku(sudokuInput_hard5)
	    if a != nil {
	      a.SolveByRowBacktracking()  
	    }
    }
}

func Benchmark_Sudoku_Hard5_Brute_Block(b *testing.B) {
    for i := 0; i < b.N; i++ {
      a := sudoku_solver.NewSudoku(sudokuInput_hard5)
	    if a != nil {
	      a.SolveByBlockBacktrackingVER2()  
	    }
    }
}
/////////////////////////////////////////////////////////

func Benchmark_Sudoku_God1_Brute_Row(b *testing.B) {
    for i := 0; i < b.N; i++ {
      a := sudoku_solver.NewSudoku(sudokuInput_GOD1)
	    if a != nil {
	      a.SolveByRowBacktracking()  
	    }
    }
}

func Benchmark_Sudoku_God1_Brute_Block(b *testing.B) {
    for i := 0; i < b.N; i++ {
      a := sudoku_solver.NewSudoku(sudokuInput_GOD1)
	    if a != nil {
	      a.SolveByBlockBacktrackingVER2()  
	    }
    }
}

/////////////////////////////////////////////////////////

func Benchmark_Sudoku_God2_Brute_Row(b *testing.B) {
    for i := 0; i < b.N; i++ {
      a := sudoku_solver.NewSudoku(sudokuInput_GOD2)
	    if a != nil {
	      a.SolveByRowBacktracking()  
	    }
    }
}

func Benchmark_Sudoku_God2_Brute_Block(b *testing.B) {
    for i := 0; i < b.N; i++ {
      a := sudoku_solver.NewSudoku(sudokuInput_GOD2)
	    if a != nil {
	      a.SolveByBlockBacktrackingVER2()  
	    }
    }
}

/////////////////////////////////////////////////////////

func Benchmark_Sudoku_God3_Brute_Row(b *testing.B) {
    for i := 0; i < b.N; i++ {
      a := sudoku_solver.NewSudoku(sudokuInput_GOD3)
	    if a != nil {
	      a.SolveByRowBacktracking()  
	    }
    }
}

func Benchmark_Sudoku_God3_Brute_Block(b *testing.B) {
    for i := 0; i < b.N; i++ {
      a := sudoku_solver.NewSudoku(sudokuInput_GOD3)
	    if a != nil {
	      a.SolveByBlockBacktrackingVER2()  
	    }
    }
}


/////////////////////////////////////////////////////////

func Benchmark_Sudoku_HARDEST_Brute_Row(b *testing.B) {
    for i := 0; i < b.N; i++ {
      a := sudoku_solver.NewSudoku(sudokuInput_HARDEST)
	    if a != nil {
	      a.SolveByRowBacktracking()  
	    }
    }
}

func Benchmark_Sudoku_HARDEST_Brute_Block(b *testing.B) {
    for i := 0; i < b.N; i++ {
      a := sudoku_solver.NewSudoku(sudokuInput_HARDEST)
	    if a != nil {
	      a.SolveByBlockBacktrackingVER2()  
	    }
    }
}


