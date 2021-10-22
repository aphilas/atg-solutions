package main

import (
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	// note: declare to be able to refer to it in RHS of definition for recursion
	var walker func(t *tree.Tree)

	walker = func(t *tree.Tree) {
		if t == nil {
			return
		}

		walker(t.Left)
		ch <- t.Value
		walker(t.Right)
	}

	walker(t)
	close(ch)
}

func main() {

}
