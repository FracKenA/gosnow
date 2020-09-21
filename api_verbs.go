package gosnow

import (
	"net/url"
)

func (connection Connection) POST(payload []byte, limit string, fields string, offset string, query string, sysID string) RequestTransitive {

	requestTransitive := RequestTransitive{}

	requestTransitive.Connection = connection

	requestTransitive.Method = "POST"

	if payload == nil {

	} else {
		requestTransitive.Payload = payload
	}

	parameters := url.Values{}

	if limit != "" {
		parameters.Add("sysparm_limit", limit)
	}
	if fields != "" {
		parameters.Add("sysparm_fields", fields)
	}
	if offset != "" {
		parameters.Add("sysparm_offset", offset)
	}
	if query != "" {
		parameters.Add("sysparm_query", query)
	}

	requestTransitive.Parameters = parameters

	if sysID == "" {

	} else {
		requestTransitive.SysID = sysID
	}

	return requestTransitive

}

func (connection Connection) GET(limit string, fields string, offset string, query string, sysID string) RequestTransitive {

	requestTransitive := RequestTransitive{}

	requestTransitive.Connection = connection

	requestTransitive.Method = "GET"

	parameters := url.Values{}

	if limit != "" {
		parameters.Add("sysparm_limit", limit)
	}
	if fields != "" {
		parameters.Add("sysparm_fields", fields)
	}
	if offset != "" {
		parameters.Add("sysparm_offset", offset)
	}
	if query != "" {
		parameters.Add("sysparm_query", query)
	}

	requestTransitive.Parameters = parameters

	requestTransitive.SysID = sysID

	return requestTransitive

}

func (connection Connection) PUT(payload []byte, limit string, fields string, offset string, query string, sysID string) RequestTransitive {


	requestTransitive := RequestTransitive{}

	requestTransitive.Connection = connection

	requestTransitive.Method = "PUT"

	requestTransitive.Payload = payload

	parameters := url.Values{}

	if limit != "" {
		parameters.Add("sysparm_limit", limit)
	}
	if fields != "" {
		parameters.Add("sysparm_fields", fields)
	}
	if offset != "" {
		parameters.Add("sysparm_offset", offset)
	}
	if query != "" {
		parameters.Add("sysparm_query", query)

	}

	requestTransitive.Parameters = parameters

	requestTransitive.SysID = sysID

	return requestTransitive

}

func (connection Connection) PATCH(payload []byte, limit string, fields string, offset string, query string, sysID string) RequestTransitive {


	requestTransitive := RequestTransitive{}

	requestTransitive.Connection = connection

	requestTransitive.Method = "PATCH"

	requestTransitive.Payload = payload

	parameters := url.Values{}

	if limit != "" {
		parameters.Add("sysparm_limit", limit)
	}
	if fields != "" {
		parameters.Add("sysparm_fields", fields)
	}
	if offset != "" {
		parameters.Add("sysparm_offset", offset)
	}
	if query != "" {
		parameters.Add("sysparm_query", query)

	}

	requestTransitive.Parameters = parameters

	requestTransitive.SysID = sysID

	return requestTransitive

}

func (connection Connection) DELETE(sysID string) RequestTransitive {

	requestTransitive := RequestTransitive{}

	requestTransitive.Connection = connection

	requestTransitive.Method = "DELETE"

	requestTransitive.SysID = sysID

	return requestTransitive

}
