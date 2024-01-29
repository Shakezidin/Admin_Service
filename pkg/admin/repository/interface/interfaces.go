package interfaces

import (
	DOM "github.com/Shakezidin/pkg/DOM/admin"
)

type RepoInterface interface {
	FetchAdmin(email string) (*DOM.Admin, error)
	UpdateaWallet(amount float64) error
}
