package main

import "fmt"

func Factorial(i int) int {
  if i <= 1 {
    return 1
  }
  return i * Factorial(i-1)
}

func TOHUtil(num int, from, to, temp string) {
  if num < 1 {
    return
  }
  TOHUtil(num - 1, from, temp, to)
  fmt.Println("Move disk ", num, " from peg ", from, " to peg ", to)
  TOHUtil(num - 1, temp, to, from)
}

func TowersOfHanoi(num int) {
  TOHUtil(num, "A", "C", "B")
}

func main() {
  TowersOfHanoi(3)
}
