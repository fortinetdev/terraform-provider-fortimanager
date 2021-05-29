// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Execute workspace lock, commit and unlock actions

package fortimanager

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceExecWorkspaceAction() *schema.Resource {
	return &schema.Resource{
		Create: resourceExecWorkspaceActionCreateUpdate,
		Read:   schema.Noop, //resourceExecWorkspaceActionRead,
		Update: resourceExecWorkspaceActionCreateUpdate,
		Delete: resourceExecWorkspaceActionDelete, //schema.Noop,

		Schema: map[string]*schema.Schema{
			"scopetype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "inherit",
				ValidateFunc: validation.StringInSlice([]string{
					"adom",
					"global",
					"inherit",
				}, false),
			},
			"adom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"action": &schema.Schema{
				Type:     schema.TypeString, //lockbegin, lockend
				Required: true,
			},
			"target": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"param": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"force_recreate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func execMain(d *schema.ResourceData, m interface{}, action string) (string, error) {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return "", err
	}

	target := d.Get("target").(string)
	param := d.Get("param").(string)

	idstr := "workspaceaction" + adomv + action + target + param
	idstr = strings.ReplaceAll(idstr, "/", ".")

	_, err = c.CreateUpdateExecWorkspaceAction(adomv, action, target, param)
	return idstr, err
}

func resourceExecWorkspaceActionCreateUpdate(d *schema.ResourceData, m interface{}) error {
	action := d.Get("action").(string)

	if action == "lockbegin" {
		idstr, err := execMain(d, m, "lock")
		if err != nil {
			return fmt.Errorf("Error exec workspace action: %v", err)
		}

		d.SetId(idstr)
	} else if action == "lockend" {
		_, err := execMain(d, m, "commit")
		if err != nil {
			return fmt.Errorf("Error exec workspace action: %v", err)
		}
		idstr, err := execMain(d, m, "unlock")
		if err != nil {
			return fmt.Errorf("Error exec workspace action: %v", err)
		}

		d.SetId(idstr)
	} else {
		return fmt.Errorf("Unknown action: %v", action)
	}

	return nil
}

func resourceExecWorkspaceActionDelete(d *schema.ResourceData, m interface{}) error {
	action := d.Get("action").(string)

	if action == "lockbegin" {
		_, err := execMain(d, m, "commit")
		if err != nil {
			return fmt.Errorf("Error exec workspace action: %v", err)
		}
		_, err = execMain(d, m, "unlock")
		if err != nil {
			return fmt.Errorf("Error exec workspace action: %v", err)
		}

		d.SetId("")
	} else if action == "lockend" {
		_, err := execMain(d, m, "lock")
		if err != nil {
			return fmt.Errorf("Error exec workspace action: %v", err)
		}

		d.SetId("")
	} else {
		return fmt.Errorf("Unknown action: %v", action)
	}

	return nil
}
