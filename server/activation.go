package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"psT10/tokenlist"
)

type ActivationRequest struct {
	SuperUserID  uint64
	ActivationID uint64 `json:"ActivationID"`
	Activated    bool
}

func (s *Server) ApproveActivation(c *gin.Context) {
	validToken, id, userType := tokenlist.CheckToken(c.Request.Header["Token"][0])
	if !validToken {
		c.JSON(http.StatusUnauthorized, "Token expired")
		return
	}
	if userType != "SuperUser" {
		c.JSON(http.StatusUnauthorized, "You don't have access rights")
		return
	}
	var activationRequest ActivationRequest
	if err := c.ShouldBindJSON(&activationRequest); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	activationRequest.SuperUserID = id
	activationRequest.Activated = true
	err := sendMessagesActivated(s.Kafka, activationRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	}
}

func (s *Server) RejectActivation(c *gin.Context) {
	validToken, id, userType := tokenlist.CheckToken(c.Request.Header["Token"][0])
	if !validToken {
		c.JSON(http.StatusUnauthorized, "Token expired")
		return
	}
	if userType != "SuperUser" {
		c.JSON(http.StatusUnauthorized, "You don't have access rights")
		return
	}
	var activationRequest ActivationRequest
	if err := c.ShouldBindJSON(&activationRequest); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	activationRequest.SuperUserID = id
	activationRequest.Activated = false
	err := sendMessagesActivated(s.Kafka, activationRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	}
}

func sendMessagesActivated(kf Kafka, msg ActivationRequest) error {
	err := kf.SendMessage(msg, "ProductActivation")
	if err != nil {
		return err
	}
	err = kf.SendMessage(msg, "ProductActivationReadDB")
	if err != nil {
		return err
	}
	return nil
}
