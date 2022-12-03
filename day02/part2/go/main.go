package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	rockValue     = 1
	paperValue    = 2
	scissorsValue = 3
	lostValue     = 0
	drawValue     = 3
	wonValue      = 6
)

func calculateRoundValue(elf int, you int) int {
	v := 0
	if elf == rockValue && you == scissorsValue {
		v = lostValue
	} else if elf == rockValue && you == paperValue {
		v = wonValue
	} else if elf == paperValue && you == rockValue {
		v = lostValue
	} else if elf == paperValue && you == scissorsValue {
		v = wonValue
	} else if elf == scissorsValue && you == rockValue {
		v = wonValue
	} else if elf == scissorsValue && you == paperValue {
		v = lostValue
	} else {
		v = drawValue
	}
	return v + you
}

func decryptElf(e string) int {
	if e == "A" {
		return rockValue
	} else if e == "B" {
		return paperValue
	} else {
		return scissorsValue
	}
}

func decryptMyShape(elf int, e string) int {
	if e == "X" {
		if elf == rockValue {
			return scissorsValue
		} else if elf == paperValue {
			return rockValue
		} else {
			return paperValue
		}
	} else if e == "Y" {
		return elf
	} else {
		if elf == rockValue {
			return paperValue
		} else if elf == paperValue {
			return scissorsValue
		} else {
			return rockValue
		}
	}
}

func calculateMyScore(f string) int {
	readFile, err := os.Open(f)
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	score := 0
	for fileScanner.Scan() {
		round := strings.Fields(fileScanner.Text())
		elf := decryptElf(round[0])
		you := decryptMyShape(elf, round[1])
		score += calculateRoundValue(elf, you)
	}
	return score
}

func main() {
	fmt.Println(calculateMyScore("../../input"))
}
