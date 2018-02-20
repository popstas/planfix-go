package planfix

import (
	"errors"
	"log"
)

// auth.login
func (a Api) AuthLogin(user, password string) (string, error) {
	requestStruct := XmlRequestAuthLogin{
		Method:   "auth.login",
		Account:  a.Account,
		Login:    a.User,
		Password: a.Password,
	}
	responseStruct := new(XmlResponseAuth)

	err := a.apiRequest(&requestStruct, responseStruct)
	if err != nil {
		log.Printf("[ERROR] %v", err)
		return "", err
	}

	return responseStruct.Sid, nil
}

// action.get
func (a *Api) ActionGet(actionId int) (XmlResponseActionGet, error) {
	a.ensureAuthenticated()
	requestStruct := XmlRequestActionGet{
		ActionId: actionId,
	}
	requestStruct.Method = "action.get"
	responseStruct := new(XmlResponseActionGet)

	err := a.apiRequest(&requestStruct, responseStruct)
	if err != nil {
		return XmlResponseActionGet{}, err
	}

	return *responseStruct, nil
}

// action.getList
func (a *Api) ActionGetList(requestStruct XmlRequestActionGetList) (XmlResponseActionGetList, error) {
	a.ensureAuthenticated()

	// defaults
	requestStruct.Method = "action.getList"
	if requestStruct.PageCurrent == 0 {
		requestStruct.PageCurrent = 1
	}
	if requestStruct.PageSize == 0 {
		requestStruct.PageSize = 100
	}

	responseStruct := new(XmlResponseActionGetList)

	err := a.apiRequest(&requestStruct, responseStruct)
	if err != nil {
		log.Printf("[ERROR] %v", err)
		return XmlResponseActionGetList{}, err
	}

	return *responseStruct, nil
}

// analitic.getList
func (a *Api) AnaliticGetList(groupId int) (XmlResponseAnaliticGetList, error) {
	a.ensureAuthenticated()
	requestStruct := XmlRequestAnaliticGetList{
		AnaliticGroupId: groupId,
	}
	requestStruct.Method = "analitic.getList"
	responseStruct := new(XmlResponseAnaliticGetList)

	err := a.apiRequest(&requestStruct, responseStruct)
	if err != nil {
		log.Printf("[ERROR] %v", err)
		return XmlResponseAnaliticGetList{}, err
	}

	return *responseStruct, nil
}

// analitic.get
func (a *Api) AnaliticGetOptions(analiticId int) (XmlResponseAnaliticGetOptions, error) {
	a.ensureAuthenticated()
	requestStruct := XmlRequestAnaliticGetOptions{
		AnaliticId: analiticId,
	}
	requestStruct.Method = "analitic.getOptions"
	responseStruct := new(XmlResponseAnaliticGetOptions)

	err := a.apiRequest(&requestStruct, responseStruct)
	if err != nil {
		return XmlResponseAnaliticGetOptions{}, err
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

	responseStruct := new(XmlResponseActionAdd)

	err := a.apiRequest(&requestStruct, responseStruct)
	if err != nil {
		log.Printf("[ERROR] %v", err)
		return XmlResponseActionAdd{}, err
	}

	return *responseStruct, nil
}

// task.get
func (a *Api) TaskGet(taskId, taskGeneral int) (XmlResponseTaskGet, error) {
	a.ensureAuthenticated()
	requestStruct := XmlRequestTaskGet{
		TaskId:      taskId,
		TaskGeneral: taskGeneral,
	}
	requestStruct.Method = "task.get"
	responseStruct := new(XmlResponseTaskGet)

	err := a.apiRequest(&requestStruct, responseStruct)
	if err != nil {
		return XmlResponseTaskGet{}, err
	}

	return *responseStruct, nil
}

// user.get
func (a *Api) UserGet(userId int) (XmlResponseUserGet, error) {
	a.ensureAuthenticated()
	requestStruct := XmlRequestUserGet{
		UserId: userId,
	}
	requestStruct.Method = "user.get"
	responseStruct := new(XmlResponseUserGet)

	err := a.apiRequest(&requestStruct, responseStruct)
	if err != nil {
		return XmlResponseUserGet{}, err
	}

	return *responseStruct, nil
}
