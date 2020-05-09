package main


import ("fmt";
        "math")


func stringInSlice(processed []string, node string) bool {
  for _, n := range processed {
    if n == node {
      return true
    }
  }
  return false
}


func findLowestCostNode(costs map[string]float64, processed []string) string{
  lowest_cost := math.Inf(1)
  var lowest_cost_node string
  for node, cost := range costs{
    if cost < lowest_cost && !stringInSlice(processed, node) {
      lowest_cost = cost
      lowest_cost_node = node
    }
  }
  return lowest_cost_node
}


func search(graph map[string]map[string]float64, costs map[string]float64,
            parents map[string]string, processed []string) (map[string]string, map[string]float64) {
    node := findLowestCostNode(costs, processed)

    for node != "" {
      cost := costs[node]
      neighbours := graph[node]

      for n, c := range neighbours {
        new_cost := cost + c
        if costs[n] > new_cost {
          costs[n] = new_cost
          parents[n] = node
        }
      }
      processed = append(processed, node)
      node = findLowestCostNode(costs, processed)
    }

    return parents, costs
}


func main() {
  var graph = map[string]map[string]float64{}
  graph["start"] = map[string]float64{}
  graph["start"]["a"] = 6
  graph["start"]["b"] = 2

  graph["a"] = map[string]float64{}
  graph["a"]["fin"] = 1

  graph["b"] = map[string]float64{}
  graph["b"]["a"] = 3
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

  parents, costs = search(graph, costs, parents, processed)
  fmt.Println("Minimal path to fin is:", costs["fin"])
}
