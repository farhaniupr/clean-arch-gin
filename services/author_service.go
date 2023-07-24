package services

import (
	"github.com/dipeshdulal/clean-gin/domains"
	"github.com/dipeshdulal/clean-gin/lib"
	"github.com/dipeshdulal/clean-gin/models"
	"github.com/dipeshdulal/clean-gin/repository"
	"gorm.io/gorm"
)

// UserService service layer
type AuthorService struct {
	logger     lib.Logger
	repository repository.AuthorRepository
}

func NewAuthorService(logger lib.Logger, repository repository.AuthorRepository) domains.AuthorService {
	return AuthorService{
		logger:     logger,
		repository: repository,
	}
}

func (s AuthorService) WithTrx(trxHandle *gorm.DB) domains.AuthorService {
	s.repository = s.repository.WithTrx(trxHandle)
	return s
}

// GetOneAuthor
func (s AuthorService) GetOneAuthor(id int) (author models.Author, err error) {
	return author, s.repository.Find(&author, id).Error
}

// // GetAllUser get all the user
// func (s UserService) GetAllUser() (users []models.User, err error) {
// 	return users, s.repository.Find(&users).Error
// }

// // CreateUser call to create the user
// func (s UserService) CreateUser(user models.User) error {
// 	return s.repository.Create(&user).Error
// }

// // UpdateUser updates the user
// func (s UserService) UpdateUser(user models.User) error {
// 	return s.repository.Save(&user).Error
// }

// // DeleteUser deletes the user
// func (s UserService) DeleteUser(id uint) error {
// 	return s.repository.Delete(&models.User{}, id).Error
// }
