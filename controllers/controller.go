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

	var data models.Status

	json.Unmarshal(bytes, &data)

	waterStatus, windStatus := helpers.GetWaterStatus(data.Status.Water), helpers.GetWindStatus(data.Status.Wind)
	waterClass, winedClass := helpers.GetWaterClass(data.Status.Water), helpers.GetWindClass(data.Status.Wind)

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"water":       data.Status.Water,
		"waterStatus": waterStatus,
		"waterClass":  waterClass,
		"wind":        data.Status.Wind,
		"windStatus":  windStatus,
		"windClass":   winedClass,
	})
}
