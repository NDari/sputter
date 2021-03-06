package builtins

import (
	a "github.com/kode4food/sputter/api"
	e "github.com/kode4food/sputter/evaluator"
)

const (
	errorName   = "error*"
	raiseName   = "raise"
	recoverName = "recover"
	doName      = "do"
	readName    = "read"
	evalName    = "eval"
)

type (
	errorFunction   struct{ BaseBuiltIn }
	raiseFunction   struct{ BaseBuiltIn }
	recoverFunction struct{ BaseBuiltIn }
	doFunction      struct{ BaseBuiltIn }
	readFunction    struct{ BaseBuiltIn }
	evalFunction    struct{ BaseBuiltIn }
)

func (*errorFunction) Apply(_ a.Context, args a.Values) a.Value {
	a.AssertArity(args, 1)
	return a.Err(toProperties(args[0].(a.Associative)))
}

func (*raiseFunction) Apply(_ a.Context, args a.Values) a.Value {
	a.AssertArity(args, 1)
	panic(args[0])
}

func (*recoverFunction) Apply(c a.Context, args a.Values) (res a.Value) {
	a.AssertArity(args, 2)

	defer func() {
		if rec := recover(); rec != nil {
			post := a.Eval(c, args[1]).(a.Applicable)
			res = post.Apply(c, a.Values{rec.(a.Value)})
		}
	}()

	return a.Eval(c, args[0])
}

func (*doFunction) Apply(c a.Context, args a.Values) a.Value {
	var res a.Value = a.Nil
	for _, f := range args {
		res = a.Eval(c, f)
	}
	return res
}

func (*readFunction) Apply(c a.Context, args a.Values) a.Value {
	a.AssertArity(args, 1)
	v := args[0]
	s := v.(a.Sequence)
	return e.ReadStr(c, a.SequenceToStr(s))
}

func (*evalFunction) Apply(c a.Context, args a.Values) a.Value {
	a.AssertArity(args, 1)
	v := args[0]
	return a.Eval(c, v)
}

func init() {
	var _error *errorFunction
	var raise *raiseFunction
	var _recover *recoverFunction
	var do *doFunction
	var read *readFunction
	var eval *evalFunction

	RegisterBuiltIn(errorName, _error)
	RegisterBuiltIn(raiseName, raise)
	RegisterBuiltIn(recoverName, _recover)
	RegisterBuiltIn(doName, do)
	RegisterBuiltIn(readName, read)
	RegisterBuiltIn(evalName, eval)
}
