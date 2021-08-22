package servicenow

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (t RequestTransitive) IncidentDetail() (IncidentResultDetail, error) {

	var i IncidentResultDetail
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

	// fmt.Println(string(body))

	var e error
	var h Result

	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		err = json.Unmarshal(body, &i)
		if err != nil {
			fmt.Printf("Error%s\n", err)
		}
		h.Status = res.Status
		h.StatusCode = res.StatusCode

	} else {
		e = fmt.Errorf("the request returned a status %s. the full details are %s", res.Status, body)
	}

	return i, e
}

func (t RequestTransitive) CaseDetail() (CaseResultDetail, error) {

	var i CaseResultDetail
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
		h.Status = res.Status
		h.StatusCode = res.StatusCode

	} else {
		e = fmt.Errorf("the request returned a status %s. the full details are %s", res.Status, body)
	}

	return i, e
}

func (t RequestTransitive) CmdbCiDetail() (CmdbCiResultDetail, error) {

	var i CmdbCiResultDetail
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
			e = fmt.Errorf("Error %s\n", err)
		}
		h.Status = res.Status
		h.StatusCode = res.StatusCode

	} else {
		e = fmt.Errorf("the request returned a status %s. the full details are %s", res.Status, body)
	}

	return i, e
}

func (t RequestTransitive) CmdbCiServerDetail() (CmdbCiServerResultDetail, string) {

	var i CmdbCiServerResultDetail
	req := AssembleRequest(t, "cmdb_ci_server")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	fmt.Println(string(body))

	defer CloseResponse(res)

	var e string
	var h Result

	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		err = json.Unmarshal(body, &i)
		if err != nil {
			fmt.Printf("Error %s\n", err)
		}
		h.Status = res.Status
		h.StatusCode = res.StatusCode

	} else {
		e = fmt.Sprintf("the request returned a status %s. the full details are %s", res.Status, body)
	}

	return i, e
}

func (t RequestTransitive) CmdbCiIpRouterDetail() (CmdbCiIpRouterResultDetail, string) {

	var i CmdbCiIpRouterResultDetail
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

	var e string
	var h Result

	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		err = json.Unmarshal(body, &i)
		if err != nil {
			fmt.Printf("Error%s\n", err)
		}
		h.Status = res.Status
		h.StatusCode = res.StatusCode

	} else {
		e = fmt.Sprintf("the request returned a status %s. the full details are %s", res.Status, body)
	}

	return i, e
}

func (t RequestTransitive) CmdbCiFirewallNetworkDetail() (CmdbCiFirewallNetworkResultDetail, string) {

	var i CmdbCiFirewallNetworkResultDetail
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

	var e string
	var h Result

	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		err = json.Unmarshal(body, &i)
		if err != nil {
			fmt.Printf("Error%s\n", err)
		}
		h.Status = res.Status
		h.StatusCode = res.StatusCode

	} else {
		e = fmt.Sprintf("the request returned a status %s. the full details are %s", res.Status, body)
	}

	return i, e
}

func (t RequestTransitive) CmdbCiNetgearDetail() (CmdbCiNetgearResultDetail, string) {

	var i CmdbCiNetgearResultDetail
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

	var e string
	var h Result

	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		err = json.Unmarshal(body, &i)
		if err != nil {
			fmt.Printf("Error%s\n", err)
		}
		h.Status = res.Status
		h.StatusCode = res.StatusCode

	} else {
		e = fmt.Sprintf("the request returned a status %s. the full details are %s", res.Status, body)
	}

	return i, e
}

func (t RequestTransitive) CoreCompanyDetail() (CoreCompanyResultDetail, error) {

	var i CoreCompanyResultDetail

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
		h.Status = res.Status
		h.StatusCode = res.StatusCode

	} else {
		e = fmt.Errorf("the request returned a status %s. the full details are %s", res.Status, body)
	}

	return i, e
}

func (t RequestTransitive) SysUserDetail() (SysUserResultDetail, string) {

	var i SysUserResultDetail
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

	var e string
	var h Result

	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		err = json.Unmarshal(body, &i)
		if err != nil {
			fmt.Printf("Error%s\n", err)
		}
		h.Status = res.Status
		h.StatusCode = res.StatusCode

	} else {
		e = fmt.Sprintf("the request returned a status %s. the full details are %s", res.Status, body)
	}

	return i, e
}

func (t RequestTransitive) UApplicationPackDetail() (UApplicationPackResultDetail, string) {

	var i UApplicationPackResultDetail
	req := AssembleRequest(t, "u_application_pack")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error %s\n", err)
	}

	defer CloseResponse(res)

	var e string
	var h Result

	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		err = json.Unmarshal(body, &i)
		if err != nil {
			fmt.Printf("Error%s\n", err)
		}
		h.Status = res.Status
		h.StatusCode = res.StatusCode

	} else {
		e = fmt.Sprintf("the request returned a status %s. the full details are %s", res.Status, body)
	}

	return i, e
}
