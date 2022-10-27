package endpoints

import (
	"api/data"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateUserHandler(writer http.ResponseWriter, request *http.Request) {
	access, err := data.Open()
	if err != nil {

		return
	}

	defer access.DB.Close()

	data := NewUserDA(access)

}

func MakeUserHandlers(router *mux.Router) error {

	function := ValididateToken(CreateUserHandler)
	router.HandleFunc("/create", function).Methods("POST")

	return nil
}
