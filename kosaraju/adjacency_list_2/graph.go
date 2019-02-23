package adjacency_list

import (
  "fmt"
  "sort"
)

type Graph struct {
  v *map[string]*vertex
  sccSizes *map[int]int
}

type vertex struct {
  value string
  position int
  explored bool
  scc int
	outgoing *[]*vertex
  incoming *[]*vertex
}

func NewGraph() *Graph {
  sizesMap := make(map[int]int)
  vertexMap := make(map[string]*vertex)
  return &Graph{v: &vertexMap, sccSizes: &sizesMap}
}

func newVertex(value string) *vertex {
  outgoing := make([]*vertex, 0)
  incoming := make([]*vertex, 0)
  return &vertex{value: value, position: 0, explored: false, scc: 0, outgoing: &outgoing, incoming: &incoming}
}

func (g *Graph) VertexExists(value string) bool {
  _, ok := (*(*g).v)[value]
  if !ok {
    return false
  }

  return true
}

func (g *Graph) AddVertex(value string) {
  (*(*g).v)[value] = newVertex(value)
}

func (g *Graph) AddOutgoing(v string, w string) {
  graphVertices := *(*g).v

  vertex, vOk := graphVertices[v]
  neighbor, nOk := graphVertices[w]

  if vOk == false || nOk == false {
    panic(fmt.Errorf("Error adding outgoing edge from  %d to %d", w, v))
  }
  *vertex.outgoing = append(*vertex.outgoing, neighbor)
}

func (g *Graph) AddIncoming(v string, w string) {
  graphVertices := *(*g).v
  vertex, vOk := graphVertices[v]
  neighbor, nOk := graphVertices[w]

  if vOk == false || nOk == false {
    panic(fmt.Errorf("Error adding incoming edge to %d from %d, w, v"))
  }
  *vertex.incoming = append(*vertex.incoming, neighbor)
}

func (g *Graph) dfsTopoKosaraju() *[]*vertex {
  verticesInPosition := []*vertex{}

  for _, vertex := range *(*g).v {
    if vertex.explored == false {
      verticesInPosition = append(g.dfsOnIncoming(vertex), verticesInPosition...)
    }
  }

  return &verticesInPosition
}

func (g *Graph) dfsOnIncoming(v *vertex) []*vertex {
  vertices := []*vertex{}
  stack := []*vertex{v}

  for len(stack) > 0 {
    currentVertex := stack[len(stack)-1]
    stack = stack[:len(stack)-1]

    if currentVertex.explored == false {
      currentVertex.explored = true
      vertices = append(vertices, currentVertex)

      for _, vertex := range *currentVertex.incoming {
        if vertex.explored == false {
          stack = append(stack, vertex)
        }
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

func (g *Graph) Kosaraju() {
  vertexPositions := g.dfsTopoKosaraju()
  g.dfsSccTopo(vertexPositions)
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
  return fiveLargest[0:5]
}

func (g *Graph) dfsSccTopo(vertices *[]*vertex) {
  g.markAllVerticesUnexplored()
  numScc := 0

  for _, vertex := range *vertices {
    numScc += 1
    if vertex.explored == false {
      sccSize := g.dfsScc(vertex, numScc)
      (*g.sccSizes)[numScc] = sccSize
    }
  }
}

func (g *Graph) dfsScc(v *vertex, numScc int) int {
  size := 0
  stack := []*vertex{v}

  for len(stack) > 0 {
    currentVertex := stack[len(stack)-1]
    stack = stack[:len(stack)-1]

    if currentVertex.explored == false {
      currentVertex.explored = true
      currentVertex.scc = numScc
      size += 1
      for _, vertex := range *currentVertex.outgoing {
        stack = append(stack, vertex)
      }
    }
  }
   return size
}