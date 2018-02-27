package planfix

import (
	"errors"
	"fmt"
)

func (a *Api) GetAnaliticByName(searchName string) (XmlResponseAnalitic, error) {
	var analiticList XmlResponseAnaliticGetList
	analiticList, err := a.AnaliticGetList(0)
	if err != nil {
		return XmlResponseAnalitic{}, err
	}
	for _, analitic := range analiticList.Analitics.Analitics {
		if analitic.Name == searchName {
			return analitic, nil
		}
	}
	return XmlResponseAnalitic{}, errors.New(fmt.Sprintf("Analitic %s not found", searchName))
}

func (a *Api) GetHandbookRecordByName(handbookId int, searchName string) (XmlResponseAnaliticHandbookRecord, error) {
	var handbook XmlResponseAnaliticGetHandbook
	handbook, err := a.AnaliticGetHandbook(handbookId)
	if err != nil {
		return XmlResponseAnaliticHandbookRecord{}, err
	}
	for _, record := range handbook.Records {
		if record.ValuesMap["Название"] == searchName {
			return record, nil
		}
	}
	return XmlResponseAnaliticHandbookRecord{}, errors.New(fmt.Sprintf("Record %s not found", searchName))
}
