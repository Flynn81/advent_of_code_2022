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

func directlyDown(tc coordinates, hc coordinates) bool {
	if tc.x == hc.x && tc.y - hc.y == 2 {
		return true
	}
	return false
}

func directlyUp(tc coordinates, hc coordinates) bool {
	if tc.x == hc.x && hc.y - tc.y == 2 {
		return true
	}
	return false
}

func directlyLeft(tc coordinates, hc coordinates) bool {
	if tc.y == hc.y && tc.x - hc.x == 2 {
		return true
	}
	return false
}

func directlyRight(tc coordinates, hc coordinates) bool {
	if tc.y == hc.y && hc.x - tc.x == 2 {
		return true
	}
	return false
}

func touching(tc coordinates, hc coordinates) bool {
	if abs(tc.x - hc.x) <= 1 && abs(tc.y - hc.y) <= 1 {
		return true
	}
	return false
}

func moveTail(tc coordinates, hc coordinates) coordinates {

	du := directlyUp(tc,hc)
	dd := directlyDown(tc,hc)
	dl := directlyLeft(tc,hc)
	dr := directlyRight(tc,hc)
	t := touching(tc,hc)
	if du {
		tc.y++
	} else if dd {
		tc.y--
	} else if dl {
		tc.x--
	} else if dr {
		tc.x++
	} else if tc.y == hc.y && tc.x == hc.x {
		return tc
	} else if !t {
			if tc.x > hc.x {
				tc.x--
			} else {
				tc.x++
			}
			if tc.y > hc.y {
				tc.y --
			} else {
				tc.y++
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
	tailVisits[createKey(tc)] = true
	for fileScanner.Scan() {
		d, m := processLine(fileScanner.Text())
		for i:=0; i<m; i++ {
			fmt.Println("***************************")
			hc = moveHead(d, hc)
			fmt.Println(hc)
			tc = moveTail(tc, hc)
			fmt.Println(tc)
			tailVisits[createKey(tc)] = true
			for y:=10; y>-10; y-- {
				for x:=-10; x<10; x++ {
					if x == tc.x && y == tc.y {
						fmt.Print("T")
					} else if x == hc.x && y == hc.y {
						fmt.Print("H")
					} else {
						fmt.Print(".")
					}
				}
				fmt.Println("")
			}
		}
	}

	return len(tailVisits)
}

func main() {
	fmt.Println(howManyUniqueVisited("../../input"))
}
