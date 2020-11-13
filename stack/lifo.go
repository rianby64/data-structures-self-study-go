package lifo

import (
	"github.com/rianby64/data-structures-self-study/doublylinkedlist"
)

// Stack interface
type Stack interface {
	Pop() interface{}
	Push(c interface{})
}

type lifo struct {
	list doublylinkedlist.DoublyLinkedList
}

func (l *lifo) Pop() interface{} {
	item := l.list.Last()
	item.Delete()
	return item.Value()
}

func (l *lifo) Push(c interface{}) {
	item := l.list.Last()
	item.Insert(c)
}

// New constructs an stack on top of a lifo
func New() Stack {
	return &lifo{
		list: doublylinkedlist.New(),
	}
}
