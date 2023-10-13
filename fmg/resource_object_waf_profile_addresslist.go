// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Black address list and white address list.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWafProfileAddressList() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWafProfileAddressListUpdate,
		Read:   resourceObjectWafProfileAddressListRead,
		Update: resourceObjectWafProfileAddressListUpdate,
		Delete: resourceObjectWafProfileAddressListDelete,

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
			"profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"blocked_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"blocked_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"severity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trusted_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectWafProfileAddressListUpdate(d *schema.ResourceData, m interface{}) error {
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

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	obj, err := getObjectObjectWafProfileAddressList(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWafProfileAddressList resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWafProfileAddressList(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWafProfileAddressList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectWafProfileAddressList")

	return resourceObjectWafProfileAddressListRead(d, m)
}

func resourceObjectWafProfileAddressListDelete(d *schema.ResourceData, m interface{}) error {
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

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	err = c.DeleteObjectWafProfileAddressList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWafProfileAddressList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWafProfileAddressListRead(d *schema.ResourceData, m interface{}) error {
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

	profile := d.Get("profile").(string)
	if profile == "" {
		profile = importOptionChecking(m.(*FortiClient).Cfg, "profile")
		if profile == "" {
			return fmt.Errorf("Parameter profile is missing")
		}
		if err = d.Set("profile", profile); err != nil {
			return fmt.Errorf("Error set params profile: %v", err)
		}
	}
	paradict["profile"] = profile

	o, err := c.ReadObjectWafProfileAddressList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWafProfileAddressList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWafProfileAddressList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWafProfileAddressList resource from API: %v", err)
	}
	return nil
}

func flattenObjectWafProfileAddressListBlockedAddress2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenObjectWafProfileAddressListBlockedLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileAddressListSeverity2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileAddressListStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWafProfileAddressListTrustedAddress2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func refreshObjectObjectWafProfileAddressList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("blocked_address", flattenObjectWafProfileAddressListBlockedAddress2edl(o["blocked-address"], d, "blocked_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["blocked-address"], "ObjectWafProfileAddressList-BlockedAddress"); ok {
			if err = d.Set("blocked_address", vv); err != nil {
				return fmt.Errorf("Error reading blocked_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading blocked_address: %v", err)
		}
	}

	if err = d.Set("blocked_log", flattenObjectWafProfileAddressListBlockedLog2edl(o["blocked-log"], d, "blocked_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["blocked-log"], "ObjectWafProfileAddressList-BlockedLog"); ok {
			if err = d.Set("blocked_log", vv); err != nil {
				return fmt.Errorf("Error reading blocked_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading blocked_log: %v", err)
		}
	}

	if err = d.Set("severity", flattenObjectWafProfileAddressListSeverity2edl(o["severity"], d, "severity")); err != nil {
		if vv, ok := fortiAPIPatch(o["severity"], "ObjectWafProfileAddressList-Severity"); ok {
			if err = d.Set("severity", vv); err != nil {
				return fmt.Errorf("Error reading severity: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectWafProfileAddressListStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectWafProfileAddressList-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("trusted_address", flattenObjectWafProfileAddressListTrustedAddress2edl(o["trusted-address"], d, "trusted_address")); err != nil {
		if vv, ok := fortiAPIPatch(o["trusted-address"], "ObjectWafProfileAddressList-TrustedAddress"); ok {
			if err = d.Set("trusted_address", vv); err != nil {
				return fmt.Errorf("Error reading trusted_address: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading trusted_address: %v", err)
		}
	}

	return nil
}

func flattenObjectWafProfileAddressListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWafProfileAddressListBlockedAddress2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileAddressListBlockedLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileAddressListSeverity2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileAddressListStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWafProfileAddressListTrustedAddress2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWafProfileAddressList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("blocked_address"); ok || d.HasChange("blocked_address") {
		t, err := expandObjectWafProfileAddressListBlockedAddress2edl(d, v, "blocked_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["blocked-address"] = t
		}
	}

	if v, ok := d.GetOk("blocked_log"); ok || d.HasChange("blocked_log") {
		t, err := expandObjectWafProfileAddressListBlockedLog2edl(d, v, "blocked_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["blocked-log"] = t
		}
	}

	if v, ok := d.GetOk("severity"); ok || d.HasChange("severity") {
		t, err := expandObjectWafProfileAddressListSeverity2edl(d, v, "severity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectWafProfileAddressListStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("trusted_address"); ok || d.HasChange("trusted_address") {
		t, err := expandObjectWafProfileAddressListTrustedAddress2edl(d, v, "trusted_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusted-address"] = t
		}
	}

	return &obj, nil
}
