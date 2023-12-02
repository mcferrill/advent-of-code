package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var constraint = map[string]int{
  "red": 12,
  "green": 13,
  "blue": 14,
}

func main() {
  data, _ := os.ReadFile("day2.txt")
  lines := strings.Split(string(data), "\n")
  var validSum int
  for _, line := range lines {
    if (len(line) < 1) {
      continue
    }
    parts := strings.Split(line, ": ")
    gameID := strings.Split(parts[0], " ")[1]
    bags := strings.Split(parts[1], "; ")
    var bagList []map[string]int
    sums := map[string]int{
      "red": 0,
      "green": 0,
      "blue": 0,
    };
    for _, bag := range bags {
      bagColors := map[string]int{
        "red": 0,
        "green": 0,
        "blue": 0,
      };
      for _, color := range strings.Split(bag, ", ") {
        parts := strings.Split(color, " ")
        count, _ := strconv.Atoi(parts[0])
        bagColors[parts[1]] = count
        sums[parts[1]] += count
      }
      bagList = append(bagList, bagColors)
    }

    if (sums["red"] <= constraint["red"] && sums["green"] <= constraint["green"] && sums["blue"] <= constraint["blue"]) {
      id, _ := strconv.Atoi(gameID)
      validSum += id
    }
  }

  fmt.Println(validSum)
}
