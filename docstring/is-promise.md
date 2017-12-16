# (promise? form+) tests whether the provided forms are promises
If all forms evaluate to a promise, then this function will return _true_. The first non-promise will result in the function returning _false_.

## An Example

  (def p1 (promise))
  (def p2 (promise))
  (promise? p1 p2 [1 2 3])

This example will return _false_ because the third form is a vector.

Like most predicates, this function can also be negated by prepending the /!/ character. This means that all of the provided forms must not be promises.

  (!promise? "hello" [1 2 3])

This example will return _true_.
