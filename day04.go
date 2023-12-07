package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
)

func existsIn(str string, list []string) bool {
	for _, s := range list {
		if str == s {
			return true
		}
	}
	return false
}

func FourOne(file *os.File) {
	scanner := bufio.NewScanner(file)

	totalPoints := 0.

	rgxNum := regexp.MustCompile("[0-9]+")

	for scanner.Scan() {
		// split into winning and losing cards
		str := scanner.Text()

		winNums := rgxNum.FindAllString(strings.Split(strings.Split(str, ":")[1], "|")[0], 999)
		ownNums := rgxNum.FindAllString(strings.Split(strings.Split(str, ":")[1], "|")[1], 999)

		count := -1.
		for _, num := range ownNums {
			if existsIn(num, winNums) {
				count++
			}
		}

		if count >= 0 {
			totalPoints += math.Pow(2, count)
		}
	}

	fmt.Println(totalPoints)
}

func FourTwo(file *os.File) {
	scanner := bufio.NewScanner(file)

	totalPoints := 0

	rgxNum := regexp.MustCompile("[0-9]+")
	var multipliers []int

	for scanner.Scan() {
		// split into winning and losing cards
		str := scanner.Text()

		winNums := rgxNum.FindAllString(strings.Split(strings.Split(str, ":")[1], "|")[0], 999)
		ownNums := rgxNum.FindAllString(strings.Split(strings.Split(str, ":")[1], "|")[1], 999)

		wins := 0
		for _, num := range ownNums {
			if existsIn(num, winNums) {
				wins++
			}
		}

		currMult := 1
		// if no multiplier in queue
		if len(multipliers) == 0 {
			// fmt.Println("noMult", 1)
			totalPoints++
		} else {
			currMult = multipliers[0] + currMult
			totalPoints += currMult
			multipliers = multipliers[1:]
		}
		// then add to multipliers
		for i := 0; i < wins; i++ {
			if len(multipliers) <= i {
				multipliers = append(multipliers, currMult)
			} else {
				multipliers[i] = multipliers[i] + currMult
			}
		}
	}

	fmt.Println(totalPoints)
}
