package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type File struct {
	name string
	size int
}

type Directory struct {
	name string
	files []File
	directories []Directory
	parentDirectory *Directory
	size int
}

func printDir(d Directory) {
	fmt.Println(d)
	for _,sd := range d.directories {
		printDir(sd)
	}
}

func calcSize(d *Directory) {
	for _,f := range d.files {
		d.size += f.size
	}
	for i:=0; i<len(d.directories); i++ {
		calcSize(&d.directories[i])
		d.size += d.directories[i].size
	}
	fmt.Println(d.name)
	fmt.Println(d.size)
}

func sumDirectories(d *Directory) int {
	sum := 0
	for i:=0; i<len(d.directories); i++ {
		if d.directories[i].size <= 100000 {
			sum += d.directories[i].size
		}
		sum += sumDirectories(&d.directories[i])
	}
	return sum
}

func processLine(line string) (string, string, File, Directory) {
		if string(line[0]) == "$" {
			command := strings.Fields(line)
			if command[1] == "ls" {
				return "ls", "", File{}, Directory{}
			} else if command[2] == "/" {
				return "returnToRoot", "", File{}, Directory{}
			} else if command[2] == ".." {
				return "upOne", "", File{}, Directory{}
			} else {
				return "changeDirectory", command[2], File{}, Directory{}
			}
		} else {
			if strings.Index(line, "dir") == 0 {
				d := Directory{strings.Fields(line)[1],make([]File,0), make([]Directory, 0),nil,0}
				return "addDirectory", "", File{}, d
			} else {
				parts := strings.Fields(line)
				size,_ := strconv.Atoi(parts[0])
				f := File{parts[1], size}
				return "addFile", "", f, Directory{}
			}
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
	root := Directory{"/",make([]File,0), make([]Directory, 0), nil, 0}
	curDir := &root
	for fileScanner.Scan() {
		whatToDo, dName, f, d := processLine(fileScanner.Text())
		if whatToDo == "addDirectory" {
			d.parentDirectory = curDir
			curDir.directories = append(curDir.directories, d)
		} else if whatToDo == "changeDirectory" {
			for i := 0; i<len(curDir.directories); i++ {
				if curDir.directories[i].name == dName {
					curDir = &curDir.directories[i]
					break
				}
			}
		} else if whatToDo == "addFile" {
			curDir.files = append(curDir.files, f)
		} else if whatToDo == "upOne" {
			curDir = curDir.parentDirectory
		} else if whatToDo == "returnToRoot" {
			curDir = &root
		}
	}
	calcSize(&root)
	printDir(root)
	return sumDirectories(&root)
}

func main() {
	fmt.Println(sumSizes("../../input"))
}
