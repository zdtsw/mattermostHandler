package main

import (
	"encoding/json"
	"log"
	"os"
)

func parseEventHandler() {
	var check eventRecord

	decoder := json.NewDecoder(os.Stdin)
	err := decoder.Decode(&check)

	payload := "| Host Name | Check Name | Output \n| :--- | :--- | :--- \n| " + check.EE.Metadata.Name + " | " + check.EC.Metadata.Name + " | " + check.EC.State + ": " + check.EC.Output + " | \n"

	text := "Unknown event : " + check.EE.Metadata.Name + " on host: " + check.EE.Metadata.Name

	if err != nil {
		log.Fatalln(err)
	} else {
		switch status := check.EC.Status; status {
		case 1:
			text = warningText + payload
		case 2:
			text = criticalText + payload
		case 0:
			text = resolveText + payload
		default:
		}
	}
	postToMMHandler(text, "sensu")
}
