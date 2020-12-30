package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type pwTest struct {
	min int
	max int
	ltr rune
}

func (t pwTest) Pass(input string) bool {
	var count int
	for _, char := range input {
		if char == t.ltr {
			count++
		}
	}

	if count < t.min || count > t.max {
		return false
	}
	return true
}

func (t pwTest) PassPart2(input string) bool {
	inputSlice := []rune(input)

	ltrA := inputSlice[t.min-1]
	ltrB := inputSlice[t.max-1]

	// fmt.Printf("search: %c, ltrA: %c, ltrB: %c \n", t.ltr, ltrA, ltrB)

	if ltrA == t.ltr && ltrB == t.ltr {
		return false
	}

	if ltrA == t.ltr || ltrB == t.ltr {
		return true
	}

	return false
}

func readTest(input string) pwTest {
	split1 := strings.Split(input, "-")
	split2 := strings.Split(split1[1], " ")

	min, err := strconv.Atoi(split1[0])
	if err != nil {
		log.Fatalf("couldn't get a number from %s: %s", split1[0], err)
	}

	max, err := strconv.Atoi(split2[0])
	if err != nil {
		log.Fatalf("couldn't get a number from %s: %s", split2[0], err)
	}

	ltr := []rune(split2[1])[0]

	return pwTest{
		min,
		max,
		ltr,
	}
}

func readFile() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func main() {
	lines := readFile()

	// part 1
	var validCount int

	for _, line := range lines {
		splitLine := strings.Split(line, ": ")
		pwTest := readTest(splitLine[0])
		passwd := splitLine[1]

		if pwTest.Pass(passwd) {
			validCount++
		}
	}

	fmt.Printf("part 1 solution: %d valid passwords \n\n", validCount)

	// part 2
	validCount = 0
	for _, line := range lines {
		splitLine := strings.Split(line, ": ")
		pwTest := readTest(splitLine[0])
		passwd := splitLine[1]

		if pwTest.PassPart2(passwd) {
			validCount++
		}
	}

	fmt.Printf("part 2 solution: %d valid passwords \n", validCount)
}
