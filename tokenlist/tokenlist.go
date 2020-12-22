package tokenlist

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"psT10/database"
	"time"
)

func AddToken(id uint64, tokenValue string, expirationTimeValue string) {
	err := database.CreateToken(id, tokenValue, expirationTimeValue)
	if err != nil {
		logrus.Fatal(err.Error())
	}
}

func isValid(token string) (bool, uint64) {
	id, tokenExpirationTime, err := database.GetTime(token)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, 0
		}
		logrus.Fatal(err.Error())
	}
	now := time.Now()
	expiration, err := time.Parse(time.RFC3339, tokenExpirationTime)
	if err != nil {
		logrus.Fatal(err.Error())
	}
	return now.Before(expiration), id
}

func CheckToken(userToken string) (bool, uint64, string) {
	valid, id := isValid(userToken)
	if valid {
		userType, err := database.GetType(id)
		if err != nil {
			logrus.Fatal(err.Error())
		} else {
			return valid, id, userType
		}
	}
	return valid, 0, ""
}
