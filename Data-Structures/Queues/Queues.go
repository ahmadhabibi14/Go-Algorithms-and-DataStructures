package main

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() (int, error) {
	if len(q.items) == 0 {
		return 0, fmt.Errorf("queue is empty")
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}

func main() {
	queue := Queue{}
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	item, err := queue.Dequeue()
	if err == nil {
		fmt.Println("Dequeued item: ", item)
	}

	item, err = queue.Dequeue()
	if err == nil {
		fmt.Println("Dequeued item: ", item)
	}

	item, err = queue.Dequeue()
	if err == nil {
		fmt.Println("Dequeued item: ", item)
	} else {
		fmt.Println("Error: ", err.Error())
	}
}

/*
In this example, we define a struct Queue that has a slice items to store the elements of the queue.
The Enqueue method appends an item to the end of the items slice, effectively adding it to the queue.
The Dequeue method removes and returns the first item from the items slice, simulating the behavior of a queue.

In the main function, we create an instance of the Queue struct, enqueue three items (10, 20, and 30), and then
dequeue and print them one by one. The Dequeue method returns an error if the queue is empty,
so we handle the error appropriately.

This implementation provides a basic example of a queue in Go using slices. However, it's worth noting that
if you need a more efficient implementation or plan to use the queue in a concurrent setting,
you might want to consider using a library or a more sophisticated data structure implementation.
*/
