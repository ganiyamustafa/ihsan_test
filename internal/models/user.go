package models

import (
	"fmt"
	"time"

	"github.com/ganiyamustafa/bts/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID            uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
	Name          string    `json:"name"`
	NIK           string    `json:"nik" gorm:"unique"`
	Phone         string    `json:"phone" gorm:"unique"`
	AccountNumber string    `json:"no_rekening" gorm:"unique"`
	Wallet        *Wallet   `json:"wallet" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()

	return nil
}

func (u *User) AfterCreate(tx *gorm.DB) error {
	return u.createEmptyWalletAfterCreateFunc(tx)
}

// create wallet before create function
func (u *User) createEmptyWalletAfterCreateFunc(tx *gorm.DB) error {
	var balance float64 = 0

	wallet := Wallet{
		UserID:   u.ID,
		Balance:  &balance,
		Currency: "IDR",
	}

	return tx.Create(&wallet).Error
}

// generate account number function
func (u *User) GenerateAccountNumber() error {
	// generate account number with time unix and random int
	accountNumber := fmt.Sprintf("%v%s", time.Now().Unix(), utils.RandStringNumber(3))

	// set user account number
	u.AccountNumber = accountNumber
	return nil
}

type UserScopes struct{}

// scopes for preload product
func (u UserScopes) PreloadWallet(scopes []func(*gorm.DB) *gorm.DB, column ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload("Wallet", func(db *gorm.DB) *gorm.DB {
			// return selected column product if there are filter on parameter
			if len(column) > 0 {
				return db.Scopes(scopes...).Select(column)
			}

			return db.Scopes(scopes...)
		})
	}
}
