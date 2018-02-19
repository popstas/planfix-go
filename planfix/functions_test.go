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
	if err != nil {
		log.Println(err)
	} else {
		t.Error("Expected error, got success")
	}
}

// auth.login
func TestApi_AuthLogin(t *testing.T) {
	api := newApi("../tests/fixtures/auth.login.xml")
	api.Sid = ""
	answer, err := api.AuthLogin(api.User, api.Password)
	if err != nil {
		t.Error(err)
	}
	if answer != "123" {
		t.Error("Expected 123, got ", answer)
	}
}

// authenticate before api request if not authenticated
func TestApi_EnsureAuthenticated(t *testing.T) {
	api := newApi("../tests/fixtures/auth.login.xml")
	api.Sid = ""
	_, _ = api.ActionGet(456)
	if api.Sid != "123" {
		t.Error("Expected api.Sid is 123, got ", api.Sid)
	}

	api.Sid = "789"
	_, _ = api.ActionGet(456)
	if api.Sid != "789" {
		t.Error("Expected api.Sid is 789, got ", api.Sid)
	}
}

// action.get
func TestApi_ActionGet(t *testing.T) {
	api := newApi("../tests/fixtures/action.get.xml")
	var action planfix.XmlResponseActionGet
	action, err := api.ActionGet(123)
	if err != nil {
		t.Error(err)
	}
	if action.Status != "ok" {
		t.Error("Expected ok, got ", action.Status)
	}
}

// action.getList
func TestApi_ActionGetList(t *testing.T) {
	api := newApi("../tests/fixtures/action.getList.xml")
	request := planfix.XmlRequestActionGetList{
		TaskGeneral: 525330,
	}
	var actionList planfix.XmlResponseActionGetList
	actionList, err := api.ActionGetList(request)
	if err != nil {
		t.Error(err)
	}
	if actionList.Status != "ok" {
		t.Error("Expected ok, got ", actionList.Status)
	}
}

// analitic.getList
func TestApi_AnaliticGetList(t *testing.T) {
	api := newApi("../tests/fixtures/analitic.getList.xml")
	var analiticList planfix.XmlResponseAnaliticGetList
	analiticList, err := api.AnaliticGetList(0)
	if err != nil {
		t.Error(err)
	}
	if analiticList.Status != "ok" {
		t.Error("Expected ok, got ", analiticList.Status)
	}
}

// analitic.getOptiions
func TestApi_AnaliticGetOptions(t *testing.T) {
	api := newApi("../tests/fixtures/analitic.getOptions.xml")
	var task planfix.XmlResponseAnaliticGetOptions
	task, err := api.AnaliticGetOptions(123)
	if err != nil {
		t.Error(err)
	}
	if task.Status != "ok" {
		t.Error("Expected ok, got ", task.Status)
	}
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
	if err != nil {
		t.Error(err)
	}
	if actionAdded.Status != "ok" {
		t.Error("Expected ok, got ", actionAdded.Status)
	}
	if actionAdded.ActionId != 123 {
		t.Error("Expected 123, got ", actionAdded.ActionId)
	}
}

// action.get
func TestApi_TaskGet(t *testing.T) {
	api := newApi("../tests/fixtures/task.get.xml")
	var task planfix.XmlResponseTaskGet
	task, err := api.TaskGet(123, 0)
	if err != nil {
		t.Error(err)
	}
	if task.Status != "ok" {
		t.Error("Expected ok, got ", task.Status)
	}
}
