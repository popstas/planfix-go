package planfix

import (
	"fmt"
)

func (a *API) GetAnaliticByName(searchName string) (XmlResponseAnalitic, error) {
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
	return XmlResponseAnalitic{}, fmt.Errorf("Analitic %s not found", searchName)
}

func (a *API) GetHandbookRecordByName(handbookID int, searchName string) (XmlResponseAnaliticHandbookRecord, error) {
	var handbook XmlResponseAnaliticGetHandbook
	handbook, err := a.AnaliticGetHandbook(handbookID)
	if err != nil {
		return XmlResponseAnaliticHandbookRecord{}, err
	}
	for _, record := range handbook.Records {
		if record.ValuesMap["Название"] == searchName {
			return record, nil
		}
	}
	return XmlResponseAnaliticHandbookRecord{}, fmt.Errorf("Record %s not found", searchName)
}
