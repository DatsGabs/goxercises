package main

import "fmt"

type Customer struct {
	name string
	next *Customer
	prev *Customer
}

type Queue struct {
	head *Customer
	tail *Customer
}

func (queue *Queue) Enqueue(name string) error {
	customer := &Customer {
		name: name,
	}
	if queue.head == nil {
		queue.head = customer
	} else {
		current := queue.tail
		current.next = customer
		customer.prev = current
	}
	queue.tail = customer
	return nil
} 

func (queue *Queue) Dequeue() error {
	if queue == nil {
		fmt.Println("The queue is empty")
		return nil
	}
	if queue.head.next != nil {
		queue.head = queue.head.next
		queue.head.prev = nil
		return nil
	}
	queue.head = nil
	return nil
}

func (queue Queue) showAllCustomers() error {
	current := queue.head
	if current == nil {
		fmt.Println("The queue is empty")
		return nil
	}

	fmt.Println(current.name)
	for current.next != nil {
		current = current.next
		fmt.Println(current.name)
	}
	return nil
}

func main()  {
	var queue Queue
	queue.Enqueue("Gabriel")
	queue.Enqueue("Alfredo")
	queue.Dequeue()
	queue.Dequeue()
	queue.Enqueue("Gabriel")
	queue.Enqueue("Alfredo")
	queue.showAllCustomers()
}