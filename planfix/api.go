package planfix

import (
	"fmt"
	"encoding/xml"
	"net/http"
	"log"
	"strings"
	"io/ioutil"
	"regexp"
)

type Api struct {
	Url       string
	ApiKey    string
	Account   string
	Sid       string
	User      string
	Password  string
	UserAgent string
}

func New(url, apiKey, account, user, password string) Api {
	return Api{
		Url:       url,
		ApiKey:    apiKey,
		Account:   account,
		User:      user,
		Password:  password,
		UserAgent: "planfix-go",
	}
}

func (a *Api) ensureAuthenticated() error {
	if a.Sid == "" {
		sid, err := a.AuthLogin(a.User, a.Password)
		if err != nil {
			log.Fatalf("Failed to authenticate to planfix.ru, %v", err)
		}
		a.Sid = sid
	}
	return nil
}

func (a Api) apiRequest(requestStruct interface{}, responseStruct interface{}) (error) {
	//xmlBytes, err := xml.MarshalIndent(requestStruct, "  ", "    ")
	xmlBytes, err := xml.Marshal(requestStruct)
	if err != nil {
		fmt.Println(err)
	}
	xmlString := xml.Header + string(xmlBytes)

	// logging
	passwordCutter := regexp.MustCompile(`<password>.*?</password>`)
	log.Printf(
		"[DEBUG] request to planfix: %s",
		passwordCutter.ReplaceAllString(string(xmlBytes), "<password>***</password>"),
	)

	httpClient := http.Client{}
	req, _ := http.NewRequest("POST", a.Url, strings.NewReader(xmlString))
	req.Header.Add("Content-Type", "application/xml; charset=utf-8")
	req.Header.Set("User-Agent", a.UserAgent)
	req.SetBasicAuth(a.ApiKey, "")
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Printf("[ERROR] Network error while request to planfix: %s", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	log.Printf(
		"[DEBUG] response from planfix: %s",
		strings.Replace(string(data), "\n", "", -1),
	)

	err = xml.Unmarshal(data, &responseStruct)
	return err
}
