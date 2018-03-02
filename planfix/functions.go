package planfix

import (
	"errors"
)

// AuthLogin = auth.login
func (a API) AuthLogin(user, password string) (string, error) {
	requestStruct := XMLRequestAuthLogin{
		Method:   "auth.login",
		Account:  a.Account,
		Login:    a.User,
		Password: a.Password,
	}
	responseStruct := new(XMLResponseAuth)

	err := a.apiRequest(&requestStruct, responseStruct)
	return responseStruct.Sid, err
}

// ActionGet = action.get
func (a *API) ActionGet(actionID int) (XMLResponseActionGet, error) {
	requestStruct := XMLRequestActionGet{
		ActionID: actionID,
	}
	requestStruct.Method = "action.get"
	responseStruct := new(XMLResponseActionGet)

	err := a.apiRequest(&requestStruct, responseStruct)
	return *responseStruct, err
}

// ActionGetList = action.getList
func (a *API) ActionGetList(requestStruct XMLRequestActionGetList) (XMLResponseActionGetList, error) {
	requestStruct.Method = "action.getList"

	// defaults
	if requestStruct.PageCurrent == 0 {
		requestStruct.PageCurrent = 1
	}
	if requestStruct.PageSize == 0 {
		requestStruct.PageSize = 100
	}

	responseStruct := new(XMLResponseActionGetList)

	err := a.apiRequest(&requestStruct, responseStruct)
	return *responseStruct, err
}

// AnaliticGetList = analitic.getList
func (a *API) AnaliticGetList(groupID int) (XMLResponseAnaliticGetList, error) {
	requestStruct := XMLRequestAnaliticGetList{
		AnaliticGroupID: groupID,
	}
	requestStruct.Method = "analitic.getList"
	responseStruct := new(XMLResponseAnaliticGetList)

	err := a.apiRequest(&requestStruct, responseStruct)
	return *responseStruct, err
}

// AnaliticGetHandbook = analitic.getHandbook
func (a *API) AnaliticGetHandbook(handbookID int) (XMLResponseAnaliticGetHandbook, error) {
	requestStruct := XMLRequestAnaliticGetHandbook{
		HandbookID: handbookID,
	}
	requestStruct.Method = "analitic.getHandbook"
	responseStruct := new(XMLResponseAnaliticGetHandbook)

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

// AnaliticGetOptions = analitic.get
func (a *API) AnaliticGetOptions(analiticID int) (XMLResponseAnaliticGetOptions, error) {
	requestStruct := XMLRequestAnaliticGetOptions{
		AnaliticID: analiticID,
	}
	requestStruct.Method = "analitic.getOptions"
	responseStruct := new(XMLResponseAnaliticGetOptions)

	err := a.apiRequest(&requestStruct, responseStruct)
	return *responseStruct, err
}

// ActionAdd = action.add
func (a *API) ActionAdd(requestStruct XMLRequestActionAdd) (XMLResponseActionAdd, error) {
	requestStruct.Method = "action.add"

	// only task or contact allowed
	if (requestStruct.TaskID > 0 || requestStruct.TaskGeneral > 0) && requestStruct.ContactGeneral > 0 {
		return XMLResponseActionAdd{}, errors.New("Both task and contact defined")
	}

	responseStruct := new(XMLResponseActionAdd)

	err := a.apiRequest(&requestStruct, responseStruct)
	return *responseStruct, err
}

// TaskGet = task.get
func (a *API) TaskGet(taskID, taskGeneral int) (XMLResponseTaskGet, error) {
	requestStruct := XMLRequestTaskGet{
		TaskID:      taskID,
		TaskGeneral: taskGeneral,
	}
	requestStruct.Method = "task.get"
	responseStruct := new(XMLResponseTaskGet)

	err := a.apiRequest(&requestStruct, responseStruct)
	return *responseStruct, err
}

// UserGet = user.get
func (a *API) UserGet(userID int) (XMLResponseUserGet, error) {
	requestStruct := XMLRequestUserGet{
		UserID: userID,
	}
	requestStruct.Method = "user.get"
	responseStruct := new(XMLResponseUserGet)

	err := a.apiRequest(&requestStruct, responseStruct)
	return *responseStruct, err
}
