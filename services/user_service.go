package services

import (
	"github.com/leslesnoa/bookstore_users-api/domain/users"
	resterr "github.com/leslesnoa/bookstore_users-api/utils/errors"
)

func GetUser(userId int64) (*users.User, *resterr.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func CreateUser(user users.User) (*users.User, *resterr.RestErr) {
	// リクエストのバリデーションチェック
	// err := user.Validate()
	// if err != nil {
	// 	return nil, err
	// }
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func UpdateUser(isPartial bool, user users.User) (*users.User, *resterr.RestErr) {
	current, err := GetUser(user.Id)
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

func DeleteUser(userId int64) *resterr.RestErr {
	user := &users.User{Id: userId}
	return user.Delete()
}

func Search(status string) ([]users.User, *resterr.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
