package main

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}

	comb := make([]int, n)
	for i := range comb {
		comb[i] = i
	}

	for {
		for i := 0; i < n; i++ {
			z01.PrintRune(rune(comb[i] + '0'))
		}

		if comb[0] == 10-n {
			break
		}
		z01.PrintRune(',')
		z01.PrintRune(' ')

		for i := n - 1; i >= 0; i-- {
			if comb[i] < 9-(n-1-i) {
				comb[i]++
				for j := i + 1; j < n; j++ {
					comb[j] = comb[j-1] + 1
				}
				break
			}
		}
	}
	z01.PrintRune('\n')
}

/*func main() {
	PrintCombN(1)
	PrintCombN(3)
	PrintCombN(9)
}*/