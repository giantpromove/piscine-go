package main

import (
	
	"os"

	piscine ".."
	"github.com/01-edu/z01"
)

func printStr(str string) {
	arrayStr := []rune(str)

	for _, s := range arrayStr {
		z01.PrintRune(s)
	}
	z01.PrintRune('\n')
}


func even(nbr int) bool {

	if nbr%2 == 0 {
		return true
	}
	return false
}
	}

	//If cell is already filled, skip it, go to next value in row
	if temp_board[row][col] != '.' {
		*board = temp_board
		return solve(board, row, col+1)
	}

	for i := 1; i <= 9; i++ {
		r := getRune(i)

		*board = temp_board
		if isValid(board, row, col, r) {
			//Place this rune to board after checking it validity
			temp_board[row][col] = r
			//Check if it works for next values
			*board = temp_board
			if solve(board, row, col+1) {
				return true
			}
			//Condition when it does not work, Return empty entry
			temp_board[row][col] = '.'
			*board = temp_board
		}
	}
	return false
}



func isEven(nbr int) bool {

	if even(nbr) == true {
		return true
	}
	return false
}

func main() {

	EvenMsg:= "I have an even number of arguments\n"
	OddMsg:= "I have an odd number of arguments\n"

	lenOfArg := piscine.Lent3(os.Args) - 1

	if isEven(lenOfArg) {
		piscine.PrintStr(EvenMsg)
	} else {
		piscine.PrintStr(OddMsg)
	}
}
