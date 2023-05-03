package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	len  int
	head *ListItem
	tail *ListItem
}

func NewList() List {
	return new(list)
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.head
}

func (l *list) Back() *ListItem {
	return l.tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	prevHead := l.head
	l.head = &ListItem{Next: prevHead, Prev: nil, Value: v}

	if l.len == 0 {
		l.tail = l.head
	}

	if prevHead != nil {
		prevHead.Prev = l.head
	}
	l.len++

	return l.head
}

func (l *list) PushBack(v interface{}) *ListItem {
	prevTail := l.tail
	l.tail = &ListItem{Prev: prevTail, Next: nil, Value: v}
	prevTail.Next = l.tail
	l.len++

	if prevTail == l.head {
		l.head.Next = l.tail
	}

	return l.tail
}

func (l *list) Remove(i *ListItem) {
	if i != l.head {
		i.Prev.Next = i.Next
	}

	if i == l.tail {
		l.tail = i.Prev
		l.tail.Next = nil
	}

	if i == l.head {
		l.head = nil
		l.head.Prev = nil
	}

	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if i == l.head {
		return
	}

	// remove from current place
	prev := i.Prev
	next := i.Next
	prev.Next = next

	if i == l.tail {
		l.tail = prev
		l.tail.Next = nil
	}

	// insert before head
	prevHead := l.head
	l.head = i

	i.Next = prevHead
	i.Prev = nil

	prevHead.Prev = i
}
