package main

import "fmt"


func sum(nums []int) int {
  sum := 0
  for _, num := range nums {
    sum += num
  }

  return sum
}


func main() {
  array := []int{2, 42, -23, 3, 32,10}

  fmt.Println("Sum of", array, "is:", sum(array))
}
