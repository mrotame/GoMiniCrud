package validators

import (
	"net/http"
	"net/mail"
)

func ValidateEmail(w http.ResponseWriter, email string) bool {
	if email == "" {
		http.Error(w, "email field is required", http.StatusBadRequest)
		return false
	}
	_, err := mail.ParseAddress(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return false
	}
	return true
}

func ValidateJsonDecode(w http.ResponseWriter, err error) bool {

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return false
	}
	return true
}

func ValidateModelFound(w http.ResponseWriter, err error, customMessage string) bool {
	if err != nil {

		http.Error(w, customMessage, http.StatusNotFound)
		return false
	}
	return true
}

func ValidateSavedModel(w http.ResponseWriter, err error) bool {
	return GenericValidator(w, err, http.StatusBadRequest)
}

func GenericValidator(w http.ResponseWriter, err error, httpStatus int) bool {

	if err != nil {
		http.Error(w, err.Error(), httpStatus)
		return false
	}
	return true
}
