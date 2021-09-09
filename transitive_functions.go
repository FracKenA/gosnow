package servicenow

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func InitializeConnection(instance, domain, username, password string) Connection {
	return Connection{
		instance,
		domain,
		username,
		password,
	}
}

func AssembleRequest(t RequestTransitive, table string) *http.Request {
	urlBuild := fmt.Sprintf("https://%s.%s", t.Connection.Instance, t.Connection.Domain)
	fmt.Sprintf(urlBuild)
	Url, err := url.Parse(urlBuild)
	if err != nil {
		fmt.Printf("Error%s\n", err)
	}

	if t.SysID != "" {
		Url.Path += "/api/now/v2/table/" + table + "/" + t.SysID
	} else {
		Url.Path += "/api/now/v2/table/" + table
	}

	z := t.Parms.Encode()
	z = strings.ReplaceAll(z, "+", "%20")

	Url.RawQuery = z

	// fmt.Printf("Encoded URL is %q\n", Url.String())

	payload := bytes.NewReader(t.Payload)

	req, err := http.NewRequest(t.Method, Url.String(), payload)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
	req.SetBasicAuth(t.Connection.Username, t.Connection.Password)
	req.Header.Add("Accept", "application/json")

	return req
}

func CloseResponse(toClose *http.Response) {
	err := toClose.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
}
