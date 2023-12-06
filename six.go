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
	wins int
}

func SixOne(file *os.File) {
	scanner := bufio.NewScanner(file)
	rgxInt := regexp.MustCompile("[0-9]+")

	// get times
	var races []race
	maxTime := 0

	scanner.Scan()
	timeStr := rgxInt.FindAllString(scanner.Text(), 999)
	scanner.Scan()
	distStr := rgxInt.FindAllString(scanner.Text(), 999)

	for idx := range timeStr {
		time, _ := strconv.Atoi(timeStr[idx])
		dist, _ := strconv.Atoi(distStr[idx])
		races = append(races, race{time, dist, 0})

		if time > maxTime {
			maxTime = time
		}
	}

	for time := 1; time < maxTime; time++ {
		for idx, rc := range races {
			distMove := (rc.time - time) * time
			if distMove > rc.dist {
				races[idx] = race{rc.time, rc.dist, rc.wins + 1}
			}
		}
	}

	total := 1
	for _, rc := range races {
		total *= rc.wins
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

	wins := 0

	for t := 1; t < time; t++ {
		distMove := (time - t) * t
		if distMove > dist {
			wins++
		}
	}

	fmt.Println(wins)
}
