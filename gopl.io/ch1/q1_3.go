package main

import (
	"fmt"
	"os"
	"strings"
  "time"
)

func main() {
	secs1 := measure(echo1)
	secs2 := measure(echo2)
  fmt.Println(fmt.Sprintf("echo1: %.7fs, echo2: %.7fs", secs1, secs2))
}

func measure(f func()) float64 {
  start := time.Now()
  f()
  return time.Since(start).Seconds()
}

func echo1() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func echo2() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
