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

func processGroup(g [3]string) int {
	for x := 0; x<len(g[0]); x++ {
		for y := 0; y<len(g[1]); y++ {
			for z := 0; z<len(g[2]); z++ {
				if g[0][x] == g[1][y] && g[0][x] == g[2][z] {
					return calculatePriority(int(g[0][x]))
				}
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
	i := 0
	group := [3]string{"","",""}
	for fileScanner.Scan() {
		group[i] = fileScanner.Text()
		i++
		if i % 3 == 0 {
			sum += processGroup(group)
			i = 0
		}
	}
	return sum
}

func main() {
	fmt.Println(sumPriorities("../../input"))
}
