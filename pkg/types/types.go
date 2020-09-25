package types

type Money int

type Currency string

const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

type PAN string

type Card struct {
	ID         int
	PAN        PAN
	Balance    Money
	MinBalance Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
}

//Category for example(car, purchases)
type Category string

//Payment представляет информацию о платеже
type Payment struct {
	ID       int
	Amount   Money
	Category Category
}
