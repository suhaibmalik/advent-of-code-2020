package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readFile() []int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var numbers []int

	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("couldn't convert to int: %s", err)
		}

		numbers = append(numbers, n)
	}

	return numbers
}

func main() {
	// grab numbers from file
	numbers := readFile()

	// part 1
	for iA, numA := range numbers {
		for iB, numB := range numbers {
			if iA <= iB {
				continue
			}

			if (numA + numB) == 2020 {
				fmt.Printf("found sum of 2020 with %d+%d\n", numA, numB)
				fmt.Printf("part 1 solution: %d*%d = %d \n\n", numA, numB, numA*numB)
				break
			}
		}
	}

	// part 2
	for iA, numA := range numbers {
		for iB, numB := range numbers {
			if iA <= iB {
				continue
			}

			for iC, numC := range numbers {
				if iB <= iC {
					continue
				}

				if (numA + numB + numC) == 2020 {
					fmt.Printf("found sum of 2020 with %d+%d+%d\n", numA, numB, numC)
					fmt.Printf("part 2 solution: %d*%d*%d = %d \n", numA, numB, numC, numA*numB*numC)
					break
				}
			}
		}
	}
}
