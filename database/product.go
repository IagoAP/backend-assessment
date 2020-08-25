package database

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
)

type ProductRequest struct {
	Description   string `json:"Description"`
	CustomerMid   int    `json:"CustomerMid"`
	CustomerEmail string `json:"CustomerEmail"`
	ExternalAppID uint64 `json:"ExternalAppID"`
	SuperuserID   uint64 `json:"SuperuserID"`
	Activated     bool   `json:"Activated"`
}

func CreateProduct(msg ProductRequest) error {
	conn := StartConnection()
	defer conn.Close()

	sqlStatement := `
		INSERT INTO product(description, customer_mid, customer_email, externalApp_id) 
		VALUES ($1,$2,$3,$4)
		RETURNING id`
	err := conn.QueryRow(sqlStatement, msg.Description, msg.CustomerMid, msg.CustomerEmail, msg.ExternalAppID)
	if err != nil {
		return err.Err()
	}
	return nil
}

func ConvertProductMessage(msg []byte) ProductRequest {
	result  :=  ProductRequest{}
	err := json.Unmarshal(msg, &result)
	if err != nil {
		logrus.Fatal(err.Error())
	}
	return result
}

//
//func ListProducts() []ProductRequest {
//	//productRequest := []ProductRequest
//
//}
