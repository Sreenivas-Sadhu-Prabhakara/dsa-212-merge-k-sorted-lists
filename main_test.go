package main

import (
	"reflect"
	"testing"
)

func TestMergeK(t *testing.T) {
	got := toSlice(MergeK([]*Node{
		fromSlice([]int{1, 4, 5}),
		fromSlice([]int{1, 3, 4}),
		fromSlice([]int{2, 6}),
	}))
	want := []int{1, 1, 2, 3, 4, 4, 5, 6}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("MergeK = %v, want %v", got, want)
	}

	if MergeK(nil) != nil {
		t.Error("no lists -> nil")
	}
	if MergeK([]*Node{nil, nil}) != nil {
		t.Error("all-nil lists -> nil")
	}
	single := toSlice(MergeK([]*Node{fromSlice([]int{7, 8})}))
	if !reflect.DeepEqual(single, []int{7, 8}) {
		t.Errorf("single list = %v", single)
	}
	mixed := toSlice(MergeK([]*Node{nil, fromSlice([]int{0}), nil}))
	if !reflect.DeepEqual(mixed, []int{0}) {
		t.Errorf("mixed nils = %v", mixed)
	}
}
