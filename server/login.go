package server

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"psT10/database"
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
	login, id, userType := database.ValidateLogin(loginRequest.Username, loginRequest.Password)
	if !login {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}
	if token, err := CreateToken(id, userType); err != nil {
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

func CreateToken(id uint64, userType string) (string, error) {
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") // FAZER ENV
	createdJwt, expirationTime := createJwt(id)
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, createdJwt)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	tokenlist.AddToken(id, token, expirationTime)
	if err != nil {
		return "", err
	}
	return token, nil
}
