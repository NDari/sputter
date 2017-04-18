package api

// ValueMapper returns a mapped representation of the specified Value
type ValueMapper func(Value) Value

// ValueFilter returns true if the Value remains part of a filtered Sequence
type ValueFilter func(Value) bool

type lazyElement struct {
	first Value
	rest  Sequence
}

// First returns the first mapped Value from the lazyElement
func (l *lazyElement) First() Value {
	return l.first
}

// Rest returns the rest of the lazyElement
func (l *lazyElement) Rest() Sequence {
	return l.rest
}

// IsSequence returns whether this instance is a consumable Sequence
func (l *lazyElement) IsSequence() bool {
	return true
}

func (l *lazyElement) Prepend(v Value) Sequence {
	return &lazyElement{
		first: v,
		rest:  l,
	}
}

type lazyMap struct {
	first    Value
	rest     Sequence
	isSeq    bool
	sequence Sequence
	mapper   ValueMapper
}

// Map creates a new lazy mapping Sequence that wraps the
// original and only maps its Values on demand
func Map(s Sequence, m ValueMapper) Sequence {
	return &lazyMap{
		rest:     EmptyList,
		sequence: s,
		mapper:   m,
	}
}

func (l *lazyMap) resolve() *lazyMap {
	if l.sequence == nil {
		return l
	}

	if l.sequence.IsSequence() {
		l.isSeq = true
		l.first = l.mapper(l.sequence.First())
		l.rest = &lazyMap{
			rest:     EmptyList,
			sequence: l.sequence.Rest(),
			mapper:   l.mapper,
		}
	}
	l.sequence = nil
	l.mapper = nil
	return l
}

// First returns the first mapped Value from the lazyMap
func (l *lazyMap) First() Value {
	return l.resolve().first
}

// Rest returns the rest of the lazyMap
func (l *lazyMap) Rest() Sequence {
	return l.resolve().rest
}

// IsSequence returns whether this instance is a consumable Sequence
func (l *lazyMap) IsSequence() bool {
	return l.resolve().isSeq
}

func (l *lazyMap) Prepend(v Value) Sequence {
	return &lazyElement{
		first: v,
		rest:  l,
	}
}

type lazyFilter struct {
	first    Value
	rest     Sequence
	isSeq    bool
	sequence Sequence
	filter   ValueFilter
}

// Filter creates a new lazy filter Sequence that wraps the
// original and only filters the next Value(s) on demand
func Filter(s Sequence, f ValueFilter) Sequence {
	return &lazyFilter{
		rest:     EmptyList,
		sequence: s,
		filter:   f,
	}
}

func (l *lazyFilter) resolve() *lazyFilter {
	if l.sequence == nil {
		return l
	}

	for i := l.sequence; i.IsSequence(); i = i.Rest() {
		if v := i.First(); l.filter(v) {
			l.isSeq = true
			l.first = v
			l.rest = &lazyFilter{
				rest:     EmptyList,
				sequence: i.Rest(),
				filter:   l.filter,
			}
			break
		}
	}
	l.sequence = nil
	l.filter = nil
	return l
}

// First returns the first mapped Value from the lazyFilter
func (l *lazyFilter) First() Value {
	return l.resolve().first
}

// Rest returns the rest of the lazyFilter
func (l *lazyFilter) Rest() Sequence {
	return l.resolve().rest
}

// IsSequence returns whether this instance is a consumable Sequence
func (l *lazyFilter) IsSequence() bool {
	return l.resolve().isSeq
}

func (l *lazyFilter) Prepend(v Value) Sequence {
	return &lazyElement{
		first: v,
		rest:  l,
	}
}

type lazyConcat struct {
	first    Value
	rest     Sequence
	isSeq    bool
	sequence Sequence
}

// Concat creates a new sequence based on the content of
// several provided Sequences. The results are computed lazily
func Concat(s Sequence) Sequence {
	return &lazyConcat{
		rest:     EmptyList,
		sequence: s,
	}
}

func (l *lazyConcat) resolve() *lazyConcat {
	if l.sequence == nil {
		return l
	}

	for i := l.sequence; i.IsSequence(); i = i.Rest() {
		if f := AssertSequence(i.First()); f.IsSequence() {
			l.first = f.First()
			l.rest = &lazyConcat{
				rest:     EmptyList,
				sequence: i.Rest().Prepend(f.Rest()),
			}
			l.isSeq = true
			l.sequence = nil
			return l
		}
	}

	l.sequence = nil
	return l
}

func (l *lazyConcat) First() Value {
	return l.resolve().first
}

func (l *lazyConcat) Rest() Sequence {
	return l.resolve().rest
}

func (l *lazyConcat) IsSequence() bool {
	return l.resolve().isSeq
}

func (l *lazyConcat) Prepend(v Value) Sequence {
	return &lazyElement{
		first: v,
		rest:  l,
	}
}

type lazyTake struct {
	first    Value
	rest     Sequence
	isSeq    bool
	sequence Sequence
	count    int
}

// Take creates a new Sequence based on the first elements of the source
// sequence. The result is computed lazily
func Take(s Sequence, count int) Sequence {
	return &lazyTake{
		rest:     EmptyList,
		sequence: s,
		count:    count,
	}
}

func (l *lazyTake) resolve() *lazyTake {
	if l.sequence == nil {
		return l
	}

	s := l.sequence
	if l.count > 0 && s.IsSequence() {
		l.isSeq = true
		l.first = s.First()
		l.rest = &lazyTake{
			rest:     EmptyList,
			sequence: s.Rest(),
			count:    l.count - 1,
		}
	}
	l.sequence = nil
	return l
}

func (l *lazyTake) First() Value {
	return l.resolve().first
}

func (l *lazyTake) Rest() Sequence {
	return l.resolve().rest
}

func (l *lazyTake) IsSequence() bool {
	return l.resolve().isSeq
}

func (l *lazyTake) Prepend(v Value) Sequence {
	return &lazyElement{
		first: v,
		rest:  l,
	}
}

type lazyDrop struct {
	first    Value
	rest     Sequence
	isSeq    bool
	sequence Sequence
	count    int
}

// Drop creates a new Sequence based on dropping the first elements of
// the source sequence. The result is computed lazily
func Drop(s Sequence, count int) Sequence {
	return &lazyDrop{
		rest:     EmptyList,
		sequence: s,
		count:    count,
	}
}

func (l *lazyDrop) resolve() *lazyDrop {
	if l.sequence == nil {
		return l
	}

	i := l.sequence
	for c := l.count; c > 0; c-- {
		if !i.IsSequence() {
			l.sequence = nil
			return l
		}
		i = i.Rest()
	}

	l.isSeq = i.IsSequence()
	l.first = i.First()
	l.rest = i.Rest()
	l.sequence = nil
	return l
}

func (l *lazyDrop) First() Value {
	return l.resolve().first
}

func (l *lazyDrop) Rest() Sequence {
	return l.resolve().rest
}

func (l *lazyDrop) IsSequence() bool {
	return l.resolve().isSeq
}

func (l *lazyDrop) Prepend(v Value) Sequence {
	return &lazyElement{
		first: v,
		rest:  l,
	}
}

// Reduce performs a reduce operation over a Sequence, starting with the
// first two elements of that sequence.
func Reduce(c Context, s Sequence, a Applicable) Value {
	AssertMinimumArity(s, 2)
	r := s.First()
	for i := s.Rest(); i.IsSequence(); i = i.Rest() {
		r = a.Apply(c, Vector{r, i.First()})
	}
	return r
}
