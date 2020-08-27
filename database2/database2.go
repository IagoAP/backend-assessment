package database2

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // sem esse import nao estabelece conexao
	"github.com/sirupsen/logrus"
	"psT10/environment"
)

type PostgresInfo struct {
	postgresHost     string
	postgresPort     string
	postgresUser     string
	postgresPassword string
	postgresDBName   string
}

func defineDB() string {
	postgreConfig := PostgresInfo{
		postgresPort:     environment.GetEnvVariables("DB_PORT"),
		postgresHost:     environment.GetEnvVariables("DB_HOST"),
		postgresUser:     environment.GetEnvVariables("DB_USER"),
		postgresPassword: environment.GetEnvVariables("DP_PASSWORD"),
		postgresDBName:   environment.GetEnvVariables("DB2_NAME"),
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		postgreConfig.postgresHost,
		postgreConfig.postgresPort,
		postgreConfig.postgresUser,
		postgreConfig.postgresPassword,
		postgreConfig.postgresDBName)
	return psqlInfo
}

func StartConnection() (*sql.DB, error) {
	var err error = nil
	var Conn *sql.DB
	Conn, err = sql.Open("postgres", defineDB())

	if err != nil {
		logrus.Infof(err.Error())
	}
	err = Conn.Ping()
	if err != nil {
		logrus.Fatal(err.Error())
	}

	return Conn, err
}
