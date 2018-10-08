package money

import "fmt"

type Currency struct {
	amount int
	kind   Kind
}

func (c *Currency) Equals(m Money) bool {
	return equals(c, m)
}

func (c *Currency) Amount() int {
	return c.amount
}

func (c *Currency) Kind() Kind {
	return c.kind
}

func (c *Currency) Plus(money Expression) Expression {
	return &Sum{Augend: c, Addend: money}
}

func (c *Currency) Times(x int) Expression {
	return newCurrency(c.Amount()*x, c.kind)
}

func (c *Currency) Reduce(b *Bank, to Kind) Money {
	a := float64(c.Amount()) * b.Rate(c.kind, to)
	return newCurrency(int(a), to)
}

func (c *Currency) String() string {
	return fmt.Sprintf("%s{%v}", c.Kind(), c.amount)
}
