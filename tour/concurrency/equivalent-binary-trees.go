package main

import (
	"fmt"

	"github.com/rezilience/readytogo/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	recurse(t, ch, 0)
}

func recurse(t *tree.Tree, ch chan int, rec int) {
	if t == nil {
		return
	}

	recurse(t.Left, ch, rec+1)
	ch <- t.Value
	recurse(t.Right, ch, rec+1)

	if rec == 0 {
		close(ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)

	go Walk(t1, c1)
	go Walk(t2, c2)

	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2

		if !ok1 && !ok2 {
			return true
		}

		if !ok1 || !ok2 || v1 != v2 {
			return false
		}
	}
}

func main() {

	c1 := make(chan int)
	go Walk(tree.New(1), c1)

	for v := range c1 {
		fmt.Print(v, " ")
	}
	fmt.Println()

	fmt.Println(Same(tree.New(1), tree.New(2)))
	fmt.Println(Same(tree.New(1), tree.New(1)))
}
