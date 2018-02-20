package planfix

import (
	"errors"
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
	return responseStruct.Sid, err
}

// action.get
func (a *Api) ActionGet(actionId int) (XmlResponseActionGet, error) {
	requestStruct := XmlRequestActionGet{
		ActionId: actionId,
	}
	requestStruct.Method = "action.get"
	responseStruct := new(XmlResponseActionGet)

	err := a.apiRequest(&requestStruct, responseStruct)
	return *responseStruct, err
}

// action.getList
func (a *Api) ActionGetList(requestStruct XmlRequestActionGetList) (XmlResponseActionGetList, error) {
	requestStruct.Method = "action.getList"

	// defaults
	if requestStruct.PageCurrent == 0 {
		requestStruct.PageCurrent = 1
	}
	if requestStruct.PageSize == 0 {
		requestStruct.PageSize = 100
	}

	responseStruct := new(XmlResponseActionGetList)

	err := a.apiRequest(&requestStruct, responseStruct)
	return *responseStruct, err
}

// analitic.getList
func (a *Api) AnaliticGetList(groupId int) (XmlResponseAnaliticGetList, error) {
	requestStruct := XmlRequestAnaliticGetList{
		AnaliticGroupId: groupId,
	}
	requestStruct.Method = "analitic.getList"
	responseStruct := new(XmlResponseAnaliticGetList)

	err := a.apiRequest(&requestStruct, responseStruct)
	return *responseStruct, err
}

// analitic.get
func (a *Api) AnaliticGetOptions(analiticId int) (XmlResponseAnaliticGetOptions, error) {
	requestStruct := XmlRequestAnaliticGetOptions{
		AnaliticId: analiticId,
	}
	requestStruct.Method = "analitic.getOptions"
	responseStruct := new(XmlResponseAnaliticGetOptions)

	err := a.apiRequest(&requestStruct, responseStruct)
	return *responseStruct, err
}

// action.add
func (a *Api) ActionAdd(requestStruct XmlRequestActionAdd) (XmlResponseActionAdd, error) {
	requestStruct.Method = "action.add"

	// only task or contact allowed
	if (requestStruct.TaskId > 0 || requestStruct.TaskGeneral > 0) && requestStruct.ContactGeneral > 0 {
		return XmlResponseActionAdd{}, errors.New("Both task and contact defined")
	}

	responseStruct := new(XmlResponseActionAdd)

	err := a.apiRequest(&requestStruct, responseStruct)
	return *responseStruct, err
}

// task.get
func (a *Api) TaskGet(taskId, taskGeneral int) (XmlResponseTaskGet, error) {
	requestStruct := XmlRequestTaskGet{
		TaskId:      taskId,
		TaskGeneral: taskGeneral,
	}
	requestStruct.Method = "task.get"
	responseStruct := new(XmlResponseTaskGet)

	err := a.apiRequest(&requestStruct, responseStruct)
	return *responseStruct, err
}

// user.get
func (a *Api) UserGet(userId int) (XmlResponseUserGet, error) {
	requestStruct := XmlRequestUserGet{
		UserId: userId,
	}
	requestStruct.Method = "user.get"
	responseStruct := new(XmlResponseUserGet)

	err := a.apiRequest(&requestStruct, responseStruct)
	return *responseStruct, err
}
