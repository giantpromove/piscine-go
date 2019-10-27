package main

import (
	"github.com/01-edu/z01"
	"os"
)

func printStr(str string) {
	arrayStr := []rune(str)
	length := 0
	for i := range arrayStr {
		length = i + 1
	}
	for i := 0; i < length; i++ {
		z01.PrintRune(arrayStr[i])
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 != 1 {
		return true
	} else {
		return false
	}
}

func main() {
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"
	args := os.Args
	args = args[1:]
	lengthOfArg := 0
	for i := range args {
		lengthOfArg = i + 1
	}

	if isEven(lengthOfArg) == true {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
