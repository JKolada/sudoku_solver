package main 

import (
	"testing"
	"./sudoku_solver"
)

func BenchmarkEasySudoku(b *testing.B) {
    for i := 0; i < b.N; i++ {
      a := sudoku_solver.NewSudoku(sudokuInput_easy) //sudokuInput_GOD2)
	    if a != nil {
	      a.ResolveWithoutPrinting()  
	    }
    }
}

func BenchmarkMediumSudoku(b *testing.B) {
    for i := 0; i < b.N; i++ {
      a := sudoku_solver.NewSudoku(sudokuInput_medium) //sudokuInput_GOD2)
	    if a != nil {
	      a.ResolveWithoutPrinting()  
	    }
    }
}

func BenchmarkHard1Sudoku(b *testing.B) {
    for i := 0; i < b.N; i++ {
      a := sudoku_solver.NewSudoku(sudokuInput_hard1) //sudokuInput_GOD2)
	    if a != nil {
	      a.ResolveWithoutPrinting()  
	    }
    }
}

func BenchmarkHard2Sudoku(b *testing.B) {
    for i := 0; i < b.N; i++ {
      a := sudoku_solver.NewSudoku(sudokuInput_hard2) //sudokuInput_GOD2)
	    if a != nil {
	      a.ResolveWithoutPrinting()  
	    }
    }
}

func BenchmarkHard3Sudoku(b *testing.B) {
    for i := 0; i < b.N; i++ {
      a := sudoku_solver.NewSudoku(sudokuInput_hard3) //sudokuInput_GOD2)
	    if a != nil {
	      a.ResolveWithoutPrinting()  
	    }
    }
}

func BenchmarkHard4Sudoku(b *testing.B) {
    for i := 0; i < b.N; i++ {
      a := sudoku_solver.NewSudoku(sudokuInput_hard2) //sudokuInput_GOD2)
	    if a != nil {
	      a.ResolveWithoutPrinting()  
	    }
    }
}

func BenchmarkHard5Sudoku(b *testing.B) {
    for i := 0; i < b.N; i++ {
      a := sudoku_solver.NewSudoku(sudokuInput_hard3) //sudokuInput_GOD2)
	    if a != nil {
	      a.ResolveWithoutPrinting()  
	    }
    }
}


