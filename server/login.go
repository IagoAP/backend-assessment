package server

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"psT10/database"
	"psT10/environment"
	"psT10/tokenlist"
	"time"
)

type LoginRequest struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

func (s *Server) Login(c *gin.Context) {
	var loginRequest LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	login, id, _, err := database.ValidateLogin(loginRequest.Username, loginRequest.Password)
	if err != nil || !login {
		logrus.Info(err.Error())
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}
	if token, err := CreateToken(id); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{"access_token": token})
	}
}

func createJwt(id uint64) (jwt.MapClaims, string) {
	jwtCreated := jwt.MapClaims{}
	jwtCreated["authorized"] = true
	jwtCreated["user_id"] = id
	var expirationTime = time.Now().Add(time.Minute * 15).Format(time.RFC3339)
	jwtCreated["exp"] = expirationTime
	return jwtCreated, expirationTime
}

func CreateToken(id uint64) (string, error) {
	createdJwt, expirationTime := createJwt(id)
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, createdJwt)
	token, err := at.SignedString([]byte(environment.GetEnvVariables("ACCESS_SECRET")))
	tokenlist.AddToken(id, token, expirationTime)
	if err != nil {
		return "", err
	}
	return token, nil
}
