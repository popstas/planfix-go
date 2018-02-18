package planfix

import "encoding/xml"

/*type XmlRequest struct {
	XMLName xml.Name `xml:"request"`
	Method  string   `xml:"method,attr"`
}

func (r xmlRequest) getMethod() string {
	return r.Method
}
func (r *xmlRequest) setMethod(m string) {
	r.Method = m
}

type xmlResponse struct {
	XMLName xml.Name `xml:"response"`
	Status  string   `xml:"status,attr"`
}*/

type XmlResponseFile struct {
	Id   int    `xml:"id"`
	Name string `xml:"name"`
}

type XmlResponseUser struct {
	Id   int    `xml:"id"`
	Name string `xml:"name"`
}

type XmlResponseAnalitic struct {
	Id   int    `xml:"id"`
	Key  int    `xml:"key"`
	Name string `xml:"name"`
}

type XmlResponseAction struct {
	Id                           int                   `xml:"id"`
	Description                  string                `xml:"description"`
	OldStatus                    int                   `xml:"statusChange>oldStatus,omitempty"`
	NewStatus                    int                   `xml:"statusChange>newStatus,omitempty"`
	IsNotRead                    bool                  `xml:"isNotRead"`
	FromEmail                    bool                  `xml:"fromEmail"`
	DateTime                     string                `xml:"dateTime"`
	TaskId                       int                   `xml:"task>id"`
	TaskTitle                    string                `xml:"task>title"`
	ContactGeneral               int                   `xml:"contact>general"`
	ContactName                  string                `xml:"contact>name"`
	Owner                        XmlResponseUser       `xml:"owner"`
	ProjectId                    int                   `xml:"project>id"`
	ProjectTitle                 string                `xml:"project>title"`
	TaskExpectDateChangedOldDate string                `xml:"taskExpectDateChanged>oldDate"`
	TaskExpectDateChangedNewDate string                `xml:"taskExpectDateChanged>newDate"`
	TaskStartTimeChangedOldDate  string                `xml:"taskStartTimeChanged>oldDate"`
	TaskStartTimeChangedNewDate  string                `xml:"taskStartTimeChanged>newDate"`
	Files                        []XmlResponseFile     `xml:"files>file"`
	NotifiedList                 []XmlResponseUser     `xml:"notifiedList>user"`
	Analitics                    []XmlResponseAnalitic `xml:"analitics>analitic"`
}

// auth.login
type XmlRequestAuth struct {
	XMLName xml.Name `xml:"request"`
	Method  string   `xml:"method,attr"`

	Account  string `xml:"account"`
	Login    string `xml:"login"`
	Password string `xml:"password"`
}

// auth.login response
type XmlResponseAuth struct {
	XMLName xml.Name `xml:"response"`
	Status  string   `xml:"status,attr"`
	Code    string   `xml:"code"`
	Message string   `xml:"message"`

	Sid string `xml:"sid"`
}

// action.get
type XmlRequestActionGet struct {
	XMLName xml.Name `xml:"request"`
	Method  string   `xml:"method,attr"`
	Account string   `xml:"account"`
	Sid     string   `xml:"sid"`

	ActionId int `xml:"action>id"`
}

// action.get response
type XmlResponseActionGet struct {
	XMLName xml.Name `xml:"response"`
	Status  string   `xml:"status,attr"`
	Code    string   `xml:"code"`
	Message string   `xml:"message"`

	Action XmlResponseAction `xml:"action"`
}

// action.getList
type XmlRequestActionGetList struct {
	XMLName xml.Name `xml:"request"`
	Method  string   `xml:"method,attr"`
	Account string   `xml:"account"`
	Sid     string   `xml:"sid"`

	TaskId         int    `xml:"task>id,omitempty"`
	TaskGeneral    int    `xml:"task>general,omitempty"`
	ContactGeneral int    `xml:"contact>general,omitempty"`
	PageCurrent    int    `xml:"pageCurrent"`
	PageSize       int    `xml:"pageSize"`
	Sort           string `xml:"sort"`
}

type XmlResponseActions struct {
}

// action.getList response
type XmlResponseActionGetList struct {
	XMLName xml.Name `xml:"response"`
	Status  string   `xml:"status,attr"`
	Code    string   `xml:"code"`
	Message string   `xml:"message"`

	Actions struct {
		ActionsCount      int                 `xml:"count,attr"`
		ActionsTotalCount int                 `xml:"totalCount,attr"`
		Actions           []XmlResponseAction `xml:"action"`
	} `xml:"actions"`
}
