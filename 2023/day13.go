package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func mirrorOffsets(lines []string) []int {
  var offsets []int
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
      offsets = append(offsets, i)
    }
  }
  return offsets
}

func valuesForGroup(group string) []int {
  lines := strings.Split(strings.TrimSpace(group), "\n")
  var offsets []int
  for _, result := range mirrorOffsets(lines) {
    if result > 0 {
      offsets = append(offsets, result * 100)
    }
  }

  columns := make([]string, len(lines[0]))
  for _, line := range lines {
    for j, char := range line {
      columns[j] += string(char)
    }
  }
  for _, result := range mirrorOffsets(columns) {
    if result > 0 {
      offsets = append(offsets, result)
    }
  }

  return offsets
}

func recalculate(group string, i int, char string) []int {
  sep := "."
  if char == "." {
    sep = "#"
  }
  s := strings.Join([]string{group[:i], group[i+1:]}, sep)
  return valuesForGroup(s)
}

func main() {
  data, _ := os.ReadFile("day13.txt")
  sum := 0
  for i, group := range strings.Split(string(data), "\n\n") {
    results := valuesForGroup(group)
    for j, char := range group {
      if strings.Contains(".#", string(char)) {
        for _, tmp := range recalculate(group, j, string(char)){
          if tmp != 0 && !slices.Contains(results, tmp) {
            results = append(results, tmp)
          }
        }
      }
    }
    sum += results[1]
    if len(results) == 0 {
      fmt.Println("No match for group", i)
    }
  }
  fmt.Println(sum)
}
