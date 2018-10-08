package money_test

import (
	"testing"

	"github.com/masa-suzu/go-tdd/money"
)

func TestTimes(t *testing.T) {
	type input struct {
		a money.Expression
		b []int
		k money.Kind
	}
	tests := []struct {
		name string
		in   input
		want *money.Currency
	}{
		{
			name: "D5*2=10",
			in:   input{a: money.NewDollar(5), b: []int{2}, k: money.USD},
			want: money.NewDollar(10)},
		{
			name: "D0*1=0",
			in:   input{a: money.NewDollar(0), b: []int{1}, k: money.USD},
			want: money.NewDollar(0)},
		{
			name: "D5*2*3=30",
			in:   input{a: money.NewDollar(5), b: []int{2, 3}, k: money.USD},
			want: money.NewDollar(30)},
		{
			name: "F5*2=10",
			in:   input{a: money.NewFranc(5), b: []int{2}, k: money.CHF},
			want: money.NewFranc(10)},
		{
			name: "F0*1=0",
			in:   input{a: money.NewFranc(0), b: []int{1}, k: money.CHF},
			want: money.NewFranc(0)},
		{
			name: "F5*2*3=30",
			in:   input{a: money.NewFranc(5), b: []int{2, 3}, k: money.CHF},
			want: money.NewFranc(30)},
		{
			name: "(D5+D2)*3=D21",
			in: input{
				a: &money.Sum{Augend: money.NewDollar(5), Addend: money.NewDollar(2)},
				b: []int{3},
				k: money.USD},
			want: money.NewDollar(21)},
	}

	for _, tt := range tests {

		bank := money.Bank{}
		t.Run(tt.name, func(t *testing.T) {

			var got money.Expression
			got = tt.in.a
			for _, x := range tt.in.b {
				got = got.Times(x)
			}

			if !got.Reduce(&bank, tt.in.k).Equals(tt.want) {
				t.Fatalf("%s.Times(%v) == %s = %v, want %v", tt.in.a, tt.in.b, tt.want, false, true)
			}
		})
	}
}

func TestPlus(t *testing.T) {
	type rate struct {
		from money.Kind
		to   money.Kind
		v    float64
	}
	type input struct {
		a    money.Expression
		b    money.Expression
		k    money.Kind
		rate rate
	}
	tests := []struct {
		name string
		in   input
		want money.Money
	}{
		{
			name: "D5+D5=D10",
			in:   input{a: money.NewDollar(5), b: money.NewDollar(5), k: money.USD, rate: rate{money.USD, money.USD, 1}},
			want: money.NewDollar(10)},
		{
			name: "F5+F5=F10",
			in:   input{a: money.NewFranc(5), b: money.NewFranc(5), k: money.CHF, rate: rate{money.CHF, money.CHF, 1}},
			want: money.NewFranc(10)},
		{
			name: "D5+F5=D10",
			in:   input{a: money.NewDollar(5), b: money.NewFranc(10), k: money.USD, rate: rate{money.CHF, money.USD, 0.5}},
			want: money.NewDollar(10)},
		{
			name: "D50+F50=F150",
			in:   input{a: money.NewDollar(50), b: money.NewFranc(50), k: money.CHF, rate: rate{money.USD, money.CHF, 2}},
			want: money.NewFranc(150)},
		{
			name: "(D50+D10)+D10=D70",
			in: input{
				a: &money.Sum{Augend: money.NewDollar(50), Addend: money.NewDollar(10)},
				b: money.NewDollar(10), k: money.USD, rate: rate{money.USD, money.USD, 1}},
			want: money.NewDollar(70)},
		{
			name: "F50+(F10+F10)=F70",
			in: input{
				a:    money.NewFranc(50),
				b:    &money.Sum{Augend: money.NewFranc(10), Addend: money.NewFranc(10)},
				k:    money.CHF,
				rate: rate{money.CHF, money.CHF, 1}},
			want: money.NewFranc(70)},
	}

	for _, tt := range tests {
		bank := money.Bank{}
		bank.UpdateRate(tt.in.rate.from, tt.in.rate.to, tt.in.rate.v)
		sum := tt.in.a.Plus(tt.in.b)
		got := bank.Reduce(sum, tt.in.k)
		if !got.Equals(tt.want) {
			t.Fatalf("%s.Equals(%s) = %v, want %v", got, tt.want, false, true)
		}
	}
}
