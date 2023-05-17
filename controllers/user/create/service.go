package createUser

import model "github.com/fikrifirmanf/go-rest-api-wedding/models"

type Service interface {
	CreateServiceUser(input *CreateUser) (*model.Users, string, []string)
}

type service struct {
	repository Repository
}

func NewServiceCreate(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) CreateServiceUser(input *CreateUser) (*model.Users, string, []string) {
	users := &model.Users{
		Name:     input.Name,
		Username: input.Username,
		Email:    input.Email,
		Role:     input.Role,
		Password: input.Password,
	}

	result, errorResult, errMessage := s.repository.CreateRepositoryUser(users)
	return result, errorResult, errMessage
}
