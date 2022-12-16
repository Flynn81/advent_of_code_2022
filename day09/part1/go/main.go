package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

const (
	UP = 1
	DOWN = 2
	RIGHT = 3
	LEFT = 4
)

type coordinates struct {
	x int
	y int
}

func moveHead(d int, hc coordinates) coordinates {
	switch d {
	case RIGHT:
		hc.x++
	case LEFT:
		hc.x--
	case UP:
		hc.y++
	case DOWN:
		hc.y--
	}
	return hc
}

func abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}

func moveTail(tc coordinates, hc coordinates) coordinates {
	moveHorizontal := true
	moveVertical := true
	if abs(tc.x - hc.x)  <= 1 {
		moveHorizontal = false
	}
	if abs(tc.y - hc.y)  <= 1 {
		moveVertical = false
	}
	if moveVertical {
		if tc.y > hc.y {
			tc.y--
		} else {
			tc.y++
		}
	}
	if moveHorizontal {
		if tc.x > hc.x {
			tc.x--
		} else {
			tc.x++
		}
	}
	return tc
}

func processLine(line string) (int, int) {
		d := 0
		parts := strings.Fields(line)
		switch parts[0] {
		case "R":
			d = RIGHT
		case "L":
			d = LEFT
		case "U":
			d = UP
		case "D":
			d = DOWN
		}
		m, _ := strconv.Atoi(parts[1])
		return d, m
}

func createKey(c coordinates) string {
	a := strconv.Itoa(c.x)
	b := strconv.Itoa(c.y)
	return a+"+"+b
}

func howManyUniqueVisited(fn string) int {
	readFile, err := os.Open(fn)
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	tc := coordinates{0,0}
	hc := coordinates{0,0}
	tailVisits := make(map[string]bool)
	for fileScanner.Scan() {
		d, m := processLine(fileScanner.Text())
		for i:=0; i<m; i++ {
			hc = moveHead(d, hc)
			fmt.Println(hc)
			tc = moveTail(tc, hc)
			fmt.Println(tc)
			tailVisits[createKey(tc)] = true
		}
	}

	return len(tailVisits)
}

func main() {
	fmt.Println(howManyUniqueVisited("../../input"))
}
