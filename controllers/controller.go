package controllers

import (
	"assignment-3/models"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStatus(ctx *gin.Context) {
	models.Data[0].Water = uint(rand.Intn(20))
	models.Data[0].Wind = uint(rand.Intn(20))

	status := models.Status{
		Water: models.Data[0].Water,
		Wind:  models.Data[0].Wind,
	}

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"water": status.Water,
		"wind":  status.Wind,
	})
}
