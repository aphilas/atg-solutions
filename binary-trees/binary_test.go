package main

import (
	"reflect"
	"testing"

	"golang.org/x/tour/tree"
)

func TestWalker(t *testing.T) {
	t.Run("runs inorder traversal", func(t *testing.T) {
		r := &tree.Tree{
			Left: &tree.Tree{
				Left:  &tree.Tree{Value: 4},
				Value: 2,
				Right: &tree.Tree{Value: 5},
			},
			Value: 1,
			Right: &tree.Tree{
				Value: 3,
			},
		}

		ch := make(chan int)
		got := make([]int, 0)
		go Walk(r, ch) // inorder traversal

		for v := range ch {
			got = append(got, v)
		}

		want := []int{4, 2, 5, 1, 3}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got: %v, want: %v", got, want)
		}
	})

	t.Run("passes tour test", func(t *testing.T) {
		ch := make(chan int)
		go Walk(tree.New(1), ch)

		got := make([]int, 0)
		for v := range ch {
			got = append(got, v)
		}

		want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		assertContains(t, got, want)
	})
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func assertContains(t testing.TB, got, want []int) {
	if !(len(got) == len(want)) {
		t.Errorf("Got is len: %v, want len: %v", len(got), len(want))
	}

	for _, v := range want {
		if !contains(got, v) {
			t.Errorf("Got: %v, want: %v", got, want)
		}
	}
}
