package dijkstra

import "fmt"

type heap []*vertex

func (h *heap) Print() {
  elements := []int{}
  for _, v := range *h {
    elements = append(elements, v.score)
  }

  fmt.Println(elements)
}

func parent(i int) int {
  return (i - 1)/2
}

func right(i int) int {
  return (2*i) + 1
}

func left(i int) int {
  return (2*i) + 2
}

func Heapify(h *heap, mappings *map[string]int, idx int) *heap {
  hp := *h
  indexMap := *mappings
  ldx := left(idx)
  rdx := right(idx)

  smallest := idx

  if (ldx < len(hp)) && (hp[ldx].score < hp[idx].score) {
    smallest = ldx
  }
  if (rdx < len(hp)) && (hp[rdx].score < hp[smallest].score) {
    smallest = rdx
  }
  if smallest != idx {
    hp[smallest], hp[idx] = hp[idx], hp[smallest]
    indexMap[hp[smallest].id] = smallest
    indexMap[hp[idx].id] = idx
    Heapify(h, mappings, smallest)
  }

  return h
}

func (h *heap) Delete(mappings *map[string]int, idx int) (*vertex, bool) {
  hp := *h

  if len(hp) <= idx {
    return nil, false
  }

  deleted := (hp)[idx]

  if idx < len(hp) - 1 {
    hp[idx] = hp[len(hp)-1]
    (*mappings)[hp[idx].id] = idx
    *h = hp[:len(hp)-1]

    parentIdx := parent(idx)

    if parentIdx >= 0 && hp[parentIdx].score > hp[idx].score {
      hp.bubbleUp(mappings, idx)
    } else {
      Heapify(h, mappings, idx)
    }
  } else {
    *h = hp[:len(hp)-1]
  }

  return deleted, true
}

func (h *heap) ExtractMin(mappings *map[string]int) (*vertex, bool) {
  return h.Delete(mappings, 0)
}

func (h *heap) Insert(mappings *map[string]int, v *vertex) {
  *h = append(*h, v)
  idx := len(*h) - 1

  h.bubbleUp(mappings, idx)
}

func (h *heap) bubbleUp(mappings *map[string]int, idx int) {
  hp := *h
  indexMap := *mappings

  for idx > 0 && hp[parent(idx)].score > hp[idx].score {
    hp[parent(idx)], hp[idx] = hp[idx], hp[parent(idx)]
    indexMap[hp[idx].id] = idx
    indexMap[hp[parent(idx)].id] = parent(idx)
    idx = parent(idx)
  }
}