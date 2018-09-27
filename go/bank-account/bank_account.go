package account

import "sync"

// Account struct keeps the balance amount of a person's bank account
type Account struct {
	open    bool
	balance int64
	mux     sync.Mutex
}

// Open initiate an account with a initial deposit amount and set the returned account as opened
func Open(initialDeposit int64) *Account {
	var acc Account

	if initialDeposit < 0 {
		return nil
	}

	acc.open = true
	acc.balance = initialDeposit

	return &acc
}

// Close returns the account balance and if it was closed sucessfully
func (acc *Account) Close() (payout int64, ok bool) {

	ok = acc.open

	acc.mux.Lock()

	if acc.open {
		payout = acc.balance
	} else {
		payout = 0
	}

	acc.open = false

	acc.mux.Unlock()

	return
}

// Deposit or withdrawal an amount into account balance
func (acc *Account) Deposit(amount int64) (newBalance int64, ok bool) {

	acc.mux.Lock()
	newBalance = acc.balance + amount

	if acc.open && newBalance >= 0 {
		ok = true
		acc.balance = newBalance
	} else {
		newBalance = acc.balance
		ok = false
	}

	acc.mux.Unlock()

	return
}

// Balance returns account balance amount
func (acc *Account) Balance() (balance int64, ok bool) {

	if acc.open {
		balance = acc.balance
		ok = true
	} else {
		balance = 0
		ok = false
	}

	return
}
