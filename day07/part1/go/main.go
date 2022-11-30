package main

import (
	"bufio"
	"fmt"
	"os"
	//"strings"
)

type File struct {
	name string
	size int
}

type Directory struct {
	name string
	files []File
	directories []Directory
}

func processLine(line string) {
	if string(line[0]) == "$" {
		fmt.Println("command")
	} else {
		fmt.Println("ls output")
	}
}

func sumSizes(fn string) int {
	readFile, err := os.Open(fn)
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		processLine(fileScanner.Text())
	}
	return 0
}

func main() {
	fmt.Println(sumSizes("../../test-input"))
}
