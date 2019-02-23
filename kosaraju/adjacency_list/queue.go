package adjacency_list

type queue struct {
  list *LinkedList
}

func NewQueue() *queue {
  return &queue{list: NewLinkedList()}
}

func (q *queue) Enqueue(v *vertex) {
  q.list.Append(v)
}

func (q *queue) Dequeue() (*vertex, bool) {
  node, ok := q.list.RemoveHead()

  if !ok {
    return nil, false
  }

  return node.Value(), true
}

func (q *queue) IsEmpty() bool {
  return q.list.size == 0
}

func (q *queue) PrintVals() {
  q.list.Print()
}