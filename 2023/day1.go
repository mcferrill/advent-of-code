package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
  data, _ := os.ReadFile("day1.txt")
  lines := strings.Split(string(data), "\n")
  var numbers []int;
  var sum int
  for _, line := range lines {
    if (len(line) < 1) {
      continue
    }
    var digits []int;
    for _, char := range line {
      digit, err := strconv.Atoi(string(char))
      if (err != nil) {
        continue
      }
      digits = append(digits, digit)
    }
    numbers = append(numbers, (digits[0] * 10) + digits[len(digits)-1])
    sum += (digits[0] * 10) + digits[len(digits)-1]
  }

  fmt.Println(sum)
}
