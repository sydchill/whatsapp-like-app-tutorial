package endpoints

import (
	"api/data"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func CreateUserHandler(writer http.ResponseWriter, request *http.Request) {
	access, err := data.Open()
	if err != nil {

		InternalServerError(writer, request, err)
	}
	defer access.DB.Close()

	var input data.CreateUser
	decoder := json.NewDecoder(request.Body)
	err = decoder.Decode(&input)

	if err != nil {
		BadRequest(writer, request, "invalid_request: Invalid request JSON")
	}

	if !pattern.MatchString(input.Email) {
		BadRequest(writer, request, "invalid_Email")
	}

	if input.Confirm != input.UserPassword {
		BadRequest(writer, request, "Invalid_password")
	}

	input.HashPassword = HashPassword(input.UserPassword)

	// dal := data.NewUserDA(access)

}

// ############################################################## HELPER FUNCTIONS

//HashPassword is used to encrypt the password before it is stored in the DB
func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}

	return string(bytes)
}

// ############################################################## HANDLERS

func MakeUserHandlers(router *mux.Router) error {

	// function := ValididateToken(CreateUserHandler)
	router.HandleFunc("/create", CreateUserHandler).Methods("POST")

	return nil
}
