package piscine

func UltimateDivMod(a *int, b *int) {
	div := *a
	*a = *a/(*b)
	*b = div % (*b)
}
