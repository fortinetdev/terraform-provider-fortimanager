package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/terraform-providers/terraform-provider-fortimanager/sdk/auth"
	"github.com/terraform-providers/terraform-provider-fortimanager/sdk/config"
	"github.com/terraform-providers/terraform-provider-fortimanager/sdk/request"
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
	Token   string
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
}

func escapeURLStringDMScope(v string) string { // doesn't support "<>()"'#"
	return strings.Replace(strings.Replace(v, "/", "\\/", -1), " ", "/", -1)
}

// NewClient initializes a new global plugin client
// It returns the created client object
func NewClient(auth *auth.Auth, client *http.Client) (c *FortiSDKClient, err error) {
	c = &FortiSDKClient{}

	c.Session = ""

	c.Config.Auth = auth
	c.Config.HTTPCon = client
	c.Config.FwTarget = auth.Hostname

	// CleanSession only works for workspace mode set to 'disabled'
	if auth.CleanSession {
		workspaceMode, _ := c.GetWorkspaceMode()
		if workspaceMode != "disabled" {
			auth.CleanSession = false
		}
	}

	if !auth.CleanSession {
		err = c.login()
	}

	return
}

func (c *FortiSDKClient) login() (err error) {
	auth := c.Config.Auth

	if auth.Token != "" {
		c.Token = auth.Token
	} else if auth.Session != "" {
		c.Session = auth.Session
	} else if auth.FMGType == "forticloud" {
		err = c.loginFortiCloud()
	} else {
		c.Session, err = c.loginSession()
	}

	if err != nil {
		return
	}

	if auth.LogSession == true && c.Session != "" {
		saveSession(c.Session)
	}
	return
}

func saveSession(session string) {
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

func (c *FortiSDKClient) loginSession() (session string, err error) {
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
						session = fmt.Sprintf("%v", result["session"])
					}
				}

			}
		}
	} else {
		err = fmt.Errorf("Response body is nil or do not contain result.")
	}

	return
}

func (c *FortiSDKClient) loginFortiCloud() (err error) {
	log.Printf("Login through FortiCloud.")
	var forticloudToken string
	if c.Config.Auth.FMGCloudToken != "" {
		forticloudToken = c.Config.Auth.FMGCloudToken
	} else {
		forticloudToken, err = generateFortiCloudToken(c)
		if err != nil {
			return
		}
	}

	if forticloudToken == "" {
		err = fmt.Errorf("Could not get FortiManager Cloud Token.")
		return
	}

	data := make(map[string]interface{})
	data["access_token"] = forticloudToken

	locJSON, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
		return
	}

	bytes := bytes.NewBuffer(locJSON)

	req := c.NewRequest("POST", "/p/forticloud_jsonrpc_login/", nil, bytes)
	err = req.Send()
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("Could not get session %v", err)
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	req.HTTPResponse.Body.Close()
	if err != nil || body == nil {
		err = fmt.Errorf("Could not get response body %v", err)
		return
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	if result != nil && result["session"] != nil {
		c.Session = fmt.Sprintf("%v", result["session"])
		if c.Session == "" {
			err = fmt.Errorf("Session is empty.")
		}
	} else {
		err = fmt.Errorf("Response body is nil or do not contain session.")
	}

	return
}

func (c *FortiSDKClient) logoutSession(session string) (err error) {
	if session == "" {
		return
	}

	data := make(map[string]interface{})
	data["method"] = "exec"
	data["params"] = make([]map[string]interface{}, 0)
	data["session"] = session

	paramItem := make(map[string]interface{})
	paramItem["url"] = "/sys/logout"
	params := make([]map[string]interface{}, 0)
	params = append(params, paramItem)
	data["params"] = params

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
					if tmp == 0 || tmp == -11 {
						c.Session = ""
					} else {
						err = fmt.Errorf("Response code could not be recognized: %v", tmp)
					}
				}

			}
		}
	} else {
		err = fmt.Errorf("Response body is nil or do not contain result.")
	}

	return
}

func generateFortiCloudToken(c *FortiSDKClient) (forticloudToken string, err error) {
	log.Printf("Generate FortiCloud Token.")
	// Get access token
	data := make(map[string]interface{})
	data["username"] = c.Config.Auth.User
	data["password"] = c.Config.Auth.Passwd
	data["client_id"] = "FortiManager"
	data["grant_type"] = "password"

	locJSON, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
		return
	}

	bytes := bytes.NewBuffer(locJSON)

	client := &http.Client{
		Timeout: time.Second * 250,
	}
	req, err := http.NewRequest("POST", "https://customerapiauth.fortinet.com/api/v1/oauth/token/", bytes)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)

	if err != nil || resp == nil {
		err = fmt.Errorf("Could not generate FortiCloud token %v", err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil || body == nil {
		err = fmt.Errorf("Could not get FortiCloud response body %v", err)
		return
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)
	if result != nil && result["access_token"] != nil {
		forticloudToken = result["access_token"].(string)
	} else {
		err = fmt.Errorf("Could not generate FortiCloud token %v", err)
		return
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
	session := ""
	if c.Config.Auth.CleanSession {
		session, err = c.loginSession()
		data["session"] = session
	} else if c.Session != "" {
		data["session"] = c.Session
	}

	paramItem := make(map[string]interface{})
	paramItem["url"] = "/cli/global/system/status"

	v2 := make([]map[string]interface{}, 0)
	v2 = append(v2, paramItem)
	data["params"] = v2

	locJSON, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
		if c.Config.Auth.CleanSession {
			err = c.logoutSession(session)
		}
		return
	}
	bytes := bytes.NewBuffer(locJSON)

	req := c.NewRequest("POST", "/jsonrpc", nil, bytes)
	err = req.Send()

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	req.HTTPResponse.Body.Close()
	if c.Config.Auth.CleanSession {
		err = c.logoutSession(session)
	}
	if err != nil || body == nil {
		err = fmt.Errorf("Cannot get response body %v", err)
		return "", err
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)
	log.Printf("Get divice version response: %v\n", string(body))
	_, err = fortiAPIErrorFormat(result, string(body))
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

// GetWorkspaceMode gets the workspace mode of FMG
// It returns workspaceMode as string
func (c *FortiSDKClient) GetWorkspaceMode() (workspaceMode string, err error) {
	data := make(map[string]interface{})
	data["method"] = "get"
	data["params"] = make([]map[string]interface{}, 0)
	data["verbose"] = 1
	session := ""
	if c.Config.Auth.CleanSession {
		session, err = c.loginSession()
		data["session"] = session
	} else if c.Session != "" {
		data["session"] = c.Session
	}

	paramItem := make(map[string]interface{})
	paramItem["url"] = "/cli/global/system/global"

	v2 := make([]map[string]interface{}, 0)
	v2 = append(v2, paramItem)
	data["params"] = v2

	locJSON, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
		if c.Config.Auth.CleanSession {
			err = c.logoutSession(session)
		}
		return
	}
	bytes := bytes.NewBuffer(locJSON)

	req := c.NewRequest("POST", "/jsonrpc", nil, bytes)
	err = req.Send()

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	req.HTTPResponse.Body.Close()
	if c.Config.Auth.CleanSession {
		err = c.logoutSession(session)
	}
	if err != nil || body == nil {
		err = fmt.Errorf("Cannot get response body %v", err)
		return "", err
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)
	_, err = fortiAPIErrorFormat(result, string(body))
	if err != nil {
		return "", err
	}

	if workspaceMode, ok := result["result"].([]interface{})[0].(map[string]interface{})["data"].(map[string]interface{})["workspace-mode"].(string); ok {
		return workspaceMode, nil
	}

	err = fmt.Errorf("Cannot get the workspace mode in response %s", string(body))
	return "", err
}
