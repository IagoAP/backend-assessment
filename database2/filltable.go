package database2

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
)

type ReadModel struct {
	IdExternalApp uint64 `json:"ExternalAppID"`
	IdProduct     string `json:"ID"`
	IdSuperUser   uint64 `json:"SuperUserID"`
	Description   string `json:"Description"`
	CustomerMid   int    `json:"CustomerMid"`
	CustomerEmail string `json:"CustomerEmail"`
	Activated     bool   `json:"Activated"`
}

func CreateRow(msg ReadModel) error {
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
		INSERT INTO model (id_product, id_externalApp, description, customer_mid, customer_email) 
		VALUES ($1,$2,$3,$4, $5)`
	err = conn.QueryRow(sqlStatement, msg.IdProduct, msg.IdExternalApp, msg.Description, msg.CustomerMid, msg.CustomerEmail).Err()
	if err != nil {
		logrus.Infof(err.Error())
	}
	return err
}

func UpdateRow (msg ReadModel) error {
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

	sqlStatement := `UPDATE model
		SET id_superUser = $2, activated = $3
		WHERE id_product = $1;`
	err = conn.QueryRow(sqlStatement, msg.IdProduct, msg.IdSuperUser, msg.Activated).Err()
	if err != nil {
		logrus.Infof(err.Error())
	}
	return err
}

func ConvertReadModel(msg []byte) (ReadModel, error) {
	result  :=  ReadModel{}
	err := json.Unmarshal(msg, &result)
	if err != nil {
		logrus.Infof(err.Error())
	}
	return result, err
}
