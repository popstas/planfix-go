package main

import (
	"fmt"
	"github.com/popstas/planfix-go/planfix"
	"os"
)

func main() {
	planfixAPI := planfix.New(
		"https://api.planfix.ru/xml/",
		os.Getenv("PLANFIX_API_KEY"),
		os.Getenv("PLANFIX_ACCOUNT"),
		os.Getenv("PLANFIX_USER_NAME"),
		os.Getenv("PLANFIX_USER_PASSWORD"),
	)
	planfixAPI.UserAgent = "planfix-toggl"

	var actionList planfix.XMLResponseActionGetList
	actionList, err := planfixAPI.ActionGetList(planfix.XMLRequestActionGetList{
		TaskGeneral: 123,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("action.getList result: %d", actionList.Actions.Actions)
}
