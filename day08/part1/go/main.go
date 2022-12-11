package main

import (
	"bufio"
	"fmt"
	"os"
	//"strings"
	"strconv"
)

func processLine(line string) []int {
		forestRow := make([]int, 0)
		for i:=0; i<len(line); i++ {
			tree,_ := strconv.Atoi(string(line[i]))
			forestRow = append(forestRow, tree)
		}
		return forestRow
}

func visibleUp(forest [][]int, x int, y int) bool {
	for i:=0; i<y; i++ {
		if forest[y][x] <= forest[i][x] {
			return false
		}
	}
	return true
}

func visibleDown(forest [][]int, x int, y int) bool {
	for i:=len(forest[0])-1; i>y; i-- {
		if forest[y][x] <= forest[i][x] {
			return false
		}
	}
	return true
}

func visibleLeft(forest [][]int, x int, y int) bool {
	for i:=0; i<x; i++ {
		if forest[y][x] <= forest[y][i] {
			return false
		}
	}
	return true
}

func visibleRight(forest [][]int, x int, y int) bool {
	for i:=len(forest)-1; i>x; i-- {
		if forest[y][x] <= forest[y][i] {
			return false
		}
	}
	return true
}

func howManyVisible(fn string) int {
	readFile, err := os.Open(fn)
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	forest := make([][]int, 0)

	for fileScanner.Scan() {
		forest = append(forest, processLine(fileScanner.Text()))
	}

	forestWidth := len(forest[0])
	visible := len(forest[0]) * 2 + len(forest) * 2 - 4
	fmt.Println(visible)
	for y:=1; y<len(forest)-1; y++ {
		for x:=1; x<forestWidth-1; x++ {
			vu := visibleUp(forest, x, y)
			vd := visibleDown(forest, x, y)
			vl := visibleLeft(forest, x, y)
			vr := visibleRight(forest, x, y)
			if vu || vd || vl || vr {
				visible++
			}
		}
	}
	return visible
}

func main() {
	fmt.Println(howManyVisible("../../input"))
}
