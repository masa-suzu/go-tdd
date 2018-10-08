package money

type Bank struct {
	rates map[string]float64
}

func (b *Bank) Reduce(e Expression, kind Kind) Money {
	return e.Reduce(b, kind)
}

func (b *Bank) UpdateRate(from Kind, to Kind, rate float64) {
	if b.rates == nil {
		b.rates = make(map[string]float64)
	}
	b.rates[from.String()+to.String()] = rate
}

func (b *Bank) Rate(from Kind, to Kind) float64 {
	rate, ok := b.rates[from.String()+to.String()]
	if !ok {
		return 1
	}
	return rate
}
