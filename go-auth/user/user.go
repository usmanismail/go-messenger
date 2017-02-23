package user

import (
	"net/http"

	"github.com/op/go-logging"
	"github.com/usmanismail/go-messenger/go-auth/database"
	"golang.org/x/crypto/bcrypt"
)

var log = logging.MustGetLogger("user")

type User interface {
	GetUserId() string
	GetPasswordHash() []byte
}

type UserS struct {
	userId       string
	passwordHash []byte
}

func NewUser(userId string, password string) (User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	} else {
		return &UserS{userId, hash}, nil
	}
}

func GetUser(userId string, passwordHash []byte) User {
	return &UserS{userId, passwordHash}
}

func RegisterUser(userData database.UserData, userId string, password string) int {
	if len(userId) > 0 && len(password) > 0 {
		userObj, _ := NewUser(userId, password)
		// Check if user already exists
		user, err := userData.GetUser(userObj.GetUserId())
		if user != nil {
			log.Warning("Duplicate username tried: %s", userId)
			return http.StatusConflict
		}
		// Save User to database
		err = userData.SaveUser(userObj.GetUserId(), userObj.GetPasswordHash())
		if err != nil {
			log.Warning("Unable to create user: %s", err.Error())
			return http.StatusInternalServerError
		} else {
			return http.StatusOK
		}
	} else {
		log.Debug("Unable to create user, username or password missing")
		return http.StatusBadRequest
	}
}

func DeleteUser(userData database.UserData, userId string, password string) int {
	if len(userId) > 0 && len(password) > 0 {
		passwordHash, _ := userData.GetUser(userId)
		if passwordHash == nil {
			log.Warning("User not found, cannot delete: %s", userId)
			return http.StatusNotFound
		} else {
			if CompareHashAndPassword(passwordHash, password) == nil {
				userData.DeleteUser(userId)
				return http.StatusOK
			} else {
				return http.StatusBadRequest
			}

		}
	} else {
		log.Debug("Unable to delete user, username or password missing")
		return http.StatusBadRequest
	}
}

func CompareHashAndPassword(hashedPassword []byte, password string) error {
	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		log.Warning("Invalid password: %s", err.Error())
		return err
	} else {
		return nil
	}
}

func (u *UserS) GetUserId() string {
	return u.userId
}

func (u *UserS) GetPasswordHash() []byte {
	return u.passwordHash
}
