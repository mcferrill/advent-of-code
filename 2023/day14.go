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

func saveState(grid [][]Point) string {
  lines := make([]string, len(grid[0]))
  for _, col := range grid {
    for y, p := range col {
      lines[y] = strings.Join([]string{lines[y], p.raw}, "")
    }
  }

  return strings.Join(lines, "\n")
}

func loadState(s string) [][]Point {
  lines := strings.Split(s, "\n")
  size := len(lines)
  size -= size % 2
  grid := make([][]Point, size)
  for y, line := range lines {
    for x, char := range line {
      if len(grid[x]) == 0 {
        grid[x] = make([]Point, size)
      }
      grid[x][y] = Point{x, y, string(char)}
    }
  }

  return grid
}

func inBounds(p Point, grid [][]Point) bool {
  if p.x < 0 || p.x > len(grid) - 1 {
    return false
  }
  if p.y < 0 || p.y > len(grid[0]) - 1 {
    return false
  }
  return true
}

func slide(p, offset Point, grid [][]Point) Point {
  if p.raw != "O" {
    return p
  }
  origin := Point{p.x, p.y, p.raw}
  for inBounds(p, grid) {
    x := p.x + offset.x
    y := p.y + offset.y
    if !inBounds(Point{x, y, ""}, grid) {
      break
    }
    if grid[x][y].raw == "." {
      p.x = x
      p.y = y
    } else {
      break
    }
  }

  grid[origin.x][origin.y].raw = "."
  grid[p.x][p.y] = p
  return p
}

func tilt(offset Point, grid [][]Point) {
  // Assumes a square grid
  for i := 0; i < len(grid); i++ {
    for j := 0; j < len(grid); j++ {
      x := i
      y := i
      if offset.x == 0 {
        y = j
        if offset.y == 1 {
          y = len(grid) - y - 1
        }
      } else {
        x = j
        if offset.x == 1 {
          x = len(grid) - x - 1
        }
      }

      slide(grid[x][y], offset, grid)
    }
  }
}

func count(grid [][]Point) int {
  sum := 0
  for x := range grid {
    for y := range grid[x] {
      p := grid[x][y]
      if p.raw == "O" {
        value := len(grid[0]) - p.y
        sum += value
      }
    }
  }
  return sum
}

func main() {
  data, _ := os.ReadFile("day14.txt")

  grid := loadState(string(data))

  offsets := []Point{
    {0, -1, ""},
    {-1, 0, ""},
    {0, 1, ""},
    {1, 0, ""},
  }

  cycles := 1000000000
  states := make(map[string]int)
  for i := 0; i < cycles; i++ {
    tilt(offsets[0], grid)
    tilt(offsets[1], grid)
    tilt(offsets[2], grid)
    tilt(offsets[3], grid)

    if states != nil {
      state := saveState(grid)
      match, ok := states[state]
      if ok {
        i = cycles - (cycles - i) % (i - match)
        states = nil
      } else {
        states[state] = i
      }
    } 
  }

  // Count
  fmt.Println(count(grid))
}
