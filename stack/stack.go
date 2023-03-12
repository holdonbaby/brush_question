package stack

// 先进后出的stack

type Stack struct {
	Content []int
}

func NewStack() *Stack {
	return &Stack{Content: make([]int, 0)}
}

func (s *Stack) Push(n int) *Stack {
	s.Content = append(s.Content, n)
	return s
}

func (s *Stack) Pull() int {
	if len(s.Content) == 0 {
		return -1
	}
	ret := s.Content[len(s.Content)-1]
	s.Content = s.Content[:len(s.Content)-1]
	return ret
}
