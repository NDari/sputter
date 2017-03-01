package util

// Value allows the Stack to take anything
type Value interface {
}

// Stack is your standard Stack interface
type Stack interface {
	Push(v Value)
	Peek() (v Value, ok bool)
	Pop() (v Value, ok bool)
}

// basicStack implements a non-concurrent Stack
type basicStack struct {
	head *entry
}

type entry struct {
	value Value
	next  *entry
}

// NewStack creates a new Stack instance
func NewStack() Stack {
	return &basicStack{}
}

// Push a Value onto the Stack
func (s *basicStack) Push(v Value) {
	if s.head == nil {
		s.head = &entry{v, nil}
		return
	}
	s.head = &entry{v, s.head}
}

// Peek the head of the Stack
func (s *basicStack) Peek() (Value, bool) {
	e := s.head
	if e != nil {
		return e.value, true
	}
	return nil, false
}

// Pop the head of the Stack
func (s *basicStack) Pop() (Value, bool) {
	e := s.head
	if e != nil {
		s.head = e.next
		return e.value, true
	}
	return nil, false
}
