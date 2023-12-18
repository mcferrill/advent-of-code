package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func count(line string, groups []int) int {
  result := 0

  // If we've satisfied all groups and there are no leftover springs.
  if len(groups) == 0 && !strings.Contains(line, "#") {
    result = 1
  }

  // For empty or unknown characters, also calculate the possibilities if we
  // skipped the first character.
  if line != "" && len(groups) > 0 && strings.Contains(".?", string(line[0])) {
    result += count(line[1:], groups)
  }

  if line != "" && len(groups) > 0 && strings.Contains("#?", string(line[0])) {
    // To proceed we need:
    // - A line at least as long as the current group
    // - A group match that doesn't include a "."
    // - A line the length of the group, or not a spring
    if groups[0] <= len(line) &&
      !strings.Contains(line[:groups[0]], ".") &&
      (len(line) == groups[0] || string(line[groups[0]]) != "#") {
      if len(line) > groups[0] {
        result += count(line[groups[0] + 1:], groups[1:])
      } else {
        result += count(line[groups[0]:], groups[1:])
      }
    }
  }

  return result
}

func main() {
  data, _ := os.ReadFile("day12.txt")
  lines := strings.Split(string(data), "\n")

  sum := 0
  for _, line := range lines {
    if line == "" {
      continue
    }
    parts := strings.Split(string(line), " ")
    raw := parts[0]
    var groups []int
    for _, length := range strings.Split(parts[1], ",") {
      value, _ := strconv.Atoi(length)
      groups = append(groups, value)
    }
    sum += count(raw, groups)
  }

  fmt.Println(sum)
}
