package main

import (
	"bytes"
	"encoding/json"
	"github.com/tatsushid/go-prettytable"
	"github.com/urfave/cli/v2"
	"log"
	"net/http"
	"os"
	"fmt"
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

const (
	urlMM        = "https://mattermost.mycompany.com/hooks/gulugulugulugulu"
	urlDC        = "https://mattermost.mycompany.com/hooks/balalalalalalal"
	resolveText  = "### :peace_symbol: The following SensuGo check has been resolved.\n"
	warningText  = "### :warning: Warning from SensuGo @channel please review the following alert.\n"
	criticalText = "### :fire: Critial from SensuGo @channel please fix the following alert.\n"
)

var (
	webhooks string
	detach   bool
	version	string
)

func main() {

	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("version=%s\n", version)
	  }

	app := &cli.App{
		Name: "Handler Mattermost (HM)",
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "Wen Zhou",
				Email: "ericchou19831101@msn.com",
			},
		},
		Version: " ",
		Usage:   "Send messagers to Mattermost channel",

		Commands: []*cli.Command{
			{
				Name:      "alert",
				Aliases:   []string{"al"},
				Usage:     "Post SensuGo alerts",
				UsageText: " Post Sensu Go Allerts to Monitoring Channel\n\tthis should be done by Sensu Handler automatically",
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
			},
			{
				Name:      "announce",
				Aliases:   []string{"an"},
				Usage:     "Post announcement",
				UsageText: "Interactive program to post announcement to channel, finish post by using 'WEN' case insensitive",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "webhook",
						Aliases:     []string{"w"},
						Value:       urlDC,
						Usage:       "URL to Mattermost Webhook",
						Destination: &webhooks,
					},
					&cli.BoolFlag{
						Name:        "detach",
						Aliases:     []string{"d"},
						Value:       false,
						Usage:       "Disable interactive, get input from env varialb: topic, time, announcement",
						Destination: &detach,
					},
				},
				Action: func(c *cli.Context) error {
					postAnnounceHandler(detach)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
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

func postToMMHandler(text, user string) (status string) {
	message := map[string]interface{}{
		"username": user,
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
