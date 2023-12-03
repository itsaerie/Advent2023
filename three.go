package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type point struct {
	x, y int
}

func findMatch(matrix []string, rgx *regexp.Regexp) []point {
	var list []point

	// search matrix for locations of "*" or "42"
	for y := 0; y < len(matrix); y++ {
		for _, v := range rgx.FindAllStringIndex(matrix[y], 999) {
			list = append(list, point{v[0], y})
		}
	}

	return list
}

func (p point) adjPoints() []point {
	var list []point

	list = append(list,
		point{p.x - 1, p.y - 1}, point{p.x - 1, p.y}, point{p.x - 1, p.y + 1},
		point{p.x, p.y - 1}, point{p.x, p.y + 1},
		point{p.x + 1, p.y - 1}, point{p.x + 1, p.y}, point{p.x + 1, p.y + 1},
	)

	return list
}

func (p point) existsIn(list []point) bool {
	for _, pnt := range list {
		if p == pnt {
			return true
		}
	}
	return false
}

func ThreeOne(file *os.File) {
	scanner := bufio.NewScanner(file)
	rgxNum := regexp.MustCompile("[0-9]")
	rgxNums := regexp.MustCompile("[0-9]+")
	rgxSym := regexp.MustCompile(`[^0-9A-Za-z\t\n\f\r\.]`)

	var rows []string
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}

	// matches for asterisks
	symMatches := findMatch(rows, rgxSym)
	numMatches := findMatch(rows, rgxNum)

	var numList []point
	sum := 0

	// find adjacent items to items
	for _, sym := range symMatches {
		// get all adjacent values
		for _, adj := range sym.adjPoints() {
			// check if theres a number at the point, and that the number wasnt used
			if adj.existsIn(numMatches) && !adj.existsIn(numList) {
				// search for the number
				ranges := rgxNums.FindAllStringIndex(rows[adj.y], 999)
				for _, ra := range ranges {
					if ra[0] <= adj.x && adj.x < ra[1] {
						// add all values to numList
						for i := ra[0]; i < ra[1]; i++ {
							numList = append(numList, point{i, adj.y})
						}
						val, _ := strconv.Atoi(rows[adj.y][ra[0]:ra[1]])
						sum += val
					}
				}
			}
		}
	}

	fmt.Println(sum)
}

func ThreeTwo(file *os.File) {
	scanner := bufio.NewScanner(file)
	rgxNum := regexp.MustCompile("[0-9]")
	rgxNums := regexp.MustCompile("[0-9]+")
	rgxSym := regexp.MustCompile(`\*`)

	// fill matrix from rows
	var rows []string
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}

	// matches for asterisks
	symMatches := findMatch(rows, rgxSym)
	numMatches := findMatch(rows, rgxNum)
	sum := 0

	// find adjacent items to items
	for _, sym := range symMatches {
		countAdj := 0
		mult := 1
		var numList []point
		for _, adj := range sym.adjPoints() {
			if adj.existsIn(numMatches) && !adj.existsIn(numList) {
				// search for the number
				ranges := rgxNums.FindAllStringIndex(rows[adj.y], 999)
				for _, ra := range ranges {
					if ra[0] <= adj.x && adj.x < ra[1] {
						// add all values to numList
						for i := ra[0]; i < ra[1]; i++ {
							numList = append(numList, point{i, adj.y})
						}
						// atoi it to sum
						val, _ := strconv.Atoi(rows[adj.y][ra[0]:ra[1]])
						mult *= val
						countAdj++
						break
					}
				}
			}
		}
		if countAdj == 2 {
			sum += mult
		}
	}

	fmt.Println(sum)
}
