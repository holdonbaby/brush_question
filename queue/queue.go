package queue

// 队列先进先出

type Queue struct {
	Content []int
}

func (q *Queue) Push(n int) {
	head := []int{n}
	q.Content = append(head, q.Content...)
}

func (q *Queue) Pull() int {
	if len(q.Content) == 0 {
		return -1
	}
	ret := q.Content[0]
	q.Content = q.Content[1:]
	return ret
}
