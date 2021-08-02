package main

import (
	"net/http"
	"strconv"

	"github.com/Willyham/takemetothestars/service"
	"github.com/gin-gonic/gin"
	"github.com/go-spatial/geom"
	"go.uber.org/zap"
)

type handler struct {
	logger           *zap.Logger
	candidateService service.CandidateService
}

func (h handler) findCandidates(c *gin.Context) {
	limit := c.DefaultQuery("limit", "5")
	limitNum, err := strconv.Atoi(limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit value"})
		return
	}

	lat := c.Query("lat")
	latNum, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid latitude value"})
		return
	}

	long := c.Query("long")
	longNum, err := strconv.ParseFloat(long, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid longitude value"})
		return
	}

	origin := geom.Point{latNum, longNum}
	candidates, err := h.candidateService.FindCandidates(c, origin, limitNum)
	if err != nil {
		h.logger.Error("failed to compute candidate points", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to compute candidates"})
		return
	}

	c.JSON(http.StatusOK, candidates)
}
