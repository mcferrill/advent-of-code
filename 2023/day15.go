package main

import (
	"fmt"
	"os"
	"strings"
)

func hash(s string) int {
  value := 0
  for _, char := range s {
    value += int(char)
    value *= 17
    value = value % 256
  }
  return value
}

func main(){
  data, _ := os.ReadFile("day15.txt")

  sum := 0
  for _, step := range strings.Split(string(data), ",") {
    result := hash(strings.TrimSpace(step))
    sum += result 
  }

  fmt.Println(sum)
}
