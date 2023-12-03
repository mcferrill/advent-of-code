package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Find surrounding characters (inc. diagonally) and return the number if
// adjacent to a symbol, or 0 if not
func partNumber(row int, column int, end int, lines []string) bool {

  // Find surrounding characters
  if (column > 1) {
    column -= 1
  }
  if (end < len(lines[0]) - 2) {
    end += 1
  }
  chars := lines[row][column:end]
  if (row > 1) {
    chars += lines[row-1][column:end]
  }
  if (row < len(lines)-2) {
    chars += lines[row+1][column:end]
  }

  // Check for valid characters.
  r, _ := regexp.Compile(`[^\d.]`)
  return len(r.FindString(chars)) > 0
}

func main() {
  data, _ := os.ReadFile("day3.txt")
  lines := strings.Split(string(data), "\n")

  // Find numbers in line
  var sum int
  r, _ := regexp.Compile("(\\d+)")
  for row, line := range lines {
    for _, offset := range r.FindAllStringIndex(line, -1) {
      if (partNumber(row, offset[0], offset[1], lines)) {
        number, _ := strconv.Atoi(line[offset[0]:offset[1]])
        sum += number
      }
    }
  }

  fmt.Println(sum)
}
