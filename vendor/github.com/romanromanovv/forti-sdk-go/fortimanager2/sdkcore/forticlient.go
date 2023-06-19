package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"net/http"
	"net/url"
	"strings"
	"regexp"

	"github.com/fortinetdev/forti-sdk-go/fortimanager2/auth"
	"github.com/fortinetdev/forti-sdk-go/fortimanager2/config"
	"github.com/fortinetdev/forti-sdk-go/fortimanager2/request"
	// "strconv"
)

// MultValue describes the nested structure in the results
type MultValue struct {
	Name string `json:"name"`
}

// MultValues describes the nested structure in the results
type MultValues []MultValue

// FortiSDKClient describes the global FMG plugin client instance
type FortiSDKClient struct {
	Config  config.Config
	Session string
	Retries int
}

// ExtractString extracts strings from result and put them into a string array,
// and return the string array
func ExtractString(members []MultValue) []string {
	vs := make([]string, 0, len(members))
	for _, v := range members {
		c := v.Name
		vs = append(vs, c)
	}
	return vs
}

// EscapeURLString escapes the string so it can be safely placed inside a URL query
func EscapeURLString(v string) string { // doesn't support "<>()"'#"
	return strings.Replace(url.QueryEscape(v), "+", "%20", -1)
}

func escapeURLString(v string) string { // doesn't support "<>()"'#"
	return strings.Replace(v, "/", "\\/", -1)
	// return strings.Replace(url.QueryEscape(v), "+", "%20", -1)
}

// NewClient initializes a new global plugin client
// It returns the created client object
func NewClient(auth *auth.Auth, client *http.Client) *FortiSDKClient {
	c := &FortiSDKClient{}

	c.Session = ""

	c.Config.Auth = auth
	c.Config.HTTPCon = client
	c.Config.FwTarget = auth.Hostname

	if auth.Session != "" {
		c.Session = auth.Session
	} else {
		login(auth, c)
	}

	saveSession(auth, c.Session)

	return c
}

func saveSession(auth *auth.Auth, session string) {
	if auth.LogSession == true {
		f, err := os.Create("presession.txt")
		if err != nil {
			return
		}
		_, err = f.WriteString(session + "\n")
		if err != nil {
			f.Close()
			return
		}

		err = f.Close()
		if err != nil {
			return
		}
	}
}

func login(auth *auth.Auth, c *FortiSDKClient) {
	data := make(map[string]interface{})
	data["method"] = "exec"
	data["params"] = make([]map[string]interface{}, 0)

	paramItem := make(map[string]interface{})
	paramItem["url"] = "sys/login/user"
	v1 := make([]map[string]interface{}, 0)

	paramItemData := make(map[string]interface{})
	paramItemData["user"] = c.Config.Auth.User
	paramItemData["passwd"] = c.Config.Auth.Passwd
	v1 = append(v1, paramItemData)
	paramItem["data"] = v1

	v2 := make([]map[string]interface{}, 0)
	v2 = append(v2, paramItem)
	data["params"] = v2


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

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	req.HTTPResponse.Body.Close()
	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %v", err)
		return
	}


	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	if result != nil && result["result"] != nil {
		v := result["result"].([]interface{})
		if v[0] != nil {
			v1 := v[0].(map[string]interface{})
			if v1["status"] != nil {
				v2 := v1["status"].(map[string]interface{})

				if vv, ok := v2["code"].(float64); ok {
					tmp := int(vv)
					if tmp == 0 {
						c.Session = fmt.Sprintf("%v", result["session"])
					}
				}

			}
		}
	}

	return
}

// NewRequest creates the request to FMG for the client
// and return it to the client
func (c *FortiSDKClient) NewRequest(method string, path string, params interface{}, data *bytes.Buffer) *request.Request {
	return request.New(c.Config, method, path, params, data)
}

// GetDeviceVersion gets the version of FMG
// It returns version as string
func (c *FortiSDKClient) GetDeviceVersion() (version string, err error) {
	data := make(map[string]interface{})
	data["method"] = "get"
	data["params"] = make([]map[string]interface{}, 0)
	data["verbose"] = 1
	data["session"] = c.Session

	paramItem := make(map[string]interface{})
	paramItem["url"] = "/cli/global/system/status"

	v2 := make([]map[string]interface{}, 0)
	v2 = append(v2, paramItem)
	data["params"] = v2

	locJSON, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
		return
	}
	bytes := bytes.NewBuffer(locJSON)

	req := c.NewRequest("POST", "/jsonrpc", nil, bytes)
	err = req.Send()

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	req.HTTPResponse.Body.Close()
	if err != nil || body == nil {
		err = fmt.Errorf("Cannot get response body %v", err)
		return "", err
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)
	log.Printf("Get divice version response: %v\n", string(body))
	err = fortiAPIErrorFormat(result, string(body))
	if err != nil {
		return "", err
	}
	
	if version, ok := result["result"].([]interface{})[0].(map[string]interface{})["data"].(map[string]interface{})["Version"].(string); ok {
		regexp, err := regexp.Compile(`v([\d.]+)`)
		match := regexp.FindStringSubmatch(version)
		if len(match) > 1 {
			return match[1], err
		}
		err = fmt.Errorf("Cannot get the version number in response %s", string(body))
	}

	err = fmt.Errorf("Cannot get the version in response %s", string(body))
	return "", err
}
