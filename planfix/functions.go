package planfix

import (
	"errors"
	"fmt"
	"log"
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

// analitic.getList
func (a *Api) AnaliticGetList(groupId int) (XmlResponseAnaliticGetList, error) {
	a.ensureAuthenticated()
	requestStruct := XmlRequestAnaliticGetList{
		Method:          "analitic.getList",
		Account:         a.Account,
		Sid:             a.Sid,
		AnaliticGroupId: groupId,
	}
	responseStruct := new(XmlResponseAnaliticGetList)

	err := a.apiRequest(requestStruct, responseStruct)
	if err != nil {
		log.Printf("[ERROR] %v", err)
		return XmlResponseAnaliticGetList{}, err
	}

	if responseStruct.Status == "error" {
		return XmlResponseAnaliticGetList{}, errors.New(fmt.Sprintf(
			"Planfix request to %s failed: %s",
			requestStruct.Method,
			a.getErrorByCode(responseStruct.Code)))
	}

	return *responseStruct, nil
}

// analitic.get
func (a *Api) AnaliticGetOptions(analiticId int) (XmlResponseAnaliticGetOptions, error) {
	a.ensureAuthenticated()
	requestStruct := XmlRequestAnaliticGetOptions{
		Method:     "analitic.getOptions",
		Account:    a.Account,
		Sid:        a.Sid,
		AnaliticId: analiticId,
	}
	responseStruct := new(XmlResponseAnaliticGetOptions)

	err := a.apiRequest(requestStruct, responseStruct)
	if err != nil {
		return XmlResponseAnaliticGetOptions{}, err
	}

	if responseStruct.Status == "error" {
		return XmlResponseAnaliticGetOptions{}, errors.New(fmt.Sprintf(
			"Planfix request to %s failed: %s, %s",
			requestStruct.Method,
			a.getErrorByCode(responseStruct.Code),
			responseStruct.Message))
	}

	return *responseStruct, nil
}

// action.add
func (a *Api) ActionAdd(requestStruct XmlRequestActionAdd) (XmlResponseActionAdd, error) {
	a.ensureAuthenticated()

	// only task or contact allowed
	if (requestStruct.TaskId > 0 || requestStruct.TaskGeneral > 0) && requestStruct.ContactGeneral > 0 {
		return XmlResponseActionAdd{}, errors.New("Both task and contact defined")
	}

	// defaults
	requestStruct.Method = "action.add"
	if requestStruct.Account == "" {
		requestStruct.Account = a.Account
	}
	if requestStruct.Sid == "" {
		requestStruct.Sid = a.Sid
	}

	responseStruct := new(XmlResponseActionAdd)

	err := a.apiRequest(requestStruct, responseStruct)
	if err != nil {
		log.Printf("[ERROR] %v", err)
		return XmlResponseActionAdd{}, err
	}

	if responseStruct.Status == "error" {
		return XmlResponseActionAdd{}, errors.New(fmt.Sprintf(
			"Planfix request to %s failed: %s, %s",
			requestStruct.Method,
			a.getErrorByCode(responseStruct.Code),
			responseStruct.Message))
	}

	return *responseStruct, nil
}

// task.get
func (a *Api) TaskGet(taskId, taskGeneral int) (XmlResponseTaskGet, error) {
	a.ensureAuthenticated()
	requestStruct := XmlRequestTaskGet{
		Method:      "task.get",
		Account:     a.Account,
		Sid:         a.Sid,
		TaskId:      taskId,
		TaskGeneral: taskGeneral,
	}
	responseStruct := new(XmlResponseTaskGet)

	err := a.apiRequest(requestStruct, responseStruct)
	if err != nil {
		return XmlResponseTaskGet{}, err
	}

	if responseStruct.Status == "error" {
		return XmlResponseTaskGet{}, errors.New(fmt.Sprintf(
			"Planfix request to %s failed: %s",
			requestStruct.Method,
			a.getErrorByCode(responseStruct.Code)))
	}

	return *responseStruct, nil
}
