package builtins

import (
	a "github.com/kode4food/sputter/api"
	d "github.com/kode4food/sputter/docstring"
	e "github.com/kode4food/sputter/evaluator"
)

// Namespace is a special Namespace for built-in identifiers
var Namespace = a.GetNamespace(a.BuiltInDomain)

func isBuiltInDomain(s a.Symbol) bool {
	return s.Domain() == a.BuiltInDomain
}

func isBuiltInCall(n a.Name, v a.Value) (a.List, bool) {
	if l, ok := v.(a.List); ok && l.Count() > 0 {
		if s, ok := l.First().(a.Symbol); ok {
			return l, isBuiltInDomain(s) && s.Name() == n
		}
	}
	return nil, false
}

func registerAnnotated(v a.Annotated) {
	n, _ := v.Metadata().Get(a.MetaName)
	Namespace.Put(n.(a.Name), v.(a.Value))
}

func do(c a.Context, args a.Sequence) a.Value {
	return a.EvalBlock(c, args)
}

func read(c a.Context, args a.Sequence) a.Value {
	a.AssertArity(args, 1)
	v := args.First()
	s := a.AssertSequence(v)
	return e.ReadStr(c, a.ToStr(s))
}

func eval(c a.Context, args a.Sequence) a.Value {
	a.AssertArity(args, 1)
	v := args.First()
	return a.Eval(c, v)
}

func init() {
	registerAnnotated(
		a.NewFunction(do).WithMetadata(a.Properties{
			a.MetaName:    a.Name("do"),
			a.MetaDoc:     d.Get("do"),
			a.MetaSpecial: a.True,
		}),
	)

	registerAnnotated(
		a.NewFunction(read).WithMetadata(a.Properties{
			a.MetaName: a.Name("read"),
			a.MetaDoc:  d.Get("read"),
		}),
	)

	registerAnnotated(
		a.NewFunction(eval).WithMetadata(a.Properties{
			a.MetaName: a.Name("eval"),
			a.MetaDoc:  d.Get("eval"),
		}),
	)
}
