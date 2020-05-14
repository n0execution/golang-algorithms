package main

import "fmt"

// Set implementation
type Set map[string]bool


func inititalizeSet(items []string) Set {
  var set = Set{}

  for _, key := range items {
    set[key] = true
  }
  return set
}


func main() {
    statesList := []string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"}
    statesNeeded := inititalizeSet(statesList)

    fmt.Println(statesNeeded)
}
