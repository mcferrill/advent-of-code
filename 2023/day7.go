package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var cardValues = map[rune]int{
  'A': 14,
  'K': 13,
  'Q': 12,
  'J': 11,
  'T': 10,
  '9': 9,
  '8': 8,
  '7': 7,
  '6': 6,
  '5': 5,
  '4': 4,
  '3': 3,
  '2': 2,
}

// Returns value of hand from 1-7. Higher number is better
func typeOfHand(hand string) int {
  cards := make(map[rune]int)
  var counts []int
  for _, card := range hand {
    cards[card] += 1
  }
  for _, v := range cards {
    counts = append(counts, v)
  }
  slices.SortFunc(counts, func(a, b int) int {
    return cmp.Compare(b, a)
  })

  if counts[0] == 5 {
    return 7
  } else if counts[0] == 4 {
    return 6
  } else if counts[0] == 3 && counts[1] == 2 {
    return 5
  } else if counts[0] == 3 {
    return 4
  } else if counts[0] == 2 && counts[1] == 2 {
    return 3
  } else if counts[0] == 2 {
    return 2
  }
  return 1
}

type Game struct {
  cards string
  bid int
  strength int
}

func main() {
  data, _ := os.ReadFile("day7.txt")
  lines := strings.Split(string(data), "\n")
  var games []Game
  for _, line := range lines {
    if line == "" {
      continue
    }
    parts := strings.Split(line, " ")
    bid, _ := strconv.Atoi(parts[1])
    games = append(games, Game{parts[0], bid, typeOfHand(parts[0])})
  }

  slices.SortFunc(games, func(a, b Game) int {
    if a.strength != b.strength {
      return cmp.Compare(b.strength, a.strength)
    } else {
      for i := 0; i <= 5; i++ {
        aCard := cardValues[rune(a.cards[i])]
        bCard := cardValues[rune(b.cards[i])]
        if aCard != bCard {
          return cmp.Compare(bCard, aCard)
        }
      }
    }
    return 0
  })

  sum := 0
  rank := 1
  for i := len(games)-1; i >= 0; i-- {
    sum += rank * games[i].bid
    rank++
  }
  fmt.Println(sum)
}
