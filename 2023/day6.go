package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func winConditions(time int, distance int) int {
  start := float64(distance) / float64(time)
  if (start != math.Trunc(start)){
    start = math.Trunc(start) + 1
  }

  // Validate start
  for i := int(start); i <= time-i; i++ {
    if (i * (time-i) > distance) {
      start = float64(i)
      break
    }
  }

  // Validate end
  var end int
  for i := time-1; i >= int(start); i-- {
    if (i * (time-i) > distance) {
      end = i
      break
    }
  }

  return end + 1 - int(start)
}

func main() {
  data, _ := os.ReadFile("day6.txt")
  r, _ := regexp.Compile(`\d+`)
  lines := strings.Split(string(data), "\n")
  time := strings.Join(r.FindAllString(lines[0], -1), "")
  distance := strings.Join(r.FindAllString(lines[1], -1), "")
  t, _ := strconv.Atoi(time)
  d, _ := strconv.Atoi(distance)
  fmt.Println(winConditions(t, d))
}

