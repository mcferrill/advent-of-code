package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
  x int
  y int
  raw string
}

var columns [][]Point
var galaxies []Point
var multiplier = 1000000

func main() {
  data, _ := os.ReadFile("day11.txt")
  verticalOffset := 0
  for y, line := range strings.Split(string(data), "\n") {
    if line == "" {
      continue
    }
    tmp := strings.TrimSpace(strings.ReplaceAll(line, ".", " "))
    if len(tmp) == 0 {
      verticalOffset += multiplier-1
    } else {
      for x, char := range line {
        p := Point{x, y + verticalOffset, string(char)}
        if len(columns) - 1 < x {
          columns = append(columns, []Point{p})
        } else {
          columns[x] = append(columns[x], p)
        }
      }
    }
  }

  horizontalOffset := 0
  for _, column := range columns {
    empty := true
    for _, point := range column {
      if point.raw == "#" {
        point.x += horizontalOffset
        galaxies = append(galaxies, point)
        empty = false
      }
    }
    if empty {
      horizontalOffset += multiplier-1
    }
  }

  sum := 0
  for i, galaxy := range galaxies {
    fmt.Println(galaxy)
    for j, other := range galaxies {
      if j <= i {
        continue
      }
      x := max(galaxy.x, other.x) - min(galaxy.x, other.x)
      y := max(galaxy.y, other.y) - min(galaxy.y, other.y)
      sum += x + y
    }
  }

  fmt.Println(sum)
}
