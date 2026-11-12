package main

import (
	"errors"
	"fmt"
)

// 1. THE INTERFACE (The Contract)
// Any type that has a `Pay(amount float64) error` method automatically implements PaymentMethod.

type PaymentMethod interface {
	Pay(amount float64) error
}

// 2. THE STRUCTS & THEIR METHODS (The Concrete Implementations)

// M-Pesa
type MPesa struct {
	PhoneNumber string
	Balance     float64
}

// Pointer receiver (*Mpesa) so we can modify the balance
func (m *MPesa) Pay(amount float64) error {
	if amount > m.Balance {
		return errors.New("Insufficient M-Pesa balance")
	}
	m.Balance -= amount
	fmt.Printf("Paid KES %.2f via M-Pesa (%s). Remaining balance: KES %.2f\n", amount, m.PhoneNumber, m.Balance)
	return nil
}

// -- Bank Account ---
type BankAccount struct {
	AccountNumber string
	Balance       float64
}

func (b *BankAccount) Pay(amount float64) error {
	if amount > b.Balance {
		return errors.New("Insufficient bank funds")
	}
	b.Balance -= amount
	fmt.Printf("Paid KES %.2f from Bank Account (%s). Remaining balance: KES %.2f\n", amount, b.AccountNumber, b.Balance)
	return nil
}

// --- Credit Card ---
type CreditCard struct {
	CardNumber  string
	CreditLimit float64
	CurrentDebt float64
}

func (c *CreditCard) Pay(amount float64) error {
	if c.CurrentDebt+amount > c.CreditLimit {
		return errors.New("credit card limit exceeded")
	}
	c.CurrentDebt += amount
	available := c.CreditLimit - c.CurrentDebt
	fmt.Printf("Charged KES %.2f to Credit Card (%s). Available credit: KES %.2f\n", amount, c.CardNumber, available)
	return nil
}

// ---------------------------------------------------------
// 3. THE FUNCTION (Polymorphism in Action)
// ---------------------------------------------------------

// ProcessCheckout accepts the INTERFACE, not a specific struct.
// It doesn't care HOW the payment is processed, only that the method exists.
func ProcessCheckout(pm PaymentMethod, amount float64) {
	fmt.Printf("\n--- Initiating checkout of KES %.2f ---\n", amount)
	err := pm.Pay(amount)
	if err != nil {
		fmt.Println("Transaction Failed:", err)
	} else {
		fmt.Println("Transaction Successful!")
	}
}

func main() {
	// Initialize our payment structs
	myMpesa := &MPesa{PhoneNumber: "0712345678", Balance: 5000.00}
	myBank := &BankAccount{AccountNumber: "ACC-98765", Balance: 25000.00}
	myCard := &CreditCard{CardNumber: "4111-XXXX-1111", CreditLimit: 50000.00, CurrentDebt: 10000.00}

	// Pass different structs into the exact same function!
	ProcessCheckout(myMpesa, 1500.00) // Works because *MPesa has Pay()
	ProcessCheckout(myBank, 12000.00) // Works because *BankAccount has Pay()
	ProcessCheckout(myCard, 35000.00) // Works because *CreditCard has Pay()

	// Trigger an error to see validation in action
	ProcessCheckout(myMpesa, 4000.00) // Fails: only 3500 left

}
