package pointers_and_errors

type Wallet struct {
	Amount int
}

func (w Wallet) Deposit(amount int) {
	w.Amount += amount
}

func (w Wallet) Balance() int {
	return w.Amount
}
