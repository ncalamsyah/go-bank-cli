package main

import (
	"encoding/base64"
	"errors"
)

func GenerateAccountNumber(p Person) string {
	email := p.User.Email
	accountHash := base64.StdEncoding.EncodeToString([]byte(email))
	return accountHash
}

func ValidateUser(email, password string) bool {
	for _, user := range Users {
		if user.Email == email && user.Password == password {
			return true
		}
	}
	return false
}

func GetData(person Person, detailType string) (*BankDetail, *[]UserTransaction, *[]TransferTransaction, error) {
	var userBankDetails BankDetail
	var userTransactions []UserTransaction
	var transferTransactions []TransferTransaction

	switch detailType {
	case "bankDetails":
		for _, bankDetail := range BankDetails {
			if bankDetail.AccountHash == person.AccountHash {
				userBankDetails = bankDetail
			}
		}
		return &userBankDetails, nil, nil, nil
	case "userTransactions":
		for _, userTransaction := range userTransactions {
			if userTransaction.AccountHash == person.AccountHash {
				userTransactions = append(userTransactions, userTransaction)
			}
		}
		return nil, &userTransactions, nil, nil
	case "transferTransactions":
		for _, transferTransaction := range TransferTransactions {
			if transferTransaction.From == person.AccountHash {
				transferTransactions = append(transferTransactions, transferTransaction)
			}
		}
		return nil, nil, &transferTransactions, nil
	default:
		return nil, nil, nil, errors.New("no valid option selected")
	}
}
