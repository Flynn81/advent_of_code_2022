package main

import (
	"bufio"
	"fmt"
 "os"
	 //"strings"
)

// const (
// 	rockValue     = 1
// 	paperValue    = 2
// 	scissorsValue = 3
// 	lostValue     = 0
// 	drawValue     = 3
// 	wonValue      = 6
// )

//fmt.Println(int("Hello"[1]))

func calculatePriority(i int) int {
	if i >= 97 {
		return i - 96
	} else {
		return i - 38
	}
}

func processRucksack(r string) int {
	compartmentSize := len(r) / 2
	for i := 0; i<compartmentSize; i++ {
		//consider creating two strings via slices and using strings.Count or strings.Contains
		for j := compartmentSize; j<len(r); j++ {
			if r[i] == r[j] {
				return calculatePriority(int(r[i]))
			}
		}
	}
	return 0
}

func sumPriorities(fn string) int {
	readFile, err := os.Open(fn)
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	sum := 0
	for fileScanner.Scan() {
		sum += processRucksack(fileScanner.Text())
	}
	return sum
}

func main() {
	fmt.Println(sumPriorities("../../input"))
}
