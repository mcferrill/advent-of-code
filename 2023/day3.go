package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Star struct {row, col int}

// Find surrounding characters (inc. diagonally) and return true if
// adjacent to a symbol, a list of corresponding * characters.
func partNumber(row int, column int, end int, lines []string) (bool, []Star) {

  // Find surrounding characters
  if (column > 1) {
    column -= 1
  }
  if (end < len(lines[0]) - 2) {
    end += 1
  }

  // Search portion of given lines in this pass
  var searchLines []string

  // Row indices for this pass
  var searchRows []int

  chars := lines[row][column:end]
  if (row > 1) {
    chars += lines[row-1][column:end]
    searchLines = append(searchLines, lines[row-1][column:end])
    searchRows = append(searchRows, row-1)
  }
  searchRows = append(searchRows, row)
  searchLines = append(searchLines, lines[row][column:end])
  if (row < len(lines)-2) {
    chars += lines[row+1][column:end]
    searchRows = append(searchRows, row+1)
    searchLines = append(searchLines, lines[row+1][column:end])
  }

  // Check for valid characters.
  r, _ := regexp.Compile(`[^\d.]`)
  match := r.FindString(chars)

  var stars []Star
  r, _ = regexp.Compile(`\*+`)
  if (len(match) > 0) {
    for i, line := range searchLines {
      for _, index := range r.FindAllStringIndex(line, -1) {
        colOffset := index[0] + column
        rowOffset := searchRows[i]
        stars = append(stars, Star{row: rowOffset, col: colOffset})
      }
    }
  }

  return len(match) > 0, stars
}

func main() {
  data, _ := os.ReadFile("day3.txt")
  lines := strings.Split(string(data), "\n")

  var sum int
  var starMap map[Star][]int = make(map[Star][]int)
  r, _ := regexp.Compile("(\\d+)")
  for row, line := range lines {
    for _, offset := range r.FindAllStringIndex(line, -1) {
      match, stars := partNumber(row, offset[0], offset[1], lines)
      if (match) {
        number, _ := strconv.Atoi(line[offset[0]:offset[1]])
        sum += number
        for _, star := range stars {
          starMap[star] = append(starMap[star], number)
        }
      }
    }
  }

  starSum := 0
  for _, numbers := range starMap {
    if (len(numbers) == 2) {
      starSum += (numbers[0] * numbers[1])
    }
  }

  fmt.Println(sum)
  fmt.Println(starSum)
}
