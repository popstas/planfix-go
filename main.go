package main

import (
	"github.com/popstas/planfix-go/planfix"
	"fmt"
	"os"
)

func main(){
	planfixApi := planfix.New(
		"https://api.planfix.ru/xml/",
		os.Getenv("PLANFIX_API_KEY"),
		os.Getenv("PLANFIX_ACCOUNT"),
		os.Getenv("PLANFIX_USER_NAME"),
		os.Getenv("PLANFIX_USER_PASSWORD"),
	)
	planfixApi.UserAgent = "planfix-toggl"

	var actionList planfix.XmlResponseActionGetList
	actionList, err := planfixApi.ActionGetList(planfix.XmlRequestActionGetList{
		TaskGeneral: 123,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("action.getList result: %d", actionList.Actions.Actions)
}
