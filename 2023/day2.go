package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Takes a line from the input file and returns the game id if valid or 0 if not.
func validGame(line string) int {
  parts := strings.Split(line, ": ")
  id, _ := strconv.Atoi(strings.Split(parts[0], " ")[1])

  peeks := strings.Split(parts[1], "; ")
  for _, bag := range peeks {
    cubes := map[string]int{"red": 0, "green": 0, "blue": 0};
    for _, color := range strings.Split(bag, ", ") {
      parts := strings.Split(color, " ")
      count, _ := strconv.Atoi(parts[0])
      cubes[parts[1]] = count
    }
    if (cubes["red"] > 12 || cubes["green"] > 13 || cubes["blue"] > 14) {
      return 0
    }
  }

  return id
}

func main() {
  data, _ := os.ReadFile("day2.txt")
  lines := strings.Split(string(data), "\n")
  var validSum int
  for _, line := range lines {
    if (len(line) < 1) {
      continue
    }
    validSum += validGame(line)

  }

  fmt.Println(validSum)
}
