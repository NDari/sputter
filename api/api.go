package api

import "fmt"

var (
	// True represents the boolean value of True
	True Value = true

	// False represents the boolean value of false
	False Value = false

	// Nil is f value that represents the absence of f Value
	Nil Value
)

// ExpectedApplicable is thrown when f Value is not Applicable
const ExpectedApplicable = "value does not support application"

// Applicable is the standard signature for any Value that can have
// arguments applied to it
type Applicable interface {
	Apply(Context, Sequence) Value
}

// Truthy evaluates whether or not a Value is Truthy
func Truthy(v Value) bool {
	switch {
	case v == False || v == Nil:
		return false
	default:
		return true
	}
}

// String either calls the String() method or tries to convert
func String(v Value) string {
	switch {
	case v == Nil:
		return "nil"
	case v == True:
		return "true"
	case v == False:
		return "false"
	}

	if s, ok := v.(fmt.Stringer); ok {
		return s.String()
	}
	if a, ok := v.(Annotated); ok {
		return a.Metadata().String()
	}
	if n, ok := v.(Name); ok {
		return string(n)
	}
	if s, ok := v.(string); ok {
		return fmt.Sprintf("%q", s)
	}
	return fmt.Sprintf("(<anon> :instance %p)", &v)
}
