package users

// user_dao.go

import (
	"fmt"

	"github.com/judesantos/go-bookstore_users_api/datasources/mysql/users_db"
	crypto_utils "github.com/judesantos/go-bookstore_users_api/utils/crypto"
	"github.com/judesantos/go-bookstore_users_api/utils/errors"
	mysql_utils "github.com/judesantos/go-bookstore_users_api/utils/mysql"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	errorNoRows      = "no rows in result set"

	queryInsertUser = "INSERT INTO users(first_name, last_name, email, status, password) " +
		"VALUES(?,?,?,?,?);"
	queryGetUser = "SELECT id, first_name, last_name, email, status, date_created " +
		"FROM users WHERE id=?;"
	queryUpdateUser = "UPDATE users SET first_name=?, last_name=?, email=?, status=?, " +
		"password=? WHERE id=?;"
	queryDeleteUser       = "DELETE FROM users where id=?"
	queryFindUserByStatus = "SELECT id, first_name, last_name, email, date_created, status " +
		"FROM users WHERE status = ?;"
	queryLoginInfo = "SELECT id, first_name, last_name, email, date_created, status " +
		"FROM users WHERE email = ? AND password = ?;"
)

var (
	usersDb = make(map[int64]*User)
)

//
//	Get - Get user by id
//
func (user *User) Get() *errors.RestError {

	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	defer stmt.Close()

	result := stmt.QueryRow(user.Id)

	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName,
		&user.Email, &user.Status, &user.DateCreated); err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil
}

//
// Save - Save new user
//
func (user *User) Save() *errors.RestError {

	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	defer stmt.Close()

	user.Password = crypto_utils.GetMd5(user.Password)

	result, err := stmt.Exec(user.FirstName, user.LastName, user.Email,
		user.Status, user.Password)
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	user.Id = userId

	return nil
}

//
// Update - Update user
//
func (user *User) Update() *errors.RestError {

	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName,
		user.Email, user.Status, user.Password, user.Id)
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil
}

//
// Delete - Delete user
//
func (user *User) Delete() *errors.RestError {

	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.Id)
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil
}

//
// FindByStatus - Find users by status
//
func (user *User) FindByStatus(status string) (Users, *errors.RestError) {

	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		return nil, mysql_utils.ParseError(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, mysql_utils.ParseError(err)
	}

	defer rows.Close()

	results := Users{}

	for rows.Next() {
		var user User
		if err := rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.DateCreated,
			&user.Status); err != nil {
			return nil, mysql_utils.ParseError(err)
		}

		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NotFoundError(
			fmt.Sprintf("no users found with matching status %s", status))
	}

	return results, nil
}

//
// FindByEmailAndPassword - Find user by email and password
//
func (user *User) FindByEmailAndPassword() *errors.RestError {

	stmt, err := users_db.Client.Prepare(queryLoginInfo)
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(user.Email, user.Password)
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	defer rows.Close()

	user.Password = crypto_utils.GetMd5(user.Password)

	result := stmt.QueryRow(user.Email, user.Password)

	if err := result.Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.DateCreated,
		&user.Status,
	); err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil
}
