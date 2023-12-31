package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func nextForRow(row []int, reverse bool) int {
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
  end := len(row) - 1
  if reverse {
    end = 0
  }
  if len(row) == 1 {
    return 0
  } else if nonZero {
    offset := nextForRow(diffs, reverse)
    if reverse {
      offset = -offset
    }
    return row[end] + offset
  } else {
    return row[end]
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

    sum += nextForRow(numbers, true)
  }

  fmt.Println(sum)
}
