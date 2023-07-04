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

	output, err = createUpdate(c, path, "exec", nil, false)
	return
}
