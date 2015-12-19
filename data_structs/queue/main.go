package main

import "fmt"

type Queue struct {
	front *Element
	rear  *Element
}

type Element struct {
	value interface{}
	next  *Element
}

func (q *Queue) Enqueue(value interface{}) {

	temp := new(Element)
	temp.value = value
	temp.next = nil
	if q.front == nil && q.rear == nil {
		q.front = temp
		q.rear = temp
		return
	}

	q.rear.next = temp
	q.rear = temp
}

func (q *Queue) Dequeue() (value interface{}) {
	temp := new(Element)
	temp = q.front
	q.front = q.front.next
	return temp
}

func main() {
	queue := new(Queue)
	queue.Enqueue("1")
	queue.Enqueue("2")
	queue.Enqueue("3")
	queue.Enqueue("4")
	queue.Enqueue("5")
	queue.Enqueue("6")
	queue.Enqueue("7")

	fmt.Println("from front")
	fmt.Println("---------")
	fmt.Println(queue.front)
	fmt.Println(queue.front.next)
	fmt.Println(queue.front.next.next)
	fmt.Println(queue.front.next.next.next)
	fmt.Println(queue.front.next.next.next.next)
	fmt.Println(queue.front.next.next.next.next.next)
	fmt.Println(queue.front.next.next.next.next.next.next)
	fmt.Println(queue.front.next.next.next.next.next.next.next)

	fmt.Println(queue.Dequeue())

}
