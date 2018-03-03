package planfix_test

import (
	"github.com/popstas/planfix-go/planfix"
	"testing"
)

func TestAPI_GetAnaliticByName(t *testing.T) {
	api := newAPI([]string{
		fixtureFromFile("analitic.getList.xml"),
		fixtureFromFile("analitic.getList.xml"),
		fixtureFromFile("error.xml"),
	})
	var analitic planfix.XMLResponseAnalitic

	// existent
	analitic, err := api.GetAnaliticByName("Выработка")
	expectSuccess(t, err, "TestAPI_GetAnaliticByName Выработка")
	assert(t, analitic.Name, "Выработка")

	// non existent
	analitic, err = api.GetAnaliticByName("ldkfgjld")
	expectError(t, err, "TestAPI_GetAnaliticByName non existent")

	// error
	analitic, err = api.GetAnaliticByName("ldkfgjld")
	expectError(t, err, "TestAPI_GetAnaliticByName error")
}

func TestAPI_GetHandbookRecordByName(t *testing.T) {
	api := newAPI([]string{
		fixtureFromFile("analitic.getHandbook.xml"),
		fixtureFromFile("analitic.getHandbook.xml"),
		fixtureFromFile("error.xml"),
	})
	var record planfix.XMLResponseAnaliticHandbookRecord

	// existent
	record, err := api.GetHandbookRecordByName(123, "Поминутная работа программиста")
	expectSuccess(t, err, "TestAPI_GetHandbookRecordByName Поминутная работа программиста")
	assert(t, record.ValuesMap["Название"], "Поминутная работа программиста")

	// non existent
	record, err = api.GetHandbookRecordByName(123, "ldkfgjld")
	expectError(t, err, "TestAPI_GetHandbookRecordByName non existent")

	// error
	record, err = api.GetHandbookRecordByName(123, "ldkfgjld")
	expectError(t, err, "TestAPI_GetHandbookRecordByName error")
}

func TestAPI_GetActiveUserByLogin(t *testing.T) {
	api := newAPI([]string{
		fixtureFromFile("user.getList.xml"),
		fixtureFromFile("user.getList.xml"),
		fixtureFromFile("error.xml"),
	})
	var user planfix.XMLResponseUser

	// existent
	user, err := api.GetActiveUserByLogin("popstas")
	expectSuccess(t, err, "TestAPI_GetActiveUserByLogin popstas")
	assert(t, user.Email, "popstas@company.ru")

	// non existent
	user, err = api.GetActiveUserByLogin("ldkfgjld")
	expectError(t, err, "TestAPI_GetActiveUserByLogin non existent")

	// error
	user, err = api.GetActiveUserByLogin("ldkfgjld")
	expectError(t, err, "TestAPI_GetActiveUserByLogin error")
}
