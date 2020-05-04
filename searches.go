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


func binarySearch(array []int, item int) int {
  lowIndex := 0


  highIndex := len(array) - 1

  for lowIndex <= highIndex {
    midIndex := (lowIndex + highIndex) / 2

    if array[midIndex] == item {
      return midIndex
    } else if item < array[midIndex] {
      highIndex = midIndex - 1
    } else if item > array[midIndex] {
      lowIndex = midIndex + 1
    }

  }
  return -1
}


func main() {
  array := []int{2, 3, 5, 8, 10, 15, 22}

  fmt.Println(binarySearch(array, 5))

}
