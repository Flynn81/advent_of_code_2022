package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func enqueue(queue []string, element string) []string {
    queue = append(queue, element)
    return queue
}

func addToFront(queue []string, element []string) []string {
    queue = append(element, queue...)
    return queue
}

func dequeue(queue []string, n int) ([]string, []string) {
    element := queue[0:n]
		e := make([]string, len(element))
		copy(e, element)
    return queue[n:], e
}

func parseStacks(containers string) [][]string {

	rows := strings.Split(containers, "\n")

	numStacks := len(strings.Fields(rows[len(rows)-2]))

	stacks := make([][]string, numStacks)
	for i := 0; i<numStacks; i++ {
		stacks[i] = make([]string, 0)
	}

	for i := 0; i<len(rows)-2; i++ {
		for j := 0; j<=(len(rows[len(rows)-2])+6)/5; j++ {
			if j*4 + 1 < len(rows[i]) && string(rows[i][j*4 + 1]) != " " {
				stacks[j] = enqueue(stacks[j], string(rows[i][j*4 + 1]))
			}
		}
	}

	return stacks
}

func makeMove(move string, stacks [][]string) {
	f := strings.Fields(move)
	numberToMove,_ := strconv.Atoi(f[1])
	fromIndex,_ := strconv.Atoi(f[3])
	fromIndex--
	toIndex,_ := strconv.Atoi(f[5])
	toIndex--
	s := make([]string,0)

	stacks[fromIndex], s = dequeue(stacks[fromIndex],numberToMove)
	stacks[toIndex] = addToFront(stacks[toIndex], s)
}

func tops(stacks [][]string) string {
	t := ""
	for _, stack := range stacks {
		if len(stack) > 0 {
			t += string(stack[0])
		} else {
			t += " "
		}
	}
	return t
}

func topOfEachStack(fn string) string {
	readFile, err := os.Open(fn)
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	containers := ""
	moves := make([]string, 0)
	readingStacks := true
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			readingStacks = false
		} else if readingStacks {
			containers += line + "\n"
		} else {
			moves = append(moves, line)
		}
	}

	stacks := parseStacks(containers)
	for _, move := range moves {
		makeMove(move, stacks)
	}

	return tops(stacks)
}

func main() {
	fmt.Println(topOfEachStack("../../input"))
}
