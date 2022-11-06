package data

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
}

func NewUserDA(connect *Connect) *UserDA {
	return &UserDA{
		Connect: connect,
	}
}

// ############################################################## Mappings

// ############################################################## Methods
