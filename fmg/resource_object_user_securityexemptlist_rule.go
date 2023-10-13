// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure rules for exempting users from captive portal authentication.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectUserSecurityExemptListRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectUserSecurityExemptListRuleCreate,
		Read:   resourceObjectUserSecurityExemptListRuleRead,
		Update: resourceObjectUserSecurityExemptListRuleUpdate,
		Delete: resourceObjectUserSecurityExemptListRuleDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
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
			"security_exempt_list": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"devices": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dstaddr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"srcaddr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectUserSecurityExemptListRuleCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	security_exempt_list := d.Get("security_exempt_list").(string)
	paradict["security_exempt_list"] = security_exempt_list

	obj, err := getObjectObjectUserSecurityExemptListRule(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectUserSecurityExemptListRule resource while getting object: %v", err)
	}

	_, err = c.CreateObjectUserSecurityExemptListRule(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectUserSecurityExemptListRule resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectUserSecurityExemptListRuleRead(d, m)
}

func resourceObjectUserSecurityExemptListRuleUpdate(d *schema.ResourceData, m interface{}) error {
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

	security_exempt_list := d.Get("security_exempt_list").(string)
	paradict["security_exempt_list"] = security_exempt_list

	obj, err := getObjectObjectUserSecurityExemptListRule(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserSecurityExemptListRule resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectUserSecurityExemptListRule(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectUserSecurityExemptListRule resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectUserSecurityExemptListRuleRead(d, m)
}

func resourceObjectUserSecurityExemptListRuleDelete(d *schema.ResourceData, m interface{}) error {
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

	security_exempt_list := d.Get("security_exempt_list").(string)
	paradict["security_exempt_list"] = security_exempt_list

	err = c.DeleteObjectUserSecurityExemptListRule(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectUserSecurityExemptListRule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectUserSecurityExemptListRuleRead(d *schema.ResourceData, m interface{}) error {
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

	security_exempt_list := d.Get("security_exempt_list").(string)
	if security_exempt_list == "" {
		security_exempt_list = importOptionChecking(m.(*FortiClient).Cfg, "security_exempt_list")
		if security_exempt_list == "" {
			return fmt.Errorf("Parameter security_exempt_list is missing")
		}
		if err = d.Set("security_exempt_list", security_exempt_list); err != nil {
			return fmt.Errorf("Error set params security_exempt_list: %v", err)
		}
	}
	paradict["security_exempt_list"] = security_exempt_list

	o, err := c.ReadObjectUserSecurityExemptListRule(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserSecurityExemptListRule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectUserSecurityExemptListRule(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectUserSecurityExemptListRule resource from API: %v", err)
	}
	return nil
}

func flattenObjectUserSecurityExemptListRuleDevices2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSecurityExemptListRuleDstaddr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSecurityExemptListRuleId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSecurityExemptListRuleService2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectUserSecurityExemptListRuleSrcaddr2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectUserSecurityExemptListRule(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("devices", flattenObjectUserSecurityExemptListRuleDevices2edl(o["devices"], d, "devices")); err != nil {
		if vv, ok := fortiAPIPatch(o["devices"], "ObjectUserSecurityExemptListRule-Devices"); ok {
			if err = d.Set("devices", vv); err != nil {
				return fmt.Errorf("Error reading devices: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading devices: %v", err)
		}
	}

	if err = d.Set("dstaddr", flattenObjectUserSecurityExemptListRuleDstaddr2edl(o["dstaddr"], d, "dstaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr"], "ObjectUserSecurityExemptListRule-Dstaddr"); ok {
			if err = d.Set("dstaddr", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr: %v", err)
		}
	}

	if err = d.Set("fosid", flattenObjectUserSecurityExemptListRuleId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectUserSecurityExemptListRule-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("service", flattenObjectUserSecurityExemptListRuleService2edl(o["service"], d, "service")); err != nil {
		if vv, ok := fortiAPIPatch(o["service"], "ObjectUserSecurityExemptListRule-Service"); ok {
			if err = d.Set("service", vv); err != nil {
				return fmt.Errorf("Error reading service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service: %v", err)
		}
	}

	if err = d.Set("srcaddr", flattenObjectUserSecurityExemptListRuleSrcaddr2edl(o["srcaddr"], d, "srcaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr"], "ObjectUserSecurityExemptListRule-Srcaddr"); ok {
			if err = d.Set("srcaddr", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr: %v", err)
		}
	}

	return nil
}

func flattenObjectUserSecurityExemptListRuleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectUserSecurityExemptListRuleDevices2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSecurityExemptListRuleDstaddr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSecurityExemptListRuleId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSecurityExemptListRuleService2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectUserSecurityExemptListRuleSrcaddr2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectUserSecurityExemptListRule(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("devices"); ok || d.HasChange("devices") {
		t, err := expandObjectUserSecurityExemptListRuleDevices2edl(d, v, "devices")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["devices"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr"); ok || d.HasChange("dstaddr") {
		t, err := expandObjectUserSecurityExemptListRuleDstaddr2edl(d, v, "dstaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectUserSecurityExemptListRuleId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("service"); ok || d.HasChange("service") {
		t, err := expandObjectUserSecurityExemptListRuleService2edl(d, v, "service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr"); ok || d.HasChange("srcaddr") {
		t, err := expandObjectUserSecurityExemptListRuleSrcaddr2edl(d, v, "srcaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr"] = t
		}
	}

	return &obj, nil
}
