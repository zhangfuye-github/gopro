package main

import "fmt"

type link struct {
	head *link
	data interface{}
	next *link
	tail *link
}

func (this *link) inserttail(data interface{}) {
	node := &link{
		data: data,
		next: nil,
	}

	if this.tail == nil && this.head == nil {
		this.tail = node
		this.head = node
		return
	}
	this.tail.next = node
	this.tail = node
}
func (this *link) insertfont(data interface{}) {
	node := &link{
		data: data,
		next: nil,
	}
	node.next = this.head
	this.head = node
}
func (this *link) trans() {
	q := this.head
	for q.next != nil {
		fmt.Println(q.data)
		q = q.next
	}
}
