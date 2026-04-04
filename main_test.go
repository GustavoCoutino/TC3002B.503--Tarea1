package main

import (
	"gustavocoutino/hashmap"
	"gustavocoutino/queue"
	"gustavocoutino/stack"
	"gustavocoutino/utils"
	"testing"
)

func TestMain_StackLIFO(t *testing.T) {
	st := stack.New[int]()
	for i := 1; i <= 5; i++ {
		st.Push(i)
	}

	want := []int{5, 4, 3, 2, 1}
	for _, expected := range want {
		got, err := st.Top()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got != expected {
			t.Errorf("got %d, want %d", got, expected)
		}
		st.Pop()
	}

	if !st.Empty() {
		t.Error("expected stack to be empty after popping all elements")
	}
}

func TestMain_QueueFIFO(t *testing.T) {
	q := queue.New[int]()
	for i := 1; i <= 5; i++ {
		q.Push(i)
	}

	want := []int{1, 2, 3, 4, 5}
	for _, expected := range want {
		got, err := q.Front()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got != expected {
			t.Errorf("got %d, want %d", got, expected)
		}
		q.Pop()
	}

	if !q.Empty() {
		t.Error("expected queue to be empty after popping all elements")
	}
}

func TestMain_QueueBack(t *testing.T) {
	q := queue.New[int]()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	q.Push(5)

	got, err := q.Back()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != 5 {
		t.Errorf("got %d, want 5", got)
	}
}

func TestMain_HashmapGetInsertRemove(t *testing.T) {
	ht := hashmap.New[int](5, utils.FNVHash)
	ht.Insert("Hello", 1)
	ht.Insert("Goodbye", 2)
	ht.Insert("Bonjour", 3)
	ht.Insert("Au revoir", 4)
	ht.Insert("Adieu", 5)

	k, v, err := ht.Get("Hello")
	if err != nil {
		t.Fatalf("unexpected error on Get: %v", err)
	}
	if k != "Hello" || v != 1 {
		t.Errorf("got key=%q value=%d, want key=%q value=%d", k, v, "Hello", 1)
	}

	err = ht.Remove("Bonjour")
	if err != nil {
		t.Fatalf("unexpected error on Remove: %v", err)
	}
	if ht.Size() != 4 {
		t.Errorf("got size %d, want 4", ht.Size())
	}

	_, _, err = ht.Get("Bonjour")
	if err == nil {
		t.Error("expected error after removing key, got nil")
	}
}
