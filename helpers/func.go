package helpers

import (
	"fmt"
	"log"
	"math/rand"
	"os"
)

func GetWaterStatus(water uint) string {
	switch {
	case water <= 5:
		return "Safe"
	case water <= 8:
		return "Alert"
	default:
		return "Danger"
	}
}

func GetWindStatus(wind uint) string {
	switch {
	case wind <= 6:
		return "Safe"
	case wind <= 15:
		return "Alert"
	default:
		return "Danger"
	}
}

func UpdateJSONData() {
	water, wind := rand.Intn(100), rand.Intn(100)

	jsonString := fmt.Sprintf(`
	{
		"status": {
			"water": %v,
			"wind": %v
		}
	}`, water, wind)

	data, err := os.Create("data.json")

	if err != nil {
		log.Fatal("Error in json update: ", err.Error())

		return
	}

	_, err = data.Write([]byte(jsonString))

	if err != nil {
		log.Println("Error in json update: ", err.Error())

		return
	}
}
