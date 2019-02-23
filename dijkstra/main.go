package main

import (
  "os"
  "fmt"
  "bufio"
  "strings"
  "strconv"
  graph "github.com/canoc62/dijkstra/graph"
)

func main() {
  // ex command: ./main text/problem9.8.txt 1 7,37,59,82,99,115,133,165,188,197
  
  f, err := os.Open(os.Args[1])
  startingVertex := os.Args[2]
  targets := strings.Split(os.Args[3], ",")
  defer f.Close()
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  scanner := bufio.NewScanner(f)

  g := graph.NewGraph()

  for scanner.Scan() {
    vals := strings.Fields(scanner.Text())
    tail := vals[0]
    g.AddVertex(tail)
    edges := vals[1:]
    for _, edge := range edges {
      nodes := strings.Split(string(edge), ",")
      head := nodes[0]
      weight, strToIErr := strconv.Atoi(nodes[1])
      if strToIErr == nil {
        g.AddEdge(tail, head, weight)
      } else {
        panic(err)
      }
    }
  }

  edgeLengths := g.Dijkstra(startingVertex, targets)
  result := []int{}
  for _, node := range targets {
    result = append(result, edgeLengths[node])
  }
  fmt.Println("result: ", result)
  fmt.Printf("Trace 197: \n%s\n", g.PrintTrace("197"))
}
