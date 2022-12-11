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

func visibleUp(forest [][]int, x int, y int) int {
	trees := 0
	for i:=y-1; i>=0; i-- {
		trees++
		if forest[y][x] <= forest[i][x] {
			return trees
		}
	}
	return trees
}

func visibleDown(forest [][]int, x int, y int) int {
	trees := 0
	for i:=y+1; i<len(forest); i++ {
		trees++
		if forest[y][x] <= forest[i][x] {
			return trees
		}
	}
	return trees
}

func visibleLeft(forest [][]int, x int, y int) int {
	trees := 0
	for i:=x-1; i>=0; i-- {
		trees++
		if forest[y][x] <= forest[y][i] {
			return trees
		}
	}
	return trees
}

func visibleRight(forest [][]int, x int, y int) int {
	trees := 0
	for i:=x+1; i<len(forest); i++ {
		trees++
		if forest[y][x] <= forest[y][i] {
			return trees
		}
	}
	return trees
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
	max := 0

	for y:=1; y<len(forest)-1; y++ {
		for x:=1; x<forestWidth-1; x++ {
			vu := visibleUp(forest, x, y)
			vd := visibleDown(forest, x, y)
			vl := visibleLeft(forest, x, y)
			vr := visibleRight(forest, x, y)
			score := vu * vd * vl * vr
			if score > max {
				max = score
			}
		}
	}
	return max
}

func main() {
	fmt.Println(howManyVisible("../../input"))
}
