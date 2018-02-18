package planfix

import (
	"log"
	"errors"
	"fmt"
)

// auth.login
func (a Api) AuthLogin(user, password string) (string, error) {
	requestStruct := XmlRequestAuth{
		Method:   "auth.login",
		Account:  a.Account,
		Login:    a.User,
		Password: a.Password,
	}
	responseStruct := new(XmlResponseAuth)

	err := a.apiRequest(requestStruct, responseStruct)
	if err != nil {
		log.Printf("[ERROR] %v", err)
		return "", err
	}

	if responseStruct.Status == "error" {
		return "", errors.New(fmt.Sprintf(
			"Planfix request to %s failed: %s",
			requestStruct.Method,
			a.getErrorByCode(responseStruct.Code)))
	}

	return responseStruct.Sid, nil
}

// action.get
func (a *Api) ActionGet(actionId int) (XmlResponseActionGet, error) {
	a.ensureAuthenticated()
	requestStruct := XmlRequestActionGet{
		Method:   "action.get",
		Account:  a.Account,
		Sid:      a.Sid,
		ActionId: actionId,
	}
	responseStruct := new(XmlResponseActionGet)

	err := a.apiRequest(requestStruct, responseStruct)
	if err != nil {
		return XmlResponseActionGet{}, err
	}

	if responseStruct.Status == "error" {
		return XmlResponseActionGet{}, errors.New(fmt.Sprintf(
			"Planfix request to %s failed: %s",
			requestStruct.Method,
			a.getErrorByCode(responseStruct.Code)))
	}

	return *responseStruct, nil
}

// action.getList
func (a *Api) ActionGetList(requestStruct XmlRequestActionGetList) (XmlResponseActionGetList, error) {
	a.ensureAuthenticated()

	// defaults
	requestStruct.Method = "action.getList"
	if requestStruct.Account == "" {
		requestStruct.Account = a.Account
	}
	if requestStruct.Sid == "" {
		requestStruct.Sid = a.Sid
	}
	if requestStruct.PageCurrent == 0 {
		requestStruct.PageCurrent = 1
	}
	if requestStruct.PageSize == 0 {
		requestStruct.PageSize = 100
	}

	responseStruct := new(XmlResponseActionGetList)

	err := a.apiRequest(requestStruct, responseStruct)
	if err != nil {
		log.Printf("[ERROR] %v", err)
		return XmlResponseActionGetList{}, err
	}

	if responseStruct.Status == "error" {
		return XmlResponseActionGetList{}, errors.New(fmt.Sprintf(
			"Planfix request to %s failed: %s",
			requestStruct.Method,
			a.getErrorByCode(responseStruct.Code)))
	}

	return *responseStruct, nil
}
