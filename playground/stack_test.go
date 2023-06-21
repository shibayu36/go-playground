package main

type Stack[T any] []T

func (s *Stack[T]) Push(v T) {
	*s = append(*s, v)
}

func (s *Stack[T]) Pop() T {
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

// func main() {
// 	var s Stack[int]
// 	println(s.IsEmpty())
// 	s.Push(1)
// 	s.Push(2)
// 	s.Push(3)
// 	println(s.Pop())
// 	println(s.Pop())
// 	println(s.Pop())
// }
