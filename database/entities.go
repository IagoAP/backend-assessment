package database

type entities struct {
	ID       uint64
	EType    string
	Username string
	Password string
}

func ValidateLogin(username, password string) (bool, uint64, string) {
	conn := StartConnection()
	defer conn.Close()

	// sqlStatement := `SELECT id, user_type, username, password FROM entities WHERE username=$1;`
	sqlStatement := `SELECT id, user_type, username, password FROM entities WHERE username=$1;`
	var entity entities
	row := conn.QueryRow(sqlStatement, username)
	err := row.Scan(&entity.ID, &entity.EType, &entity.Username, &entity.Password)
	if err == nil && entity.Password == password {
		return true, entity.ID, entity.EType
	}
	return false, 0, ""
}

func GetType(id uint64) (string, error) {
	conn := StartConnection()
	defer conn.Close()

	sqlStatement := `SELECT user_type FROM entities WHERE id=$1;`
	var userType string
	row := conn.QueryRow(sqlStatement, id)
	err := row.Scan(&userType)
	if err != nil {
		return "", err
	}
	return userType, nil
}
