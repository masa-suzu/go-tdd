package money

type Expression interface {
	Reduce(*Bank, Kind) Money
	Plus(money Expression) Expression
	Times(int) Expression
}

type Sum struct {
	Augend Expression
	Addend Expression
}

func (s *Sum) Reduce(b *Bank, kind Kind) Money {
	x := s.Augend.Reduce(b, kind)
	y := s.Addend.Reduce(b, kind)

	return newCurrency(x.Amount()+y.Amount(), kind)
}

func (s *Sum) Plus(e Expression) Expression {
	return &Sum{s, e}
}

func (s *Sum) Times(x int) Expression {
	return &Sum{s.Addend.Times(x), s.Augend.Times(x)}
}
