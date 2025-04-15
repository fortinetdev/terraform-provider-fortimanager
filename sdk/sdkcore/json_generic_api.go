package forticlient

import (
	"encoding/json"
	"fmt"
	"log"
)

func (c *FortiSDKClient) JsonGenericAPI(data string, wsParams map[string]string) (output string, err error) {

	rev := map[string]interface{}{}
	json.Unmarshal([]byte(data), &rev)
	rev["verbose"] = 1
	session := ""
	// Lock if needed
	if c.WorkspaceMode == "normal" {
		session, err = lockWorkspace(c, wsParams)
		if err != nil {
			return
		}
	} else if c.Config.Auth.CleanSession {
		session, err = c.loginSession()
	} else if c.Session != "" {
		session = c.Session
	}
	rev["session"] = session

	_, result, err := sendRequest(c, rev)

	// Unlock if needed
	if c.WorkspaceMode == "normal" {
		err = unlockWorkspace(c, wsParams, session)
		if err != nil {
			return
		}
	}
	if c.Config.Auth.CleanSession {
		c.logoutSession(session)
	}

	if err == nil {
		output = string(fmt.Sprintf("%v", result))
		log.Printf("Successful\n")
	}

	return
}
