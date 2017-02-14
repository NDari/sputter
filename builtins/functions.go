package builtins

import a "github.com/kode4food/sputter/api"

func define(name a.Name, argNames a.Iterable, body a.Iterable) *a.Function {
	ac := argCount(argNames)

	return &a.Function{
		Name: name,
		Exec: func(c *a.Context, args a.Iterable) a.Value {
			AssertArity(args, ac)
			l := c.Child()
			anIter := argNames.Iterate()
			aIter := args.Iterate()
			for ns, nok := anIter.Next(); nok; {
				an := ns.(*a.Symbol).Name
				av, aok := aIter.Next()
				if aok {
					l.Put(an, av)
				} else {
					l.Put(an, a.EmptyList)
				}
				ns, nok = anIter.Next()
			}
			return a.EvaluateIterator(l, body.Iterate())
		},
	}
}

func defunCommand(c *a.Context, args a.Iterable) a.Value {
	AssertMinimumArity(args, 3)
	g := c.Globals()
	i := args.Iterate()

	fv, _ := i.Next()
	fn := fv.(*a.Symbol).Name

	av, _ := i.Next()
	an := av.(a.Iterable)

	b := i.Iterable()

	d := define(fn, an, b)
	g.PutFunction(d)
	return d
}

func init() {
	BuiltIns.PutFunction(&a.Function{Name: "defun", Exec: defunCommand})
}
