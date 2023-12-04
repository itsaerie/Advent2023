package main

import (
	"bufio"
	"fmt"
	"os"
)

func partOne(file *os.File) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func partTwo(file *os.File) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	// open input file
	file, _ := os.Open(os.Args[1] + ".txt")
	defer file.Close()

	// arg parsing for which day to run; default to copy_paste
	part1 := partOne
	part2 := partTwo

	switch day := os.Args[1]; day {
	case "One":
		part1 = OneOne
		part2 = OneTwo
	case "Two":
		part1 = TwoOne
		part2 = TwoTwo
	case "Three":
		part1 = ThreeOne
		part2 = ThreeTwo
	case "Four":
		part1 = FourOne
		part2 = FourTwo
	default:
		part1 = partOne
		part2 = partTwo
	}

	// arg parsing for part 1 or part 2; default to part 1
	switch part := os.Args[2]; part {
	case "1":
		part1(file)
	case "2":
		part2(file)
	default:
		part1(file)
		file.Close()
		// open input file, again, since ingest of file is greedy
		file, _ = os.Open(os.Args[1] + ".txt")
		part2(file)
	}
}
