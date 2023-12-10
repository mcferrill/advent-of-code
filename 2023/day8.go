package main

import (
	"fmt"
	"os"
	"strings"
)

var nodeMap map[string][]string
var directions []int

func part1() {
  currentNode := "AAA"
  step := 0
  for currentNode != "ZZZ" {
    leftOrRight := directions[step % len(directions)]
    currentNode = nodeMap[currentNode][leftOrRight]
    step++
  }

  fmt.Println(step)
}

// Adapted from https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
  for b != 0 {
    t := b
    b = a % b
    a = t
  }
  return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
  result := a * b / GCD(a, b)

  for i := 0; i < len(integers); i++ {
    result = LCM(result, integers[i])
  }

  return result
}

func part2() {
  // Make array of nodes ending with "A"
  var nodes []string
  for k := range nodeMap {
    if k[2] == 'A' {
      nodes = append(nodes, k)
    }
  }

  var steps []int
  for _, node := range nodes {
    step := 0
    for node[2] != 'Z' {
      leftOrRight := directions[step % len(directions)]
      node = nodeMap[node][leftOrRight]
      step++
    }
    steps = append(steps, step)
  }
  fmt.Println(LCM(steps[0], steps[1], steps[2:]...))
}

func main() {
  data, _ := os.ReadFile("day8.txt")
  lines := strings.Split(string(data), "\n")

  for _, step := range lines[0] {
    if step == 'R' {
      directions = append(directions, 1)
    } else {
      directions = append(directions, 0)
    }
  }

  nodeMap = make(map[string][]string)
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

  part2()
}
