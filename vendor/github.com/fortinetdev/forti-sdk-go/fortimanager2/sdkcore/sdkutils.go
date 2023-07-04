package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"regexp"
)

////////////////////////////////////////////////////////////////////////

func createUpdate(c *FortiSDKClient, path string, method string, params *map[string]interface{}, move bool) (output map[string]interface{}, err error) {
	log.Printf("[INFO] Request infomation: %v:    %v\n", path, params)

	data := encodeData(c, path, method, params, move)
	locJSON, err := json.Marshal(data)
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
		output = decodeData(result)
		log.Printf("Successful\n")
	}

	return
}

func delete(c *FortiSDKClient, path string, method string, move bool) (err error) {
	data := encodeData(c, path, method, nil, move)
	locJSON, err := json.Marshal(data)
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
	log.Printf("FMG response: %s", string(body))


	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	err = fortiAPIErrorFormat(result, string(body))

	return
}

func read(c *FortiSDKClient, path string, method string, move bool) (mapTmp map[string]interface{}, err error) {
	data := encodeData(c, path, method, nil, move)
	locJSON, err := json.Marshal(data)
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
	log.Printf("FMG reading response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	err = fortiAPIErrorFormat(result, string(body))

	if err == nil {
		mapTmp = decodeData(result)
	}

	return
}

func readMove(c *FortiSDKClient, path string, method string, params *map[string]interface{}, move bool) (listTmp []interface{}, err error) {
	data := encodeData(c, path, method, params, move)
	locJSON, err := json.Marshal(data)
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
	log.Printf("FMG reading response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	err = fortiAPIErrorFormat(result, string(body))

	if err == nil {
		listTmp = decodeDataMove(result)
	}

	return
}

////////////////////////////////////////////////////////////////////////
func encodeData(c *FortiSDKClient, path, method string, params *map[string]interface{}, move bool) map[string]interface{} {
	data := make(map[string]interface{})
	data["method"] = method
	data["params"] = make([]map[string]interface{}, 0)
	data["verbose"] = 1
	data["session"] = c.Session

	paramItem := make(map[string]interface{})
	paramItem["url"] = path //"/cli/global/system/admin/setting"

	log.Printf("[INFO] Request URL: %v\n", paramItem["url"])

	if move == false {
		paramItem["data"] = params
	} else {
		for k,v := range *params {
			paramItem[k] = v
		}
	}

	v2 := make([]map[string]interface{}, 0)
	v2 = append(v2, paramItem)
	data["params"] = v2
	return data
}


func decodeData(result map[string]interface{}) (map[string]interface{}) {
	v := result["result"]

	// fortiapi intercepted all the exceptions
	l := v.([]interface{})
	v2 := l[0].(map[string]interface{})
	if v2["data"] != nil {
		mapTmp := v2["data"].(map[string]interface{})
		return mapTmp
	}

	return nil
}

func decodeDataMove(result map[string]interface{}) ([]interface{}) {
	v := result["result"]

	// fortiapi intercepted all the exceptions
	l := v.([]interface{})
	v2 := l[0].(map[string]interface{})

	if v2["data"] != nil {
		v3 := v2["data"]

		if v, ok := v3.([]interface{}); ok {
			return v
		}
	}

	return nil
}

func fortiAPIErrorFormat(result map[string]interface{}, body string) (err error) {
	if result != nil {
		if result["result"] != nil {
			v := result["result"]

			l := v.([]interface{})
			if len(l) > 0 || l[0] != nil {
				v2 := l[0].(map[string]interface{})
				if v2["status"] != nil {
					v3 := v2["status"].(map[string]interface{})

					if v3["code"] != nil && v3["message"] != nil {
						code := fortiIntValue(v3["code"])
						message := fortiStringValue(v3["message"])

						if code == 0 {
							err = nil
							return
						}

						err = fmt.Errorf("\nerr %d: %v", code, message)
						return
					}
				}
			}
		}
	}

	err = fmt.Errorf("\n%v", body)
	return
}
////////////////////////////////////////////////////////////////////////

func fortiAPIHttpStatus404Checking(result map[string]interface{}) (b404 bool) {
	b404 = false

	if result != nil {
		if result["http_status"] != nil && result["http_status"] == 404.0 {
			b404 = true
			return
		}
	}

	return
}

func fortiStringValue(t interface{}) string {
	if v, ok := t.(string); ok {
		return v
	} else {
		return ""
	}
}

func fortiIntValue(t interface{}) int {
	if v, ok := t.(float64); ok {
		return int(v)
	} else {
		return 0
	}
}

func replaceParaWithValue(path string, paradict map[string]string) (string, error) {
	if !strings.ContainsAny(path, "[ | {") {
		return path, nil
	}
	if paradict == nil {
		return "", fmt.Errorf("Missing path parameters error %v, %v", path, paradict)
	}
	rstPath := path

	if adomv, ok := paradict["adom"]; ok && adomv != "" {
		rstPath = strings.ReplaceAll(rstPath, "[*]", adomv)
	}

	re := regexp.MustCompile(`{.*?}`)

	argUrl := re.FindAllString(path, -1)
	argUrlMap := make(map[string]string)
	for _, argName := range argUrl {
		re = regexp.MustCompile(`[.()+/ ]`)
		curTfName := re.ReplaceAllString(argName, "")
		curTfName = strings.ReplaceAll(curTfName, "-", "_")
		curTfName = strings.ToLower(curTfName)
		argUrlMap[curTfName] = argName
	}

	for argName, argVal := range paradict {
		if argName == "adom" {
			argVal = strings.Replace(argVal, "adom/", "", 1)
		}
		curArgName := "{" + argName + "}"
		if curTfName, ok := argUrlMap[curArgName]; ok {
			rstPath = strings.ReplaceAll(rstPath, curTfName, argVal)
		} else {
			log.Printf("[Warning] argument map: %v do not contain argument: %v", argUrlMap, curArgName)
		}
	}

	if strings.ContainsAny(rstPath, "[ | {") {
		return "", fmt.Errorf("path parameters error %v, %v", path, paradict)
	}

	return rstPath, nil
}
