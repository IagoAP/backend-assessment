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
	conn := StartConnection()
	defer conn.Close()

	sqlStatement := `UPDATE product
		SET superuser_id = $2, activated = $3
		WHERE id = $1;`
	err := conn.QueryRow(sqlStatement, msg.ActivationID, msg.SuperUserID, msg.Activated).Err()
	if err != nil {
		return err
	}
	return nil
}

func ConvertActivationMessage(msg []byte) ActivationRequest {
	result  :=  ActivationRequest{}
	err := json.Unmarshal(msg, &result)
	if err != nil {
		logrus.Fatal(err.Error())
	}
	return result
}


