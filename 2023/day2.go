package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Takes a line from the input file and returns the game id and power of
// "Minimum Viable Cubes" if valid or 0 if not.
func validGame(line string) (int, int) {
  parts := strings.Split(line, ": ")
  id, _ := strconv.Atoi(strings.Split(parts[0], " ")[1])

  peeks := strings.Split(parts[1], "; ")
  // minimum viable cubes
  MVC := map[string]int{"red": 0, "green": 0, "blue": 0};
  for _, peek := range peeks {
    cubes := map[string]int{"red": 0, "green": 0, "blue": 0};
    for _, color := range strings.Split(peek, ", ") {
      parts := strings.Split(color, " ")
      count, _ := strconv.Atoi(parts[0])
      cubes[parts[1]] = count
      MVC[parts[1]] = max(count, MVC[parts[1]])
    }
    if (cubes["red"] > 12 || cubes["green"] > 13 || cubes["blue"] > 14) {
      id = 0
    }

  }

  return id, (MVC["red"] * MVC["green"] * MVC["blue"])
}

func main() {
  data, _ := os.ReadFile("day2.txt")
  lines := strings.Split(string(data), "\n")
  var validSum, MVCSum int
  for _, line := range lines {
    if (len(line) < 1) {
      continue
    }
    sum, mvc := validGame(line)
    validSum += sum
    MVCSum += mvc
  }

  fmt.Println(validSum)
  fmt.Println(MVCSum)
}
