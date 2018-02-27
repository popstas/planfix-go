package planfix_test

import (
	"encoding/xml"
	"github.com/popstas/planfix-go/planfix"
	"io/ioutil"
	"log"
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
	body := string(lastRequest)

	rs := requestStruct{}
	err = xml.Unmarshal(lastRequest, &rs)
	if err != nil {
		log.Println(body)
		panic(err)
	}

	if err != nil {
		log.Println(body)
		panic(err)
	}
	s.Requests = append(s.Requests, lastRequest)
	answer := s.Responses[0]
	s.Responses = s.Responses[1:]
	resp.Write([]byte(answer))
}

func newApi(responses []string) planfix.Api {
	ms := NewMockedServer(responses)
	api := planfix.New(ms.URL, "apiKey", "account", "user", "password")
	api.Sid = "123"
	return api
}

func TestApi_ErrorCode(t *testing.T) {
	api := newApi([]string{fixtureFromFile("error.xml")})
	_, err := api.AuthLogin(api.User, api.Password)
	expectError(t, err, "TestApi_ErrorCode")
}

func TestApi_ErrorCodeUnknown(t *testing.T) {
	api := newApi([]string{fixtureFromFile("error.unknown.xml")})
	_, err := api.AuthLogin(api.User, api.Password)
	if !strings.Contains(string(err.Error()), "Неизвестная ошибка") {
		t.Error("Failed to output unknown error")
	}
	expectError(t, err, "TestApi_ErrorCodeUnknown")
}

// auth.login
func TestApi_AuthLogin(t *testing.T) {
	api := newApi([]string{fixtureFromFile("auth.login.xml")})
	api.Sid = ""
	answer, err := api.AuthLogin(api.User, api.Password)

	expectSuccess(t, err, "TestApi_AuthLogin")
	assert(t, answer, "sid_after_login")
}

// authenticate before api request if not authenticated
func TestApi_EnsureAuthenticated(t *testing.T) {
	api := newApi([]string{
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
func TestApi_AuthenticatedExpire(t *testing.T) {
	api := newApi([]string{
		fixtureFromFile("error.sessionExpired.xml"),
		fixtureFromFile("auth.login.xml"),
		fixtureFromFile("action.get.xml"),
	})
	action, err := api.ActionGet(456)

	expectSuccess(t, err, "TestApi_AuthenticatedExpire")
	assert(t, action.Action.TaskId, 1128468)
}

// error response after reauthenticated session
func TestApi_AuthenticatedExpireFailed(t *testing.T) {
	api := newApi([]string{
		fixtureFromFile("error.sessionExpired.xml"),
		fixtureFromFile("auth.login.xml"),
		fixtureFromFile("error.xml"),
	})
	_, err := api.ActionGet(456)

	expectError(t, err, "TestApi_AuthenticatedExpireFailed")
}

// action.get
func TestApi_ActionGet(t *testing.T) {
	api := newApi([]string{fixtureFromFile("action.get.xml")})
	var action planfix.XmlResponseActionGet
	action, err := api.ActionGet(123)

	expectSuccess(t, err, "TestApi_ActionGet")
	assert(t, action.Action.TaskId, 1128468)
}

// action.getList
func TestApi_ActionGetList(t *testing.T) {
	api := newApi([]string{fixtureFromFile("action.getList.xml")})
	request := planfix.XmlRequestActionGetList{
		TaskGeneral: 525330,
	}
	var actionList planfix.XmlResponseActionGetList
	actionList, err := api.ActionGetList(request)

	expectSuccess(t, err, "TestApi_ActionGetList")
	assert(t, actionList.Actions.ActionsTotalCount, 31)
}

// analitic.getList
func TestApi_AnaliticGetList(t *testing.T) {
	api := newApi([]string{fixtureFromFile("analitic.getList.xml")})
	var analiticList planfix.XmlResponseAnaliticGetList
	analiticList, err := api.AnaliticGetList(0)

	expectSuccess(t, err, "TestApi_AnaliticGetList")
	assert(t, analiticList.Analitics.AnaliticsTotalCount, 2)
}

// analitic.getOptiions
func TestApi_AnaliticGetOptions(t *testing.T) {
	api := newApi([]string{fixtureFromFile("analitic.getOptions.xml")})
	var analitic planfix.XmlResponseAnaliticGetOptions
	analitic, err := api.AnaliticGetOptions(123)

	expectSuccess(t, err, "TestApi_AnaliticGetOptions")
	assert(t, analitic.Analitic.GroupId, 1)
	assert(t, analitic.Analitic.Fields[0].HandbookId, 131)
}

// action.add
func TestApi_ActionAdd(t *testing.T) {
	api := newApi([]string{fixtureFromFile("action.add.xml")})
	request := planfix.XmlRequestActionAdd{
		TaskGeneral: 123,
		Description: "asdf",
	}
	var actionAdded planfix.XmlResponseActionAdd
	actionAdded, err := api.ActionAdd(request)

	expectSuccess(t, err, "TestApi_ActionAdd")
	assert(t, actionAdded.ActionId, 123)
}

// action.add both task and contact defined
func TestApi_ActionAddBothTaskContact(t *testing.T) {
	api := newApi([]string{fixtureFromFile("action.add.xml")})
	request := planfix.XmlRequestActionAdd{
		TaskGeneral:    123,
		ContactGeneral: 123,
		Description:    "asdf",
	}
	_, err := api.ActionAdd(request)

	expectError(t, err, "TestApi_ActionAddBothTaskContact")
}

// task.get
func TestApi_TaskGet(t *testing.T) {
	api := newApi([]string{fixtureFromFile("task.get.xml")})
	var task planfix.XmlResponseTaskGet
	task, err := api.TaskGet(123, 0)

	expectSuccess(t, err, "TestApi_TaskGet")
	assert(t, task.Task.ProjectId, 9830)
}

// user.get
func TestApi_UserGet(t *testing.T) {
	api := newApi([]string{fixtureFromFile("user.get.xml")})
	var user planfix.XmlResponseUserGet
	user, err := api.UserGet(0)

	expectSuccess(t, err, "TestApi_UserGet")
	assert(t, user.User.Login, "popstas")
}
