package money_test

import (
	"testing"

	"github.com/masa-suzu/go-tdd/money"
)

func TestNewCurrency(t *testing.T) {
	type want struct {
		amount  int
		kind    money.Kind
		literal string
	}
	tests := []struct {
		name string
		in   *money.Currency
		want want
	}{
		{name: "D5", in: money.NewDollar(5), want: want{amount: 5, kind: money.USD, literal: "USD{5}"}},
		{name: "F5", in: money.NewFranc(4), want: want{amount: 4, kind: money.CHF, literal: "CHF{4}"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := tt.in

			if got.Amount() != tt.want.amount {
				t.Fatalf("%s.Amount() = %v, want %v", got, got.Amount(), tt.want.amount)
			}
			if got.Kind() != tt.want.kind {
				t.Fatalf("%s.Kind() = %v, want %v", got, got.Kind(), tt.want.kind)
			}
			if got.String() != tt.want.literal {
				t.Fatalf("%s.String() = %v, want %v", got, got.String(), tt.want.literal)
			}
		})
	}
}

func TestEquals(t *testing.T) {
	type input struct {
		a money.Money
		b money.Money
	}
	tests := []struct {
		name string
		in   input
		want bool
	}{
		{name: "D5=D5", in: input{a: money.NewDollar(5), b: money.NewDollar(5)}, want: true},
		{name: "D4=D4", in: input{a: money.NewDollar(4), b: money.NewDollar(4)}, want: true},
		{name: "D4!=D5", in: input{a: money.NewDollar(4), b: money.NewDollar(5)}, want: false},
		{name: "D5!=D4", in: input{a: money.NewDollar(5), b: money.NewDollar(4)}, want: false},
		{name: "F5=F5", in: input{a: money.NewFranc(5), b: money.NewFranc(5)}, want: true},
		{name: "F4=F4", in: input{a: money.NewFranc(4), b: money.NewFranc(4)}, want: true},
		{name: "F4!=F5", in: input{a: money.NewFranc(4), b: money.NewFranc(5)}, want: false},
		{name: "F5!=F4", in: input{a: money.NewFranc(5), b: money.NewFranc(4)}, want: false},

		{name: "D4!=F4", in: input{a: money.NewDollar(4), b: money.NewFranc(4)}, want: false},
		{name: "D4!=F5", in: input{a: money.NewDollar(4), b: money.NewFranc(5)}, want: false},
		{name: "F4!=D4", in: input{a: money.NewFranc(4), b: money.NewDollar(4)}, want: false},
		{name: "F4!=D5", in: input{a: money.NewFranc(4), b: money.NewDollar(5)}, want: false},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			got := tt.in.a.Equals(tt.in.b)

			if got != tt.want {
				t.Fatalf("%v.Equals(%v) = %v, want %v", tt.in.a, tt.in.b, got, tt.want)
			}
		})
	}
}
