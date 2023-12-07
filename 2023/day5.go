package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func find(source int, sourceMap [][]int) int {
  for _, row := range sourceMap {
    dStart := row[2]
    // dEnd := row[3]
    sStart := row[0]
    sEnd := row[1]

    if (source > sEnd || source < sStart) {
      continue
    }

    return dStart + source - sStart
  }
  return source
}

func validSeed(number int, seeds [][]int) bool {
  for _, seedRange := range seeds {
    if (number >= seedRange[0] && number <= seedRange[1]) {
      return true
    }
  }
  return false
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

  var seeds [][]int
  numbers := strings.Split(sections[0], " ")[1:]
  for i := 0; i <= len(numbers) - 1; i += 2 {
    start, _ := strconv.Atoi(numbers[i])
    length, _ := strconv.Atoi(numbers[i + 1])
    seeds = append(seeds, []int{start, start + length})
  }

  locations := maps[len(maps)-1]
  slices.SortFunc(locations, func(a, b []int) int {
    return cmp.Compare(a[0], b[0])
  })

  for _, locationRange := range locations {
    for number := locationRange[0]; number <= locationRange[1]; number++ {
      traversalNumber := number
      for i := len(maps) -1; i > -1; i-- {
        traversalNumber = find(traversalNumber, maps[i])
      }
      valid := validSeed(number ,seeds)
      if (valid) {
        fmt.Println(number)
        return
      }
    }
  }
}

