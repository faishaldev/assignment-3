package controllers

import (
	"assignment-3/helpers"
	"assignment-3/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetStatus(ctx *gin.Context) {
	helpers.UpdateJSONData()

	jsonData, err := os.Open("data.json")

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)

		return
	}

	bytes, err := ioutil.ReadAll(jsonData)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)

		return
	}

	var status models.Status

	json.Unmarshal(bytes, &status)

	waterStatus, windStatus := helpers.GetWaterStatus(status.Status.Water), helpers.GetWindStatus(status.Status.Wind)

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"water":       status.Status.Water,
		"waterStatus": waterStatus,
		"wind":        status.Status.Wind,
		"windStatus":  windStatus,
	})
}
