package service

import (
	"project-template/internal/entity"
	"project-template/internal/repo"
)

type UserService interface {
	CreateUser(user *entity.User) error
	GetUserById(id int) (*[]entity.User, error)
	GetListOfUsers() (*entity.User, error)
	PutUser(user *entity.User) (*entity.User, error)
	DeleteUser(id int) error
}

type UserUseCase struct {
	userRepo repo.UserRepo
}

func (uc *UserUseCase) Create(user *entity.User) error {
	return uc.userRepo.Create(user)
}

func (uc *UserUseCase) GetUserById(id int) (*entity.User, error) {
	return uc.userRepo.GetById(id)
}

func (uc *UserUseCase) GetListOfUsers() (*[]entity.User, error) {
	return uc.userRepo.GetAll()
}

func (uc *UserUseCase) PutUser(user *entity.User) (*entity.User, error) {
	return uc.userRepo.PutUser(user)
}

func (uc *UserUseCase) DeleteUser(id int) error {
	return uc.DeleteUser(id)
}
