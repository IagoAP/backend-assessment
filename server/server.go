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
	r.PUT("/IssueProductActivation", s.IssueProductActivation)
	//r.GET("/ActivationRequests", )
	var err = r.Run()
	if err != nil {
		logrus.Fatal(err.Error())
	}
}
