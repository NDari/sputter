package parser

import (
	"regexp"

	a "github.com/kode4food/sputter/api"
)

// UnmatchedState is the error returned when the lexer state is invalid
const UnmatchedState = "unmatched lexing state"

// TokenType is an opaque type for lexer tokens
type TokenType int

// Token Types
const (
	Error TokenType = iota
	Identifier
	String
	Number
	Ratio
	ListStart
	ListEnd
	VectorStart
	VectorEnd
	MapStart
	MapEnd
	QuoteMarker
	Whitespace
	Comment
	endOfFile
)

type lexer struct {
	once  a.Do
	src   a.Str
	isSeq bool
	first a.Value
	rest  a.Sequence
}

// Token is a lexer value
type Token struct {
	Type  TokenType
	Value a.Value
}

type tokenizer func(s a.Str) *Token

type matchEntry struct {
	pattern  *regexp.Regexp
	function tokenizer
}

type mactchEntries []matchEntry

var (
	escaped    = regexp.MustCompile(`\\\\|\\"|\\[^\\"]`)
	escapedMap = map[string]string{
		`\\`: `\`,
		`\"`: `"`,
	}
	matchers mactchEntries
)

// NewLexer creates a new lexer Sequence
func NewLexer(src a.Str) a.Sequence {
	l := &lexer{
		once: a.Once(),
		src:  src,
	}

	return a.Filter(l, func(v a.Value) bool {
		return isNotWhitespace(v.(*Token))
	})
}

func (l *lexer) resolve() *lexer {
	l.once(func() {
		t, rsrc := l.matchToken()

		if t.Type != endOfFile {
			l.isSeq = true
			l.first = t
		}
		l.rest = &lexer{
			once: a.Once(),
			rest: a.EmptyList,
			src:  rsrc,
		}
	})
	return l
}

func (l *lexer) matchToken() (*Token, a.Str) {
	src := string(l.src)
	for _, s := range matchers {
		if i := s.pattern.FindStringIndex(src); i != nil {
			f := src[:i[1]]
			r := src[len(f):]
			return s.function(a.Str(f)), a.Str(r)
		}
	}
	// Shouldn't happen because of the patterns that are defined,
	// but is here as a safety net
	panic(UnmatchedState)
}

func (l *lexer) IsSequence() bool {
	return l.resolve().isSeq
}

func (l *lexer) First() a.Value {
	return l.resolve().first
}

func (l *lexer) Rest() a.Sequence {
	return l.resolve().rest
}

func (l *lexer) Prepend(v a.Value) a.Sequence {
	// insulated by a filter
	panic("not implemented")
}

func (l *lexer) Str() a.Str {
	// insulated by a filter
	panic("not implemented")
}

// Str converts this Value into a Str
func (t *Token) Str() a.Str {
	return a.NewVector(a.NewFloat(float64(t.Type)), t.Value).Str()
}

func isNotWhitespace(t *Token) bool {
	return t.Type != Whitespace && t.Type != Comment
}

func makeToken(t TokenType, v a.Value) *Token {
	return &Token{
		Type:  t,
		Value: v,
	}
}

func tokenState(t TokenType) tokenizer {
	return func(s a.Str) *Token {
		return makeToken(t, a.Str(s))
	}
}

func endState(t TokenType) tokenizer {
	return func(_ a.Str) *Token {
		return makeToken(t, a.Nil)
	}
}

func unescape(s a.Str) a.Str {
	r := escaped.ReplaceAllStringFunc(string(s), func(e string) string {
		if r, ok := escapedMap[e]; ok {
			return r
		}
		return e
	})
	return a.Str(r)
}

func stringState(r a.Str) *Token {
	s := unescape(r[1 : len(r)-1])
	return makeToken(String, a.Str(s))
}

func ratioState(s a.Str) *Token {
	v := a.ParseNumber(s)
	return makeToken(Ratio, v)
}

func numberState(s a.Str) *Token {
	v := a.ParseNumber(s)
	return makeToken(Number, v)
}

func init() {
	pattern := func(p string, s tokenizer) matchEntry {
		return matchEntry{
			pattern:  regexp.MustCompile(p),
			function: s,
		}
	}

	matchers = mactchEntries{
		pattern(`^$`, endState(endOfFile)),
		pattern(`^;[^\n]*[\n]`, tokenState(Comment)),
		pattern(`^[\s,]+`, tokenState(Whitespace)),
		pattern(`^\(`, tokenState(ListStart)),
		pattern(`^\[`, tokenState(VectorStart)),
		pattern(`^{`, tokenState(MapStart)),
		pattern(`^\)`, tokenState(ListEnd)),
		pattern(`^]`, tokenState(VectorEnd)),
		pattern(`^}`, tokenState(MapEnd)),
		pattern(`^'`, tokenState(QuoteMarker)),

		pattern(`^"(\\\\|\\"|\\[^\\"]|[^"\\])*"`, stringState),

		pattern(`^[+-]?[1-9]\d*/[1-9]\d*`, ratioState),
		pattern(`^[+-]?((0|[1-9]\d*)(\.\d+)?([eE][+-]?\d+)?)`, numberState),

		pattern(`^[^(){}\[\]\s,'";]+`, tokenState(Identifier)),

		pattern(`^.`, endState(Error)),
	}
}
