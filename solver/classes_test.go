package solver

import "testing"

func TestIsoString(t *testing.T) {
	for _, tc := range []struct {
		in    string
		want  string
		first string
		next  string
	}{
		{"CAT", "C,AT", "C", "AT"},
		{"PIPE", "PP,EI", "P", "P,EI"},
		{"QACK", "KQ,C,A", "K", "Q,C,A"},
		{"X", "X", "X", ""},
		{"", "", "", ""},
	} {
		i := ToIsochars(tc.in)
		if got := i.String(); got != tc.want {
			t.Errorf("ToIsochars(%q) = %v, want %v", tc.in, got, tc.want)
		}
		if got := i.First(); got != tc.first {
			t.Errorf("ToIsochars(%q).First() = %v, want %v", tc.in, got, tc.first)
		}
		if got := i.Next().String(); got != tc.next {
			t.Errorf("ToIsochars(%q).Next() = %v, want %v", tc.in, got, tc.next)
		}
	}
}
