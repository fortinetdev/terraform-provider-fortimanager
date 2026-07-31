// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: SDK for FortiManager Provider

package forticlient

import (
	"fmt"
	"time"
)

// CreateDevice API operation for FortiManager add a new device.
// Returns the index value of the Device and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dvm - cmd add device chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateDevice(params *map[string]interface{}, paradict, wsParams map[string]string) (output map[string]interface{}, err error) {
	path := "/dvm/cmd/add/device"
	path, err = replaceParaWithValue(path, paradict)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	requestInput := &requestInput{}

	requestInput.fortiSDKClient = c
	requestInput.method = "exec"
	requestInput.path = path
	requestInput.bodyParams = params
	requestInput.wsParams = wsParams
	requestInput.bMove = false
	requestInput.waitTime = 2

	output, err = createUpdate(requestInput)
	return
}

// UpdateDevice API operation for FortiManager updates the specified device.
// Returns the index value of the Device and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dvm - cmd update device chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateDevice(params *map[string]interface{}, mkey string, paradict, wsParams map[string]string) (output map[string]interface{}, err error) {
	path := "/dvm/cmd/update/device"
	path, err = replaceParaWithValue(path, paradict)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	requestInput := &requestInput{}

	requestInput.fortiSDKClient = c
	requestInput.method = "exec"
	requestInput.path = path
	requestInput.bodyParams = params
	requestInput.wsParams = wsParams
	requestInput.bMove = false

	output, err = createUpdate(requestInput)
	return
}

// DeleteDevice API operation for FortiManager deletes the specified device.
// Returns error for service API and SDK errors.
// See the dvm - cmd del device chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteDevice(params *map[string]interface{}, mkey string, paradict, wsParams map[string]string) (output map[string]interface{}, err error) {
	path := "/dvm/cmd/del/device"
	path, err = replaceParaWithValue(path, paradict)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	requestInput := &requestInput{}

	requestInput.fortiSDKClient = c
	requestInput.method = "exec"
	requestInput.path = path
	requestInput.bodyParams = params
	requestInput.wsParams = wsParams
	requestInput.bMove = false

	output, err = createUpdate(requestInput)
	return
}

// ReadDevice API operation for FortiManager gets the Device
// with the specified index value.
// Returns the requested Device value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dvmdb - device chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadDevice(mkey string, paradict map[string]string) (mapTmp map[string]interface{}, err error) {
	path := "/dvmdb/device"
	path, err = replaceParaWithValue(path, paradict)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	requestInput := &requestInput{}

	requestInput.fortiSDKClient = c
	requestInput.method = "get"
	requestInput.path = path
	requestInput.bMove = false

	mapTmp, err = read(requestInput)
	return
}

// ReadTask API operation for FortiManager gets the task info
// with the specified index value.
// Returns the requested Device value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the task - task chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadTask(mkey string, paradict map[string]string) (mapTmp map[string]interface{}, err error) {
	path := "/task/task"
	path, err = replaceParaWithValue(path, paradict)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	requestInput := &requestInput{}

	requestInput.fortiSDKClient = c
	requestInput.method = "get"
	requestInput.path = path
	requestInput.bMove = false

	mapTmp, err = read(requestInput)
	return
}

func (c *FortiSDKClient) WaitTask(taskID string) (err error) {
	if taskID == "" {
		return fmt.Errorf("Task ID is empty.")
	}

	const timeoutSeconds = 300 // 5 min

	var percent float64
	noChangeCount := 0

	for i := 0; i < timeoutSeconds; i++ {
		taskInfo, err := c.ReadTask(taskID, nil)
		if err != nil {
			return fmt.Errorf("read task %s: %w", taskID, err)
		}

		curPercent, ok := taskInfo["percent"].(float64)
		if !ok {
			return fmt.Errorf("task %s: missing or invalid 'percent' field", taskID)
		}

		curState, ok := taskInfo["state"].(string)
		if !ok {
			return fmt.Errorf("task %s: missing or invalid 'state' field", taskID)
		}

		if curState == "error" {
			return fmt.Errorf("Task %s failed.", taskID)
		}

		if curState == "done" && curPercent == 100 {
			return nil
		}

		if curPercent == percent {
			noChangeCount++
		} else {
			percent = curPercent
			noChangeCount = 0
		}

		// Task has not progressed for 60 seconds
		if noChangeCount >= 60 {
			return fmt.Errorf("task %s stuck at %.0f%%", taskID, percent)
		}

		time.Sleep(time.Second)
	}
	return fmt.Errorf("Timeout, task did not finish in 5 mins.")
}
