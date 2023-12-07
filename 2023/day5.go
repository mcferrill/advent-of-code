package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func find(start int, end int, sourceMap [][]int) int {
  for _, row := range sourceMap {
    dStart := row[0]
    // dEnd := row[1]
    sStart := row[2]
    sEnd := row[3]

    if (start > sEnd || end < sStart) {
      continue
    }

    offset := sStart - max(start, sStart)
    return dStart + offset

  }
  return 0
}

func main() {
  data, _ := os.ReadFile("day5.txt")
  sections := strings.Split(string(data), "\n\n")
  var maps [][][]int
  for _, section := range sections[1:] {
    var sectionMap [][]int
    for _, line := range strings.Split(section, "\n")[1:] {
      if (line == "") {
        continue
      }
      parts := strings.Split(string(line), " ")
      destinationStart, _ := strconv.Atoi(parts[0])
      sourceStart, _ := strconv.Atoi(parts[1])
      rangeLength, _ := strconv.Atoi(parts[2])
      row := []int{destinationStart, destinationStart + rangeLength, sourceStart, sourceStart + rangeLength}
      sectionMap = append(sectionMap, row)
    }
    maps = append(maps, sectionMap)
  }

  seeds := strings.Split(sections[0], " ")[1:]

  var results []int
  for i := 0; i <= len(seeds) - 1; i += 2 {
    start, _ := strconv.Atoi(seeds[i])
    length, _ := strconv.Atoi(seeds[i+1])

    number := find(start, start + length, maps[0])
    for _, section := range maps[1:] {
      number = find(number, number, section)
      if (number == 0) {
        break
      }
    }
    results = append(results, number)

  }

  fmt.Println(slices.Min(results))
}

