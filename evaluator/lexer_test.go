package evaluator_test

import (
	"testing"

	a "github.com/kode4food/sputter/api"
	"github.com/kode4food/sputter/assert"
	e "github.com/kode4food/sputter/evaluator"
)

func makeToken(t e.TokenType, v a.Value) *e.Token {
	return &e.Token{
		Type:  t,
		Value: v,
	}
}

func assertToken(as *assert.Wrapper, like *e.Token, value *e.Token) {
	as.Number(float64(like.Type), float64(value.Type))
}

func assertTokenSequence(as *assert.Wrapper, s a.Sequence, tokens []*e.Token) {
	iter := a.Iterate(s)
	for _, l := range tokens {
		v, ok := iter.Next()
		as.True(ok)
		assertToken(as, l, v.(*e.Token))
	}
	v, ok := iter.Next()
	as.False(ok)
	as.Nil(v)
}

func TestCreateLexer(t *testing.T) {
	as := assert.New(t)
	l := e.Scan("hello")
	as.NotNil(l)
	as.String(`([1 "hello"])`, a.MakeSequenceStr(l))
}

func TestWhitespace(t *testing.T) {
	as := assert.New(t)
	l := e.Scan("   \t ")
	assertTokenSequence(as, l, []*e.Token{})
}

func TestEmptyList(t *testing.T) {
	as := assert.New(t)
	l := e.Scan(" ( \t ) ")
	assertTokenSequence(as, l, []*e.Token{
		makeToken(e.ListStart, s("(")),
		makeToken(e.ListEnd, s(")")),
	})
}

func TestNumbers(t *testing.T) {
	as := assert.New(t)
	l := e.Scan(" 10 12.8 8E+10 99.598e+10 54e+12 1/2")
	assertTokenSequence(as, l, []*e.Token{
		makeToken(e.Number, f(10)),
		makeToken(e.Number, f(12.8)),
		makeToken(e.Number, f(8E+10)),
		makeToken(e.Number, f(99.598e+10)),
		makeToken(e.Number, f(54e+12)),
		makeToken(e.Ratio, a.NewRatio(1, 2)),
	})
}

func TestStrings(t *testing.T) {
	as := assert.New(t)
	l := e.Scan(` "hello there" "how's \"life\"?"  `)
	assertTokenSequence(as, l, []*e.Token{
		makeToken(e.String, s(`hello there`)),
		makeToken(e.String, s(`how's "life"?`)),
	})
}

func TestMultiLine(t *testing.T) {
	as := assert.New(t)
	l := e.Scan(` "hello there"
  "how's life?"
99`)

	assertTokenSequence(as, l, []*e.Token{
		makeToken(e.String, s(`hello there`)),
		makeToken(e.String, s(`how's life?`)),
		makeToken(e.Number, f(99)),
	})
}

func TestComments(t *testing.T) {
	as := assert.New(t)
	l := e.Scan(`"hello" ; (this is commented)`)
	assertTokenSequence(as, l, []*e.Token{
		makeToken(e.String, s(`hello`)),
	})
}
