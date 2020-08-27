package database

import (
	"database/sql"
	"github.com/sirupsen/logrus"
)

func CreateToken(id uint64, token string, expirationTime string) error {
	var err error = nil
	conn, err := StartConnection()
	if err != nil {
		logrus.Infof(err.Error())
		return err
	}

	defer func() {
		err = conn.Close()
		if err != nil {
			logrus.Infof(err.Error())
		}
	}()

	sqlStatement := `
		INSERT INTO tokens(id_entities, token, expiration_time) 
		VALUES ($1,$2,$3)`
	error := conn.QueryRow(sqlStatement, id, token, expirationTime)
	err = error.Err()
	if err != nil {
		logrus.Infof(err.Error())
	}
	return err
}

func GetTime(token string) (uint64, string, error) {
	var err error = nil
	conn, err := StartConnection()
	if err != nil {
		logrus.Infof(err.Error())
		return 0, "", err
	}

	defer func() {
		err = conn.Close()
		if err != nil {
			logrus.Infof(err.Error())
		}
	}()

	sqlStatement := `SELECT id_entities, expiration_time FROM tokens WHERE token=$1;`
	row := conn.QueryRow(sqlStatement, token)
	var expirationTime string
	var id uint64
	err = row.Scan(&id, &expirationTime)
	switch err {
	case sql.ErrNoRows:
		return 0, "", err
	case nil:
		return id, expirationTime, nil
	default:
		logrus.Fatal(err.Error())
		return 0, "", err
	}
}
