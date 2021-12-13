package users

import (
	"strings"

	resterr "github.com/leslesnoa/bookstore_users-api/utils/errors"
)

// ユーザAPIのフィールドを定義
type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

// ユーザAPIのバリデーションを定義
func (user *User) Validate() *resterr.RestErr {
	// 受け取ったメールアドレスを小文字にして整形
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		// メールが空白の場合エラーを返す
		return resterr.NewBadRequestError("invalid email address")
	}
	return nil
}
