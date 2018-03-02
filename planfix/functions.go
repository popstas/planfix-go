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
func (a *Api) ActionGet(actionID int) (XmlResponseActionGet, error) {
	requestStruct := XmlRequestActionGet{
		ActionID: actionID,
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
func (a *Api) AnaliticGetList(groupID int) (XmlResponseAnaliticGetList, error) {
	requestStruct := XmlRequestAnaliticGetList{
		AnaliticGroupID: groupID,
	}
	requestStruct.Method = "analitic.getList"
	responseStruct := new(XmlResponseAnaliticGetList)

	err := a.apiRequest(&requestStruct, responseStruct)
	return *responseStruct, err
}

// analitic.getHandbook
func (a *Api) AnaliticGetHandbook(handbookID int) (XmlResponseAnaliticGetHandbook, error) {
	requestStruct := XmlRequestAnaliticGetHandbook{
		HandbookID: handbookID,
	}
	requestStruct.Method = "analitic.getHandbook"
	responseStruct := new(XmlResponseAnaliticGetHandbook)

	err := a.apiRequest(&requestStruct, responseStruct)

	// map from values list
	for rid, record := range responseStruct.Records {
		record.ValuesMap = make(map[string]string)
		for _, value := range record.Values {
			record.ValuesMap[value.Name] = value.Value
		}
		responseStruct.Records[rid] = record
	}

	return *responseStruct, err
}

// analitic.get
func (a *Api) AnaliticGetOptions(analiticID int) (XmlResponseAnaliticGetOptions, error) {
	requestStruct := XmlRequestAnaliticGetOptions{
		AnaliticID: analiticID,
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
	if (requestStruct.TaskID > 0 || requestStruct.TaskGeneral > 0) && requestStruct.ContactGeneral > 0 {
		return XmlResponseActionAdd{}, errors.New("Both task and contact defined")
	}

	responseStruct := new(XmlResponseActionAdd)

	err := a.apiRequest(&requestStruct, responseStruct)
	return *responseStruct, err
}

// task.get
func (a *Api) TaskGet(taskID, taskGeneral int) (XmlResponseTaskGet, error) {
	requestStruct := XmlRequestTaskGet{
		TaskID:      taskID,
		TaskGeneral: taskGeneral,
	}
	requestStruct.Method = "task.get"
	responseStruct := new(XmlResponseTaskGet)

	err := a.apiRequest(&requestStruct, responseStruct)
	return *responseStruct, err
}

// user.get
func (a *Api) UserGet(userID int) (XmlResponseUserGet, error) {
	requestStruct := XmlRequestUserGet{
		UserID: userID,
	}
	requestStruct.Method = "user.get"
	responseStruct := new(XmlResponseUserGet)

	err := a.apiRequest(&requestStruct, responseStruct)
	return *responseStruct, err
}
