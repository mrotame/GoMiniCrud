package validators

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestValidateEmail(t *testing.T) {
	valid_email := "test@test.com"
	invalid_email_1 := "test"
	invalid_email_2 := "test.test"
	invalid_email_3 := "test@"
	w := *httptest.NewRecorder()

	if !ValidateEmail(&w, valid_email) {
		t.Error(fmt.Sprintf("Error, ValidateEmail returned false for valid email `%s`", valid_email))
	}
	if ValidateEmail(&w, invalid_email_1) {
		t.Error(fmt.Sprintf("Error, ValidateEmail returned true for invalid email 1 `%s`", invalid_email_1))
	}
	if ValidateEmail(&w, invalid_email_2) {
		t.Error(fmt.Sprintf("Error, ValidateEmail returned true for invalid email 2 `%s`", invalid_email_2))
	}
	if ValidateEmail(&w, invalid_email_3) {
		t.Error(fmt.Sprintf("Error, ValidateEmail returned true for invalid email 3 `%s`", invalid_email_3))
	}
}

func TestValidateJsonDecode(t *testing.T) {
	w := *httptest.NewRecorder()
	err := errors.New("test error")

	if !ValidateJsonDecode(&w, nil) {
		t.Error("Error, validator returned false for `nil` err")
	}
	if ValidateJsonDecode(&w, err) {
		t.Error(fmt.Sprintf("Error, validator returned true. expected false for error `%v`", err))
	}

	if w.Result().StatusCode != http.StatusBadRequest {
		t.Error(fmt.Sprintf("Error, response status code is not `%v` for error `%v`", http.StatusBadRequest, err))
	}
}

func TestValidateModelFound(t *testing.T) {
	w := *httptest.NewRecorder()
	err := errors.New("test error")

	if !ValidateModelFound(&w, nil, "") {
		t.Error("Error, validator returned false for `nil` err")
	}

	if ValidateModelFound(&w, err, err.Error()) {
		t.Error(fmt.Sprintf("Error, validator returned true. expected false for error `%v`", err))
	}

	if strings.TrimSpace(w.Body.String()) != err.Error() {
		t.Error(fmt.Sprintf("Error, message written by validator `%v` is different from custom message `%v` ", w.Body, err.Error()))
	}

	if w.Result().StatusCode != http.StatusNotFound {
		t.Error(fmt.Sprintf("Error, status code defined by validator is not `%v`", http.StatusNotFound))
	}
}

func TestValidateSavedModel(t *testing.T) {
	w := *httptest.NewRecorder()
	err := errors.New("test error")

	if !ValidateSavedModel(&w, nil) {
		t.Error("Error, validator returned false for `nil` err")
	}

	if ValidateSavedModel(&w, err) {
		t.Error(fmt.Sprintf("Error, validator returned true. expected false for error `%v`", err))
	}

	if strings.TrimSpace(w.Body.String()) != err.Error() {
		t.Error(fmt.Sprintf("Error, message written by validator `%v` is different from custom message `%v` ", w.Body, err.Error()))
	}

	if w.Result().StatusCode != http.StatusBadRequest {
		t.Error(fmt.Sprintf("Error, status code defined by validator is not `%v`", http.StatusBadRequest))
	}
}

func TestGenericValidator(t *testing.T) {
	w := *httptest.NewRecorder()
	err := errors.New("test error")

	if !GenericValidator(&w, nil, http.StatusForbidden) {
		t.Error("Error, generic validator returned false for `nil` err")
	}

	if GenericValidator(&w, err, http.StatusForbidden) {
		t.Error(fmt.Sprintf("Error, generic validator returned true. expected false for error `%v`", err))
	}

	if strings.TrimSpace(w.Body.String()) != err.Error() {
		t.Error(fmt.Sprintf("Error, message written by generic validator `%v` is different from custom message `%v` ", w.Body, err.Error()))
	}

	if w.Result().StatusCode != http.StatusForbidden {
		t.Error(fmt.Sprintf("Error, status code defined by generic validator is not `%v`", http.StatusForbidden))
	}

}
