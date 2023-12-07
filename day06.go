package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type race struct {
	time int
	dist int
}

func countRaceWins(time int, dist int) int {
	// to get the wins, you really only need to find the lowest/highest start time
	// because every start time between the lowest and highest are going to have higher values
	// (see: area of a rectangle if h+w = x)
	// this calculation is symmetrical as well
	// and the midpoint is approximately time/2
	// so, if x+y = time and you want x*y >= dist

	// start the binary search from the midpoint
	low := 0
	high := time / 2
	for low+1 < high {
		mid := (low + high) / 2
		// too high
		if mid*(time-mid) > dist {
			high = mid
		} else {
			low = mid
		}
	}
	// high is going to be the value of our objective
	// the +1 is because it's an inclusive selection
	return (time - high) - high + 1
}

func SixOne(file *os.File) {
	scanner := bufio.NewScanner(file)
	rgxInt := regexp.MustCompile("[0-9]+")
	var races []race

	scanner.Scan()
	timeStr := rgxInt.FindAllString(scanner.Text(), 999)
	scanner.Scan()
	distStr := rgxInt.FindAllString(scanner.Text(), 999)

	for idx := range timeStr {
		time, _ := strconv.Atoi(timeStr[idx])
		dist, _ := strconv.Atoi(distStr[idx])
		races = append(races, race{time, dist})
	}

	total := 1
	for _, rc := range races {
		total *= countRaceWins(rc.time, rc.dist)
	}

	fmt.Println(total)
}

func SixTwo(file *os.File) {
	scanner := bufio.NewScanner(file)
	rgxInt := regexp.MustCompile("[0-9]+")

	scanner.Scan()
	timeList := rgxInt.FindAllString(scanner.Text(), 999)
	scanner.Scan()
	distList := rgxInt.FindAllString(scanner.Text(), 999)

	timeStr := ""
	distStr := ""

	for idx := range timeList {
		timeStr += timeList[idx]
		distStr += distList[idx]
	}

	time, _ := strconv.Atoi(timeStr)
	dist, _ := strconv.Atoi(distStr)

	fmt.Println(countRaceWins(time, dist))
}
