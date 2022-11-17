package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
	Clear()
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	len   int
	front *ListItem
	back  *ListItem
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	i := &ListItem{Value: v, Next: nil, Prev: nil}
	if l.len == 0 {
		l.front = i
		l.back = i
	} else {
		i.Next = l.front
		l.front.Prev = i
		l.front = i
	}
	l.len++
	return i
}

func (l *list) PushBack(v interface{}) *ListItem {
	i := &ListItem{Value: v, Next: nil, Prev: nil}
	if l.len == 0 {
		l.front = i
		l.back = i
	} else {
		l.back.Next = i
		i.Prev = l.back
		l.back = i
	}
	l.len++
	return i
}

func (l *list) Remove(i *ListItem) {
	if l.front == i {
		l.front = i.Next
	}
	if l.back == i {
		l.back = i.Prev
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove(i)
	l.PushFront(i.Value)
}

func (l *list) Clear() {
	l.len = 0
	l.front = nil
	l.back = nil
}

func NewList() List {
	return new(list)
}
