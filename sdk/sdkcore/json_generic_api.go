package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func (c *FortiSDKClient) JsonGenericAPI(data string) (output string, err error) {

	rev := map[string]interface{}{}
	json.Unmarshal([]byte(data), &rev)
	rev["verbose"] = 1
	session := ""
	if c.Config.Auth.CleanSession {
		session, err = c.loginSession()
		rev["session"] = session
	} else if c.Session != "" {
		rev["session"] = c.Session
	}

	locJSON, err := json.Marshal(rev)
	if err != nil {
		log.Fatal(err)
		return
	}

	bytes := bytes.NewBuffer(locJSON)

	req := c.NewRequest("POST", "/jsonrpc", nil, bytes)
	err = req.Send()
	if c.Config.Auth.CleanSession {
		err = c.logoutSession(session)
	}
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request %v", err)
		return
	}

	defer req.HTTPResponse.Body.Close()
	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %v", err)
		return
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)
	_, err = fortiAPIErrorFormat(result, string(body))

	if err == nil {
		output = string(body)
		log.Printf("Successful\n")
	}

	return
}
