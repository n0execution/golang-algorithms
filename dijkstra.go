package main


import "fmt"


func main() {
  var graph = map[string]map[string]int{}
  graph["start"] = map[string]int{}
  graph["start"]["a"] = 6
  graph["start"]["b"] = 2

  graph["a"] = map[string]int{}
  graph["a"]["fin"] = 1

  graph["b"] = map[string]int{}
  graph["b"]["fin"] = 5

  graph["fin"] = map[string]int{}

  fmt.Println(graph)
}
