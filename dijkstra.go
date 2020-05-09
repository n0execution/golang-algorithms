package main


import ("fmt";
        "math")


func find_lowest_cost_node(costs map[string]float64, processed []string) string{
  lowest_cost:= math.Inf(1)
  var lowest_cost_node string
  for node, cost := range costs{
    if cost < lowest_cost {
      lowest_cost = cost
      lowest_cost_node = node
    }
  }
  return lowest_cost_node
}


func main() {
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
  costs["fin"] = math.Inf(1)

  var parents = map[string]string{}
  parents["a"] = "start"
  parents["b"] = "start"

  processed := make([]string, 0)

  fmt.Println(find_lowest_cost_node(costs, processed))
}
