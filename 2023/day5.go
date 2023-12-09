package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	// "sync"
)

// Traverse a single number through a single map.
func find(source int, sourceMap [][]int) int {
  for _, sourceRange := range sourceMap {
    dStart := sourceRange[0]
    // dEnd := sourceRange[1]
    sStart := sourceRange[2]
    sEnd := sourceRange[3]

    if (source > sEnd || source < sStart) {
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

// Find (start, end) intersection with (destination start, destination end)
// and return array of corresponding (source start, source end)
func findFromRange(start, end int, sourceMap [][]int) [][]int {
  var result [][]int
  for _, sourceRange := range sourceMap {
    if (start > sourceRange[1] || end < sourceRange[0]) {
      continue
    }

    startOffset := max(start, sourceRange[0]) - sourceRange[0]
    endOffset := sourceRange[1] - min(end, sourceRange[1])

    start = sourceRange[2] + startOffset
    end = sourceRange[3] + endOffset

    result = append(result, []int{start, end})
  }
  return result
}

func validSeed(number int) bool {
  for _, seedRange := range seeds {
    if (number >= seedRange[0] && number <= seedRange[1]) {
      return true
    }
  }
  return false
}

var maps [][][]int
var seeds [][]int

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
    maps = append(maps, sectionMap)
  }

  // Parse seeds as rows of [rangeStart, rangeEnd]
  numbers := strings.Split(sections[0], " ")[1:]
  for i := 0; i <= len(numbers) - 1; i += 2 {
    start, _ := strconv.Atoi(numbers[i])
    length, _ := strconv.Atoi(numbers[i + 1])
    seeds = append(seeds, []int{start, start + length})
  }

  // recreating part 1
  // var results []int
  // for _, number := range strings.Split(sections[0], " ")[1:] {
  //   seed, _ := strconv.Atoi(number)
  //   results = append(results, seedToLocation(seed))
  // }
  // fmt.Println(slices.Min(results))
  // return

  // brute force
  lowest := -1
  for _, seedRange := range seeds {
    for i := seedRange[0]; i <= seedRange[1]; i++ {
      number := seedToLocation(i)
      if lowest < 0 || number < lowest {
        lowest = number
      }
    }
  }
  fmt.Println(lowest)
  return

  // Sort locations by lowest first
  locations := maps[len(maps)-1]
  slices.SortFunc(locations, func(a, b []int) int {
    return cmp.Compare(a[0], b[0])
  })

  // var results []int
  lowest = -1
  for _, locationRange := range locations {
    start := locationRange[0]
    end := locationRange[1]

    var nextRanges [][]int = [][]int{{start, end}}

    for i := len(maps)-2; i > -1; i-- {
      ranges := nextRanges[:]
      nextRanges = [][]int{{}}

      for _, r := range ranges {
        if len(r) == 0 {
          continue
        }
        result := findFromRange(r[0], r[1], maps[i])
        nextRanges = append(nextRanges, result...)
      }
    }

    for _, r := range nextRanges {
      if len(r) == 0 {
        continue
      }
      for i := r[0]; i < r[1]; i++ {
        if validSeed(i) {
          result := seedToLocation(i)
          if lowest > -1 && result < lowest {
            lowest = result
          }
        }
      }
    }
  }

  fmt.Println(lowest)
}

