package exp

import "testing"

func TestBool(t *testing.T) {
	var p = Map{
		"foo": "true",
		"bar": "false",
	}
	for _, test := range []struct {
		exp Exp
		out bool
	}{
		{Boolean("foo", true), true},
		{Boolean("foo", false), false},
		{Boolean("bar", true), false},
		{Boolean("bar", false), true},
	} {
		if test.exp.Eval(p) != test.out {
			t.Error("unexpected output")
		}
	}
}
