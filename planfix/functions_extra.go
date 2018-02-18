package planfix

import (
	"errors"
	"fmt"
)

func (a *Api) GetAnaliticByName(searchName string) (XmlResponseAnalitic, error) {
	var analiticList XmlResponseAnaliticGetList
	analiticList, err := a.AnaliticGetList(0)
	if err != nil {
		fmt.Println(err)
	}
	for _, analitic := range analiticList.Analitics.Analitics {
		if analitic.Name == searchName {
			return analitic, nil
		}
	}
	return XmlResponseAnalitic{}, errors.New(fmt.Sprintf("Analitic %s not found", searchName))
}
