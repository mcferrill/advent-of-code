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

func elevate(p Point) Point {
  if p.y == 0 {
    return p
  }
  for y := p.y-1; y >= 0; y-- {
    if grid[p.x][y].raw == "." {
      p.y = y
    } else {
      break
    }
  }

  return p
}

func printGrid() {
  lines := make([]string, len(grid[0]))
  for _, col := range grid {
    for y, p := range col {
      lines[y] = strings.Join([]string{lines[y], p.raw}, "")
    }
  }

  for _, line := range lines {
    fmt.Println(line)
  }
}

func main() {
  data, _ := os.ReadFile("day14.txt")
  lines := strings.Split(string(data), "\n")

  grid = make([][]Point, len(lines[0]))
  for y, line := range lines {
    for x, char := range line {
      if len(grid[x]) == 0 {
        grid[x] = make([]Point, len(lines))
      }
      grid[x][y] = Point{x, y, string(char)}
    }
  }

  sum := 0
  for x := range grid {
    for y := range grid[x] {
      p := grid[x][y]
      if p.raw == "O" {
        p = elevate(p)
        grid[x][y].raw = "."
        grid[x][p.y] = p
        value := len(grid[0]) - 1 - p.y
        sum += value
      }
    }
  }
  fmt.Println(sum)
}
