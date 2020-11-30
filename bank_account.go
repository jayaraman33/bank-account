package account

import "sync"

type Account struct {
	balance int
	open    bool
	wg      sync.Mutex
}

func Open(amount int) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{
		balance: amount,
		open:    true,
	}
}

func (a *Account) Balance() (int, bool) {
	return a.balance, a.open
}

func (a *Account) Close() (int, bool) {
	a.wg.Lock()
	defer func() {
		a.balance = 0
		a.open = false
		a.wg.Unlock()
	}()
	return a.balance, a.open
}

func (a *Account) Deposit(amount int) (int, bool) {
	a.wg.Lock()
	defer a.wg.Unlock()
	if a.balance+amount < 0 {
		return a.balance, false
	}
	a.balance += amount
	return a.balance, a.open
}
