package api_test

import (
	"testing"

	a "github.com/kode4food/sputter/api"
	"github.com/stretchr/testify/assert"
)

func TestLazyMapper(t *testing.T) {
	as := assert.New(t)

	c := 0
	l := a.NewList("last").Prepend("middle").Prepend("first")
	w := a.NewLazyMapper(l, func(v a.Value) a.Value {
		c++
		return "this is the " + v.(string)
	})

	as.Equal(0, c, "nothing has been processed")

	v1 := w.First()
	as.Equal(1, c, "first element has been processed")
	as.Equal("this is the first", v1, "first element mapped")

	v2 := w.Rest().First()
	as.Equal(2, c, "second element has been processed")
	as.Equal("this is the middle", v2, "second element mapped")

	v3 := w.Rest().Rest().First()
	as.Equal(3, c, "third element has been processed")
	as.Equal("this is the last", v3, "third element mapped")

	r1 := w.Rest().Rest().Rest()
	as.False(r1.IsSequence(), "finished")

	p := w.Prepend("not mapped")
	v4 := p.First()
	r2 := p.Rest()
	as.Equal(3, c, "prepend doesn't trigger mapper")
	as.Equal("not mapped", v4, "prepended element retrieved")
	as.Equal(w, r2, "prepended rest is the original sequence")
}

func TestLazyFilter(t *testing.T) {
	as := assert.New(t)

	c := 0
	l := a.NewList("last").Prepend("filtered out").Prepend("first")
	w := a.NewLazyFilter(l, func(v a.Value) bool {
		c++
		return v.(string) != "filtered out"
	})

	as.Equal(0, c, "nothing has been processed")

	v1 := w.First()
	as.Equal(1, c, "first element has been processed")
	as.Equal("first", v1, "first element passed")

	v2 := w.Rest().First()
	as.Equal(3, c, "second element has been skipped")
	as.Equal("last", v2, "third element passed")

	r1 := w.Rest().Rest()
	as.False(r1.IsSequence(), "finished")

	p := w.Prepend("filtered out")
	v4 := p.First()
	r2 := p.Rest()
	as.Equal(3, c, "prepend doesn't trigger filter")
	as.Equal("filtered out", v4, "prepended element retrieved")
	as.Equal(w, r2, "prepended rest is the original sequence")
}

func TestLazyFilteredAndMapped(t *testing.T) {
	as := assert.New(t)

	c1 := 0
	c2 := 0
	l := a.NewList("last").Prepend("middle").Prepend("first")
	w1 := a.NewLazyFilter(l, func(v a.Value) bool {
		c1++
		return v.(string) != "middle"
	})
	w2 := a.NewLazyMapper(w1, func(v a.Value) a.Value {
		c2++
		return "this is the " + v.(string)
	})

	as.Equal(0, c1, "nothing has been processed")
	as.Equal(0, c2, "nothing has been processed")

	v1 := w2.First()
	as.Equal(1, c1, "first element has been processed")
	as.Equal(1, c2, "first element has been processed")
	as.Equal("this is the first", v1, "first element mapped")

	v2 := w2.Rest().First()
	as.Equal(3, c1, "second element has been skipped")
	as.Equal(2, c2, "second element is first mapper sees")
	as.Equal("this is the last", v2, "last element mapped")

	r1 := w2.Rest().Rest()
	as.False(r1.IsSequence(), "finished")
}
