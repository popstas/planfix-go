package planfix

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type API struct {
	URL       string
	APIKey    string
	Account   string
	Sid       string
	User      string
	Password  string
	UserAgent string
	Logger    *log.Logger
}

func New(url, apiKey, account, user, password string) API {
	return API{
		URL:       url,
		APIKey:    apiKey,
		Account:   account,
		User:      user,
		Password:  password,
		UserAgent: "planfix-go",
		Logger:    log.New(os.Stderr, "[planfix-go] ", log.LstdFlags),
	}
}

func (a *API) ensureAuthenticated() error {
	if a.Sid == "" {
		sid, err := a.AuthLogin(a.User, a.Password)
		if err != nil {
			a.Logger.Printf("[ERROR] Failed to authenticate to planfix.ru, %v", err)
			return err
		}
		a.Sid = sid
	}
	return nil
}

func (a API) tryRequest(requestStruct XmlRequester) (status XmlResponseStatus, data []byte, err error) {
	//xmlBytes, err := xml.MarshalIndent(requestStruct, "  ", "    ")
	xmlBytes, _ := xml.Marshal(requestStruct)
	xmlString := xml.Header + string(xmlBytes)

	// logging
	passwordCutter := regexp.MustCompile(`<password>.*?</password>`)
	a.Logger.Printf(
		"[DEBUG] request to planfix: %s",
		passwordCutter.ReplaceAllString(string(xmlBytes), "<password>***</password>"),
	)

	httpClient := http.Client{}
	req, _ := http.NewRequest("POST", a.URL, strings.NewReader(xmlString))
	req.Header.Add("Content-Type", "application/xml; charset=utf-8")
	req.Header.Set("User-Agent", a.UserAgent)
	req.SetBasicAuth(a.APIKey, "")

	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Printf("[ERROR] Network error while request to planfix: %s", err)
		return status, data, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return status, data, fmt.Errorf("[ERROR] status error: %v", resp.StatusCode)
	}

	data, err = ioutil.ReadAll(resp.Body)
	a.Logger.Printf(
		"[DEBUG] response from planfix: %s",
		strings.Replace(string(data), "\n", "", -1),
	)

	err = xml.Unmarshal(data, &status)
	return status, data, err
}

func (a *API) apiRequest(requestStruct XmlRequester, responseStruct interface{}) error {
	requestStruct.SetAccount(a.Account)
	if requestStruct.GetMethod() != "auth.login" {
		if err := a.ensureAuthenticated(); err != nil {
			return err
		}
		requestStruct.SetSid(a.Sid)
	}

	var (
		status XmlResponseStatus
		data   []byte
		err    error
	)

	for try := 0; try < 2; try++ {
		status, data, err = a.tryRequest(requestStruct)
		if err != nil {
			return err
		}

		if status.Status == "ok" {
			break
		} else {
			if status.Code == "0005" { // session expired
				a.Logger.Println("[INFO] session expired, relogin")
				a.Sid = ""
				if err := a.ensureAuthenticated(); err != nil {
					return err
				}
				requestStruct.SetSid(a.Sid)
			} else {
				return fmt.Errorf(
					"%s: response status: %s, %s, %s",
					requestStruct.GetMethod(),
					status.Status,
					a.getErrorByCode(status.Code),
					status.Message,
				)
			}
		}
	}

	err = xml.Unmarshal(data, &responseStruct)
	return err
}
