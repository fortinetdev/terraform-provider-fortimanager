// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure allowaccess list for mgmt and internal interfaces on managed FortiSwitch units.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSwitchControllerSecurityPolicyLocalAccess() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSwitchControllerSecurityPolicyLocalAccessCreate,
		Read:   resourceObjectSwitchControllerSecurityPolicyLocalAccessRead,
		Update: resourceObjectSwitchControllerSecurityPolicyLocalAccessUpdate,
		Delete: resourceObjectSwitchControllerSecurityPolicyLocalAccessDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"scopetype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "inherit",
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"adom",
					"global",
					"inherit",
				}, false),
			},
			"adom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"internal_allowaccess": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"mgmt_allowaccess": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceObjectSwitchControllerSecurityPolicyLocalAccessCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectSwitchControllerSecurityPolicyLocalAccess(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSwitchControllerSecurityPolicyLocalAccess resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadObjectSwitchControllerSecurityPolicyLocalAccess(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateObjectSwitchControllerSecurityPolicyLocalAccess(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating ObjectSwitchControllerSecurityPolicyLocalAccess resource: %v", err)
			}
		}
	}

	if !existing {
		_, err = c.CreateObjectSwitchControllerSecurityPolicyLocalAccess(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating ObjectSwitchControllerSecurityPolicyLocalAccess resource: %v", err)
		}

	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSwitchControllerSecurityPolicyLocalAccessRead(d, m)
}

func resourceObjectSwitchControllerSecurityPolicyLocalAccessUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectSwitchControllerSecurityPolicyLocalAccess(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerSecurityPolicyLocalAccess resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectSwitchControllerSecurityPolicyLocalAccess(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSwitchControllerSecurityPolicyLocalAccess resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectSwitchControllerSecurityPolicyLocalAccessRead(d, m)
}

func resourceObjectSwitchControllerSecurityPolicyLocalAccessDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	wsParams["adom"] = adomv

	err = c.DeleteObjectSwitchControllerSecurityPolicyLocalAccess(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSwitchControllerSecurityPolicyLocalAccess resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSwitchControllerSecurityPolicyLocalAccessRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	o, err := c.ReadObjectSwitchControllerSecurityPolicyLocalAccess(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectSwitchControllerSecurityPolicyLocalAccess resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSwitchControllerSecurityPolicyLocalAccess(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSwitchControllerSecurityPolicyLocalAccess resource from API: %v", err)
	}
	return nil
}

func flattenObjectSwitchControllerSecurityPolicyLocalAccessInternalAllowaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSwitchControllerSecurityPolicyLocalAccessMgmtAllowaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectSwitchControllerSecurityPolicyLocalAccessName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSwitchControllerSecurityPolicyLocalAccess(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("internal_allowaccess", flattenObjectSwitchControllerSecurityPolicyLocalAccessInternalAllowaccess(o["internal-allowaccess"], d, "internal_allowaccess")); err != nil {
		if vv, ok := fortiAPIPatch(o["internal-allowaccess"], "ObjectSwitchControllerSecurityPolicyLocalAccess-InternalAllowaccess"); ok {
			if err = d.Set("internal_allowaccess", vv); err != nil {
				return fmt.Errorf("Error reading internal_allowaccess: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading internal_allowaccess: %v", err)
		}
	}

	if err = d.Set("mgmt_allowaccess", flattenObjectSwitchControllerSecurityPolicyLocalAccessMgmtAllowaccess(o["mgmt-allowaccess"], d, "mgmt_allowaccess")); err != nil {
		if vv, ok := fortiAPIPatch(o["mgmt-allowaccess"], "ObjectSwitchControllerSecurityPolicyLocalAccess-MgmtAllowaccess"); ok {
			if err = d.Set("mgmt_allowaccess", vv); err != nil {
				return fmt.Errorf("Error reading mgmt_allowaccess: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mgmt_allowaccess: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectSwitchControllerSecurityPolicyLocalAccessName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectSwitchControllerSecurityPolicyLocalAccess-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenObjectSwitchControllerSecurityPolicyLocalAccessFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSwitchControllerSecurityPolicyLocalAccessInternalAllowaccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSwitchControllerSecurityPolicyLocalAccessMgmtAllowaccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectSwitchControllerSecurityPolicyLocalAccessName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSwitchControllerSecurityPolicyLocalAccess(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("internal_allowaccess"); ok || d.HasChange("internal_allowaccess") {
		t, err := expandObjectSwitchControllerSecurityPolicyLocalAccessInternalAllowaccess(d, v, "internal_allowaccess")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internal-allowaccess"] = t
		}
	}

	if v, ok := d.GetOk("mgmt_allowaccess"); ok || d.HasChange("mgmt_allowaccess") {
		t, err := expandObjectSwitchControllerSecurityPolicyLocalAccessMgmtAllowaccess(d, v, "mgmt_allowaccess")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mgmt-allowaccess"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectSwitchControllerSecurityPolicyLocalAccessName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
