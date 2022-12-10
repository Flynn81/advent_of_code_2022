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
}

func getAllDirs(d *Directory) []*Directory {
	dirs := make([]*Directory, 0)
	dirs = append(dirs, d)
	for i:=0; i<len(d.directories); i++ {
		dirs = append(dirs, getAllDirs(&d.directories[i])...)
	}
	return dirs
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

func sizeToDelete(fn string) int {
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
	unused := 70000000 - root.size
	smallest := 70000000
	allDirs := getAllDirs(&root)
	for i:=0; i<len(allDirs); i++ {
		if allDirs[i].size + unused >= 30000000 && allDirs[i].size <= smallest {
			smallest = allDirs[i].size
		}
	}
	return smallest
}

func main() {
	fmt.Println(sizeToDelete("../../input"))
}
