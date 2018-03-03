package planfix

import (
	"fmt"
)

// GetAnaliticByName возвращает аналитику по названию
func (a *API) GetAnaliticByName(searchName string) (XMLResponseAnalitic, error) {
	var analiticList XMLResponseAnaliticGetList
	analiticList, err := a.AnaliticGetList(0)
	if err != nil {
		return XMLResponseAnalitic{}, err
	}
	for _, analitic := range analiticList.Analitics.Analitics {
		if analitic.Name == searchName {
			return analitic, nil
		}
	}
	return XMLResponseAnalitic{}, fmt.Errorf("Analitic %s not found", searchName)
}

// GetHandbookRecordByName возвращает запись справочника ID справочника и названию
func (a *API) GetHandbookRecordByName(handbookID int, searchName string) (XMLResponseAnaliticHandbookRecord, error) {
	var handbook XMLResponseAnaliticGetHandbook
	handbook, err := a.AnaliticGetHandbook(handbookID)
	if err != nil {
		return XMLResponseAnaliticHandbookRecord{}, err
	}
	for _, record := range handbook.Records {
		if record.ValuesMap["Название"] == searchName {
			return record, nil
		}
	}
	return XMLResponseAnaliticHandbookRecord{}, fmt.Errorf("Record %s not found", searchName)
}

// GetHandbookRecordByName возвращает юзера по логину
func (a *API) GetActiveUserByLogin(login string) (XMLResponseUser, error) {
	users, err := a.UserGetList(XMLRequestUserGetList{Status: "ACTIVE"})
	if err != nil {
		return XMLResponseUser{}, err
	}
	for _, user := range users.Users.Users {
		if user.Login == login {
			return user, nil
		}
	}
	return XMLResponseUser{}, fmt.Errorf("User %s not found", login)
}
