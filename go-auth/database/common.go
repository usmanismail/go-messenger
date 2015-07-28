package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("database")

func Connect(provider string, user string, password string, dbHost string,
	dbPort int, dbname string) (UserData, TokenData, error) {

	db, _ := sql.Open(provider, fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, dbHost, dbPort, dbname))
	err := db.Ping()
	if err != nil {
		return nil, nil, err
	} else {
		log.Debug("Connected to DB %s:%d/%s", dbHost, dbPort, dbname)
		userData := UserDataS{db}
		tokenData := TokenDataS{db}

		err := userData.Init()
		if err != nil {
			return nil, nil, err
		}

		err = tokenData.Init()
		if err != nil {
			return nil, nil, err
		}
		return &userData, &tokenData, nil
	}
}
