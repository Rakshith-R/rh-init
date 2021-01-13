package submodule

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct{ in, out string }{
		{"abc", "cba"},
		{"a", "a"},
	}
	for _, test := range tests {
		if test.out != Reverse(test.in) {
			t.Errorf("Wrong ans %v, for input %v\n", test.out, test.in)
		} else {
			t.Logf("Correct ans %v, for input %v\n", test.out, test.in)
		}
	}
}

func TestReverse2(t *testing.T) {
	tests := []struct{ in, out string }{
		{"abc", "cba"},
		{"a", "a"},
	}
	for _, test := range tests {
		if test.out == Reverse(test.in) {
			t.Errorf("Wrong ans %v, for input %v\n", test.out, test.in)
		} else {
			t.Logf("Correct ans %v, for input %v\n", test.out, test.in)
		}
	}
}
