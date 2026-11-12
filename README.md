# Go Interface Lesson: Payment Methods

This project demonstrates how Go interfaces allow different payment types to share the same behavior without forcing them to have the same internal data.

## How the pieces fit together

- Structs hold state:
  - `MPesa`, `BankAccount`, and `CreditCard` each store different data.
  - Example: phone numbers, account balances, and credit limits.

- Methods define behavior:
  - Each struct implements its own `Pay(amount float64) error` method.
  - `MPesa` and `BankAccount` reduce a balance.
  - `CreditCard` increases `CurrentDebt`.

- Interfaces decouple the code:
  - The `PaymentMethod` interface only requires one thing: `Pay(amount float64) error`.
  - It does not care about the concrete fields inside each struct.

- Functions become reusable:
  - `ProcessCheckout` accepts a `PaymentMethod` instead of a specific payment type.
  - This means one checkout function can work with many payment methods.
  - If a new provider like PayPal or Crypto is added later, `ProcessCheckout` does not need to change.

## Why pointers matter

In `main()`, the code creates values like `&MPesa{...}` instead of `MPesa{...}` because the methods use pointer receivers:

```go
func (m *MPesa) Pay(amount float64) error
```

This matters because `Pay()` changes the struct's balance. If a value copy were used instead of a pointer, the method would modify the copy, not the original object.

## Interface rule

When a method uses a pointer receiver, only a pointer to that type satisfies the interface.

For example:

- `*MPesa` satisfies `PaymentMethod`
- `MPesa` does not satisfy `PaymentMethod`

This is important because Go requires the exact method set to match the interface.

## Summary

This example shows the core strength of interfaces in Go:

- different types can share the same behavior,
- code can depend on abstractions instead of concrete implementations,
- and new payment methods can be added without rewriting checkout logic.
