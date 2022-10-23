package endpoints

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
)

type HandleFunc func(http.ResponseWriter, *http.Request)

func ValididateToken(handler HandleFunc) HandleFunc {

	return func(writer http.ResponseWriter, request *http.Request) {

		handler(writer, request)

	}
}

var (
	_, filename = path.Split(os.Args[0])
	HTTPLogger  = log.New(os.Stdout, filename+" ", log.LstdFlags)
)

//InternalServerError 500 provide bad request
func InternalServerError(writer http.ResponseWriter, request *http.Request, err error) {

	writer.Header().Set("Content-Type", "text/plain")
	writer.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(writer, "internal_server_error")
	HTTPLogger.Printf("ERROR %v %v [500] internal_server_error %v\n", request.RemoteAddr, request.RequestURI, err)

}

//BadRequest provide  400 bad request
func BadRequest(writer http.ResponseWriter, request *http.Request, message error) {
	writer.Header().Set("Content-Type", "text/plain")
	writer.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(writer, message)
	HTTPLogger.Printf("INFO %v %v [400] %v\n", request.RemoteAddr, request.RequestURI, message)
}

//AccessDenied 401 access denied response
func AccessDenied(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/plain")
	writer.WriteHeader(http.StatusUnauthorized)
	HTTPLogger.Printf("ERROR %v %v [401] access_denied\n", request.RemoteAddr, request.RequestURI)
}

//StatusOK 200 response
func StatusOK(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/plain")
	writer.WriteHeader(http.StatusOK)
	HTTPLogger.Printf("ERROR %v %v [200]\n", request.RemoteAddr, request.RequestURI)

}

//JSONResponse provide json response 200
func JSONResponse(writer http.ResponseWriter, request *http.Request, value interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(value)
	if err != nil {
		HTTPLogger.Printf("INFO %v %v [200] Unable to Marshal value\n", request.RemoteAddr, request.RequestURI)
		return
	}
	HTTPLogger.Printf("INFO %v %v [200]\n", request.RemoteAddr, request.RequestURI)
}
