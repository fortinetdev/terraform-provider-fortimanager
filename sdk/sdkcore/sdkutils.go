package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand/v2"
	"regexp"
	"strings"
	"time"
)

////////////////////////////////////////////////////////////////////////

func createUpdate(c *FortiSDKClient, path string, method string, params *map[string]interface{}, move bool, wsParams map[string]string) (output map[string]interface{}, err error) {
	log.Printf("[INFO] Request infomation: %v:    %v\n", path, params)
	session := ""
	// Lock if needed
	if c.WorkspaceMode == "normal" {
		session, err = lockWorkspace(c, wsParams)
		if err != nil {
			return
		}
	}
	if c.Config.Auth.CleanSession && session == "" {
		session, err = c.loginSession()
	}
	// Send the request
	data := encodeData(c, path, method, session, params, move)
	_, result, err := sendRequest(c, data)
	// Unlock if needed
	if c.WorkspaceMode == "normal" {
		unlockErr := unlockWorkspace(c, wsParams, session)
		if unlockErr != nil {
			if err == nil {
				err = unlockErr
			} else {
				err = fmt.Errorf("%v, %v", err, unlockErr)
			}
		}
	}
	if c.Config.Auth.CleanSession {
		c.logoutSession(session)
	}

	if err == nil {
		output, err = decodeData(result)
	}

	return
}

func delete(c *FortiSDKClient, path string, method string, move bool, wsParams map[string]string) (err error) {
	session := ""
	// Lock if needed
	if c.WorkspaceMode == "normal" {
		session, err = lockWorkspace(c, wsParams)
		if err != nil {
			return
		}
	}
	if c.Config.Auth.CleanSession && session == "" {
		session, err = c.loginSession()
	}
	// Send the request
	data := encodeData(c, path, method, session, nil, move)
	_, _, err = sendRequest(c, data)
	// Unlock if needed
	if c.WorkspaceMode == "normal" {
		unlockErr := unlockWorkspace(c, wsParams, session)
		if unlockErr != nil {
			if err == nil {
				err = unlockErr
			} else {
				err = fmt.Errorf("%v, %v", err, unlockErr)
			}
		}
	}
	if c.Config.Auth.CleanSession {
		c.logoutSession(session)
	}

	return
}

func read(c *FortiSDKClient, path string, method string, move bool) (mapTmp map[string]interface{}, err error) {
	session := ""
	if c.Config.Auth.CleanSession {
		session, err = c.loginSession()
	}
	data := encodeData(c, path, method, session, nil, move)
	code, result, err := sendRequest(c, data)
	if c.Config.Auth.CleanSession {
		c.logoutSession(session)
	}

	if err != nil && code == -3 {
		err = nil
	}

	if err == nil {
		mapTmp, err = decodeData(result)
	}

	return
}

func readMove(c *FortiSDKClient, path string, method string, params *map[string]interface{}, move bool) (listTmp []interface{}, err error) {
	session := ""
	if c.Config.Auth.CleanSession {
		session, err = c.loginSession()
	}
	data := encodeData(c, path, method, session, params, move)
	_, result, err := sendRequest(c, data)
	if c.Config.Auth.CleanSession {
		c.logoutSession(session)
	}

	if err == nil {
		listTmp = decodeDataMove(result)
	}

	return
}

func sendRequest(c *FortiSDKClient, bodyData map[string]interface{}) (code int, result map[string]interface{}, err error) {
	locJSON, err := json.Marshal(bodyData)
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

	json.Unmarshal([]byte(string(body)), &result)

	code, err = fortiAPIErrorFormat(result, string(body))
	return
}

func lockWorkspace(c *FortiSDKClient, wsParams map[string]string) (session string, err error) {
	if lockMode, ok := wsParams["lockMode"]; ok {
		if lockMode == "skip_lock" {
			return "", nil
		}
	}
	adom := wsParams["adom"]
	path := "/dvmdb/" + adom + "/workspace/lock"
	maxLoop := 500
	for _ = range maxLoop {
		session, err = c.loginSession()
		if err != nil {
			err = fmt.Errorf("Login failed: %v", err)
			continue
		}
		data := encodeData(c, path, "exec", session, nil, false)
		code, _, errmsg := sendRequest(c, data)
		if code == 0 && err == nil {
			// Check whether locked session ID is current session ID
			curSessionID, _ := getCurSessionID(c, session)
			lockedSessionID, _ := getLockedSessionID(c, wsParams, session)
			if curSessionID == lockedSessionID {
				err = nil
				break
			} else {
				err = fmt.Errorf("Locked session not match current session")
			}
		} else {
			err = fmt.Errorf("Failed to lock the workspace: %v", errmsg)
		}
		c.logoutSession(session)
		session = ""
		waitTime := (0.3 + rand.Float64()*1.4) * 1500.0
		time.Sleep(time.Duration(int(waitTime)) * time.Millisecond)
		continue
	}
	if err != nil {
		err = fmt.Errorf("Error lock workspace: %v", err)
	}

	return session, err
}

func unlockWorkspace(c *FortiSDKClient, wsParams map[string]string, session string) (err error) {
	if lockMode, ok := wsParams["lockMode"]; ok {
		if lockMode == "skip_lock" || lockMode == "skip_unlock" {
			return nil
		}
	}
	adom := wsParams["adom"]
	// commit
	path := "/dvmdb/" + adom + "/workspace/commit"
	data := encodeData(c, path, "exec", session, nil, false)
	_, _, err = sendRequest(c, data)
	if err != nil {
		err = fmt.Errorf("Error commit workspace: %v", err)
		return
	}
	// unlock
	path = "/dvmdb/" + adom + "/workspace/unlock"
	data = encodeData(c, path, "exec", session, nil, false)
	_, _, err = sendRequest(c, data)
	if err != nil {
		err = fmt.Errorf("Error unlock workspace: %v", err)
		return
	}

	return err
}

func getCurSessionID(c *FortiSDKClient, session string) (sessionID int, err error) {
	path := "/pm/pkg/adom/not-existing"
	data := encodeData(c, path, "get", session, nil, false)
	_, result, err := sendRequest(c, data)
	sessionID = fortiIntValue(result["session"])
	return
}

func getLockedSessionID(c *FortiSDKClient, wsParams map[string]string, session string) (sessionID int, err error) {
	adom := wsParams["adom"]
	path := "/dvmdb/" + adom + "/workspace/lockinfo"
	data := encodeData(c, path, "get", session, nil, false)
	_, result, err := sendRequest(c, data)
	mapTmp, err := decodeData(result)
	if mapTmp == nil {
		sessionID = 0
		return
	}
	sessionID = fortiIntValue(mapTmp["lock_sid"])
	return
}

// //////////////////////////////////////////////////////////////////////
func encodeData(c *FortiSDKClient, path, method, session string, params *map[string]interface{}, move bool) map[string]interface{} {
	data := make(map[string]interface{})
	data["method"] = method
	data["params"] = make([]map[string]interface{}, 0)
	if method == "get" {
		data["verbose"] = 1
	}
	if session != "" {
		data["session"] = session
	} else if c.Session != "" {
		data["session"] = c.Session
	}

	paramItem := make(map[string]interface{})
	paramItem["url"] = path //"/cli/global/system/admin/setting"

	log.Printf("[INFO] Request URL: %v\n", paramItem["url"])

	if move == false {
		paramItem["data"] = params
	} else {
		for k, v := range *params {
			paramItem[k] = v
		}
	}

	v2 := make([]map[string]interface{}, 0)
	v2 = append(v2, paramItem)
	data["params"] = v2
	return data
}

func decodeData(result map[string]interface{}) (dataMap map[string]interface{}, err error) {
	v := result["result"]

	// fortiapi intercepted all the exceptions
	l := v.([]interface{})
	v2 := l[0].(map[string]interface{})
	if v2["data"] != nil {
		if dataList, ok := v2["data"].([]interface{}); ok {
			if len(dataList) == 0 {
				err = fmt.Errorf("API response is empty.")
				return
			}
			if dataMap, ok = dataList[0].(map[string]interface{}); ok {
				return
			} else {
				err = fmt.Errorf("The element of data list is not map for the API response. Please file an issue on provider's Github repository.")
			}
		} else if dataMap, ok = v2["data"].(map[string]interface{}); ok {
			return
		} else {
			err = fmt.Errorf("Could not identify the type of parameter 'data' of API response. Please file an issue on provider's Github repository.")
		}
	}

	return
}

func decodeDataMove(result map[string]interface{}) []interface{} {
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

func fortiAPIErrorFormat(result map[string]interface{}, body string) (code int, err error) {
	if result != nil {
		if result["result"] != nil {
			v := result["result"]

			l := v.([]interface{})
			if len(l) > 0 && l[0] != nil {
				v2 := l[0].(map[string]interface{})
				if v2["status"] != nil {
					v3 := v2["status"].(map[string]interface{})

					if v3["code"] != nil && v3["message"] != nil {
						code = fortiIntValue(v3["code"])
						message := fortiStringValue(v3["message"])

						if code == 0 {
							err = nil
							return
						}

						err = fmt.Errorf("\nerr %d: %v", code, message)
						return code, err
					}
				}
			}
		}
	}

	err = fmt.Errorf("\n%v", body)
	return code, err
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
			continue
		}
		curArgName := "{" + argName + "}"
		if curTfName, ok := argUrlMap[curArgName]; ok {
			rstPath = strings.ReplaceAll(rstPath, curTfName, argVal)
		} else {
			log.Printf("[Warning] argument map: %v do not contain argument: %v", argUrlMap, curArgName)
		}
	}
	rstPath = strings.ReplaceAll(rstPath, "//", "/")
	rstPath = strings.TrimRight(rstPath, "/")

	if strings.ContainsAny(rstPath, "[ | {") {
		return "", fmt.Errorf("path parameters error %v, %v", path, paradict)
	}

	return rstPath, nil
}
