package dijkstra

import (
  "math"
  "bytes"
  "fmt"
)

type vertex struct {
  score int // key
  id string
  list *linkedList
  prev *vertex
}

type graph struct {
  vertices *map[string]*vertex
}

func NewGraph() *graph {
  newMap := make(map[string]*vertex)
  return &graph{vertices: &newMap}
}

func (g *graph) AddVertex(id string) *vertex {
  graphVertices := *(*g).vertices

  _, ok := graphVertices[id]
  if !ok {
    graphVertices[id] = &vertex{score: math.MaxInt32, id: id, list: newLinkedList()}
  }
  return graphVertices[id]
}

func (g *graph) AddEdge(tailId string, headId string, weight int) {
  graphVertices := *(*g).vertices

  tail, tailOk := graphVertices[tailId]
  if !tailOk {
    tail = g.AddVertex(tailId)
  }
  head, headOk := graphVertices[headId]
  if !headOk {
    head = g.AddVertex(headId)
  }
  (*(*tail).list).put(head, weight)
}

func (g *graph) putVerticesIntoUnderlyingHeapArray(mappings *map[string]int, initId string) *heap {
  h := heap{}
  indexMap := *mappings

  for _, v := range *(*g).vertices {
    h = append(h, v)
    indexMap[v.id] = len(h)- 1
  }

  currInitIdIdx := indexMap[initId]
  h[currInitIdIdx], h[0] = h[0], h[currInitIdIdx]
  indexMap[h[currInitIdIdx].id] = currInitIdIdx
  indexMap[h[0].id] = 0

  return &h
}

func (g *graph) Dijkstra(vertexId string, targets []string) map[string]int {
  seen :=  map[string]struct{}{}
  mappings := map[string]int{} // mapping from vertex to position in heap, to allow for deletion
  (*(*g).vertices)[vertexId].score = 0
  heapArray:= g.putVerticesIntoUnderlyingHeapArray(&mappings, vertexId)
  dijkstraHeap := *(Heapify(heapArray, &mappings, 0))

  for len(dijkstraHeap) > 0 {
    currVertex, ok := dijkstraHeap.ExtractMin(&mappings)

    if !ok {
      panic("Error extracting min from heap.")
      break
    }
    seen[currVertex.id] = struct{}{}
    for node := currVertex.list.head; node != nil; node = node.next {
      vertex := node.data
      _, ok := seen[vertex.id]

      if !ok {
        dijkstraScore := currVertex.score + node.weight
      
        if dijkstraScore < vertex.score {
          _, deleteOk := dijkstraHeap.Delete(&mappings, mappings[vertex.id])

          if !deleteOk {
            panic("Error deleting from heap.")
          }
          vertex.score = dijkstraScore
          vertex.prev = currVertex
          dijkstraHeap.Insert(&mappings, vertex)
        }
      }
    }
  }

  results := map[string]int{}
  for _, vId := range targets {
    results[vId] = (*(*g).vertices)[vId].score
  }
  return results
}

func (g *graph) PrintTrace(vertexId string) string {
  start := (*(*g).vertices)[vertexId]
  var trace bytes.Buffer
  
  for ptr := start; ptr != nil; ptr = ptr.prev {
    trace.WriteString(fmt.Sprintf("%s: %d, ", ptr.id, ptr.score))
  }

  return trace.String()
}