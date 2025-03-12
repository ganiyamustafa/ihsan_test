package services

import (
	"github.com/ganiyamustafa/bts/internal/models"
	"github.com/ganiyamustafa/bts/internal/requests"
	"github.com/ganiyamustafa/bts/utils"
	apperror "github.com/ganiyamustafa/bts/utils/app_error"
	"github.com/jinzhu/copier"
)

type UserService struct {
	Handler *utils.Handler
}

// create user function
func (u *UserService) CreateUser(payload requests.RegisterRequest) (*models.User, *apperror.AppError) {
	// copy payload to user models
	var user models.User
	copier.Copy(&user, &payload)

	// generate account number
	user.GenerateAccountNumber()

	return &user, apperror.FromError(u.Handler.Postgre.Create(&user).Error)
}

// get user by nik or phone
func (u *UserService) GetUserByNIKOrPhone(nik, phone string) (*models.User, *apperror.AppError) {
	var user models.User
	return &user, apperror.FromError(u.Handler.Postgre.Where("nik = ?", nik).Or("phone = ?", phone).First(&user).Error)
}

// get user by account number, nik and phone
func (u *UserService) GetUserByAccountNumberNIKAndPhone(accountNumber, nik, phone string) (*models.User, *apperror.AppError) {
	var user models.User
	return &user, apperror.FromError(u.Handler.Postgre.Where("account_number = ? and nik = ? and phone = ?", accountNumber, nik, phone).First(&user).Error)
}

// get user by account number
func (u *UserService) GetUserByAccountNumber(accountNumber string) (*models.User, *apperror.AppError) {
	var user models.User
	var db = u.Handler.Postgre.Table("users")

	db = db.Scopes(
		models.UserScopes{}.PreloadWallet(nil),
	)

	return &user, apperror.FromError(db.Where("account_number = ?", accountNumber).First(&user).Error)
}
