package api

import "bytes"

// EmptyList represents an empty List
var EmptyList = &List{nil, nil}

// List contains a node to a singly-linked List
type List struct {
	first Value
	rest  *List
}

// NewList creates a new List instance
func NewList(v Value) Sequence {
	return &List{v, EmptyList}
}

// First returns the first element of a List
func (l *List) First() Value {
	return l.first
}

// Rest returns the rest of the List as a Sequence
func (l *List) Rest() Sequence {
	return l.rest
}

// IsSequence returns whether this instance is a consumable Sequence
func (l *List) IsSequence() bool {
	return l != EmptyList
}

// Prepend creates a new Sequence by prepending a Value
func (l *List) Prepend(v Value) Sequence {
	return &List{first: v, rest: l}
}

// Count returns the length of the List
func (l *List) Count() int {
	r := 0
	for e := l; e != EmptyList; e = e.rest {
		r++
	}
	return r
}

// Get returns the Value at the indexed position in the List
func (l *List) Get(index int) Value {
	var i = 0
	for e := l; e != EmptyList; e = e.rest {
		if i == index {
			return e.first
		}
		i++
	}
	return Nil
}

// Eval makes a List Evaluable
func (l *List) Eval(ctx Context) Value {
	if l == EmptyList {
		return EmptyList
	}
	if f, ok := ResolveAsFunction(ctx, l.first); ok {
		return f.Apply(ctx, l.rest)
	}
	panic(ExpectedFunction)
}

func (l *List) String() string {
	if l == EmptyList {
		return "()"
	}

	var b bytes.Buffer
	b.WriteString("(")
	for e := l; e != EmptyList; e = e.rest {
		b.WriteString(String(e.first))
		if e.rest != EmptyList {
			b.WriteString(" ")
		}
	}
	b.WriteString(")")
	return b.String()
}