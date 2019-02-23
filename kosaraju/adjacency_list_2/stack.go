package adjacency_list

type stack []*vertex

func newStack() *stack {
  var initStack stack

 return &initStack
}

func (s *stack) push(v *vertex) {
  *s = append(*s, v)
}

func (s *stack) pop() (*vertex, bool) {
  if len(*s) < 1 {
    return nil, false
  }
  lastIndex := len(*s) - 1
  v := (*s)[lastIndex]
  *s = (*s)[:lastIndex]
  return v, true
}

func (s *stack) isEmpty() bool {
	return len(*s) > 0
}
