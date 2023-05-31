package main

import "golang.org/x/tour/tree"

func walkInternal(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	walkInternal(t.Left, ch)
	ch <- t.Value
	walkInternal(t.Right, ch)
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walkInternal(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for val1 := range ch1 {
		val2, ok := <-ch2

		if !ok {
			return false
		}

		if val1 != val2 {
			return false
		}
	}

	_, ok := <-ch2
	if ok {
		return false
	} else {
		return true
	}
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for val := range ch {
		println(val)
	}

	println(Same(tree.New(1), tree.New(1)))
	println(Same(tree.New(1), tree.New(2)))
}
