package main

import (
	"fmt"
	"os"
	"slices"
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

var grid [][]Point
var traversed []Point

func traverse(from Point, offset Point) {
  x := from.x + offset.x
  y := from.y + offset.y
  if !inBounds(Point{x, y, ""}, grid) {
    return
  }
  next := grid[x][y]

  if !slices.Contains(traversed, next) {
    traversed = append(traversed, next)
  } else if slices.Index(traversed, next) == 0 || traversed[slices.Index(traversed, next) - 1] == from {
    return
  }

  if offset.x == 1 {
    // Moving right
    if next.raw == "/" {
      offset = Point{0, -1, ""}
    } else if next.raw == "\\" {
      offset = Point{0, 1, ""}
    } else if next.raw == "|" {
      traverse(next, Point{0, -1, ""})
      offset = Point{0, 1, ""}
    }
  } else if offset.x == -1 {
  // Moving left
    if next.raw == "/" {
      offset = Point{0, 1, ""}
    } else if next.raw == "\\" {
      offset = Point{0, -1, ""}
    } else if next.raw == "|" {
      traverse(next, Point{0, -1, ""})
      offset = Point{0, 1, ""}
    }
  } else if offset.y == -1 {
  // Moving up
    if next.raw == "/" {
      offset = Point{1, 0, ""}
    } else if next.raw == "\\" {
      offset = Point{-1, 0, ""}
    } else if next.raw == "-" {
      traverse(next, Point{-1, 0, ""})
      offset = Point{1, 0, ""}
    }
  } else if offset.y == 1 {
  // Moving down
    if next.raw == "/" {
      offset = Point{-1, 0, ""}
    } else if next.raw == "\\" {
      offset = Point{1, 0, ""}
    } else if next.raw == "-" {
      traverse(next, Point{-1, 0, ""})
      offset = Point{1, 0, ""}
    }
  }
  traverse(next, offset)
}

func main(){
  data, _ := os.ReadFile("day16.txt")
  grid = loadState(string(data))

  bestRoute := 0
  for x := range grid {
    traversed = []Point{}
    traverse(Point{x, 0, ""}, Point{0, 1, ""})
    bestRoute = max(len(traversed), bestRoute)
    traversed = []Point{}
    traverse(Point{x, len(grid[0]), ""}, Point{0, -1, ""})
    bestRoute = max(len(traversed), bestRoute)
  }

  for y := range grid[0] {
    traversed = []Point{}
    traverse(Point{0, y, ""}, Point{-1, 0, ""})
    bestRoute = max(len(traversed), bestRoute)
    traversed = []Point{}
    traverse(Point{0, len(grid), ""}, Point{1, 0, ""})
    bestRoute = max(len(traversed), bestRoute)
  }

  fmt.Println(bestRoute)
}
