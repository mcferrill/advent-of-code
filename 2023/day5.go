package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func find(source int, sourceMap []string) int {
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

    offset := source - sourceStart
    return destinationStart + offset
  }
  return source
}

func main() {
  data, _ := os.ReadFile("day5.txt")
  sections := strings.Split(string(data), "\n\n")
  seeds := strings.Split(sections[0], " ")[1:]

  var results []int
  for i := 0; i <= len(seeds); i += 2 {
    start, _ := strconv.Atoi(seeds[i])
    length, _ := strconv.Atoi(seeds[i+1])

    for i := start; i <= start + length; i++ {
      number := i
      for _, section := range sections {
        lines := strings.Split(section, "\n")[1:]
        number = find(number, lines)
      }
      results = append(results, number)
    }

    fmt.Println(results)
    return

  }

  fmt.Println(slices.Min(results))
}

