package getUser

import (
	model "github.com/fikrifirmanf/go-rest-api-wedding/models"
	"gorm.io/gorm"
)

type Repository interface {
	GetRepositoryUsers() (*[]model.Users, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetRepositoryUsers() (*[]model.Users, string) {
	var users []model.Users
	errorCode := make(chan string, 1)

	results := r.db.Find(&users)
	if results.Error != nil {
		errorCode <- "USERS_NOT_FOUND"
		return &users, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &users, <-errorCode
}
