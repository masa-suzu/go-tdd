package money_test

import (
	"testing"

	"github.com/masa-suzu/go-tdd/money"
)

func TestReduceMoneyFromBank(t *testing.T) {
	type input struct {
		ex money.Expression
		k  money.Kind
	}
	tests := []struct {
		name string
		in   input
		want money.Money
	}{
		{
			name: "D5->D5",
			in:   input{ex: money.NewDollar(5), k: money.USD},
			want: money.NewDollar(5)},
		{
			name: "F4->F4",
			in:   input{ex: money.NewFranc(4), k: money.CHF},
			want: money.NewFranc(4)},
	}

	bank := money.Bank{}

	for _, tt := range tests {
		got := bank.Reduce(tt.in.ex, tt.in.k)
		if !got.Equals(tt.want) {
			t.Fatalf("bank.Reduce(%s, %s) = %s, want %s", tt.in.ex, tt.in.k, got, tt.want)
		}
	}
}

func TestReduceMoneyDifferentCurrency(t *testing.T) {
	type input struct {
		c    *money.Currency
		to   money.Kind
		rate float64
	}
	tests := []struct {
		in   input
		want *money.Currency
	}{
		{in: input{c: money.NewDollar(2), to: money.CHF, rate: 2}, want: money.NewFranc(4)},
		{in: input{c: money.NewDollar(2), to: money.CHF, rate: 3}, want: money.NewFranc(6)},
		{in: input{c: money.NewFranc(2), to: money.USD, rate: 0.5}, want: money.NewDollar(1)},
		{in: input{c: money.NewDollar(2), to: money.USD, rate: 1}, want: money.NewDollar(2)},
		{in: input{c: money.NewFranc(2), to: money.CHF, rate: 1}, want: money.NewFranc(2)},
	}

	for _, tt := range tests {
		bank := money.Bank{}
		bank.UpdateRate(tt.in.c.Kind(), tt.in.to, tt.in.rate)
		got := bank.Reduce(tt.in.c, tt.in.to)

		if !got.Equals(tt.want) {
			t.Fatalf("bank.Reduce(%s, %s) = %s, want %s", tt.in.c, tt.in.to, got, tt.want)
		}
	}
}
