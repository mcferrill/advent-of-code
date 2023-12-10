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

  nodeMap := make(map[string][]string)
  for _, line := range lines[2:] {
    if line == "" {
      continue
    }

    parts := strings.Split(line, " = ")
    name := parts[0]
    sides := strings.Split(parts[1], ", ")
    left := sides[0][1:]
    right := sides[1][:3]

    nodeMap[name] = []string{left, right}
  }

  currentNode := "AAA"
  step := 0
  for currentNode != "ZZZ" {
    leftOrRight := directions[step % len(directions)]
    currentNode = nodeMap[currentNode][leftOrRight]
    step++
  }

  fmt.Println(step)
}
