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

// func processLine(line string, curDir Directory, rootDir Directory) Directory {
// 	fmt.Println("yoda " + curDir.name)
// 	if string(line[0]) == "$" {
// 		command := strings.Fields(line)
// 		if command[1] == "ls" {
// 			return curDir//ignore
// 		} else if command[2] == "/" {
// 			curDir = rootDir
// 		} else if command[2] == ".." {
// 			curDir = *curDir.parentDirectory
// 		} else {
// 			for _,d := range curDir.directories {
// 				if d.name == command[2] {
// 					curDir = d
// 					break;
// 				}
// 			}
// 		}
// 	} else {
// 		if strings.Index(line, "dir") == 0 {
// 			d := Directory{strings.Fields(line)[1],make([]File,0), make([]Directory, 0),&curDir,0}
// 			curDir.directories = append(curDir.directories, d)
// 		} else {
// 			parts := strings.Fields(line)
// 			size,_ := strconv.Atoi(parts[0])
// 			f := File{parts[1], size}
// 			curDir.files = append(curDir.files, f)
// 		}
// 	}
// 	fmt.Println(len(curDir.directories))
// 	fmt.Println(len(curDir.files))
//
// 	return curDir
// }
//
// func sizeOf(d Directory) {
// 	fmt.Println(d.name)
// 	for _,subDir := range d.directories {
// 		fmt.Println(subDir.name)
// 		sizeOf(subDir)
// 		d.size += subDir.size
// 	}
// 	for _,f := range d.files {
// 		d.size += f.size
// 	}
// 	fmt.Println(d.size)
// }
//
func printDir(d Directory) {
	fmt.Println(d)
	for _,sd := range d.directories {
		printDir(sd)
	}
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
			d.parentDirectory = curDir //maybechange to :=
			(*curDir).directories = append((*curDir).directories, d)
		} else if whatToDo == "changeDirectory" {
			for _, sd := range (*curDir).directories {
				if sd.name == dName {
					fmt.Println(len((*curDir).directories))
					curDir = &sd
					break
				}
			}
		} else if whatToDo == "addFile" {
			(*curDir).files = append((*curDir).files, f)
		} else if whatToDo == "upOne" {
			curDir = (*curDir).parentDirectory
		} else if whatToDo == "returnToRoot" {
			curDir = &root
		}
	}

	printDir(root)
	// sizeOf(rootDir)
	return root.size
}

func main() {
	fmt.Println(sumSizes("../../test-input"))
}
