package endpoints

import (
	"api/data"
	"encoding/json"
	"fmt"
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
		return
	}

	if !pattern.MatchString(input.Email) {
		BadRequest(writer, request, "invalid_Email")
		return
	}

	if input.Confirm != input.UserPassword {
		BadRequest(writer, request, "Invalid_password")
		return
	}

	input.HashPassword = HashPassword(input.UserPassword)

	dal := data.NewUserDA(access)

	userID, err := dal.CreateUser(input.HashPassword, input.Name, input.Surname, input.Username, input.Email, input.Token, input.RefreshToken)

	if err != nil {

		InternalServerError(writer, request, err)
		return
	}

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

//VerifyPassword checks the input password while verifying it with the passward in the DB.
func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	match := true
	var message string

	if err != nil {
		message = fmt.Sprintf("password incorrect")
		match = false
	}

	return match, message
}

// ############################################################## HANDLERS

func MakeUserHandlers(router *mux.Router) error {

	// function := ValididateToken(CreateUserHandler)
	router.HandleFunc("/create", CreateUserHandler).Methods("POST")
	router.HandleFunc("/login", CreateUserHandler).Methods("POST")

	return nil
}
