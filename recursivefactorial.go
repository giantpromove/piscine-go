package main 

import "fmt"

func RecursiveFactorial(nb int) int {

	if nb == 1 || nb == 0 {
		return 1
	}
	if nb > 50 || nb < 0 {
		return 0
	} else {
		return nb * RecursiveFactorial(nb-1)
	}
	return 0
}
