package services

import (
	"github.com/leslesnoa/bookstore_users-api/domain/users"
	resterr "github.com/leslesnoa/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *resterr.RestErr) {
	// リクエストのバリデーションチェック
	if err := user.Validate(); err != nil {
		return nil, err
	}

	return nil, nil
}
