package expensebook

import "testing"

func TestMoneyString(t *testing.T) {
	for _, test := range []struct {
		value Money
		want  string
	}{
		{value: 0, want: "0.00"},
		{value: 1250, want: "12.50"},
		{value: 7, want: "0.07"},
	} {
		if got := test.value.String(); got != test.want {
			t.Errorf("Money(%d).String() = %q, want %q", test.value, got, test.want)
		}
	}
}
