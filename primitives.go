package sputter

import "fmt"

// UnknownSymbol is thrown if a symbol cannot be resolved
const UnknownSymbol = "Symbol has not been defined"

// Value is the generic interface for all 'Values' in the VM
type Value interface {
}

// Iterable values can be used in loops and comprehensions
type Iterable interface {
	Iterate() Iterator
}

// Iterator interfaces are stateful iteration interfaces
type Iterator interface {
	Next() (Value, bool)
	Iterable() Iterable
}

// Literal identifies a Value as being a literal reference
type Literal struct {
	value Value
}

// Evaluate makes a Literal Evaluable
func (l *Literal) Evaluate(c *Context) Value {
	return l.value
}

func (l *Literal) String() string {
	if str, ok := l.value.(fmt.Stringer); ok {
		return str.String()
	}
	return l.value.(string)
}

// Symbol is an Identifier that can be resolved
type Symbol struct {
	name string
}

// Evaluate makes a Symbol Evaluable
func (s *Symbol) Evaluate(c *Context) Value {
	if resolved, ok := c.Get(s.name); ok {
		return resolved
	}
	panic(UnknownSymbol)
}

func (s *Symbol) String() string {
	return s.name
}

// ArgumentProcessor is the standard signature for a function that is
// capable of processing an Iterable (like Lists)
type ArgumentProcessor func(*Context, Iterable) Value

// Function is a Value that can be invoked
type Function struct {
	name string
	exec ArgumentProcessor
}

func (f *Function) String() string {
	return f.name
}