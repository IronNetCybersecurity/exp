package exp

import "strconv"

// Boolean

// Bool is a wrapper for the native bool type which provides an Eval function so
// that it satisfies the Exp interface.
type Bool bool

// Eval will always return the boolean value of b and disregard p.
func (b Bool) Eval(p Params) bool {
	return bool(b)
}

func (b Bool) String() string {
	if bool(b) {
		return "T"
	}
	return "F"
}

var (
	// True is an expression that always evaluates to true.
	True = Bool(true)
	// False is an expression that always evaluates to false.
	False = Bool(false)
)

// Boolean exp that actually read the stupid field value and matches
type expBool struct {
	key   string
	value bool
}

func (eq expBool) Eval(p Params) bool {
	value, err := strconv.ParseBool(p.Get(eq.key))
	if err != nil {
		return false
	}

	return value == eq.value
}

func (eq expBool) String() string {
	if eq.value {
		return "T"
	}
	return "F"
}

// Boolean builds a Boolean expression
func Boolean(k string, v bool) Exp {
	return expBool{k, v}
}
