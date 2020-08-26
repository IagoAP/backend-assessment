package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Server struct {
	Kafka Kafka
}

func (s *Server) Run() {
	r := gin.Default()
	r.POST("/RequestToken", s.Login)
	r.POST("/IssueProductActivation", s.IssueProductActivation)
	r.POST("/ApproveActivation", s.ApproveActivation)
	r.POST("/RejectActivation", s.RejectActivation)
	r.GET("/ActivationRequests", s.ActivationRequests)
	var err = r.Run()
	if err != nil {
		logrus.Fatal(err.Error())
	}
}
