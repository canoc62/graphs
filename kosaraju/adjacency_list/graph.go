package adjacency_list

import (
  "fmt"
  "sort"
)

type Graph struct {
  v *map[string]*vertex
  sccSizes *map[int]int
  numVertices int
}

type vertex struct {
  val string
  position int
  explored bool
  scc int
	outgoing *LinkedList
  incoming *LinkedList
}

func NewGraph() *Graph {
  sizesMap := make(map[int]int)
  vertexMap := make(map[string]*vertex)
  return &Graph{v: &vertexMap, sccSizes: &sizesMap}
}

func (g *Graph) SCCs() map[int]int {
  return *(g.sccSizes)
}

func (g *Graph) PrintVertices() {
  fmt.Println("Vertices: ")
  for _, v := range *(*g).v {
    fmt.Println(v.val)
  }
}

func (g *Graph) PrintEdges() {
  for _, v := range *(*g).v {
    edges := []string{}

    for n := v.outgoing.head; n != nil; n = n.next {
      edges = append(edges, n.value.val)
    }
  }
}

func newVertex(val string) *vertex {
  return &vertex{val: val, position: 0, explored: false, scc: 0, outgoing: NewLinkedList(), incoming: NewLinkedList()}
}

func (g *Graph) VertexExists(val string) bool {
  _, ok := (*(*g).v)[val]
  if !ok {
    return false
  }

  return true
}

func (g * Graph) AddVertex(val string) {
  graphVertices := *(*g).v
  graphVertices[val] = newVertex(val)
  g.numVertices += 1
}

func (g *Graph) AddOutgoing(v string, w string) {
  graphVertices := *(*g).v
  vertex, vOk := graphVertices[v]
  neighbor, nOk := graphVertices[w]

  if vOk == false || nOk == false {
    panic(fmt.Errorf("Error adding outgoing edge from  %d to %d", w, v))
  }
  vertex.outgoing.Append(neighbor)

}

func (g *Graph) AddIncoming(v string, w string) {
  graphVertices := *(*g).v
  vertex, vOk := graphVertices[v]
  neighbor, nOk := graphVertices[w]

  if vOk == false || nOk == false {
    panic(fmt.Errorf("Error adding incoming edge to %d from %d\n", w, v))
  }
  vertex.incoming.Append(neighbor)
}

func (g *Graph) DfsTopoKosaraju() *[]*vertex {
  verticesInPosition := []*vertex{}
  currLabel := g.numVertices

  for _, vertex := range *(*g).v {
    if vertex.explored == false {
      verticesInPosition = append(g.DfsOnIncoming(vertex, &currLabel), verticesInPosition...)
    }
  }

  return &verticesInPosition
}

func (g *Graph) DfsOnIncoming(v *vertex, currLabel *int) []*vertex {
  vertices := []*vertex{}
  s := NewStack()
  s.Push(v)

  for s.IsEmpty() == false {
    currentVertex, ok := s.Pop()

    if !ok {
      panic(fmt.Errorf("Error popping from DFS stack. currentVertex value: %v", currentVertex.val))
    }

    if currentVertex.explored == false {
      currentVertex.explored = true
      vertices = append(vertices, currentVertex)

      for n := currentVertex.incoming.head; n != nil; n = n.next {
        vertex := n.value
        s.Push(vertex)
      }
    }
  }

  return vertices
}

func (g *Graph) markAllVerticesUnexplored() {
  for _, v := range *(*g).v {
    v.explored = false
  }
}

func vertexPositionsReversed(vertices *[]*vertex) *[]*vertex {
  reversed := []*vertex{}

  for i := len(*vertices)-1; i >= 0; i-- {
    reversed = append(reversed, (*vertices)[i])
  }

  return &reversed
}

func (g *Graph) Kosaraju() {
  vertexPositions := g.DfsTopoKosaraju()
  g.DfsSccTopo(vertexPositions)
}

func (g *Graph) FiveLargestSccSizes() []int {
  sizes := *g.sccSizes
  fiveLargest := []int{}

  for _, s := range sizes {
    fiveLargest = append(fiveLargest, s)
  }

  if len(fiveLargest) < 5 {
    diff := 5 - len(fiveLargest)
    appendedSlice := make([]int, diff)
    fiveLargest = append(fiveLargest, appendedSlice...)
  }
  
  sort.Sort(sort.Reverse(sort.IntSlice(fiveLargest)))
  return fiveLargest
}

func (g *Graph) DfsSccTopo(vertices *[]*vertex) {
  g.markAllVerticesUnexplored()
  numScc := 0

  for _, vertex := range *vertices {
    numScc += 1
    if vertex.explored == false {
      sccSize := g.DfsScc(vertex, numScc)
      (*g.sccSizes)[numScc] = sccSize
    }
  }
}

func (g *Graph) DfsScc(v *vertex, numScc int) int {
  size := 0
  s := NewStack()
  s.Push(v)

  for s.IsEmpty() == false {
    currentVertex, ok := s.Pop()

    if !ok {
      panic(fmt.Errorf("Error popping from DFS stack. currentVertex value: %v", currentVertex.val))
    }

    if currentVertex.explored == false {
      currentVertex.explored = true
      currentVertex.scc = numScc
      size += 1
       for n := currentVertex.outgoing.head; n != nil; n = n.next {
         vertex := n.value
         s.Push(vertex)
       }
     }
   }
   return size
}
