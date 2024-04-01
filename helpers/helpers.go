package helpers

import (
	"coffee-app/services"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
)

type Envelop map[string]interface{}

type Message struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

// MessageLogs is a global variable that holds the loggers
var infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
var errorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

var MessageLogs = &Message{
	InfoLog:  infoLog,
	ErrorLog: errorLog,
}

func ReadJson(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxByte := 1048576
	// read the body of the request
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxByte))
	// decode the body into the data interface
	decode := json.NewDecoder(r.Body)
	// decode the body into the data interface
	err := decode.Decode(data)
	if err != nil {
		return err
	}

	// check if the body has only a single JSON value
	err = decode.Decode(&struct{}{})
	if err != nil {
		return errors.New("body must only have a single JSON value")
	}

	return nil
}

func WriteJson(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	// set the content type to application/json
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// set the status code
	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	// write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func ErrorJSON(w http.ResponseWriter, err error, status ...int) {
	// write the error response
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}

	// write the error response
	var payload services.JsonResponse
	payload.Error = true
	payload.Messge = err.Error()
	WriteJson(w, statusCode, payload)

}
