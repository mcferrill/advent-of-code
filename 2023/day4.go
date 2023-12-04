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

  sum := 0
  for _, line := range lines {
    if (len(line) < 1) {
      continue
    }
    parts := strings.Split(strings.Split(line, ": ")[1], "|")
    winningNumbers := r.FindAllString(parts[0], -1)
    myNumbers := r.FindAllString(parts[1], -1)
    total := 0
    for _, myNumber := range myNumbers {
      for _, winningNumber := range winningNumbers {
        if (myNumber == winningNumber) {
          if (total == 0) {
            total = 1
          } else {
            total *= 2
          }
        }
      }
    }
    sum += total
  }

  fmt.Println(sum)
}
