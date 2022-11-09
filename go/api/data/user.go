package data

import "database/sql"

type UserDA struct {
	Connect *Connect
}

// ############################################################## Struct

type CreateUser struct {
	UserPassword string `json:"user_password"`
	HashPassword string `json:"hash_password"`
	Confirm      string `json:"confirm"`
	Username     string `json:"username"`
	Name         string `json:"name"`
	Surname      string `json:"surname"`
	Email        string `json:"email"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func NewUserDA(connect *Connect) *UserDA {
	return &UserDA{
		Connect: connect,
	}
}

// ############################################################## Mappings

func mapInt(rows *sql.Rows) (int64, error) {

	var i int64
	err := rows.Scan(
		&i,
	)
	if err != nil {
		return 0, err
	}

	return i, nil
}

// ############################################################## Methods

func (da *UserDA) CreateUser(hashPassword, name, surname, username, email, token, refreshToken string) ([]int64, error) {

	return Query(da.Connect, `SELECT * hidotcom.create_user($1, $2, $3, $4, $5, $6, $7)`,
		mapInt, name, surname, email, hashPassword, token, refreshToken)

}
