package main

import (
  "fmt"
  "os"
  "bufio"
  "strconv"
  "sort"
)

func main() {
    readFile, err := os.Open("../../input")
    defer readFile.Close()

    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)

    fileScanner.Split(bufio.ScanLines)

    topThree := []int{0,0,0}
    ct := 0
    for fileScanner.Scan() {
        t := fileScanner.Text()
        if t != "" {
          c, err := strconv.Atoi(t)
          if err != nil {
            fmt.Println(err)
          }
          ct += c
        } else {
          for i, m := range topThree {
            if ct > m {
              topThree[i] = ct
              sort.Ints(topThree)
              break
            }
          }
          ct = 0
        }
    }
    total := topThree[0] + topThree[1] + topThree[2]
    fmt.Println(total)
}
