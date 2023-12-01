package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
  digitWords := map[string]string{
    "one": "o1e",
    "two": "t2o",
    "three": "th3ee",
    "four": "f4ur",
    "five": "f5ve",
    "six": "s6x",
    "seven": "se7en",
    "eight": "ei8ht",
    "nine": "n9ne",
  }

  data, _ := os.ReadFile("day1.txt")
  lines := strings.Split(string(data), "\n")
  var sum int
  for _, line := range lines {
    if (len(line) < 1) {
      continue
    }

    fmt.Println(line)

    wordsFound := make(map[int]int)
    for word, value := range digitWords {
      line = strings.Replace(line, word, value, -1)
    }

    var digits []int;
    for i, char := range line {
      value, ok := wordsFound[i]
      if (ok) {
        digits = append(digits, value)
        fmt.Println(i, value)
        continue
      }

      digit, err := strconv.Atoi(string(char))
      if (err != nil) {
        continue
      }
      digits = append(digits, digit)
      fmt.Println(i, digit)
    }

    fmt.Println(digits[0], digits[len(digits)-1], (digits[0] * 10) + digits[len(digits)-1])
    fmt.Println()
    sum += (digits[0] * 10) + digits[len(digits)-1]
  }

  fmt.Println(sum)
}
