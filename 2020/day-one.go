package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)


func find_matching_rows(rows []int) [3]int {
  for _, row1 := range rows {
    for _, row2 := range rows {
      if (row1 == row2) {
        continue
      }
      for _, row3 := range rows {
        if (row3 == row1 || row3 == row2) {
          continue
        }
        if (row1 + row2 + row3 == 2020) {
          result := [3]int{row1, row2, row3}
          return result
        }
      }
    }
  }
  panic("no match found!")
}

func main() {
  data, _ := os.ReadFile("day1.txt")
  lines := strings.Split(string(data), "\n")
  var rows []int
  for _, line := range lines {
    n, _ := strconv.Atoi(line)
    rows = append(rows, n)
  }
  result := find_matching_rows(rows)

  fmt.Println(result[0] * result[1] * result[2])
}
