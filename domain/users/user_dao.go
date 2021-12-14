package users

import (
	"fmt"

	"github.com/leslesnoa/bookstore_users-api/datasources/mysql/users_db"
	"github.com/leslesnoa/bookstore_users-api/logger"
	"github.com/leslesnoa/bookstore_users-api/utils/date_utils"
	resterr "github.com/leslesnoa/bookstore_users-api/utils/errors"
	"github.com/leslesnoa/bookstore_users-api/utils/mysql_utils"
)

const (
	queryInsertUser       = "INSERT INTO users (first_name, last_name, email, date_created, status, password) VALUES(?,?,?,?,?,?);"
	queryGetUser          = "SELECT id, first_name, last_name, email, date_created, status, password FROM users WHERE id=?;"
	queryUpdateUser       = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser       = "DELETE FROM users WHERE id=?;"
	queryFindUserByStatus = "SELECT id, first_name, last_name, email, data_created, status FROM users WHERE status=?;"
)

// SQL導入前のDBスタブ
// var (
// 	usersDB = make(map[int64]*User)
// )

// ポインタにしない場合
func somethingGet() {
	user := User{}

	if err := user.Get(); err != nil {
		fmt.Println(err)
		return
	}
}

func (user *User) Get() *resterr.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return resterr.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow()
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status, &user.Password); err != nil {
		logger.Error("error when trying to prepare get user by id", getErr)
		return mysql_utils.ParseError(getErr)
		// MySQLのエラーかどうか確認
		// sqlErr, ok := getErr.(*mysql.MySQLError)
		// if !ok {
		// 	return resterr.NewInternalServerError(fmt.Sprintf("error when trying to get user %d: %s", user.Id, getErr.Error()))
		// }
		// SQLのエラーハンドリング
		// fmt.Println(sqlErr.Number)
		// fmt.Println(sqlErr.Message)
	}
	return nil

	// SQL導入前
	// if err := users_db.Client.Ping(); err != nil {
	// 	panic(err)
	// }
	// result := usersDB[user.Id]
	// if result == nil {
	// 	return resterr.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	// }

	// user.Id = result.Id
	// user.FirstName = result.FirstName
	// user.LastName = result.LastName
	// user.Email = result.Email
	// user.DateCreated = result.DateCreated
}

func (user *User) Save() *resterr.RestErr {
	stmt, saveErr := users_db.Client.Prepare(queryInsertUser)
	if saveErr != nil {
		logger.Error("error when trying to prepare get user by id", saveErr)
		return resterr.NewInternalServerError("database error")
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if saveErr != nil {
		logger.Error("error when trying to prepare create user", saveErr)
		return mysql_utils.ParseError(saveErr)
		// MySQLのエラーかどうか確認
		// sqlErr, ok := saveErr.(*mysql.MySQLError)
		// if !ok {
		// 	return resterr.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", saveErr.Error()))
		// }
		// SQLのエラーハンドリング
		// fmt.Println(sqlErr.Number)
		// fmt.Println(sqlErr.Message)
		// switch sqlErr.Number {
		// case 1062:
		// 	return resterr.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
		// }
		// return resterr.NewInternalServerError(
		// 	fmt.Sprintf("error when trying to save user: %s", saveErr.Error()),
		// )
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to prepare insert result", err)
		return resterr.NewInternalServerError("database error")
	}
	user.Id = userId

	return nil

	// current := usersDB[user.Id]
	// if current != nil {
	// 	if current.Email == user.Email {
	// 		return resterr.NewBadRequestError(fmt.Sprintf("email %s already registered", user.Email))
	// 	}
	// 	return resterr.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	// }

	// 現在時刻をセット
	// user.DateCreated = date_utils.GetNowString()

	// usersDB[user.Id] = user
}

func (user *User) Update() *resterr.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when trying to prepare update user by id", err)
		return resterr.NewInternalServerError("database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		logger.Error("error when trying to prepare update user exec", err)
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) Delete() *resterr.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare delete user by id", err)
		return mysql_utils.ParseError(err)
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.Id); err != nil {
		logger.Error("error when trying to prepare delete user exec", err)
		return mysql_utils.ParseError(err)
	}

	return nil
}

func (user *User) FindByStatus(status string) ([]User, *resterr.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		logger.Error("error when trying to prepare find user by id", err)
		return nil, resterr.NewInternalServerError("database error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error when trying to prepare find user by id", err)
		return nil, resterr.NewInternalServerError("database error")
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("error when trying to prepare find user scan", err)
			return nil, mysql_utils.ParseError(err)
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, resterr.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}

	return results, nil
}
