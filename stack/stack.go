package main

import "fmt"

type stackValueType int

type stackMember struct {
	value stackValueType
	next  *stackValueType
}

type StackHolder struct {
	pointer *Stack
}

type Stack struct {
	value stackValueType
	next  *Stack
}

func (Stack) new(item stackValueType) *Stack {
	return &Stack{item, nil}
}

func (StackHolder) New() *StackHolder {
	return &StackHolder{nil}
}

func (stackholder *StackHolder) push(item stackValueType) {
	pointer := &Stack{item, nil}

	if stackholder.pointer == nil {
		stackholder.pointer = pointer
		return
	}

	lastItem := stackholder.pointer
	previous := stackholder.pointer
	for lastItem != nil {
		previous = lastItem
		lastItem = lastItem.next
	}
	previous.next = pointer
}

func (stackholder *StackHolder) pop() {

	if stackholder.pointer == nil {
		fmt.Println("stack.next == nil")
		return
	}

	if stackholder.pointer.next == nil {
		fmt.Println("stack.next.next == nil => stack.next = nil")
		stackholder.pointer = nil
		return
	}

	previous := stackholder.pointer //remove parren needs to be done
	next := stackholder.pointer
	depth := 0
	for next.next != nil {
		depth++
		previous = next
		next = next.next
	}

	previous.next = nil
}

func (stackholder *StackHolder) Get() stackValueType {
	currentItem := stackholder.pointer
	if currentItem == nil {
		return 0
	}

	previous := currentItem
	for currentItem != nil {
		previous = currentItem
		currentItem = currentItem.next
	}

	return previous.value
}

func (stackholder *StackHolder) String() string {

	if stackholder.pointer == nil {
		return "{}"
	}

	var result string = fmt.Sprintf("%d", stackholder.pointer.value)
	for item := stackholder.pointer.next; item != nil; {
		result += "->" + fmt.Sprintf("%d", item.value)
		item = item.next
	}

	return result
}

func main() {
	stck := StackHolder{}.New()
	fmt.Println(stck, "last ", stck.Get())

	stck.push(1)
	fmt.Println(stck, "last ", stck.Get())

	stck.push(2)
	fmt.Println(stck, "last ", stck.Get())

	stck.pop()
	fmt.Println(stck, "last ", stck.Get())

	stck.push(3)
	fmt.Println(stck, "last ", stck.Get())

	stck.push(4)
	fmt.Println(stck, "last ", stck.Get())

	stck.pop()
	fmt.Println(stck, "last ", stck.Get())

	stck.pop()
	fmt.Println(stck, "last ", stck.Get())
}
