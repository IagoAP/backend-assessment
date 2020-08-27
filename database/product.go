package database

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
)

type ProductRequest struct {
	ID   		  string `json:"ID"`
	Description   string `json:"Description"`
	CustomerMid   int    `json:"CustomerMid"`
	CustomerEmail string `json:"CustomerEmail"`
	ExternalAppID uint64 `json:"ExternalAppID"`
	SuperuserID   uint64 `json:"SuperuserID"`
	Activated     bool   `json:"Activated"`
}

func CreateProduct(msg ProductRequest) error {
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
		INSERT INTO product(id, description, customer_mid, customer_email, externalApp_id) 
		VALUES ($1,$2,$3,$4, $5)`
	err = conn.QueryRow(sqlStatement, msg.ID, msg.Description, msg.CustomerMid, msg.CustomerEmail, msg.ExternalAppID).Err()
	if err != nil {
		return err
	}
	return err
}

func ConvertProductMessage(msg []byte) (ProductRequest, error) {
	result  :=  ProductRequest{}
	err := json.Unmarshal(msg, &result)
	if err != nil {
		logrus.Error(err.Error())
	}
	return result, err
}

