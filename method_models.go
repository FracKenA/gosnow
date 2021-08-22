package servicenow

import (
	"net/url"
)

type Connection struct {
	Instance string
	Domain   string
	Username string
	Password string
}

type RequestTransitive struct {
	Payload []byte
	Method  string
	Connection
	Parms url.Values
	SysID string
}

type Result struct {
	Status     string
	StatusCode int
	Results    int
}
