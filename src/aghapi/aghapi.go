package aghapi

import (
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"strings"
)

type aghapi struct {
	AghHost string
	AghUser string
	AghPass string
	Verbose bool
	Client  *http.Client
}

func New(host, user, pass string, verbose bool) aghapi {
	jar, _ := cookiejar.New(nil)
	res := aghapi{
		host,
		user,
		pass,
		verbose,
		&http.Client{
			Jar: jar,
		},
	}
	return res
}

func (a aghapi) Login() bool {
	suc, resp := a.JsonRequest("login", `{"name": "`+a.AghUser+`", "password": "`+a.AghPass+`" }`)
	if suc && resp == "OK" {
		return true
	}
	return false
}

func (a aghapi) GetDnsInfo() (bool, string) {
	req, err := http.NewRequest("GET", a.AghHost+"dns_info", nil)
	suc, resp := a.ProcessRequest(req, err)
	if suc && resp == "Forbidden" {
		if a.Login() {
			req, err := http.NewRequest("GET", a.AghHost+"dns_info", nil)
			suc, resp := a.ProcessRequest(req, err)
			if resp == "Forbidden" {
				return false, resp
			}
			return suc, resp
		}
	}
	return suc, resp
}

func (a aghapi) SetDnsInfo(data string) bool {
	suc, resp := a.JsonRequest("dns_config", data)
	if suc && resp == "Forbidden" {
		if a.Login() {
			suc, resp := a.JsonRequest("dns_config", data)
			if resp == "Forbidden" {
				return false
			}
			return suc
		}
	}
	return suc
}

func (a aghapi) JsonRequest(path, bodydata string) (bool, string) {
	payload := strings.NewReader(bodydata)
	req, err := http.NewRequest("POST", a.AghHost+path, payload)
	req.Header.Add("Content-Type", "application/json")
	return a.ProcessRequest(req, err)
}

func (a aghapi) GetRequest(path string) (bool, string) {
	req, err := http.NewRequest("GET", a.AghHost+path, nil)
	return a.ProcessRequest(req, err)
}

func (a aghapi) ProcessRequest(req *http.Request, err error) (bool, string) {
	if err == nil {
		res, err := a.Client.Do(req)
		if err == nil {
			defer res.Body.Close()
			body, err := ioutil.ReadAll(res.Body)
			if err == nil {
				return true, strings.TrimSpace(string(body))
			}
		}
	}
	return false, ""
}
