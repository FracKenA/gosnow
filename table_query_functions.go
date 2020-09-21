package gosnow

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (requestTransitive RequestTransitive) Incident() (map[string]interface{}, Result, []error) {

	var returnedPayload map[string]interface{}
	request := AssembleRequest(requestTransitive, "incident")
	
	requestResult, err := http.DefaultClient.Do(request)
	if err != nil {
		requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
	}

	body, err := ioutil.ReadAll(requestResult.Body)
	if err != nil {
		requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
	}

	err = requestResult.Body.Close()
	if err != nil {
		requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
	}
	var result Result

	if requestResult.StatusCode >= 200 && requestResult.StatusCode <= 299 {
		err = json.Unmarshal(body, &returnedPayload)
		if err != nil {
			requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
		}

		resultCount, err := strconv.Atoi(requestResult.Header.Get("X-Total-Count"))
		if err != nil {
			requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
		}

		result.Status = requestResult.Status
		result.StatusCode = requestResult.StatusCode
		result.ResultCount = resultCount

	} else {
		requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: the request returned a status %s. the full details are %s", requestResult.Status, body))
	}

	return returnedPayload, result, requestTransitive.err
}

func (requestTransitive RequestTransitive) Case() (map[string]interface{}, Result, []error) {

	var returnedPayload map[string]interface{}
	request := AssembleRequest(requestTransitive, "sn_customerservice_case")

	requestResult, err := http.DefaultClient.Do(request)
	if err != nil {
		requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
	}

	body, err := ioutil.ReadAll(requestResult.Body)
	if err != nil {
		requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
	}

	err = requestResult.Body.Close()
	if err != nil {
		requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
	}
	var result Result

	if requestResult.StatusCode >= 200 && requestResult.StatusCode <= 299 {
		err = json.Unmarshal(body, &returnedPayload)
		if err != nil {
			requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
		}

		resultCount, err := strconv.Atoi(requestResult.Header.Get("X-Total-Count"))
		if err != nil {
			requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
		}

		result.Status = requestResult.Status
		result.StatusCode = requestResult.StatusCode
		result.ResultCount = resultCount

	} else {
		requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: the request returned a status %s. the full details are %s", requestResult.Status, body))
	}

	return returnedPayload, result, requestTransitive.err
}

func (requestTransitive RequestTransitive) CmdbCi() (map[string]interface{}, Result, []error) {

	var returnedPayload map[string]interface{}
	request := AssembleRequest(requestTransitive, "cmdb_ci")

	requestResult, err := http.DefaultClient.Do(request)
	if err != nil {
		requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
	}

	body, err := ioutil.ReadAll(requestResult.Body)
	if err != nil {
		requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
	}

	err = requestResult.Body.Close()
	if err != nil {
		requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
	}
	var result Result

	if requestResult.StatusCode >= 200 && requestResult.StatusCode <= 299 {
		err = json.Unmarshal(body, &returnedPayload)
		if err != nil {
			requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
		}

		resultCount, err := strconv.Atoi(requestResult.Header.Get("X-Total-Count"))
		if err != nil {
			requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
		}

		result.Status = requestResult.Status
		result.StatusCode = requestResult.StatusCode
		result.ResultCount = resultCount

	} else {
		requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: the request returned a status %s. the full details are %s", requestResult.Status, body))
	}

	return returnedPayload, result, requestTransitive.err
}

func (requestTransitive RequestTransitive) CmdbCiServer() (map[string]interface{}, Result, []error) {
	
	var returnedPayload map[string]interface{}
	request := AssembleRequest(requestTransitive, "cmdb_ci_server")

	requestResult, err := http.DefaultClient.Do(request)
	if err != nil {
		requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
	}

	body, err := ioutil.ReadAll(requestResult.Body)
	if err != nil {
		requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
	}

	err = requestResult.Body.Close()
	if err != nil {
		requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
	}
	var result Result

	if requestResult.StatusCode >= 200 && requestResult.StatusCode <= 299 {
		err = json.Unmarshal(body, &returnedPayload)
		if err != nil {
			requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
		}

		resultCount, err := strconv.Atoi(requestResult.Header.Get("X-Total-Count"))
		if err != nil {
			requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
		}

		result.Status = requestResult.Status
		result.StatusCode = requestResult.StatusCode
		result.ResultCount = resultCount

	} else {
		requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: the request returned a status %s. the full details are %s", requestResult.Status, body))
	}

	return returnedPayload, result, requestTransitive.err
}

func (requestTransitive RequestTransitive) CmdbCiIpRouter() (map[string]interface{}, Result, []error) {

	var returnedPayload map[string]interface{}
	request := AssembleRequest(requestTransitive, "cmdb_ci_ip_router")

	requestResult, err := http.DefaultClient.Do(request)
	if err != nil {
		requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
	}

	body, err := ioutil.ReadAll(requestResult.Body)
	if err != nil {
		requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
	}

	err = requestResult.Body.Close()
	if err != nil {
		requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
	}
	var result Result

	if requestResult.StatusCode >= 200 && requestResult.StatusCode <= 299 {
		err = json.Unmarshal(body, &returnedPayload)
		if err != nil {
			requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
		}

		resultCount, err := strconv.Atoi(requestResult.Header.Get("X-Total-Count"))
		if err != nil {
			requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
		}

		result.Status = requestResult.Status
		result.StatusCode = requestResult.StatusCode
		result.ResultCount = resultCount

	} else {
		requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: the request returned a status %s. the full details are %s", requestResult.Status, body))
	}

	return returnedPayload, result, requestTransitive.err
}

func (requestTransitive RequestTransitive) CmdbCiIpSwitch() (map[string]interface{}, Result, []error) {
	var returnedPayload map[string]interface{}
	request := AssembleRequest(requestTransitive, "cmdb_ci_ip_switch")

	requestResult, err := http.DefaultClient.Do(request)
	if err != nil {
		requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
	}

	body, err := ioutil.ReadAll(requestResult.Body)
	if err != nil {
		requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
	}

	err = requestResult.Body.Close()
	if err != nil {
		requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
	}
	var result Result

	if requestResult.StatusCode >= 200 && requestResult.StatusCode <= 299 {
		err = json.Unmarshal(body, &returnedPayload)
		if err != nil {
			requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
		}

		resultCount, err := strconv.Atoi(requestResult.Header.Get("X-Total-Count"))
		if err != nil {
			requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
		}

		result.Status = requestResult.Status
		result.StatusCode = requestResult.StatusCode
		result.ResultCount = resultCount

	} else {
		requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: the request returned a status %s. the full details are %s", requestResult.Status, body))
	}

	return returnedPayload, result, requestTransitive.err
}

func (requestTransitive RequestTransitive) CmdbCiFirewallNetwork() (map[string]interface{}, Result, []error) {
	var returnedPayload map[string]interface{}
	request := AssembleRequest(requestTransitive, "cmdb_ci_firewall_network")

	requestResult, err := http.DefaultClient.Do(request)
	if err != nil {
		requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
	}

	body, err := ioutil.ReadAll(requestResult.Body)
	if err != nil {
		requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
	}

	err = requestResult.Body.Close()
	if err != nil {
		requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
	}
	var result Result

	if requestResult.StatusCode >= 200 && requestResult.StatusCode <= 299 {
		err = json.Unmarshal(body, &returnedPayload)
		if err != nil {
			requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
		}

		resultCount, err := strconv.Atoi(requestResult.Header.Get("X-Total-Count"))
		if err != nil {
			requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
		}

		result.Status = requestResult.Status
		result.StatusCode = requestResult.StatusCode
		result.ResultCount = resultCount

	} else {
		requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: the request returned a status %s. the full details are %s", requestResult.Status, body))
	}

	return returnedPayload, result, requestTransitive.err
}

func (requestTransitive RequestTransitive) CmdbCiNetgear() (map[string]interface{}, Result, []error) {
	var returnedPayload map[string]interface{}
	request := AssembleRequest(requestTransitive, "cmdb_ci_firewall_network")

	requestResult, err := http.DefaultClient.Do(request)
	if err != nil {
		requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
	}

	body, err := ioutil.ReadAll(requestResult.Body)
	if err != nil {
		requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
	}

	err = requestResult.Body.Close()
	if err != nil {
		requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
	}
	var result Result

	if requestResult.StatusCode >= 200 && requestResult.StatusCode <= 299 {
		err = json.Unmarshal(body, &returnedPayload)
		if err != nil {
			requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
		}

		resultCount, err := strconv.Atoi(requestResult.Header.Get("X-Total-Count"))
		if err != nil {
			requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
		}

		result.Status = requestResult.Status
		result.StatusCode = requestResult.StatusCode
		result.ResultCount = resultCount

	} else {
		requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: the request returned a status %s. the full details are %s", requestResult.Status, body))
	}

	return returnedPayload, result, requestTransitive.err
}

func (requestTransitive RequestTransitive) CoreCompany() (map[string]interface{}, Result, []error) {
	var returnedPayload map[string]interface{}
	request := AssembleRequest(requestTransitive, "core_company")

	requestResult, err := http.DefaultClient.Do(request)
	if err != nil {
		requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
	}

	body, err := ioutil.ReadAll(requestResult.Body)
	if err != nil {
		requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
	}

	err = requestResult.Body.Close()
	if err != nil {
		requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
	}
	var result Result

	if requestResult.StatusCode >= 200 && requestResult.StatusCode <= 299 {
		err = json.Unmarshal(body, &returnedPayload)
		if err != nil {
			requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
		}

		resultCount, err := strconv.Atoi(requestResult.Header.Get("X-Total-Count"))
		if err != nil {
			requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
		}

		result.Status = requestResult.Status
		result.StatusCode = requestResult.StatusCode
		result.ResultCount = resultCount

	} else {
		requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: the request returned a status %s. the full details are %s", requestResult.Status, body))
	}

	return returnedPayload, result, requestTransitive.err
}

func (requestTransitive RequestTransitive) SysUser() (map[string]interface{}, Result, []error) {
	var returnedPayload map[string]interface{}
	request := AssembleRequest(requestTransitive, "sys_user")

	requestResult, err := http.DefaultClient.Do(request)
	if err != nil {
		requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
	}

	body, err := ioutil.ReadAll(requestResult.Body)
	if err != nil {
		requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
	}

	err = requestResult.Body.Close()
	if err != nil {
		requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
	}

	var result Result

	if requestResult.StatusCode >= 200 && requestResult.StatusCode <= 299 {
		err = json.Unmarshal(body, &returnedPayload)
		if err != nil {
			requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
		}

		resultCount, err := strconv.Atoi(requestResult.Header.Get("X-Total-Count"))
		if err != nil {
			requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: %s\n", err))
		}

		result.Status = requestResult.Status
		result.StatusCode = requestResult.StatusCode
		result.ResultCount = resultCount

	} else {
		requestTransitive.err = append(requestTransitive.err, fmt.Errorf("warning: the request returned a status %s. the full details are %s", requestResult.Status, body))
	}

	return returnedPayload, result, requestTransitive.err
}
