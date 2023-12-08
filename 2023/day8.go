package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
  data, _ := os.ReadFile("day8.txt")
  lines := strings.Split(string(data), "\n")
  var directions []int
  for _, step := range lines[0] {
    if step == 'R' {
      directions = append(directions, 1)
    } else {
      directions = append(directions, 0)
    }
  }

  var nodeList [][]string
  nodeMap := make(map[string]int)
  for i, line := range lines[2:] {
    if line == "" {
      continue
    }

    parts := strings.Split(line, " = ")
    name := parts[0]
    sides := strings.Split(parts[1], ", ")
    left := sides[0][1:]
    right := sides[1][:3]

    nodeMap[name] = i
    nodeList = append(nodeList, []string{name, left, right})
  }

  stepCount := 0
  currentNode := 0
  step := 0
  for currentNode != len(nodeList) - 1 {
    next := nodeList[currentNode][directions[step]+1]
    currentNode = nodeMap[string(next)]

    stepCount++
    step++
    if step >= len(directions) {
      step = 0
    }
  }

  fmt.Println(stepCount)
}
