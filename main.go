package main

import (
	"fmt"
	"gustavocoutino/hashmap"
	"gustavocoutino/queue"
	"gustavocoutino/stack"
)

func main() {
	// Stacks
	fmt.Println("Stacks")
	st := stack.New[int]()
	st.Push(1)
	st.Push(2)
	st.Push(3)
	st.Push(4)
	st.Push(5)

	// Prints 5 to 1
	size := len(st.Items)
	for i := 0; i < size; i++ {
		v, _ := st.Top()
		fmt.Println(v)
		st.Pop()
	}
	fmt.Println("--------------------------------------------------------")

	// Queues
	fmt.Println("Queues")

	q := queue.New[int]()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	q.Push(5)

	// Prints 1 to 5
	size = len(q.Items)
	for i := 0; i < size; i++ {
		v, _ := q.Front()
		fmt.Println(v)
		q.Pop()
	}

	fmt.Println("--------------------------------------------------------")

	q = queue.New[int]()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	q.Push(5)
	v, _ := q.Back()
	// Prints 5
	fmt.Println(v)
	fmt.Println("--------------------------------------------------------")

	// Hashmap
	fmt.Println("Hashmap")
	ht := hashmap.New[string, int](5)
	ht.Insert("Hello", 1)
	ht.Insert("Goodbye", 2)
	ht.Insert("Bonjour", 3)
	ht.Insert("Au revoir", 4)
	ht.Insert("Adieu", 5)

	k, v, err := ht.Get("Hello")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Key: %q, value: %d\n", k, v)
	}

	err = ht.Remove("Bonjour")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Removed %q, size is now %d\n", "Bonjour", ht.Size())
	}

	_, _, err = ht.Get("Bonjour")
	if err != nil {
		fmt.Println(err)
	}
}
