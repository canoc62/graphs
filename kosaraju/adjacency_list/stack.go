package adjacency_list

type stack struct {
  list *LinkedList
}

func NewStack() *stack {
 return &stack{list: NewLinkedList()}
}

func (s *stack) Push(v *vertex) {
  s.list.Append(v)
}

func (s *stack) Pop() (*vertex, bool) {
  node, ok := s.list.RemoveTail()

  if !ok {
    return nil, false
  }

  return node.value, true
}

func (s *stack) IsEmpty() bool {
	return s.list.size == 0
}
