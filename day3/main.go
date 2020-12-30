package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

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

func onTree(line string, index int) bool {
	length := len(line)

	i := index - ((index / length) * length)
	if line[i] == '#' {
		return true
	}
	return false
}

func main() {
	lines := readFile()

	// part 1
	var onTreeCount int
	for i, ln := range lines {
		p := i * 3
		if onTree(ln, p) {
			onTreeCount++
		}
	}

	fmt.Printf("part 1 solution: %d tree collisions \n\n", onTreeCount)

	// part 2
	var treeCounts []int

	for _, n := range []int{1, 3, 5, 7} {
		onTreeCount = 0
		for i, ln := range lines {
			p := i * n
			if onTree(ln, p) {
				onTreeCount++
			}
		}
		treeCounts = append(treeCounts, onTreeCount)
	}

	// special loop for right 1, down 2
	onTreeCount = 0
	for i, ln := range lines {
		if i%2 != 0 { // number is odd
			continue
		}

		p := i / 2
		if onTree(ln, p) {
			onTreeCount++
		}
	}
	treeCounts = append(treeCounts, onTreeCount)

	treeProduct := 1
	for _, n := range treeCounts {
		treeProduct = treeProduct * n
	}

	fmt.Printf("part 2 solution: product of %v = %d \n", treeCounts, treeProduct)
}
