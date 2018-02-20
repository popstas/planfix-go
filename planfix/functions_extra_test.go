package planfix_test

import (
	"github.com/popstas/planfix-go/planfix"
	"testing"
)

func TestApi_GetAnaliticByName(t *testing.T) {
	api := newApi([]string{
		fixtureFromFile("analitic.getList.xml"),
		fixtureFromFile("analitic.getList.xml"),
		fixtureFromFile("error.xml"),
	})
	var analitic planfix.XmlResponseAnalitic

	// existent
	analitic, err := api.GetAnaliticByName("Выработка")
	expectSuccess(t, err, "TestApi_GetAnaliticByName Выработка")
	assert(t, analitic.Name, "Выработка")

	// non existent
	analitic, err = api.GetAnaliticByName("ldkfgjld")
	expectError(t, err, "TestApi_GetAnaliticByName non existent")

	// error
	analitic, err = api.GetAnaliticByName("ldkfgjld")
	expectError(t, err, "TestApi_GetAnaliticByName error")
}
