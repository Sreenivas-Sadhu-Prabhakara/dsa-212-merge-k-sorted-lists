package main

import (
	"container/heap"
	"fmt"
)

// Node is a singly linked list node.
type Node struct {
	Val  int
	Next *Node
}

type nodeHeap []*Node

func (h nodeHeap) Len() int           { return len(h) }
func (h nodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h nodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *nodeHeap) Push(x any)        { *h = append(*h, x.(*Node)) }
func (h *nodeHeap) Pop() any {
	old := *h
	v := old[len(old)-1]
	*h = old[:len(old)-1]
	return v
}

// MergeK splices k sorted lists into one sorted list, reusing the nodes.
// A min-heap holds each list's current head; popping the smallest advances
// only that list.
func MergeK(lists []*Node) *Node {
	h := &nodeHeap{}
	for _, l := range lists {
		if l != nil {
			*h = append(*h, l)
		}
	}
	heap.Init(h)
	dummy := &Node{}
	tail := dummy
	for h.Len() > 0 {
		n := heap.Pop(h).(*Node)
		if n.Next != nil {
			heap.Push(h, n.Next)
		}
		tail.Next = n
		tail = n
	}
	tail.Next = nil
	return dummy.Next
}

func fromSlice(vals []int) *Node {
	dummy := &Node{}
	cur := dummy
	for _, v := range vals {
		cur.Next = &Node{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

func toSlice(head *Node) []int {
	out := []int{}
	for n := head; n != nil; n = n.Next {
		out = append(out, n.Val)
	}
	return out
}

func main() {
	merged := MergeK([]*Node{fromSlice([]int{1, 4, 5}), fromSlice([]int{1, 3, 4}), fromSlice([]int{2, 6})})
	fmt.Println(toSlice(merged))
}
