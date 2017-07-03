package api

import "bytes"

// ExpectedVector is raised if a value is not a Vector
const ExpectedVector = "value is not a vector: %s"

// Vector is a fixed-length Array of Values
type Vector interface {
	Conjoiner
	Indexed
	Counted
	Applicable
	Evaluable
	IsVector() bool
}

type vector []Value

var emptyVector = vector{}

// NewVector instantiates a new Vector
func NewVector(v ...Value) Vector {
	return vector(v)
}

func (v vector) IsVector() bool {
	return true
}

func (v vector) Count() int {
	return len(v)
}

func (v vector) ElementAt(index int) (Value, bool) {
	if index >= 0 && index < len(v) {
		return v[index], true
	}
	return Nil, false
}

func (v vector) Apply(_ Context, args Sequence) Value {
	return IndexedApply(v, args)
}

func (v vector) Eval(c Context) Value {
	l := len(v)
	r := make(vector, l)
	for i := 0; i < l; i++ {
		r[i] = Eval(c, v[i])
	}
	return r
}

func (v vector) First() Value {
	if len(v) > 0 {
		return v[0]
	}
	return Nil
}

func (v vector) Rest() Sequence {
	if len(v) > 1 {
		return Sequence(v[1:])
	}
	return emptyVector
}

func (v vector) Prepend(p Value) Sequence {
	return append(vector{p}, v...)
}

func (v vector) Conjoin(a Value) Sequence {
	return append(v, a)
}

func (v vector) IsSequence() bool {
	return len(v) > 0
}

func (v vector) Str() Str {
	var b bytes.Buffer
	l := len(v)

	b.WriteString("[")
	for i := 0; i < l; i++ {
		if i > 0 {
			b.WriteString(" ")
		}
		b.WriteString(string(v[i].Str()))
	}
	b.WriteString("]")
	return Str(b.String())
}

// AssertVector will cast the Value into a Vector or die trying
func AssertVector(v Value) Vector {
	if r, ok := v.(Vector); ok {
		return r
	}
	panic(Err(ExpectedVector, v))
}
