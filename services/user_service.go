package services

import (
	"github.com/leslesnoa/bookstore_users-api/domain/users"
	"github.com/leslesnoa/bookstore_users-api/utils/crypto_utils"
	resterr "github.com/leslesnoa/bookstore_users-api/utils/errors"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct {
}

type usersServiceInterface interface {
	GetUser(int64) (*users.User, *resterr.RestErr)
	CreateUser(users.User) (*users.User, *resterr.RestErr)
	UpdateUser(bool, users.User) (*users.User, *resterr.RestErr)
	DeleteUser(int64) *resterr.RestErr
	Search(string) ([]users.User, *resterr.RestErr)
}

func (s *usersService) GetUser(userId int64) (*users.User, *resterr.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *usersService) CreateUser(user users.User) (*users.User, *resterr.RestErr) {
	// リクエストのバリデーションチェック
	// err := user.Validate()
	// if err != nil {
	// 	return nil, err
	// }

	user.Status = users.StatusActive
	user.Password = crypto_utils.GetMd5(user.Password)

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *usersService) UpdateUser(isPartial bool, user users.User) (*users.User, *resterr.RestErr) {
	current, err := UsersService.GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}
	// PATCHリクエストの場合
	if !isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}

		if user.LastName != "" {
			current.LastName = user.LastName
		}

		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func (s *usersService) DeleteUser(userId int64) *resterr.RestErr {
	user := &users.User{Id: userId}
	return user.Delete()
}

func (s *usersService) Search(status string) ([]users.User, *resterr.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
