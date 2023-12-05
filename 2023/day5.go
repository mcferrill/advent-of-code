package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func find(source int, sourceMap []string) int {
  var matches []int
  for _, line := range sourceMap {
    if (line == "") {
      continue
    }
    parts := strings.Split(string(line), " ")
    destinationStart, _ := strconv.Atoi(parts[0])
    sourceStart, _ := strconv.Atoi(parts[1])
    rangeLength, _ := strconv.Atoi(parts[2])

    if (sourceStart > source || source > sourceStart + rangeLength) {
      continue
    }

    for i := 0; i <= 0 + rangeLength; i++ {
      if (sourceStart + i == source) {
        matches = append(matches, destinationStart + i)
      }
    }
  }
  if (len(matches) > 0) {
    return slices.Min(matches)
  }
  return source
}

func main() {
  data, _ := os.ReadFile("day5.txt")
  sections := strings.Split(string(data), "\n\n")
  seeds := strings.Split(sections[0], " ")[1:]

  var results []int
  for _, seed := range seeds {
    number, _ := strconv.Atoi(seed)
    for _, section := range sections {
      lines := strings.Split(section, "\n")[1:]
      number = find(number, lines)
    }
    results = append(results, number)
  }

  fmt.Println(slices.Min(results))
}

