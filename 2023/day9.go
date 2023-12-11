package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func nextForRow(row []int) int {
  var diffs []int
  nonZero := false
  for i, number := range row {
    if number != 0 {
      nonZero = true
    }
    if i > 0 {
      last := row[i-1]
      diffs = append(diffs, number - last)
    }
  }
  if len(row) == 1 {
    return 0
  } else if nonZero {
    return row[len(row)-1] + nextForRow(diffs)
  } else {
    return row[len(row)-1]
  }
}

func main() {
  data, _ := os.ReadFile("day9.txt")
  lines := strings.Split(string(data), "\n")

  sum := 0
  for _, line := range lines {
    if line == "" {
      continue
    }
    var numbers []int
    for _, number := range strings.Split(line, " ") {
      number, _ := strconv.Atoi(number)
      numbers = append(numbers, number)
    }

    sum += nextForRow(numbers)
  }

  fmt.Println(sum)
}
