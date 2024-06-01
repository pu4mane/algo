package linkedlist

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	// Test Read method
	t.Run("Read from empty list", func(t *testing.T) {
		list := LinkedList{}
		_, ok := list.Read(0)
		if ok {
			t.Error("Expected Read from empty list to return false")
		}
	})

	// Test IndexOf method
	t.Run("IndexOf existing element", func(t *testing.T) {
		list := LinkedList{}
		list.InsertAtIndex(0, 10)
		list.InsertAtIndex(1, 20)
		index, ok := list.IndexOf(20)
		if !ok || index != 1 {
			t.Error("Expected IndexOf to return index 1")
		}
	})

	// Test InsertAtIndex method
	t.Run("Insert at beginning of empty list", func(t *testing.T) {
		list := LinkedList{}
		list.InsertAtIndex(0, 10)
		data, ok := list.Read(0)
		if !ok || data != 10 {
			t.Error("Expected inserted element to be at index 0")
		}
	})
}
