package user

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type UserDataTest struct {
}

func TestRegisterOk(t *testing.T) {

	status := RegisterUser(&UserDataTest{}, "new", "password")
	assert.Equal(t, 200, status, "Expected 200 OK")
}

func TestRegisterInternalError(t *testing.T) {

	status := RegisterUser(&UserDataTest{}, "error-user", "password")
	assert.Equal(t, 500, status, "Expected 500 Internal Error")
}

func TestRegisterConflict(t *testing.T) {

	status := RegisterUser(&UserDataTest{}, "old", "password")
	assert.Equal(t, 409, status, "Expected 409 Conflict")
}

func (userData *UserDataTest) Init() error {
	return nil
}

func (userData *UserDataTest) DeleteUser(userId string) error {
	return nil
}

func (userData *UserDataTest) SaveUser(userId string, passwordHash []byte) error {
	if userId == "error-user" {
		return errors.New("Some Error")
	} else {
		return nil
	}
}

func (userData *UserDataTest) GetUser(userId string) ([]byte, error) {
	if userId == "new" {
		return nil, errors.New("Not Found")
	} else if userId == "old" {
		return []byte("someHash"), nil
	} else {
		return nil, nil
	}
}
