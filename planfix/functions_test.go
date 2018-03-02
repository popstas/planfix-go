package planfix_test

import (
	"encoding/xml"
	"github.com/popstas/planfix-go/planfix"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func assert(t *testing.T, data interface{}, expected interface{}) {
	if data != expected {
		t.Errorf("Expected %v, got, %v", expected, data)
	}
}
func expectError(t *testing.T, err error, msg string) {
	if err == nil {
		t.Errorf("Expected error, got success %v", msg)
	}
}
func expectSuccess(t *testing.T, err error, msg string) {
	if err != nil {
		t.Errorf("Expected success, got %v %v", err, msg)
	}
}

type requestStruct struct {
	XMLName xml.Name `xml:"request"`
	Method  string   `xml:"method,attr"`
	Account string   `xml:"account"`
	Sid     string   `xml:"sid"`
}

func fixtureFromFile(fixtureName string) string {
	buf, _ := ioutil.ReadFile("../tests/fixtures/" + fixtureName)
	return string(buf)
}

type MockedServer struct {
	*httptest.Server
	Requests  [][]byte
	Responses []string // fifo queue of answers
}

func NewMockedServer(responses []string) *MockedServer {
	s := &MockedServer{
		Requests:  [][]byte{},
		Responses: responses,
	}

	s.Server = httptest.NewServer(s)
	return s
}

func (s *MockedServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	lastRequest, err := ioutil.ReadAll(req.Body)
	//body := string(lastRequest)

	rs := requestStruct{}
	err = xml.Unmarshal(lastRequest, &rs)
	if err != nil {
		panic(err)
	}
	s.Requests = append(s.Requests, lastRequest)
	answer := s.Responses[0]

	// simulate network error
	if answer == "panic" {
		panic(err)
	}

	// simulate 502
	if answer == "502" {
		resp.WriteHeader(http.StatusBadGateway)
	}
	s.Responses = s.Responses[1:]
	resp.Write([]byte(answer))
}

func newAPI(responses []string) planfix.API {
	ms := NewMockedServer(responses)
	api := planfix.New(ms.URL, "apiKey", "account", "user", "password")
	api.Sid = "123"
	return api
}

func TestAPI_ErrorCode(t *testing.T) {
	api := newAPI([]string{fixtureFromFile("error.xml")})
	_, err := api.AuthLogin(api.User, api.Password)
	expectError(t, err, "TestAPI_ErrorCode")
}

func TestAPI_ErrorCodeUnknown(t *testing.T) {
	api := newAPI([]string{fixtureFromFile("error.unknown.xml")})
	_, err := api.AuthLogin(api.User, api.Password)
	if !strings.Contains(string(err.Error()), "Неизвестная ошибка") {
		t.Error("Failed to output unknown error")
	}
	expectError(t, err, "TestAPI_ErrorCodeUnknown")
}

// auth.login
func TestAPI_AuthLogin(t *testing.T) {
	api := newAPI([]string{fixtureFromFile("auth.login.xml")})
	api.Sid = ""
	answer, err := api.AuthLogin(api.User, api.Password)

	expectSuccess(t, err, "TestAPI_AuthLogin")
	assert(t, answer, "sid_after_login")
}

// fail to authenticate
func TestAPI_EnsureAuthenticatedFailed(t *testing.T) {
	api := newAPI([]string{fixtureFromFile("error.xml")})
	api.Sid = ""
	_, err := api.ActionGet(456)

	expectError(t, err, "TestAPI_EnsureAuthenticatedFailed")
}

// authenticate before api request if not authenticated
func TestAPI_EnsureAuthenticated(t *testing.T) {
	api := newAPI([]string{
		fixtureFromFile("auth.login.xml"),
		fixtureFromFile("action.get.xml"),
		fixtureFromFile("action.get.xml"),
	})
	api.Sid = ""
	_, _ = api.ActionGet(456)
	assert(t, api.Sid, "sid_after_login")

	api.Sid = "789"
	_, _ = api.ActionGet(456)
	assert(t, api.Sid, "789")
}

// reauthenticate if session expired
func TestAPI_AuthenticatedExpire(t *testing.T) {
	api := newAPI([]string{
		fixtureFromFile("error.sessionExpired.xml"),
		fixtureFromFile("auth.login.xml"),
		fixtureFromFile("action.get.xml"),
	})
	action, err := api.ActionGet(456)

	expectSuccess(t, err, "TestAPI_AuthenticatedExpire")
	assert(t, action.Action.TaskID, 1128468)
}

// error response while reauthenticate
func TestAPI_AuthenticatedExpireFailed(t *testing.T) {
	api := newAPI([]string{
		fixtureFromFile("error.sessionExpired.xml"),
		fixtureFromFile("error.xml"),
	})
	_, err := api.ActionGet(456)

	expectError(t, err, "TestAPI_AuthenticatedExpireFailed")
}

// error response after reauthenticated session
func TestAPI_AuthenticatedExpireFailedAfter(t *testing.T) {
	api := newAPI([]string{
		fixtureFromFile("error.sessionExpired.xml"),
		fixtureFromFile("auth.login.xml"),
		fixtureFromFile("error.xml"),
	})
	_, err := api.ActionGet(456)

	expectError(t, err, "TestAPI_AuthenticatedExpireFailedAfter")
}

// invalid xml
func TestAPI_InvalidXML(t *testing.T) {
	api := newAPI([]string{fixtureFromFile("error.invalidXML.xml")})
	_, err := api.ActionGet(123)
	expectError(t, err, "TestAPI_InvalidXML")
}

// network error
func TestAPI_NetworkError(t *testing.T) {
	api := newAPI([]string{"panic"})
	_, err := api.ActionGet(123)
	expectError(t, err, "TestAPI_NetworkError")
}

// 502 error
func TestAPI_502Error(t *testing.T) {
	api := newAPI([]string{"502"})
	_, err := api.ActionGet(123)
	expectError(t, err, "TestAPI_502Error")
}

// action.get
func TestAPI_ActionGet(t *testing.T) {
	api := newAPI([]string{fixtureFromFile("action.get.xml")})
	var action planfix.XMLResponseActionGet
	action, err := api.ActionGet(123)

	expectSuccess(t, err, "TestAPI_ActionGet")
	assert(t, action.Action.TaskID, 1128468)
}

// action.getList
func TestAPI_ActionGetList(t *testing.T) {
	api := newAPI([]string{fixtureFromFile("action.getList.xml")})
	request := planfix.XMLRequestActionGetList{
		TaskGeneral: 525330,
	}
	var actionList planfix.XMLResponseActionGetList
	actionList, err := api.ActionGetList(request)

	expectSuccess(t, err, "TestAPI_ActionGetList")
	assert(t, actionList.Actions.ActionsTotalCount, 31)
}

// analitic.getList
func TestAPI_AnaliticGetList(t *testing.T) {
	api := newAPI([]string{fixtureFromFile("analitic.getList.xml")})
	var analiticList planfix.XMLResponseAnaliticGetList
	analiticList, err := api.AnaliticGetList(0)

	expectSuccess(t, err, "TestAPI_AnaliticGetList")
	assert(t, analiticList.Analitics.AnaliticsTotalCount, 2)
}

// analitic.getOptiions
func TestAPI_AnaliticGetOptions(t *testing.T) {
	api := newAPI([]string{fixtureFromFile("analitic.getOptions.xml")})
	var analitic planfix.XMLResponseAnaliticGetOptions
	analitic, err := api.AnaliticGetOptions(123)

	expectSuccess(t, err, "TestAPI_AnaliticGetOptions")
	assert(t, analitic.Analitic.GroupID, 1)
	assert(t, analitic.Analitic.Fields[0].HandbookID, 131)
}

// analitic.getHandbook
func TestAPI_AnaliticGetHandbook(t *testing.T) {
	api := newAPI([]string{fixtureFromFile("analitic.getHandbook.xml")})
	var handbook planfix.XMLResponseAnaliticGetHandbook
	handbook, err := api.AnaliticGetHandbook(123)

	expectSuccess(t, err, "TestAPI_AnaliticGetHandbook")
	assert(t, handbook.Records[4].Values[0].Value, "Поминутная работа программиста")
	assert(t, handbook.Records[4].ValuesMap["Название"], "Поминутная работа программиста")
}

// action.add
func TestAPI_ActionAdd(t *testing.T) {
	api := newAPI([]string{fixtureFromFile("action.add.xml")})
	request := planfix.XMLRequestActionAdd{
		TaskGeneral: 123,
		Description: "asdf",
	}
	var actionAdded planfix.XMLResponseActionAdd
	actionAdded, err := api.ActionAdd(request)

	expectSuccess(t, err, "TestAPI_ActionAdd")
	assert(t, actionAdded.ActionID, 123)
}

// action.add both task and contact defined
func TestAPI_ActionAddBothTaskContact(t *testing.T) {
	api := newAPI([]string{fixtureFromFile("action.add.xml")})
	request := planfix.XMLRequestActionAdd{
		TaskGeneral:    123,
		ContactGeneral: 123,
		Description:    "asdf",
	}
	_, err := api.ActionAdd(request)

	expectError(t, err, "TestAPI_ActionAddBothTaskContact")
}

// task.get
func TestAPI_TaskGet(t *testing.T) {
	api := newAPI([]string{fixtureFromFile("task.get.xml")})
	var task planfix.XMLResponseTaskGet
	task, err := api.TaskGet(123, 0)

	expectSuccess(t, err, "TestAPI_TaskGet")
	assert(t, task.Task.ProjectID, 9830)
}

// user.get
func TestAPI_UserGet(t *testing.T) {
	api := newAPI([]string{fixtureFromFile("user.get.xml")})
	var user planfix.XMLResponseUserGet
	user, err := api.UserGet(0)

	expectSuccess(t, err, "TestAPI_UserGet")
	assert(t, user.User.Login, "popstas")
}
