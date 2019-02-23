package adjacency_list

import "fmt"

type LinkedList struct {
  head, tail *Node
  size int
}

type Node struct {
	value *vertex
	prev, next *Node
}

func (node Node) Value() *vertex {
  return node.value
}

func NewLinkedList() *LinkedList {
  return &LinkedList{ head: nil, tail: nil, size: 0 }
}

func (ll *LinkedList) Print() {
  for node := ll.head; node != nil; node = node.next {
    fmt.Println(node.Value())
  }
}

func (ll *LinkedList) PrintReverse() {
  for node := ll.tail; node != nil; node = node.prev {
    fmt.Println(node.Value())
  }
}

func(ll *LinkedList) LinearSearch(w *vertex) bool {
  for node := ll.head; node != nil; node = node.next {
		if (*node).value == w {
			return true
		}
	}

	return false
}

func (ll *LinkedList) Prepend(w *vertex) {
  newNode := Node{ value: w, prev: nil, next: ll.head }
  if ll.head != nil {
    ll.head.prev = &newNode
  }
  if ll.tail == nil {
    ll.tail = &newNode
  }
  ll.head = &newNode
  ll.size += 1
}

func (ll *LinkedList) Append(w *vertex) {
  newNode := Node{ value: w, prev: ll.tail, next: nil }
  if ll.tail != nil {
    ll.tail.next = &newNode
  }
  if ll.head == nil {
    ll.head = &newNode
  }
  ll.tail = &newNode
  ll.size += 1
}

func (ll *LinkedList) RemoveHead() (*Node, bool) {
  if ll.head == nil {
    return nil, false
  }

  removed := ll.head
  if removed == ll.tail {
    ll.tail = nil
  }

  ll.head = ll.head.next
  if ll.head != nil {
    ll.head.prev = nil
  }
  ll.size -= 1

  return removed, true
}

func (ll *LinkedList) RemoveTail() (*Node, bool) {
  if ll.tail == nil {
    return nil, false
  }

  removed := ll.tail
  if removed == ll.head {
    ll.head = nil
  } else {
    ll.tail = ll.tail.prev
  }
  if ll.tail != nil {
    ll.tail.next = nil
  }
  ll.size -= 1

  return removed, true
}
