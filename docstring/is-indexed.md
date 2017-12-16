# (indexed? form+) tests whether the provided forms are indexed sequences
If all forms evaluate to a valid sequence than can be accessed by index, then this function will return _true_. The first non-indexed sequence will result in the function returning _false_.

## An Example

  (indexed? '(1 2 3 4) [5 6 7 8])

This example will return _true_.

Like most predicates, this function can also be negated by prepending the /!/ character. This means that all of the provided forms must not be valid indexed sequences.

  (!indexed? "hello" 99)

This example will return _false_ because strings are indexed sequences.
