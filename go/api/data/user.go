package data

type UserDA struct {
	Connect *Connect
}

func NewUserDA(connect *Connect) *UserDA {
	return &UserDA{
		Connect: connect,
	}
}
