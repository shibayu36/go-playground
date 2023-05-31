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
	return false
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for val := range ch {
		println(val)
	}
}
