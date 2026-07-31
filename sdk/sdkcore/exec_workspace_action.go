package forticlient

// CreateUpdateExecWorkspaceAction API operation for locking ADOMs, devices, or policy packages so that an administrator can prevent other administrators from making changes to the elements that they are working in.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) CreateUpdateExecWorkspaceAction(globaladom, action, target, param string) (output map[string]interface{}, err error) {
	path := "/dvmdb/" + globaladom + "/workspace"
	path += "/"
	path += action
	if target != "" {
		path += "/"
		path += target

		if param != "" {
			path += "/"
			path += param
		}
	}

	requestInput := &requestInput{}

	requestInput.fortiSDKClient = c
	requestInput.method = "exec"
	requestInput.path = path
	requestInput.bMove = false

	output, err = createUpdate(requestInput)
	return
}

// GetExecWorkspaceAction API operation for getting lock information.
// Returns lock information and error for service API and SDK errors.
func (c *FortiSDKClient) GetExecWorkspaceAction(globaladom, target, param string) (output map[string]interface{}, err error) {
	path := "/dvmdb/" + globaladom + "/workspace/lockinfo"
	path += "/"
	if target != "" {
		path += "/"
		path += target

		if param != "" {
			path += "/"
			path += param
		}
	}

	requestInput := &requestInput{}

	requestInput.fortiSDKClient = c
	requestInput.method = "get"
	requestInput.path = path
	requestInput.bMove = false

	output, err = read(requestInput)
	return
}
