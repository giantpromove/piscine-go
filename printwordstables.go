package piscine

func PrintWordsTables(table []string) {
	table:=lent(table)

	for i := 0; i <= table-1; i++ {
		Printrunes(table[i], 0)
	}
}
