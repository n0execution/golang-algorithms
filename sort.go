package main


import ("fmt"
        "math/rand")


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


func quickSort(array []int) []int{
  if len(array) >= 2 {
    sortedArray := make([]int, 0)
    index := rand.Intn(len(array))
    pivot := array[index]
    leftArray := make([]int, 0)
    rightArray := make([]int, 0)

    array = append(array[:index], array[index+1:]...)
    for _, num := range array {
      if num <= pivot {
        leftArray = append(leftArray, num)
      } else if num > pivot{
        rightArray = append(rightArray, num)
      }
    }
    sortedArray = append(quickSort(leftArray), pivot)
    sortedArray = append(sortedArray, quickSort(rightArray)...)

    return sortedArray
  }

  return array
}


func main() {
  array := []int{2, 42, -234, 3, 42, 32, 10}
  fmt.Println(quickSort(array))
}
