package exp

import "strconv"

// Eq

type expEq struct {
	key   string
	value float64
}

func (eq expEq) Eval(p Params) bool {
	value, err := strconv.ParseFloat(p.Get(eq.key), 64)
	if err != nil {
		return false
	}
	return value == eq.value
}

func (eq expEq) String() string {
	return sprintf("[%s==%.2f]", eq.key, eq.value)
}

// Equal evaluates to true if the value pointed to by key is equal in value to
// v. The value pointed to by k is parsed into a float64 before comparing. If a
// parse error occurs false is returned.
func Equal(k string, v float64) Exp {
	return expEq{k, v}
}

// Eq is an alias for Equal.
func Eq(k string, v float64) Exp {
	return Equal(k, v)
}

// NotEqual is a shorthand for Not(Eq(k, v)).
func NotEqual(k string, v float64) Exp {
	return Neq(k, v)
}

// Neq is an alias for NotEqual.
func Neq(k string, v float64) Exp {
	return Not(Eq(k, v))
}

// Gte
type expGte struct {
	key   string
	value float64
}

func (gte expGte) Eval(p Params) bool {
	value, err := strconv.ParseFloat(p.Get(gte.key), 64)
	if err != nil {
		return false
	}
	return value >= gte.value
}

func (gte expGte) String() string {
	return sprintf("[%s≥%.2f]", gte.key, gte.value)
}

// GreaterThanOrEqual evaluates to true if the value pointed to by key is greater or equal to
// the value of v. The value is parsed as float before performing the comparison.
func GreaterThanOrEqual(k string, v float64) Exp {
	return expGte{k, v}
}

// Gte is an alias for GreaterOrEqual.
func Gte(k string, v float64) Exp {
	return GreaterThanOrEqual(k, v)
}

// Gt

type expGt struct {
	key   string
	value float64
}

func (gt expGt) Eval(p Params) bool {
	value, err := strconv.ParseFloat(p.Get(gt.key), 64)
	if err != nil {
		return false
	}
	return value > gt.value
}

func (gt expGt) String() string {
	return sprintf("[%s>%.2f]", gt.key, gt.value)
}


// GreaterThan evaluates to true if the value pointed to by key is greater in
// value than v. The value is parsed as float before performing the comparison.
func GreaterThan(k string, v float64) Exp {
	return expGt{k, v}
}

// Gt is an alias for GreaterThan.
func Gt(k string, v float64) Exp {
	return GreaterThan(k, v)
}

// Lte
type expLte struct {
	key   string
	value float64
}

func (lte expLte) Eval(p Params) bool {
	value, err := strconv.ParseFloat(p.Get(lte.key), 64)
	if err != nil {
		return false
	}
	return value <= lte.value
}

func (lte expLte) String() string {
	return sprintf("[%s≤%.2f]", lte.key, lte.value)
}

// LessThanOrEqual evaluates to true if the value pointed to by key is less or equal to
// the value of v. The value is parsed as float before performing the comparison.
func LessThanOrEqual(k string, v float64) Exp {
	return expLte{k, v}
}

// Lte is an alias for LessThanOrEqual.
func Lte(k string, v float64) Exp {
	return LessThanOrEqual(k, v)
}


// Lt

type expLt struct {
	key   string
	value float64
}

func (lt expLt) Eval(p Params) bool {
	value, err := strconv.ParseFloat(p.Get(lt.key), 64)
	if err != nil {
		return false
	}
	return value < lt.value
}

func (lt expLt) String() string {
	return sprintf("[%s<%.2f]", lt.key, lt.value)
}

// LessThan evaluates to true if the value pointed to by key is less in value
// than v. The value is parsed as float before performing the comparison.
func LessThan(k string, v float64) Exp {
	return expLt{k, v}
}

// Lt is an alias for LessThan.
func Lt(k string, v float64) Exp {
	return LessThan(k, v)
}

