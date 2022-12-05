package main

import (
	"bufio"
	"fmt"
 "os"
	 "strings"
	 "strconv"
)

// const (
// 	rockValue     = 1
// 	paperValue    = 2
// 	scissorsValue = 3
// 	lostValue     = 0
// 	drawValue     = 3
// 	wonValue      = 6
// )



func isThereOverlap(p string) bool {
	ranges := strings.Split(p, ",")
	range1 := strings.Split(ranges[0],"-")
	r1Lower,_ := strconv.Atoi(range1[0])
	r1Upper,_ := strconv.Atoi(range1[1])
	range2 := strings.Split(ranges[1],"-")
	r2Lower,_ := strconv.Atoi(range2[0])
	r2Upper,_ := strconv.Atoi(range2[1])
	if r1Lower <= r2Lower && r1Upper >= r2Upper {
		return true
	} else if r2Lower <= r1Lower && r2Upper >= r1Upper {
		return true
	} else if r1Lower <= r2Lower && r1Upper >= r2Lower {
		return true
	} else if r2Lower <= r1Lower && r2Upper >= r1Lower {
		return true
	} else if r1Lower <= r2Upper && r1Upper >= r2Upper {
		return true
	} else if r2Lower <= r1Upper && r2Upper >= r1Upper {
		return true
	}
	return false
}

func howManyRangesOverlap(fn string) int {
	readFile, err := os.Open(fn)
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	count := 0
	for fileScanner.Scan() {
		if isThereOverlap(fileScanner.Text()) {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(howManyRangesOverlap("../../input"))
}
