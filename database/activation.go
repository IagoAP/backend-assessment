package database

import (
"encoding/json"
"github.com/sirupsen/logrus"
)

type ActivationRequest struct {
	SuperUserID  uint64 `json:"SuperUserID"`
	ActivationID string `json:"ID"`
	Activated    bool   `json:"Activated"`
}

func ActivateProduct(msg ActivationRequest) error {
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

	sqlStatement := `UPDATE product
						SET superuser_id = $2, activated = $3
						WHERE id = $1;`
	err = conn.QueryRow(sqlStatement, msg.ActivationID, msg.SuperUserID, msg.Activated).Err()
	if err != nil {
		logrus.Infof(err.Error())
		return err
	}
	return nil
}

func GetEmail(msg ActivationRequest) (string, error) {
	var err error = nil
	conn, err := StartConnection()
	if err != nil {
		logrus.Infof(err.Error())
		return "", err
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			logrus.Infof(err.Error())
		}
	}()

	var email string

	sqlStatement := `SELECT customer_email 
		FROM product
		WHERE id = $1;`
	err = conn.QueryRow(sqlStatement, msg.ActivationID).Scan(&email)
	if err != nil {
		logrus.Infof(err.Error())
		return "", err
	}
	return email, err
}

func ConvertActivationMessage(msg []byte) (ActivationRequest, error) {
	result  :=  ActivationRequest{}
	err := json.Unmarshal(msg, &result)
	if err != nil {
		logrus.Infof(err.Error())
	}
	return result, err
}


