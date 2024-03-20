package helpers

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/FerryDwiZ12/FGA-Golang/tree/master/assigments/models"
)

func updateElement(min, max int) (result *models.Element) {
	result = &models.Element{
		Water: rune(rand.Intn(max-min) + min),
		Wind:  rune(rand.Intn(max-min) + min),
	}
	return
}

func checkCondition(element *models.Element) (status *models.Status) {
	var waterStatus string
	var windStatus string

	// Water Condition
	switch {
	case element.Water <= 5:
		waterStatus = "Aman"
	case element.Water >= 6 && element.Water <= 8:
		waterStatus = "Siaga"
	case element.Water > 8:
		waterStatus = "Bahaya"
	default:
		waterStatus = "Water Value Not Defined"
	}

	// Wind Condition
	switch {
	case element.Wind <= 6:
		windStatus = "Aman"
	case element.Wind >= 7 && element.Wind <= 15:
		windStatus = "Siaga"
	case element.Wind > 15:
		windStatus = "Bahaya"
	default:
		windStatus = "Wind Value Not Defined"
	}

	status = &models.Status{
		WaterStatus: waterStatus,
		WindStatus:  windStatus,
	}
	return
}

func writeJSON(element *models.Element) {
    jsonData, err := json.Marshal(element)
    if err != nil {
        log.Println("Error marshalling JSON:", err)
        return
    }

    file, err := os.Create("data/element.json")
    if err != nil {
        log.Println("Error creating file:", err)
        return
    }

    defer file.Close()

    _, err = file.Write(jsonData)
    if err != nil {
        log.Println("Error writing JSON to file:", err)
        return
    }
}

func RunCronJob() {
    s := gocron.NewScheduler(time.UTC)

    s.Every(5).Seconds().Do(func() {
        Element := updateElement(1, 100)
        // Condition := checkCondition(Element)

        writeJSON(Element)
        // log.Printf("Element Water: %v m, Status %v\n", Element.Water, Condition.WaterStatus)
        // log.Printf("Element Wind: %v m/s, Status %v\n", Element.Wind, Condition.WindStatus)
    })

    s.StartBlocking()
}
