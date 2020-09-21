package gosnow

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// InitializeConnection is used to import the instance, username and password to connect to a ServiceNow instance.
// No error is returned.
// It returns the connection strings required for all methods and functions in this module.
// The required variables are:
//   - Your ServiceNow instance name. Exclude the https:// prefix as well as .service-now.com
//   - Your ServiceNow instance username.
//   - Your ServiceNow instance password.
func InitializeConnection(instance, username, password string) *Connection {
	return &Connection{
		instance,
		username,
		password,
	}
}

// AssembleRequest will take the API Request details and create a propperly formatted API request for ServiceNow
// An error is returned
// It returns the HTTP Request details.
// It receives the RequestTransitive Type and the desired table to be queried.
func AssembleRequest(requestTransitive RequestTransitive, table string) *http.Request {
	// Creates the ServiceNow URL from the requestTransitive Type that is populated by the InitializeConnection Function.
	Url, err := url.Parse("https://" + requestTransitive.Connection.Instance + ".service-now.com")
	if err != nil {
		requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
	}

	// Appends the v2 API endpoint based on if the desired endpoint is a table or a specific SysID
	if requestTransitive.SysID != "" {
		Url.Path += "/api/now/v2/table/" + table + "/" + requestTransitive.SysID
	} else {
		Url.Path += "/api/now/v2/table/" + table
	}

	// Encode the URL parameters.
	encode := requestTransitive.Parameters.Encode()

	// Fix a disconnect with URL encoding per RFC3986.
	encode = strings.ReplaceAll(encode,"+","%20")

	// Add the corrected encoded parameters to the URL.
	Url.RawQuery = encode

	var request *http.Request

	// Create the API Request's Payload if exists.
	var payload *bytes.Reader
	if requestTransitive.Payload != nil {
		payload = bytes.NewReader(requestTransitive.Payload)
		request, err = http.NewRequest(requestTransitive.Method, Url.String(), payload)
		if err != nil {
			requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
		}
	} else {
		request, err = http.NewRequest(requestTransitive.Method, Url.String(), nil)
		if err != nil {
			requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
		}
	}

	// Create the full request.


	// Set the request's authentication
	request.SetBasicAuth(requestTransitive.Connection.Username, requestTransitive.Connection.Password)

	// Set the request's header
	request.Header.Add("Accept", "application/json")

	// return the request to the table functions
	return request
}
