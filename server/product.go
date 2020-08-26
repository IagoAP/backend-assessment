package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"psT10/tokenlist"
)

type ProductRequest struct {
	ExternalAppID uint64
	Description   string `json:"Description"`
	CustomerMid   int    `json:"CustomerMid"`
	CustomerEmail string `json:"CustomerEmail"`
}

func (s *Server) IssueProductActivation(c *gin.Context) {
	validToken, id, userType := tokenlist.CheckToken(c.Request.Header["Token"][0])
	if !validToken {
		c.JSON(http.StatusUnauthorized, "Token expired")
		return
	}
	if userType != "ExternalApp" {
		c.JSON(http.StatusUnauthorized, "You don't have access rights")
		return
	}
	var productRequest ProductRequest
	if err := c.ShouldBindJSON(&productRequest); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	productRequest.ExternalAppID = id
	err := s.Kafka.SendMessage(productRequest, "ProductCreate")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	}
}

func (s *Server) ActivationRequests(c *gin.Context) {
	validToken, _, _ := tokenlist.CheckToken(c.Request.Header["Token"][0])
	if !validToken {
		c.JSON(http.StatusUnauthorized, "Token expired")
		return
	}

}
