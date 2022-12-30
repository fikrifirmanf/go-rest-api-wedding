package createUser

import (
	"fmt"

	model "github.com/fikrifirmanf/go-rest-api-wedding/models"
	util "github.com/fikrifirmanf/go-rest-api-wedding/utils"
	"gorm.io/gorm"
)

type Repository interface {
	CreateRepositoryUser(input *model.Users) (*model.Users, string, []string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryCreate(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateRepositoryUser(input *model.Users) (*model.Users, string, []string) {
	var users model.Users
	errorCode := make(chan string, 1)

	user := CreateUser{
		Name:     input.Name,
		Username: input.Username,
		Email:    input.Email,
		Role:     input.Role,
		Password: input.Password,
	}

	errs := util.TranslateError(user)
	if errs != nil {
		errorCode <- "VALIDATION_ERROR"
		fmt.Println(errs)
		return &users, <-errorCode, errs
	}
	checkDuplicateAccount := r.db.Where("email = ? OR username = ? ", input.Email, input.Username).First(&users)
	if checkDuplicateAccount.RowsAffected > 0 {
		errorCode <- "ACCOUNT_ALREADY_EXISTS"
		return &users, <-errorCode, []string{"Email or Username already exists"}
	}

	users.Name = input.Name
	users.Username = input.Username
	users.Email = input.Email
	users.Role = input.Role
	users.Password = input.Password

	results := r.db.Create(&users)
	if results.Error != nil {
		errorCode <- "CREATE_USER_FAILED"
		return &users, <-errorCode, []string{"Create user failed"}
	} else {
		errorCode <- "nil"
	}

	return &users, <-errorCode, nil
}
