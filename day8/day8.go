package day8

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


/**
  0:      1:      2:      3:      4:
 aaaa    ....    aaaa    aaaa    ....
b    c  .    c  .    c  .    c  b    c
b    c  .    c  .    c  .    c  b    c
 ....    ....    dddd    dddd    dddd
e    f  .    f  e    .  .    f  .    f
e    f  .    f  e    .  .    f  .    f
 gggg    ....    gggg    gggg    ....

  5:      6:      7:      8:      9:
 aaaa    aaaa    aaaa    aaaa    aaaa
b    .  b    .  .    c  b    c  b    c
b    .  b    .  .    c  b    c  b    c
 dddd    dddd    ....    dddd    dddd
.    f  e    f  .    f  e    f  .    f
.    f  e    f  .    f  e    f  .    f
 gggg    gggg    ....    gggg    gggg
 */

func ContainsAll(digit, signal string) bool {
	ExternalLoop:
	for _, dC := range digit {
		for _, sC := range signal {
			if dC == sC {
				continue ExternalLoop
			}
		}
		return false
	}
	return true
}

func CountContained(digit, signal string) int {
	count := 0
	Loop:
	for _, dC := range digit {
		for _, sC := range signal {
			if dC == sC {
				count++
				continue Loop
			}
		}
	}
	return count

}

func ExtractMap(signals, digits []string) int {
	result := ""
	mapDigits := make(map[string]string, 0)
	for _, signal := range signals {
		s := len(signal)
		switch s {
		case 2:
			mapDigits["1"] = signal
			break
		case 4:
			mapDigits["4"] = signal
			break
		case 3:
			mapDigits["7"] = signal
			break
		default:
			// Nothing
		}
	}

	for i := 0; i < 4; i++ {
		digit := digits[i]
		s := len(digit)
		switch s {
		case 2:
			result += "1"
			break
		case 4:
			result += "4"
			break
		case 3:
			result += "7"
			break
		case 7:
			result += "8"
			break
		case 6:
			if ContainsAll(mapDigits["4"], digit) {
				result += "9"
			} else if ContainsAll(mapDigits["1"], digit) {
				result += "0"
			} else {
				result += "6"
			}
			break
		default:
			if ContainsAll(mapDigits["1"], digit) {
				result += "3"
			} else if CountContained(mapDigits["4"], digit) == 3 {
				result += "5"
			} else {
				result += "2"
			}
		}

	}
	r, _ := strconv.Atoi(result)
	return r
}

func Solution() {
	file, _ := os.Open("day8/input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	count := 0

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		parts := strings.Split(line, "|")
		signals := strings.Fields(parts[0])
		digits := strings.Fields(parts[1])
		r := ExtractMap(signals, digits)
		fmt.Println(r)
		count += r

	}

	fmt.Println(count)
}
