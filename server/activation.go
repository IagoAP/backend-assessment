package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"psT10/tokenlist"
)

type ActivationRequest struct {
	SuperUserID  uint64
	ActivationID string `json:"ID"`
	Activated    bool
}

func (s *Server) ApproveActivation(c *gin.Context) {
	activate(c, s.Kafka,true)
}

func (s *Server) RejectActivation(c *gin.Context) {
	activate(c, s.Kafka, false)
}

func activate(c *gin.Context, kf Kafka, action bool) {
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
	activationRequest.Activated = action
	err := kf.SendMessage(activationRequest, "ProductActivation")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	}
}