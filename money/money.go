package money

type Kind int

const (
	USD Kind = iota
	CHF
)

func (k Kind) String() string {
	switch k {
	case USD:
		return "USD"
	case CHF:
		return "CHF"
	default:
		return "UNDEFINED"
	}
}

type Money interface {
	Equals(money Money) bool
	Amount() int
	Kind() Kind
}

func NewDollar(amount int) *Currency {
	return &Currency{amount: amount, kind: USD}
}

func NewFranc(amount int) *Currency {
	return &Currency{amount: amount, kind: CHF}
}

func newCurrency(amount int, kind Kind) *Currency {
	switch kind {
	case USD:
		return NewDollar(amount)
	case CHF:
		return NewFranc(amount)
	}
	return nil
}

func equals(m1 Money, m2 Money) bool {
	return m1.Kind() == m2.Kind() && m1.Amount() == m2.Amount()
}
