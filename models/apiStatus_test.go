package models

import (
	"testing"
)

func TestApiOk(t *testing.T) {
	ok := ApiOk{}
	if _, ok := interface{}(&ok).(ApiStatus); !ok {
		t.Errorf("Expected ApiOk to implement interface ApiStatus")
	}
}

func TestApiOk_Error(t *testing.T) {
	ok := ApiOk{}
	err := ok.Status()
	if err.Code != 404 {
		t.Errorf("Expected ApiOk to have status 200")
	}
}

func TestApiNotFoundError(t *testing.T) {
	nfe := ApiNotFoundError{}
	if _, ok := interface{}(&nfe).(ApiStatus); !ok {
		t.Errorf("Expected ApiNotFoundError to implement interface ApiStatus")
	}
}

func TestApiNotFoundError_Error(t *testing.T) {
	nfe := ApiNotFoundError{}
	err := nfe.Status()
	if err.Code != 404 {
		t.Errorf("Expected ApiNotFoundError to have status 404")
	}
}

func TestApiForbiddenError(t *testing.T) {
	fe := ApiForbiddenError{}
	if _, ok := interface{}(&fe).(ApiStatus); !ok {
		t.Errorf("Expected ApiForbiddenError to implement interface ApiStatus")
	}
}

func TestApiForbiddenError_Error(t *testing.T) {
	fe := ApiForbiddenError{}
	err := fe.Status()
	if err.Code != 403 {
		t.Errorf("Expected ApiForbiddenError to have status 403")
	}
}

func TestApiBadRequestError(t *testing.T) {
	bre := ApiBadRequestError{}
	if _, ok := interface{}(&bre).(ApiStatus); !ok {
		t.Errorf("Expected ApiBadRequestError to implement interface ApiStatus")
	}
}

func TestApiBadRequestError_Error(t *testing.T) {
	bre := ApiBadRequestError{}
	err := bre.Status()
	if err.Code != 400 {
		t.Errorf("Expected ApiBadRequestError to have status 400")
	}
}

func TestApiInternalServerError(t *testing.T) {
	ise := ApiInternalServerError{}
	if _, ok := interface{}(&ise).(ApiStatus); !ok {
		t.Errorf("Expected ApiInternalServerError to implement interface ApiStatus")
	}
}

func TestApiInternalServerError_Error(t *testing.T) {
	ise := ApiInternalServerError{}
	if ise.Status().Code != 500 {
		t.Errorf("Expected ApiInternalServerError to have status 400")
	}

	ise.Message = "" //default
	if len(ise.Status().Message) == 0 {
		t.Errorf("Expected ApiInternalServerError to have a default error message")
	}

	definedMessage := "USER_DEFINED_MESSAGE"
	ise = ApiInternalServerError{Message: definedMessage}
	if ise.Status().Message != definedMessage {
		t.Errorf("Expected ApiInternalServerError to have a message defined by user")
	}
}
