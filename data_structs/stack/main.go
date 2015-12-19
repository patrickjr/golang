// shamelessy stole from https://gist.github.com/bemasher/1777766

package main

import "fmt"

type Stack struct {
	top  *Element
	size int
}

type Element struct {
	value interface{}
	next  *Element
}

func (s *Stack) Length() int {
	return s.size
}

func (s *Stack) Push(value interface{}) {
	s.top = &Element{value, s.top}
	s.size++
}

// Remove the top element from the stack and return it's value
// If the stack is empty, return nil
func (s *Stack) Pop() (value interface{}) {
	if s.size > 0 {
		value = s.top.value
		s.top = s.top.next
		s.size--
		return
	}
	return nil
}

func (s *Stack) Peek() (value interface{}) {
	if s.size == 0 {
		return nil
	}
	return s.top.value
}

func main() {
	stack := new(Stack)

	stack.Push("Things")
	stack.Push("and")
	stack.Push("Stuff")

	info := stack.Peek()
	fmt.Println(info.(string))

	for stack.Length() > 0 {
		// We have to do a type assertion because we get back a variable of type
		// interface{} while the underlying type is a string.
		fmt.Printf("%s ", stack.Pop().(string))
	}
	fmt.Println()

}
