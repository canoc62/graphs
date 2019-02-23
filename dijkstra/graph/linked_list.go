package dijkstra

type linkedList struct {
  head, tail *node
  size int
}

type node struct {
	data *vertex
  weight int
	prev, next *node
}

func newLinkedList() *linkedList {
  return &linkedList{ head: nil, tail: nil, size: 0 }
}

func (ll *linkedList) put(w *vertex, weight int) {
  newnode := node{ data: w, weight: weight, prev: ll.tail, next: nil }
  if ll.tail != nil {
    ll.tail.next = &newnode
  }
  if ll.head == nil {
    ll.head = &newnode
  }
  ll.tail = &newnode
  ll.size += 1
}
