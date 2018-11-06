package main 

import (
	"testing"
	"./sudoku_solver"
)

func Benchmark_Sudoku_VeryEasy_Deduction(b *testing.B) {
    for i := 0; i < b.N; i++ {
      a := sudoku_solver.NewSudoku(sudokuInput_veryeasy)
	    if a != nil {
	      a.ResolveWithoutPrinting()  
	    }
    }
}

func Benchmark_Sudoku_Medium_Deduction(b *testing.B) {
    for i := 0; i < b.N; i++ {
      a := sudoku_solver.NewSudoku(sudokuInput_medium)
	    if a != nil {
	      a.ResolveWithoutPrinting()  
	    }
    }
}

func Benchmark_Sudoku_Hard1_Deduction(b *testing.B) {
    for i := 0; i < b.N; i++ {
      a := sudoku_solver.NewSudoku(sudokuInput_hard1)
	    if a != nil {
	      a.ResolveWithoutPrinting()  
	    }
    }
}

func Benchmark_Sudoku_Hard2_Deduction(b *testing.B) {
    for i := 0; i < b.N; i++ {
      a := sudoku_solver.NewSudoku(sudokuInput_hard2)
	    if a != nil {
	      a.ResolveWithoutPrinting()  
	    }
    }
}

func Benchmark_Sudoku_Hard3_Deduction(b *testing.B) {
    for i := 0; i < b.N; i++ {
      a := sudoku_solver.NewSudoku(sudokuInput_hard3)
	    if a != nil {
	      a.ResolveWithoutPrinting()  
	    }
    }
}

func Benchmark_Sudoku_Hard4_Deduction(b *testing.B) {
    for i := 0; i < b.N; i++ {
      a := sudoku_solver.NewSudoku(sudokuInput_hard4)
	    if a != nil {
	      a.ResolveWithoutPrinting()  
	    }
    }
}

func Benchmark_Sudoku_Hard5_Deduction(b *testing.B) {
    for i := 0; i < b.N; i++ {
      a := sudoku_solver.NewSudoku(sudokuInput_hard5)
	    if a != nil {
	      a.ResolveWithoutPrinting()  
	    }
    }
}


