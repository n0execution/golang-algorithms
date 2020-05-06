package main


import "fmt"


//Queue implementation
type Queue []string


func (q *Queue) popLeft() string{
  firstPerson := (*q)[0]
  *q = (*q)[1:]
  return firstPerson
}


func (q *Queue) addToQueue(items []string) {
  *q = append(*q, items...)
}


func isPersonSeller(name string) bool {
  return name[len(name)-1] == 'm'
}


func search(q *Queue, graph map[string][]string) bool {
  for len(*q) != 0 {
    person := (*q).popLeft()
    if isPersonSeller(person) {
      fmt.Println(person, "is a mango seller")
      return true
    } else {
      (*q).addToQueue(graph[person])
    }
  }

  return false
}


func main() {
  var graph = make(map[string][]string)
  graph["you"] = []string{"alice", "bob", "claire"}
  graph["alice"] = []string{"peggy"}
  graph["bob"] = []string{"anuj", "peggy"}
  graph["claire"] = []string{"thom", "jonny"}

  var queue Queue
  queue = graph["you"]

  search(&queue, graph)
}
