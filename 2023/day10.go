package main

import (
	"fmt"
	"os"
	"reflect"
	"slices"
	"strings"
)

type Point struct {
  x int
  y int
  raw string
}

var grid [][]Point

// Values corresponding to opening in a given direction
var directions map[string][]int = map[string][]int{
  "|LJ": {0, -1}, // Up
  "|7F": {0, 1}, // Down
  "-J7": {-1, 0}, // Left
  "-LF": {1, 0}, // Right
}

func offsets(from Point) [][]int {
  var results [][]int
  for match, offset := range directions {
    other := []int{from.x + offset[0], from.y + offset[1]}
    if other[0] < 0 || other[0] > len(grid) || other[1] < 0 || other[1] > len(grid[0]) {
      continue
    }
    if from.raw == "S" || strings.Contains(match, from.raw) {
      results = append(results, offset)
    }
  }
  return results
}

func findNext(start Point) []Point {
  var result []Point
  for _, offset := range offsets(start) {
    other := grid[start.x + offset[0]][start.y + offset[1]]
    for _, otherOffset := range offsets(other) {
      if slices.Equal(otherOffset, []int{-offset[0], -offset[1]}) {
        result = append(result, other)
      }
    }
  }
  return result
}

func area(perimeter []Point) int {
  value := 0.0

  last := perimeter[len(perimeter) - 1]

  for _, point := range perimeter {
    value += float64(last.x + point.x) * float64(last.y - point.y)
    last = point
  }

  return int(value / 2)
}

func main() {
  data, _ := os.ReadFile("day10.txt")
  lines := strings.Split(string(data), "\n")

  var start Point
  var ground []Point

  grid = make([][]Point, len(lines[0]))
  for y, line := range lines {
    for x, char := range line {
      if len(grid[x]) == 0 {
        grid[x] = make([]Point, len(lines))
      }
      p := Point{x, y, string(char)}
      if char == 'S' {
        start = p 
      } else if char == '.' {
        ground = append(ground, p)
      }
      grid[x][y] = p
    }
  }

  var steps []Point
  pos := start
  last := start
  for len(steps) == 0 || pos.raw != "S" {
    if len(steps) > 0 && reflect.DeepEqual(pos, last) {
      break
    }
    steps = append(steps, pos)
    for _, neighbor := range findNext(pos) {
      if !reflect.DeepEqual(neighbor, last) {
        last = pos
        pos = neighbor
        break
      }
    }
  }
  fmt.Println(len(steps) / 2)

  loopArea := area(steps)
  fmt.Println(loopArea - len(steps) / 2 + 1)
}
