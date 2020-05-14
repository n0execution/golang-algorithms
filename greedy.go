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


func itemInSet(item string, set Set) bool {
  for key := range set {
    if key == item {
      return true
    }
  }
  return false
}


func intersection(set1 Set, set2 Set) Set{
    var intersectionSet = Set{}

    for item := range set1 {
      if itemInSet(item, set2) {
        intersectionSet[item] = true
      }
    }

    return intersectionSet
}


func main() {
    statesList := []string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"}
    statesNeeded := inititalizeSet(statesList)

    fmt.Println(statesNeeded)

    var stations = map[string]Set{}
    stations["kone"] = inititalizeSet([]string{"id", "nv", "ut"})
    stations["ktwo"] = inititalizeSet([]string{"wa", "id", "mt"})
    stations["kthree"] = inititalizeSet([]string{"or", "nv", "ca"})
    stations["kfour"] = inititalizeSet([]string{"nv", "ut"})
    stations["kfive"] = inititalizeSet([]string{"ca", "az"})

    fmt.Println(stations)

    var finalStations = Set{}
    fmt.Println(finalStations)

    fmt.Println(intersection(stations["kone"], stations["kfour"]))
}
