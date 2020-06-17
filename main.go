package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/urfave/cli/v2"
)

type everntRecord struct {
	Name   string
	Output string
	Status int
}

var payload = "\n| Host               | Check Name      | Output                    |\n| :----------------- |:--------------- | :-------------------------|\n| #{check_fqdn}      | #{check_name}   | #{check_output}           |\n"

const urlMM = "https://mattermost.mycompany.com/hooks/ayn7yhrtqtnyfguqz5cbbiuptr"

var resolveText = "### :white_check_mark: The following sensu check has been resolved" + payload
var warningText = "### :warning: Warning from sensu @channel please review the following alert." + payload
var criticalText = "### :critical: Critial from sensu @channel please fix the following alert." + payload

var webhooks string

func main() {
	app := &cli.App{
		Name: "Handler Mattermost (H&M)",
		Authors: []*cli.Author{
			&cli.Author{
			  Name:  "Wen Zhou",
			  Email: "ericchou19831101@msn.com",
			},
		},
		Usage: "Send messaget to Mattermost channel",
		Flags: []cli.Flag {
			&cli.StringFlag{
			  Name: "webhook",
			  Aliases:  []string{"w"},
			  Value:  urlMM,
			  Usage:  "URL to Mattermost Webhook",
			  Destination: &webhooks,
			},
		  },
		Action: func(c *cli.Context) error {
			parseEventHandler()
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func parseEventHandler() {

	var check everntRecord
	//decoder := json.NewDecoder(os.Stdin)
	//err := decoder.Decode(&check)

	//test data
	jsonData := []byte(`
	{
		"name": "srvumgmt88",
		"status": 1,
		"output": "this is warning "
	}`)
	err := json.Unmarshal(jsonData, &check)
	fmt.Println(check.Name)
	fmt.Println(check.Output)
	fmt.Println(check.Status)
	//test done
	text := "unknown:" + check.Name

	if err != nil {
		log.Fatalln(err)
	} else {
		switch status := check.Status; status {
		case 1:
			fmt.Printf("set payload warning: %s", check.Output)
			text = warningText
		case 2:
			fmt.Printf("set payload critial: %s", check.Output)
			text = criticalText
		case 0:
			fmt.Printf("set payload  ok: %s", check.Output)
			text = resolveText
		default:
			fmt.Print("unknown"
		}
	}
	postToMMHandler(text)
}

func postToMMHandler(text string) (status string) {

	message := map[string]interface{}{
		"username": "sensu",
		"text":     text,
	}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post(urlMM, "application/json; charset=utf-8", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}
	return resp.Status
}
