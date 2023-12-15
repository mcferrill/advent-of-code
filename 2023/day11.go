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

var grid [][]Point
var galaxies []Point

func main() {
  data, _ := os.ReadFile("tmp.txt")
  var lines []string
  for _, line := range strings.Split(string(data), "\n") {
    if line == "" {
      continue
    }
    tmp := strings.TrimSpace(strings.ReplaceAll(line, ".", " "))
    lines = append(lines, line)
    if len(tmp) == 0 {
      lines = append(lines, line)
    }
  }

  grid = make([][]Point, len(lines[0]))
  for y := 0; y <= len(lines)-1; y++ {
    line := strings.TrimSpace(lines[y])
    for x, char := range line {
      if len(grid[x]) == 0 {
        grid[x] = make([]Point, len(lines))
      }
      p := Point{x, y, string(char)}
      grid[x][y] = p
    }
  }

  for x := 0; x < len(grid); x++ {
    col := grid[x]
    emptyCol := true
    for _, point := range col {
      if point.raw == "#" {
        point.x = x
        emptyCol = false
        galaxies = append(galaxies, point)
      }
    }
    if emptyCol {
      grid = append(grid[:x+1], grid[x:]...)
      x++
    }
  }

  sum := 0
  for i, galaxy := range galaxies {
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
