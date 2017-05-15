package builtins_test

import (
	"testing"

	a "github.com/kode4food/sputter/api"
)

func TestList(t *testing.T) {
	testCode(t, `(list? '(1 2 3))`, a.True)
	testCode(t, `(list? ())`, a.True)
	testCode(t, `(list? [1 2 3])`, a.False)
	testCode(t, `(list? 42)`, a.False)
	testCode(t, `(list? (list 1 2 3))`, a.True)
	testCode(t, `(list)`, a.EmptyList)
	testCode(t, `(list? (to-list (vector 1 2 3)))`, a.True)

	testCode(t, `
		(first (apply list (map (fn [x] (* x 2)) [1 2 3 4])))
	`, f(2))

	a.GetNamespace(a.UserDomain).Delete("x")
	testCode(t, `
		(def x '(1 2 3 4))
		(x 2)
	`, f(3))
}
