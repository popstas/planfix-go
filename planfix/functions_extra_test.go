package planfix_test

import (
	"github.com/popstas/planfix-go/planfix"
	"testing"
)

func TestApi_GetAnaliticByName(t *testing.T) {
	api := newApi("../tests/fixtures/analitic.getList.xml")
	var analitic planfix.XmlResponseAnalitic
	analitic, err := api.GetAnaliticByName("Выработка")
	if err != nil {
		t.Error(err)
	}
	if analitic.Name != "Выработка" {
		t.Error("Expected Выработка, got ", analitic.Name)
	}
}
