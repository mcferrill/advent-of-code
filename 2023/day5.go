package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var maps [][][]int
var seeds [][]int

// Traverse a single number through a single map.
func find(source int, sourceMap [][]int) int {
  for _, sourceRange := range sourceMap {
    dStart := sourceRange[0]
    // dEnd := sourceRange[1]
    sStart := sourceRange[2]
    sEnd := sourceRange[3]

    if source < sStart || source >= sEnd {
      continue
    }

    offset := source - sStart
    return dStart + offset
  }
  return source
}

func seedToLocation(seed int) int {
  number := seed
  for _, section := range maps {
    number = find(number, section)
  }
  return number
}

func main() {
  data, _ := os.ReadFile("day5.txt")
  sections := strings.Split(string(data), "\n\n")

  // Convert maps to rows of [destinationStart, destinationEnd, sourceStart, sourceEnd]
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
    slices.SortFunc(sectionMap, func(a, b []int) int {
      return cmp.Compare(a[2], b[2])
    })
    maps = append(maps, sectionMap)
  }

  // Parse seeds as rows of [rangeStart, rangeEnd]
  numbers := strings.Split(sections[0], " ")[1:]
  for i := 0; i <= len(numbers) - 1; i += 2 {
    start, _ := strconv.Atoi(numbers[i])
    length, _ := strconv.Atoi(numbers[i + 1])
    seeds = append(seeds, []int{start, start + length})
  }

  // recreating part 1 = 226172555
  // lowest : = -1
  // for _, number := range strings.Split(sections[0], " ")[1:] {
  //   seed, _ := strconv.Atoi(number)
  //   results = append(results, seedToLocation(seed))
  // }
  // fmt.Println(slices.Min(results))
  // return

  // brute force part 2
  lowest := -1
  lowestSeed := -1
  for _, seedRange := range seeds {
    for i := seedRange[0]; i <= seedRange[1]; i++ {
      number := seedToLocation(i)
      if lowest < 0 || number < lowest {
        lowest = number
        lowestSeed = i
      }
    }
  }
  fmt.Println(lowest, lowestSeed)
}

