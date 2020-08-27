package database

import "github.com/sirupsen/logrus"

type entities struct {
	ID       uint64
	EType    string
	Username string
	Password string
}

func ValidateLogin(username, password string) (bool, uint64, string, error) {
	var err error = nil
	conn, err := StartConnection()
	if err != nil {
		logrus.Info(err.Error())
		return false, 0, "", err
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			logrus.Infof(err.Error())
		}
	}()

	sqlStatement := `SELECT id, user_type, username, password FROM entities WHERE username=$1;`
	var entity entities
	row := conn.QueryRow(sqlStatement, username)
	err = row.Scan(&entity.ID, &entity.EType, &entity.Username, &entity.Password)
	if err == nil && entity.Password == password {
		return true, entity.ID, entity.EType, nil
	}
	return false, 0, "", err
}

func GetType(id uint64) (string, error) {
	var err error = nil
	conn, err := StartConnection()
	if err != nil {
		logrus.Info(err.Error())
		return "", err
	}

	defer func() {
		err = conn.Close()
		if err != nil {
			logrus.Infof(err.Error())
		}
	}()

	sqlStatement := `SELECT user_type FROM entities WHERE id=$1;`
	var userType string
	row := conn.QueryRow(sqlStatement, id)
	err = row.Scan(&userType)
	if err != nil {
		logrus.Infof(err.Error())
		return "", err
	}
	return userType, err
}
