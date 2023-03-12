package queue

// 队列先进先出

type Queue struct {
	Content []int
}

func NewQueue() *Queue {
	return &Queue{Content: make([]int, 0)}
}

func (q *Queue) Push(n int) *Queue {
	q.Content = append(q.Content, n)
	return q
}

func (q *Queue) Pull() int {
	if len(q.Content) == 0 {
		return -1
	}
	ret := q.Content[0]
	q.Content = q.Content[1:]
	return ret
}
