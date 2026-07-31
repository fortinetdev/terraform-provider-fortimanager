// Copyright 2026 Fortinet, Inc. All rights reserved.
// Author: Xing Li (@lix-fortinet), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Xing Li (@lix-fortinet), Hongbin Lu (@fgtdev-hblu), Yue Wang (@yuew-ftnt)

// Description: SDK for resource fortimanager_dvmdb_group_objectmember

package forticlient

import (
	"fmt"
)

// CreateDvmdbGroupObjectMember API operation for FortiManager creates a new GroupObject Member.
// Returns the index value of the GroupObject Member and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dvmdb - group object member chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateDvmdbGroupObjectMember(params *map[string]interface{}, paradict, wsParams map[string]string) (output map[string]interface{}, err error) {
	path := "/dvmdb/[*]/group/{group}/object member"
	path, err = replaceParaWithValue(path, paradict)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	requestInput := &requestInput{}

	requestInput.fortiSDKClient = c
	requestInput.method = "add"
	requestInput.path = path
	requestInput.bodyParams = params
	requestInput.wsParams = wsParams
	requestInput.bMove = false

	output, err = createUpdate(requestInput)
	return
}

// DeleteDvmdbGroupObjectMember API operation for FortiManager deletes the specified GroupObject Member.
// Returns error for service API and SDK errors.
// See the dvmdb - group object member chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteDvmdbGroupObjectMember(params *map[string]interface{}, paradict, wsParams map[string]string) (output map[string]interface{}, err error) {
	path := "/dvmdb/[*]/group/{group}/object member"
	path, err = replaceParaWithValue(path, paradict)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	requestInput := &requestInput{}

	requestInput.fortiSDKClient = c
	requestInput.method = "delete"
	requestInput.path = path
	requestInput.bodyParams = params
	requestInput.wsParams = wsParams
	requestInput.bMove = false

	output, err = createUpdate(requestInput)
	return
}
