package planfix

import (
	"encoding/xml"
)

// XMLRequester - любая структура запроса
type XMLRequester interface {
	SetSid(sid string)
	SetAccount(account string)
	GetMethod() string
}

// XMLRequestAuth - запрос авторизации
type XMLRequestAuth struct {
	XMLName xml.Name `xml:"request"`
	Method  string   `xml:"method,attr"`
	Account string   `xml:"account"`
	Sid     string   `xml:"sid"`
}

// SetSid задает sid
func (a *XMLRequestAuth) SetSid(sid string) {
	a.Sid = sid
}

// SetAccount задает account
func (a *XMLRequestAuth) SetAccount(account string) {
	a.Account = account
}

// GetMethod возвращает метод API (название)
func (a *XMLRequestAuth) GetMethod() string {
	return a.Method
}

// XMLResponseStatus - любой ответ, подразумевает статус ответа и код ошибки в случае ошибки
type XMLResponseStatus struct {
	XMLName xml.Name `xml:"response"`
	Status  string   `xml:"status,attr"`
	Code    string   `xml:"code"`
	Message string   `xml:"message"`
}

// XMLResponseFile - файл в ответе analitic.get
type XMLResponseFile struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// XMLResponseActionUser - юзер в ответе analitic.get
type XMLResponseActionUser struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

// XMLResponseActionAnalitic - аналитика в ответе analitic.get
type XMLResponseActionAnalitic struct {
	ID   int    `xml:"id"`
	Key  int    `xml:"key"`
	Name string `xml:"name"`
}

// XMLResponseAnaliticOptions - аналитика в ответе analitic.getOptions
type XMLResponseAnaliticOptions struct {
	ID      int                               `xml:"id"`
	Name    string                            `xml:"name"`
	GroupID int                               `xml:"group>id"`
	Fields  []XMLResponseAnaliticOptionsField `xml:"fields>field"`
}

// XMLResponseAnaliticOptionsField - поле аналитики в ответе analitic.getOptions
type XMLResponseAnaliticOptionsField struct {
	ID         int      `xml:"id"`
	Num        int      `xml:"num"`
	Name       string   `xml:"name"`
	Type       string   `xml:"type"`
	ListValues []string `xml:"list>value"`
	HandbookID int      `xml:"handbook>id"`
}

// XMLResponseAnaliticHandbookRecord - запись справочника в ответе analitic.getHandbook
type XMLResponseAnaliticHandbookRecord struct {
	Key       int                                       `xml:"key"`
	ParentKey int                                       `xml:"parentKey"`
	IsGroup   int                                       `xml:"isGroup"`
	Values    []XMLResponseAnaliticHandbookRecordValues `xml:"value"`
	ValuesMap map[string]string
}

// XMLResponseAnaliticHandbookRecordValues - значения полей в записи справочника в ответе analitic.getHandbook
type XMLResponseAnaliticHandbookRecordValues struct {
	Name        string `xml:"name,attr"`
	Value       string `xml:"value,attr"`
	IsDisplayed int    `xml:"isDisplayed,attr"`
}

// XMLResponseAction - действие в ответе action.getList
type XMLResponseAction struct {
	ID                           int                         `xml:"id"`
	Description                  string                      `xml:"description"`
	OldStatus                    int                         `xml:"statusChange>oldStatus,omitempty"`
	NewStatus                    int                         `xml:"statusChange>newStatus,omitempty"`
	IsNotRead                    bool                        `xml:"isNotRead"`
	FromEmail                    bool                        `xml:"fromEmail"`
	DateTime                     string                      `xml:"dateTime"`
	TaskID                       int                         `xml:"task>id"`
	TaskTitle                    string                      `xml:"task>title"`
	ContactGeneral               int                         `xml:"contact>general"`
	ContactName                  string                      `xml:"contact>name"`
	Owner                        XMLResponseActionUser       `xml:"owner"`
	ProjectID                    int                         `xml:"project>id"`
	ProjectTitle                 string                      `xml:"project>title"`
	TaskExpectDateChangedOldDate string                      `xml:"taskExpectDateChanged>oldDate"`
	TaskExpectDateChangedNewDate string                      `xml:"taskExpectDateChanged>newDate"`
	TaskStartTimeChangedOldDate  string                      `xml:"taskStartTimeChanged>oldDate"`
	TaskStartTimeChangedNewDate  string                      `xml:"taskStartTimeChanged>newDate"`
	Files                        []XMLResponseFile           `xml:"files>file"`
	NotifiedList                 []XMLResponseActionUser     `xml:"notifiedList>user"`
	Analitics                    []XMLResponseActionAnalitic `xml:"analitics>analitic"`
}

// XMLResponseTask - задача в ответе task.get
// TODO: добавить все поля из https://planfix.ru/docs/ПланФикс_API_task.get
type XMLResponseTask struct {
	ID           int    `xml:"id"`
	Title        string `xml:"title"`
	Description  string `xml:"description"`
	General      int    `xml:"general"`
	ProjectID    int    `xml:"project>id"`
	ProjectTitle string `xml:"project>title"`
}

// XMLResponseAnalitic - аналитика в ответе analitic.getList
type XMLResponseAnalitic struct {
	ID        int    `xml:"id"`
	Name      string `xml:"name"`
	GroupID   int    `xml:"group>id"`
	GroupName string `xml:"group>name"`
}

// XMLRequestActionAnalitic - аналитика в запросе на добавление действия analitic.add
type XMLRequestActionAnalitic struct {
	ID       int                       `xml:"id"`
	ItemData []XMLRequestAnaliticField `xml:"analiticData>itemData"`
}

// XMLRequestAnaliticField - поле аналитики в запросе на добавление действия analitic.add
type XMLRequestAnaliticField struct {
	FieldID int         `xml:"fieldId"`
	Value   interface{} `xml:"value"`
}

// XMLResponseUser - юзер в ответе user.get
// TODO: добавить все поля из https://planfix.ru/docs/ПланФикс_API_user.get
type XMLResponseUser struct {
	ID       int    `xml:"id"`
	Name     string `xml:"name"`
	LastName string `xml:"lastName"`
	Login    string `xml:"login"`
	Email    string `xml:"email"`
}

// XMLRequestAuthLogin - запрос auth.login
type XMLRequestAuthLogin struct {
	XMLName xml.Name `xml:"request"`
	Method  string   `xml:"method,attr"`

	Account  string `xml:"account"`
	Login    string `xml:"login"`
	Password string `xml:"password"`
}

// - SetSid - заглушка, чтобы реализовать интерфейс
func (a *XMLRequestAuthLogin) SetSid(sid string) {}

// SetAccount задает account
func (a *XMLRequestAuthLogin) SetAccount(account string) {
	a.Account = account
}

// GetMethod возвращает метод API (название)
func (a *XMLRequestAuthLogin) GetMethod() string {
	return a.Method
}

// XMLResponseAuth - ответ auth.login
type XMLResponseAuth struct {
	XMLName xml.Name `xml:"response"`
	Sid     string   `xml:"sid"`
}

// XMLRequestActionGet - запрос action.get
type XMLRequestActionGet struct {
	XMLRequestAuth
	XMLName  xml.Name `xml:"request"`
	ActionID int      `xml:"action>id"`
}

// XMLResponseActionGet - ответ action.get
type XMLResponseActionGet struct {
	XMLName xml.Name          `xml:"response"`
	Action  XMLResponseAction `xml:"action"`
}

// XMLRequestActionGetList - запрос action.getList
type XMLRequestActionGetList struct {
	XMLRequestAuth
	XMLName xml.Name `xml:"request"`

	TaskID         int    `xml:"task>id,omitempty"`
	TaskGeneral    int    `xml:"task>general,omitempty"`
	ContactGeneral int    `xml:"contact>general,omitempty"`
	PageCurrent    int    `xml:"pageCurrent"`
	PageSize       int    `xml:"pageSize"`
	Sort           string `xml:"sort"`
}

// XMLResponseActionGetList - ответ action.getList
type XMLResponseActionGetList struct {
	XMLName xml.Name `xml:"response"`
	Actions struct {
		ActionsCount      int                 `xml:"count,attr"`
		ActionsTotalCount int                 `xml:"totalCount,attr"`
		Actions           []XMLResponseAction `xml:"action"`
	} `xml:"actions"`
}

// XMLRequestActionAdd - запрос action.add
type XMLRequestActionAdd struct {
	XMLRequestAuth
	XMLName xml.Name `xml:"request"`

	Description    string                     `xml:"action>description"`
	TaskID         int                        `xml:"action>task>id,omitempty"`
	TaskGeneral    int                        `xml:"action>task>general,omitempty"`
	ContactGeneral int                        `xml:"action>contact>general,omitempty"`
	TaskNewStatus  int                        `xml:"action>taskNewStatus,omitempty"`
	NotifiedList   []XMLResponseUser          `xml:"action>notifiedList>user,omitempty"`
	IsHidden       int                        `xml:"action>isHidden"`
	Owner          XMLResponseUser            `xml:"action>owner,omitempty"`
	DateTime       string                     `xml:"action>dateTime,omitempty"`
	Analitics      []XMLRequestActionAnalitic `xml:"action>analitics>analitic,omitempty"`
}

// XMLResponseActionAdd - ответ action.add
type XMLResponseActionAdd struct {
	XMLName  xml.Name `xml:"response"`
	ActionID int      `xml:"action>id"`
}

// XMLRequestAnaliticGetList - запрос analitic.getList
type XMLRequestAnaliticGetList struct {
	XMLRequestAuth
	XMLName xml.Name `xml:"request"`

	AnaliticGroupID int `xml:"analiticGroupId,omitempty"`
}

// XMLResponseAnaliticGetList - ответ analitic.getList
type XMLResponseAnaliticGetList struct {
	XMLName   xml.Name `xml:"response"`
	Analitics struct {
		AnaliticsCount      int                   `xml:"count,attr"`
		AnaliticsTotalCount int                   `xml:"totalCount,attr"`
		Analitics           []XMLResponseAnalitic `xml:"analitic"`
	} `xml:"analitics"`
}

// XMLRequestAnaliticGetOptions - запрос analitic.getOptions
type XMLRequestAnaliticGetOptions struct {
	XMLRequestAuth
	XMLName xml.Name `xml:"request"`

	AnaliticID int `xml:"analitic>id"`
}

// XMLResponseAnaliticGetOptions - ответ analitic.getOptions
type XMLResponseAnaliticGetOptions struct {
	XMLName  xml.Name                   `xml:"response"`
	Analitic XMLResponseAnaliticOptions `xml:"analitic"`
}

// XMLRequestAnaliticGetHandbook - запрос analitic.getHandbook
type XMLRequestAnaliticGetHandbook struct {
	XMLRequestAuth
	XMLName xml.Name `xml:"request"`

	HandbookID int `xml:"handbook>id"`
}

// XMLResponseAnaliticGetHandbook - ответ analitic.getHandbook
type XMLResponseAnaliticGetHandbook struct {
	XMLName xml.Name                            `xml:"response"`
	Records []XMLResponseAnaliticHandbookRecord `xml:"records>record"`
}

// XMLRequestTaskGet - запрос task.get
type XMLRequestTaskGet struct {
	XMLRequestAuth
	XMLName xml.Name `xml:"request"`

	TaskID      int `xml:"task>id,omitempty"`
	TaskGeneral int `xml:"task>general,omitempty"`
}

// XMLResponseTaskGet - ответ task.get response
type XMLResponseTaskGet struct {
	XMLName xml.Name        `xml:"response"`
	Task    XMLResponseTask `xml:"task"`
}

// XMLRequestUserGet - запрос user.get
type XMLRequestUserGet struct {
	XMLRequestAuth
	XMLName xml.Name `xml:"request"`

	UserID int `xml:"user>id,omitempty"`
}

// XMLResponseUserGet - ответ user.get
type XMLResponseUserGet struct {
	XMLName xml.Name        `xml:"response"`
	User    XMLResponseUser `xml:"user"`
}
