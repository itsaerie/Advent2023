package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func OneOne(file *os.File) {
	scanner := bufio.NewScanner(file)
	sum := 0

	// regex
	re := regexp.MustCompile("[0-9]")

	tens := -1
	ones := -1

	for scanner.Scan() {
		// take a number out of the text
		str := scanner.Text()

		// take out left
		leftStr := re.FindString(str)
		// take out right
		rightStr := ""
		index := len(str)
		for rightStr == "" {
			rightStr = re.FindString(str[index:])
			index--
		}

		// atoi them
		tens, _ = strconv.Atoi(leftStr)
		ones, _ = strconv.Atoi(rightStr)
		sum += tens*10 + ones
	}

	fmt.Println(sum)
}

func atoi(str string) int {
	i, err := strconv.Atoi(str)

	if err != nil {
		switch str {
		case "one":
			return 1
		case "two":
			return 2
		case "three":
			return 3
		case "four":
			return 4
		case "five":
			return 5
		case "six":
			return 6
		case "seven":
			return 7
		case "eight":
			return 8
		case "nine":
			return 9
		default:
			return -1
		}
	}

	return i
}

func OneTwo(file *os.File) {
	scanner := bufio.NewScanner(file)
	sum := 0

	// regex
	re := regexp.MustCompile("[0-9]|(one)|(two)|(three)|(four)|(five)|(six)|(seven)|(eight)|(nine)")

	tens := -1
	ones := -1

	for scanner.Scan() {
		// take a number out of the text
		str := scanner.Text()

		// take out left
		leftStr := re.FindString(str)
		// take out right
		rightStr := ""
		index := len(str)
		for rightStr == "" {
			rightStr = re.FindString(str[index:])
			index--
		}

		// atoi them
		tens = atoi(leftStr)
		ones = atoi(rightStr)
		sum += tens*10 + ones
	}

	fmt.Println(sum)
}
