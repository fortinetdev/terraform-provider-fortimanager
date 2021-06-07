
package forticlient

import (
	"fmt"
	"encoding/json"
	"bytes"
	"io/ioutil"
	"log"
)

func (c *FortiSDKClient) JsonGenericAPI(data string) (output string, err error) {

	rev := map[string]interface{}{}
	json.Unmarshal([]byte(data), &rev)
	rev["session"] = c.Session

	locJSON, err := json.Marshal(rev)
	if err != nil {
		log.Fatal(err)
		return
	}

	bytes := bytes.NewBuffer(locJSON)

	req := c.NewRequest("POST", "/jsonrpc", nil, bytes)
	err = req.Send()
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
	err = fortiAPIErrorFormat(result, string(body))

	if err == nil {
		output = string(body)
		log.Printf("Successful\n")
	}

	return
}
