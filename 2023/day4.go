package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
  data, _ := os.ReadFile("day4.txt")
  lines := strings.Split(string(data), "\n")

  r, _ := regexp.Compile(`\d+`)

  copies := make(map[int]int)
  for i, line := range lines {
    if (len(line) < 1) {
      continue
    }

    parts := strings.Split(strings.Split(line, ": ")[1], "|")
    winningNumbers := r.FindAllString(parts[0], -1)
    myNumbers := r.FindAllString(parts[1], -1)

    copies[i] += 1
    total := 0
    for _, myNumber := range myNumbers {
      for _, winningNumber := range winningNumbers {
        if (myNumber == winningNumber) {
          total += 1
        }
      }
    }
    for j := 0; j < total; j++ {
      copies[i+j+1] += copies[i]
    }
    // fmt.Println(i, total, copies)
    // return
  }

  sum := 0
  for _, v := range copies {
    sum += v
  }
  fmt.Println(sum)
}
