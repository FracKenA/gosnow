package servicenow

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (t RequestTransitive) Incident() (IncidentResultsArray, Result, error) {

	var i IncidentResultsArray
	req := AssembleRequest(t, "incident")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	defer CloseResponse(res)

	var e error
	var h Result

	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		err = json.Unmarshal(body, &i)
		if err != nil {
			fmt.Printf("Error%s\n", err)
		}
		l := res.Header.Get("X-Total-Count")

		m, err := strconv.Atoi(l)
		if err != nil {
			fmt.Println("Error discovering number of results")
		}

		h.Status = res.Status
		h.StatusCode = res.StatusCode
		h.Results = m

	} else {
		e = fmt.Errorf("the request returned a status %s. the full details are %s", res.Status, body)
	}

	return i, h, e
}

func (t RequestTransitive) Case() (CaseResultsArray, Result, error) {

	var i CaseResultsArray
	req := AssembleRequest(t, "sn_customerservice_case")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	defer CloseResponse(res)

	var e error
	var h Result

	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		err = json.Unmarshal(body, &i)
		if err != nil {
			fmt.Printf("Error%s\n", err)
		}
		l := res.Header.Get("X-Total-Count")

		m, err := strconv.Atoi(l)
		if err != nil {
			fmt.Println("Error discovering number of results")
		}

		h.Status = res.Status
		h.StatusCode = res.StatusCode
		h.Results = m

	} else {
		e = fmt.Errorf("the request returned a status %s. the full details are %s", res.Status, body)
	}

	return i, h, e
}

func (t RequestTransitive) CmdbCi() (CmdbCiResultsArray, Result, error) {

	var i CmdbCiResultsArray
	req := AssembleRequest(t, "cmdb_ci")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	defer CloseResponse(res)

	var e error
	var h Result

	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		err = json.Unmarshal(body, &i)
		if err != nil {
			fmt.Printf("Error %s\n", err)
		}
		l := res.Header.Get("X-Total-Count")

		m, err := strconv.Atoi(l)
		if err != nil {
			fmt.Println("Error discovering number of results")
		}

		h.Status = res.Status
		h.StatusCode = res.StatusCode
		h.Results = m

	} else {
		e = fmt.Errorf("the request returned a status %s. the full details are %s", res.Status, body)
	}

	return i, h, e
}

func (t RequestTransitive) CmdbCiServer() (CmdbCiServerResultsArray, Result, error) {

	var i CmdbCiServerResultsArray
	req := AssembleRequest(t, "cmdb_ci_server")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	defer CloseResponse(res)

	var e error
	var h Result

	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		err = json.Unmarshal(body, &i)
		if err != nil {
			fmt.Printf("Error %s\n", err)
		}
		l := res.Header.Get("X-Total-Count")

		m, err := strconv.Atoi(l)
		if err != nil {
			fmt.Println("Error discovering number of results")
		}

		h.Status = res.Status
		h.StatusCode = res.StatusCode
		h.Results = m

	} else {
		e = fmt.Errorf("the request returned a status %s. the full details are %s", res.Status, body)
	}

	return i, h, e
}

func (t RequestTransitive) CmdbCiIpRouter() (CmdbCiIpRouterResultsArray, Result, error) {

	var i CmdbCiIpRouterResultsArray
	req := AssembleRequest(t, "cmdb_ci_ip_router")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	defer CloseResponse(res)

	var e error
	var h Result

	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		err = json.Unmarshal(body, &i)
		if err != nil {
			fmt.Printf("Error%s\n", err)
		}
		l := res.Header.Get("X-Total-Count")

		m, err := strconv.Atoi(l)
		if err != nil {
			fmt.Println("Error discovering number of results")
		}

		h.Status = res.Status
		h.StatusCode = res.StatusCode
		h.Results = m

	} else {
		e = fmt.Errorf("the request returned a status %s. the full details are %s", res.Status, body)
	}

	return i, h, e
}

func (t RequestTransitive) CmdbCiIpSwitch() (CmdbCiIpSwitchResultsArray, Result, error) {

	var i CmdbCiIpSwitchResultsArray
	req := AssembleRequest(t, "cmdb_ci_ip_switch")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	defer CloseResponse(res)

	var e error
	var h Result

	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		err = json.Unmarshal(body, &i)
		if err != nil {
			fmt.Printf("Error%s\n", err)
		}
		l := res.Header.Get("X-Total-Count")

		m, err := strconv.Atoi(l)
		if err != nil {
			fmt.Println("Error discovering number of results")
		}

		h.Status = res.Status
		h.StatusCode = res.StatusCode
		h.Results = m

	} else {
		e = fmt.Errorf("the request returned a status %s. the full details are %s", res.Status, body)
	}

	return i, h, e
}

func (t RequestTransitive) CmdbCiFirewallNetwork() (CmdbCiFirewallNetworkResultsArray, Result, error) {

	var i CmdbCiFirewallNetworkResultsArray
	req := AssembleRequest(t, "cmdb_ci_firewall_network")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	defer CloseResponse(res)

	var e error
	var h Result

	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		err = json.Unmarshal(body, &i)
		if err != nil {
			fmt.Printf("Error%s\n", err)
		}
		l := res.Header.Get("X-Total-Count")

		m, err := strconv.Atoi(l)
		if err != nil {
			fmt.Println("Error discovering number of results")
		}

		h.Status = res.Status
		h.StatusCode = res.StatusCode
		h.Results = m

	} else {
		e = fmt.Errorf("the request returned a status %s. the full details are %s", res.Status, body)
	}

	return i, h, e
}

func (t RequestTransitive) CmdbCiNetgear() (CmdbCiNetgearResultsArray, Result, error) {

	var i CmdbCiNetgearResultsArray
	req := AssembleRequest(t, "cmdb_ci_firewall_network")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	defer CloseResponse(res)

	var e error
	var h Result

	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		err = json.Unmarshal(body, &i)
		if err != nil {
			fmt.Printf("Error%s\n", err)
		}
		l := res.Header.Get("X-Total-Count")

		m, err := strconv.Atoi(l)
		if err != nil {
			fmt.Println("Error discovering number of results")
		}

		h.Status = res.Status
		h.StatusCode = res.StatusCode
		h.Results = m

	} else {
		e = fmt.Errorf("the request returned a status %s. the full details are %s", res.Status, body)
	}

	return i, h, e
}

func (t RequestTransitive) CoreCompany() (CoreCompanyResultsArray, Result, error) {

	var i CoreCompanyResultsArray
	req := AssembleRequest(t, "core_company")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	defer CloseResponse(res)

	var e error
	var h Result

	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		err = json.Unmarshal(body, &i)
		if err != nil {
			fmt.Printf("Error%s\n", err)
		}
		l := res.Header.Get("X-Total-Count")

		m, err := strconv.Atoi(l)
		if err != nil {
			fmt.Println("Error discovering number of results")
		}

		h.Status = res.Status
		h.StatusCode = res.StatusCode
		h.Results = m

	} else {
		e = fmt.Errorf("the request returned a status %s. the full details are %s", res.Status, body)
	}

	return i, h, e
}

func (t RequestTransitive) SysUser() (SysUserResultsArray, Result, error) {

	var i SysUserResultsArray
	req := AssembleRequest(t, "sys_user")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	defer CloseResponse(res)

	var e error
	var h Result

	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		err = json.Unmarshal(body, &i)
		if err != nil {
			fmt.Printf("Error%s\n", err)
		}
		l := res.Header.Get("X-Total-Count")

		m, err := strconv.Atoi(l)
		if err != nil {
			fmt.Println("Error discovering number of results")
		}

		h.Status = res.Status
		h.StatusCode = res.StatusCode
		h.Results = m

	} else {
		e = fmt.Errorf("the request returned a status %s. the full details are %s", res.Status, body)
	}

	return i, h, e
}

func (t RequestTransitive) Custom(tableName string) (map[string]interface{} , Result, error) {

	var i map[string]interface{}
	req := AssembleRequest(t, tableName)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	defer CloseResponse(res)

	var e error
	var h Result

	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		err = json.Unmarshal(body, &i)
		if err != nil {
			fmt.Printf("Error%s\n", err)
		}
		l := res.Header.Get("X-Total-Count")

		m, err := strconv.Atoi(l)
		if err != nil {
			fmt.Println("Error discovering number of results")
		}

		h.Status = res.Status
		h.StatusCode = res.StatusCode
		h.Results = m

	} else {
		e = fmt.Errorf("the request returned a status %s. the full details are %s", res.Status, body)
	}

	return i, h, e
}