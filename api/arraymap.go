package api

import "bytes"

const (
	// ExpectedPair is thrown if you prepend to a Map incorrectly
	ExpectedPair = "expected a two element vector when prepending"

	// ExpectedMapped is thrown if the Value is not a Mapped item
	ExpectedMapped = "expected a mapped value"
)

// Mapped interfaces allow a Sequence item to be retrieved by Name
type Mapped interface {
	Sequence
	Get(key Value) Value
}

// ArrayMap is a Mappable that is implemented atop an array
type ArrayMap []Vector

// Count returns the number of key/value pairs in this ArrayMap
func (m ArrayMap) Count() int {
	return len(m)
}

// Get returns the Value corresponding to the key in the ArrayMap
func (m ArrayMap) Get(key Value) Value {
	l := len(m)
	for i := 0; i < l; i++ {
		mp := m[i]
		if mp[0] == key {
			return mp[1]
		}
	}
	return Nil
}

// Apply makes ArrayMap applicable
func (m ArrayMap) Apply(c Context, args Sequence) Value {
	AssertArity(args, 1)
	key := Eval(c, args.First())
	return m.Get(key)
}

// Eval makes an ArrayMap Evaluable
func (m ArrayMap) Eval(c Context) Value {
	l := len(m)
	r := make(ArrayMap, l)
	for i := 0; i < l; i++ {
		mp := m[i]
		r[i] = Vector{
			Eval(c, mp[0]),
			Eval(c, mp[1]),
		}
	}
	return r
}

// First returns the first key/value pair of an ArrayMap
func (m ArrayMap) First() Value {
	return m[0]
}

// Rest returns the remaining elements of an ArrayMap as a Sequence
func (m ArrayMap) Rest() Sequence {
	return m[1:]
}

// Prepend creates a new Sequence by prepending a Value
func (m ArrayMap) Prepend(v Value) Sequence {
	if mp, ok := v.(Vector); ok {
		AssertArity(mp, 2)
		return append(ArrayMap{mp}, m...)
	}
	panic(ExpectedPair)
}

// IsSequence returns whether this instance is a consumable Sequence
func (m ArrayMap) IsSequence() bool {
	return len(m) > 0
}

func (m ArrayMap) String() string {
	var b bytes.Buffer
	l := len(m)

	b.WriteString("{")
	for i := 0; i < l; i++ {
		if i > 0 {
			b.WriteString(", ")
		}
		mp := m[i]
		b.WriteString(String(mp[0]))
		b.WriteString(" ")
		b.WriteString(String(mp[1]))
	}
	b.WriteString("}")
	return b.String()
}
