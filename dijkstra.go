package main


import ("fmt";
        "math")


func main() {
  infinity := math.Inf(1)

  var graph = map[string]map[string]float64{}
  graph["start"] = map[string]float64{}
  graph["start"]["a"] = 6
  graph["start"]["b"] = 2

  graph["a"] = map[string]float64{}
  graph["a"]["fin"] = 1

  graph["b"] = map[string]float64{}
  graph["b"]["fin"] = 5

  graph["fin"] = map[string]float64{}

  var costs = map[string]float64{}
  costs["a"] = 6
  costs["b"] = 2
  costs["fin"] = infinity

  fmt.Println(costs)
}
