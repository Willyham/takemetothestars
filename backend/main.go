package main

import (
	"github.com/Willyham/takemetothestars/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	r := gin.Default()

	logger, _ := zap.NewProduction()
	candService := service.StubCandidates{}

	hand := handler{
		logger:           logger,
		candidateService: candService,
	}

	r.GET("/candidates", hand.findCandidates)

	r.Run()
}
