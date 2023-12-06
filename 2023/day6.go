package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
  data, _ := os.ReadFile("day6.txt")
  r, _ := regexp.Compile(`\d+`)
  lines := strings.Split(string(data), "\n")
  time := r.FindAllString(lines[0], -1)
  distance := r.FindAllString(lines[1], -1)
  total := 0
  for i, race := range time {
    t, _ := strconv.Atoi(race)
    d, _ := strconv.Atoi(distance[i])
    start := float64(d) / float64(t)
    if (start != math.Trunc(start)){
      start = math.Trunc(start) + 1
    }

    var winCondition []int
    for i := int(start); i <= t-1; i++ {
      if (i * (t-i) > d) {
        winCondition = append(winCondition, i)
      }
    }
    if (total == 0) {
      total = len(winCondition)
    } else {
      total *= len(winCondition)
    }
  }
  fmt.Println(total)
}

