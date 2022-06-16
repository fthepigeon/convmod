package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func digitToSymbols(d int) string {
	if d == 0 {
		return "---"
	}

	var sb strings.Builder
	symbols := [3]rune{'r', 'w', 'x'}

	for d > 0 {
		for i, k := range []int{4, 2, 1} {
			if d-k >= 0 {
				sb.WriteRune(symbols[i])
				d -= k
			} else {
				sb.WriteRune('-')
			}
		}
	}
	return sb.String()
}

func helpAndExit() {
	fmt.Println("help me!")
	os.Exit(0)
}

func symbolicToOct(v string) {
	fmt.Println("oct repr")
}

func octToSymbolic(v string) {
	octV, err := strconv.Atoi(v)

	if err != nil {
		fmt.Println("wrong permission format:", v)
		os.Exit(1)
	}
	if octV > 777 || octV < 0 {
		fmt.Println("wrong permission value:", v)
		os.Exit(1)
	}

	symbolicPerms := ""

	for i := 0; i < 3; i++ {
		digit := octV % 10

		if digit > 7 {
			fmt.Println("wrong permission value:", v)
			os.Exit(1)
		}
		symbolicPerms = digitToSymbols(digit) + symbolicPerms
		octV /= 10
	}
	fmt.Println(symbolicPerms)
}

func main() {
	switch len(os.Args) {
	case 2:
		symbolicToOct(os.Args[1])
	case 3:
		if os.Args[2] == "--oct" {
			octToSymbolic(os.Args[1])
		}
	default:
		helpAndExit()
	}
}
