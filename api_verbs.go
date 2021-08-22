package servicenow

import (
	"net/url"
)

func (c Connection) POST(payload []byte, limit string, fields string, offset string, query string, sysID string) RequestTransitive {

	t := RequestTransitive{}

	t.Connection = c

	t.Method = "POST"

	t.Payload = payload

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

	t.Parms = parameters

	t.SysID = sysID

	return t

}

func (c Connection) GET(limit string, fields string, offset string, query string, sysID string) RequestTransitive {

	t := RequestTransitive{}

	t.Connection = c

	t.Method = "GET"

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

	t.Parms = parameters

	t.SysID = sysID

	return t

}

func (c Connection) PUT(payload []byte, limit string, fields string, offset string, query string, sysID string) RequestTransitive {

	t := RequestTransitive{}

	t.Connection = c

	t.Method = "PUT"

	t.Payload = payload

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

	t.Parms = parameters

	t.SysID = sysID

	return t

}

func (c Connection) PATCH(payload []byte, limit string, fields string, offset string, query string, sysID string) RequestTransitive {

	t := RequestTransitive{}

	t.Connection = c

	t.Method = "PATCH"

	t.Payload = payload

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

	t.Parms = parameters

	t.SysID = sysID

	return t

}

func (c Connection) DELETE(sysID string) RequestTransitive {

	t := RequestTransitive{}

	t.Connection = c

	t.Method = "DELETE"

	t.SysID = sysID

	return t

}
