package gosnow

import (
	"net/url"
)

type Connection struct {
	Instance string
	Username string
	Password string
}

type RequestTransitive struct {
	Payload []byte
	Method  string
	Connection
	Parameters url.Values
	SysID      string
	err []error
}

type Result struct {
	Status      string
	StatusCode  int
	ResultCount int
}
