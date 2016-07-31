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
}

func coinflip() string {
  rand.Seed(time.Now().UnixNano())
  n := rand.Intn(2)
  fmt.Println(n)
  if n > 0 {
    return "heads"
  } else {
    return "tails"
  }
}
