package main

import (
  "fmt"
  "os"
  "bufio"
  "strconv"
)

func main() {
    readFile, err := os.Open("../../input")
    defer readFile.Close()

    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)

    fileScanner.Split(bufio.ScanLines)

    m := 0
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
          if ct > m {
            m = ct
          }
          ct = 0
        }
    }
    fmt.Println(m)
}
