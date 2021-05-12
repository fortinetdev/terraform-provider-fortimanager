// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Description: SDK for FortiManager Provider

package forticlient

import (
	"fmt"
)



// UpdatePackagesFirewallPolicyMove API operation for FortiManager updates the specified FirewallPolicyMove.
// Returns the index value of the FirewallPolicyMove and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall policy move chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdatePackagesFirewallPolicyMove(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/policy/{policy}"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "move", params, output, true)
	return
}


// ReadPackagesFirewallPolicyMove API operation for FortiManager gets the FirewallPolicyMove
// with the specified index value.
// Returns the requested FirewallPolicyMove value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall policy move chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadPackagesFirewallPolicyMove(globaladom, mkey string, paralist []string) (listTmp []interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	listTmp, err = readMove(c, globaladom, path, "get", false)
	return
}


// UpdatePackagesFirewallMulticastPolicyMove API operation for FortiManager updates the specified FirewallMulticast PolicyMove.
// Returns the index value of the FirewallMulticast PolicyMove and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall multicast-policy move chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdatePackagesFirewallMulticastPolicyMove(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/multicast-policy/{multicast-policy}"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "move", params, output, true)
	return
}


// ReadPackagesFirewallMulticastPolicyMove API operation for FortiManager gets the FirewallMulticast PolicyMove
// with the specified index value.
// Returns the requested FirewallMulticast PolicyMove value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall multicast-policy move chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadPackagesFirewallMulticastPolicyMove(globaladom, mkey string, paralist []string) (listTmp []interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/multicast-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	listTmp, err = readMove(c, globaladom, path, "get", false)
	return
}


// UpdatePackagesFirewallInterfacePolicyMove API operation for FortiManager updates the specified FirewallInterface PolicyMove.
// Returns the index value of the FirewallInterface PolicyMove and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall interface-policy move chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdatePackagesFirewallInterfacePolicyMove(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/interface-policy/{interface-policy}"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "move", params, output, true)
	return
}


// ReadPackagesFirewallInterfacePolicyMove API operation for FortiManager gets the FirewallInterface PolicyMove
// with the specified index value.
// Returns the requested FirewallInterface PolicyMove value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall interface-policy move chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadPackagesFirewallInterfacePolicyMove(globaladom, mkey string, paralist []string) (listTmp []interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/interface-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	listTmp, err = readMove(c, globaladom, path, "get", false)
	return
}


// UpdatePackagesFirewallInterfacePolicy6Move API operation for FortiManager updates the specified FirewallInterface Policy6Move.
// Returns the index value of the FirewallInterface Policy6Move and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall interface-policy6 move chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdatePackagesFirewallInterfacePolicy6Move(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/interface-policy6/{interface-policy6}"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "move", params, output, true)
	return
}


// ReadPackagesFirewallInterfacePolicy6Move API operation for FortiManager gets the FirewallInterface Policy6Move
// with the specified index value.
// Returns the requested FirewallInterface Policy6Move value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall interface-policy6 move chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadPackagesFirewallInterfacePolicy6Move(globaladom, mkey string, paralist []string) (listTmp []interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/interface-policy6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	listTmp, err = readMove(c, globaladom, path, "get", false)
	return
}


// UpdatePackagesFirewallLocalInPolicyMove API operation for FortiManager updates the specified FirewallLocal In PolicyMove.
// Returns the index value of the FirewallLocal In PolicyMove and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall local-in-policy move chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdatePackagesFirewallLocalInPolicyMove(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/local-in-policy/{local-in-policy}"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "move", params, output, true)
	return
}


// ReadPackagesFirewallLocalInPolicyMove API operation for FortiManager gets the FirewallLocal In PolicyMove
// with the specified index value.
// Returns the requested FirewallLocal In PolicyMove value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall local-in-policy move chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadPackagesFirewallLocalInPolicyMove(globaladom, mkey string, paralist []string) (listTmp []interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/local-in-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	listTmp, err = readMove(c, globaladom, path, "get", false)
	return
}


// UpdatePackagesFirewallLocalInPolicy6Move API operation for FortiManager updates the specified FirewallLocal In Policy6Move.
// Returns the index value of the FirewallLocal In Policy6Move and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall local-in-policy6 move chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdatePackagesFirewallLocalInPolicy6Move(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/local-in-policy6/{local-in-policy6}"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "move", params, output, true)
	return
}


// ReadPackagesFirewallLocalInPolicy6Move API operation for FortiManager gets the FirewallLocal In Policy6Move
// with the specified index value.
// Returns the requested FirewallLocal In Policy6Move value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall local-in-policy6 move chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadPackagesFirewallLocalInPolicy6Move(globaladom, mkey string, paralist []string) (listTmp []interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/local-in-policy6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	listTmp, err = readMove(c, globaladom, path, "get", false)
	return
}


// UpdatePackagesFirewallPolicy64Move API operation for FortiManager updates the specified FirewallPolicy64Move.
// Returns the index value of the FirewallPolicy64Move and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall policy64 move chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdatePackagesFirewallPolicy64Move(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/policy64/{policy64}"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "move", params, output, true)
	return
}


// ReadPackagesFirewallPolicy64Move API operation for FortiManager gets the FirewallPolicy64Move
// with the specified index value.
// Returns the requested FirewallPolicy64Move value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall policy64 move chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadPackagesFirewallPolicy64Move(globaladom, mkey string, paralist []string) (listTmp []interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/policy64"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	listTmp, err = readMove(c, globaladom, path, "get", false)
	return
}


// UpdatePackagesFirewallDosPolicyMove API operation for FortiManager updates the specified FirewallDos PolicyMove.
// Returns the index value of the FirewallDos PolicyMove and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall DoS-policy move chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdatePackagesFirewallDosPolicyMove(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/DoS-policy/{DoS-policy}"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "move", params, output, true)
	return
}


// ReadPackagesFirewallDosPolicyMove API operation for FortiManager gets the FirewallDos PolicyMove
// with the specified index value.
// Returns the requested FirewallDos PolicyMove value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall DoS-policy move chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadPackagesFirewallDosPolicyMove(globaladom, mkey string, paralist []string) (listTmp []interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/DoS-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	listTmp, err = readMove(c, globaladom, path, "get", false)
	return
}


// UpdatePackagesFirewallPolicy46Move API operation for FortiManager updates the specified FirewallPolicy46Move.
// Returns the index value of the FirewallPolicy46Move and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall policy46 move chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdatePackagesFirewallPolicy46Move(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/policy46/{policy46}"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "move", params, output, true)
	return
}


// ReadPackagesFirewallPolicy46Move API operation for FortiManager gets the FirewallPolicy46Move
// with the specified index value.
// Returns the requested FirewallPolicy46Move value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall policy46 move chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadPackagesFirewallPolicy46Move(globaladom, mkey string, paralist []string) (listTmp []interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/policy46"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	listTmp, err = readMove(c, globaladom, path, "get", false)
	return
}


// UpdatePackagesFirewallDosPolicy6Move API operation for FortiManager updates the specified FirewallDos Policy6Move.
// Returns the index value of the FirewallDos Policy6Move and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall DoS-policy6 move chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdatePackagesFirewallDosPolicy6Move(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/DoS-policy6/{DoS-policy6}"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "move", params, output, true)
	return
}


// ReadPackagesFirewallDosPolicy6Move API operation for FortiManager gets the FirewallDos Policy6Move
// with the specified index value.
// Returns the requested FirewallDos Policy6Move value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall DoS-policy6 move chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadPackagesFirewallDosPolicy6Move(globaladom, mkey string, paralist []string) (listTmp []interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/DoS-policy6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	listTmp, err = readMove(c, globaladom, path, "get", false)
	return
}


// UpdatePackagesFirewallShapingPolicyMove API operation for FortiManager updates the specified FirewallShaping PolicyMove.
// Returns the index value of the FirewallShaping PolicyMove and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall shaping-policy move chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdatePackagesFirewallShapingPolicyMove(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/shaping-policy/{shaping-policy}"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "move", params, output, true)
	return
}


// ReadPackagesFirewallShapingPolicyMove API operation for FortiManager gets the FirewallShaping PolicyMove
// with the specified index value.
// Returns the requested FirewallShaping PolicyMove value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall shaping-policy move chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadPackagesFirewallShapingPolicyMove(globaladom, mkey string, paralist []string) (listTmp []interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/shaping-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	listTmp, err = readMove(c, globaladom, path, "get", false)
	return
}


// UpdatePackagesFirewallMulticastPolicy6Move API operation for FortiManager updates the specified FirewallMulticast Policy6Move.
// Returns the index value of the FirewallMulticast Policy6Move and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall multicast-policy6 move chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdatePackagesFirewallMulticastPolicy6Move(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/multicast-policy6/{multicast-policy6}"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "move", params, output, true)
	return
}


// ReadPackagesFirewallMulticastPolicy6Move API operation for FortiManager gets the FirewallMulticast Policy6Move
// with the specified index value.
// Returns the requested FirewallMulticast Policy6Move value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall multicast-policy6 move chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadPackagesFirewallMulticastPolicy6Move(globaladom, mkey string, paralist []string) (listTmp []interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/multicast-policy6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	listTmp, err = readMove(c, globaladom, path, "get", false)
	return
}


// UpdatePackagesFirewallCentralSnatMapMove API operation for FortiManager updates the specified FirewallCentral Snat MapMove.
// Returns the index value of the FirewallCentral Snat MapMove and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall central-snat-map move chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdatePackagesFirewallCentralSnatMapMove(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/central-snat-map/{central-snat-map}"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "move", params, output, true)
	return
}


// ReadPackagesFirewallCentralSnatMapMove API operation for FortiManager gets the FirewallCentral Snat MapMove
// with the specified index value.
// Returns the requested FirewallCentral Snat MapMove value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall central-snat-map move chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadPackagesFirewallCentralSnatMapMove(globaladom, mkey string, paralist []string) (listTmp []interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/central-snat-map"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	listTmp, err = readMove(c, globaladom, path, "get", false)
	return
}


// UpdatePackagesFirewallProxyPolicyMove API operation for FortiManager updates the specified FirewallProxy PolicyMove.
// Returns the index value of the FirewallProxy PolicyMove and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall proxy-policy move chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdatePackagesFirewallProxyPolicyMove(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/proxy-policy/{proxy-policy}"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "move", params, output, true)
	return
}


// ReadPackagesFirewallProxyPolicyMove API operation for FortiManager gets the FirewallProxy PolicyMove
// with the specified index value.
// Returns the requested FirewallProxy PolicyMove value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall proxy-policy move chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadPackagesFirewallProxyPolicyMove(globaladom, mkey string, paralist []string) (listTmp []interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/proxy-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	listTmp, err = readMove(c, globaladom, path, "get", false)
	return
}


// UpdatePackagesFirewallSecurityPolicyMove API operation for FortiManager updates the specified FirewallSecurity PolicyMove.
// Returns the index value of the FirewallSecurity PolicyMove and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall security-policy move chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdatePackagesFirewallSecurityPolicyMove(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/security-policy/{security-policy}"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "move", params, output, true)
	return
}


// ReadPackagesFirewallSecurityPolicyMove API operation for FortiManager gets the FirewallSecurity PolicyMove
// with the specified index value.
// Returns the requested FirewallSecurity PolicyMove value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall security-policy move chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadPackagesFirewallSecurityPolicyMove(globaladom, mkey string, paralist []string) (listTmp []interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/security-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	listTmp, err = readMove(c, globaladom, path, "get", false)
	return
}

// CreatePackagesPkg API operation for FortiManager creates a new Pkg.
// Returns the index value of the Pkg and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - pkg chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreatePackagesPkg(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/pkg/[*]"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdatePackagesPkg API operation for FortiManager updates the specified Pkg.
// Returns the index value of the Pkg and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - pkg chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdatePackagesPkg(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/pkg/[*]"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeletePackagesPkg API operation for FortiManager deletes the specified Pkg.
// Returns error for service API and SDK errors.
// See the packages - pkg chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeletePackagesPkg(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/pkg/[*]"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadPackagesPkg API operation for FortiManager gets the Pkg
// with the specified index value.
// Returns the requested Pkg value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - pkg chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadPackagesPkg(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/pkg/[*]"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreatePackagesFirewallPolicy API operation for FortiManager creates a new FirewallPolicy.
// Returns the index value of the FirewallPolicy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreatePackagesFirewallPolicy(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdatePackagesFirewallPolicy API operation for FortiManager updates the specified FirewallPolicy.
// Returns the index value of the FirewallPolicy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdatePackagesFirewallPolicy(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeletePackagesFirewallPolicy API operation for FortiManager deletes the specified FirewallPolicy.
// Returns error for service API and SDK errors.
// See the packages - firewall policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeletePackagesFirewallPolicy(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadPackagesFirewallPolicy API operation for FortiManager gets the FirewallPolicy
// with the specified index value.
// Returns the requested FirewallPolicy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadPackagesFirewallPolicy(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreatePackagesGlobalHeaderPolicy API operation for FortiManager creates a new GlobalHeaderPolicy.
// Returns the index value of the GlobalHeaderPolicy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - global header policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreatePackagesGlobalHeaderPolicy(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/global/header/policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdatePackagesGlobalHeaderPolicy API operation for FortiManager updates the specified GlobalHeaderPolicy.
// Returns the index value of the GlobalHeaderPolicy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - global header policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdatePackagesGlobalHeaderPolicy(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/global/header/policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeletePackagesGlobalHeaderPolicy API operation for FortiManager deletes the specified GlobalHeaderPolicy.
// Returns error for service API and SDK errors.
// See the packages - global header policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeletePackagesGlobalHeaderPolicy(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/pkg/{pkg}/global/header/policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadPackagesGlobalHeaderPolicy API operation for FortiManager gets the GlobalHeaderPolicy
// with the specified index value.
// Returns the requested GlobalHeaderPolicy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - global header policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadPackagesGlobalHeaderPolicy(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/global/header/policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreatePackagesGlobalFooterPolicy API operation for FortiManager creates a new GlobalFooterPolicy.
// Returns the index value of the GlobalFooterPolicy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - global footer policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreatePackagesGlobalFooterPolicy(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/global/footer/policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdatePackagesGlobalFooterPolicy API operation for FortiManager updates the specified GlobalFooterPolicy.
// Returns the index value of the GlobalFooterPolicy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - global footer policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdatePackagesGlobalFooterPolicy(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/global/footer/policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeletePackagesGlobalFooterPolicy API operation for FortiManager deletes the specified GlobalFooterPolicy.
// Returns error for service API and SDK errors.
// See the packages - global footer policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeletePackagesGlobalFooterPolicy(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/pkg/{pkg}/global/footer/policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadPackagesGlobalFooterPolicy API operation for FortiManager gets the GlobalFooterPolicy
// with the specified index value.
// Returns the requested GlobalFooterPolicy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - global footer policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadPackagesGlobalFooterPolicy(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/global/footer/policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreatePackagesGlobalHeaderPolicy6 API operation for FortiManager creates a new GlobalHeaderPolicy6.
// Returns the index value of the GlobalHeaderPolicy6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - global header policy6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreatePackagesGlobalHeaderPolicy6(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/global/header/policy6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdatePackagesGlobalHeaderPolicy6 API operation for FortiManager updates the specified GlobalHeaderPolicy6.
// Returns the index value of the GlobalHeaderPolicy6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - global header policy6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdatePackagesGlobalHeaderPolicy6(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/global/header/policy6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeletePackagesGlobalHeaderPolicy6 API operation for FortiManager deletes the specified GlobalHeaderPolicy6.
// Returns error for service API and SDK errors.
// See the packages - global header policy6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeletePackagesGlobalHeaderPolicy6(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/pkg/{pkg}/global/header/policy6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadPackagesGlobalHeaderPolicy6 API operation for FortiManager gets the GlobalHeaderPolicy6
// with the specified index value.
// Returns the requested GlobalHeaderPolicy6 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - global header policy6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadPackagesGlobalHeaderPolicy6(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/global/header/policy6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreatePackagesGlobalFooterPolicy6 API operation for FortiManager creates a new GlobalFooterPolicy6.
// Returns the index value of the GlobalFooterPolicy6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - global footer policy6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreatePackagesGlobalFooterPolicy6(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/global/footer/policy6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdatePackagesGlobalFooterPolicy6 API operation for FortiManager updates the specified GlobalFooterPolicy6.
// Returns the index value of the GlobalFooterPolicy6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - global footer policy6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdatePackagesGlobalFooterPolicy6(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/global/footer/policy6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeletePackagesGlobalFooterPolicy6 API operation for FortiManager deletes the specified GlobalFooterPolicy6.
// Returns error for service API and SDK errors.
// See the packages - global footer policy6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeletePackagesGlobalFooterPolicy6(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/pkg/{pkg}/global/footer/policy6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadPackagesGlobalFooterPolicy6 API operation for FortiManager gets the GlobalFooterPolicy6
// with the specified index value.
// Returns the requested GlobalFooterPolicy6 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - global footer policy6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadPackagesGlobalFooterPolicy6(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/global/footer/policy6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreatePackagesFirewallMulticastPolicy API operation for FortiManager creates a new FirewallMulticast Policy.
// Returns the index value of the FirewallMulticast Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall multicast-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreatePackagesFirewallMulticastPolicy(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/multicast-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdatePackagesFirewallMulticastPolicy API operation for FortiManager updates the specified FirewallMulticast Policy.
// Returns the index value of the FirewallMulticast Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall multicast-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdatePackagesFirewallMulticastPolicy(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/multicast-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeletePackagesFirewallMulticastPolicy API operation for FortiManager deletes the specified FirewallMulticast Policy.
// Returns error for service API and SDK errors.
// See the packages - firewall multicast-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeletePackagesFirewallMulticastPolicy(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/multicast-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadPackagesFirewallMulticastPolicy API operation for FortiManager gets the FirewallMulticast Policy
// with the specified index value.
// Returns the requested FirewallMulticast Policy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall multicast-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadPackagesFirewallMulticastPolicy(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/multicast-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreatePackagesFirewallInterfacePolicy API operation for FortiManager creates a new FirewallInterface Policy.
// Returns the index value of the FirewallInterface Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall interface-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreatePackagesFirewallInterfacePolicy(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/interface-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdatePackagesFirewallInterfacePolicy API operation for FortiManager updates the specified FirewallInterface Policy.
// Returns the index value of the FirewallInterface Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall interface-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdatePackagesFirewallInterfacePolicy(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/interface-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeletePackagesFirewallInterfacePolicy API operation for FortiManager deletes the specified FirewallInterface Policy.
// Returns error for service API and SDK errors.
// See the packages - firewall interface-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeletePackagesFirewallInterfacePolicy(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/interface-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadPackagesFirewallInterfacePolicy API operation for FortiManager gets the FirewallInterface Policy
// with the specified index value.
// Returns the requested FirewallInterface Policy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall interface-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadPackagesFirewallInterfacePolicy(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/interface-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreatePackagesFirewallInterfacePolicy6 API operation for FortiManager creates a new FirewallInterface Policy6.
// Returns the index value of the FirewallInterface Policy6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall interface-policy6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreatePackagesFirewallInterfacePolicy6(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/interface-policy6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdatePackagesFirewallInterfacePolicy6 API operation for FortiManager updates the specified FirewallInterface Policy6.
// Returns the index value of the FirewallInterface Policy6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall interface-policy6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdatePackagesFirewallInterfacePolicy6(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/interface-policy6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeletePackagesFirewallInterfacePolicy6 API operation for FortiManager deletes the specified FirewallInterface Policy6.
// Returns error for service API and SDK errors.
// See the packages - firewall interface-policy6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeletePackagesFirewallInterfacePolicy6(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/interface-policy6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadPackagesFirewallInterfacePolicy6 API operation for FortiManager gets the FirewallInterface Policy6
// with the specified index value.
// Returns the requested FirewallInterface Policy6 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall interface-policy6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadPackagesFirewallInterfacePolicy6(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/interface-policy6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreatePackagesFirewallLocalInPolicy API operation for FortiManager creates a new FirewallLocal In Policy.
// Returns the index value of the FirewallLocal In Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall local-in-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreatePackagesFirewallLocalInPolicy(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/local-in-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdatePackagesFirewallLocalInPolicy API operation for FortiManager updates the specified FirewallLocal In Policy.
// Returns the index value of the FirewallLocal In Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall local-in-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdatePackagesFirewallLocalInPolicy(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/local-in-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeletePackagesFirewallLocalInPolicy API operation for FortiManager deletes the specified FirewallLocal In Policy.
// Returns error for service API and SDK errors.
// See the packages - firewall local-in-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeletePackagesFirewallLocalInPolicy(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/local-in-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadPackagesFirewallLocalInPolicy API operation for FortiManager gets the FirewallLocal In Policy
// with the specified index value.
// Returns the requested FirewallLocal In Policy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall local-in-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadPackagesFirewallLocalInPolicy(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/local-in-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreatePackagesFirewallLocalInPolicy6 API operation for FortiManager creates a new FirewallLocal In Policy6.
// Returns the index value of the FirewallLocal In Policy6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall local-in-policy6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreatePackagesFirewallLocalInPolicy6(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/local-in-policy6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdatePackagesFirewallLocalInPolicy6 API operation for FortiManager updates the specified FirewallLocal In Policy6.
// Returns the index value of the FirewallLocal In Policy6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall local-in-policy6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdatePackagesFirewallLocalInPolicy6(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/local-in-policy6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeletePackagesFirewallLocalInPolicy6 API operation for FortiManager deletes the specified FirewallLocal In Policy6.
// Returns error for service API and SDK errors.
// See the packages - firewall local-in-policy6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeletePackagesFirewallLocalInPolicy6(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/local-in-policy6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadPackagesFirewallLocalInPolicy6 API operation for FortiManager gets the FirewallLocal In Policy6
// with the specified index value.
// Returns the requested FirewallLocal In Policy6 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall local-in-policy6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadPackagesFirewallLocalInPolicy6(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/local-in-policy6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreatePackagesFirewallPolicy64 API operation for FortiManager creates a new FirewallPolicy64.
// Returns the index value of the FirewallPolicy64 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall policy64 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreatePackagesFirewallPolicy64(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/policy64"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdatePackagesFirewallPolicy64 API operation for FortiManager updates the specified FirewallPolicy64.
// Returns the index value of the FirewallPolicy64 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall policy64 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdatePackagesFirewallPolicy64(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/policy64"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeletePackagesFirewallPolicy64 API operation for FortiManager deletes the specified FirewallPolicy64.
// Returns error for service API and SDK errors.
// See the packages - firewall policy64 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeletePackagesFirewallPolicy64(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/policy64"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadPackagesFirewallPolicy64 API operation for FortiManager gets the FirewallPolicy64
// with the specified index value.
// Returns the requested FirewallPolicy64 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall policy64 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadPackagesFirewallPolicy64(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/policy64"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreatePackagesFirewallDosPolicy API operation for FortiManager creates a new FirewallDos Policy.
// Returns the index value of the FirewallDos Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall DoS-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreatePackagesFirewallDosPolicy(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/DoS-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdatePackagesFirewallDosPolicy API operation for FortiManager updates the specified FirewallDos Policy.
// Returns the index value of the FirewallDos Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall DoS-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdatePackagesFirewallDosPolicy(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/DoS-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeletePackagesFirewallDosPolicy API operation for FortiManager deletes the specified FirewallDos Policy.
// Returns error for service API and SDK errors.
// See the packages - firewall DoS-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeletePackagesFirewallDosPolicy(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/DoS-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadPackagesFirewallDosPolicy API operation for FortiManager gets the FirewallDos Policy
// with the specified index value.
// Returns the requested FirewallDos Policy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall DoS-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadPackagesFirewallDosPolicy(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/DoS-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreatePackagesFirewallPolicy46 API operation for FortiManager creates a new FirewallPolicy46.
// Returns the index value of the FirewallPolicy46 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall policy46 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreatePackagesFirewallPolicy46(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/policy46"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdatePackagesFirewallPolicy46 API operation for FortiManager updates the specified FirewallPolicy46.
// Returns the index value of the FirewallPolicy46 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall policy46 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdatePackagesFirewallPolicy46(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/policy46"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeletePackagesFirewallPolicy46 API operation for FortiManager deletes the specified FirewallPolicy46.
// Returns error for service API and SDK errors.
// See the packages - firewall policy46 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeletePackagesFirewallPolicy46(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/policy46"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadPackagesFirewallPolicy46 API operation for FortiManager gets the FirewallPolicy46
// with the specified index value.
// Returns the requested FirewallPolicy46 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall policy46 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadPackagesFirewallPolicy46(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/policy46"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreatePackagesFirewallDosPolicy6 API operation for FortiManager creates a new FirewallDos Policy6.
// Returns the index value of the FirewallDos Policy6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall DoS-policy6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreatePackagesFirewallDosPolicy6(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/DoS-policy6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdatePackagesFirewallDosPolicy6 API operation for FortiManager updates the specified FirewallDos Policy6.
// Returns the index value of the FirewallDos Policy6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall DoS-policy6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdatePackagesFirewallDosPolicy6(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/DoS-policy6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeletePackagesFirewallDosPolicy6 API operation for FortiManager deletes the specified FirewallDos Policy6.
// Returns error for service API and SDK errors.
// See the packages - firewall DoS-policy6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeletePackagesFirewallDosPolicy6(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/DoS-policy6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadPackagesFirewallDosPolicy6 API operation for FortiManager gets the FirewallDos Policy6
// with the specified index value.
// Returns the requested FirewallDos Policy6 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall DoS-policy6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadPackagesFirewallDosPolicy6(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/DoS-policy6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreatePackagesFirewallShapingPolicy API operation for FortiManager creates a new FirewallShaping Policy.
// Returns the index value of the FirewallShaping Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall shaping-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreatePackagesFirewallShapingPolicy(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/shaping-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdatePackagesFirewallShapingPolicy API operation for FortiManager updates the specified FirewallShaping Policy.
// Returns the index value of the FirewallShaping Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall shaping-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdatePackagesFirewallShapingPolicy(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/shaping-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeletePackagesFirewallShapingPolicy API operation for FortiManager deletes the specified FirewallShaping Policy.
// Returns error for service API and SDK errors.
// See the packages - firewall shaping-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeletePackagesFirewallShapingPolicy(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/shaping-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadPackagesFirewallShapingPolicy API operation for FortiManager gets the FirewallShaping Policy
// with the specified index value.
// Returns the requested FirewallShaping Policy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall shaping-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadPackagesFirewallShapingPolicy(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/shaping-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreatePackagesGlobalHeaderShapingPolicy API operation for FortiManager creates a new GlobalHeaderShaping Policy.
// Returns the index value of the GlobalHeaderShaping Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - global header shaping-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreatePackagesGlobalHeaderShapingPolicy(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/global/header/shaping-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdatePackagesGlobalHeaderShapingPolicy API operation for FortiManager updates the specified GlobalHeaderShaping Policy.
// Returns the index value of the GlobalHeaderShaping Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - global header shaping-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdatePackagesGlobalHeaderShapingPolicy(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/global/header/shaping-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeletePackagesGlobalHeaderShapingPolicy API operation for FortiManager deletes the specified GlobalHeaderShaping Policy.
// Returns error for service API and SDK errors.
// See the packages - global header shaping-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeletePackagesGlobalHeaderShapingPolicy(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/pkg/{pkg}/global/header/shaping-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadPackagesGlobalHeaderShapingPolicy API operation for FortiManager gets the GlobalHeaderShaping Policy
// with the specified index value.
// Returns the requested GlobalHeaderShaping Policy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - global header shaping-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadPackagesGlobalHeaderShapingPolicy(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/global/header/shaping-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreatePackagesGlobalFooterShapingPolicy API operation for FortiManager creates a new GlobalFooterShaping Policy.
// Returns the index value of the GlobalFooterShaping Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - global footer shaping-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreatePackagesGlobalFooterShapingPolicy(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/global/footer/shaping-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdatePackagesGlobalFooterShapingPolicy API operation for FortiManager updates the specified GlobalFooterShaping Policy.
// Returns the index value of the GlobalFooterShaping Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - global footer shaping-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdatePackagesGlobalFooterShapingPolicy(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/global/footer/shaping-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeletePackagesGlobalFooterShapingPolicy API operation for FortiManager deletes the specified GlobalFooterShaping Policy.
// Returns error for service API and SDK errors.
// See the packages - global footer shaping-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeletePackagesGlobalFooterShapingPolicy(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/pkg/{pkg}/global/footer/shaping-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadPackagesGlobalFooterShapingPolicy API operation for FortiManager gets the GlobalFooterShaping Policy
// with the specified index value.
// Returns the requested GlobalFooterShaping Policy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - global footer shaping-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadPackagesGlobalFooterShapingPolicy(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/global/footer/shaping-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreatePackagesFirewallMulticastPolicy6 API operation for FortiManager creates a new FirewallMulticast Policy6.
// Returns the index value of the FirewallMulticast Policy6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall multicast-policy6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreatePackagesFirewallMulticastPolicy6(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/multicast-policy6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdatePackagesFirewallMulticastPolicy6 API operation for FortiManager updates the specified FirewallMulticast Policy6.
// Returns the index value of the FirewallMulticast Policy6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall multicast-policy6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdatePackagesFirewallMulticastPolicy6(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/multicast-policy6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeletePackagesFirewallMulticastPolicy6 API operation for FortiManager deletes the specified FirewallMulticast Policy6.
// Returns error for service API and SDK errors.
// See the packages - firewall multicast-policy6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeletePackagesFirewallMulticastPolicy6(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/multicast-policy6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadPackagesFirewallMulticastPolicy6 API operation for FortiManager gets the FirewallMulticast Policy6
// with the specified index value.
// Returns the requested FirewallMulticast Policy6 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall multicast-policy6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadPackagesFirewallMulticastPolicy6(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/multicast-policy6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreatePackagesFirewallCentralSnatMap API operation for FortiManager creates a new FirewallCentral Snat Map.
// Returns the index value of the FirewallCentral Snat Map and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall central-snat-map chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreatePackagesFirewallCentralSnatMap(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/central-snat-map"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdatePackagesFirewallCentralSnatMap API operation for FortiManager updates the specified FirewallCentral Snat Map.
// Returns the index value of the FirewallCentral Snat Map and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall central-snat-map chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdatePackagesFirewallCentralSnatMap(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/central-snat-map"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeletePackagesFirewallCentralSnatMap API operation for FortiManager deletes the specified FirewallCentral Snat Map.
// Returns error for service API and SDK errors.
// See the packages - firewall central-snat-map chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeletePackagesFirewallCentralSnatMap(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/central-snat-map"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadPackagesFirewallCentralSnatMap API operation for FortiManager gets the FirewallCentral Snat Map
// with the specified index value.
// Returns the requested FirewallCentral Snat Map value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall central-snat-map chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadPackagesFirewallCentralSnatMap(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/central-snat-map"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreatePackagesFirewallProxyPolicy API operation for FortiManager creates a new FirewallProxy Policy.
// Returns the index value of the FirewallProxy Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall proxy-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreatePackagesFirewallProxyPolicy(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/proxy-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdatePackagesFirewallProxyPolicy API operation for FortiManager updates the specified FirewallProxy Policy.
// Returns the index value of the FirewallProxy Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall proxy-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdatePackagesFirewallProxyPolicy(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/proxy-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeletePackagesFirewallProxyPolicy API operation for FortiManager deletes the specified FirewallProxy Policy.
// Returns error for service API and SDK errors.
// See the packages - firewall proxy-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeletePackagesFirewallProxyPolicy(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/proxy-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadPackagesFirewallProxyPolicy API operation for FortiManager gets the FirewallProxy Policy
// with the specified index value.
// Returns the requested FirewallProxy Policy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall proxy-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadPackagesFirewallProxyPolicy(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/proxy-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreatePackagesAuthenticationRule API operation for FortiManager creates a new AuthenticationRule.
// Returns the index value of the AuthenticationRule and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - authentication rule chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreatePackagesAuthenticationRule(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/authentication/rule"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdatePackagesAuthenticationRule API operation for FortiManager updates the specified AuthenticationRule.
// Returns the index value of the AuthenticationRule and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - authentication rule chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdatePackagesAuthenticationRule(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/authentication/rule"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeletePackagesAuthenticationRule API operation for FortiManager deletes the specified AuthenticationRule.
// Returns error for service API and SDK errors.
// See the packages - authentication rule chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeletePackagesAuthenticationRule(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/pkg/{pkg}/authentication/rule"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadPackagesAuthenticationRule API operation for FortiManager gets the AuthenticationRule
// with the specified index value.
// Returns the requested AuthenticationRule value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - authentication rule chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadPackagesAuthenticationRule(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/authentication/rule"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdatePackagesAuthenticationSetting API operation for FortiManager updates the specified AuthenticationSetting.
// Returns the index value of the AuthenticationSetting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - authentication setting chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdatePackagesAuthenticationSetting(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/authentication/setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeletePackagesAuthenticationSetting API operation for FortiManager deletes the specified AuthenticationSetting.
// Returns error for service API and SDK errors.
// See the packages - authentication setting chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeletePackagesAuthenticationSetting(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for packages - authentication setting
	return
}

// ReadPackagesAuthenticationSetting API operation for FortiManager gets the AuthenticationSetting
// with the specified index value.
// Returns the requested AuthenticationSetting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - authentication setting chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadPackagesAuthenticationSetting(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/authentication/setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreatePackagesGlobalHeaderConsolidatedPolicy API operation for FortiManager creates a new GlobalHeaderConsolidatedPolicy.
// Returns the index value of the GlobalHeaderConsolidatedPolicy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - global header consolidated policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreatePackagesGlobalHeaderConsolidatedPolicy(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/global/header/consolidated/policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdatePackagesGlobalHeaderConsolidatedPolicy API operation for FortiManager updates the specified GlobalHeaderConsolidatedPolicy.
// Returns the index value of the GlobalHeaderConsolidatedPolicy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - global header consolidated policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdatePackagesGlobalHeaderConsolidatedPolicy(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/global/header/consolidated/policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeletePackagesGlobalHeaderConsolidatedPolicy API operation for FortiManager deletes the specified GlobalHeaderConsolidatedPolicy.
// Returns error for service API and SDK errors.
// See the packages - global header consolidated policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeletePackagesGlobalHeaderConsolidatedPolicy(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/pkg/{pkg}/global/header/consolidated/policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadPackagesGlobalHeaderConsolidatedPolicy API operation for FortiManager gets the GlobalHeaderConsolidatedPolicy
// with the specified index value.
// Returns the requested GlobalHeaderConsolidatedPolicy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - global header consolidated policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadPackagesGlobalHeaderConsolidatedPolicy(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/global/header/consolidated/policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreatePackagesGlobalFooterConsolidatedPolicy API operation for FortiManager creates a new GlobalFooterConsolidatedPolicy.
// Returns the index value of the GlobalFooterConsolidatedPolicy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - global footer consolidated policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreatePackagesGlobalFooterConsolidatedPolicy(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/global/footer/consolidated/policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdatePackagesGlobalFooterConsolidatedPolicy API operation for FortiManager updates the specified GlobalFooterConsolidatedPolicy.
// Returns the index value of the GlobalFooterConsolidatedPolicy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - global footer consolidated policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdatePackagesGlobalFooterConsolidatedPolicy(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/global/footer/consolidated/policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeletePackagesGlobalFooterConsolidatedPolicy API operation for FortiManager deletes the specified GlobalFooterConsolidatedPolicy.
// Returns error for service API and SDK errors.
// See the packages - global footer consolidated policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeletePackagesGlobalFooterConsolidatedPolicy(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/pkg/{pkg}/global/footer/consolidated/policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadPackagesGlobalFooterConsolidatedPolicy API operation for FortiManager gets the GlobalFooterConsolidatedPolicy
// with the specified index value.
// Returns the requested GlobalFooterConsolidatedPolicy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - global footer consolidated policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadPackagesGlobalFooterConsolidatedPolicy(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/global/footer/consolidated/policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreatePackagesFirewallSecurityPolicy API operation for FortiManager creates a new FirewallSecurity Policy.
// Returns the index value of the FirewallSecurity Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall security-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreatePackagesFirewallSecurityPolicy(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/security-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdatePackagesFirewallSecurityPolicy API operation for FortiManager updates the specified FirewallSecurity Policy.
// Returns the index value of the FirewallSecurity Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall security-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdatePackagesFirewallSecurityPolicy(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/security-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeletePackagesFirewallSecurityPolicy API operation for FortiManager deletes the specified FirewallSecurity Policy.
// Returns error for service API and SDK errors.
// See the packages - firewall security-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeletePackagesFirewallSecurityPolicy(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/security-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadPackagesFirewallSecurityPolicy API operation for FortiManager gets the FirewallSecurity Policy
// with the specified index value.
// Returns the requested FirewallSecurity Policy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the packages - firewall security-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadPackagesFirewallSecurityPolicy(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/pkg/{pkg}/firewall/security-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateDvmdbScriptExecute API operation for FortiManager updates the specified ScriptExecute.
// Returns the index value of the ScriptExecute and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dvmdb - script execute chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateDvmdbScriptExecute(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/dvmdb/[*]/script/execute"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "exec", params, output, false)
	return
}




// UpdateDvmCmdAddDevice API operation for FortiManager updates the specified CmdAddDevice.
// Returns the index value of the CmdAddDevice and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dvm - cmd add device chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateDvmCmdAddDevice(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/dvm/cmd/add/device"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "exec", params, output, false)
	return
}




// UpdateDvmCmdAddDevList API operation for FortiManager updates the specified CmdAddDev List.
// Returns the index value of the CmdAddDev List and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dvm - cmd add dev-list chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateDvmCmdAddDevList(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/dvm/cmd/add/dev-list"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "exec", params, output, false)
	return
}




// UpdateDvmCmdDiscoverDevice API operation for FortiManager updates the specified CmdDiscoverDevice.
// Returns the index value of the CmdDiscoverDevice and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dvm - cmd discover device chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateDvmCmdDiscoverDevice(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/dvm/cmd/discover/device"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "exec", params, output, false)
	return
}




// UpdateDvmCmdUpdateDevice API operation for FortiManager updates the specified CmdUpdateDevice.
// Returns the index value of the CmdUpdateDevice and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dvm - cmd update device chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateDvmCmdUpdateDevice(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/dvm/cmd/update/device"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "exec", params, output, false)
	return
}




// UpdateDvmCmdUpdateDevList API operation for FortiManager updates the specified CmdUpdateDev List.
// Returns the index value of the CmdUpdateDev List and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dvm - cmd update dev-list chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateDvmCmdUpdateDevList(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/dvm/cmd/update/dev-list"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "exec", params, output, false)
	return
}




// UpdateDvmCmdReloadDevList API operation for FortiManager updates the specified CmdReloadDev List.
// Returns the index value of the CmdReloadDev List and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dvm - cmd reload dev-list chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateDvmCmdReloadDevList(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/dvm/cmd/reload/dev-list"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "exec", params, output, false)
	return
}




// UpdateDvmCmdDelDevice API operation for FortiManager updates the specified CmdDelDevice.
// Returns the index value of the CmdDelDevice and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dvm - cmd del device chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateDvmCmdDelDevice(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/dvm/cmd/del/device"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "exec", params, output, false)
	return
}




// UpdateDvmCmdDelDevList API operation for FortiManager updates the specified CmdDelDev List.
// Returns the index value of the CmdDelDev List and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dvm - cmd del dev-list chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateDvmCmdDelDevList(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/dvm/cmd/del/dev-list"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "exec", params, output, false)
	return
}




// UpdateSecurityconsoleAbort API operation for FortiManager updates the specified Abort.
// Returns the index value of the Abort and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the securityconsole - abort chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSecurityconsoleAbort(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/securityconsole/abort"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "exec", params, output, false)
	return
}




// UpdateSecurityconsoleInstallPreview API operation for FortiManager updates the specified InstallPreview.
// Returns the index value of the InstallPreview and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the securityconsole - install preview chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSecurityconsoleInstallPreview(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/securityconsole/install/preview"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "exec", params, output, false)
	return
}




// UpdateSecurityconsolePreviewResult API operation for FortiManager updates the specified PreviewResult.
// Returns the index value of the PreviewResult and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the securityconsole - preview result chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSecurityconsolePreviewResult(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/securityconsole/preview/result"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "exec", params, output, false)
	return
}




// UpdateSecurityconsoleInstallPackage API operation for FortiManager updates the specified InstallPackage.
// Returns the index value of the InstallPackage and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the securityconsole - install package chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSecurityconsoleInstallPackage(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/securityconsole/install/package"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "exec", params, output, false)
	return
}




// UpdateSecurityconsoleInstallDevice API operation for FortiManager updates the specified InstallDevice.
// Returns the index value of the InstallDevice and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the securityconsole - install device chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSecurityconsoleInstallDevice(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/securityconsole/install/device"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "exec", params, output, false)
	return
}




// UpdateSecurityconsoleImportDevObjs API operation for FortiManager updates the specified ImportDevObjs.
// Returns the index value of the ImportDevObjs and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the securityconsole - import dev objs chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSecurityconsoleImportDevObjs(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/securityconsole/import/dev/objs"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "exec", params, output, false)
	return
}




// UpdateSecurityconsolePackageCommit API operation for FortiManager updates the specified PackageCommit.
// Returns the index value of the PackageCommit and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the securityconsole - package commit chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSecurityconsolePackageCommit(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/securityconsole/package/commit"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "exec", params, output, false)
	return
}




// UpdateSecurityconsolePackageCancelInstall API operation for FortiManager updates the specified PackageCancelInstall.
// Returns the index value of the PackageCancelInstall and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the securityconsole - package cancel install chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSecurityconsolePackageCancelInstall(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/securityconsole/package/cancel/install"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "exec", params, output, false)
	return
}




// UpdateSecurityconsolePackageClone API operation for FortiManager updates the specified PackageClone.
// Returns the index value of the PackageClone and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the securityconsole - package clone chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSecurityconsolePackageClone(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/securityconsole/package/clone"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "exec", params, output, false)
	return
}




// UpdateSecurityconsolePackageMove API operation for FortiManager updates the specified PackageMove.
// Returns the index value of the PackageMove and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the securityconsole - package move chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSecurityconsolePackageMove(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/securityconsole/package/move"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "exec", params, output, false)
	return
}




// UpdateSecurityconsoleReinstallPackage API operation for FortiManager updates the specified ReinstallPackage.
// Returns the index value of the ReinstallPackage and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the securityconsole - reinstall package chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSecurityconsoleReinstallPackage(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/securityconsole/reinstall/package"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "exec", params, output, false)
	return
}




// UpdateSecurityconsoleSignCertificateTemplate API operation for FortiManager updates the specified SignCertificateTemplate.
// Returns the index value of the SignCertificateTemplate and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the securityconsole - sign certificate template chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSecurityconsoleSignCertificateTemplate(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/securityconsole/sign/certificate/template"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "exec", params, output, false)
	return
}




// UpdateSecurityconsolePblockClone API operation for FortiManager updates the specified PblockClone.
// Returns the index value of the PblockClone and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the securityconsole - pblock clone chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSecurityconsolePblockClone(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/securityconsole/pblock/clone"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "exec", params, output, false)
	return
}



// CreateObjectWebfilterFtgdLocalCat API operation for FortiManager creates a new Ftgd Local Cat.
// Returns the index value of the Ftgd Local Cat and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object webfilter - ftgd-local-cat chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectWebfilterFtgdLocalCat(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/webfilter/ftgd-local-cat"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectWebfilterFtgdLocalCat API operation for FortiManager updates the specified Ftgd Local Cat.
// Returns the index value of the Ftgd Local Cat and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object webfilter - ftgd-local-cat chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectWebfilterFtgdLocalCat(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/webfilter/ftgd-local-cat"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectWebfilterFtgdLocalCat API operation for FortiManager deletes the specified Ftgd Local Cat.
// Returns error for service API and SDK errors.
// See the object webfilter - ftgd-local-cat chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectWebfilterFtgdLocalCat(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/webfilter/ftgd-local-cat"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectWebfilterFtgdLocalCat API operation for FortiManager gets the Ftgd Local Cat
// with the specified index value.
// Returns the requested Ftgd Local Cat value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object webfilter - ftgd-local-cat chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectWebfilterFtgdLocalCat(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/webfilter/ftgd-local-cat"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectWebfilterUrlfilter API operation for FortiManager creates a new Urlfilter.
// Returns the index value of the Urlfilter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object webfilter - urlfilter chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectWebfilterUrlfilter(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/webfilter/urlfilter"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectWebfilterUrlfilter API operation for FortiManager updates the specified Urlfilter.
// Returns the index value of the Urlfilter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object webfilter - urlfilter chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectWebfilterUrlfilter(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/webfilter/urlfilter"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectWebfilterUrlfilter API operation for FortiManager deletes the specified Urlfilter.
// Returns error for service API and SDK errors.
// See the object webfilter - urlfilter chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectWebfilterUrlfilter(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/webfilter/urlfilter"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectWebfilterUrlfilter API operation for FortiManager gets the Urlfilter
// with the specified index value.
// Returns the requested Urlfilter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object webfilter - urlfilter chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectWebfilterUrlfilter(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/webfilter/urlfilter"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectWebfilterFtgdLocalRating API operation for FortiManager creates a new Ftgd Local Rating.
// Returns the index value of the Ftgd Local Rating and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object webfilter - ftgd-local-rating chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectWebfilterFtgdLocalRating(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/webfilter/ftgd-local-rating"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectWebfilterFtgdLocalRating API operation for FortiManager updates the specified Ftgd Local Rating.
// Returns the index value of the Ftgd Local Rating and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object webfilter - ftgd-local-rating chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectWebfilterFtgdLocalRating(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/webfilter/ftgd-local-rating"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectWebfilterFtgdLocalRating API operation for FortiManager deletes the specified Ftgd Local Rating.
// Returns error for service API and SDK errors.
// See the object webfilter - ftgd-local-rating chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectWebfilterFtgdLocalRating(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/webfilter/ftgd-local-rating"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectWebfilterFtgdLocalRating API operation for FortiManager gets the Ftgd Local Rating
// with the specified index value.
// Returns the requested Ftgd Local Rating value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object webfilter - ftgd-local-rating chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectWebfilterFtgdLocalRating(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/webfilter/ftgd-local-rating"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectVpnCertificateCa API operation for FortiManager creates a new CertificateCa.
// Returns the index value of the CertificateCa and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object vpn - certificate ca chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectVpnCertificateCa(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/vpn/certificate/ca"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectVpnCertificateCa API operation for FortiManager updates the specified CertificateCa.
// Returns the index value of the CertificateCa and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object vpn - certificate ca chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectVpnCertificateCa(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/vpn/certificate/ca"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectVpnCertificateCa API operation for FortiManager deletes the specified CertificateCa.
// Returns error for service API and SDK errors.
// See the object vpn - certificate ca chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectVpnCertificateCa(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/vpn/certificate/ca"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectVpnCertificateCa API operation for FortiManager gets the CertificateCa
// with the specified index value.
// Returns the requested CertificateCa value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object vpn - certificate ca chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectVpnCertificateCa(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/vpn/certificate/ca"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectVpnCertificateRemote API operation for FortiManager creates a new CertificateRemote.
// Returns the index value of the CertificateRemote and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object vpn - certificate remote chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectVpnCertificateRemote(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/vpn/certificate/remote"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectVpnCertificateRemote API operation for FortiManager updates the specified CertificateRemote.
// Returns the index value of the CertificateRemote and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object vpn - certificate remote chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectVpnCertificateRemote(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/vpn/certificate/remote"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectVpnCertificateRemote API operation for FortiManager deletes the specified CertificateRemote.
// Returns error for service API and SDK errors.
// See the object vpn - certificate remote chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectVpnCertificateRemote(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/vpn/certificate/remote"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectVpnCertificateRemote API operation for FortiManager gets the CertificateRemote
// with the specified index value.
// Returns the requested CertificateRemote value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object vpn - certificate remote chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectVpnCertificateRemote(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/vpn/certificate/remote"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateObjectSystemFortiguard API operation for FortiManager updates the specified Fortiguard.
// Returns the index value of the Fortiguard and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object system - fortiguard chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectSystemFortiguard(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/system/fortiguard"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectSystemFortiguard API operation for FortiManager deletes the specified Fortiguard.
// Returns error for service API and SDK errors.
// See the object system - fortiguard chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectSystemFortiguard(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for object system - fortiguard
	return
}

// ReadObjectSystemFortiguard API operation for FortiManager gets the Fortiguard
// with the specified index value.
// Returns the requested Fortiguard value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object system - fortiguard chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectSystemFortiguard(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/system/fortiguard"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectIpsCustom API operation for FortiManager creates a new Custom.
// Returns the index value of the Custom and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object ips - custom chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectIpsCustom(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/ips/custom"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectIpsCustom API operation for FortiManager updates the specified Custom.
// Returns the index value of the Custom and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object ips - custom chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectIpsCustom(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/ips/custom"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectIpsCustom API operation for FortiManager deletes the specified Custom.
// Returns error for service API and SDK errors.
// See the object ips - custom chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectIpsCustom(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/ips/custom"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectIpsCustom API operation for FortiManager gets the Custom
// with the specified index value.
// Returns the requested Custom value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object ips - custom chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectIpsCustom(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/ips/custom"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectSystemDhcpServer API operation for FortiManager creates a new DhcpServer.
// Returns the index value of the DhcpServer and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object system - dhcp server chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectSystemDhcpServer(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/system/dhcp/server"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectSystemDhcpServer API operation for FortiManager updates the specified DhcpServer.
// Returns the index value of the DhcpServer and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object system - dhcp server chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectSystemDhcpServer(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/system/dhcp/server"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectSystemDhcpServer API operation for FortiManager deletes the specified DhcpServer.
// Returns error for service API and SDK errors.
// See the object system - dhcp server chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectSystemDhcpServer(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/system/dhcp/server"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectSystemDhcpServer API operation for FortiManager gets the DhcpServer
// with the specified index value.
// Returns the requested DhcpServer value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object system - dhcp server chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectSystemDhcpServer(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/system/dhcp/server"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectFirewallAddress API operation for FortiManager creates a new Address.
// Returns the index value of the Address and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - address chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectFirewallAddress(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/address"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectFirewallAddress API operation for FortiManager updates the specified Address.
// Returns the index value of the Address and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - address chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectFirewallAddress(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/address"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectFirewallAddress API operation for FortiManager deletes the specified Address.
// Returns error for service API and SDK errors.
// See the object firewall - address chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectFirewallAddress(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/firewall/address"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectFirewallAddress API operation for FortiManager gets the Address
// with the specified index value.
// Returns the requested Address value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - address chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectFirewallAddress(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/address"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectFirewallAddress6 API operation for FortiManager creates a new Address6.
// Returns the index value of the Address6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - address6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectFirewallAddress6(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/address6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectFirewallAddress6 API operation for FortiManager updates the specified Address6.
// Returns the index value of the Address6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - address6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectFirewallAddress6(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/address6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectFirewallAddress6 API operation for FortiManager deletes the specified Address6.
// Returns error for service API and SDK errors.
// See the object firewall - address6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectFirewallAddress6(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/firewall/address6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectFirewallAddress6 API operation for FortiManager gets the Address6
// with the specified index value.
// Returns the requested Address6 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - address6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectFirewallAddress6(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/address6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectFirewallAddrgrp API operation for FortiManager creates a new Addrgrp.
// Returns the index value of the Addrgrp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - addrgrp chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectFirewallAddrgrp(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/addrgrp"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectFirewallAddrgrp API operation for FortiManager updates the specified Addrgrp.
// Returns the index value of the Addrgrp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - addrgrp chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectFirewallAddrgrp(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/addrgrp"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectFirewallAddrgrp API operation for FortiManager deletes the specified Addrgrp.
// Returns error for service API and SDK errors.
// See the object firewall - addrgrp chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectFirewallAddrgrp(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/firewall/addrgrp"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectFirewallAddrgrp API operation for FortiManager gets the Addrgrp
// with the specified index value.
// Returns the requested Addrgrp value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - addrgrp chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectFirewallAddrgrp(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/addrgrp"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectUserAdgrp API operation for FortiManager creates a new Adgrp.
// Returns the index value of the Adgrp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - adgrp chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectUserAdgrp(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/adgrp"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectUserAdgrp API operation for FortiManager updates the specified Adgrp.
// Returns the index value of the Adgrp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - adgrp chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectUserAdgrp(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/adgrp"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectUserAdgrp API operation for FortiManager deletes the specified Adgrp.
// Returns error for service API and SDK errors.
// See the object user - adgrp chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectUserAdgrp(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/user/adgrp"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectUserAdgrp API operation for FortiManager gets the Adgrp
// with the specified index value.
// Returns the requested Adgrp value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - adgrp chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectUserAdgrp(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/adgrp"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectUserRadius API operation for FortiManager creates a new Radius.
// Returns the index value of the Radius and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - radius chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectUserRadius(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/radius"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectUserRadius API operation for FortiManager updates the specified Radius.
// Returns the index value of the Radius and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - radius chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectUserRadius(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/radius"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectUserRadius API operation for FortiManager deletes the specified Radius.
// Returns error for service API and SDK errors.
// See the object user - radius chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectUserRadius(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/user/radius"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectUserRadius API operation for FortiManager gets the Radius
// with the specified index value.
// Returns the requested Radius value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - radius chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectUserRadius(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/radius"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectUserLdap API operation for FortiManager creates a new Ldap.
// Returns the index value of the Ldap and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - ldap chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectUserLdap(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/ldap"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectUserLdap API operation for FortiManager updates the specified Ldap.
// Returns the index value of the Ldap and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - ldap chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectUserLdap(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/ldap"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectUserLdap API operation for FortiManager deletes the specified Ldap.
// Returns error for service API and SDK errors.
// See the object user - ldap chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectUserLdap(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/user/ldap"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectUserLdap API operation for FortiManager gets the Ldap
// with the specified index value.
// Returns the requested Ldap value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - ldap chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectUserLdap(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/ldap"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectUserLocal API operation for FortiManager creates a new Local.
// Returns the index value of the Local and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - local chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectUserLocal(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/local"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectUserLocal API operation for FortiManager updates the specified Local.
// Returns the index value of the Local and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - local chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectUserLocal(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/local"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectUserLocal API operation for FortiManager deletes the specified Local.
// Returns error for service API and SDK errors.
// See the object user - local chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectUserLocal(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/user/local"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectUserLocal API operation for FortiManager gets the Local
// with the specified index value.
// Returns the requested Local value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - local chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectUserLocal(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/local"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectUserPeer API operation for FortiManager creates a new Peer.
// Returns the index value of the Peer and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - peer chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectUserPeer(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/peer"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectUserPeer API operation for FortiManager updates the specified Peer.
// Returns the index value of the Peer and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - peer chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectUserPeer(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/peer"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectUserPeer API operation for FortiManager deletes the specified Peer.
// Returns error for service API and SDK errors.
// See the object user - peer chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectUserPeer(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/user/peer"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectUserPeer API operation for FortiManager gets the Peer
// with the specified index value.
// Returns the requested Peer value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - peer chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectUserPeer(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/peer"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectUserPeergrp API operation for FortiManager creates a new Peergrp.
// Returns the index value of the Peergrp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - peergrp chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectUserPeergrp(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/peergrp"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectUserPeergrp API operation for FortiManager updates the specified Peergrp.
// Returns the index value of the Peergrp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - peergrp chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectUserPeergrp(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/peergrp"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectUserPeergrp API operation for FortiManager deletes the specified Peergrp.
// Returns error for service API and SDK errors.
// See the object user - peergrp chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectUserPeergrp(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/user/peergrp"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectUserPeergrp API operation for FortiManager gets the Peergrp
// with the specified index value.
// Returns the requested Peergrp value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - peergrp chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectUserPeergrp(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/peergrp"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectUserGroup API operation for FortiManager creates a new Group.
// Returns the index value of the Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectUserGroup(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectUserGroup API operation for FortiManager updates the specified Group.
// Returns the index value of the Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectUserGroup(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectUserGroup API operation for FortiManager deletes the specified Group.
// Returns error for service API and SDK errors.
// See the object user - group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectUserGroup(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/user/group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectUserGroup API operation for FortiManager gets the Group
// with the specified index value.
// Returns the requested Group value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectUserGroup(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectFirewallServiceCustom API operation for FortiManager creates a new ServiceCustom.
// Returns the index value of the ServiceCustom and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - service custom chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectFirewallServiceCustom(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/service/custom"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectFirewallServiceCustom API operation for FortiManager updates the specified ServiceCustom.
// Returns the index value of the ServiceCustom and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - service custom chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectFirewallServiceCustom(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/service/custom"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectFirewallServiceCustom API operation for FortiManager deletes the specified ServiceCustom.
// Returns error for service API and SDK errors.
// See the object firewall - service custom chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectFirewallServiceCustom(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/firewall/service/custom"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectFirewallServiceCustom API operation for FortiManager gets the ServiceCustom
// with the specified index value.
// Returns the requested ServiceCustom value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - service custom chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectFirewallServiceCustom(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/service/custom"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectFirewallServiceGroup API operation for FortiManager creates a new ServiceGroup.
// Returns the index value of the ServiceGroup and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - service group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectFirewallServiceGroup(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/service/group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectFirewallServiceGroup API operation for FortiManager updates the specified ServiceGroup.
// Returns the index value of the ServiceGroup and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - service group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectFirewallServiceGroup(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/service/group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectFirewallServiceGroup API operation for FortiManager deletes the specified ServiceGroup.
// Returns error for service API and SDK errors.
// See the object firewall - service group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectFirewallServiceGroup(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/firewall/service/group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectFirewallServiceGroup API operation for FortiManager gets the ServiceGroup
// with the specified index value.
// Returns the requested ServiceGroup value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - service group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectFirewallServiceGroup(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/service/group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectFirewallScheduleOnetime API operation for FortiManager creates a new ScheduleOnetime.
// Returns the index value of the ScheduleOnetime and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - schedule onetime chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectFirewallScheduleOnetime(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/schedule/onetime"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectFirewallScheduleOnetime API operation for FortiManager updates the specified ScheduleOnetime.
// Returns the index value of the ScheduleOnetime and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - schedule onetime chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectFirewallScheduleOnetime(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/schedule/onetime"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectFirewallScheduleOnetime API operation for FortiManager deletes the specified ScheduleOnetime.
// Returns error for service API and SDK errors.
// See the object firewall - schedule onetime chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectFirewallScheduleOnetime(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/firewall/schedule/onetime"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectFirewallScheduleOnetime API operation for FortiManager gets the ScheduleOnetime
// with the specified index value.
// Returns the requested ScheduleOnetime value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - schedule onetime chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectFirewallScheduleOnetime(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/schedule/onetime"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectFirewallScheduleRecurring API operation for FortiManager creates a new ScheduleRecurring.
// Returns the index value of the ScheduleRecurring and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - schedule recurring chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectFirewallScheduleRecurring(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/schedule/recurring"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectFirewallScheduleRecurring API operation for FortiManager updates the specified ScheduleRecurring.
// Returns the index value of the ScheduleRecurring and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - schedule recurring chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectFirewallScheduleRecurring(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/schedule/recurring"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectFirewallScheduleRecurring API operation for FortiManager deletes the specified ScheduleRecurring.
// Returns error for service API and SDK errors.
// See the object firewall - schedule recurring chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectFirewallScheduleRecurring(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/firewall/schedule/recurring"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectFirewallScheduleRecurring API operation for FortiManager gets the ScheduleRecurring
// with the specified index value.
// Returns the requested ScheduleRecurring value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - schedule recurring chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectFirewallScheduleRecurring(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/schedule/recurring"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectFirewallIppool API operation for FortiManager creates a new Ippool.
// Returns the index value of the Ippool and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - ippool chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectFirewallIppool(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/ippool"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectFirewallIppool API operation for FortiManager updates the specified Ippool.
// Returns the index value of the Ippool and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - ippool chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectFirewallIppool(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/ippool"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectFirewallIppool API operation for FortiManager deletes the specified Ippool.
// Returns error for service API and SDK errors.
// See the object firewall - ippool chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectFirewallIppool(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/firewall/ippool"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectFirewallIppool API operation for FortiManager gets the Ippool
// with the specified index value.
// Returns the requested Ippool value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - ippool chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectFirewallIppool(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/ippool"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectFirewallVip API operation for FortiManager creates a new Vip.
// Returns the index value of the Vip and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - vip chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectFirewallVip(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/vip"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectFirewallVip API operation for FortiManager updates the specified Vip.
// Returns the index value of the Vip and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - vip chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectFirewallVip(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/vip"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectFirewallVip API operation for FortiManager deletes the specified Vip.
// Returns error for service API and SDK errors.
// See the object firewall - vip chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectFirewallVip(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/firewall/vip"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectFirewallVip API operation for FortiManager gets the Vip
// with the specified index value.
// Returns the requested Vip value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - vip chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectFirewallVip(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/vip"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectFirewallVipgrp API operation for FortiManager creates a new Vipgrp.
// Returns the index value of the Vipgrp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - vipgrp chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectFirewallVipgrp(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/vipgrp"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectFirewallVipgrp API operation for FortiManager updates the specified Vipgrp.
// Returns the index value of the Vipgrp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - vipgrp chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectFirewallVipgrp(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/vipgrp"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectFirewallVipgrp API operation for FortiManager deletes the specified Vipgrp.
// Returns error for service API and SDK errors.
// See the object firewall - vipgrp chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectFirewallVipgrp(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/firewall/vipgrp"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectFirewallVipgrp API operation for FortiManager gets the Vipgrp
// with the specified index value.
// Returns the requested Vipgrp value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - vipgrp chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectFirewallVipgrp(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/vipgrp"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectFirewallAddrgrp6 API operation for FortiManager creates a new Addrgrp6.
// Returns the index value of the Addrgrp6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - addrgrp6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectFirewallAddrgrp6(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/addrgrp6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectFirewallAddrgrp6 API operation for FortiManager updates the specified Addrgrp6.
// Returns the index value of the Addrgrp6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - addrgrp6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectFirewallAddrgrp6(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/addrgrp6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectFirewallAddrgrp6 API operation for FortiManager deletes the specified Addrgrp6.
// Returns error for service API and SDK errors.
// See the object firewall - addrgrp6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectFirewallAddrgrp6(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/firewall/addrgrp6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectFirewallAddrgrp6 API operation for FortiManager gets the Addrgrp6
// with the specified index value.
// Returns the requested Addrgrp6 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - addrgrp6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectFirewallAddrgrp6(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/addrgrp6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectIpsSensor API operation for FortiManager creates a new Sensor.
// Returns the index value of the Sensor and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object ips - sensor chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectIpsSensor(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/ips/sensor"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectIpsSensor API operation for FortiManager updates the specified Sensor.
// Returns the index value of the Sensor and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object ips - sensor chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectIpsSensor(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/ips/sensor"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectIpsSensor API operation for FortiManager deletes the specified Sensor.
// Returns error for service API and SDK errors.
// See the object ips - sensor chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectIpsSensor(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/ips/sensor"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectIpsSensor API operation for FortiManager gets the Sensor
// with the specified index value.
// Returns the requested Sensor value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object ips - sensor chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectIpsSensor(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/ips/sensor"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectLogCustomField API operation for FortiManager creates a new Custom Field.
// Returns the index value of the Custom Field and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object log - custom-field chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectLogCustomField(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/log/custom-field"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectLogCustomField API operation for FortiManager updates the specified Custom Field.
// Returns the index value of the Custom Field and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object log - custom-field chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectLogCustomField(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/log/custom-field"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectLogCustomField API operation for FortiManager deletes the specified Custom Field.
// Returns error for service API and SDK errors.
// See the object log - custom-field chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectLogCustomField(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/log/custom-field"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectLogCustomField API operation for FortiManager gets the Custom Field
// with the specified index value.
// Returns the requested Custom Field value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object log - custom-field chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectLogCustomField(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/log/custom-field"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectUserTacacs API operation for FortiManager creates a new Tacacs+.
// Returns the index value of the Tacacs+ and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - tacacs+ chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectUserTacacs(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/tacacs+"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectUserTacacs API operation for FortiManager updates the specified Tacacs+.
// Returns the index value of the Tacacs+ and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - tacacs+ chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectUserTacacs(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/tacacs+"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectUserTacacs API operation for FortiManager deletes the specified Tacacs+.
// Returns error for service API and SDK errors.
// See the object user - tacacs+ chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectUserTacacs(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/user/tacacs+"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectUserTacacs API operation for FortiManager gets the Tacacs+
// with the specified index value.
// Returns the requested Tacacs+ value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - tacacs+ chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectUserTacacs(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/tacacs+"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectFirewallLdbMonitor API operation for FortiManager creates a new Ldb Monitor.
// Returns the index value of the Ldb Monitor and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - ldb-monitor chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectFirewallLdbMonitor(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/ldb-monitor"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectFirewallLdbMonitor API operation for FortiManager updates the specified Ldb Monitor.
// Returns the index value of the Ldb Monitor and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - ldb-monitor chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectFirewallLdbMonitor(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/ldb-monitor"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectFirewallLdbMonitor API operation for FortiManager deletes the specified Ldb Monitor.
// Returns error for service API and SDK errors.
// See the object firewall - ldb-monitor chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectFirewallLdbMonitor(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/firewall/ldb-monitor"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectFirewallLdbMonitor API operation for FortiManager gets the Ldb Monitor
// with the specified index value.
// Returns the requested Ldb Monitor value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - ldb-monitor chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectFirewallLdbMonitor(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/ldb-monitor"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectApplicationList API operation for FortiManager creates a new List.
// Returns the index value of the List and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object application - list chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectApplicationList(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/application/list"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectApplicationList API operation for FortiManager updates the specified List.
// Returns the index value of the List and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object application - list chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectApplicationList(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/application/list"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectApplicationList API operation for FortiManager deletes the specified List.
// Returns error for service API and SDK errors.
// See the object application - list chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectApplicationList(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/application/list"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectApplicationList API operation for FortiManager gets the List
// with the specified index value.
// Returns the requested List value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object application - list chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectApplicationList(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/application/list"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectDlpSensor API operation for FortiManager creates a new Sensor.
// Returns the index value of the Sensor and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object dlp - sensor chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectDlpSensor(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/dlp/sensor"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectDlpSensor API operation for FortiManager updates the specified Sensor.
// Returns the index value of the Sensor and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object dlp - sensor chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectDlpSensor(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/dlp/sensor"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectDlpSensor API operation for FortiManager deletes the specified Sensor.
// Returns error for service API and SDK errors.
// See the object dlp - sensor chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectDlpSensor(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/dlp/sensor"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectDlpSensor API operation for FortiManager gets the Sensor
// with the specified index value.
// Returns the requested Sensor value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object dlp - sensor chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectDlpSensor(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/dlp/sensor"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectWanoptPeer API operation for FortiManager creates a new Peer.
// Returns the index value of the Peer and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wanopt - peer chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectWanoptPeer(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wanopt/peer"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectWanoptPeer API operation for FortiManager updates the specified Peer.
// Returns the index value of the Peer and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wanopt - peer chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectWanoptPeer(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wanopt/peer"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectWanoptPeer API operation for FortiManager deletes the specified Peer.
// Returns error for service API and SDK errors.
// See the object wanopt - peer chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectWanoptPeer(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/wanopt/peer"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectWanoptPeer API operation for FortiManager gets the Peer
// with the specified index value.
// Returns the requested Peer value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wanopt - peer chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectWanoptPeer(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wanopt/peer"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectWanoptAuthGroup API operation for FortiManager creates a new Auth Group.
// Returns the index value of the Auth Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wanopt - auth-group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectWanoptAuthGroup(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wanopt/auth-group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectWanoptAuthGroup API operation for FortiManager updates the specified Auth Group.
// Returns the index value of the Auth Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wanopt - auth-group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectWanoptAuthGroup(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wanopt/auth-group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectWanoptAuthGroup API operation for FortiManager deletes the specified Auth Group.
// Returns error for service API and SDK errors.
// See the object wanopt - auth-group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectWanoptAuthGroup(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/wanopt/auth-group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectWanoptAuthGroup API operation for FortiManager gets the Auth Group
// with the specified index value.
// Returns the requested Auth Group value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wanopt - auth-group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectWanoptAuthGroup(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wanopt/auth-group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectVpnSslWebPortal API operation for FortiManager creates a new SslWebPortal.
// Returns the index value of the SslWebPortal and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object vpn - ssl web portal chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectVpnSslWebPortal(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/vpn/ssl/web/portal"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectVpnSslWebPortal API operation for FortiManager updates the specified SslWebPortal.
// Returns the index value of the SslWebPortal and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object vpn - ssl web portal chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectVpnSslWebPortal(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/vpn/ssl/web/portal"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectVpnSslWebPortal API operation for FortiManager deletes the specified SslWebPortal.
// Returns error for service API and SDK errors.
// See the object vpn - ssl web portal chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectVpnSslWebPortal(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/vpn/ssl/web/portal"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectVpnSslWebPortal API operation for FortiManager gets the SslWebPortal
// with the specified index value.
// Returns the requested SslWebPortal value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object vpn - ssl web portal chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectVpnSslWebPortal(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/vpn/ssl/web/portal"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectSystemReplacemsgImage API operation for FortiManager creates a new Replacemsg Image.
// Returns the index value of the Replacemsg Image and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object system - replacemsg-image chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectSystemReplacemsgImage(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/system/replacemsg-image"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectSystemReplacemsgImage API operation for FortiManager updates the specified Replacemsg Image.
// Returns the index value of the Replacemsg Image and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object system - replacemsg-image chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectSystemReplacemsgImage(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/system/replacemsg-image"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectSystemReplacemsgImage API operation for FortiManager deletes the specified Replacemsg Image.
// Returns error for service API and SDK errors.
// See the object system - replacemsg-image chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectSystemReplacemsgImage(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/system/replacemsg-image"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectSystemReplacemsgImage API operation for FortiManager gets the Replacemsg Image
// with the specified index value.
// Returns the requested Replacemsg Image value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object system - replacemsg-image chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectSystemReplacemsgImage(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/system/replacemsg-image"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectSystemReplacemsgGroup API operation for FortiManager creates a new Replacemsg Group.
// Returns the index value of the Replacemsg Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object system - replacemsg-group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectSystemReplacemsgGroup(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/system/replacemsg-group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectSystemReplacemsgGroup API operation for FortiManager updates the specified Replacemsg Group.
// Returns the index value of the Replacemsg Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object system - replacemsg-group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectSystemReplacemsgGroup(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/system/replacemsg-group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectSystemReplacemsgGroup API operation for FortiManager deletes the specified Replacemsg Group.
// Returns error for service API and SDK errors.
// See the object system - replacemsg-group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectSystemReplacemsgGroup(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/system/replacemsg-group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectSystemReplacemsgGroup API operation for FortiManager gets the Replacemsg Group
// with the specified index value.
// Returns the requested Replacemsg Group value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object system - replacemsg-group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectSystemReplacemsgGroup(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/system/replacemsg-group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectFirewallGtp API operation for FortiManager creates a new Gtp.
// Returns the index value of the Gtp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - gtp chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectFirewallGtp(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/gtp"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectFirewallGtp API operation for FortiManager updates the specified Gtp.
// Returns the index value of the Gtp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - gtp chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectFirewallGtp(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/gtp"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectFirewallGtp API operation for FortiManager deletes the specified Gtp.
// Returns error for service API and SDK errors.
// See the object firewall - gtp chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectFirewallGtp(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/firewall/gtp"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectFirewallGtp API operation for FortiManager gets the Gtp
// with the specified index value.
// Returns the requested Gtp value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - gtp chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectFirewallGtp(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/gtp"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectWebfilterContent API operation for FortiManager creates a new Content.
// Returns the index value of the Content and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object webfilter - content chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectWebfilterContent(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/webfilter/content"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectWebfilterContent API operation for FortiManager updates the specified Content.
// Returns the index value of the Content and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object webfilter - content chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectWebfilterContent(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/webfilter/content"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectWebfilterContent API operation for FortiManager deletes the specified Content.
// Returns error for service API and SDK errors.
// See the object webfilter - content chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectWebfilterContent(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/webfilter/content"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectWebfilterContent API operation for FortiManager gets the Content
// with the specified index value.
// Returns the requested Content value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object webfilter - content chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectWebfilterContent(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/webfilter/content"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectWebfilterContentHeader API operation for FortiManager creates a new Content Header.
// Returns the index value of the Content Header and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object webfilter - content-header chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectWebfilterContentHeader(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/webfilter/content-header"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectWebfilterContentHeader API operation for FortiManager updates the specified Content Header.
// Returns the index value of the Content Header and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object webfilter - content-header chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectWebfilterContentHeader(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/webfilter/content-header"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectWebfilterContentHeader API operation for FortiManager deletes the specified Content Header.
// Returns error for service API and SDK errors.
// See the object webfilter - content-header chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectWebfilterContentHeader(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/webfilter/content-header"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectWebfilterContentHeader API operation for FortiManager gets the Content Header
// with the specified index value.
// Returns the requested Content Header value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object webfilter - content-header chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectWebfilterContentHeader(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/webfilter/content-header"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectFirewallScheduleGroup API operation for FortiManager creates a new ScheduleGroup.
// Returns the index value of the ScheduleGroup and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - schedule group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectFirewallScheduleGroup(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/schedule/group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectFirewallScheduleGroup API operation for FortiManager updates the specified ScheduleGroup.
// Returns the index value of the ScheduleGroup and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - schedule group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectFirewallScheduleGroup(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/schedule/group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectFirewallScheduleGroup API operation for FortiManager deletes the specified ScheduleGroup.
// Returns error for service API and SDK errors.
// See the object firewall - schedule group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectFirewallScheduleGroup(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/firewall/schedule/group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectFirewallScheduleGroup API operation for FortiManager gets the ScheduleGroup
// with the specified index value.
// Returns the requested ScheduleGroup value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - schedule group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectFirewallScheduleGroup(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/schedule/group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectFirewallShaperTrafficShaper API operation for FortiManager creates a new ShaperTraffic Shaper.
// Returns the index value of the ShaperTraffic Shaper and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - shaper traffic-shaper chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectFirewallShaperTrafficShaper(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/shaper/traffic-shaper"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectFirewallShaperTrafficShaper API operation for FortiManager updates the specified ShaperTraffic Shaper.
// Returns the index value of the ShaperTraffic Shaper and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - shaper traffic-shaper chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectFirewallShaperTrafficShaper(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/shaper/traffic-shaper"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectFirewallShaperTrafficShaper API operation for FortiManager deletes the specified ShaperTraffic Shaper.
// Returns error for service API and SDK errors.
// See the object firewall - shaper traffic-shaper chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectFirewallShaperTrafficShaper(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/firewall/shaper/traffic-shaper"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectFirewallShaperTrafficShaper API operation for FortiManager gets the ShaperTraffic Shaper
// with the specified index value.
// Returns the requested ShaperTraffic Shaper value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - shaper traffic-shaper chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectFirewallShaperTrafficShaper(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/shaper/traffic-shaper"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectFirewallShaperPerIpShaper API operation for FortiManager creates a new ShaperPer Ip Shaper.
// Returns the index value of the ShaperPer Ip Shaper and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - shaper per-ip-shaper chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectFirewallShaperPerIpShaper(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/shaper/per-ip-shaper"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectFirewallShaperPerIpShaper API operation for FortiManager updates the specified ShaperPer Ip Shaper.
// Returns the index value of the ShaperPer Ip Shaper and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - shaper per-ip-shaper chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectFirewallShaperPerIpShaper(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/shaper/per-ip-shaper"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectFirewallShaperPerIpShaper API operation for FortiManager deletes the specified ShaperPer Ip Shaper.
// Returns error for service API and SDK errors.
// See the object firewall - shaper per-ip-shaper chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectFirewallShaperPerIpShaper(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/firewall/shaper/per-ip-shaper"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectFirewallShaperPerIpShaper API operation for FortiManager gets the ShaperPer Ip Shaper
// with the specified index value.
// Returns the requested ShaperPer Ip Shaper value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - shaper per-ip-shaper chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectFirewallShaperPerIpShaper(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/shaper/per-ip-shaper"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectVpnSslWebHostCheckSoftware API operation for FortiManager creates a new SslWebHost Check Software.
// Returns the index value of the SslWebHost Check Software and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object vpn - ssl web host-check-software chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectVpnSslWebHostCheckSoftware(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/vpn/ssl/web/host-check-software"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectVpnSslWebHostCheckSoftware API operation for FortiManager updates the specified SslWebHost Check Software.
// Returns the index value of the SslWebHost Check Software and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object vpn - ssl web host-check-software chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectVpnSslWebHostCheckSoftware(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/vpn/ssl/web/host-check-software"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectVpnSslWebHostCheckSoftware API operation for FortiManager deletes the specified SslWebHost Check Software.
// Returns error for service API and SDK errors.
// See the object vpn - ssl web host-check-software chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectVpnSslWebHostCheckSoftware(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/vpn/ssl/web/host-check-software"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectVpnSslWebHostCheckSoftware API operation for FortiManager gets the SslWebHost Check Software
// with the specified index value.
// Returns the requested SslWebHost Check Software value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object vpn - ssl web host-check-software chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectVpnSslWebHostCheckSoftware(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/vpn/ssl/web/host-check-software"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectWirelessControllerVap API operation for FortiManager creates a new Vap.
// Returns the index value of the Vap and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - vap chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectWirelessControllerVap(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/vap"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectWirelessControllerVap API operation for FortiManager updates the specified Vap.
// Returns the index value of the Vap and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - vap chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectWirelessControllerVap(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/vap"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectWirelessControllerVap API operation for FortiManager deletes the specified Vap.
// Returns error for service API and SDK errors.
// See the object wireless-controller - vap chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectWirelessControllerVap(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/wireless-controller/vap"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectWirelessControllerVap API operation for FortiManager gets the Vap
// with the specified index value.
// Returns the requested Vap value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - vap chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectWirelessControllerVap(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/vap"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectWirelessControllerVapGroup API operation for FortiManager creates a new Vap Group.
// Returns the index value of the Vap Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - vap-group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectWirelessControllerVapGroup(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/vap-group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectWirelessControllerVapGroup API operation for FortiManager updates the specified Vap Group.
// Returns the index value of the Vap Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - vap-group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectWirelessControllerVapGroup(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/vap-group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectWirelessControllerVapGroup API operation for FortiManager deletes the specified Vap Group.
// Returns error for service API and SDK errors.
// See the object wireless-controller - vap-group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectWirelessControllerVapGroup(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/wireless-controller/vap-group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectWirelessControllerVapGroup API operation for FortiManager gets the Vap Group
// with the specified index value.
// Returns the requested Vap Group value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - vap-group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectWirelessControllerVapGroup(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/vap-group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectWebfilterProfile API operation for FortiManager creates a new Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object webfilter - profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectWebfilterProfile(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/webfilter/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectWebfilterProfile API operation for FortiManager updates the specified Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object webfilter - profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectWebfilterProfile(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/webfilter/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectWebfilterProfile API operation for FortiManager deletes the specified Profile.
// Returns error for service API and SDK errors.
// See the object webfilter - profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectWebfilterProfile(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/webfilter/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectWebfilterProfile API operation for FortiManager gets the Profile
// with the specified index value.
// Returns the requested Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object webfilter - profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectWebfilterProfile(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/webfilter/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectAntivirusProfile API operation for FortiManager creates a new Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object antivirus - profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectAntivirusProfile(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/antivirus/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectAntivirusProfile API operation for FortiManager updates the specified Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object antivirus - profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectAntivirusProfile(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/antivirus/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectAntivirusProfile API operation for FortiManager deletes the specified Profile.
// Returns error for service API and SDK errors.
// See the object antivirus - profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectAntivirusProfile(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/antivirus/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectAntivirusProfile API operation for FortiManager gets the Profile
// with the specified index value.
// Returns the requested Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object antivirus - profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectAntivirusProfile(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/antivirus/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectFirewallProfileProtocolOptions API operation for FortiManager creates a new Profile Protocol Options.
// Returns the index value of the Profile Protocol Options and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - profile-protocol-options chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectFirewallProfileProtocolOptions(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/profile-protocol-options"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectFirewallProfileProtocolOptions API operation for FortiManager updates the specified Profile Protocol Options.
// Returns the index value of the Profile Protocol Options and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - profile-protocol-options chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectFirewallProfileProtocolOptions(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/profile-protocol-options"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectFirewallProfileProtocolOptions API operation for FortiManager deletes the specified Profile Protocol Options.
// Returns error for service API and SDK errors.
// See the object firewall - profile-protocol-options chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectFirewallProfileProtocolOptions(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/firewall/profile-protocol-options"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectFirewallProfileProtocolOptions API operation for FortiManager gets the Profile Protocol Options
// with the specified index value.
// Returns the requested Profile Protocol Options value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - profile-protocol-options chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectFirewallProfileProtocolOptions(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/profile-protocol-options"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectFirewallProfileGroup API operation for FortiManager creates a new Profile Group.
// Returns the index value of the Profile Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - profile-group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectFirewallProfileGroup(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/profile-group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectFirewallProfileGroup API operation for FortiManager updates the specified Profile Group.
// Returns the index value of the Profile Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - profile-group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectFirewallProfileGroup(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/profile-group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectFirewallProfileGroup API operation for FortiManager deletes the specified Profile Group.
// Returns error for service API and SDK errors.
// See the object firewall - profile-group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectFirewallProfileGroup(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/firewall/profile-group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectFirewallProfileGroup API operation for FortiManager gets the Profile Group
// with the specified index value.
// Returns the requested Profile Group value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - profile-group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectFirewallProfileGroup(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/profile-group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectVoipProfile API operation for FortiManager creates a new Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object voip - profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectVoipProfile(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/voip/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectVoipProfile API operation for FortiManager updates the specified Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object voip - profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectVoipProfile(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/voip/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectVoipProfile API operation for FortiManager deletes the specified Profile.
// Returns error for service API and SDK errors.
// See the object voip - profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectVoipProfile(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/voip/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectVoipProfile API operation for FortiManager gets the Profile
// with the specified index value.
// Returns the requested Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object voip - profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectVoipProfile(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/voip/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectUserFortitoken API operation for FortiManager creates a new Fortitoken.
// Returns the index value of the Fortitoken and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - fortitoken chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectUserFortitoken(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/fortitoken"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectUserFortitoken API operation for FortiManager updates the specified Fortitoken.
// Returns the index value of the Fortitoken and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - fortitoken chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectUserFortitoken(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/fortitoken"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectUserFortitoken API operation for FortiManager deletes the specified Fortitoken.
// Returns error for service API and SDK errors.
// See the object user - fortitoken chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectUserFortitoken(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/user/fortitoken"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectUserFortitoken API operation for FortiManager gets the Fortitoken
// with the specified index value.
// Returns the requested Fortitoken value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - fortitoken chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectUserFortitoken(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/fortitoken"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectWebProxyForwardServer API operation for FortiManager creates a new Forward Server.
// Returns the index value of the Forward Server and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object web-proxy - forward-server chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectWebProxyForwardServer(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/web-proxy/forward-server"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectWebProxyForwardServer API operation for FortiManager updates the specified Forward Server.
// Returns the index value of the Forward Server and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object web-proxy - forward-server chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectWebProxyForwardServer(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/web-proxy/forward-server"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectWebProxyForwardServer API operation for FortiManager deletes the specified Forward Server.
// Returns error for service API and SDK errors.
// See the object web-proxy - forward-server chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectWebProxyForwardServer(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/web-proxy/forward-server"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectWebProxyForwardServer API operation for FortiManager gets the Forward Server
// with the specified index value.
// Returns the requested Forward Server value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object web-proxy - forward-server chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectWebProxyForwardServer(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/web-proxy/forward-server"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectWirelessControllerWtpProfile API operation for FortiManager creates a new Wtp Profile.
// Returns the index value of the Wtp Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - wtp-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectWirelessControllerWtpProfile(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/wtp-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectWirelessControllerWtpProfile API operation for FortiManager updates the specified Wtp Profile.
// Returns the index value of the Wtp Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - wtp-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectWirelessControllerWtpProfile(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/wtp-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectWirelessControllerWtpProfile API operation for FortiManager deletes the specified Wtp Profile.
// Returns error for service API and SDK errors.
// See the object wireless-controller - wtp-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectWirelessControllerWtpProfile(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/wireless-controller/wtp-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectWirelessControllerWtpProfile API operation for FortiManager gets the Wtp Profile
// with the specified index value.
// Returns the requested Wtp Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - wtp-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectWirelessControllerWtpProfile(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/wtp-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectDlpFilepattern API operation for FortiManager creates a new Filepattern.
// Returns the index value of the Filepattern and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object dlp - filepattern chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectDlpFilepattern(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/dlp/filepattern"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectDlpFilepattern API operation for FortiManager updates the specified Filepattern.
// Returns the index value of the Filepattern and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object dlp - filepattern chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectDlpFilepattern(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/dlp/filepattern"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectDlpFilepattern API operation for FortiManager deletes the specified Filepattern.
// Returns error for service API and SDK errors.
// See the object dlp - filepattern chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectDlpFilepattern(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/dlp/filepattern"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectDlpFilepattern API operation for FortiManager gets the Filepattern
// with the specified index value.
// Returns the requested Filepattern value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object dlp - filepattern chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectDlpFilepattern(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/dlp/filepattern"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectIcapServer API operation for FortiManager creates a new Server.
// Returns the index value of the Server and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object icap - server chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectIcapServer(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/icap/server"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectIcapServer API operation for FortiManager updates the specified Server.
// Returns the index value of the Server and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object icap - server chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectIcapServer(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/icap/server"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectIcapServer API operation for FortiManager deletes the specified Server.
// Returns error for service API and SDK errors.
// See the object icap - server chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectIcapServer(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/icap/server"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectIcapServer API operation for FortiManager gets the Server
// with the specified index value.
// Returns the requested Server value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object icap - server chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectIcapServer(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/icap/server"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectIcapProfile API operation for FortiManager creates a new Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object icap - profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectIcapProfile(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/icap/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectIcapProfile API operation for FortiManager updates the specified Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object icap - profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectIcapProfile(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/icap/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectIcapProfile API operation for FortiManager deletes the specified Profile.
// Returns error for service API and SDK errors.
// See the object icap - profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectIcapProfile(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/icap/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectIcapProfile API operation for FortiManager gets the Profile
// with the specified index value.
// Returns the requested Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object icap - profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectIcapProfile(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/icap/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectUserFsso API operation for FortiManager creates a new Fsso.
// Returns the index value of the Fsso and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - fsso chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectUserFsso(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/fsso"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectUserFsso API operation for FortiManager updates the specified Fsso.
// Returns the index value of the Fsso and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - fsso chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectUserFsso(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/fsso"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectUserFsso API operation for FortiManager deletes the specified Fsso.
// Returns error for service API and SDK errors.
// See the object user - fsso chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectUserFsso(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/user/fsso"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectUserFsso API operation for FortiManager gets the Fsso
// with the specified index value.
// Returns the requested Fsso value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - fsso chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectUserFsso(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/fsso"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectSystemSmsServer API operation for FortiManager creates a new Sms Server.
// Returns the index value of the Sms Server and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object system - sms-server chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectSystemSmsServer(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/system/sms-server"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectSystemSmsServer API operation for FortiManager updates the specified Sms Server.
// Returns the index value of the Sms Server and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object system - sms-server chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectSystemSmsServer(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/system/sms-server"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectSystemSmsServer API operation for FortiManager deletes the specified Sms Server.
// Returns error for service API and SDK errors.
// See the object system - sms-server chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectSystemSmsServer(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/system/sms-server"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectSystemSmsServer API operation for FortiManager gets the Sms Server
// with the specified index value.
// Returns the requested Sms Server value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object system - sms-server chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectSystemSmsServer(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/system/sms-server"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectVpnCertificateOcspServer API operation for FortiManager creates a new CertificateOcsp Server.
// Returns the index value of the CertificateOcsp Server and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object vpn - certificate ocsp-server chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectVpnCertificateOcspServer(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/vpn/certificate/ocsp-server"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectVpnCertificateOcspServer API operation for FortiManager updates the specified CertificateOcsp Server.
// Returns the index value of the CertificateOcsp Server and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object vpn - certificate ocsp-server chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectVpnCertificateOcspServer(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/vpn/certificate/ocsp-server"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectVpnCertificateOcspServer API operation for FortiManager deletes the specified CertificateOcsp Server.
// Returns error for service API and SDK errors.
// See the object vpn - certificate ocsp-server chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectVpnCertificateOcspServer(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/vpn/certificate/ocsp-server"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectVpnCertificateOcspServer API operation for FortiManager gets the CertificateOcsp Server
// with the specified index value.
// Returns the requested CertificateOcsp Server value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object vpn - certificate ocsp-server chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectVpnCertificateOcspServer(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/vpn/certificate/ocsp-server"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectUserPasswordPolicy API operation for FortiManager creates a new Password Policy.
// Returns the index value of the Password Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - password-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectUserPasswordPolicy(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/password-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectUserPasswordPolicy API operation for FortiManager updates the specified Password Policy.
// Returns the index value of the Password Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - password-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectUserPasswordPolicy(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/password-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectUserPasswordPolicy API operation for FortiManager deletes the specified Password Policy.
// Returns error for service API and SDK errors.
// See the object user - password-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectUserPasswordPolicy(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/user/password-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectUserPasswordPolicy API operation for FortiManager gets the Password Policy
// with the specified index value.
// Returns the requested Password Policy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - password-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectUserPasswordPolicy(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/password-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectUserFssoPolling API operation for FortiManager creates a new Fsso Polling.
// Returns the index value of the Fsso Polling and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - fsso-polling chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectUserFssoPolling(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/fsso-polling"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectUserFssoPolling API operation for FortiManager updates the specified Fsso Polling.
// Returns the index value of the Fsso Polling and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - fsso-polling chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectUserFssoPolling(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/fsso-polling"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectUserFssoPolling API operation for FortiManager deletes the specified Fsso Polling.
// Returns error for service API and SDK errors.
// See the object user - fsso-polling chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectUserFssoPolling(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/user/fsso-polling"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectUserFssoPolling API operation for FortiManager gets the Fsso Polling
// with the specified index value.
// Returns the requested Fsso Polling value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - fsso-polling chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectUserFssoPolling(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/fsso-polling"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectFirewallIppool6 API operation for FortiManager creates a new Ippool6.
// Returns the index value of the Ippool6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - ippool6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectFirewallIppool6(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/ippool6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectFirewallIppool6 API operation for FortiManager updates the specified Ippool6.
// Returns the index value of the Ippool6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - ippool6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectFirewallIppool6(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/ippool6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectFirewallIppool6 API operation for FortiManager deletes the specified Ippool6.
// Returns error for service API and SDK errors.
// See the object firewall - ippool6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectFirewallIppool6(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/firewall/ippool6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectFirewallIppool6 API operation for FortiManager gets the Ippool6
// with the specified index value.
// Returns the requested Ippool6 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - ippool6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectFirewallIppool6(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/ippool6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectFirewallVip6 API operation for FortiManager creates a new Vip6.
// Returns the index value of the Vip6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - vip6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectFirewallVip6(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/vip6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectFirewallVip6 API operation for FortiManager updates the specified Vip6.
// Returns the index value of the Vip6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - vip6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectFirewallVip6(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/vip6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectFirewallVip6 API operation for FortiManager deletes the specified Vip6.
// Returns error for service API and SDK errors.
// See the object firewall - vip6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectFirewallVip6(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/firewall/vip6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectFirewallVip6 API operation for FortiManager gets the Vip6
// with the specified index value.
// Returns the requested Vip6 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - vip6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectFirewallVip6(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/vip6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectFirewallVipgrp6 API operation for FortiManager creates a new Vipgrp6.
// Returns the index value of the Vipgrp6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - vipgrp6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectFirewallVipgrp6(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/vipgrp6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectFirewallVipgrp6 API operation for FortiManager updates the specified Vipgrp6.
// Returns the index value of the Vipgrp6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - vipgrp6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectFirewallVipgrp6(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/vipgrp6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectFirewallVipgrp6 API operation for FortiManager deletes the specified Vipgrp6.
// Returns error for service API and SDK errors.
// See the object firewall - vipgrp6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectFirewallVipgrp6(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/firewall/vipgrp6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectFirewallVipgrp6 API operation for FortiManager gets the Vipgrp6
// with the specified index value.
// Returns the requested Vipgrp6 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - vipgrp6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectFirewallVipgrp6(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/vipgrp6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectFirewallIdentityBasedRoute API operation for FortiManager creates a new Identity Based Route.
// Returns the index value of the Identity Based Route and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - identity-based-route chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectFirewallIdentityBasedRoute(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/identity-based-route"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectFirewallIdentityBasedRoute API operation for FortiManager updates the specified Identity Based Route.
// Returns the index value of the Identity Based Route and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - identity-based-route chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectFirewallIdentityBasedRoute(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/identity-based-route"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectFirewallIdentityBasedRoute API operation for FortiManager deletes the specified Identity Based Route.
// Returns error for service API and SDK errors.
// See the object firewall - identity-based-route chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectFirewallIdentityBasedRoute(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/firewall/identity-based-route"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectFirewallIdentityBasedRoute API operation for FortiManager gets the Identity Based Route
// with the specified index value.
// Returns the requested Identity Based Route value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - identity-based-route chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectFirewallIdentityBasedRoute(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/identity-based-route"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectWirelessControllerWidsProfile API operation for FortiManager creates a new Wids Profile.
// Returns the index value of the Wids Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - wids-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectWirelessControllerWidsProfile(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/wids-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectWirelessControllerWidsProfile API operation for FortiManager updates the specified Wids Profile.
// Returns the index value of the Wids Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - wids-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectWirelessControllerWidsProfile(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/wids-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectWirelessControllerWidsProfile API operation for FortiManager deletes the specified Wids Profile.
// Returns error for service API and SDK errors.
// See the object wireless-controller - wids-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectWirelessControllerWidsProfile(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/wireless-controller/wids-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectWirelessControllerWidsProfile API operation for FortiManager gets the Wids Profile
// with the specified index value.
// Returns the requested Wids Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - wids-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectWirelessControllerWidsProfile(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/wids-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectSwitchControllerManagedSwitch API operation for FortiManager creates a new Managed Switch.
// Returns the index value of the Managed Switch and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object switch-controller - managed-switch chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectSwitchControllerManagedSwitch(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/switch-controller/managed-switch"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectSwitchControllerManagedSwitch API operation for FortiManager updates the specified Managed Switch.
// Returns the index value of the Managed Switch and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object switch-controller - managed-switch chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectSwitchControllerManagedSwitch(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/switch-controller/managed-switch"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectSwitchControllerManagedSwitch API operation for FortiManager deletes the specified Managed Switch.
// Returns error for service API and SDK errors.
// See the object switch-controller - managed-switch chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectSwitchControllerManagedSwitch(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/switch-controller/managed-switch"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectSwitchControllerManagedSwitch API operation for FortiManager gets the Managed Switch
// with the specified index value.
// Returns the requested Managed Switch value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object switch-controller - managed-switch chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectSwitchControllerManagedSwitch(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/switch-controller/managed-switch"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectWanoptProfile API operation for FortiManager creates a new Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wanopt - profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectWanoptProfile(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wanopt/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectWanoptProfile API operation for FortiManager updates the specified Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wanopt - profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectWanoptProfile(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wanopt/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectWanoptProfile API operation for FortiManager deletes the specified Profile.
// Returns error for service API and SDK errors.
// See the object wanopt - profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectWanoptProfile(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/wanopt/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectWanoptProfile API operation for FortiManager gets the Profile
// with the specified index value.
// Returns the requested Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wanopt - profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectWanoptProfile(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wanopt/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectFirewallMulticastAddress API operation for FortiManager creates a new Multicast Address.
// Returns the index value of the Multicast Address and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - multicast-address chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectFirewallMulticastAddress(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/multicast-address"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectFirewallMulticastAddress API operation for FortiManager updates the specified Multicast Address.
// Returns the index value of the Multicast Address and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - multicast-address chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectFirewallMulticastAddress(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/multicast-address"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectFirewallMulticastAddress API operation for FortiManager deletes the specified Multicast Address.
// Returns error for service API and SDK errors.
// See the object firewall - multicast-address chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectFirewallMulticastAddress(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/firewall/multicast-address"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectFirewallMulticastAddress API operation for FortiManager gets the Multicast Address
// with the specified index value.
// Returns the requested Multicast Address value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - multicast-address chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectFirewallMulticastAddress(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/multicast-address"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectFirewallServiceCategory API operation for FortiManager creates a new ServiceCategory.
// Returns the index value of the ServiceCategory and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - service category chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectFirewallServiceCategory(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/service/category"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectFirewallServiceCategory API operation for FortiManager updates the specified ServiceCategory.
// Returns the index value of the ServiceCategory and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - service category chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectFirewallServiceCategory(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/service/category"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectFirewallServiceCategory API operation for FortiManager deletes the specified ServiceCategory.
// Returns error for service API and SDK errors.
// See the object firewall - service category chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectFirewallServiceCategory(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/firewall/service/category"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectFirewallServiceCategory API operation for FortiManager gets the ServiceCategory
// with the specified index value.
// Returns the requested ServiceCategory value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - service category chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectFirewallServiceCategory(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/service/category"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectSystemGeoipOverride API operation for FortiManager creates a new Geoip Override.
// Returns the index value of the Geoip Override and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object system - geoip-override chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectSystemGeoipOverride(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/system/geoip-override"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectSystemGeoipOverride API operation for FortiManager updates the specified Geoip Override.
// Returns the index value of the Geoip Override and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object system - geoip-override chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectSystemGeoipOverride(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/system/geoip-override"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectSystemGeoipOverride API operation for FortiManager deletes the specified Geoip Override.
// Returns error for service API and SDK errors.
// See the object system - geoip-override chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectSystemGeoipOverride(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/system/geoip-override"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectSystemGeoipOverride API operation for FortiManager gets the Geoip Override
// with the specified index value.
// Returns the requested Geoip Override value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object system - geoip-override chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectSystemGeoipOverride(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/system/geoip-override"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectApplicationCustom API operation for FortiManager creates a new Custom.
// Returns the index value of the Custom and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object application - custom chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectApplicationCustom(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/application/custom"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectApplicationCustom API operation for FortiManager updates the specified Custom.
// Returns the index value of the Custom and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object application - custom chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectApplicationCustom(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/application/custom"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectApplicationCustom API operation for FortiManager deletes the specified Custom.
// Returns error for service API and SDK errors.
// See the object application - custom chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectApplicationCustom(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/application/custom"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectApplicationCustom API operation for FortiManager gets the Custom
// with the specified index value.
// Returns the requested Custom value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object application - custom chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectApplicationCustom(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/application/custom"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectVpnSslWebRealm API operation for FortiManager creates a new SslWebRealm.
// Returns the index value of the SslWebRealm and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object vpn - ssl web realm chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectVpnSslWebRealm(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/vpn/ssl/web/realm"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectVpnSslWebRealm API operation for FortiManager updates the specified SslWebRealm.
// Returns the index value of the SslWebRealm and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object vpn - ssl web realm chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectVpnSslWebRealm(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/vpn/ssl/web/realm"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectVpnSslWebRealm API operation for FortiManager deletes the specified SslWebRealm.
// Returns error for service API and SDK errors.
// See the object vpn - ssl web realm chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectVpnSslWebRealm(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/vpn/ssl/web/realm"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectVpnSslWebRealm API operation for FortiManager gets the SslWebRealm
// with the specified index value.
// Returns the requested SslWebRealm value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object vpn - ssl web realm chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectVpnSslWebRealm(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/vpn/ssl/web/realm"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectFirewallVip64 API operation for FortiManager creates a new Vip64.
// Returns the index value of the Vip64 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - vip64 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectFirewallVip64(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/vip64"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectFirewallVip64 API operation for FortiManager updates the specified Vip64.
// Returns the index value of the Vip64 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - vip64 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectFirewallVip64(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/vip64"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectFirewallVip64 API operation for FortiManager deletes the specified Vip64.
// Returns error for service API and SDK errors.
// See the object firewall - vip64 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectFirewallVip64(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/firewall/vip64"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectFirewallVip64 API operation for FortiManager gets the Vip64
// with the specified index value.
// Returns the requested Vip64 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - vip64 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectFirewallVip64(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/vip64"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectFirewallVip46 API operation for FortiManager creates a new Vip46.
// Returns the index value of the Vip46 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - vip46 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectFirewallVip46(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/vip46"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectFirewallVip46 API operation for FortiManager updates the specified Vip46.
// Returns the index value of the Vip46 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - vip46 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectFirewallVip46(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/vip46"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectFirewallVip46 API operation for FortiManager deletes the specified Vip46.
// Returns error for service API and SDK errors.
// See the object firewall - vip46 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectFirewallVip46(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/firewall/vip46"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectFirewallVip46 API operation for FortiManager gets the Vip46
// with the specified index value.
// Returns the requested Vip46 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - vip46 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectFirewallVip46(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/vip46"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectFirewallVipgrp64 API operation for FortiManager creates a new Vipgrp64.
// Returns the index value of the Vipgrp64 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - vipgrp64 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectFirewallVipgrp64(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/vipgrp64"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectFirewallVipgrp64 API operation for FortiManager updates the specified Vipgrp64.
// Returns the index value of the Vipgrp64 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - vipgrp64 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectFirewallVipgrp64(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/vipgrp64"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectFirewallVipgrp64 API operation for FortiManager deletes the specified Vipgrp64.
// Returns error for service API and SDK errors.
// See the object firewall - vipgrp64 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectFirewallVipgrp64(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/firewall/vipgrp64"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectFirewallVipgrp64 API operation for FortiManager gets the Vipgrp64
// with the specified index value.
// Returns the requested Vipgrp64 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - vipgrp64 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectFirewallVipgrp64(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/vipgrp64"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectFirewallVipgrp46 API operation for FortiManager creates a new Vipgrp46.
// Returns the index value of the Vipgrp46 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - vipgrp46 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectFirewallVipgrp46(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/vipgrp46"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectFirewallVipgrp46 API operation for FortiManager updates the specified Vipgrp46.
// Returns the index value of the Vipgrp46 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - vipgrp46 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectFirewallVipgrp46(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/vipgrp46"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectFirewallVipgrp46 API operation for FortiManager deletes the specified Vipgrp46.
// Returns error for service API and SDK errors.
// See the object firewall - vipgrp46 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectFirewallVipgrp46(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/firewall/vipgrp46"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectFirewallVipgrp46 API operation for FortiManager gets the Vipgrp46
// with the specified index value.
// Returns the requested Vipgrp46 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - vipgrp46 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectFirewallVipgrp46(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/vipgrp46"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectWebProxyForwardServerGroup API operation for FortiManager creates a new Forward Server Group.
// Returns the index value of the Forward Server Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object web-proxy - forward-server-group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectWebProxyForwardServerGroup(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/web-proxy/forward-server-group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectWebProxyForwardServerGroup API operation for FortiManager updates the specified Forward Server Group.
// Returns the index value of the Forward Server Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object web-proxy - forward-server-group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectWebProxyForwardServerGroup(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/web-proxy/forward-server-group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectWebProxyForwardServerGroup API operation for FortiManager deletes the specified Forward Server Group.
// Returns error for service API and SDK errors.
// See the object web-proxy - forward-server-group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectWebProxyForwardServerGroup(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/web-proxy/forward-server-group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectWebProxyForwardServerGroup API operation for FortiManager gets the Forward Server Group
// with the specified index value.
// Returns the requested Forward Server Group value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object web-proxy - forward-server-group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectWebProxyForwardServerGroup(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/web-proxy/forward-server-group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectUserPop3 API operation for FortiManager creates a new Pop3.
// Returns the index value of the Pop3 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - pop3 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectUserPop3(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/pop3"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectUserPop3 API operation for FortiManager updates the specified Pop3.
// Returns the index value of the Pop3 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - pop3 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectUserPop3(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/pop3"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectUserPop3 API operation for FortiManager deletes the specified Pop3.
// Returns error for service API and SDK errors.
// See the object user - pop3 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectUserPop3(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/user/pop3"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectUserPop3 API operation for FortiManager gets the Pop3
// with the specified index value.
// Returns the requested Pop3 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - pop3 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectUserPop3(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/pop3"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectWebProxyProfile API operation for FortiManager creates a new Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object web-proxy - profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectWebProxyProfile(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/web-proxy/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectWebProxyProfile API operation for FortiManager updates the specified Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object web-proxy - profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectWebProxyProfile(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/web-proxy/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectWebProxyProfile API operation for FortiManager deletes the specified Profile.
// Returns error for service API and SDK errors.
// See the object web-proxy - profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectWebProxyProfile(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/web-proxy/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectWebProxyProfile API operation for FortiManager gets the Profile
// with the specified index value.
// Returns the requested Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object web-proxy - profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectWebProxyProfile(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/web-proxy/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectExtenderControllerSimProfile API operation for FortiManager creates a new SimProfile.
// Returns the index value of the SimProfile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object extender-controller - sim profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectExtenderControllerSimProfile(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/extender-controller/sim_profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectExtenderControllerSimProfile API operation for FortiManager updates the specified SimProfile.
// Returns the index value of the SimProfile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object extender-controller - sim profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectExtenderControllerSimProfile(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/extender-controller/sim_profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectExtenderControllerSimProfile API operation for FortiManager deletes the specified SimProfile.
// Returns error for service API and SDK errors.
// See the object extender-controller - sim profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectExtenderControllerSimProfile(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/extender-controller/sim_profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectExtenderControllerSimProfile API operation for FortiManager gets the SimProfile
// with the specified index value.
// Returns the requested SimProfile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object extender-controller - sim profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectExtenderControllerSimProfile(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/extender-controller/sim_profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectUserSecurityExemptList API operation for FortiManager creates a new Security Exempt List.
// Returns the index value of the Security Exempt List and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - security-exempt-list chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectUserSecurityExemptList(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/security-exempt-list"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectUserSecurityExemptList API operation for FortiManager updates the specified Security Exempt List.
// Returns the index value of the Security Exempt List and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - security-exempt-list chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectUserSecurityExemptList(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/security-exempt-list"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectUserSecurityExemptList API operation for FortiManager deletes the specified Security Exempt List.
// Returns error for service API and SDK errors.
// See the object user - security-exempt-list chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectUserSecurityExemptList(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/user/security-exempt-list"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectUserSecurityExemptList API operation for FortiManager gets the Security Exempt List
// with the specified index value.
// Returns the requested Security Exempt List value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - security-exempt-list chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectUserSecurityExemptList(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/security-exempt-list"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectFirewallSslSshProfile API operation for FortiManager creates a new Ssl Ssh Profile.
// Returns the index value of the Ssl Ssh Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - ssl-ssh-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectFirewallSslSshProfile(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/ssl-ssh-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectFirewallSslSshProfile API operation for FortiManager updates the specified Ssl Ssh Profile.
// Returns the index value of the Ssl Ssh Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - ssl-ssh-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectFirewallSslSshProfile(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/ssl-ssh-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectFirewallSslSshProfile API operation for FortiManager deletes the specified Ssl Ssh Profile.
// Returns error for service API and SDK errors.
// See the object firewall - ssl-ssh-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectFirewallSslSshProfile(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/firewall/ssl-ssh-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectFirewallSslSshProfile API operation for FortiManager gets the Ssl Ssh Profile
// with the specified index value.
// Returns the requested Ssl Ssh Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - ssl-ssh-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectFirewallSslSshProfile(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/ssl-ssh-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectSystemCustomLanguage API operation for FortiManager creates a new Custom Language.
// Returns the index value of the Custom Language and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object system - custom-language chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectSystemCustomLanguage(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/system/custom-language"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectSystemCustomLanguage API operation for FortiManager updates the specified Custom Language.
// Returns the index value of the Custom Language and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object system - custom-language chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectSystemCustomLanguage(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/system/custom-language"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectSystemCustomLanguage API operation for FortiManager deletes the specified Custom Language.
// Returns error for service API and SDK errors.
// See the object system - custom-language chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectSystemCustomLanguage(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/system/custom-language"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectSystemCustomLanguage API operation for FortiManager gets the Custom Language
// with the specified index value.
// Returns the requested Custom Language value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object system - custom-language chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectSystemCustomLanguage(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/system/custom-language"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectSystemGeoipCountry API operation for FortiManager creates a new Geoip Country.
// Returns the index value of the Geoip Country and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object system - geoip-country chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectSystemGeoipCountry(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/system/geoip-country"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectSystemGeoipCountry API operation for FortiManager updates the specified Geoip Country.
// Returns the index value of the Geoip Country and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object system - geoip-country chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectSystemGeoipCountry(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/system/geoip-country"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectSystemGeoipCountry API operation for FortiManager deletes the specified Geoip Country.
// Returns error for service API and SDK errors.
// See the object system - geoip-country chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectSystemGeoipCountry(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/system/geoip-country"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectSystemGeoipCountry API operation for FortiManager gets the Geoip Country
// with the specified index value.
// Returns the requested Geoip Country value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object system - geoip-country chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectSystemGeoipCountry(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/system/geoip-country"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectWebfilterCategories API operation for FortiManager creates a new Categories.
// Returns the index value of the Categories and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object webfilter - categories chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectWebfilterCategories(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/webfilter/categories"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectWebfilterCategories API operation for FortiManager updates the specified Categories.
// Returns the index value of the Categories and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object webfilter - categories chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectWebfilterCategories(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/webfilter/categories"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectWebfilterCategories API operation for FortiManager deletes the specified Categories.
// Returns error for service API and SDK errors.
// See the object webfilter - categories chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectWebfilterCategories(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/webfilter/categories"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectWebfilterCategories API operation for FortiManager gets the Categories
// with the specified index value.
// Returns the requested Categories value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object webfilter - categories chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectWebfilterCategories(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/webfilter/categories"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectSystemVirtualWirePair API operation for FortiManager creates a new Virtual Wire Pair.
// Returns the index value of the Virtual Wire Pair and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object system - virtual-wire-pair chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectSystemVirtualWirePair(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/system/virtual-wire-pair"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectSystemVirtualWirePair API operation for FortiManager updates the specified Virtual Wire Pair.
// Returns the index value of the Virtual Wire Pair and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object system - virtual-wire-pair chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectSystemVirtualWirePair(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/system/virtual-wire-pair"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectSystemVirtualWirePair API operation for FortiManager deletes the specified Virtual Wire Pair.
// Returns error for service API and SDK errors.
// See the object system - virtual-wire-pair chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectSystemVirtualWirePair(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/system/virtual-wire-pair"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectSystemVirtualWirePair API operation for FortiManager gets the Virtual Wire Pair
// with the specified index value.
// Returns the requested Virtual Wire Pair value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object system - virtual-wire-pair chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectSystemVirtualWirePair(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/system/virtual-wire-pair"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectFirewallMulticastAddress6 API operation for FortiManager creates a new Multicast Address6.
// Returns the index value of the Multicast Address6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - multicast-address6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectFirewallMulticastAddress6(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/multicast-address6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectFirewallMulticastAddress6 API operation for FortiManager updates the specified Multicast Address6.
// Returns the index value of the Multicast Address6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - multicast-address6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectFirewallMulticastAddress6(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/multicast-address6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectFirewallMulticastAddress6 API operation for FortiManager deletes the specified Multicast Address6.
// Returns error for service API and SDK errors.
// See the object firewall - multicast-address6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectFirewallMulticastAddress6(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/firewall/multicast-address6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectFirewallMulticastAddress6 API operation for FortiManager gets the Multicast Address6
// with the specified index value.
// Returns the requested Multicast Address6 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - multicast-address6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectFirewallMulticastAddress6(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/multicast-address6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectWafMainClass API operation for FortiManager creates a new Main Class.
// Returns the index value of the Main Class and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object waf - main-class chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectWafMainClass(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/waf/main-class"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectWafMainClass API operation for FortiManager updates the specified Main Class.
// Returns the index value of the Main Class and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object waf - main-class chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectWafMainClass(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/waf/main-class"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectWafMainClass API operation for FortiManager deletes the specified Main Class.
// Returns error for service API and SDK errors.
// See the object waf - main-class chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectWafMainClass(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/waf/main-class"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectWafMainClass API operation for FortiManager gets the Main Class
// with the specified index value.
// Returns the requested Main Class value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object waf - main-class chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectWafMainClass(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/waf/main-class"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectWafSubClass API operation for FortiManager creates a new Sub Class.
// Returns the index value of the Sub Class and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object waf - sub-class chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectWafSubClass(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/waf/sub-class"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectWafSubClass API operation for FortiManager updates the specified Sub Class.
// Returns the index value of the Sub Class and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object waf - sub-class chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectWafSubClass(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/waf/sub-class"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectWafSubClass API operation for FortiManager deletes the specified Sub Class.
// Returns error for service API and SDK errors.
// See the object waf - sub-class chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectWafSubClass(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/waf/sub-class"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectWafSubClass API operation for FortiManager gets the Sub Class
// with the specified index value.
// Returns the requested Sub Class value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object waf - sub-class chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectWafSubClass(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/waf/sub-class"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectWafSignature API operation for FortiManager creates a new Signature.
// Returns the index value of the Signature and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object waf - signature chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectWafSignature(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/waf/signature"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectWafSignature API operation for FortiManager updates the specified Signature.
// Returns the index value of the Signature and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object waf - signature chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectWafSignature(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/waf/signature"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectWafSignature API operation for FortiManager deletes the specified Signature.
// Returns error for service API and SDK errors.
// See the object waf - signature chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectWafSignature(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/waf/signature"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectWafSignature API operation for FortiManager gets the Signature
// with the specified index value.
// Returns the requested Signature value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object waf - signature chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectWafSignature(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/waf/signature"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectWafProfile API operation for FortiManager creates a new Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object waf - profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectWafProfile(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/waf/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectWafProfile API operation for FortiManager updates the specified Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object waf - profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectWafProfile(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/waf/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectWafProfile API operation for FortiManager deletes the specified Profile.
// Returns error for service API and SDK errors.
// See the object waf - profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectWafProfile(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/waf/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectWafProfile API operation for FortiManager gets the Profile
// with the specified index value.
// Returns the requested Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object waf - profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectWafProfile(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/waf/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectDnsfilterProfile API operation for FortiManager creates a new Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object dnsfilter - profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectDnsfilterProfile(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/dnsfilter/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectDnsfilterProfile API operation for FortiManager updates the specified Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object dnsfilter - profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectDnsfilterProfile(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/dnsfilter/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectDnsfilterProfile API operation for FortiManager deletes the specified Profile.
// Returns error for service API and SDK errors.
// See the object dnsfilter - profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectDnsfilterProfile(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/dnsfilter/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectDnsfilterProfile API operation for FortiManager gets the Profile
// with the specified index value.
// Returns the requested Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object dnsfilter - profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectDnsfilterProfile(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/dnsfilter/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectWebProxyWisp API operation for FortiManager creates a new Wisp.
// Returns the index value of the Wisp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object web-proxy - wisp chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectWebProxyWisp(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/web-proxy/wisp"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectWebProxyWisp API operation for FortiManager updates the specified Wisp.
// Returns the index value of the Wisp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object web-proxy - wisp chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectWebProxyWisp(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/web-proxy/wisp"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectWebProxyWisp API operation for FortiManager deletes the specified Wisp.
// Returns error for service API and SDK errors.
// See the object web-proxy - wisp chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectWebProxyWisp(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/web-proxy/wisp"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectWebProxyWisp API operation for FortiManager gets the Wisp
// with the specified index value.
// Returns the requested Wisp value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object web-proxy - wisp chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectWebProxyWisp(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/web-proxy/wisp"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectUserKrbKeytab API operation for FortiManager creates a new Krb Keytab.
// Returns the index value of the Krb Keytab and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - krb-keytab chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectUserKrbKeytab(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/krb-keytab"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectUserKrbKeytab API operation for FortiManager updates the specified Krb Keytab.
// Returns the index value of the Krb Keytab and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - krb-keytab chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectUserKrbKeytab(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/krb-keytab"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectUserKrbKeytab API operation for FortiManager deletes the specified Krb Keytab.
// Returns error for service API and SDK errors.
// See the object user - krb-keytab chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectUserKrbKeytab(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/user/krb-keytab"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectUserKrbKeytab API operation for FortiManager gets the Krb Keytab
// with the specified index value.
// Returns the requested Krb Keytab value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - krb-keytab chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectUserKrbKeytab(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/krb-keytab"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectGtpApn API operation for FortiManager creates a new Apn.
// Returns the index value of the Apn and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object gtp - apn chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectGtpApn(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/gtp/apn"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectGtpApn API operation for FortiManager updates the specified Apn.
// Returns the index value of the Apn and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object gtp - apn chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectGtpApn(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/gtp/apn"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectGtpApn API operation for FortiManager deletes the specified Apn.
// Returns error for service API and SDK errors.
// See the object gtp - apn chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectGtpApn(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/gtp/apn"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectGtpApn API operation for FortiManager gets the Apn
// with the specified index value.
// Returns the requested Apn value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object gtp - apn chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectGtpApn(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/gtp/apn"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectGtpApngrp API operation for FortiManager creates a new Apngrp.
// Returns the index value of the Apngrp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object gtp - apngrp chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectGtpApngrp(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/gtp/apngrp"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectGtpApngrp API operation for FortiManager updates the specified Apngrp.
// Returns the index value of the Apngrp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object gtp - apngrp chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectGtpApngrp(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/gtp/apngrp"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectGtpApngrp API operation for FortiManager deletes the specified Apngrp.
// Returns error for service API and SDK errors.
// See the object gtp - apngrp chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectGtpApngrp(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/gtp/apngrp"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectGtpApngrp API operation for FortiManager gets the Apngrp
// with the specified index value.
// Returns the requested Apngrp value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object gtp - apngrp chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectGtpApngrp(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/gtp/apngrp"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectApplicationCategories API operation for FortiManager creates a new Categories.
// Returns the index value of the Categories and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object application - categories chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectApplicationCategories(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/application/categories"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectApplicationCategories API operation for FortiManager updates the specified Categories.
// Returns the index value of the Categories and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object application - categories chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectApplicationCategories(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/application/categories"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectApplicationCategories API operation for FortiManager deletes the specified Categories.
// Returns error for service API and SDK errors.
// See the object application - categories chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectApplicationCategories(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/application/categories"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectApplicationCategories API operation for FortiManager gets the Categories
// with the specified index value.
// Returns the requested Categories value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object application - categories chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectApplicationCategories(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/application/categories"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectApplicationGroup API operation for FortiManager creates a new Group.
// Returns the index value of the Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object application - group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectApplicationGroup(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/application/group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectApplicationGroup API operation for FortiManager updates the specified Group.
// Returns the index value of the Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object application - group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectApplicationGroup(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/application/group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectApplicationGroup API operation for FortiManager deletes the specified Group.
// Returns error for service API and SDK errors.
// See the object application - group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectApplicationGroup(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/application/group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectApplicationGroup API operation for FortiManager gets the Group
// with the specified index value.
// Returns the requested Group value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object application - group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectApplicationGroup(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/application/group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectFirewallShapingProfile API operation for FortiManager creates a new Shaping Profile.
// Returns the index value of the Shaping Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - shaping-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectFirewallShapingProfile(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/shaping-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectFirewallShapingProfile API operation for FortiManager updates the specified Shaping Profile.
// Returns the index value of the Shaping Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - shaping-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectFirewallShapingProfile(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/shaping-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectFirewallShapingProfile API operation for FortiManager deletes the specified Shaping Profile.
// Returns error for service API and SDK errors.
// See the object firewall - shaping-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectFirewallShapingProfile(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/firewall/shaping-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectFirewallShapingProfile API operation for FortiManager gets the Shaping Profile
// with the specified index value.
// Returns the requested Shaping Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - shaping-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectFirewallShapingProfile(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/shaping-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectDnsfilterDomainFilter API operation for FortiManager creates a new Domain Filter.
// Returns the index value of the Domain Filter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object dnsfilter - domain-filter chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectDnsfilterDomainFilter(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/dnsfilter/domain-filter"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectDnsfilterDomainFilter API operation for FortiManager updates the specified Domain Filter.
// Returns the index value of the Domain Filter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object dnsfilter - domain-filter chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectDnsfilterDomainFilter(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/dnsfilter/domain-filter"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectDnsfilterDomainFilter API operation for FortiManager deletes the specified Domain Filter.
// Returns error for service API and SDK errors.
// See the object dnsfilter - domain-filter chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectDnsfilterDomainFilter(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/dnsfilter/domain-filter"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectDnsfilterDomainFilter API operation for FortiManager gets the Domain Filter
// with the specified index value.
// Returns the requested Domain Filter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object dnsfilter - domain-filter chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectDnsfilterDomainFilter(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/dnsfilter/domain-filter"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateObjectFirewallInternetService API operation for FortiManager updates the specified Internet Service.
// Returns the index value of the Internet Service and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - internet-service chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectFirewallInternetService(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/internet-service"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectFirewallInternetService API operation for FortiManager deletes the specified Internet Service.
// Returns error for service API and SDK errors.
// See the object firewall - internet-service chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectFirewallInternetService(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for object firewall - internet-service
	return
}

// ReadObjectFirewallInternetService API operation for FortiManager gets the Internet Service
// with the specified index value.
// Returns the requested Internet Service value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - internet-service chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectFirewallInternetService(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/internet-service"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectFirewallInternetServiceCustom API operation for FortiManager creates a new Internet Service Custom.
// Returns the index value of the Internet Service Custom and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - internet-service-custom chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectFirewallInternetServiceCustom(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/internet-service-custom"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectFirewallInternetServiceCustom API operation for FortiManager updates the specified Internet Service Custom.
// Returns the index value of the Internet Service Custom and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - internet-service-custom chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectFirewallInternetServiceCustom(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/internet-service-custom"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectFirewallInternetServiceCustom API operation for FortiManager deletes the specified Internet Service Custom.
// Returns error for service API and SDK errors.
// See the object firewall - internet-service-custom chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectFirewallInternetServiceCustom(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/firewall/internet-service-custom"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectFirewallInternetServiceCustom API operation for FortiManager gets the Internet Service Custom
// with the specified index value.
// Returns the requested Internet Service Custom value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - internet-service-custom chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectFirewallInternetServiceCustom(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/internet-service-custom"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectFirewallProxyAddress API operation for FortiManager creates a new Proxy Address.
// Returns the index value of the Proxy Address and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - proxy-address chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectFirewallProxyAddress(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/proxy-address"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectFirewallProxyAddress API operation for FortiManager updates the specified Proxy Address.
// Returns the index value of the Proxy Address and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - proxy-address chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectFirewallProxyAddress(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/proxy-address"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectFirewallProxyAddress API operation for FortiManager deletes the specified Proxy Address.
// Returns error for service API and SDK errors.
// See the object firewall - proxy-address chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectFirewallProxyAddress(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/firewall/proxy-address"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectFirewallProxyAddress API operation for FortiManager gets the Proxy Address
// with the specified index value.
// Returns the requested Proxy Address value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - proxy-address chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectFirewallProxyAddress(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/proxy-address"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectFirewallProxyAddrgrp API operation for FortiManager creates a new Proxy Addrgrp.
// Returns the index value of the Proxy Addrgrp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - proxy-addrgrp chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectFirewallProxyAddrgrp(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/proxy-addrgrp"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectFirewallProxyAddrgrp API operation for FortiManager updates the specified Proxy Addrgrp.
// Returns the index value of the Proxy Addrgrp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - proxy-addrgrp chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectFirewallProxyAddrgrp(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/proxy-addrgrp"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectFirewallProxyAddrgrp API operation for FortiManager deletes the specified Proxy Addrgrp.
// Returns error for service API and SDK errors.
// See the object firewall - proxy-addrgrp chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectFirewallProxyAddrgrp(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/firewall/proxy-addrgrp"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectFirewallProxyAddrgrp API operation for FortiManager gets the Proxy Addrgrp
// with the specified index value.
// Returns the requested Proxy Addrgrp value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - proxy-addrgrp chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectFirewallProxyAddrgrp(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/proxy-addrgrp"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectWirelessControllerBonjourProfile API operation for FortiManager creates a new Bonjour Profile.
// Returns the index value of the Bonjour Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - bonjour-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectWirelessControllerBonjourProfile(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/bonjour-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectWirelessControllerBonjourProfile API operation for FortiManager updates the specified Bonjour Profile.
// Returns the index value of the Bonjour Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - bonjour-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectWirelessControllerBonjourProfile(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/bonjour-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectWirelessControllerBonjourProfile API operation for FortiManager deletes the specified Bonjour Profile.
// Returns error for service API and SDK errors.
// See the object wireless-controller - bonjour-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectWirelessControllerBonjourProfile(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/wireless-controller/bonjour-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectWirelessControllerBonjourProfile API operation for FortiManager gets the Bonjour Profile
// with the specified index value.
// Returns the requested Bonjour Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - bonjour-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectWirelessControllerBonjourProfile(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/bonjour-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectSwitchControllerSecurityPolicy8021X API operation for FortiManager creates a new Security Policy802 1X.
// Returns the index value of the Security Policy802 1X and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object switch-controller - security-policy 802-1X chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectSwitchControllerSecurityPolicy8021X(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/switch-controller/security-policy/802-1X"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectSwitchControllerSecurityPolicy8021X API operation for FortiManager updates the specified Security Policy802 1X.
// Returns the index value of the Security Policy802 1X and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object switch-controller - security-policy 802-1X chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectSwitchControllerSecurityPolicy8021X(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/switch-controller/security-policy/802-1X"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectSwitchControllerSecurityPolicy8021X API operation for FortiManager deletes the specified Security Policy802 1X.
// Returns error for service API and SDK errors.
// See the object switch-controller - security-policy 802-1X chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectSwitchControllerSecurityPolicy8021X(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/switch-controller/security-policy/802-1X"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectSwitchControllerSecurityPolicy8021X API operation for FortiManager gets the Security Policy802 1X
// with the specified index value.
// Returns the requested Security Policy802 1X value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object switch-controller - security-policy 802-1X chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectSwitchControllerSecurityPolicy8021X(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/switch-controller/security-policy/802-1X"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectSwitchControllerLldpProfile API operation for FortiManager creates a new Lldp Profile.
// Returns the index value of the Lldp Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object switch-controller - lldp-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectSwitchControllerLldpProfile(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/switch-controller/lldp-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectSwitchControllerLldpProfile API operation for FortiManager updates the specified Lldp Profile.
// Returns the index value of the Lldp Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object switch-controller - lldp-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectSwitchControllerLldpProfile(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/switch-controller/lldp-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectSwitchControllerLldpProfile API operation for FortiManager deletes the specified Lldp Profile.
// Returns error for service API and SDK errors.
// See the object switch-controller - lldp-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectSwitchControllerLldpProfile(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/switch-controller/lldp-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectSwitchControllerLldpProfile API operation for FortiManager gets the Lldp Profile
// with the specified index value.
// Returns the requested Lldp Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object switch-controller - lldp-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectSwitchControllerLldpProfile(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/switch-controller/lldp-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectWirelessControllerQosProfile API operation for FortiManager creates a new Qos Profile.
// Returns the index value of the Qos Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - qos-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectWirelessControllerQosProfile(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/qos-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectWirelessControllerQosProfile API operation for FortiManager updates the specified Qos Profile.
// Returns the index value of the Qos Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - qos-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectWirelessControllerQosProfile(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/qos-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectWirelessControllerQosProfile API operation for FortiManager deletes the specified Qos Profile.
// Returns error for service API and SDK errors.
// See the object wireless-controller - qos-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectWirelessControllerQosProfile(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/wireless-controller/qos-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectWirelessControllerQosProfile API operation for FortiManager gets the Qos Profile
// with the specified index value.
// Returns the requested Qos Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - qos-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectWirelessControllerQosProfile(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/qos-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectAuthenticationScheme API operation for FortiManager creates a new Scheme.
// Returns the index value of the Scheme and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object authentication - scheme chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectAuthenticationScheme(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/authentication/scheme"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectAuthenticationScheme API operation for FortiManager updates the specified Scheme.
// Returns the index value of the Scheme and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object authentication - scheme chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectAuthenticationScheme(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/authentication/scheme"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectAuthenticationScheme API operation for FortiManager deletes the specified Scheme.
// Returns error for service API and SDK errors.
// See the object authentication - scheme chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectAuthenticationScheme(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/authentication/scheme"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectAuthenticationScheme API operation for FortiManager gets the Scheme
// with the specified index value.
// Returns the requested Scheme value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object authentication - scheme chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectAuthenticationScheme(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/authentication/scheme"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectGtpMessageFilterV0V1 API operation for FortiManager creates a new Message Filter V0V1.
// Returns the index value of the Message Filter V0V1 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object gtp - message-filter-v0v1 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectGtpMessageFilterV0V1(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/gtp/message-filter-v0v1"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectGtpMessageFilterV0V1 API operation for FortiManager updates the specified Message Filter V0V1.
// Returns the index value of the Message Filter V0V1 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object gtp - message-filter-v0v1 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectGtpMessageFilterV0V1(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/gtp/message-filter-v0v1"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectGtpMessageFilterV0V1 API operation for FortiManager deletes the specified Message Filter V0V1.
// Returns error for service API and SDK errors.
// See the object gtp - message-filter-v0v1 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectGtpMessageFilterV0V1(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/gtp/message-filter-v0v1"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectGtpMessageFilterV0V1 API operation for FortiManager gets the Message Filter V0V1
// with the specified index value.
// Returns the requested Message Filter V0V1 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object gtp - message-filter-v0v1 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectGtpMessageFilterV0V1(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/gtp/message-filter-v0v1"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectGtpMessageFilterV2 API operation for FortiManager creates a new Message Filter V2.
// Returns the index value of the Message Filter V2 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object gtp - message-filter-v2 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectGtpMessageFilterV2(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/gtp/message-filter-v2"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectGtpMessageFilterV2 API operation for FortiManager updates the specified Message Filter V2.
// Returns the index value of the Message Filter V2 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object gtp - message-filter-v2 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectGtpMessageFilterV2(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/gtp/message-filter-v2"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectGtpMessageFilterV2 API operation for FortiManager deletes the specified Message Filter V2.
// Returns error for service API and SDK errors.
// See the object gtp - message-filter-v2 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectGtpMessageFilterV2(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/gtp/message-filter-v2"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectGtpMessageFilterV2 API operation for FortiManager gets the Message Filter V2
// with the specified index value.
// Returns the requested Message Filter V2 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object gtp - message-filter-v2 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectGtpMessageFilterV2(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/gtp/message-filter-v2"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectGtpIeWhiteList API operation for FortiManager creates a new Ie White List.
// Returns the index value of the Ie White List and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object gtp - ie-white-list chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectGtpIeWhiteList(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/gtp/ie-white-list"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectGtpIeWhiteList API operation for FortiManager updates the specified Ie White List.
// Returns the index value of the Ie White List and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object gtp - ie-white-list chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectGtpIeWhiteList(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/gtp/ie-white-list"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectGtpIeWhiteList API operation for FortiManager deletes the specified Ie White List.
// Returns error for service API and SDK errors.
// See the object gtp - ie-white-list chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectGtpIeWhiteList(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/gtp/ie-white-list"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectGtpIeWhiteList API operation for FortiManager gets the Ie White List
// with the specified index value.
// Returns the requested Ie White List value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object gtp - ie-white-list chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectGtpIeWhiteList(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/gtp/ie-white-list"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectWirelessControllerBleProfile API operation for FortiManager creates a new Ble Profile.
// Returns the index value of the Ble Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - ble-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectWirelessControllerBleProfile(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/ble-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectWirelessControllerBleProfile API operation for FortiManager updates the specified Ble Profile.
// Returns the index value of the Ble Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - ble-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectWirelessControllerBleProfile(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/ble-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectWirelessControllerBleProfile API operation for FortiManager deletes the specified Ble Profile.
// Returns error for service API and SDK errors.
// See the object wireless-controller - ble-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectWirelessControllerBleProfile(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/wireless-controller/ble-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectWirelessControllerBleProfile API operation for FortiManager gets the Ble Profile
// with the specified index value.
// Returns the requested Ble Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - ble-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectWirelessControllerBleProfile(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/ble-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectSwitchControllerQosDot1PMap API operation for FortiManager creates a new QosDot1P Map.
// Returns the index value of the QosDot1P Map and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object switch-controller - qos dot1p-map chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectSwitchControllerQosDot1PMap(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/switch-controller/qos/dot1p-map"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectSwitchControllerQosDot1PMap API operation for FortiManager updates the specified QosDot1P Map.
// Returns the index value of the QosDot1P Map and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object switch-controller - qos dot1p-map chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectSwitchControllerQosDot1PMap(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/switch-controller/qos/dot1p-map"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectSwitchControllerQosDot1PMap API operation for FortiManager deletes the specified QosDot1P Map.
// Returns error for service API and SDK errors.
// See the object switch-controller - qos dot1p-map chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectSwitchControllerQosDot1PMap(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/switch-controller/qos/dot1p-map"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectSwitchControllerQosDot1PMap API operation for FortiManager gets the QosDot1P Map
// with the specified index value.
// Returns the requested QosDot1P Map value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object switch-controller - qos dot1p-map chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectSwitchControllerQosDot1PMap(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/switch-controller/qos/dot1p-map"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectSwitchControllerQosIpDscpMap API operation for FortiManager creates a new QosIp Dscp Map.
// Returns the index value of the QosIp Dscp Map and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object switch-controller - qos ip-dscp-map chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectSwitchControllerQosIpDscpMap(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/switch-controller/qos/ip-dscp-map"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectSwitchControllerQosIpDscpMap API operation for FortiManager updates the specified QosIp Dscp Map.
// Returns the index value of the QosIp Dscp Map and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object switch-controller - qos ip-dscp-map chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectSwitchControllerQosIpDscpMap(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/switch-controller/qos/ip-dscp-map"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectSwitchControllerQosIpDscpMap API operation for FortiManager deletes the specified QosIp Dscp Map.
// Returns error for service API and SDK errors.
// See the object switch-controller - qos ip-dscp-map chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectSwitchControllerQosIpDscpMap(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/switch-controller/qos/ip-dscp-map"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectSwitchControllerQosIpDscpMap API operation for FortiManager gets the QosIp Dscp Map
// with the specified index value.
// Returns the requested QosIp Dscp Map value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object switch-controller - qos ip-dscp-map chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectSwitchControllerQosIpDscpMap(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/switch-controller/qos/ip-dscp-map"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectSwitchControllerQosQueuePolicy API operation for FortiManager creates a new QosQueue Policy.
// Returns the index value of the QosQueue Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object switch-controller - qos queue-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectSwitchControllerQosQueuePolicy(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/switch-controller/qos/queue-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectSwitchControllerQosQueuePolicy API operation for FortiManager updates the specified QosQueue Policy.
// Returns the index value of the QosQueue Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object switch-controller - qos queue-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectSwitchControllerQosQueuePolicy(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/switch-controller/qos/queue-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectSwitchControllerQosQueuePolicy API operation for FortiManager deletes the specified QosQueue Policy.
// Returns error for service API and SDK errors.
// See the object switch-controller - qos queue-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectSwitchControllerQosQueuePolicy(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/switch-controller/qos/queue-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectSwitchControllerQosQueuePolicy API operation for FortiManager gets the QosQueue Policy
// with the specified index value.
// Returns the requested QosQueue Policy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object switch-controller - qos queue-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectSwitchControllerQosQueuePolicy(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/switch-controller/qos/queue-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectSwitchControllerQosQosPolicy API operation for FortiManager creates a new QosQos Policy.
// Returns the index value of the QosQos Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object switch-controller - qos qos-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectSwitchControllerQosQosPolicy(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/switch-controller/qos/qos-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectSwitchControllerQosQosPolicy API operation for FortiManager updates the specified QosQos Policy.
// Returns the index value of the QosQos Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object switch-controller - qos qos-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectSwitchControllerQosQosPolicy(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/switch-controller/qos/qos-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectSwitchControllerQosQosPolicy API operation for FortiManager deletes the specified QosQos Policy.
// Returns error for service API and SDK errors.
// See the object switch-controller - qos qos-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectSwitchControllerQosQosPolicy(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/switch-controller/qos/qos-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectSwitchControllerQosQosPolicy API operation for FortiManager gets the QosQos Policy
// with the specified index value.
// Returns the requested QosQos Policy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object switch-controller - qos qos-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectSwitchControllerQosQosPolicy(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/switch-controller/qos/qos-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectWirelessControllerHotspot20AnqpVenueName API operation for FortiManager creates a new Hotspot20Anqp Venue Name.
// Returns the index value of the Hotspot20Anqp Venue Name and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - hotspot20 anqp-venue-name chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectWirelessControllerHotspot20AnqpVenueName(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/hotspot20/anqp-venue-name"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectWirelessControllerHotspot20AnqpVenueName API operation for FortiManager updates the specified Hotspot20Anqp Venue Name.
// Returns the index value of the Hotspot20Anqp Venue Name and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - hotspot20 anqp-venue-name chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectWirelessControllerHotspot20AnqpVenueName(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/hotspot20/anqp-venue-name"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectWirelessControllerHotspot20AnqpVenueName API operation for FortiManager deletes the specified Hotspot20Anqp Venue Name.
// Returns error for service API and SDK errors.
// See the object wireless-controller - hotspot20 anqp-venue-name chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectWirelessControllerHotspot20AnqpVenueName(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/wireless-controller/hotspot20/anqp-venue-name"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectWirelessControllerHotspot20AnqpVenueName API operation for FortiManager gets the Hotspot20Anqp Venue Name
// with the specified index value.
// Returns the requested Hotspot20Anqp Venue Name value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - hotspot20 anqp-venue-name chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectWirelessControllerHotspot20AnqpVenueName(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/hotspot20/anqp-venue-name"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectWirelessControllerHotspot20AnqpNetworkAuthType API operation for FortiManager creates a new Hotspot20Anqp Network Auth Type.
// Returns the index value of the Hotspot20Anqp Network Auth Type and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - hotspot20 anqp-network-auth-type chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectWirelessControllerHotspot20AnqpNetworkAuthType(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/hotspot20/anqp-network-auth-type"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectWirelessControllerHotspot20AnqpNetworkAuthType API operation for FortiManager updates the specified Hotspot20Anqp Network Auth Type.
// Returns the index value of the Hotspot20Anqp Network Auth Type and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - hotspot20 anqp-network-auth-type chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectWirelessControllerHotspot20AnqpNetworkAuthType(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/hotspot20/anqp-network-auth-type"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectWirelessControllerHotspot20AnqpNetworkAuthType API operation for FortiManager deletes the specified Hotspot20Anqp Network Auth Type.
// Returns error for service API and SDK errors.
// See the object wireless-controller - hotspot20 anqp-network-auth-type chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectWirelessControllerHotspot20AnqpNetworkAuthType(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/wireless-controller/hotspot20/anqp-network-auth-type"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectWirelessControllerHotspot20AnqpNetworkAuthType API operation for FortiManager gets the Hotspot20Anqp Network Auth Type
// with the specified index value.
// Returns the requested Hotspot20Anqp Network Auth Type value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - hotspot20 anqp-network-auth-type chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectWirelessControllerHotspot20AnqpNetworkAuthType(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/hotspot20/anqp-network-auth-type"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectWirelessControllerHotspot20AnqpRoamingConsortium API operation for FortiManager creates a new Hotspot20Anqp Roaming Consortium.
// Returns the index value of the Hotspot20Anqp Roaming Consortium and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - hotspot20 anqp-roaming-consortium chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectWirelessControllerHotspot20AnqpRoamingConsortium(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/hotspot20/anqp-roaming-consortium"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectWirelessControllerHotspot20AnqpRoamingConsortium API operation for FortiManager updates the specified Hotspot20Anqp Roaming Consortium.
// Returns the index value of the Hotspot20Anqp Roaming Consortium and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - hotspot20 anqp-roaming-consortium chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectWirelessControllerHotspot20AnqpRoamingConsortium(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/hotspot20/anqp-roaming-consortium"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectWirelessControllerHotspot20AnqpRoamingConsortium API operation for FortiManager deletes the specified Hotspot20Anqp Roaming Consortium.
// Returns error for service API and SDK errors.
// See the object wireless-controller - hotspot20 anqp-roaming-consortium chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectWirelessControllerHotspot20AnqpRoamingConsortium(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/wireless-controller/hotspot20/anqp-roaming-consortium"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectWirelessControllerHotspot20AnqpRoamingConsortium API operation for FortiManager gets the Hotspot20Anqp Roaming Consortium
// with the specified index value.
// Returns the requested Hotspot20Anqp Roaming Consortium value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - hotspot20 anqp-roaming-consortium chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectWirelessControllerHotspot20AnqpRoamingConsortium(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/hotspot20/anqp-roaming-consortium"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectWirelessControllerHotspot20AnqpNaiRealm API operation for FortiManager creates a new Hotspot20Anqp Nai Realm.
// Returns the index value of the Hotspot20Anqp Nai Realm and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - hotspot20 anqp-nai-realm chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectWirelessControllerHotspot20AnqpNaiRealm(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/hotspot20/anqp-nai-realm"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectWirelessControllerHotspot20AnqpNaiRealm API operation for FortiManager updates the specified Hotspot20Anqp Nai Realm.
// Returns the index value of the Hotspot20Anqp Nai Realm and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - hotspot20 anqp-nai-realm chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectWirelessControllerHotspot20AnqpNaiRealm(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/hotspot20/anqp-nai-realm"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectWirelessControllerHotspot20AnqpNaiRealm API operation for FortiManager deletes the specified Hotspot20Anqp Nai Realm.
// Returns error for service API and SDK errors.
// See the object wireless-controller - hotspot20 anqp-nai-realm chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectWirelessControllerHotspot20AnqpNaiRealm(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/wireless-controller/hotspot20/anqp-nai-realm"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectWirelessControllerHotspot20AnqpNaiRealm API operation for FortiManager gets the Hotspot20Anqp Nai Realm
// with the specified index value.
// Returns the requested Hotspot20Anqp Nai Realm value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - hotspot20 anqp-nai-realm chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectWirelessControllerHotspot20AnqpNaiRealm(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/hotspot20/anqp-nai-realm"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectWirelessControllerHotspot20Anqp3GppCellular API operation for FortiManager creates a new Hotspot20Anqp 3Gpp Cellular.
// Returns the index value of the Hotspot20Anqp 3Gpp Cellular and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - hotspot20 anqp-3gpp-cellular chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectWirelessControllerHotspot20Anqp3GppCellular(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/hotspot20/anqp-3gpp-cellular"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectWirelessControllerHotspot20Anqp3GppCellular API operation for FortiManager updates the specified Hotspot20Anqp 3Gpp Cellular.
// Returns the index value of the Hotspot20Anqp 3Gpp Cellular and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - hotspot20 anqp-3gpp-cellular chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectWirelessControllerHotspot20Anqp3GppCellular(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/hotspot20/anqp-3gpp-cellular"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectWirelessControllerHotspot20Anqp3GppCellular API operation for FortiManager deletes the specified Hotspot20Anqp 3Gpp Cellular.
// Returns error for service API and SDK errors.
// See the object wireless-controller - hotspot20 anqp-3gpp-cellular chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectWirelessControllerHotspot20Anqp3GppCellular(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/wireless-controller/hotspot20/anqp-3gpp-cellular"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectWirelessControllerHotspot20Anqp3GppCellular API operation for FortiManager gets the Hotspot20Anqp 3Gpp Cellular
// with the specified index value.
// Returns the requested Hotspot20Anqp 3Gpp Cellular value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - hotspot20 anqp-3gpp-cellular chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectWirelessControllerHotspot20Anqp3GppCellular(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/hotspot20/anqp-3gpp-cellular"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectWirelessControllerHotspot20AnqpIpAddressType API operation for FortiManager creates a new Hotspot20Anqp Ip Address Type.
// Returns the index value of the Hotspot20Anqp Ip Address Type and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - hotspot20 anqp-ip-address-type chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectWirelessControllerHotspot20AnqpIpAddressType(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/hotspot20/anqp-ip-address-type"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectWirelessControllerHotspot20AnqpIpAddressType API operation for FortiManager updates the specified Hotspot20Anqp Ip Address Type.
// Returns the index value of the Hotspot20Anqp Ip Address Type and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - hotspot20 anqp-ip-address-type chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectWirelessControllerHotspot20AnqpIpAddressType(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/hotspot20/anqp-ip-address-type"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectWirelessControllerHotspot20AnqpIpAddressType API operation for FortiManager deletes the specified Hotspot20Anqp Ip Address Type.
// Returns error for service API and SDK errors.
// See the object wireless-controller - hotspot20 anqp-ip-address-type chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectWirelessControllerHotspot20AnqpIpAddressType(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/wireless-controller/hotspot20/anqp-ip-address-type"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectWirelessControllerHotspot20AnqpIpAddressType API operation for FortiManager gets the Hotspot20Anqp Ip Address Type
// with the specified index value.
// Returns the requested Hotspot20Anqp Ip Address Type value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - hotspot20 anqp-ip-address-type chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectWirelessControllerHotspot20AnqpIpAddressType(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/hotspot20/anqp-ip-address-type"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectWirelessControllerHotspot20H2QpOperatorName API operation for FortiManager creates a new Hotspot20H2Qp Operator Name.
// Returns the index value of the Hotspot20H2Qp Operator Name and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - hotspot20 h2qp-operator-name chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectWirelessControllerHotspot20H2QpOperatorName(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/hotspot20/h2qp-operator-name"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectWirelessControllerHotspot20H2QpOperatorName API operation for FortiManager updates the specified Hotspot20H2Qp Operator Name.
// Returns the index value of the Hotspot20H2Qp Operator Name and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - hotspot20 h2qp-operator-name chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectWirelessControllerHotspot20H2QpOperatorName(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/hotspot20/h2qp-operator-name"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectWirelessControllerHotspot20H2QpOperatorName API operation for FortiManager deletes the specified Hotspot20H2Qp Operator Name.
// Returns error for service API and SDK errors.
// See the object wireless-controller - hotspot20 h2qp-operator-name chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectWirelessControllerHotspot20H2QpOperatorName(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/wireless-controller/hotspot20/h2qp-operator-name"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectWirelessControllerHotspot20H2QpOperatorName API operation for FortiManager gets the Hotspot20H2Qp Operator Name
// with the specified index value.
// Returns the requested Hotspot20H2Qp Operator Name value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - hotspot20 h2qp-operator-name chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectWirelessControllerHotspot20H2QpOperatorName(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/hotspot20/h2qp-operator-name"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectWirelessControllerHotspot20H2QpWanMetric API operation for FortiManager creates a new Hotspot20H2Qp Wan Metric.
// Returns the index value of the Hotspot20H2Qp Wan Metric and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - hotspot20 h2qp-wan-metric chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectWirelessControllerHotspot20H2QpWanMetric(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/hotspot20/h2qp-wan-metric"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectWirelessControllerHotspot20H2QpWanMetric API operation for FortiManager updates the specified Hotspot20H2Qp Wan Metric.
// Returns the index value of the Hotspot20H2Qp Wan Metric and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - hotspot20 h2qp-wan-metric chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectWirelessControllerHotspot20H2QpWanMetric(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/hotspot20/h2qp-wan-metric"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectWirelessControllerHotspot20H2QpWanMetric API operation for FortiManager deletes the specified Hotspot20H2Qp Wan Metric.
// Returns error for service API and SDK errors.
// See the object wireless-controller - hotspot20 h2qp-wan-metric chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectWirelessControllerHotspot20H2QpWanMetric(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/wireless-controller/hotspot20/h2qp-wan-metric"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectWirelessControllerHotspot20H2QpWanMetric API operation for FortiManager gets the Hotspot20H2Qp Wan Metric
// with the specified index value.
// Returns the requested Hotspot20H2Qp Wan Metric value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - hotspot20 h2qp-wan-metric chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectWirelessControllerHotspot20H2QpWanMetric(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/hotspot20/h2qp-wan-metric"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectWirelessControllerHotspot20H2QpConnCapability API operation for FortiManager creates a new Hotspot20H2Qp Conn Capability.
// Returns the index value of the Hotspot20H2Qp Conn Capability and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - hotspot20 h2qp-conn-capability chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectWirelessControllerHotspot20H2QpConnCapability(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/hotspot20/h2qp-conn-capability"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectWirelessControllerHotspot20H2QpConnCapability API operation for FortiManager updates the specified Hotspot20H2Qp Conn Capability.
// Returns the index value of the Hotspot20H2Qp Conn Capability and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - hotspot20 h2qp-conn-capability chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectWirelessControllerHotspot20H2QpConnCapability(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/hotspot20/h2qp-conn-capability"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectWirelessControllerHotspot20H2QpConnCapability API operation for FortiManager deletes the specified Hotspot20H2Qp Conn Capability.
// Returns error for service API and SDK errors.
// See the object wireless-controller - hotspot20 h2qp-conn-capability chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectWirelessControllerHotspot20H2QpConnCapability(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/wireless-controller/hotspot20/h2qp-conn-capability"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectWirelessControllerHotspot20H2QpConnCapability API operation for FortiManager gets the Hotspot20H2Qp Conn Capability
// with the specified index value.
// Returns the requested Hotspot20H2Qp Conn Capability value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - hotspot20 h2qp-conn-capability chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectWirelessControllerHotspot20H2QpConnCapability(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/hotspot20/h2qp-conn-capability"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectWirelessControllerHotspot20H2QpOsuProvider API operation for FortiManager creates a new Hotspot20H2Qp Osu Provider.
// Returns the index value of the Hotspot20H2Qp Osu Provider and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - hotspot20 h2qp-osu-provider chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectWirelessControllerHotspot20H2QpOsuProvider(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/hotspot20/h2qp-osu-provider"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectWirelessControllerHotspot20H2QpOsuProvider API operation for FortiManager updates the specified Hotspot20H2Qp Osu Provider.
// Returns the index value of the Hotspot20H2Qp Osu Provider and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - hotspot20 h2qp-osu-provider chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectWirelessControllerHotspot20H2QpOsuProvider(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/hotspot20/h2qp-osu-provider"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectWirelessControllerHotspot20H2QpOsuProvider API operation for FortiManager deletes the specified Hotspot20H2Qp Osu Provider.
// Returns error for service API and SDK errors.
// See the object wireless-controller - hotspot20 h2qp-osu-provider chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectWirelessControllerHotspot20H2QpOsuProvider(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/wireless-controller/hotspot20/h2qp-osu-provider"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectWirelessControllerHotspot20H2QpOsuProvider API operation for FortiManager gets the Hotspot20H2Qp Osu Provider
// with the specified index value.
// Returns the requested Hotspot20H2Qp Osu Provider value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - hotspot20 h2qp-osu-provider chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectWirelessControllerHotspot20H2QpOsuProvider(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/hotspot20/h2qp-osu-provider"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectWirelessControllerHotspot20QosMap API operation for FortiManager creates a new Hotspot20Qos Map.
// Returns the index value of the Hotspot20Qos Map and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - hotspot20 qos-map chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectWirelessControllerHotspot20QosMap(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/hotspot20/qos-map"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectWirelessControllerHotspot20QosMap API operation for FortiManager updates the specified Hotspot20Qos Map.
// Returns the index value of the Hotspot20Qos Map and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - hotspot20 qos-map chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectWirelessControllerHotspot20QosMap(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/hotspot20/qos-map"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectWirelessControllerHotspot20QosMap API operation for FortiManager deletes the specified Hotspot20Qos Map.
// Returns error for service API and SDK errors.
// See the object wireless-controller - hotspot20 qos-map chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectWirelessControllerHotspot20QosMap(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/wireless-controller/hotspot20/qos-map"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectWirelessControllerHotspot20QosMap API operation for FortiManager gets the Hotspot20Qos Map
// with the specified index value.
// Returns the requested Hotspot20Qos Map value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - hotspot20 qos-map chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectWirelessControllerHotspot20QosMap(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/hotspot20/qos-map"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectWirelessControllerHotspot20HsProfile API operation for FortiManager creates a new Hotspot20Hs Profile.
// Returns the index value of the Hotspot20Hs Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - hotspot20 hs-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectWirelessControllerHotspot20HsProfile(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/hotspot20/hs-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectWirelessControllerHotspot20HsProfile API operation for FortiManager updates the specified Hotspot20Hs Profile.
// Returns the index value of the Hotspot20Hs Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - hotspot20 hs-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectWirelessControllerHotspot20HsProfile(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/hotspot20/hs-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectWirelessControllerHotspot20HsProfile API operation for FortiManager deletes the specified Hotspot20Hs Profile.
// Returns error for service API and SDK errors.
// See the object wireless-controller - hotspot20 hs-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectWirelessControllerHotspot20HsProfile(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/wireless-controller/hotspot20/hs-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectWirelessControllerHotspot20HsProfile API operation for FortiManager gets the Hotspot20Hs Profile
// with the specified index value.
// Returns the requested Hotspot20Hs Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - hotspot20 hs-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectWirelessControllerHotspot20HsProfile(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/hotspot20/hs-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectSystemSdnConnector API operation for FortiManager creates a new Sdn Connector.
// Returns the index value of the Sdn Connector and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object system - sdn-connector chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectSystemSdnConnector(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/system/sdn-connector"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectSystemSdnConnector API operation for FortiManager updates the specified Sdn Connector.
// Returns the index value of the Sdn Connector and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object system - sdn-connector chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectSystemSdnConnector(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/system/sdn-connector"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectSystemSdnConnector API operation for FortiManager deletes the specified Sdn Connector.
// Returns error for service API and SDK errors.
// See the object system - sdn-connector chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectSystemSdnConnector(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/system/sdn-connector"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectSystemSdnConnector API operation for FortiManager gets the Sdn Connector
// with the specified index value.
// Returns the requested Sdn Connector value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object system - sdn-connector chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectSystemSdnConnector(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/system/sdn-connector"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectUserDomainController API operation for FortiManager creates a new Domain Controller.
// Returns the index value of the Domain Controller and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - domain-controller chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectUserDomainController(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/domain-controller"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectUserDomainController API operation for FortiManager updates the specified Domain Controller.
// Returns the index value of the Domain Controller and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - domain-controller chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectUserDomainController(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/domain-controller"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectUserDomainController API operation for FortiManager deletes the specified Domain Controller.
// Returns error for service API and SDK errors.
// See the object user - domain-controller chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectUserDomainController(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/user/domain-controller"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectUserDomainController API operation for FortiManager gets the Domain Controller
// with the specified index value.
// Returns the requested Domain Controller value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - domain-controller chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectUserDomainController(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/domain-controller"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectSystemObjectTagging API operation for FortiManager creates a new Object Tagging.
// Returns the index value of the Object Tagging and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object system - object-tagging chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectSystemObjectTagging(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/system/object-tagging"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectSystemObjectTagging API operation for FortiManager updates the specified Object Tagging.
// Returns the index value of the Object Tagging and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object system - object-tagging chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectSystemObjectTagging(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/system/object-tagging"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectSystemObjectTagging API operation for FortiManager deletes the specified Object Tagging.
// Returns error for service API and SDK errors.
// See the object system - object-tagging chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectSystemObjectTagging(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/system/object-tagging"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectSystemObjectTagging API operation for FortiManager gets the Object Tagging
// with the specified index value.
// Returns the requested Object Tagging value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object system - object-tagging chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectSystemObjectTagging(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/system/object-tagging"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectFirewallAddress6Template API operation for FortiManager creates a new Address6 Template.
// Returns the index value of the Address6 Template and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - address6-template chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectFirewallAddress6Template(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/address6-template"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectFirewallAddress6Template API operation for FortiManager updates the specified Address6 Template.
// Returns the index value of the Address6 Template and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - address6-template chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectFirewallAddress6Template(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/address6-template"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectFirewallAddress6Template API operation for FortiManager deletes the specified Address6 Template.
// Returns error for service API and SDK errors.
// See the object firewall - address6-template chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectFirewallAddress6Template(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/firewall/address6-template"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectFirewallAddress6Template API operation for FortiManager gets the Address6 Template
// with the specified index value.
// Returns the requested Address6 Template value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - address6-template chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectFirewallAddress6Template(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/address6-template"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectFirewallWildcardFqdnCustom API operation for FortiManager creates a new Wildcard FqdnCustom.
// Returns the index value of the Wildcard FqdnCustom and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - wildcard-fqdn custom chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectFirewallWildcardFqdnCustom(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/wildcard-fqdn/custom"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectFirewallWildcardFqdnCustom API operation for FortiManager updates the specified Wildcard FqdnCustom.
// Returns the index value of the Wildcard FqdnCustom and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - wildcard-fqdn custom chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectFirewallWildcardFqdnCustom(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/wildcard-fqdn/custom"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectFirewallWildcardFqdnCustom API operation for FortiManager deletes the specified Wildcard FqdnCustom.
// Returns error for service API and SDK errors.
// See the object firewall - wildcard-fqdn custom chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectFirewallWildcardFqdnCustom(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/firewall/wildcard-fqdn/custom"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectFirewallWildcardFqdnCustom API operation for FortiManager gets the Wildcard FqdnCustom
// with the specified index value.
// Returns the requested Wildcard FqdnCustom value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - wildcard-fqdn custom chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectFirewallWildcardFqdnCustom(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/wildcard-fqdn/custom"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectFirewallWildcardFqdnGroup API operation for FortiManager creates a new Wildcard FqdnGroup.
// Returns the index value of the Wildcard FqdnGroup and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - wildcard-fqdn group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectFirewallWildcardFqdnGroup(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/wildcard-fqdn/group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectFirewallWildcardFqdnGroup API operation for FortiManager updates the specified Wildcard FqdnGroup.
// Returns the index value of the Wildcard FqdnGroup and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - wildcard-fqdn group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectFirewallWildcardFqdnGroup(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/wildcard-fqdn/group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectFirewallWildcardFqdnGroup API operation for FortiManager deletes the specified Wildcard FqdnGroup.
// Returns error for service API and SDK errors.
// See the object firewall - wildcard-fqdn group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectFirewallWildcardFqdnGroup(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/firewall/wildcard-fqdn/group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectFirewallWildcardFqdnGroup API operation for FortiManager gets the Wildcard FqdnGroup
// with the specified index value.
// Returns the requested Wildcard FqdnGroup value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - wildcard-fqdn group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectFirewallWildcardFqdnGroup(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/wildcard-fqdn/group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectFirewallInternetServiceGroup API operation for FortiManager creates a new Internet Service Group.
// Returns the index value of the Internet Service Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - internet-service-group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectFirewallInternetServiceGroup(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/internet-service-group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectFirewallInternetServiceGroup API operation for FortiManager updates the specified Internet Service Group.
// Returns the index value of the Internet Service Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - internet-service-group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectFirewallInternetServiceGroup(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/internet-service-group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectFirewallInternetServiceGroup API operation for FortiManager deletes the specified Internet Service Group.
// Returns error for service API and SDK errors.
// See the object firewall - internet-service-group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectFirewallInternetServiceGroup(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/firewall/internet-service-group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectFirewallInternetServiceGroup API operation for FortiManager gets the Internet Service Group
// with the specified index value.
// Returns the requested Internet Service Group value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - internet-service-group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectFirewallInternetServiceGroup(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/internet-service-group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectFirewallInternetServiceCustomGroup API operation for FortiManager creates a new Internet Service Custom Group.
// Returns the index value of the Internet Service Custom Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - internet-service-custom-group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectFirewallInternetServiceCustomGroup(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/internet-service-custom-group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectFirewallInternetServiceCustomGroup API operation for FortiManager updates the specified Internet Service Custom Group.
// Returns the index value of the Internet Service Custom Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - internet-service-custom-group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectFirewallInternetServiceCustomGroup(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/internet-service-custom-group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectFirewallInternetServiceCustomGroup API operation for FortiManager deletes the specified Internet Service Custom Group.
// Returns error for service API and SDK errors.
// See the object firewall - internet-service-custom-group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectFirewallInternetServiceCustomGroup(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/firewall/internet-service-custom-group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectFirewallInternetServiceCustomGroup API operation for FortiManager gets the Internet Service Custom Group
// with the specified index value.
// Returns the requested Internet Service Custom Group value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - internet-service-custom-group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectFirewallInternetServiceCustomGroup(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/internet-service-custom-group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectSystemExternalResource API operation for FortiManager creates a new External Resource.
// Returns the index value of the External Resource and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object system - external-resource chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectSystemExternalResource(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/system/external-resource"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectSystemExternalResource API operation for FortiManager updates the specified External Resource.
// Returns the index value of the External Resource and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object system - external-resource chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectSystemExternalResource(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/system/external-resource"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectSystemExternalResource API operation for FortiManager deletes the specified External Resource.
// Returns error for service API and SDK errors.
// See the object system - external-resource chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectSystemExternalResource(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/system/external-resource"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectSystemExternalResource API operation for FortiManager gets the External Resource
// with the specified index value.
// Returns the requested External Resource value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object system - external-resource chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectSystemExternalResource(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/system/external-resource"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectSshFilterProfile API operation for FortiManager creates a new Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object ssh-filter - profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectSshFilterProfile(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/ssh-filter/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectSshFilterProfile API operation for FortiManager updates the specified Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object ssh-filter - profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectSshFilterProfile(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/ssh-filter/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectSshFilterProfile API operation for FortiManager deletes the specified Profile.
// Returns error for service API and SDK errors.
// See the object ssh-filter - profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectSshFilterProfile(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/ssh-filter/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectSshFilterProfile API operation for FortiManager gets the Profile
// with the specified index value.
// Returns the requested Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object ssh-filter - profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectSshFilterProfile(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/ssh-filter/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectGtpTunnelLimit API operation for FortiManager creates a new Tunnel Limit.
// Returns the index value of the Tunnel Limit and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object gtp - tunnel-limit chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectGtpTunnelLimit(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/gtp/tunnel-limit"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectGtpTunnelLimit API operation for FortiManager updates the specified Tunnel Limit.
// Returns the index value of the Tunnel Limit and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object gtp - tunnel-limit chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectGtpTunnelLimit(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/gtp/tunnel-limit"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectGtpTunnelLimit API operation for FortiManager deletes the specified Tunnel Limit.
// Returns error for service API and SDK errors.
// See the object gtp - tunnel-limit chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectGtpTunnelLimit(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/gtp/tunnel-limit"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectGtpTunnelLimit API operation for FortiManager gets the Tunnel Limit
// with the specified index value.
// Returns the requested Tunnel Limit value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object gtp - tunnel-limit chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectGtpTunnelLimit(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/gtp/tunnel-limit"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectWirelessControllerUtmProfile API operation for FortiManager creates a new Utm Profile.
// Returns the index value of the Utm Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - utm-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectWirelessControllerUtmProfile(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/utm-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectWirelessControllerUtmProfile API operation for FortiManager updates the specified Utm Profile.
// Returns the index value of the Utm Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - utm-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectWirelessControllerUtmProfile(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/utm-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectWirelessControllerUtmProfile API operation for FortiManager deletes the specified Utm Profile.
// Returns error for service API and SDK errors.
// See the object wireless-controller - utm-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectWirelessControllerUtmProfile(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/wireless-controller/utm-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectWirelessControllerUtmProfile API operation for FortiManager gets the Utm Profile
// with the specified index value.
// Returns the requested Utm Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - utm-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectWirelessControllerUtmProfile(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/utm-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectFirewallSshLocalCa API operation for FortiManager creates a new SshLocal Ca.
// Returns the index value of the SshLocal Ca and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - ssh local-ca chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectFirewallSshLocalCa(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/ssh/local-ca"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectFirewallSshLocalCa API operation for FortiManager updates the specified SshLocal Ca.
// Returns the index value of the SshLocal Ca and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - ssh local-ca chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectFirewallSshLocalCa(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/ssh/local-ca"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectFirewallSshLocalCa API operation for FortiManager deletes the specified SshLocal Ca.
// Returns error for service API and SDK errors.
// See the object firewall - ssh local-ca chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectFirewallSshLocalCa(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/firewall/ssh/local-ca"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectFirewallSshLocalCa API operation for FortiManager gets the SshLocal Ca
// with the specified index value.
// Returns the requested SshLocal Ca value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - ssh local-ca chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectFirewallSshLocalCa(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/ssh/local-ca"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectWirelessControllerRegion API operation for FortiManager creates a new Region.
// Returns the index value of the Region and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - region chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectWirelessControllerRegion(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/region"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectWirelessControllerRegion API operation for FortiManager updates the specified Region.
// Returns the index value of the Region and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - region chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectWirelessControllerRegion(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/region"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectWirelessControllerRegion API operation for FortiManager deletes the specified Region.
// Returns error for service API and SDK errors.
// See the object wireless-controller - region chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectWirelessControllerRegion(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/wireless-controller/region"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectWirelessControllerRegion API operation for FortiManager gets the Region
// with the specified index value.
// Returns the requested Region value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - region chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectWirelessControllerRegion(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/region"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectDlpSensitivity API operation for FortiManager creates a new Sensitivity.
// Returns the index value of the Sensitivity and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object dlp - sensitivity chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectDlpSensitivity(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/dlp/sensitivity"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectDlpSensitivity API operation for FortiManager updates the specified Sensitivity.
// Returns the index value of the Sensitivity and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object dlp - sensitivity chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectDlpSensitivity(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/dlp/sensitivity"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectDlpSensitivity API operation for FortiManager deletes the specified Sensitivity.
// Returns error for service API and SDK errors.
// See the object dlp - sensitivity chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectDlpSensitivity(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/dlp/sensitivity"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectDlpSensitivity API operation for FortiManager gets the Sensitivity
// with the specified index value.
// Returns the requested Sensitivity value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object dlp - sensitivity chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectDlpSensitivity(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/dlp/sensitivity"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectEmailfilterBword API operation for FortiManager creates a new Bword.
// Returns the index value of the Bword and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object emailfilter - bword chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectEmailfilterBword(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/emailfilter/bword"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectEmailfilterBword API operation for FortiManager updates the specified Bword.
// Returns the index value of the Bword and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object emailfilter - bword chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectEmailfilterBword(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/emailfilter/bword"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectEmailfilterBword API operation for FortiManager deletes the specified Bword.
// Returns error for service API and SDK errors.
// See the object emailfilter - bword chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectEmailfilterBword(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/emailfilter/bword"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectEmailfilterBword API operation for FortiManager gets the Bword
// with the specified index value.
// Returns the requested Bword value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object emailfilter - bword chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectEmailfilterBword(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/emailfilter/bword"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectEmailfilterBwl API operation for FortiManager creates a new Bwl.
// Returns the index value of the Bwl and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object emailfilter - bwl chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectEmailfilterBwl(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/emailfilter/bwl"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectEmailfilterBwl API operation for FortiManager updates the specified Bwl.
// Returns the index value of the Bwl and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object emailfilter - bwl chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectEmailfilterBwl(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/emailfilter/bwl"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectEmailfilterBwl API operation for FortiManager deletes the specified Bwl.
// Returns error for service API and SDK errors.
// See the object emailfilter - bwl chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectEmailfilterBwl(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/emailfilter/bwl"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectEmailfilterBwl API operation for FortiManager gets the Bwl
// with the specified index value.
// Returns the requested Bwl value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object emailfilter - bwl chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectEmailfilterBwl(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/emailfilter/bwl"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectEmailfilterMheader API operation for FortiManager creates a new Mheader.
// Returns the index value of the Mheader and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object emailfilter - mheader chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectEmailfilterMheader(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/emailfilter/mheader"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectEmailfilterMheader API operation for FortiManager updates the specified Mheader.
// Returns the index value of the Mheader and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object emailfilter - mheader chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectEmailfilterMheader(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/emailfilter/mheader"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectEmailfilterMheader API operation for FortiManager deletes the specified Mheader.
// Returns error for service API and SDK errors.
// See the object emailfilter - mheader chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectEmailfilterMheader(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/emailfilter/mheader"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectEmailfilterMheader API operation for FortiManager gets the Mheader
// with the specified index value.
// Returns the requested Mheader value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object emailfilter - mheader chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectEmailfilterMheader(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/emailfilter/mheader"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectEmailfilterDnsbl API operation for FortiManager creates a new Dnsbl.
// Returns the index value of the Dnsbl and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object emailfilter - dnsbl chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectEmailfilterDnsbl(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/emailfilter/dnsbl"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectEmailfilterDnsbl API operation for FortiManager updates the specified Dnsbl.
// Returns the index value of the Dnsbl and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object emailfilter - dnsbl chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectEmailfilterDnsbl(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/emailfilter/dnsbl"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectEmailfilterDnsbl API operation for FortiManager deletes the specified Dnsbl.
// Returns error for service API and SDK errors.
// See the object emailfilter - dnsbl chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectEmailfilterDnsbl(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/emailfilter/dnsbl"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectEmailfilterDnsbl API operation for FortiManager gets the Dnsbl
// with the specified index value.
// Returns the requested Dnsbl value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object emailfilter - dnsbl chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectEmailfilterDnsbl(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/emailfilter/dnsbl"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectEmailfilterIptrust API operation for FortiManager creates a new Iptrust.
// Returns the index value of the Iptrust and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object emailfilter - iptrust chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectEmailfilterIptrust(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/emailfilter/iptrust"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectEmailfilterIptrust API operation for FortiManager updates the specified Iptrust.
// Returns the index value of the Iptrust and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object emailfilter - iptrust chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectEmailfilterIptrust(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/emailfilter/iptrust"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectEmailfilterIptrust API operation for FortiManager deletes the specified Iptrust.
// Returns error for service API and SDK errors.
// See the object emailfilter - iptrust chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectEmailfilterIptrust(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/emailfilter/iptrust"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectEmailfilterIptrust API operation for FortiManager gets the Iptrust
// with the specified index value.
// Returns the requested Iptrust value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object emailfilter - iptrust chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectEmailfilterIptrust(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/emailfilter/iptrust"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectEmailfilterProfile API operation for FortiManager creates a new Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object emailfilter - profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectEmailfilterProfile(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/emailfilter/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectEmailfilterProfile API operation for FortiManager updates the specified Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object emailfilter - profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectEmailfilterProfile(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/emailfilter/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectEmailfilterProfile API operation for FortiManager deletes the specified Profile.
// Returns error for service API and SDK errors.
// See the object emailfilter - profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectEmailfilterProfile(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/emailfilter/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectEmailfilterProfile API operation for FortiManager gets the Profile
// with the specified index value.
// Returns the requested Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object emailfilter - profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectEmailfilterProfile(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/emailfilter/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateObjectEmailfilterFortishield API operation for FortiManager updates the specified Fortishield.
// Returns the index value of the Fortishield and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object emailfilter - fortishield chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectEmailfilterFortishield(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/emailfilter/fortishield"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectEmailfilterFortishield API operation for FortiManager deletes the specified Fortishield.
// Returns error for service API and SDK errors.
// See the object emailfilter - fortishield chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectEmailfilterFortishield(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for object emailfilter - fortishield
	return
}

// ReadObjectEmailfilterFortishield API operation for FortiManager gets the Fortishield
// with the specified index value.
// Returns the requested Fortishield value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object emailfilter - fortishield chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectEmailfilterFortishield(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/emailfilter/fortishield"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateObjectEmailfilterOptions API operation for FortiManager updates the specified Options.
// Returns the index value of the Options and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object emailfilter - options chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectEmailfilterOptions(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/emailfilter/options"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectEmailfilterOptions API operation for FortiManager deletes the specified Options.
// Returns error for service API and SDK errors.
// See the object emailfilter - options chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectEmailfilterOptions(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for object emailfilter - options
	return
}

// ReadObjectEmailfilterOptions API operation for FortiManager gets the Options
// with the specified index value.
// Returns the requested Options value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object emailfilter - options chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectEmailfilterOptions(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/emailfilter/options"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectUserExchange API operation for FortiManager creates a new Exchange.
// Returns the index value of the Exchange and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - exchange chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectUserExchange(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/exchange"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectUserExchange API operation for FortiManager updates the specified Exchange.
// Returns the index value of the Exchange and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - exchange chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectUserExchange(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/exchange"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectUserExchange API operation for FortiManager deletes the specified Exchange.
// Returns error for service API and SDK errors.
// See the object user - exchange chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectUserExchange(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/user/exchange"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectUserExchange API operation for FortiManager gets the Exchange
// with the specified index value.
// Returns the requested Exchange value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - exchange chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectUserExchange(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/exchange"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectFirewallInternetServiceAddition API operation for FortiManager creates a new Internet Service Addition.
// Returns the index value of the Internet Service Addition and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - internet-service-addition chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectFirewallInternetServiceAddition(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/internet-service-addition"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectFirewallInternetServiceAddition API operation for FortiManager updates the specified Internet Service Addition.
// Returns the index value of the Internet Service Addition and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - internet-service-addition chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectFirewallInternetServiceAddition(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/internet-service-addition"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectFirewallInternetServiceAddition API operation for FortiManager deletes the specified Internet Service Addition.
// Returns error for service API and SDK errors.
// See the object firewall - internet-service-addition chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectFirewallInternetServiceAddition(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/firewall/internet-service-addition"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectFirewallInternetServiceAddition API operation for FortiManager gets the Internet Service Addition
// with the specified index value.
// Returns the requested Internet Service Addition value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - internet-service-addition chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectFirewallInternetServiceAddition(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/internet-service-addition"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectCifsProfile API operation for FortiManager creates a new Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object cifs - profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectCifsProfile(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/cifs/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectCifsProfile API operation for FortiManager updates the specified Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object cifs - profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectCifsProfile(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/cifs/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectCifsProfile API operation for FortiManager deletes the specified Profile.
// Returns error for service API and SDK errors.
// See the object cifs - profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectCifsProfile(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/cifs/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectCifsProfile API operation for FortiManager gets the Profile
// with the specified index value.
// Returns the requested Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object cifs - profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectCifsProfile(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/cifs/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectUserSaml API operation for FortiManager creates a new Saml.
// Returns the index value of the Saml and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - saml chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectUserSaml(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/saml"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectUserSaml API operation for FortiManager updates the specified Saml.
// Returns the index value of the Saml and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - saml chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectUserSaml(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/saml"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectUserSaml API operation for FortiManager deletes the specified Saml.
// Returns error for service API and SDK errors.
// See the object user - saml chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectUserSaml(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/user/saml"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectUserSaml API operation for FortiManager gets the Saml
// with the specified index value.
// Returns the requested Saml value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - saml chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectUserSaml(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/saml"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectFirewallTrafficClass API operation for FortiManager creates a new Traffic Class.
// Returns the index value of the Traffic Class and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - traffic-class chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectFirewallTrafficClass(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/traffic-class"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectFirewallTrafficClass API operation for FortiManager updates the specified Traffic Class.
// Returns the index value of the Traffic Class and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - traffic-class chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectFirewallTrafficClass(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/traffic-class"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectFirewallTrafficClass API operation for FortiManager deletes the specified Traffic Class.
// Returns error for service API and SDK errors.
// See the object firewall - traffic-class chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectFirewallTrafficClass(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/firewall/traffic-class"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectFirewallTrafficClass API operation for FortiManager gets the Traffic Class
// with the specified index value.
// Returns the requested Traffic Class value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - traffic-class chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectFirewallTrafficClass(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/traffic-class"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectWirelessControllerWagProfile API operation for FortiManager creates a new Wag Profile.
// Returns the index value of the Wag Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - wag-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectWirelessControllerWagProfile(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/wag-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectWirelessControllerWagProfile API operation for FortiManager updates the specified Wag Profile.
// Returns the index value of the Wag Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - wag-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectWirelessControllerWagProfile(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/wag-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectWirelessControllerWagProfile API operation for FortiManager deletes the specified Wag Profile.
// Returns error for service API and SDK errors.
// See the object wireless-controller - wag-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectWirelessControllerWagProfile(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/wireless-controller/wag-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectWirelessControllerWagProfile API operation for FortiManager gets the Wag Profile
// with the specified index value.
// Returns the requested Wag Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - wag-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectWirelessControllerWagProfile(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/wag-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectCredentialStoreDomainController API operation for FortiManager creates a new Domain Controller.
// Returns the index value of the Domain Controller and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object credential-store - domain-controller chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectCredentialStoreDomainController(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/credential-store/domain-controller"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectCredentialStoreDomainController API operation for FortiManager updates the specified Domain Controller.
// Returns the index value of the Domain Controller and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object credential-store - domain-controller chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectCredentialStoreDomainController(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/credential-store/domain-controller"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectCredentialStoreDomainController API operation for FortiManager deletes the specified Domain Controller.
// Returns error for service API and SDK errors.
// See the object credential-store - domain-controller chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectCredentialStoreDomainController(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/credential-store/domain-controller"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectCredentialStoreDomainController API operation for FortiManager gets the Domain Controller
// with the specified index value.
// Returns the requested Domain Controller value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object credential-store - domain-controller chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectCredentialStoreDomainController(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/credential-store/domain-controller"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectWirelessControllerApcfgProfile API operation for FortiManager creates a new Apcfg Profile.
// Returns the index value of the Apcfg Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - apcfg-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectWirelessControllerApcfgProfile(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/apcfg-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectWirelessControllerApcfgProfile API operation for FortiManager updates the specified Apcfg Profile.
// Returns the index value of the Apcfg Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - apcfg-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectWirelessControllerApcfgProfile(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/apcfg-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectWirelessControllerApcfgProfile API operation for FortiManager deletes the specified Apcfg Profile.
// Returns error for service API and SDK errors.
// See the object wireless-controller - apcfg-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectWirelessControllerApcfgProfile(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/wireless-controller/apcfg-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectWirelessControllerApcfgProfile API operation for FortiManager gets the Apcfg Profile
// with the specified index value.
// Returns the requested Apcfg Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - apcfg-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectWirelessControllerApcfgProfile(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/apcfg-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectFirewallInternetServiceName API operation for FortiManager creates a new Internet Service Name.
// Returns the index value of the Internet Service Name and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - internet-service-name chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectFirewallInternetServiceName(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/internet-service-name"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectFirewallInternetServiceName API operation for FortiManager updates the specified Internet Service Name.
// Returns the index value of the Internet Service Name and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - internet-service-name chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectFirewallInternetServiceName(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/internet-service-name"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectFirewallInternetServiceName API operation for FortiManager deletes the specified Internet Service Name.
// Returns error for service API and SDK errors.
// See the object firewall - internet-service-name chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectFirewallInternetServiceName(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/firewall/internet-service-name"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectFirewallInternetServiceName API operation for FortiManager gets the Internet Service Name
// with the specified index value.
// Returns the requested Internet Service Name value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - internet-service-name chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectFirewallInternetServiceName(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/internet-service-name"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectFirewallDecryptedTrafficMirror API operation for FortiManager creates a new Decrypted Traffic Mirror.
// Returns the index value of the Decrypted Traffic Mirror and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - decrypted-traffic-mirror chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectFirewallDecryptedTrafficMirror(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/decrypted-traffic-mirror"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectFirewallDecryptedTrafficMirror API operation for FortiManager updates the specified Decrypted Traffic Mirror.
// Returns the index value of the Decrypted Traffic Mirror and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - decrypted-traffic-mirror chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectFirewallDecryptedTrafficMirror(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/decrypted-traffic-mirror"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectFirewallDecryptedTrafficMirror API operation for FortiManager deletes the specified Decrypted Traffic Mirror.
// Returns error for service API and SDK errors.
// See the object firewall - decrypted-traffic-mirror chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectFirewallDecryptedTrafficMirror(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/firewall/decrypted-traffic-mirror"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectFirewallDecryptedTrafficMirror API operation for FortiManager gets the Decrypted Traffic Mirror
// with the specified index value.
// Returns the requested Decrypted Traffic Mirror value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object firewall - decrypted-traffic-mirror chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectFirewallDecryptedTrafficMirror(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/firewall/decrypted-traffic-mirror"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectFileFilterProfile API operation for FortiManager creates a new Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object file-filter - profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectFileFilterProfile(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/file-filter/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectFileFilterProfile API operation for FortiManager updates the specified Profile.
// Returns the index value of the Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object file-filter - profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectFileFilterProfile(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/file-filter/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectFileFilterProfile API operation for FortiManager deletes the specified Profile.
// Returns error for service API and SDK errors.
// See the object file-filter - profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectFileFilterProfile(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/file-filter/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectFileFilterProfile API operation for FortiManager gets the Profile
// with the specified index value.
// Returns the requested Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object file-filter - profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectFileFilterProfile(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/file-filter/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectExtenderControllerDataplan API operation for FortiManager creates a new Dataplan.
// Returns the index value of the Dataplan and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object extender-controller - dataplan chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectExtenderControllerDataplan(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/extender-controller/dataplan"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectExtenderControllerDataplan API operation for FortiManager updates the specified Dataplan.
// Returns the index value of the Dataplan and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object extender-controller - dataplan chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectExtenderControllerDataplan(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/extender-controller/dataplan"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectExtenderControllerDataplan API operation for FortiManager deletes the specified Dataplan.
// Returns error for service API and SDK errors.
// See the object extender-controller - dataplan chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectExtenderControllerDataplan(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/extender-controller/dataplan"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectExtenderControllerDataplan API operation for FortiManager gets the Dataplan
// with the specified index value.
// Returns the requested Dataplan value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object extender-controller - dataplan chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectExtenderControllerDataplan(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/extender-controller/dataplan"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectWirelessControllerMpskProfile API operation for FortiManager creates a new Mpsk Profile.
// Returns the index value of the Mpsk Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - mpsk-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectWirelessControllerMpskProfile(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/mpsk-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectWirelessControllerMpskProfile API operation for FortiManager updates the specified Mpsk Profile.
// Returns the index value of the Mpsk Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - mpsk-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectWirelessControllerMpskProfile(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/mpsk-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectWirelessControllerMpskProfile API operation for FortiManager deletes the specified Mpsk Profile.
// Returns error for service API and SDK errors.
// See the object wireless-controller - mpsk-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectWirelessControllerMpskProfile(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/wireless-controller/mpsk-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectWirelessControllerMpskProfile API operation for FortiManager gets the Mpsk Profile
// with the specified index value.
// Returns the requested Mpsk Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object wireless-controller - mpsk-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectWirelessControllerMpskProfile(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/wireless-controller/mpsk-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectDynamicMulticastInterface API operation for FortiManager creates a new MulticastInterface.
// Returns the index value of the MulticastInterface and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object dynamic - multicast interface chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectDynamicMulticastInterface(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/dynamic/multicast/interface"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectDynamicMulticastInterface API operation for FortiManager updates the specified MulticastInterface.
// Returns the index value of the MulticastInterface and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object dynamic - multicast interface chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectDynamicMulticastInterface(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/dynamic/multicast/interface"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectDynamicMulticastInterface API operation for FortiManager deletes the specified MulticastInterface.
// Returns error for service API and SDK errors.
// See the object dynamic - multicast interface chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectDynamicMulticastInterface(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/dynamic/multicast/interface"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectDynamicMulticastInterface API operation for FortiManager gets the MulticastInterface
// with the specified index value.
// Returns the requested MulticastInterface value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object dynamic - multicast interface chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectDynamicMulticastInterface(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/dynamic/multicast/interface"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectDynamicInterface API operation for FortiManager creates a new Interface.
// Returns the index value of the Interface and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object dynamic - interface chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectDynamicInterface(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/dynamic/interface"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectDynamicInterface API operation for FortiManager updates the specified Interface.
// Returns the index value of the Interface and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object dynamic - interface chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectDynamicInterface(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/dynamic/interface"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectDynamicInterface API operation for FortiManager deletes the specified Interface.
// Returns error for service API and SDK errors.
// See the object dynamic - interface chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectDynamicInterface(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/dynamic/interface"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectDynamicInterface API operation for FortiManager gets the Interface
// with the specified index value.
// Returns the requested Interface value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object dynamic - interface chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectDynamicInterface(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/dynamic/interface"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectDynamicAddress API operation for FortiManager creates a new Address.
// Returns the index value of the Address and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object dynamic - address chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectDynamicAddress(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/dynamic/address"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectDynamicAddress API operation for FortiManager updates the specified Address.
// Returns the index value of the Address and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object dynamic - address chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectDynamicAddress(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/dynamic/address"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectDynamicAddress API operation for FortiManager deletes the specified Address.
// Returns error for service API and SDK errors.
// See the object dynamic - address chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectDynamicAddress(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/dynamic/address"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectDynamicAddress API operation for FortiManager gets the Address
// with the specified index value.
// Returns the requested Address value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object dynamic - address chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectDynamicAddress(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/dynamic/address"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectVpnmgrVpntable API operation for FortiManager creates a new Vpntable.
// Returns the index value of the Vpntable and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object vpnmgr - vpntable chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectVpnmgrVpntable(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/vpnmgr/vpntable"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectVpnmgrVpntable API operation for FortiManager updates the specified Vpntable.
// Returns the index value of the Vpntable and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object vpnmgr - vpntable chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectVpnmgrVpntable(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/vpnmgr/vpntable"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectVpnmgrVpntable API operation for FortiManager deletes the specified Vpntable.
// Returns error for service API and SDK errors.
// See the object vpnmgr - vpntable chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectVpnmgrVpntable(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/vpnmgr/vpntable"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectVpnmgrVpntable API operation for FortiManager gets the Vpntable
// with the specified index value.
// Returns the requested Vpntable value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object vpnmgr - vpntable chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectVpnmgrVpntable(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/vpnmgr/vpntable"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectVpnmgrNode API operation for FortiManager creates a new Node.
// Returns the index value of the Node and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object vpnmgr - node chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectVpnmgrNode(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/vpnmgr/node"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectVpnmgrNode API operation for FortiManager updates the specified Node.
// Returns the index value of the Node and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object vpnmgr - node chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectVpnmgrNode(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/vpnmgr/node"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectVpnmgrNode API operation for FortiManager deletes the specified Node.
// Returns error for service API and SDK errors.
// See the object vpnmgr - node chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectVpnmgrNode(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/vpnmgr/node"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectVpnmgrNode API operation for FortiManager gets the Node
// with the specified index value.
// Returns the requested Node value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object vpnmgr - node chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectVpnmgrNode(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/vpnmgr/node"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectSystemMeta API operation for FortiManager creates a new Meta.
// Returns the index value of the Meta and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object system - meta chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectSystemMeta(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/system/meta"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectSystemMeta API operation for FortiManager updates the specified Meta.
// Returns the index value of the Meta and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object system - meta chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectSystemMeta(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/system/meta"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectSystemMeta API operation for FortiManager deletes the specified Meta.
// Returns error for service API and SDK errors.
// See the object system - meta chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectSystemMeta(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/system/meta"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectSystemMeta API operation for FortiManager gets the Meta
// with the specified index value.
// Returns the requested Meta value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object system - meta chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectSystemMeta(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/system/meta"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateObjectAdomOptions API operation for FortiManager updates the specified Options.
// Returns the index value of the Options and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object adom - options chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectAdomOptions(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/adom/options"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectAdomOptions API operation for FortiManager deletes the specified Options.
// Returns error for service API and SDK errors.
// See the object adom - options chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectAdomOptions(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for object adom - options
	return
}

// ReadObjectAdomOptions API operation for FortiManager gets the Options
// with the specified index value.
// Returns the requested Options value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object adom - options chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectAdomOptions(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/adom/options"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectDynamicVip API operation for FortiManager creates a new Vip.
// Returns the index value of the Vip and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object dynamic - vip chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectDynamicVip(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/dynamic/vip"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectDynamicVip API operation for FortiManager updates the specified Vip.
// Returns the index value of the Vip and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object dynamic - vip chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectDynamicVip(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/dynamic/vip"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectDynamicVip API operation for FortiManager deletes the specified Vip.
// Returns error for service API and SDK errors.
// See the object dynamic - vip chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectDynamicVip(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/dynamic/vip"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectDynamicVip API operation for FortiManager gets the Vip
// with the specified index value.
// Returns the requested Vip value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object dynamic - vip chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectDynamicVip(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/dynamic/vip"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectDynamicIppool API operation for FortiManager creates a new Ippool.
// Returns the index value of the Ippool and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object dynamic - ippool chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectDynamicIppool(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/dynamic/ippool"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectDynamicIppool API operation for FortiManager updates the specified Ippool.
// Returns the index value of the Ippool and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object dynamic - ippool chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectDynamicIppool(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/dynamic/ippool"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectDynamicIppool API operation for FortiManager deletes the specified Ippool.
// Returns error for service API and SDK errors.
// See the object dynamic - ippool chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectDynamicIppool(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/dynamic/ippool"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectDynamicIppool API operation for FortiManager gets the Ippool
// with the specified index value.
// Returns the requested Ippool value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object dynamic - ippool chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectDynamicIppool(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/dynamic/ippool"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectDynamicCertificateLocal API operation for FortiManager creates a new CertificateLocal.
// Returns the index value of the CertificateLocal and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object dynamic - certificate local chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectDynamicCertificateLocal(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/dynamic/certificate/local"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectDynamicCertificateLocal API operation for FortiManager updates the specified CertificateLocal.
// Returns the index value of the CertificateLocal and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object dynamic - certificate local chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectDynamicCertificateLocal(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/dynamic/certificate/local"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectDynamicCertificateLocal API operation for FortiManager deletes the specified CertificateLocal.
// Returns error for service API and SDK errors.
// See the object dynamic - certificate local chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectDynamicCertificateLocal(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/dynamic/certificate/local"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectDynamicCertificateLocal API operation for FortiManager gets the CertificateLocal
// with the specified index value.
// Returns the requested CertificateLocal value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object dynamic - certificate local chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectDynamicCertificateLocal(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/dynamic/certificate/local"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectDynamicVpntunnel API operation for FortiManager creates a new Vpntunnel.
// Returns the index value of the Vpntunnel and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object dynamic - vpntunnel chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectDynamicVpntunnel(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/dynamic/vpntunnel"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectDynamicVpntunnel API operation for FortiManager updates the specified Vpntunnel.
// Returns the index value of the Vpntunnel and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object dynamic - vpntunnel chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectDynamicVpntunnel(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/dynamic/vpntunnel"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectDynamicVpntunnel API operation for FortiManager deletes the specified Vpntunnel.
// Returns error for service API and SDK errors.
// See the object dynamic - vpntunnel chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectDynamicVpntunnel(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/dynamic/vpntunnel"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectDynamicVpntunnel API operation for FortiManager gets the Vpntunnel
// with the specified index value.
// Returns the requested Vpntunnel value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object dynamic - vpntunnel chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectDynamicVpntunnel(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/dynamic/vpntunnel"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectCertificateTemplate API operation for FortiManager creates a new Template.
// Returns the index value of the Template and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object certificate - template chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectCertificateTemplate(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/certificate/template"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectCertificateTemplate API operation for FortiManager updates the specified Template.
// Returns the index value of the Template and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object certificate - template chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectCertificateTemplate(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/certificate/template"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectCertificateTemplate API operation for FortiManager deletes the specified Template.
// Returns error for service API and SDK errors.
// See the object certificate - template chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectCertificateTemplate(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/certificate/template"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectCertificateTemplate API operation for FortiManager gets the Template
// with the specified index value.
// Returns the requested Template value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object certificate - template chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectCertificateTemplate(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/certificate/template"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectDynamicVirtualWanLinkMembers API operation for FortiManager creates a new Virtual Wan LinkMembers.
// Returns the index value of the Virtual Wan LinkMembers and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object dynamic - virtual-wan-link members chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectDynamicVirtualWanLinkMembers(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/dynamic/virtual-wan-link/members"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectDynamicVirtualWanLinkMembers API operation for FortiManager updates the specified Virtual Wan LinkMembers.
// Returns the index value of the Virtual Wan LinkMembers and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object dynamic - virtual-wan-link members chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectDynamicVirtualWanLinkMembers(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/dynamic/virtual-wan-link/members"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectDynamicVirtualWanLinkMembers API operation for FortiManager deletes the specified Virtual Wan LinkMembers.
// Returns error for service API and SDK errors.
// See the object dynamic - virtual-wan-link members chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectDynamicVirtualWanLinkMembers(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/dynamic/virtual-wan-link/members"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectDynamicVirtualWanLinkMembers API operation for FortiManager gets the Virtual Wan LinkMembers
// with the specified index value.
// Returns the requested Virtual Wan LinkMembers value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object dynamic - virtual-wan-link members chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectDynamicVirtualWanLinkMembers(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/dynamic/virtual-wan-link/members"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectDynamicVirtualWanLinkServer API operation for FortiManager creates a new Virtual Wan LinkServer.
// Returns the index value of the Virtual Wan LinkServer and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object dynamic - virtual-wan-link server chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectDynamicVirtualWanLinkServer(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/dynamic/virtual-wan-link/server"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectDynamicVirtualWanLinkServer API operation for FortiManager updates the specified Virtual Wan LinkServer.
// Returns the index value of the Virtual Wan LinkServer and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object dynamic - virtual-wan-link server chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectDynamicVirtualWanLinkServer(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/dynamic/virtual-wan-link/server"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectDynamicVirtualWanLinkServer API operation for FortiManager deletes the specified Virtual Wan LinkServer.
// Returns error for service API and SDK errors.
// See the object dynamic - virtual-wan-link server chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectDynamicVirtualWanLinkServer(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/dynamic/virtual-wan-link/server"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectDynamicVirtualWanLinkServer API operation for FortiManager gets the Virtual Wan LinkServer
// with the specified index value.
// Returns the requested Virtual Wan LinkServer value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object dynamic - virtual-wan-link server chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectDynamicVirtualWanLinkServer(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/dynamic/virtual-wan-link/server"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectDynamicVirtualWanLinkNeighbor API operation for FortiManager creates a new Virtual Wan LinkNeighbor.
// Returns the index value of the Virtual Wan LinkNeighbor and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object dynamic - virtual-wan-link neighbor chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectDynamicVirtualWanLinkNeighbor(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/dynamic/virtual-wan-link/neighbor"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectDynamicVirtualWanLinkNeighbor API operation for FortiManager updates the specified Virtual Wan LinkNeighbor.
// Returns the index value of the Virtual Wan LinkNeighbor and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object dynamic - virtual-wan-link neighbor chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectDynamicVirtualWanLinkNeighbor(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/dynamic/virtual-wan-link/neighbor"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectDynamicVirtualWanLinkNeighbor API operation for FortiManager deletes the specified Virtual Wan LinkNeighbor.
// Returns error for service API and SDK errors.
// See the object dynamic - virtual-wan-link neighbor chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectDynamicVirtualWanLinkNeighbor(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/dynamic/virtual-wan-link/neighbor"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectDynamicVirtualWanLinkNeighbor API operation for FortiManager gets the Virtual Wan LinkNeighbor
// with the specified index value.
// Returns the requested Virtual Wan LinkNeighbor value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object dynamic - virtual-wan-link neighbor chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectDynamicVirtualWanLinkNeighbor(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/dynamic/virtual-wan-link/neighbor"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectUserPxgrid API operation for FortiManager creates a new Pxgrid.
// Returns the index value of the Pxgrid and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - pxgrid chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectUserPxgrid(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/pxgrid"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectUserPxgrid API operation for FortiManager updates the specified Pxgrid.
// Returns the index value of the Pxgrid and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - pxgrid chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectUserPxgrid(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/pxgrid"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectUserPxgrid API operation for FortiManager deletes the specified Pxgrid.
// Returns error for service API and SDK errors.
// See the object user - pxgrid chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectUserPxgrid(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/user/pxgrid"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectUserPxgrid API operation for FortiManager gets the Pxgrid
// with the specified index value.
// Returns the requested Pxgrid value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - pxgrid chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectUserPxgrid(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/pxgrid"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectUserClearpass API operation for FortiManager creates a new Clearpass.
// Returns the index value of the Clearpass and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - clearpass chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectUserClearpass(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/clearpass"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectUserClearpass API operation for FortiManager updates the specified Clearpass.
// Returns the index value of the Clearpass and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - clearpass chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectUserClearpass(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/clearpass"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectUserClearpass API operation for FortiManager deletes the specified Clearpass.
// Returns error for service API and SDK errors.
// See the object user - clearpass chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectUserClearpass(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/user/clearpass"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectUserClearpass API operation for FortiManager gets the Clearpass
// with the specified index value.
// Returns the requested Clearpass value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - clearpass chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectUserClearpass(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/clearpass"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectUserNsx API operation for FortiManager creates a new Nsx.
// Returns the index value of the Nsx and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - nsx chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectUserNsx(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/nsx"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectUserNsx API operation for FortiManager updates the specified Nsx.
// Returns the index value of the Nsx and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - nsx chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectUserNsx(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/nsx"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectUserNsx API operation for FortiManager deletes the specified Nsx.
// Returns error for service API and SDK errors.
// See the object user - nsx chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectUserNsx(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/user/nsx"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectUserNsx API operation for FortiManager gets the Nsx
// with the specified index value.
// Returns the requested Nsx value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - nsx chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectUserNsx(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/nsx"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectUserVcenter API operation for FortiManager creates a new Vcenter.
// Returns the index value of the Vcenter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - vcenter chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectUserVcenter(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/vcenter"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectUserVcenter API operation for FortiManager updates the specified Vcenter.
// Returns the index value of the Vcenter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - vcenter chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectUserVcenter(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/vcenter"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectUserVcenter API operation for FortiManager deletes the specified Vcenter.
// Returns error for service API and SDK errors.
// See the object user - vcenter chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectUserVcenter(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/user/vcenter"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectUserVcenter API operation for FortiManager gets the Vcenter
// with the specified index value.
// Returns the requested Vcenter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object user - vcenter chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectUserVcenter(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/user/vcenter"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectCliTemplate API operation for FortiManager creates a new Template.
// Returns the index value of the Template and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object cli - template chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectCliTemplate(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/cli/template"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectCliTemplate API operation for FortiManager updates the specified Template.
// Returns the index value of the Template and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object cli - template chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectCliTemplate(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/cli/template"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectCliTemplate API operation for FortiManager deletes the specified Template.
// Returns error for service API and SDK errors.
// See the object cli - template chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectCliTemplate(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/cli/template"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectCliTemplate API operation for FortiManager gets the Template
// with the specified index value.
// Returns the requested Template value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object cli - template chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectCliTemplate(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/cli/template"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectCliTemplateGroup API operation for FortiManager creates a new Template Group.
// Returns the index value of the Template Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object cli - template-group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectCliTemplateGroup(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/cli/template-group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectCliTemplateGroup API operation for FortiManager updates the specified Template Group.
// Returns the index value of the Template Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object cli - template-group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectCliTemplateGroup(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/cli/template-group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectCliTemplateGroup API operation for FortiManager deletes the specified Template Group.
// Returns error for service API and SDK errors.
// See the object cli - template-group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectCliTemplateGroup(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/cli/template-group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectCliTemplateGroup API operation for FortiManager gets the Template Group
// with the specified index value.
// Returns the requested Template Group value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object cli - template-group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectCliTemplateGroup(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/cli/template-group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateObjectFspVlan API operation for FortiManager creates a new Vlan.
// Returns the index value of the Vlan and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object fsp - vlan chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateObjectFspVlan(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/fsp/vlan"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateObjectFspVlan API operation for FortiManager updates the specified Vlan.
// Returns the index value of the Vlan and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object fsp - vlan chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateObjectFspVlan(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/fsp/vlan"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteObjectFspVlan API operation for FortiManager deletes the specified Vlan.
// Returns error for service API and SDK errors.
// See the object fsp - vlan chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteObjectFspVlan(globaladom, mkey string, paralist []string) (err error) {
	path := "/pm/config/[*]/obj/fsp/vlan"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadObjectFspVlan API operation for FortiManager gets the Vlan
// with the specified index value.
// Returns the requested Vlan value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the object fsp - vlan chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadObjectFspVlan(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/pm/config/[*]/obj/fsp/vlan"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemGlobal API operation for FortiManager updates the specified Global.
// Returns the index value of the Global and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - global chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemGlobal(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/global"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemGlobal API operation for FortiManager deletes the specified Global.
// Returns error for service API and SDK errors.
// See the system - global chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemGlobal(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - global
	return
}

// ReadSystemGlobal API operation for FortiManager gets the Global
// with the specified index value.
// Returns the requested Global value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - global chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemGlobal(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/global"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemGlobalMcPolicyDisabledAdoms API operation for FortiManager creates a new GlobalMc Policy Disabled Adoms.
// Returns the index value of the GlobalMc Policy Disabled Adoms and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - global mc-policy-disabled-adoms chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemGlobalMcPolicyDisabledAdoms(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/global/mc-policy-disabled-adoms"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateSystemGlobalMcPolicyDisabledAdoms API operation for FortiManager updates the specified GlobalMc Policy Disabled Adoms.
// Returns the index value of the GlobalMc Policy Disabled Adoms and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - global mc-policy-disabled-adoms chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemGlobalMcPolicyDisabledAdoms(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/global/mc-policy-disabled-adoms"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemGlobalMcPolicyDisabledAdoms API operation for FortiManager deletes the specified GlobalMc Policy Disabled Adoms.
// Returns error for service API and SDK errors.
// See the system - global mc-policy-disabled-adoms chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemGlobalMcPolicyDisabledAdoms(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/global/mc-policy-disabled-adoms"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemGlobalMcPolicyDisabledAdoms API operation for FortiManager gets the GlobalMc Policy Disabled Adoms
// with the specified index value.
// Returns the requested GlobalMc Policy Disabled Adoms value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - global mc-policy-disabled-adoms chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemGlobalMcPolicyDisabledAdoms(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/global/mc-policy-disabled-adoms"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemInterface API operation for FortiManager creates a new Interface.
// Returns the index value of the Interface and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - interface chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemInterface(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/interface"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateSystemInterface API operation for FortiManager updates the specified Interface.
// Returns the index value of the Interface and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - interface chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemInterface(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/interface"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemInterface API operation for FortiManager deletes the specified Interface.
// Returns error for service API and SDK errors.
// See the system - interface chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemInterface(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/interface"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemInterface API operation for FortiManager gets the Interface
// with the specified index value.
// Returns the requested Interface value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - interface chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemInterface(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/interface"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemSnmpSysinfo API operation for FortiManager updates the specified SnmpSysinfo.
// Returns the index value of the SnmpSysinfo and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - snmp sysinfo chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSnmpSysinfo(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/snmp/sysinfo"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemSnmpSysinfo API operation for FortiManager deletes the specified SnmpSysinfo.
// Returns error for service API and SDK errors.
// See the system - snmp sysinfo chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSnmpSysinfo(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - snmp sysinfo
	return
}

// ReadSystemSnmpSysinfo API operation for FortiManager gets the SnmpSysinfo
// with the specified index value.
// Returns the requested SnmpSysinfo value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - snmp sysinfo chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSnmpSysinfo(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/snmp/sysinfo"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemSnmpCommunity API operation for FortiManager creates a new SnmpCommunity.
// Returns the index value of the SnmpCommunity and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - snmp community chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemSnmpCommunity(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/snmp/community"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateSystemSnmpCommunity API operation for FortiManager updates the specified SnmpCommunity.
// Returns the index value of the SnmpCommunity and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - snmp community chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSnmpCommunity(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/snmp/community"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemSnmpCommunity API operation for FortiManager deletes the specified SnmpCommunity.
// Returns error for service API and SDK errors.
// See the system - snmp community chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSnmpCommunity(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/snmp/community"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemSnmpCommunity API operation for FortiManager gets the SnmpCommunity
// with the specified index value.
// Returns the requested SnmpCommunity value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - snmp community chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSnmpCommunity(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/snmp/community"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemSnmpUser API operation for FortiManager creates a new SnmpUser.
// Returns the index value of the SnmpUser and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - snmp user chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemSnmpUser(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/snmp/user"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateSystemSnmpUser API operation for FortiManager updates the specified SnmpUser.
// Returns the index value of the SnmpUser and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - snmp user chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSnmpUser(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/snmp/user"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemSnmpUser API operation for FortiManager deletes the specified SnmpUser.
// Returns error for service API and SDK errors.
// See the system - snmp user chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSnmpUser(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/snmp/user"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemSnmpUser API operation for FortiManager gets the SnmpUser
// with the specified index value.
// Returns the requested SnmpUser value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - snmp user chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSnmpUser(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/snmp/user"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemRoute API operation for FortiManager creates a new Route.
// Returns the index value of the Route and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - route chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemRoute(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/route"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateSystemRoute API operation for FortiManager updates the specified Route.
// Returns the index value of the Route and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - route chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemRoute(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/route"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemRoute API operation for FortiManager deletes the specified Route.
// Returns error for service API and SDK errors.
// See the system - route chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemRoute(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/route"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemRoute API operation for FortiManager gets the Route
// with the specified index value.
// Returns the requested Route value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - route chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemRoute(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/route"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemRoute6 API operation for FortiManager creates a new Route6.
// Returns the index value of the Route6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - route6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemRoute6(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/route6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateSystemRoute6 API operation for FortiManager updates the specified Route6.
// Returns the index value of the Route6 and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - route6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemRoute6(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/route6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemRoute6 API operation for FortiManager deletes the specified Route6.
// Returns error for service API and SDK errors.
// See the system - route6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemRoute6(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/route6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemRoute6 API operation for FortiManager gets the Route6
// with the specified index value.
// Returns the requested Route6 value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - route6 chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemRoute6(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/route6"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemDns API operation for FortiManager updates the specified Dns.
// Returns the index value of the Dns and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - dns chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemDns(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/dns"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemDns API operation for FortiManager deletes the specified Dns.
// Returns error for service API and SDK errors.
// See the system - dns chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemDns(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - dns
	return
}

// ReadSystemDns API operation for FortiManager gets the Dns
// with the specified index value.
// Returns the requested Dns value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - dns chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemDns(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/dns"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemConnector API operation for FortiManager updates the specified Connector.
// Returns the index value of the Connector and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - connector chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemConnector(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/connector"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemConnector API operation for FortiManager deletes the specified Connector.
// Returns error for service API and SDK errors.
// See the system - connector chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemConnector(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - connector
	return
}

// ReadSystemConnector API operation for FortiManager gets the Connector
// with the specified index value.
// Returns the requested Connector value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - connector chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemConnector(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/connector"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemHa API operation for FortiManager updates the specified Ha.
// Returns the index value of the Ha and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ha chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemHa(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/ha"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemHa API operation for FortiManager deletes the specified Ha.
// Returns error for service API and SDK errors.
// See the system - ha chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemHa(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - ha
	return
}

// ReadSystemHa API operation for FortiManager gets the Ha
// with the specified index value.
// Returns the requested Ha value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ha chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemHa(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/ha"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemHaPeer API operation for FortiManager creates a new HaPeer.
// Returns the index value of the HaPeer and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ha peer chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemHaPeer(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/ha/peer"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateSystemHaPeer API operation for FortiManager updates the specified HaPeer.
// Returns the index value of the HaPeer and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ha peer chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemHaPeer(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/ha/peer"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemHaPeer API operation for FortiManager deletes the specified HaPeer.
// Returns error for service API and SDK errors.
// See the system - ha peer chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemHaPeer(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/ha/peer"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemHaPeer API operation for FortiManager gets the HaPeer
// with the specified index value.
// Returns the requested HaPeer value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ha peer chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemHaPeer(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/ha/peer"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemCertificateCa API operation for FortiManager creates a new CertificateCa.
// Returns the index value of the CertificateCa and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - certificate ca chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemCertificateCa(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/certificate/ca"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateSystemCertificateCa API operation for FortiManager updates the specified CertificateCa.
// Returns the index value of the CertificateCa and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - certificate ca chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemCertificateCa(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/certificate/ca"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemCertificateCa API operation for FortiManager deletes the specified CertificateCa.
// Returns error for service API and SDK errors.
// See the system - certificate ca chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemCertificateCa(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/certificate/ca"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemCertificateCa API operation for FortiManager gets the CertificateCa
// with the specified index value.
// Returns the requested CertificateCa value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - certificate ca chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemCertificateCa(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/certificate/ca"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemCertificateLocal API operation for FortiManager creates a new CertificateLocal.
// Returns the index value of the CertificateLocal and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - certificate local chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemCertificateLocal(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/certificate/local"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateSystemCertificateLocal API operation for FortiManager updates the specified CertificateLocal.
// Returns the index value of the CertificateLocal and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - certificate local chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemCertificateLocal(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/certificate/local"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemCertificateLocal API operation for FortiManager deletes the specified CertificateLocal.
// Returns error for service API and SDK errors.
// See the system - certificate local chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemCertificateLocal(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/certificate/local"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemCertificateLocal API operation for FortiManager gets the CertificateLocal
// with the specified index value.
// Returns the requested CertificateLocal value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - certificate local chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemCertificateLocal(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/certificate/local"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemCertificateCrl API operation for FortiManager creates a new CertificateCrl.
// Returns the index value of the CertificateCrl and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - certificate crl chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemCertificateCrl(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/certificate/crl"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateSystemCertificateCrl API operation for FortiManager updates the specified CertificateCrl.
// Returns the index value of the CertificateCrl and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - certificate crl chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemCertificateCrl(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/certificate/crl"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemCertificateCrl API operation for FortiManager deletes the specified CertificateCrl.
// Returns error for service API and SDK errors.
// See the system - certificate crl chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemCertificateCrl(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/certificate/crl"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemCertificateCrl API operation for FortiManager gets the CertificateCrl
// with the specified index value.
// Returns the requested CertificateCrl value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - certificate crl chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemCertificateCrl(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/certificate/crl"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemCertificateRemote API operation for FortiManager creates a new CertificateRemote.
// Returns the index value of the CertificateRemote and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - certificate remote chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemCertificateRemote(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/certificate/remote"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateSystemCertificateRemote API operation for FortiManager updates the specified CertificateRemote.
// Returns the index value of the CertificateRemote and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - certificate remote chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemCertificateRemote(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/certificate/remote"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemCertificateRemote API operation for FortiManager deletes the specified CertificateRemote.
// Returns error for service API and SDK errors.
// See the system - certificate remote chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemCertificateRemote(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/certificate/remote"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemCertificateRemote API operation for FortiManager gets the CertificateRemote
// with the specified index value.
// Returns the requested CertificateRemote value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - certificate remote chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemCertificateRemote(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/certificate/remote"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemCertificateOftp API operation for FortiManager updates the specified CertificateOftp.
// Returns the index value of the CertificateOftp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - certificate oftp chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemCertificateOftp(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/certificate/oftp"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemCertificateOftp API operation for FortiManager deletes the specified CertificateOftp.
// Returns error for service API and SDK errors.
// See the system - certificate oftp chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemCertificateOftp(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - certificate oftp
	return
}

// ReadSystemCertificateOftp API operation for FortiManager gets the CertificateOftp
// with the specified index value.
// Returns the requested CertificateOftp value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - certificate oftp chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemCertificateOftp(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/certificate/oftp"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemCertificateSsh API operation for FortiManager creates a new CertificateSsh.
// Returns the index value of the CertificateSsh and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - certificate ssh chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemCertificateSsh(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/certificate/ssh"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateSystemCertificateSsh API operation for FortiManager updates the specified CertificateSsh.
// Returns the index value of the CertificateSsh and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - certificate ssh chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemCertificateSsh(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/certificate/ssh"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemCertificateSsh API operation for FortiManager deletes the specified CertificateSsh.
// Returns error for service API and SDK errors.
// See the system - certificate ssh chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemCertificateSsh(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/certificate/ssh"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemCertificateSsh API operation for FortiManager gets the CertificateSsh
// with the specified index value.
// Returns the requested CertificateSsh value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - certificate ssh chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemCertificateSsh(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/certificate/ssh"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemSaml API operation for FortiManager updates the specified Saml.
// Returns the index value of the Saml and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - saml chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSaml(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/saml"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemSaml API operation for FortiManager deletes the specified Saml.
// Returns error for service API and SDK errors.
// See the system - saml chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSaml(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - saml
	return
}

// ReadSystemSaml API operation for FortiManager gets the Saml
// with the specified index value.
// Returns the requested Saml value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - saml chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSaml(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/saml"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemSamlFabricIdp API operation for FortiManager creates a new SamlFabric Idp.
// Returns the index value of the SamlFabric Idp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - saml fabric-idp chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemSamlFabricIdp(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/saml/fabric-idp"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateSystemSamlFabricIdp API operation for FortiManager updates the specified SamlFabric Idp.
// Returns the index value of the SamlFabric Idp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - saml fabric-idp chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSamlFabricIdp(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/saml/fabric-idp"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemSamlFabricIdp API operation for FortiManager deletes the specified SamlFabric Idp.
// Returns error for service API and SDK errors.
// See the system - saml fabric-idp chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSamlFabricIdp(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/saml/fabric-idp"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemSamlFabricIdp API operation for FortiManager gets the SamlFabric Idp
// with the specified index value.
// Returns the requested SamlFabric Idp value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - saml fabric-idp chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSamlFabricIdp(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/saml/fabric-idp"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemSamlServiceProviders API operation for FortiManager creates a new SamlService Providers.
// Returns the index value of the SamlService Providers and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - saml service-providers chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemSamlServiceProviders(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/saml/service-providers"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateSystemSamlServiceProviders API operation for FortiManager updates the specified SamlService Providers.
// Returns the index value of the SamlService Providers and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - saml service-providers chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSamlServiceProviders(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/saml/service-providers"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemSamlServiceProviders API operation for FortiManager deletes the specified SamlService Providers.
// Returns error for service API and SDK errors.
// See the system - saml service-providers chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSamlServiceProviders(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/saml/service-providers"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemSamlServiceProviders API operation for FortiManager gets the SamlService Providers
// with the specified index value.
// Returns the requested SamlService Providers value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - saml service-providers chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSamlServiceProviders(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/saml/service-providers"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemNtp API operation for FortiManager updates the specified Ntp.
// Returns the index value of the Ntp and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ntp chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemNtp(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/ntp"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemNtp API operation for FortiManager deletes the specified Ntp.
// Returns error for service API and SDK errors.
// See the system - ntp chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemNtp(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - ntp
	return
}

// ReadSystemNtp API operation for FortiManager gets the Ntp
// with the specified index value.
// Returns the requested Ntp value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ntp chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemNtp(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/ntp"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemNtpNtpserver API operation for FortiManager creates a new NtpNtpserver.
// Returns the index value of the NtpNtpserver and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ntp ntpserver chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemNtpNtpserver(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/ntp/ntpserver"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateSystemNtpNtpserver API operation for FortiManager updates the specified NtpNtpserver.
// Returns the index value of the NtpNtpserver and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ntp ntpserver chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemNtpNtpserver(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/ntp/ntpserver"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemNtpNtpserver API operation for FortiManager deletes the specified NtpNtpserver.
// Returns error for service API and SDK errors.
// See the system - ntp ntpserver chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemNtpNtpserver(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/ntp/ntpserver"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemNtpNtpserver API operation for FortiManager gets the NtpNtpserver
// with the specified index value.
// Returns the requested NtpNtpserver value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - ntp ntpserver chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemNtpNtpserver(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/ntp/ntpserver"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemBackupAllSettings API operation for FortiManager updates the specified BackupAll Settings.
// Returns the index value of the BackupAll Settings and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - backup all-settings chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemBackupAllSettings(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/backup/all-settings"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemBackupAllSettings API operation for FortiManager deletes the specified BackupAll Settings.
// Returns error for service API and SDK errors.
// See the system - backup all-settings chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemBackupAllSettings(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - backup all-settings
	return
}

// ReadSystemBackupAllSettings API operation for FortiManager gets the BackupAll Settings
// with the specified index value.
// Returns the requested BackupAll Settings value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - backup all-settings chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemBackupAllSettings(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/backup/all-settings"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemGuiact API operation for FortiManager updates the specified Guiact.
// Returns the index value of the Guiact and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - guiact chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemGuiact(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/guiact"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemGuiact API operation for FortiManager deletes the specified Guiact.
// Returns error for service API and SDK errors.
// See the system - guiact chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemGuiact(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - guiact
	return
}

// ReadSystemGuiact API operation for FortiManager gets the Guiact
// with the specified index value.
// Returns the requested Guiact value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - guiact chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemGuiact(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/guiact"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemMetadataAdmins API operation for FortiManager creates a new MetadataAdmins.
// Returns the index value of the MetadataAdmins and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - metadata admins chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemMetadataAdmins(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/metadata/admins"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateSystemMetadataAdmins API operation for FortiManager updates the specified MetadataAdmins.
// Returns the index value of the MetadataAdmins and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - metadata admins chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemMetadataAdmins(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/metadata/admins"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemMetadataAdmins API operation for FortiManager deletes the specified MetadataAdmins.
// Returns error for service API and SDK errors.
// See the system - metadata admins chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemMetadataAdmins(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/metadata/admins"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemMetadataAdmins API operation for FortiManager gets the MetadataAdmins
// with the specified index value.
// Returns the requested MetadataAdmins value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - metadata admins chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemMetadataAdmins(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/metadata/admins"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemAdminProfile API operation for FortiManager creates a new AdminProfile.
// Returns the index value of the AdminProfile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemAdminProfile(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/admin/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateSystemAdminProfile API operation for FortiManager updates the specified AdminProfile.
// Returns the index value of the AdminProfile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAdminProfile(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/admin/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemAdminProfile API operation for FortiManager deletes the specified AdminProfile.
// Returns error for service API and SDK errors.
// See the system - admin profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAdminProfile(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/admin/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemAdminProfile API operation for FortiManager gets the AdminProfile
// with the specified index value.
// Returns the requested AdminProfile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAdminProfile(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/admin/profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemAdminRadius API operation for FortiManager creates a new AdminRadius.
// Returns the index value of the AdminRadius and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin radius chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemAdminRadius(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/admin/radius"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateSystemAdminRadius API operation for FortiManager updates the specified AdminRadius.
// Returns the index value of the AdminRadius and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin radius chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAdminRadius(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/admin/radius"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemAdminRadius API operation for FortiManager deletes the specified AdminRadius.
// Returns error for service API and SDK errors.
// See the system - admin radius chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAdminRadius(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/admin/radius"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemAdminRadius API operation for FortiManager gets the AdminRadius
// with the specified index value.
// Returns the requested AdminRadius value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin radius chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAdminRadius(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/admin/radius"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemAdminLdap API operation for FortiManager creates a new AdminLdap.
// Returns the index value of the AdminLdap and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin ldap chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemAdminLdap(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/admin/ldap"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateSystemAdminLdap API operation for FortiManager updates the specified AdminLdap.
// Returns the index value of the AdminLdap and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin ldap chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAdminLdap(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/admin/ldap"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemAdminLdap API operation for FortiManager deletes the specified AdminLdap.
// Returns error for service API and SDK errors.
// See the system - admin ldap chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAdminLdap(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/admin/ldap"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemAdminLdap API operation for FortiManager gets the AdminLdap
// with the specified index value.
// Returns the requested AdminLdap value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin ldap chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAdminLdap(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/admin/ldap"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemAdminTacacs API operation for FortiManager creates a new AdminTacacs.
// Returns the index value of the AdminTacacs and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin tacacs chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemAdminTacacs(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/admin/tacacs"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateSystemAdminTacacs API operation for FortiManager updates the specified AdminTacacs.
// Returns the index value of the AdminTacacs and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin tacacs chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAdminTacacs(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/admin/tacacs"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemAdminTacacs API operation for FortiManager deletes the specified AdminTacacs.
// Returns error for service API and SDK errors.
// See the system - admin tacacs chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAdminTacacs(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/admin/tacacs"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemAdminTacacs API operation for FortiManager gets the AdminTacacs
// with the specified index value.
// Returns the requested AdminTacacs value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin tacacs chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAdminTacacs(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/admin/tacacs"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemAdminGroup API operation for FortiManager creates a new AdminGroup.
// Returns the index value of the AdminGroup and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemAdminGroup(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/admin/group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateSystemAdminGroup API operation for FortiManager updates the specified AdminGroup.
// Returns the index value of the AdminGroup and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAdminGroup(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/admin/group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemAdminGroup API operation for FortiManager deletes the specified AdminGroup.
// Returns error for service API and SDK errors.
// See the system - admin group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAdminGroup(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/admin/group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemAdminGroup API operation for FortiManager gets the AdminGroup
// with the specified index value.
// Returns the requested AdminGroup value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAdminGroup(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/admin/group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemAdminUser API operation for FortiManager creates a new AdminUser.
// Returns the index value of the AdminUser and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin user chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemAdminUser(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/admin/user"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateSystemAdminUser API operation for FortiManager updates the specified AdminUser.
// Returns the index value of the AdminUser and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin user chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAdminUser(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/admin/user"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemAdminUser API operation for FortiManager deletes the specified AdminUser.
// Returns error for service API and SDK errors.
// See the system - admin user chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAdminUser(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/admin/user"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemAdminUser API operation for FortiManager gets the AdminUser
// with the specified index value.
// Returns the requested AdminUser value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin user chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAdminUser(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/admin/user"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemAdminSetting API operation for FortiManager updates the specified AdminSetting.
// Returns the index value of the AdminSetting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin setting chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAdminSetting(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/admin/setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemAdminSetting API operation for FortiManager deletes the specified AdminSetting.
// Returns error for service API and SDK errors.
// See the system - admin setting chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAdminSetting(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - admin setting
	return
}

// ReadSystemAdminSetting API operation for FortiManager gets the AdminSetting
// with the specified index value.
// Returns the requested AdminSetting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - admin setting chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAdminSetting(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/admin/setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemPasswordPolicy API operation for FortiManager updates the specified Password Policy.
// Returns the index value of the Password Policy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - password-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemPasswordPolicy(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/password-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemPasswordPolicy API operation for FortiManager deletes the specified Password Policy.
// Returns error for service API and SDK errors.
// See the system - password-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemPasswordPolicy(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - password-policy
	return
}

// ReadSystemPasswordPolicy API operation for FortiManager gets the Password Policy
// with the specified index value.
// Returns the requested Password Policy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - password-policy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemPasswordPolicy(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/password-policy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemAlertemail API operation for FortiManager updates the specified Alertemail.
// Returns the index value of the Alertemail and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - alertemail chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAlertemail(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/alertemail"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemAlertemail API operation for FortiManager deletes the specified Alertemail.
// Returns error for service API and SDK errors.
// See the system - alertemail chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAlertemail(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - alertemail
	return
}

// ReadSystemAlertemail API operation for FortiManager gets the Alertemail
// with the specified index value.
// Returns the requested Alertemail value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - alertemail chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAlertemail(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/alertemail"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemSyslog API operation for FortiManager creates a new Syslog.
// Returns the index value of the Syslog and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - syslog chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemSyslog(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/syslog"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateSystemSyslog API operation for FortiManager updates the specified Syslog.
// Returns the index value of the Syslog and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - syslog chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSyslog(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/syslog"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemSyslog API operation for FortiManager deletes the specified Syslog.
// Returns error for service API and SDK errors.
// See the system - syslog chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSyslog(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/syslog"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemSyslog API operation for FortiManager gets the Syslog
// with the specified index value.
// Returns the requested Syslog value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - syslog chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSyslog(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/syslog"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemMail API operation for FortiManager creates a new Mail.
// Returns the index value of the Mail and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - mail chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemMail(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/mail"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateSystemMail API operation for FortiManager updates the specified Mail.
// Returns the index value of the Mail and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - mail chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemMail(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/mail"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemMail API operation for FortiManager deletes the specified Mail.
// Returns error for service API and SDK errors.
// See the system - mail chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemMail(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/mail"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemMail API operation for FortiManager gets the Mail
// with the specified index value.
// Returns the requested Mail value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - mail chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemMail(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/mail"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemAlertEvent API operation for FortiManager creates a new Alert Event.
// Returns the index value of the Alert Event and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - alert-event chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemAlertEvent(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/alert-event"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateSystemAlertEvent API operation for FortiManager updates the specified Alert Event.
// Returns the index value of the Alert Event and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - alert-event chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAlertEvent(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/alert-event"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemAlertEvent API operation for FortiManager deletes the specified Alert Event.
// Returns error for service API and SDK errors.
// See the system - alert-event chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAlertEvent(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/alert-event"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemAlertEvent API operation for FortiManager gets the Alert Event
// with the specified index value.
// Returns the requested Alert Event value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - alert-event chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAlertEvent(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/alert-event"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemAlertConsole API operation for FortiManager updates the specified Alert Console.
// Returns the index value of the Alert Console and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - alert-console chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAlertConsole(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/alert-console"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemAlertConsole API operation for FortiManager deletes the specified Alert Console.
// Returns error for service API and SDK errors.
// See the system - alert-console chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAlertConsole(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - alert-console
	return
}

// ReadSystemAlertConsole API operation for FortiManager gets the Alert Console
// with the specified index value.
// Returns the requested Alert Console value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - alert-console chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAlertConsole(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/alert-console"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLocallogDiskSetting API operation for FortiManager updates the specified LocallogDiskSetting.
// Returns the index value of the LocallogDiskSetting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog disk setting chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLocallogDiskSetting(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/disk/setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemLocallogDiskSetting API operation for FortiManager deletes the specified LocallogDiskSetting.
// Returns error for service API and SDK errors.
// See the system - locallog disk setting chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLocallogDiskSetting(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - locallog disk setting
	return
}

// ReadSystemLocallogDiskSetting API operation for FortiManager gets the LocallogDiskSetting
// with the specified index value.
// Returns the requested LocallogDiskSetting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog disk setting chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLocallogDiskSetting(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/disk/setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLocallogDiskFilter API operation for FortiManager updates the specified LocallogDiskFilter.
// Returns the index value of the LocallogDiskFilter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog disk filter chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLocallogDiskFilter(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/disk/filter"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemLocallogDiskFilter API operation for FortiManager deletes the specified LocallogDiskFilter.
// Returns error for service API and SDK errors.
// See the system - locallog disk filter chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLocallogDiskFilter(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - locallog disk filter
	return
}

// ReadSystemLocallogDiskFilter API operation for FortiManager gets the LocallogDiskFilter
// with the specified index value.
// Returns the requested LocallogDiskFilter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog disk filter chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLocallogDiskFilter(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/disk/filter"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLocallogMemorySetting API operation for FortiManager updates the specified LocallogMemorySetting.
// Returns the index value of the LocallogMemorySetting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog memory setting chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLocallogMemorySetting(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/memory/setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemLocallogMemorySetting API operation for FortiManager deletes the specified LocallogMemorySetting.
// Returns error for service API and SDK errors.
// See the system - locallog memory setting chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLocallogMemorySetting(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - locallog memory setting
	return
}

// ReadSystemLocallogMemorySetting API operation for FortiManager gets the LocallogMemorySetting
// with the specified index value.
// Returns the requested LocallogMemorySetting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog memory setting chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLocallogMemorySetting(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/memory/setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLocallogMemoryFilter API operation for FortiManager updates the specified LocallogMemoryFilter.
// Returns the index value of the LocallogMemoryFilter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog memory filter chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLocallogMemoryFilter(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/memory/filter"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemLocallogMemoryFilter API operation for FortiManager deletes the specified LocallogMemoryFilter.
// Returns error for service API and SDK errors.
// See the system - locallog memory filter chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLocallogMemoryFilter(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - locallog memory filter
	return
}

// ReadSystemLocallogMemoryFilter API operation for FortiManager gets the LocallogMemoryFilter
// with the specified index value.
// Returns the requested LocallogMemoryFilter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog memory filter chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLocallogMemoryFilter(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/memory/filter"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLocallogFortianalyzerFilter API operation for FortiManager updates the specified LocallogFortianalyzerFilter.
// Returns the index value of the LocallogFortianalyzerFilter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog fortianalyzer filter chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLocallogFortianalyzerFilter(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/fortianalyzer/filter"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemLocallogFortianalyzerFilter API operation for FortiManager deletes the specified LocallogFortianalyzerFilter.
// Returns error for service API and SDK errors.
// See the system - locallog fortianalyzer filter chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLocallogFortianalyzerFilter(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - locallog fortianalyzer filter
	return
}

// ReadSystemLocallogFortianalyzerFilter API operation for FortiManager gets the LocallogFortianalyzerFilter
// with the specified index value.
// Returns the requested LocallogFortianalyzerFilter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog fortianalyzer filter chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLocallogFortianalyzerFilter(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/fortianalyzer/filter"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLocallogFortianalyzerSetting API operation for FortiManager updates the specified LocallogFortianalyzerSetting.
// Returns the index value of the LocallogFortianalyzerSetting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog fortianalyzer setting chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLocallogFortianalyzerSetting(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/fortianalyzer/setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemLocallogFortianalyzerSetting API operation for FortiManager deletes the specified LocallogFortianalyzerSetting.
// Returns error for service API and SDK errors.
// See the system - locallog fortianalyzer setting chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLocallogFortianalyzerSetting(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - locallog fortianalyzer setting
	return
}

// ReadSystemLocallogFortianalyzerSetting API operation for FortiManager gets the LocallogFortianalyzerSetting
// with the specified index value.
// Returns the requested LocallogFortianalyzerSetting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog fortianalyzer setting chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLocallogFortianalyzerSetting(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/fortianalyzer/setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLocallogFortianalyzer2Filter API operation for FortiManager updates the specified LocallogFortianalyzer2Filter.
// Returns the index value of the LocallogFortianalyzer2Filter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog fortianalyzer2 filter chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLocallogFortianalyzer2Filter(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/fortianalyzer2/filter"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemLocallogFortianalyzer2Filter API operation for FortiManager deletes the specified LocallogFortianalyzer2Filter.
// Returns error for service API and SDK errors.
// See the system - locallog fortianalyzer2 filter chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLocallogFortianalyzer2Filter(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - locallog fortianalyzer2 filter
	return
}

// ReadSystemLocallogFortianalyzer2Filter API operation for FortiManager gets the LocallogFortianalyzer2Filter
// with the specified index value.
// Returns the requested LocallogFortianalyzer2Filter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog fortianalyzer2 filter chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLocallogFortianalyzer2Filter(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/fortianalyzer2/filter"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLocallogFortianalyzer2Setting API operation for FortiManager updates the specified LocallogFortianalyzer2Setting.
// Returns the index value of the LocallogFortianalyzer2Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog fortianalyzer2 setting chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLocallogFortianalyzer2Setting(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/fortianalyzer2/setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemLocallogFortianalyzer2Setting API operation for FortiManager deletes the specified LocallogFortianalyzer2Setting.
// Returns error for service API and SDK errors.
// See the system - locallog fortianalyzer2 setting chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLocallogFortianalyzer2Setting(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - locallog fortianalyzer2 setting
	return
}

// ReadSystemLocallogFortianalyzer2Setting API operation for FortiManager gets the LocallogFortianalyzer2Setting
// with the specified index value.
// Returns the requested LocallogFortianalyzer2Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog fortianalyzer2 setting chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLocallogFortianalyzer2Setting(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/fortianalyzer2/setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLocallogFortianalyzer3Filter API operation for FortiManager updates the specified LocallogFortianalyzer3Filter.
// Returns the index value of the LocallogFortianalyzer3Filter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog fortianalyzer3 filter chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLocallogFortianalyzer3Filter(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/fortianalyzer3/filter"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemLocallogFortianalyzer3Filter API operation for FortiManager deletes the specified LocallogFortianalyzer3Filter.
// Returns error for service API and SDK errors.
// See the system - locallog fortianalyzer3 filter chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLocallogFortianalyzer3Filter(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - locallog fortianalyzer3 filter
	return
}

// ReadSystemLocallogFortianalyzer3Filter API operation for FortiManager gets the LocallogFortianalyzer3Filter
// with the specified index value.
// Returns the requested LocallogFortianalyzer3Filter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog fortianalyzer3 filter chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLocallogFortianalyzer3Filter(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/fortianalyzer3/filter"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLocallogFortianalyzer3Setting API operation for FortiManager updates the specified LocallogFortianalyzer3Setting.
// Returns the index value of the LocallogFortianalyzer3Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog fortianalyzer3 setting chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLocallogFortianalyzer3Setting(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/fortianalyzer3/setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemLocallogFortianalyzer3Setting API operation for FortiManager deletes the specified LocallogFortianalyzer3Setting.
// Returns error for service API and SDK errors.
// See the system - locallog fortianalyzer3 setting chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLocallogFortianalyzer3Setting(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - locallog fortianalyzer3 setting
	return
}

// ReadSystemLocallogFortianalyzer3Setting API operation for FortiManager gets the LocallogFortianalyzer3Setting
// with the specified index value.
// Returns the requested LocallogFortianalyzer3Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog fortianalyzer3 setting chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLocallogFortianalyzer3Setting(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/fortianalyzer3/setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLocallogSyslogdSetting API operation for FortiManager updates the specified LocallogSyslogdSetting.
// Returns the index value of the LocallogSyslogdSetting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog syslogd setting chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLocallogSyslogdSetting(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/syslogd/setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemLocallogSyslogdSetting API operation for FortiManager deletes the specified LocallogSyslogdSetting.
// Returns error for service API and SDK errors.
// See the system - locallog syslogd setting chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLocallogSyslogdSetting(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - locallog syslogd setting
	return
}

// ReadSystemLocallogSyslogdSetting API operation for FortiManager gets the LocallogSyslogdSetting
// with the specified index value.
// Returns the requested LocallogSyslogdSetting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog syslogd setting chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLocallogSyslogdSetting(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/syslogd/setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLocallogSyslogdFilter API operation for FortiManager updates the specified LocallogSyslogdFilter.
// Returns the index value of the LocallogSyslogdFilter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog syslogd filter chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLocallogSyslogdFilter(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/syslogd/filter"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemLocallogSyslogdFilter API operation for FortiManager deletes the specified LocallogSyslogdFilter.
// Returns error for service API and SDK errors.
// See the system - locallog syslogd filter chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLocallogSyslogdFilter(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - locallog syslogd filter
	return
}

// ReadSystemLocallogSyslogdFilter API operation for FortiManager gets the LocallogSyslogdFilter
// with the specified index value.
// Returns the requested LocallogSyslogdFilter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog syslogd filter chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLocallogSyslogdFilter(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/syslogd/filter"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLocallogSyslogd2Setting API operation for FortiManager updates the specified LocallogSyslogd2Setting.
// Returns the index value of the LocallogSyslogd2Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog syslogd2 setting chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLocallogSyslogd2Setting(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/syslogd2/setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemLocallogSyslogd2Setting API operation for FortiManager deletes the specified LocallogSyslogd2Setting.
// Returns error for service API and SDK errors.
// See the system - locallog syslogd2 setting chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLocallogSyslogd2Setting(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - locallog syslogd2 setting
	return
}

// ReadSystemLocallogSyslogd2Setting API operation for FortiManager gets the LocallogSyslogd2Setting
// with the specified index value.
// Returns the requested LocallogSyslogd2Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog syslogd2 setting chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLocallogSyslogd2Setting(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/syslogd2/setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLocallogSyslogd2Filter API operation for FortiManager updates the specified LocallogSyslogd2Filter.
// Returns the index value of the LocallogSyslogd2Filter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog syslogd2 filter chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLocallogSyslogd2Filter(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/syslogd2/filter"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemLocallogSyslogd2Filter API operation for FortiManager deletes the specified LocallogSyslogd2Filter.
// Returns error for service API and SDK errors.
// See the system - locallog syslogd2 filter chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLocallogSyslogd2Filter(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - locallog syslogd2 filter
	return
}

// ReadSystemLocallogSyslogd2Filter API operation for FortiManager gets the LocallogSyslogd2Filter
// with the specified index value.
// Returns the requested LocallogSyslogd2Filter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog syslogd2 filter chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLocallogSyslogd2Filter(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/syslogd2/filter"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLocallogSyslogd3Setting API operation for FortiManager updates the specified LocallogSyslogd3Setting.
// Returns the index value of the LocallogSyslogd3Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog syslogd3 setting chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLocallogSyslogd3Setting(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/syslogd3/setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemLocallogSyslogd3Setting API operation for FortiManager deletes the specified LocallogSyslogd3Setting.
// Returns error for service API and SDK errors.
// See the system - locallog syslogd3 setting chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLocallogSyslogd3Setting(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - locallog syslogd3 setting
	return
}

// ReadSystemLocallogSyslogd3Setting API operation for FortiManager gets the LocallogSyslogd3Setting
// with the specified index value.
// Returns the requested LocallogSyslogd3Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog syslogd3 setting chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLocallogSyslogd3Setting(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/syslogd3/setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLocallogSyslogd3Filter API operation for FortiManager updates the specified LocallogSyslogd3Filter.
// Returns the index value of the LocallogSyslogd3Filter and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog syslogd3 filter chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLocallogSyslogd3Filter(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/syslogd3/filter"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemLocallogSyslogd3Filter API operation for FortiManager deletes the specified LocallogSyslogd3Filter.
// Returns error for service API and SDK errors.
// See the system - locallog syslogd3 filter chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLocallogSyslogd3Filter(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - locallog syslogd3 filter
	return
}

// ReadSystemLocallogSyslogd3Filter API operation for FortiManager gets the LocallogSyslogd3Filter
// with the specified index value.
// Returns the requested LocallogSyslogd3Filter value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog syslogd3 filter chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLocallogSyslogd3Filter(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/syslogd3/filter"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLocallogSetting API operation for FortiManager updates the specified LocallogSetting.
// Returns the index value of the LocallogSetting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog setting chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLocallogSetting(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemLocallogSetting API operation for FortiManager deletes the specified LocallogSetting.
// Returns error for service API and SDK errors.
// See the system - locallog setting chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLocallogSetting(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - locallog setting
	return
}

// ReadSystemLocallogSetting API operation for FortiManager gets the LocallogSetting
// with the specified index value.
// Returns the requested LocallogSetting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - locallog setting chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLocallogSetting(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/locallog/setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemFips API operation for FortiManager updates the specified Fips.
// Returns the index value of the Fips and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - fips chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemFips(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/fips"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemFips API operation for FortiManager deletes the specified Fips.
// Returns error for service API and SDK errors.
// See the system - fips chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemFips(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - fips
	return
}

// ReadSystemFips API operation for FortiManager gets the Fips
// with the specified index value.
// Returns the requested Fips value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - fips chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemFips(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/fips"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemWorkflowApprovalMatrix API operation for FortiManager creates a new WorkflowApproval Matrix.
// Returns the index value of the WorkflowApproval Matrix and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - workflow approval-matrix chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemWorkflowApprovalMatrix(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/workflow/approval-matrix"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateSystemWorkflowApprovalMatrix API operation for FortiManager updates the specified WorkflowApproval Matrix.
// Returns the index value of the WorkflowApproval Matrix and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - workflow approval-matrix chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemWorkflowApprovalMatrix(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/workflow/approval-matrix"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemWorkflowApprovalMatrix API operation for FortiManager deletes the specified WorkflowApproval Matrix.
// Returns error for service API and SDK errors.
// See the system - workflow approval-matrix chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemWorkflowApprovalMatrix(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/workflow/approval-matrix"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemWorkflowApprovalMatrix API operation for FortiManager gets the WorkflowApproval Matrix
// with the specified index value.
// Returns the requested WorkflowApproval Matrix value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - workflow approval-matrix chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemWorkflowApprovalMatrix(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/workflow/approval-matrix"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemDocker API operation for FortiManager updates the specified Docker.
// Returns the index value of the Docker and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - docker chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemDocker(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/docker"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemDocker API operation for FortiManager deletes the specified Docker.
// Returns error for service API and SDK errors.
// See the system - docker chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemDocker(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - docker
	return
}

// ReadSystemDocker API operation for FortiManager gets the Docker
// with the specified index value.
// Returns the requested Docker value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - docker chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemDocker(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/docker"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemSniffer API operation for FortiManager creates a new Sniffer.
// Returns the index value of the Sniffer and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sniffer chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemSniffer(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/sniffer"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateSystemSniffer API operation for FortiManager updates the specified Sniffer.
// Returns the index value of the Sniffer and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sniffer chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSniffer(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/sniffer"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemSniffer API operation for FortiManager deletes the specified Sniffer.
// Returns error for service API and SDK errors.
// See the system - sniffer chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSniffer(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/sniffer"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemSniffer API operation for FortiManager gets the Sniffer
// with the specified index value.
// Returns the requested Sniffer value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sniffer chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSniffer(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/sniffer"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemDm API operation for FortiManager updates the specified Dm.
// Returns the index value of the Dm and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - dm chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemDm(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/dm"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemDm API operation for FortiManager deletes the specified Dm.
// Returns error for service API and SDK errors.
// See the system - dm chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemDm(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - dm
	return
}

// ReadSystemDm API operation for FortiManager gets the Dm
// with the specified index value.
// Returns the requested Dm value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - dm chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemDm(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/dm"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLogAlert API operation for FortiManager updates the specified LogAlert.
// Returns the index value of the LogAlert and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log alert chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLogAlert(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/alert"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemLogAlert API operation for FortiManager deletes the specified LogAlert.
// Returns error for service API and SDK errors.
// See the system - log alert chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLogAlert(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - log alert
	return
}

// ReadSystemLogAlert API operation for FortiManager gets the LogAlert
// with the specified index value.
// Returns the requested LogAlert value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log alert chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLogAlert(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/alert"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLogInterfaceStats API operation for FortiManager updates the specified LogInterface Stats.
// Returns the index value of the LogInterface Stats and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log interface-stats chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLogInterfaceStats(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/interface-stats"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemLogInterfaceStats API operation for FortiManager deletes the specified LogInterface Stats.
// Returns error for service API and SDK errors.
// See the system - log interface-stats chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLogInterfaceStats(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - log interface-stats
	return
}

// ReadSystemLogInterfaceStats API operation for FortiManager gets the LogInterface Stats
// with the specified index value.
// Returns the requested LogInterface Stats value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log interface-stats chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLogInterfaceStats(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/interface-stats"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLogIoc API operation for FortiManager updates the specified LogIoc.
// Returns the index value of the LogIoc and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log ioc chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLogIoc(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/ioc"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemLogIoc API operation for FortiManager deletes the specified LogIoc.
// Returns error for service API and SDK errors.
// See the system - log ioc chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLogIoc(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - log ioc
	return
}

// ReadSystemLogIoc API operation for FortiManager gets the LogIoc
// with the specified index value.
// Returns the requested LogIoc value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log ioc chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLogIoc(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/ioc"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemLogMailDomain API operation for FortiManager creates a new LogMail Domain.
// Returns the index value of the LogMail Domain and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log mail-domain chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemLogMailDomain(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/mail-domain"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateSystemLogMailDomain API operation for FortiManager updates the specified LogMail Domain.
// Returns the index value of the LogMail Domain and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log mail-domain chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLogMailDomain(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/mail-domain"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemLogMailDomain API operation for FortiManager deletes the specified LogMail Domain.
// Returns error for service API and SDK errors.
// See the system - log mail-domain chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLogMailDomain(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/log/mail-domain"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemLogMailDomain API operation for FortiManager gets the LogMail Domain
// with the specified index value.
// Returns the requested LogMail Domain value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log mail-domain chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLogMailDomain(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/mail-domain"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemLogDeviceDisable API operation for FortiManager creates a new LogDevice Disable.
// Returns the index value of the LogDevice Disable and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log device-disable chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemLogDeviceDisable(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/device-disable"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateSystemLogDeviceDisable API operation for FortiManager updates the specified LogDevice Disable.
// Returns the index value of the LogDevice Disable and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log device-disable chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLogDeviceDisable(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/device-disable"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemLogDeviceDisable API operation for FortiManager deletes the specified LogDevice Disable.
// Returns error for service API and SDK errors.
// See the system - log device-disable chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLogDeviceDisable(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/log/device-disable"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemLogDeviceDisable API operation for FortiManager gets the LogDevice Disable
// with the specified index value.
// Returns the requested LogDevice Disable value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log device-disable chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLogDeviceDisable(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/device-disable"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLogSettings API operation for FortiManager updates the specified LogSettings.
// Returns the index value of the LogSettings and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log settings chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLogSettings(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/settings"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemLogSettings API operation for FortiManager deletes the specified LogSettings.
// Returns error for service API and SDK errors.
// See the system - log settings chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLogSettings(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - log settings
	return
}

// ReadSystemLogSettings API operation for FortiManager gets the LogSettings
// with the specified index value.
// Returns the requested LogSettings value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log settings chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLogSettings(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/settings"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLogSettingsRollingAnalyzer API operation for FortiManager updates the specified LogSettingsRolling Analyzer.
// Returns the index value of the LogSettingsRolling Analyzer and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log settings rolling-analyzer chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLogSettingsRollingAnalyzer(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/settings/rolling-analyzer"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemLogSettingsRollingAnalyzer API operation for FortiManager deletes the specified LogSettingsRolling Analyzer.
// Returns error for service API and SDK errors.
// See the system - log settings rolling-analyzer chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLogSettingsRollingAnalyzer(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - log settings rolling-analyzer
	return
}

// ReadSystemLogSettingsRollingAnalyzer API operation for FortiManager gets the LogSettingsRolling Analyzer
// with the specified index value.
// Returns the requested LogSettingsRolling Analyzer value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log settings rolling-analyzer chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLogSettingsRollingAnalyzer(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/settings/rolling-analyzer"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLogSettingsRollingLocal API operation for FortiManager updates the specified LogSettingsRolling Local.
// Returns the index value of the LogSettingsRolling Local and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log settings rolling-local chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLogSettingsRollingLocal(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/settings/rolling-local"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemLogSettingsRollingLocal API operation for FortiManager deletes the specified LogSettingsRolling Local.
// Returns error for service API and SDK errors.
// See the system - log settings rolling-local chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLogSettingsRollingLocal(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - log settings rolling-local
	return
}

// ReadSystemLogSettingsRollingLocal API operation for FortiManager gets the LogSettingsRolling Local
// with the specified index value.
// Returns the requested LogSettingsRolling Local value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log settings rolling-local chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLogSettingsRollingLocal(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/settings/rolling-local"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLogSettingsRollingRegular API operation for FortiManager updates the specified LogSettingsRolling Regular.
// Returns the index value of the LogSettingsRolling Regular and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log settings rolling-regular chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLogSettingsRollingRegular(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/settings/rolling-regular"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemLogSettingsRollingRegular API operation for FortiManager deletes the specified LogSettingsRolling Regular.
// Returns error for service API and SDK errors.
// See the system - log settings rolling-regular chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLogSettingsRollingRegular(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - log settings rolling-regular
	return
}

// ReadSystemLogSettingsRollingRegular API operation for FortiManager gets the LogSettingsRolling Regular
// with the specified index value.
// Returns the requested LogSettingsRolling Regular value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log settings rolling-regular chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLogSettingsRollingRegular(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/log/settings/rolling-regular"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemLogFetchServerSettings API operation for FortiManager updates the specified Log FetchServer Settings.
// Returns the index value of the Log FetchServer Settings and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log-fetch server-settings chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLogFetchServerSettings(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/log-fetch/server-settings"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemLogFetchServerSettings API operation for FortiManager deletes the specified Log FetchServer Settings.
// Returns error for service API and SDK errors.
// See the system - log-fetch server-settings chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLogFetchServerSettings(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - log-fetch server-settings
	return
}

// ReadSystemLogFetchServerSettings API operation for FortiManager gets the Log FetchServer Settings
// with the specified index value.
// Returns the requested Log FetchServer Settings value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log-fetch server-settings chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLogFetchServerSettings(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/log-fetch/server-settings"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemLogFetchClientProfile API operation for FortiManager creates a new Log FetchClient Profile.
// Returns the index value of the Log FetchClient Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log-fetch client-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemLogFetchClientProfile(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/log-fetch/client-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateSystemLogFetchClientProfile API operation for FortiManager updates the specified Log FetchClient Profile.
// Returns the index value of the Log FetchClient Profile and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log-fetch client-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemLogFetchClientProfile(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/log-fetch/client-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemLogFetchClientProfile API operation for FortiManager deletes the specified Log FetchClient Profile.
// Returns error for service API and SDK errors.
// See the system - log-fetch client-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemLogFetchClientProfile(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/log-fetch/client-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemLogFetchClientProfile API operation for FortiManager gets the Log FetchClient Profile
// with the specified index value.
// Returns the requested Log FetchClient Profile value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - log-fetch client-profile chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemLogFetchClientProfile(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/log-fetch/client-profile"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemSql API operation for FortiManager updates the specified Sql.
// Returns the index value of the Sql and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sql chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSql(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/sql"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemSql API operation for FortiManager deletes the specified Sql.
// Returns error for service API and SDK errors.
// See the system - sql chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSql(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - sql
	return
}

// ReadSystemSql API operation for FortiManager gets the Sql
// with the specified index value.
// Returns the requested Sql value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sql chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSql(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/sql"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemSqlTsIndexField API operation for FortiManager creates a new SqlTs Index Field.
// Returns the index value of the SqlTs Index Field and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sql ts-index-field chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemSqlTsIndexField(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/sql/ts-index-field"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateSystemSqlTsIndexField API operation for FortiManager updates the specified SqlTs Index Field.
// Returns the index value of the SqlTs Index Field and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sql ts-index-field chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSqlTsIndexField(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/sql/ts-index-field"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemSqlTsIndexField API operation for FortiManager deletes the specified SqlTs Index Field.
// Returns error for service API and SDK errors.
// See the system - sql ts-index-field chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSqlTsIndexField(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/sql/ts-index-field"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemSqlTsIndexField API operation for FortiManager gets the SqlTs Index Field
// with the specified index value.
// Returns the requested SqlTs Index Field value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sql ts-index-field chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSqlTsIndexField(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/sql/ts-index-field"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemSqlCustomSkipidx API operation for FortiManager creates a new SqlCustom Skipidx.
// Returns the index value of the SqlCustom Skipidx and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sql custom-skipidx chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemSqlCustomSkipidx(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/sql/custom-skipidx"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateSystemSqlCustomSkipidx API operation for FortiManager updates the specified SqlCustom Skipidx.
// Returns the index value of the SqlCustom Skipidx and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sql custom-skipidx chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSqlCustomSkipidx(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/sql/custom-skipidx"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemSqlCustomSkipidx API operation for FortiManager deletes the specified SqlCustom Skipidx.
// Returns error for service API and SDK errors.
// See the system - sql custom-skipidx chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSqlCustomSkipidx(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/sql/custom-skipidx"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemSqlCustomSkipidx API operation for FortiManager gets the SqlCustom Skipidx
// with the specified index value.
// Returns the requested SqlCustom Skipidx value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sql custom-skipidx chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSqlCustomSkipidx(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/sql/custom-skipidx"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemSqlCustomIndex API operation for FortiManager creates a new SqlCustom Index.
// Returns the index value of the SqlCustom Index and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sql custom-index chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemSqlCustomIndex(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/sql/custom-index"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateSystemSqlCustomIndex API operation for FortiManager updates the specified SqlCustom Index.
// Returns the index value of the SqlCustom Index and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sql custom-index chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSqlCustomIndex(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/sql/custom-index"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemSqlCustomIndex API operation for FortiManager deletes the specified SqlCustom Index.
// Returns error for service API and SDK errors.
// See the system - sql custom-index chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemSqlCustomIndex(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/sql/custom-index"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemSqlCustomIndex API operation for FortiManager gets the SqlCustom Index
// with the specified index value.
// Returns the requested SqlCustom Index value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - sql custom-index chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSqlCustomIndex(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/sql/custom-index"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemReportEstBrowseTime API operation for FortiManager updates the specified ReportEst Browse Time.
// Returns the index value of the ReportEst Browse Time and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - report est-browse-time chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemReportEstBrowseTime(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/report/est-browse-time"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemReportEstBrowseTime API operation for FortiManager deletes the specified ReportEst Browse Time.
// Returns error for service API and SDK errors.
// See the system - report est-browse-time chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemReportEstBrowseTime(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - report est-browse-time
	return
}

// ReadSystemReportEstBrowseTime API operation for FortiManager gets the ReportEst Browse Time
// with the specified index value.
// Returns the requested ReportEst Browse Time value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - report est-browse-time chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemReportEstBrowseTime(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/report/est-browse-time"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemReportAutoCache API operation for FortiManager updates the specified ReportAuto Cache.
// Returns the index value of the ReportAuto Cache and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - report auto-cache chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemReportAutoCache(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/report/auto-cache"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemReportAutoCache API operation for FortiManager deletes the specified ReportAuto Cache.
// Returns error for service API and SDK errors.
// See the system - report auto-cache chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemReportAutoCache(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - report auto-cache
	return
}

// ReadSystemReportAutoCache API operation for FortiManager gets the ReportAuto Cache
// with the specified index value.
// Returns the requested ReportAuto Cache value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - report auto-cache chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemReportAutoCache(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/report/auto-cache"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemReportSetting API operation for FortiManager updates the specified ReportSetting.
// Returns the index value of the ReportSetting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - report setting chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemReportSetting(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/report/setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemReportSetting API operation for FortiManager deletes the specified ReportSetting.
// Returns error for service API and SDK errors.
// See the system - report setting chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemReportSetting(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - report setting
	return
}

// ReadSystemReportSetting API operation for FortiManager gets the ReportSetting
// with the specified index value.
// Returns the requested ReportSetting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - report setting chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemReportSetting(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/report/setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateSystemReportGroup API operation for FortiManager creates a new ReportGroup.
// Returns the index value of the ReportGroup and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - report group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemReportGroup(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/report/group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateSystemReportGroup API operation for FortiManager updates the specified ReportGroup.
// Returns the index value of the ReportGroup and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - report group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemReportGroup(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/report/group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemReportGroup API operation for FortiManager deletes the specified ReportGroup.
// Returns error for service API and SDK errors.
// See the system - report group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemReportGroup(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/system/report/group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadSystemReportGroup API operation for FortiManager gets the ReportGroup
// with the specified index value.
// Returns the requested ReportGroup value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - report group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemReportGroup(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/report/group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemFortiviewSetting API operation for FortiManager updates the specified FortiviewSetting.
// Returns the index value of the FortiviewSetting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - fortiview setting chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemFortiviewSetting(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/fortiview/setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemFortiviewSetting API operation for FortiManager deletes the specified FortiviewSetting.
// Returns error for service API and SDK errors.
// See the system - fortiview setting chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemFortiviewSetting(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - fortiview setting
	return
}

// ReadSystemFortiviewSetting API operation for FortiManager gets the FortiviewSetting
// with the specified index value.
// Returns the requested FortiviewSetting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - fortiview setting chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemFortiviewSetting(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/fortiview/setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemFortiviewAutoCache API operation for FortiManager updates the specified FortiviewAuto Cache.
// Returns the index value of the FortiviewAuto Cache and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - fortiview auto-cache chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemFortiviewAutoCache(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/fortiview/auto-cache"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemFortiviewAutoCache API operation for FortiManager deletes the specified FortiviewAuto Cache.
// Returns error for service API and SDK errors.
// See the system - fortiview auto-cache chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemFortiviewAutoCache(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - fortiview auto-cache
	return
}

// ReadSystemFortiviewAutoCache API operation for FortiManager gets the FortiviewAuto Cache
// with the specified index value.
// Returns the requested FortiviewAuto Cache value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - fortiview auto-cache chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemFortiviewAutoCache(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/fortiview/auto-cache"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemAutoDelete API operation for FortiManager updates the specified Auto Delete.
// Returns the index value of the Auto Delete and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - auto-delete chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAutoDelete(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/auto-delete"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemAutoDelete API operation for FortiManager deletes the specified Auto Delete.
// Returns error for service API and SDK errors.
// See the system - auto-delete chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAutoDelete(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - auto-delete
	return
}

// ReadSystemAutoDelete API operation for FortiManager gets the Auto Delete
// with the specified index value.
// Returns the requested Auto Delete value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - auto-delete chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAutoDelete(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/auto-delete"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemAutoDeleteQuarantineFilesAutoDeletion API operation for FortiManager updates the specified Auto DeleteQuarantine Files Auto Deletion.
// Returns the index value of the Auto DeleteQuarantine Files Auto Deletion and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - auto-delete quarantine-files-auto-deletion chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAutoDeleteQuarantineFilesAutoDeletion(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/auto-delete/quarantine-files-auto-deletion"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemAutoDeleteQuarantineFilesAutoDeletion API operation for FortiManager deletes the specified Auto DeleteQuarantine Files Auto Deletion.
// Returns error for service API and SDK errors.
// See the system - auto-delete quarantine-files-auto-deletion chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAutoDeleteQuarantineFilesAutoDeletion(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - auto-delete quarantine-files-auto-deletion
	return
}

// ReadSystemAutoDeleteQuarantineFilesAutoDeletion API operation for FortiManager gets the Auto DeleteQuarantine Files Auto Deletion
// with the specified index value.
// Returns the requested Auto DeleteQuarantine Files Auto Deletion value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - auto-delete quarantine-files-auto-deletion chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAutoDeleteQuarantineFilesAutoDeletion(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/auto-delete/quarantine-files-auto-deletion"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemAutoDeleteDlpFilesAutoDeletion API operation for FortiManager updates the specified Auto DeleteDlp Files Auto Deletion.
// Returns the index value of the Auto DeleteDlp Files Auto Deletion and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - auto-delete dlp-files-auto-deletion chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAutoDeleteDlpFilesAutoDeletion(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/auto-delete/dlp-files-auto-deletion"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemAutoDeleteDlpFilesAutoDeletion API operation for FortiManager deletes the specified Auto DeleteDlp Files Auto Deletion.
// Returns error for service API and SDK errors.
// See the system - auto-delete dlp-files-auto-deletion chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAutoDeleteDlpFilesAutoDeletion(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - auto-delete dlp-files-auto-deletion
	return
}

// ReadSystemAutoDeleteDlpFilesAutoDeletion API operation for FortiManager gets the Auto DeleteDlp Files Auto Deletion
// with the specified index value.
// Returns the requested Auto DeleteDlp Files Auto Deletion value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - auto-delete dlp-files-auto-deletion chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAutoDeleteDlpFilesAutoDeletion(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/auto-delete/dlp-files-auto-deletion"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemAutoDeleteReportAutoDeletion API operation for FortiManager updates the specified Auto DeleteReport Auto Deletion.
// Returns the index value of the Auto DeleteReport Auto Deletion and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - auto-delete report-auto-deletion chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAutoDeleteReportAutoDeletion(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/auto-delete/report-auto-deletion"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemAutoDeleteReportAutoDeletion API operation for FortiManager deletes the specified Auto DeleteReport Auto Deletion.
// Returns error for service API and SDK errors.
// See the system - auto-delete report-auto-deletion chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAutoDeleteReportAutoDeletion(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - auto-delete report-auto-deletion
	return
}

// ReadSystemAutoDeleteReportAutoDeletion API operation for FortiManager gets the Auto DeleteReport Auto Deletion
// with the specified index value.
// Returns the requested Auto DeleteReport Auto Deletion value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - auto-delete report-auto-deletion chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAutoDeleteReportAutoDeletion(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/auto-delete/report-auto-deletion"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateSystemAutoDeleteLogAutoDeletion API operation for FortiManager updates the specified Auto DeleteLog Auto Deletion.
// Returns the index value of the Auto DeleteLog Auto Deletion and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - auto-delete log-auto-deletion chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemAutoDeleteLogAutoDeletion(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/system/auto-delete/log-auto-deletion"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteSystemAutoDeleteLogAutoDeletion API operation for FortiManager deletes the specified Auto DeleteLog Auto Deletion.
// Returns error for service API and SDK errors.
// See the system - auto-delete log-auto-deletion chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemAutoDeleteLogAutoDeletion(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for system - auto-delete log-auto-deletion
	return
}

// ReadSystemAutoDeleteLogAutoDeletion API operation for FortiManager gets the Auto DeleteLog Auto Deletion
// with the specified index value.
// Returns the requested Auto DeleteLog Auto Deletion value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - auto-delete log-auto-deletion chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemAutoDeleteLogAutoDeletion(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/system/auto-delete/log-auto-deletion"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateFmupdateAvIpsWebProxy API operation for FortiManager updates the specified Av IpsWeb Proxy.
// Returns the index value of the Av IpsWeb Proxy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - av-ips web-proxy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFmupdateAvIpsWebProxy(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/av-ips/web-proxy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteFmupdateAvIpsWebProxy API operation for FortiManager deletes the specified Av IpsWeb Proxy.
// Returns error for service API and SDK errors.
// See the fmupdate - av-ips web-proxy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFmupdateAvIpsWebProxy(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for fmupdate - av-ips web-proxy
	return
}

// ReadFmupdateAvIpsWebProxy API operation for FortiManager gets the Av IpsWeb Proxy
// with the specified index value.
// Returns the requested Av IpsWeb Proxy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - av-ips web-proxy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFmupdateAvIpsWebProxy(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/av-ips/web-proxy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateFmupdateAvIpsAdvancedLog API operation for FortiManager updates the specified Av IpsAdvanced Log.
// Returns the index value of the Av IpsAdvanced Log and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - av-ips advanced-log chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFmupdateAvIpsAdvancedLog(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/av-ips/advanced-log"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteFmupdateAvIpsAdvancedLog API operation for FortiManager deletes the specified Av IpsAdvanced Log.
// Returns error for service API and SDK errors.
// See the fmupdate - av-ips advanced-log chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFmupdateAvIpsAdvancedLog(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for fmupdate - av-ips advanced-log
	return
}

// ReadFmupdateAvIpsAdvancedLog API operation for FortiManager gets the Av IpsAdvanced Log
// with the specified index value.
// Returns the requested Av IpsAdvanced Log value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - av-ips advanced-log chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFmupdateAvIpsAdvancedLog(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/av-ips/advanced-log"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateFmupdateWebSpamWebProxy API operation for FortiManager updates the specified Web SpamWeb Proxy.
// Returns the index value of the Web SpamWeb Proxy and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - web-spam web-proxy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFmupdateWebSpamWebProxy(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/web-spam/web-proxy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteFmupdateWebSpamWebProxy API operation for FortiManager deletes the specified Web SpamWeb Proxy.
// Returns error for service API and SDK errors.
// See the fmupdate - web-spam web-proxy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFmupdateWebSpamWebProxy(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for fmupdate - web-spam web-proxy
	return
}

// ReadFmupdateWebSpamWebProxy API operation for FortiManager gets the Web SpamWeb Proxy
// with the specified index value.
// Returns the requested Web SpamWeb Proxy value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - web-spam web-proxy chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFmupdateWebSpamWebProxy(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/web-spam/web-proxy"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateFmupdateWebSpamFgdSetting API operation for FortiManager updates the specified Web SpamFgd Setting.
// Returns the index value of the Web SpamFgd Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - web-spam fgd-setting chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFmupdateWebSpamFgdSetting(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/web-spam/fgd-setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteFmupdateWebSpamFgdSetting API operation for FortiManager deletes the specified Web SpamFgd Setting.
// Returns error for service API and SDK errors.
// See the fmupdate - web-spam fgd-setting chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFmupdateWebSpamFgdSetting(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for fmupdate - web-spam fgd-setting
	return
}

// ReadFmupdateWebSpamFgdSetting API operation for FortiManager gets the Web SpamFgd Setting
// with the specified index value.
// Returns the requested Web SpamFgd Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - web-spam fgd-setting chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFmupdateWebSpamFgdSetting(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/web-spam/fgd-setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateFmupdateWebSpamFgdSettingServerOverride API operation for FortiManager updates the specified Web SpamFgd SettingServer Override.
// Returns the index value of the Web SpamFgd SettingServer Override and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - web-spam fgd-setting server-override chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFmupdateWebSpamFgdSettingServerOverride(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/web-spam/fgd-setting/server-override"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteFmupdateWebSpamFgdSettingServerOverride API operation for FortiManager deletes the specified Web SpamFgd SettingServer Override.
// Returns error for service API and SDK errors.
// See the fmupdate - web-spam fgd-setting server-override chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFmupdateWebSpamFgdSettingServerOverride(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for fmupdate - web-spam fgd-setting server-override
	return
}

// ReadFmupdateWebSpamFgdSettingServerOverride API operation for FortiManager gets the Web SpamFgd SettingServer Override
// with the specified index value.
// Returns the requested Web SpamFgd SettingServer Override value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - web-spam fgd-setting server-override chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFmupdateWebSpamFgdSettingServerOverride(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/web-spam/fgd-setting/server-override"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateFmupdateWebSpamFgdSettingServerOverrideServlist API operation for FortiManager creates a new Web SpamFgd SettingServer OverrideServlist.
// Returns the index value of the Web SpamFgd SettingServer OverrideServlist and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - web-spam fgd-setting server-override servlist chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFmupdateWebSpamFgdSettingServerOverrideServlist(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/web-spam/fgd-setting/server-override/servlist"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateFmupdateWebSpamFgdSettingServerOverrideServlist API operation for FortiManager updates the specified Web SpamFgd SettingServer OverrideServlist.
// Returns the index value of the Web SpamFgd SettingServer OverrideServlist and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - web-spam fgd-setting server-override servlist chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFmupdateWebSpamFgdSettingServerOverrideServlist(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/web-spam/fgd-setting/server-override/servlist"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteFmupdateWebSpamFgdSettingServerOverrideServlist API operation for FortiManager deletes the specified Web SpamFgd SettingServer OverrideServlist.
// Returns error for service API and SDK errors.
// See the fmupdate - web-spam fgd-setting server-override servlist chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFmupdateWebSpamFgdSettingServerOverrideServlist(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/fmupdate/web-spam/fgd-setting/server-override/servlist"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadFmupdateWebSpamFgdSettingServerOverrideServlist API operation for FortiManager gets the Web SpamFgd SettingServer OverrideServlist
// with the specified index value.
// Returns the requested Web SpamFgd SettingServer OverrideServlist value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - web-spam fgd-setting server-override servlist chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFmupdateWebSpamFgdSettingServerOverrideServlist(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/web-spam/fgd-setting/server-override/servlist"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateFmupdateFctServices API operation for FortiManager updates the specified Fct Services.
// Returns the index value of the Fct Services and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - fct-services chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFmupdateFctServices(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/fct-services"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteFmupdateFctServices API operation for FortiManager deletes the specified Fct Services.
// Returns error for service API and SDK errors.
// See the fmupdate - fct-services chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFmupdateFctServices(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for fmupdate - fct-services
	return
}

// ReadFmupdateFctServices API operation for FortiManager gets the Fct Services
// with the specified index value.
// Returns the requested Fct Services value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - fct-services chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFmupdateFctServices(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/fct-services"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateFmupdateAnalyzerVirusreport API operation for FortiManager updates the specified AnalyzerVirusreport.
// Returns the index value of the AnalyzerVirusreport and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - analyzer virusreport chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFmupdateAnalyzerVirusreport(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/analyzer/virusreport"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteFmupdateAnalyzerVirusreport API operation for FortiManager deletes the specified AnalyzerVirusreport.
// Returns error for service API and SDK errors.
// See the fmupdate - analyzer virusreport chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFmupdateAnalyzerVirusreport(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for fmupdate - analyzer virusreport
	return
}

// ReadFmupdateAnalyzerVirusreport API operation for FortiManager gets the AnalyzerVirusreport
// with the specified index value.
// Returns the requested AnalyzerVirusreport value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - analyzer virusreport chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFmupdateAnalyzerVirusreport(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/analyzer/virusreport"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateFmupdateService API operation for FortiManager updates the specified Service.
// Returns the index value of the Service and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - service chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFmupdateService(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/service"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteFmupdateService API operation for FortiManager deletes the specified Service.
// Returns error for service API and SDK errors.
// See the fmupdate - service chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFmupdateService(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for fmupdate - service
	return
}

// ReadFmupdateService API operation for FortiManager gets the Service
// with the specified index value.
// Returns the requested Service value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - service chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFmupdateService(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/service"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateFmupdatePublicnetwork API operation for FortiManager updates the specified Publicnetwork.
// Returns the index value of the Publicnetwork and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - publicnetwork chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFmupdatePublicnetwork(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/publicnetwork"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteFmupdatePublicnetwork API operation for FortiManager deletes the specified Publicnetwork.
// Returns error for service API and SDK errors.
// See the fmupdate - publicnetwork chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFmupdatePublicnetwork(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for fmupdate - publicnetwork
	return
}

// ReadFmupdatePublicnetwork API operation for FortiManager gets the Publicnetwork
// with the specified index value.
// Returns the requested Publicnetwork value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - publicnetwork chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFmupdatePublicnetwork(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/publicnetwork"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateFmupdateDiskQuota API operation for FortiManager updates the specified Disk Quota.
// Returns the index value of the Disk Quota and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - disk-quota chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFmupdateDiskQuota(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/disk-quota"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteFmupdateDiskQuota API operation for FortiManager deletes the specified Disk Quota.
// Returns error for service API and SDK errors.
// See the fmupdate - disk-quota chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFmupdateDiskQuota(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for fmupdate - disk-quota
	return
}

// ReadFmupdateDiskQuota API operation for FortiManager gets the Disk Quota
// with the specified index value.
// Returns the requested Disk Quota value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - disk-quota chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFmupdateDiskQuota(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/disk-quota"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateFmupdateServerAccessPriorities API operation for FortiManager updates the specified Server Access Priorities.
// Returns the index value of the Server Access Priorities and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - server-access-priorities chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFmupdateServerAccessPriorities(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/server-access-priorities"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteFmupdateServerAccessPriorities API operation for FortiManager deletes the specified Server Access Priorities.
// Returns error for service API and SDK errors.
// See the fmupdate - server-access-priorities chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFmupdateServerAccessPriorities(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for fmupdate - server-access-priorities
	return
}

// ReadFmupdateServerAccessPriorities API operation for FortiManager gets the Server Access Priorities
// with the specified index value.
// Returns the requested Server Access Priorities value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - server-access-priorities chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFmupdateServerAccessPriorities(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/server-access-priorities"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateFmupdateServerAccessPrioritiesPrivateServer API operation for FortiManager creates a new Server Access PrioritiesPrivate Server.
// Returns the index value of the Server Access PrioritiesPrivate Server and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - server-access-priorities private-server chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFmupdateServerAccessPrioritiesPrivateServer(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/server-access-priorities/private-server"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateFmupdateServerAccessPrioritiesPrivateServer API operation for FortiManager updates the specified Server Access PrioritiesPrivate Server.
// Returns the index value of the Server Access PrioritiesPrivate Server and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - server-access-priorities private-server chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFmupdateServerAccessPrioritiesPrivateServer(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/server-access-priorities/private-server"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteFmupdateServerAccessPrioritiesPrivateServer API operation for FortiManager deletes the specified Server Access PrioritiesPrivate Server.
// Returns error for service API and SDK errors.
// See the fmupdate - server-access-priorities private-server chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFmupdateServerAccessPrioritiesPrivateServer(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/fmupdate/server-access-priorities/private-server"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadFmupdateServerAccessPrioritiesPrivateServer API operation for FortiManager gets the Server Access PrioritiesPrivate Server
// with the specified index value.
// Returns the requested Server Access PrioritiesPrivate Server value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - server-access-priorities private-server chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFmupdateServerAccessPrioritiesPrivateServer(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/server-access-priorities/private-server"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateFmupdateCustomUrlList API operation for FortiManager updates the specified Custom Url List.
// Returns the index value of the Custom Url List and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - custom-url-list chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFmupdateCustomUrlList(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/custom-url-list"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteFmupdateCustomUrlList API operation for FortiManager deletes the specified Custom Url List.
// Returns error for service API and SDK errors.
// See the fmupdate - custom-url-list chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFmupdateCustomUrlList(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for fmupdate - custom-url-list
	return
}

// ReadFmupdateCustomUrlList API operation for FortiManager gets the Custom Url List
// with the specified index value.
// Returns the requested Custom Url List value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - custom-url-list chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFmupdateCustomUrlList(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/custom-url-list"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateFmupdateServerOverrideStatus API operation for FortiManager updates the specified Server Override Status.
// Returns the index value of the Server Override Status and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - server-override-status chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFmupdateServerOverrideStatus(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/server-override-status"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteFmupdateServerOverrideStatus API operation for FortiManager deletes the specified Server Override Status.
// Returns error for service API and SDK errors.
// See the fmupdate - server-override-status chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFmupdateServerOverrideStatus(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for fmupdate - server-override-status
	return
}

// ReadFmupdateServerOverrideStatus API operation for FortiManager gets the Server Override Status
// with the specified index value.
// Returns the requested Server Override Status value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - server-override-status chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFmupdateServerOverrideStatus(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/server-override-status"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateFmupdateMultilayer API operation for FortiManager updates the specified Multilayer.
// Returns the index value of the Multilayer and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - multilayer chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFmupdateMultilayer(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/multilayer"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteFmupdateMultilayer API operation for FortiManager deletes the specified Multilayer.
// Returns error for service API and SDK errors.
// See the fmupdate - multilayer chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFmupdateMultilayer(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for fmupdate - multilayer
	return
}

// ReadFmupdateMultilayer API operation for FortiManager gets the Multilayer
// with the specified index value.
// Returns the requested Multilayer value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - multilayer chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFmupdateMultilayer(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/multilayer"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateFmupdateFdsSetting API operation for FortiManager updates the specified Fds Setting.
// Returns the index value of the Fds Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - fds-setting chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFmupdateFdsSetting(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/fds-setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteFmupdateFdsSetting API operation for FortiManager deletes the specified Fds Setting.
// Returns error for service API and SDK errors.
// See the fmupdate - fds-setting chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFmupdateFdsSetting(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for fmupdate - fds-setting
	return
}

// ReadFmupdateFdsSetting API operation for FortiManager gets the Fds Setting
// with the specified index value.
// Returns the requested Fds Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - fds-setting chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFmupdateFdsSetting(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/fds-setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateFmupdateFdsSettingServerOverride API operation for FortiManager updates the specified Fds SettingServer Override.
// Returns the index value of the Fds SettingServer Override and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - fds-setting server-override chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFmupdateFdsSettingServerOverride(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/fds-setting/server-override"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteFmupdateFdsSettingServerOverride API operation for FortiManager deletes the specified Fds SettingServer Override.
// Returns error for service API and SDK errors.
// See the fmupdate - fds-setting server-override chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFmupdateFdsSettingServerOverride(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for fmupdate - fds-setting server-override
	return
}

// ReadFmupdateFdsSettingServerOverride API operation for FortiManager gets the Fds SettingServer Override
// with the specified index value.
// Returns the requested Fds SettingServer Override value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - fds-setting server-override chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFmupdateFdsSettingServerOverride(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/fds-setting/server-override"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateFmupdateFdsSettingServerOverrideServlist API operation for FortiManager creates a new Fds SettingServer OverrideServlist.
// Returns the index value of the Fds SettingServer OverrideServlist and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - fds-setting server-override servlist chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFmupdateFdsSettingServerOverrideServlist(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/fds-setting/server-override/servlist"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateFmupdateFdsSettingServerOverrideServlist API operation for FortiManager updates the specified Fds SettingServer OverrideServlist.
// Returns the index value of the Fds SettingServer OverrideServlist and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - fds-setting server-override servlist chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFmupdateFdsSettingServerOverrideServlist(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/fds-setting/server-override/servlist"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteFmupdateFdsSettingServerOverrideServlist API operation for FortiManager deletes the specified Fds SettingServer OverrideServlist.
// Returns error for service API and SDK errors.
// See the fmupdate - fds-setting server-override servlist chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFmupdateFdsSettingServerOverrideServlist(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/fmupdate/fds-setting/server-override/servlist"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadFmupdateFdsSettingServerOverrideServlist API operation for FortiManager gets the Fds SettingServer OverrideServlist
// with the specified index value.
// Returns the requested Fds SettingServer OverrideServlist value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - fds-setting server-override servlist chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFmupdateFdsSettingServerOverrideServlist(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/fds-setting/server-override/servlist"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateFmupdateFdsSettingUpdateSchedule API operation for FortiManager updates the specified Fds SettingUpdate Schedule.
// Returns the index value of the Fds SettingUpdate Schedule and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - fds-setting update-schedule chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFmupdateFdsSettingUpdateSchedule(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/fds-setting/update-schedule"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteFmupdateFdsSettingUpdateSchedule API operation for FortiManager deletes the specified Fds SettingUpdate Schedule.
// Returns error for service API and SDK errors.
// See the fmupdate - fds-setting update-schedule chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFmupdateFdsSettingUpdateSchedule(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for fmupdate - fds-setting update-schedule
	return
}

// ReadFmupdateFdsSettingUpdateSchedule API operation for FortiManager gets the Fds SettingUpdate Schedule
// with the specified index value.
// Returns the requested Fds SettingUpdate Schedule value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - fds-setting update-schedule chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFmupdateFdsSettingUpdateSchedule(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/fds-setting/update-schedule"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateFmupdateFdsSettingPushOverrideToClient API operation for FortiManager updates the specified Fds SettingPush Override To Client.
// Returns the index value of the Fds SettingPush Override To Client and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - fds-setting push-override-to-client chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFmupdateFdsSettingPushOverrideToClient(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/fds-setting/push-override-to-client"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteFmupdateFdsSettingPushOverrideToClient API operation for FortiManager deletes the specified Fds SettingPush Override To Client.
// Returns error for service API and SDK errors.
// See the fmupdate - fds-setting push-override-to-client chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFmupdateFdsSettingPushOverrideToClient(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for fmupdate - fds-setting push-override-to-client
	return
}

// ReadFmupdateFdsSettingPushOverrideToClient API operation for FortiManager gets the Fds SettingPush Override To Client
// with the specified index value.
// Returns the requested Fds SettingPush Override To Client value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - fds-setting push-override-to-client chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFmupdateFdsSettingPushOverrideToClient(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/fds-setting/push-override-to-client"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateFmupdateFdsSettingPushOverrideToClientAnnounceIp API operation for FortiManager creates a new Fds SettingPush Override To ClientAnnounce Ip.
// Returns the index value of the Fds SettingPush Override To ClientAnnounce Ip and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - fds-setting push-override-to-client announce-ip chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateFmupdateFdsSettingPushOverrideToClientAnnounceIp(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/fds-setting/push-override-to-client/announce-ip"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateFmupdateFdsSettingPushOverrideToClientAnnounceIp API operation for FortiManager updates the specified Fds SettingPush Override To ClientAnnounce Ip.
// Returns the index value of the Fds SettingPush Override To ClientAnnounce Ip and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - fds-setting push-override-to-client announce-ip chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFmupdateFdsSettingPushOverrideToClientAnnounceIp(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/fds-setting/push-override-to-client/announce-ip"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteFmupdateFdsSettingPushOverrideToClientAnnounceIp API operation for FortiManager deletes the specified Fds SettingPush Override To ClientAnnounce Ip.
// Returns error for service API and SDK errors.
// See the fmupdate - fds-setting push-override-to-client announce-ip chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFmupdateFdsSettingPushOverrideToClientAnnounceIp(globaladom, mkey string, paralist []string) (err error) {
	path := "/cli/[*]/fmupdate/fds-setting/push-override-to-client/announce-ip"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadFmupdateFdsSettingPushOverrideToClientAnnounceIp API operation for FortiManager gets the Fds SettingPush Override To ClientAnnounce Ip
// with the specified index value.
// Returns the requested Fds SettingPush Override To ClientAnnounce Ip value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - fds-setting push-override-to-client announce-ip chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFmupdateFdsSettingPushOverrideToClientAnnounceIp(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/fds-setting/push-override-to-client/announce-ip"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateFmupdateFdsSettingPushOverride API operation for FortiManager updates the specified Fds SettingPush Override.
// Returns the index value of the Fds SettingPush Override and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - fds-setting push-override chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFmupdateFdsSettingPushOverride(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/fds-setting/push-override"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteFmupdateFdsSettingPushOverride API operation for FortiManager deletes the specified Fds SettingPush Override.
// Returns error for service API and SDK errors.
// See the fmupdate - fds-setting push-override chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFmupdateFdsSettingPushOverride(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for fmupdate - fds-setting push-override
	return
}

// ReadFmupdateFdsSettingPushOverride API operation for FortiManager gets the Fds SettingPush Override
// with the specified index value.
// Returns the requested Fds SettingPush Override value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - fds-setting push-override chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFmupdateFdsSettingPushOverride(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/fds-setting/push-override"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}


// UpdateFmupdateFwmSetting API operation for FortiManager updates the specified Fwm Setting.
// Returns the index value of the Fwm Setting and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - fwm-setting chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateFmupdateFwmSetting(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/fwm-setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteFmupdateFwmSetting API operation for FortiManager deletes the specified Fwm Setting.
// Returns error for service API and SDK errors.
// See the fmupdate - fwm-setting chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteFmupdateFwmSetting(globaladom, mkey string, paralist []string) (err error) {

	//No unset API for fmupdate - fwm-setting
	return
}

// ReadFmupdateFwmSetting API operation for FortiManager gets the Fwm Setting
// with the specified index value.
// Returns the requested Fwm Setting value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the fmupdate - fwm-setting chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadFmupdateFwmSetting(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/cli/[*]/fmupdate/fwm-setting"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}


	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateDvmdbDevice API operation for FortiManager creates a new Device.
// Returns the index value of the Device and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dvmdb - device chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateDvmdbDevice(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/dvmdb/[*]/device"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateDvmdbDevice API operation for FortiManager updates the specified Device.
// Returns the index value of the Device and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dvmdb - device chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateDvmdbDevice(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/dvmdb/[*]/device"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteDvmdbDevice API operation for FortiManager deletes the specified Device.
// Returns error for service API and SDK errors.
// See the dvmdb - device chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteDvmdbDevice(globaladom, mkey string, paralist []string) (err error) {
	path := "/dvmdb/[*]/device"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadDvmdbDevice API operation for FortiManager gets the Device
// with the specified index value.
// Returns the requested Device value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dvmdb - device chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadDvmdbDevice(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/dvmdb/[*]/device"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateDvmdbGroup API operation for FortiManager creates a new Group.
// Returns the index value of the Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dvmdb - group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateDvmdbGroup(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/dvmdb/[*]/group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateDvmdbGroup API operation for FortiManager updates the specified Group.
// Returns the index value of the Group and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dvmdb - group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateDvmdbGroup(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/dvmdb/[*]/group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteDvmdbGroup API operation for FortiManager deletes the specified Group.
// Returns error for service API and SDK errors.
// See the dvmdb - group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteDvmdbGroup(globaladom, mkey string, paralist []string) (err error) {
	path := "/dvmdb/[*]/group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadDvmdbGroup API operation for FortiManager gets the Group
// with the specified index value.
// Returns the requested Group value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dvmdb - group chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadDvmdbGroup(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/dvmdb/[*]/group"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateDvmdbScript API operation for FortiManager creates a new Script.
// Returns the index value of the Script and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dvmdb - script chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateDvmdbScript(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/dvmdb/[*]/script"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateDvmdbScript API operation for FortiManager updates the specified Script.
// Returns the index value of the Script and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dvmdb - script chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateDvmdbScript(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/dvmdb/[*]/script"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteDvmdbScript API operation for FortiManager deletes the specified Script.
// Returns error for service API and SDK errors.
// See the dvmdb - script chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteDvmdbScript(globaladom, mkey string, paralist []string) (err error) {
	path := "/dvmdb/[*]/script"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadDvmdbScript API operation for FortiManager gets the Script
// with the specified index value.
// Returns the requested Script value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dvmdb - script chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadDvmdbScript(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/dvmdb/[*]/script"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateDvmdbRevision API operation for FortiManager creates a new Revision.
// Returns the index value of the Revision and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dvmdb - revision chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateDvmdbRevision(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/dvmdb/[*]/revision"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateDvmdbRevision API operation for FortiManager updates the specified Revision.
// Returns the index value of the Revision and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dvmdb - revision chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateDvmdbRevision(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/dvmdb/[*]/revision"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteDvmdbRevision API operation for FortiManager deletes the specified Revision.
// Returns error for service API and SDK errors.
// See the dvmdb - revision chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteDvmdbRevision(globaladom, mkey string, paralist []string) (err error) {
	path := "/dvmdb/[*]/revision"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadDvmdbRevision API operation for FortiManager gets the Revision
// with the specified index value.
// Returns the requested Revision value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dvmdb - revision chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadDvmdbRevision(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/dvmdb/[*]/revision"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateDvmdbFolder API operation for FortiManager creates a new Folder.
// Returns the index value of the Folder and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dvmdb - folder chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateDvmdbFolder(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/dvmdb/[*]/folder"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateDvmdbFolder API operation for FortiManager updates the specified Folder.
// Returns the index value of the Folder and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dvmdb - folder chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateDvmdbFolder(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/dvmdb/[*]/folder"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteDvmdbFolder API operation for FortiManager deletes the specified Folder.
// Returns error for service API and SDK errors.
// See the dvmdb - folder chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteDvmdbFolder(globaladom, mkey string, paralist []string) (err error) {
	path := "/dvmdb/[*]/folder"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadDvmdbFolder API operation for FortiManager gets the Folder
// with the specified index value.
// Returns the requested Folder value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dvmdb - folder chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadDvmdbFolder(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/dvmdb/[*]/folder"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}

// CreateDvmdbAdom API operation for FortiManager creates a new Adom.
// Returns the index value of the Adom and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dvmdb - adom chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) CreateDvmdbAdom(params *map[string]interface{}, globaladom string, paralist []string) (output map[string]interface{}, err error) {
	path := "/dvmdb/adom"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "add", params, output, false)
	return
}

// UpdateDvmdbAdom API operation for FortiManager updates the specified Adom.
// Returns the index value of the Adom and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dvmdb - adom chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateDvmdbAdom(params *map[string]interface{}, globaladom, mkey string, paralist []string) (output map[string]interface{}, err error) {
	path := "/dvmdb/adom"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)
	output = make(map[string]interface{})

	err = createUpdate(c, globaladom, path, "set", params, output, false)
	return
}

// DeleteDvmdbAdom API operation for FortiManager deletes the specified Adom.
// Returns error for service API and SDK errors.
// See the dvmdb - adom chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteDvmdbAdom(globaladom, mkey string, paralist []string) (err error) {
	path := "/dvmdb/adom"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	err = delete(c, globaladom, path, "delete", false)
	return
}

// ReadDvmdbAdom API operation for FortiManager gets the Adom
// with the specified index value.
// Returns the requested Adom value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the dvmdb - adom chapter in the FortiManager Handbook - CLI Reference.
func (c *FortiSDKClient) ReadDvmdbAdom(globaladom, mkey string, paralist []string) (mapTmp map[string]interface{}, err error) {
	path := "/dvmdb/adom"
	path, err = replaceParaWithValue(path, paralist)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	path += "/" + escapeURLString(mkey)

	mapTmp, err = read(c, globaladom, path, "get", false)
	return
}
