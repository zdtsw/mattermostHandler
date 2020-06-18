package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"github.com/tatsushid/go-prettytable"
	"github.com/urfave/cli/v2"
)

type eventRecord struct {
	EC EventCheck  `json:"check"`
	EE EventEntity `json:"entity"`
}

// EventCheck for sensugo json event .check.*
type EventCheck struct {
	Output   string `json:"output"`
	Metadata MD     `json:"metadata"`
	State    string `json:"state"`
	Status   int    `json:"status"`
}

// EventEntity for sensugo json event .entity.*
type EventEntity struct {
	Metadata MD `json:"metadata"`
}

// MD for sensugo json event .check.metadata.name and .entity.metadata.name
type MD struct {
	Name string `json:"name"`
}

const urlMM = "https://mattermost.mycompany.com/hooks/balabalahookwebalabala"

//below emoj need to be able in mattermost.mycompany.com or you can replace with any you want
var resolveText = "### :peace_symbol: The following SensuGo check has been resolved.\n"
var warningText = "### :warning: Warning from SensuGo @channel please review the following alert.\n"
var criticalText = "### :fire: Critial from SensuGo @channel please fix the following alert.\n"
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
		Version: "1.0.0",
		Usage:   "Send messaget to Mycompany's Mattermost channel",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "webhook",
				Aliases:     []string{"w"},
				Value:       urlMM,
				Usage:       "URL to Mattermost Webhook",
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

// prettytalbe not work in this case, because return type is []byte, but good to call Print() for debug
func formatTable(h, c, o string) []byte {
	tbl, err := prettytable.NewTable([]prettytable.Column{
		{Header: "Host", MinWidth: 10},
		{Header: "Check", MinWidth: 8},
		{Header: "Output"},
	}...)
	if err != nil {
		panic(err)
	}
	tbl.Separator = " | "
	tbl.AddRow(h, c, o)
	return tbl.Bytes()
}

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

	resp, err := http.Post(webhooks, "application/json; charset=utf-8", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}
	return resp.Status
}
