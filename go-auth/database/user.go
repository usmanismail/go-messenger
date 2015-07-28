package database

import (
	"database/sql"
)

type UserData interface {
	Init() error
	SaveUser(string, []byte) error
	DeleteUser(string) error
	GetUser(string) ([]byte, error)
}

type UserDataS struct {
	db *sql.DB
}

func (userData *UserDataS) Init() error {
	// Create Table if Missing
	_, err := userData.db.Exec(
		`CREATE TABLE IF NOT EXISTS user (
			userid VARCHAR(20) NOT NULL,
			password VARBINARY(60) NOT NULL,
			PRIMARY KEY (userid)
		)`)
	if err != nil {
		return err
	} else {
		log.Debug("Created User Table")
		return nil
	}
}

func (userData *UserDataS) DeleteUser(userId string) error {
	_, err := userData.db.Exec(`DELETE from user where userid = ?`, userId)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (userData *UserDataS) SaveUser(userId string, passwordHash []byte) error {

	_, err := userData.db.Exec(`INSERT INTO user Values (?,?)`, userId, passwordHash)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (userData *UserDataS) GetUser(userId string) ([]byte, error) {

	var passwordHash []byte
	err := userData.db.QueryRow("select password from user where userid = ?", userId).Scan(&passwordHash)
	if err != nil {
		log.Info("Error %s", err.Error())
		return nil, err
	} else {
		log.Info("User Exists %s", userId)
		return passwordHash, nil
	}
}
