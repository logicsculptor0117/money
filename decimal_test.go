package money

import (
	"fmt"
	"testing"
)

func ExampleNew() {
	fmt.Print(New(50))
	// Output: 50.00
}

func ExampleNewCents() {
	fmt.Print(NewCents(1234))
	// Output: 12.34
}

func ExampleNewScalar() {
	fmt.Print(NewScalar(42, 0))
	//Output: 42
}

func ExamplePc() {
	fmt.Print(Pc(50))
	// Output: 0.5
}

func ExamplePm() {
	fmt.Print(Pm(20))
	// Output: 0.02
}

func ExampleBp() {
	fmt.Print(Bp(10))
	// Output: 0.001
}

func ExampleDecimal_Equals() {
	a := New(1)
	b := NewCents(100)
	fmt.Print(a.Equals(b))
	// Output: true
}

func ExampleDecimal_Format() {
	print := func(format string, d Decimal) {
		fmt.Printf(format+"\n", d)
	}

	print("%s", NewCents(123))
	print("%s", NewScalar(2345, -6))
	print("%v", NewScalar(2345, -6))
	print("%.2f", New(5678))
	print("`%5.2f`", New(7).Div(New(3)))
	print("'%-10.f'", NewCents(-80808))

	// Output:
	// 1.23
	// 2.345e+9
	// 2345000000
	// 5678.00
	// ` 2.33`
	// '-808      '
}

func TestDecimalFormat(t *testing.T) {
	tests := []struct {
		fs   string
		d    Decimal
		want string
	}{
		{"%s", NewCents(123), "1.23"},
		{"%s", NewScalar(2345, -6), "2.345e+9"},
		{"%v", NewScalar(2345, -6), "2345000000"},
		{"%s", NewScalar(3, 9), "3e-9"},
		{"%.1f", NewCents(4567), "45.7"},
		{"%.2f", New(5678), "5678.00"},
		{"%.4f", NewCents(6789), "67.8900"},
		{"`%5.2f`", New(7).Div(New(3)), "` 2.33`"},
		{"'%-10.f'", NewCents(-80808), "'-808      '"},
		{"%.2f", NewScalar(9, 0), "9.00"},
		{"%.2f", NewScalar(10, 2), "0.10"},
		{"%.2f", NewScalar(11, -2), "1100.00"},
		{"%.2f", NewScalar(12, 4), "0.00"},
	}

	for _, tc := range tests {
		got := fmt.Sprintf(tc.fs, tc.d)
		if got != tc.want {
			t.Errorf("\n got: %v\nwant: %v\n", got, tc.want)
		}
	}
}
