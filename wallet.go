package main

import "fmt"

type Wallet struct {
	balance int
}

// these methods take pointers to a variable of type Wallet to keep consistency
// this will allow the function to change the variable as held by the caller
// not a copy of the variable if this receiver were not for a pointer
func (w *Wallet) Deposit(amount int) {
	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	w.balance += amount
}

// balance doesn't change anything,
// it doesn't need to take a reference,
// but for consistency we keep it
// Also note that `return w.balance` is automatically dereferenced
func (w *Wallet) Balance() int {
	return w.balance
}
