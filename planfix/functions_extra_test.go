package planfix_test

import (
	"github.com/popstas/planfix-go/planfix"
	"testing"
)

func TestApi_GetAnaliticByName(t *testing.T) {
	api := newApi("../tests/fixtures/analitic.getList.xml")
	var analitic planfix.XmlResponseAnalitic

	// existent
	analitic, err := api.GetAnaliticByName("Выработка")
	expectSuccess(t, err, "TestApi_GetAnaliticByName Выработка")
	assert(t, analitic.Name, "Выработка")

	// non existent
	analitic, err = api.GetAnaliticByName("ldkfgjld")
	expectError(t, err, "TestApi_GetAnaliticByName non existent")

	// error
	api = newApi("../tests/fixtures/error.xml")
	analitic, err = api.GetAnaliticByName("ldkfgjld")
	expectError(t, err, "TestApi_GetAnaliticByName error")
}
