package main


import "fmt"


func findSmallest(array []int) int{
  smallest := array[0]
  smallestIndex := 0

  for i, num := range array {
    if num < smallest {
      smallest = num
      smallestIndex = i
    }
  }
  return smallestIndex
}


func selectionSort(array []int) []int {
  sortedArray := make([]int, 0)
  for range array {
      smallestIndex := findSmallest(array)
      smallest := array[smallestIndex]
      sortedArray = append(sortedArray, smallest)

      array = append(array[:smallestIndex], array[smallestIndex+1:]...)

  }

  return sortedArray
}


func main() {
  array := []int{2, 42, -234, 3, 42, 32, 10}
  fmt.Println(selectionSort(array))
}
