package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func postAnnounceHandler(detach bool) {

	if detach {
		detachPost()
	} else {
		interactivePost()
	}
}
func detachPost() {
	var announcement string

	announcement = "@channel\n## ANNOUNCEMENT: " + os.Getenv("topic") + "\n"
	announcement += "__" + os.Getenv("time") + "__ \n"
	announcement += os.Getenv("announcement") + "\nXaaS Team Delivery"
	fmt.Println(announcement)
	postToMMHandler(announcement, "Wen")

}

func interactivePost() {

	scanner := bufio.NewScanner(os.Stdin)
	var announcement, text string

	fmt.Print("Enter Announcement Topic: (e.g Jenkins Upgrade, LDAP change, etc) ")
	scanner.Scan()
	announcement = "@channel\n## ANNOUNCEMENT: " + scanner.Text() + "\n"

	fmt.Print("Enter Announcement detail(stop by 'XaaS Team Delivery'):\n")
	fmt.Print("\twhen (e.g yyyy-mm-dd@hh:mm or near future etc): ")
	scanner.Scan()
	announcement += "__" + scanner.Text() + "__ "

	for !strings.EqualFold(text, "WEN") {
		fmt.Print("\tDetail line: ")
		scanner.Scan()
		text = scanner.Text()
		announcement += scanner.Text() + "\n"
	}

	postToMMHandler(announcement, "Wen")
}
