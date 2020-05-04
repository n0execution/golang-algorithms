package main

import "fmt"


func simpleSearch(array []int, item int) int {
  for i, num := range array {
    if num == item {
      return i
    }
  }
  return -1
}


func main() {
  array := []int{2, 3, 5, 8, 10, 15, 22}

  fmt.Println(simpleSearch(array, 5))

}
