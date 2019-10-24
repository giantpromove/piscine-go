package piscine

import piscine ".."

func PrintWordsTables(table []string) {
	table := len(table)

	for i := 0; i <= table-1; i++ {
		PrintRunes(table[i], 0)
	}
}
