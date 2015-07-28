package database

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"net/http"
	"time"
)

type TokenData interface {
	Init() error
	CreateToken(string) (string, error)
	ValidateToken(string, string) (int, error)
}

type TokenDataS struct {
	db *sql.DB
}

func (tokenData *TokenDataS) Init() error {
	// Create Table if Missing
	_, err := tokenData.db.Exec(
		`CREATE TABLE IF NOT EXISTS token (
			username VARCHAR(20) NOT NULL,
			token VARBINARY(256) NOT NULL,
			expiry TIMESTAMP NOT NULL,
			PRIMARY KEY (username)
		)`)
	if err != nil {
		return err
	} else {
		log.Debug("Created Token Table")
		return nil
	}
}

func (tokenData *TokenDataS) CreateToken(userId string) (string, error) {

	token := make([]byte, 256)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}

	timestamp := time.Now().Add(time.Hour).UTC().Format("2006-01-02 15:04:05")

	_, err = tokenData.db.Exec(`REPLACE INTO token Values (?,?,?)`, userId, token, timestamp)
	if err != nil {
		log.Debug("Error %s", err.Error())
		return "", err
	} else {
		return base64.StdEncoding.EncodeToString(token), nil
	}
}

func (tokenData *TokenDataS) ValidateToken(token string, username string) (int, error) {

	var usernameDB string
	var expiry string

	tokenBytes, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		log.Info("Error %s", err.Error())
		return http.StatusBadRequest, err
	}

	err = tokenData.db.QueryRow("select username,expiry from token where token.token = ?", tokenBytes).Scan(&usernameDB, &expiry)
	if err != nil {
		log.Info("Error %s", err.Error())
		return http.StatusBadRequest, err
	}

	if usernameDB != username {
		log.Info("Token for wrong user")
		return http.StatusBadRequest, err
	}

	t, _ := time.Parse("2006-01-02 15:04:05", expiry)
	log.Debug("Expiry %d Now %d", t.Unix(), time.Now().Unix())
	if t.Before(time.Now()) {
		log.Debug("Token Expired")
		return http.StatusBadRequest, err
	}
	return http.StatusOK, nil
}
