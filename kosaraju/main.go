package main

import (
	"fmt"
  "os"
  "bufio"
  "strings"
  "time"
  ag "github.com/canoc62/kosaraju/adjacency_list"
  ag2 "github.com/canoc62/kosaraju/adjacency_list_2"
)

func main() {
  // ex commnad: ./main text/problem8.10test1.txt
  
  graph1 := ag.NewGraph() // uses custom linked-list implemented stack
  processAndRun1(graph1)
  graph2 := ag2.NewGraph() // uses go native slices for stack implementation
  processAndRun2(graph2)
}

func processAndRun1(graph *ag.Graph) {
 fmt.Println("Running kosaraju on graph1 starting at...", time.Now().Format("15:04:05\n"))
 f, err := os.Open(os.Args[1])
 defer f.Close()
 if err != nil {
   fmt.Println(err)
   os.Exit(1)
 }
 scanner := bufio.NewScanner(f)

 for scanner.Scan() {
   vertexIds := strings.Split(scanner.Text(), " ")
   if graph.VertexExists(vertexIds[0]) == false {
    graph.AddVertex(vertexIds[0])
   }

   if graph.VertexExists(vertexIds[1]) == false {
     graph.AddVertex(vertexIds[1])
   }

   graph.AddOutgoing(vertexIds[0], vertexIds[1])
   graph.AddIncoming(vertexIds[1], vertexIds[0])
 }

 graph.Kosaraju()
 result := graph.FiveLargestSccSizes()
 fmt.Println(result)
 message := fmt.Sprintf("kosaraju1 finished at: %s\n", time.Now().Format("15:04:05"))
 fmt.Println(message)
}

func processAndRun2(graph *ag2.Graph) {
  fmt.Println("Running kosaraju on graph2 starting at...", time.Now().Format("15:04:05\n"))
  f, err := os.Open(os.Args[1])
  defer f.Close()
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  scanner := bufio.NewScanner(f)

  for scanner.Scan() {
    vertexIds := strings.Split(scanner.Text(), " ")
    if graph.VertexExists(vertexIds[0]) == false {
     graph.AddVertex(vertexIds[0])
    }

    if graph.VertexExists(vertexIds[1]) == false {
      graph.AddVertex(vertexIds[1])
    }

    graph.AddOutgoing(vertexIds[0], vertexIds[1])
    graph.AddIncoming(vertexIds[1], vertexIds[0])
  }

  graph.Kosaraju()
  result := graph.FiveLargestSccSizes()
  fmt.Println(result)
  message := fmt.Sprintf("kosaraju2 finished at: %s\n", time.Now().Format("15:04:05"))
  fmt.Println(message)
}
