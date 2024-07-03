package main

import (
	"errors"
	"fmt"
)

// This is a type created from an int,
// so it is an int, but can receive methods.
type Bitcoin int

// this is an interface implementing the String() method
// "Stringer is implemented by any value that has a String method, which defines the “native” format for that value."
// In other words, this method already exists in the fmt package.
// We are instantiating it here so we can use it with Bitcoin.
type Stringer interface {
	String() string
}

// Method declaration on a type receiver is the same as on a struct receiver
// this makes the test read: wallet_test.go|20| got 10 BTC want 20 BTC
// because we're defining that printing this type returns words (BTC)
// not just the value (10 and 20).
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

// these methods take pointers to a variable of type Wallet to keep consistency
// this will allow the function to change the variable as held by the caller
// not a copy of the variable if this receiver were not for a pointer
func (w *Wallet) Deposit(amount Bitcoin) {
	// fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	w.balance += amount
}

// balance doesn't change anything,
// it doesn't need to take a reference,
// but for consistency we keep it
// Also note that `return w.balance` is automatically dereferenced
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Define a global variable to make an error message available everywhere
var ErrInsufficentFunds = errors.New("cannot withdraw, insufficient funds")

// now for a withdraw method
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficentFunds
	}
	w.balance -= amount
	return nil
}
