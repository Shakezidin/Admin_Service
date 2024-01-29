package repository

import (
	DOM "github.com/Shakezidin/pkg/DOM/admin"
	inter "github.com/Shakezidin/pkg/admin/repository/interface"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (u *UserRepository) FetchAdmin(email string) (*DOM.Admin, error) {
	var admin DOM.Admin
	result := u.db.Where("email = ?", email).First(&admin)
	if result.Error != nil {
		return nil, result.Error
	}
	return &admin, nil
}

func (u *UserRepository) UpdateaWallet(amount float64) error {
	var admin DOM.Admin
	u.db.Where("id = ?", 1).First(&admin)
	result := u.db.Model(&admin).Update("wallet", admin.Wallet+amount)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func NewAdminRepository(db *gorm.DB) inter.RepoInterface {
	return &UserRepository{
		db: db,
	}
}
