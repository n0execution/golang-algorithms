package main

import "fmt"


func sum(nums []int) int {
  sum := 0
  for _, num := range nums {
    sum += num
  }

  return sum
}


func recursiveSum(nums []int) int {
  if len(nums) == 0 {
    return 0
  }
  return nums[0] + recursiveSum(nums[1:])
}


func main() {
  array := []int{2, 42, -23, 3, 32,10}

  fmt.Println("Sum of", array, "is:", sum(array))
  fmt.Println("Recursive sum of", array, "is:", recursiveSum(array))
}
