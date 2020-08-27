package database2

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"strconv"
)


type ReadModelNullable struct {
	IdExternalApp sql.NullInt64  `json:"ExternalAppID"`
	IdProduct     sql.NullString `json:"ID"`
	IdSuperUser   sql.NullInt64  `json:"SuperUserID"`
	Description   sql.NullString `json:"Description"`
	CustomerMid   sql.NullInt64  `json:"CustomerMid"`
	CustomerEmail sql.NullString `json:"CustomerEmail"`
	Activated     sql.NullBool   `json:"Activated"`
}

type ReadModuleResult struct {
	IdExternalApp string `json:"ExternalAppID"`
	IdProduct     string `json:"ID"`
	IdSuperUser   string `json:"SuperUserID"`
	Description   string `json:"Description"`
	CustomerMid   string `json:"CustomerMid"`
	CustomerEmail string `json:"CustomerEmail"`
	Activated     string `json:"Activated"`
}

func GetAllRows () ([]ReadModuleResult, error) {
	var err error = nil
	conn, err := StartConnection()
	if err != nil {
		logrus.Infof(err.Error())
		return nil, err
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			logrus.Infof(err.Error())
		}
	}()

	sqlStatement := `SELECT id_externalApp, id_product, id_superUser, description, customer_mid, customer_email, activated
			FROM model;`
	rows, err := conn.Query(sqlStatement)
	if err != nil {
		logrus.Infof(err.Error())
		return nil, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			logrus.Infof(err.Error())
		}
	}()
	var readModels []ReadModuleResult
	for rows.Next() {
		var readModelNullable = ReadModelNullable{}
		err = rows.Scan(&readModelNullable.IdExternalApp, &readModelNullable.IdProduct,
			&readModelNullable.IdSuperUser, &readModelNullable.Description,
			&readModelNullable.CustomerMid, &readModelNullable.CustomerEmail,
			&readModelNullable.Activated)
		if err != nil {
			logrus.Infof(err.Error())
			return nil, err
		}
		readModels = append(readModels, validateValues(readModelNullable))
	}
	return readModels, err
}


func validateValues (values ReadModelNullable) ReadModuleResult {
	result := ReadModuleResult{}
	if values.IdExternalApp.Valid {
		result.IdExternalApp = strconv.Itoa(int(values.IdExternalApp.Int64))
	}else{
		result.IdExternalApp = "Nenhum ExternalApp Encontrado"
	}
	if values.IdProduct.Valid {
		result.IdProduct = values.IdProduct.String
	}else{
		result.IdProduct = "Nenhum produto Econtrado"
	}
	if values.IdSuperUser.Valid {
		result.IdSuperUser = strconv.Itoa(int(values.IdSuperUser.Int64))
	}else{
		result.IdSuperUser = "Produto nao avaliado."
	}
	if values.Description.Valid {
		result.Description = values.Description.String
	}else{
		result.Description = "Sem descricao"
	}
	if values.CustomerMid.Valid {
		result.CustomerMid = strconv.Itoa(int(values.CustomerMid.Int64))
	}else{
		result.CustomerMid = "Nenhum customer encotnrado"
	}
	if values.CustomerEmail.Valid {
		result.CustomerEmail = values.CustomerEmail.String
	}else{
		result.CustomerEmail = "Nenhum email cadastrado"
	}
	if values.Activated.Valid {
		if values.Activated.Bool {
			result.Activated = "Produto Aprovado"
		}else{
			result.Activated = "Produto Recusado"
		}
	}else{
		result.Activated = "Produto nao avaliado"
	}
	return result
}
