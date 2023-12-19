package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func mirrorOffset(lines []string) int {
  for i := range lines {
    length := min(i, len(lines) - i)
    if length < 1 {
      continue
    }
    left := lines[i-length:i]
    right := make([]string, length)
    copy(right, lines[i:i+length])
    slices.Reverse(right)
    if strings.Join(left, "") == strings.Join(right, "") {
      return i
    }
  }
  return 0
}

func valueForGroup(group string) int {
  lines := strings.Split(strings.TrimSpace(group), "\n")
  result := mirrorOffset(lines)
  if result > 0 {
    return result * 100
  }

  columns := make([]string, len(lines[0]))
  for _, line := range lines {
    for j, char := range line {
      columns[j] += string(char)
    }
  }
  return mirrorOffset(columns)
}

func main() {
  data, _ := os.ReadFile("day13.txt")
  sum := 0
  for i, group := range strings.Split(string(data), "\n\n") {
    result := valueForGroup(group)
    sum += result
    if result == 0 {
      fmt.Println("No match for group", i)
    }
  }
  fmt.Println(sum)
}
