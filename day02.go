package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func TwoOne(file *os.File) {
	scanner := bufio.NewScanner(file)

	// consts limits
	maxRed := 12
	maxGreen := 13
	maxBlue := 14

	idSum := 0

	for scanner.Scan() {
		str := scanner.Text()

		isValidGame := true
		// get the game number
		rgxInt := regexp.MustCompile("[0-9]+") // numbers
		gameNum, _ := strconv.Atoi(rgxInt.FindString(str))

		// get regexes for colors
		rgxRed := regexp.MustCompile("red")
		rgxGrn := regexp.MustCompile("green")
		rgxBlu := regexp.MustCompile("blue")

		// get the text after the colon
		gameData := strings.Split(str, ":")[1]
		for _, hand := range strings.Split(gameData, ";") {
			// verify vals in each hand
			for _, val := range strings.Split(hand, ",") {
				// redcheck
				if rgxRed.MatchString(val) {
					num, _ := strconv.Atoi(rgxInt.FindString(val))
					if num > maxRed {
						isValidGame = false
					}
				}
				// greencheck
				if rgxGrn.MatchString(val) {
					num, _ := strconv.Atoi(rgxInt.FindString(val))
					if num > maxGreen {
						isValidGame = false
					}
				}
				// bluecheck
				if rgxBlu.MatchString(val) {
					num, _ := strconv.Atoi(rgxInt.FindString(val))
					if num > maxBlue {
						isValidGame = false
					}
				}
				if !isValidGame {
					break
				}
			}
			if !isValidGame {
				break
			}
		}

		if isValidGame {
			idSum += gameNum
		}
	}

	fmt.Println(idSum)
}

func TwoTwo(file *os.File) {
	scanner := bufio.NewScanner(file)

	powSum := 0

	for scanner.Scan() {
		str := scanner.Text()

		rgxInt := regexp.MustCompile("[0-9]+") // numbers
		rgxRed := regexp.MustCompile("red")
		rgxGrn := regexp.MustCompile("green")
		rgxBlu := regexp.MustCompile("blue")

		// maxVals
		maxRed, maxGrn, maxBlu := 0, 0, 0

		// get the text after the colon
		gameData := strings.Split(str, ":")[1]
		for _, hand := range strings.Split(gameData, ";") {
			// verify vals in each hand
			for _, val := range strings.Split(hand, ",") {
				// redcheck
				if rgxRed.MatchString(val) {
					num, _ := strconv.Atoi(rgxInt.FindString(val))
					if num > maxRed {
						maxRed = num
					}
				}
				// greencheck
				if rgxGrn.MatchString(val) {
					num, _ := strconv.Atoi(rgxInt.FindString(val))
					if num > maxGrn {
						maxGrn = num
					}
				}
				// bluecheck
				if rgxBlu.MatchString(val) {
					num, _ := strconv.Atoi(rgxInt.FindString(val))
					if num > maxBlu {
						maxBlu = num
					}
				}
			}
		}

		powSum += maxRed * maxGrn * maxBlu
	}

	fmt.Println(powSum)
}
