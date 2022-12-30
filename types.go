package main

// Auth user
type User struct {
	Email, Password string
}

// General user
type Person struct {
	User              User
	Name, AccountHash string
}

type BankDetail struct {
	AccountHash    string
	CurrentBalance float64
}

// User's transaction with self
type UserTransaction struct {
	AccountHash    string
	WithdrawAmount float64
	DepositAmount  float64
}

// User's transaction with another user
type TransferTransaction struct {
	Amount float64
	From   string // Sending user's account number
	To     string // Receiving user's account number
}
