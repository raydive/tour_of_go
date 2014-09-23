package main

import "code.google.com/p/go-tour/tree"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	recusive_work(t, ch)
	close(ch)
}

func recusive_work(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	recusive_work(t.Left, ch)
	ch <- t.Value
	recusive_work(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	go Walk(t1, ch1)
	ch2 := make(chan int)
	go Walk(t2, ch2)

	for n := range ch1 {
		if n != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for n := range ch {
		print(n, " ")
	}
	println()

	println(Same(tree.New(1), tree.New(1)))
	println(Same(tree.New(1), tree.New(2)))
}
