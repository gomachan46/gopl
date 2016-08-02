package main

import (
  "math/rand"
  "fmt"
  "time"
)

func main() {
  switch coinflip() {
  case "heads":
    fmt.Println("coinflip heads")
  case "tails":
    fmt.Println("coinflip tails")
  default:
    fmt.Println("landed on edge!")
  }
  fmt.Println(janken())
}

func coinflip() string {
  rand.Seed(time.Now().UnixNano())
  n := rand.Intn(2)
  if n > 0 {
    return "heads"
  } else {
    return "tails"
  }
}

func janken() string {
  rand.Seed(time.Now().UnixNano())
  n := rand.Intn(3)
  switch {
  case n == 0:
    return "グー"
  case n == 1:
    return "チョキ"
  case n == 2:
    return "パー"
  default:
    return "反則奴〜"
  }
}
