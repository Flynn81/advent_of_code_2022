package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (msgLength = 4) //change to 14 for part2

func findMarkerStart(s string) int {
	index := 0
	for i,_ := range s {
		marker := string(s[i:i+msgLength])
		valid := true
		for j:=0; j<len(marker); j++ {
			if strings.Count(marker, string(marker[j])) > 1 {
				valid = false
				break
			}
		}
		if valid {
			index = i
			break
		}
	}
	return index + msgLength
}

func charactersToProcess(fn string) int {
	readFile, err := os.Open(fn)
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	fileScanner.Scan()
	return findMarkerStart(fileScanner.Text())
}

func main() {
	fmt.Println(charactersToProcess("../../input"))
}
