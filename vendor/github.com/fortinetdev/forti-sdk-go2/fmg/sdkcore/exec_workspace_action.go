package forticlient

// CreateUpdateExecWorkspaceAction API operation for locking ADOMs, devices, or policy packages so that an administrator can prevent other administrators from making changes to the elements that they are working in.
// Returns error for service API and SDK errors.
func (c *FortiSDKClient) CreateUpdateExecWorkspaceAction(globaladom, action, target, param string) (output map[string]interface{}, err error) {
	path := "/dvmdb/[*]/workspace"
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
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "exec", nil, output, false)
	return
}


// // ReadFirewallSecurityPolicySort API operation for FortiOS to read the firewall policies sort results
// // Returns sort status
// // Returns error for service API and SDK errors.
// func (c *FortiSDKClient) ReadFirewallSecurityPolicySort(sortby, sortdirection string) (sorted bool, err error) {
// 	idlist, err := getPolicyList(c)
// 	if err != nil {
// 		err = fmt.Errorf("sort err %s", err)
// 		return
// 	}

// 	bsorted := bPolicyListSorted(idlist, sortby, sortdirection)
// 	log.Printf("shengh: %v", bsorted)
// 	if bsorted == true {
// 		sorted = true
// 		return
// 	}

// 	sorted = false

// 	return
// }


