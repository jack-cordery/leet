package leet_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"jack.com/leet/leet"
)

func TestPrint(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		head := &leet.ListNode{Val: 1, Next: &leet.ListNode{Val: 2, Next: &leet.ListNode{Val: 3}}}
		l := leet.LinkedList{head}
		l.Print()
		assert.Equal(t, 1, 1)
	})
}

func TestFromVec(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		in := []int{1, 2, 3}
		expected := &leet.ListNode{Val: 1, Next: &leet.ListNode{Val: 2, Next: &leet.ListNode{Val: 3}}}
		out := leet.LinkedList{}.FromVec(in)
		out.Print()
		assert.Equal(t, expected, out.Head)

	})
}
