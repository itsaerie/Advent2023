package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func existsInInt(i int, list []int) bool {
	for _, s := range list {
		if i == s {
			return true
		}
	}
	return false
}

func minInt(list []int) int {
	smallest := list[0]
	for _, i := range list {
		if i < smallest {
			smallest = i
		}
	}
	return smallest
}

func FiveOne(file *os.File) {
	scanner := bufio.NewScanner(file)
	rgxNum := regexp.MustCompile("[0-9]+")

	// LINE ONE, GET SEEDS
	scanner.Scan()
	seedsStr := rgxNum.FindAllString(scanner.Text(), 999)

	var oldValues, newValues, idxToSkip []int
	for _, seedStr := range seedsStr {
		val, _ := strconv.Atoi(seedStr)
		oldValues = append(oldValues, val)
	}

	for scanner.Scan() {
		line := scanner.Text()
		if rgxNum.FindString(line) == "" {
			// append all leftover oldValues
			if len(idxToSkip) < len(oldValues) {
				for idx, val := range oldValues {
					if !existsInInt(idx, idxToSkip) {
						idxToSkip = append(idxToSkip, idx)
						newValues = append(newValues, val)
					}
				}
			}
			// if there's no numbers, then we can bounce the newValues down to oldValues
			if len(newValues) > 0 {
				oldValues = newValues
				newValues = []int{}
				idxToSkip = []int{}
			}
		} else {
			// check if each val is within the range
			dstStart, _ := strconv.Atoi(rgxNum.FindAllString(line, 5)[0])
			srcStart, _ := strconv.Atoi(rgxNum.FindAllString(line, 5)[1])
			rangeLen, _ := strconv.Atoi(rgxNum.FindAllString(line, 5)[2])
			// indexes to remove
			for idx, val := range oldValues {
				if srcStart <= val && val < srcStart+rangeLen {
					newValues = append(newValues, val+(dstStart-srcStart))
					idxToSkip = append(idxToSkip, idx)
				}
			}
		}
	}

	// append all leftover oldValues
	if len(idxToSkip) < len(oldValues) {
		for idx, val := range oldValues {
			if !existsInInt(idx, idxToSkip) {
				idxToSkip = append(idxToSkip, idx)
				newValues = append(newValues, val)
			}
		}
	}
	// if there's no numbers, then we can bounce the newValues down to oldValues
	if len(newValues) > 0 {
		oldValues = newValues
	}

	fmt.Println(minInt(oldValues))
}

type rng struct {
	start int
	dist  int
}

func lowestInt(list []rng) int {
	smallest := list[0].start
	for _, ra := range list {
		if ra.start < smallest {
			smallest = ra.start
		}
	}
	return smallest
}

func FiveTwo(file *os.File) {
	scanner := bufio.NewScanner(file)
	rgxNum := regexp.MustCompile("[0-9]+")

	// LINE ONE, GET SEEDS
	scanner.Scan()
	seedsStr := rgxNum.FindAllString(scanner.Text(), 999)

	var oldValues, newValues []rng
	for idx := range seedsStr {
		if idx%2 == 1 {
			continue
		}
		val, _ := strconv.Atoi(seedsStr[idx])
		ra, _ := strconv.Atoi(seedsStr[idx+1])
		oldValues = append(oldValues, rng{val, ra})
	}

	for scanner.Scan() {
		line := scanner.Text()
		if rgxNum.FindString(line) == "" {
			// append all leftover oldValues
			for _, ra := range oldValues {
				if ra.dist == -1 {
					continue
				}
				newValues = append(newValues, ra)
			}
			// if there's no numbers, then we can bounce the newValues down to oldValues
			if len(newValues) > 0 {
				oldValues = newValues
				newValues = []rng{}
			}
		} else {
			// parsing rule
			dstStart, _ := strconv.Atoi(rgxNum.FindAllString(line, 5)[0])
			srcStart, _ := strconv.Atoi(rgxNum.FindAllString(line, 5)[1])
			rangeLen, _ := strconv.Atoi(rgxNum.FindAllString(line, 5)[2])
			diff := dstStart - srcStart
			// convert applicable parts of slices
			for idx, ra := range oldValues {
				if ra.dist == -1 {
					continue
				}
				// check if ra is outside first
				if ra.start+ra.dist < srcStart || srcStart+rangeLen < ra.start {
					continue
				}
				// if ra is fully contained in rule
				if srcStart <= ra.start && ra.start+ra.dist <= srcStart+rangeLen {
					newValues = append(newValues, rng{ra.start + diff, ra.dist})
					// mark for ignore
					ra.dist = -1
					// if left side of ra is outside the rule, but right isnt
				} else if ra.start < srcStart && ra.start+ra.dist <= srcStart+rangeLen {
					newValues = append(newValues, rng{srcStart + diff, ra.dist - (srcStart - ra.start)})
					// update inplace
					ra.dist = ra.dist - (srcStart - ra.start)
					// if right side of ra is outside rule, but left isnt
				} else if srcStart <= ra.start && srcStart+rangeLen < ra.start+ra.dist {
					newValues = append(newValues, rng{ra.start + diff, (srcStart + rangeLen) - ra.start})
					// update inplace
					ra.dist = (ra.dist + ra.start) - (srcStart + rangeLen)
					ra.start = srcStart + rangeLen
					// if ra fully contains rule
				} else if ra.start < srcStart && srcStart+rangeLen < ra.start+ra.dist {
					newValues = append(newValues, rng{srcStart + diff, rangeLen})
					// append right hand side to oldValues
					oldValues = append(oldValues, rng{srcStart + rangeLen, (ra.dist + ra.start) - (srcStart + rangeLen)})
					// update inplace for old
					ra.dist = ra.dist - (srcStart - ra.start)
				}

				oldValues[idx] = ra
			}
		}
	}

	// append all leftover oldValues
	for _, ra := range oldValues {
		if ra.dist == -1 {
			continue
		}
		newValues = append(newValues, ra)
	}

	// fmt.Println(newValues)
	fmt.Println(lowestInt(newValues))
}
