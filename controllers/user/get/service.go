package getUser

import model "github.com/fikrifirmanf/go-rest-api-wedding/models"

type Service interface {
	GetServiceUsers() (*[]model.Users, string)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetServiceUsers() (*[]model.Users, string) {
	getUsers, errorGetusers := s.repository.GetRepositoryUsers()

	return getUsers, errorGetusers
}
