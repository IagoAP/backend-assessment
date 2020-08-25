package database

import (
	"database/sql"
	"github.com/sirupsen/logrus"
)

func CreateToken(id uint64, token string, expirationTime string) error {
	conn := StartConnection()
	defer conn.Close()

	sqlStatement := `
		INSERT INTO tokens(id_entities, token, expiration_time) 
		VALUES ($1,$2,$3)`
	err := conn.QueryRow(sqlStatement, id, token, expirationTime)
	if err != nil {
		return err.Err()
	}
	return nil
}

func GetTime(token string) (uint64, string, error) {
	conn := StartConnection()
	defer conn.Close()

	sqlStatement := `SELECT id_entities, expiration_time FROM tokens WHERE token=$1;`
	row := conn.QueryRow(sqlStatement, token)
	var expirationTime string
	var id uint64
	err := row.Scan(&id, &expirationTime)
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
