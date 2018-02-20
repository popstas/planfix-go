package planfix_test

import (
	"github.com/popstas/planfix-go/planfix"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockedServer struct {
	*httptest.Server
	Requests [][]byte
	Response string
}

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

func NewMockedServer(fileName string) *MockedServer {
	buf, _ := ioutil.ReadFile(fileName)

	s := &MockedServer{
		Requests: [][]byte{},
		Response: string(buf),
	}

	s.Server = httptest.NewServer(s)

	return s
}

func (s *MockedServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	lastRequest, err := ioutil.ReadAll(req.Body)
	body := string(lastRequest)
	if err != nil {
		log.Println(body)
		panic(err)
	}
	s.Requests = append(s.Requests, lastRequest)
	resp.Write([]byte(s.Response))
}

func newApi(fixtureFileName string) planfix.Api {
	ms := NewMockedServer(fixtureFileName)
	api := planfix.New(ms.URL, "apiKey", "account", "user", "password")
	api.Sid = "123"
	return api
}

func TestApi_ErrorCode(t *testing.T) {
	api := newApi("../tests/fixtures/error.xml")
	_, err := api.AuthLogin(api.User, api.Password)
	expectError(t, err, "TestApi_ErrorCode")
}

// auth.login
func TestApi_AuthLogin(t *testing.T) {
	api := newApi("../tests/fixtures/auth.login.xml")
	api.Sid = ""
	answer, err := api.AuthLogin(api.User, api.Password)

	expectSuccess(t, err, "TestApi_AuthLogin")
	assert(t, answer, "123")
}

// authenticate before api request if not authenticated
func TestApi_EnsureAuthenticated(t *testing.T) {
	api := newApi("../tests/fixtures/auth.login.xml")
	api.Sid = ""
	_, _ = api.ActionGet(456)
	assert(t, api.Sid, "123")

	api.Sid = "789"
	_, _ = api.ActionGet(456)
	assert(t, api.Sid, "789")
}

// action.get
func TestApi_ActionGet(t *testing.T) {
	api := newApi("../tests/fixtures/action.get.xml")
	var action planfix.XmlResponseActionGet
	action, err := api.ActionGet(123)

	expectSuccess(t, err, "TestApi_ActionGet")
	assert(t, action.Action.TaskId, 1128468)
}

// action.getList
func TestApi_ActionGetList(t *testing.T) {
	api := newApi("../tests/fixtures/action.getList.xml")
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
	api := newApi("../tests/fixtures/analitic.getList.xml")
	var analiticList planfix.XmlResponseAnaliticGetList
	analiticList, err := api.AnaliticGetList(0)

	expectSuccess(t, err, "TestApi_AnaliticGetList")
	assert(t, analiticList.Analitics.AnaliticsTotalCount, 2)
}

// analitic.getOptiions
func TestApi_AnaliticGetOptions(t *testing.T) {
	api := newApi("../tests/fixtures/analitic.getOptions.xml")
	var analitic planfix.XmlResponseAnaliticGetOptions
	analitic, err := api.AnaliticGetOptions(123)

	expectSuccess(t, err, "TestApi_AnaliticGetOptions")
	assert(t, analitic.Analitic.GroupId, 1)
}

// action.add
func TestApi_ActionAdd(t *testing.T) {
	api := newApi("../tests/fixtures/action.add.xml")
	request := planfix.XmlRequestActionAdd{
		TaskGeneral: 123,
		Description: "asdf",
	}
	var actionAdded planfix.XmlResponseActionAdd
	actionAdded, err := api.ActionAdd(request)

	expectSuccess(t, err, "TestApi_ActionAdd")
	assert(t, actionAdded.ActionId, 123)
}

// task.get
func TestApi_TaskGet(t *testing.T) {
	api := newApi("../tests/fixtures/task.get.xml")
	var task planfix.XmlResponseTaskGet
	task, err := api.TaskGet(123, 0)

	expectSuccess(t, err, "TestApi_TaskGet")
	assert(t, task.Task.ProjectId, 9830)
}

// user.get
func TestApi_UserGet(t *testing.T) {
	api := newApi("../tests/fixtures/user.get.xml")
	var user planfix.XmlResponseUserGet
	user, err := api.UserGet(0)

	expectSuccess(t, err, "TestApi_UserGet")
	assert(t, user.User.Login, "popstas")
}
