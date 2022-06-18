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

func symbolicToOct(mod string) (string, error) {
	if len(mod) < 9 || len(mod) > 10 {
		return "", fmt.Errorf("wrong permission value '%s': invalid length", mod)
	}
	if len(mod) == 10 {
		mod = mod[1:]
	}

	var sb strings.Builder
	symbols := [4]rune{'r', 'w', 'x', '-'}
	symVals := map[rune]int{
		'r': 4,
		'w': 2,
		'x': 1,
		'-': 0,
	}

	for i := 0; i < 3; i++ {
		sum := 0

		for j, c := range mod[i*3 : (i+1)*3] {
			if c != symbols[j] && c != symbols[3] {
				return "", fmt.Errorf("wrong permission value '%s': invalid character at index %d", mod, (i*3)+j)
			}
			sum += symVals[c]
		}
		sb.WriteString(strconv.FormatInt(int64(sum), 10))
	}
	return sb.String(), nil
}

func octToSymbolic(mod string) (string, error) {
	octMod, err := strconv.Atoi(mod)

	if err != nil {
		return "", fmt.Errorf("unable to convert permission value: %s", mod)
	}
	if octMod > 777 || octMod < 0 {
		return "", fmt.Errorf("wrong permission value: %s", mod)
	}

	symbolicPerms := ""

	for i := 0; i < 3; i++ {
		digit := octMod % 10
		symbolicPerms = digitToSymbols(digit) + symbolicPerms
		octMod /= 10
	}
	return symbolicPerms, nil
}

func main() {
	var ret string
	var err error

	switch len(os.Args) {
	case 2:
		ret, err = symbolicToOct(os.Args[1])
	case 3:
		if os.Args[1] == "--oct" {
			ret, err = octToSymbolic(os.Args[2])
		} else {
			helpAndExit()
		}
	default:
		helpAndExit()
	}

	if err != nil {
		fmt.Printf("convmod: %s\n", err)
		os.Exit(1)
	}

	fmt.Println(ret)
}
