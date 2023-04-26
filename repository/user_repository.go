package repository

import (
	"belajar-echo/model"
	"database/sql"
)

type UserRepository struct {
	Sql *sql.DB
}

func (r *UserRepository) GetAllUsers() (result []model.User, err error) {
	var user model.User
	rows, err := r.Sql.Query("SELECT Name FROM user")
	if err != nil {
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&user.Username)
		result = append(result, user)
	}

	return result, err
}

func (r *UserRepository) LoginUser(username, password string) (res bool, err error) {
	row := r.Sql.QueryRow("SELECT EXIST(SELECT 1 FROM user WHERE username = ? AND password = ?)", username, password)
	if err := row.Scan(&res); err != nil {
		return false, err
	} else {
		return true, nil
	}

}
