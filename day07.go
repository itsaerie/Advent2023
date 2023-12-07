package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type handBid struct {
	hand     string
	bid      int
	handType int
}

// --- Part 1 Helper Functions ---

func handTypeOne(hand string) int {
	// ...do i need to make this map? dunno
	m := make(map[rune]int)
	for _, char := range hand {
		if m[char] == 0 {
			m[char] = 1
		} else {
			m[char] = m[char] + 1
		}
	}

	// find highest and next highest
	high := 0
	next := 0
	for char := range m {
		if m[char] > high {
			next = high
			high = m[char]
		} else if m[char] > next {
			next = m[char]
		}
	}

	if (high == 2 || high == 3) && next == 2 {
		return high*2 + 1
	} else {
		return high * 2
	}
}

func cardByteOneLessThan(a, b byte) bool {
	cardRating := map[byte]int{
		byte('A'): 14,
		byte('K'): 13,
		byte('Q'): 12,
		byte('J'): 11,
		byte('T'): 10,
		byte('9'): 9,
		byte('8'): 8,
		byte('7'): 7,
		byte('6'): 6,
		byte('5'): 5,
		byte('4'): 4,
		byte('3'): 3,
		byte('2'): 2,
	}
	return cardRating[a] < cardRating[b]
}

func handOneLessThan(a, b string) bool {
	for i := 0; i < 5; i++ {
		if a[i] != b[i] {
			return cardByteOneLessThan(a[i], b[i])
		}
	}

	return false
}

func (handA handBid) bidOneLessThan(handB handBid) bool {
	if handA.handType != handB.handType {
		return handA.handType < handB.handType
	}

	// only need to compare the highest occurrence
	return handOneLessThan(handA.hand, handB.hand)
}

// --- Part 2 Helper Functions
func handTypeTwo(hand string) int {
	j := 0 // count instances of 'j'
	// ...do i need to make this map? dunno
	m := make(map[rune]int)
	for _, char := range hand {
		if char == 'J' {
			j++
			continue
		}
		if m[char] == 0 {
			m[char] = 1
		} else {
			m[char] = m[char] + 1
		}
	}

	// find highest and next highest
	high := 0
	next := 0
	for char := range m {
		if m[char] > high {
			next = high
			high = m[char]
		} else if m[char] > next {
			next = m[char]
		}
	}

	// just add J to highest
	high += j

	if (high == 2 || high == 3) && next == 2 {
		return high*2 + 1
	} else {
		return high * 2
	}
}

func cardByteTwoLessThan(a, b byte) bool {
	cardRating := map[byte]int{
		byte('A'): 14,
		byte('K'): 13,
		byte('Q'): 12,
		byte('T'): 10,
		byte('9'): 9,
		byte('8'): 8,
		byte('7'): 7,
		byte('6'): 6,
		byte('5'): 5,
		byte('4'): 4,
		byte('3'): 3,
		byte('2'): 2,
		byte('J'): 1,
	}
	return cardRating[a] < cardRating[b]
}

func handTwoLessThan(a, b string) bool {
	for i := 0; i < 5; i++ {
		if a[i] != b[i] {
			return cardByteTwoLessThan(a[i], b[i])
		}
	}

	return false
}

func (handA handBid) bidTwoLessThan(handB handBid) bool {
	if handA.handType != handB.handType {
		return handA.handType < handB.handType
	}

	// only need to compare the highest occurrence
	return handTwoLessThan(handA.hand, handB.hand)
}

// --- the cake ---

func SevenOne(file *os.File) {
	scanner := bufio.NewScanner(file)

	var bids []handBid

	for scanner.Scan() {
		// make new handbid
		handData := strings.Split(scanner.Text(), " ")
		bidVal, _ := strconv.Atoi(handData[1])
		bids = append(bids, handBid{handData[0], bidVal, handTypeOne(handData[0])})
	}

	sort.Slice(bids, func(i, j int) bool {
		return bids[i].bidOneLessThan(bids[j])
	})

	total := 0
	for idx, bid := range bids {
		total += bid.bid * (idx + 1)
	}

	fmt.Println(total)
}

func SevenTwo(file *os.File) {
	scanner := bufio.NewScanner(file)

	var bids []handBid

	for scanner.Scan() {
		// make new handbid
		handData := strings.Split(scanner.Text(), " ")
		bidVal, _ := strconv.Atoi(handData[1])
		bids = append(bids, handBid{handData[0], bidVal, handTypeTwo(handData[0])})
	}

	sort.Slice(bids, func(i, j int) bool {
		return bids[i].bidTwoLessThan(bids[j])
	})

	total := 0
	for idx, bid := range bids {
		total += bid.bid * (idx + 1)
	}

	fmt.Println(total)
}
