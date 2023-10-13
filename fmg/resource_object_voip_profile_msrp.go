// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: MSRP.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectVoipProfileMsrp() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectVoipProfileMsrpUpdate,
		Read:   resourceObjectVoipProfileMsrpRead,
		Update: resourceObjectVoipProfileMsrpUpdate,
		Delete: resourceObjectVoipProfileMsrpDelete,

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
			"log_violations": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_msg_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_msg_size_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectVoipProfileMsrpUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectVoipProfileMsrp(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVoipProfileMsrp resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectVoipProfileMsrp(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVoipProfileMsrp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectVoipProfileMsrp")

	return resourceObjectVoipProfileMsrpRead(d, m)
}

func resourceObjectVoipProfileMsrpDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectVoipProfileMsrp(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectVoipProfileMsrp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectVoipProfileMsrpRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectVoipProfileMsrp(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVoipProfileMsrp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectVoipProfileMsrp(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVoipProfileMsrp resource from API: %v", err)
	}
	return nil
}

func flattenObjectVoipProfileMsrpLogViolations2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileMsrpMaxMsgSize2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileMsrpMaxMsgSizeAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileMsrpStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectVoipProfileMsrp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("log_violations", flattenObjectVoipProfileMsrpLogViolations2edl(o["log-violations"], d, "log_violations")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-violations"], "ObjectVoipProfileMsrp-LogViolations"); ok {
			if err = d.Set("log_violations", vv); err != nil {
				return fmt.Errorf("Error reading log_violations: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_violations: %v", err)
		}
	}

	if err = d.Set("max_msg_size", flattenObjectVoipProfileMsrpMaxMsgSize2edl(o["max-msg-size"], d, "max_msg_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-msg-size"], "ObjectVoipProfileMsrp-MaxMsgSize"); ok {
			if err = d.Set("max_msg_size", vv); err != nil {
				return fmt.Errorf("Error reading max_msg_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_msg_size: %v", err)
		}
	}

	if err = d.Set("max_msg_size_action", flattenObjectVoipProfileMsrpMaxMsgSizeAction2edl(o["max-msg-size-action"], d, "max_msg_size_action")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-msg-size-action"], "ObjectVoipProfileMsrp-MaxMsgSizeAction"); ok {
			if err = d.Set("max_msg_size_action", vv); err != nil {
				return fmt.Errorf("Error reading max_msg_size_action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_msg_size_action: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectVoipProfileMsrpStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectVoipProfileMsrp-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenObjectVoipProfileMsrpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectVoipProfileMsrpLogViolations2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileMsrpMaxMsgSize2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileMsrpMaxMsgSizeAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileMsrpStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectVoipProfileMsrp(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("log_violations"); ok || d.HasChange("log_violations") {
		t, err := expandObjectVoipProfileMsrpLogViolations2edl(d, v, "log_violations")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-violations"] = t
		}
	}

	if v, ok := d.GetOk("max_msg_size"); ok || d.HasChange("max_msg_size") {
		t, err := expandObjectVoipProfileMsrpMaxMsgSize2edl(d, v, "max_msg_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-msg-size"] = t
		}
	}

	if v, ok := d.GetOk("max_msg_size_action"); ok || d.HasChange("max_msg_size_action") {
		t, err := expandObjectVoipProfileMsrpMaxMsgSizeAction2edl(d, v, "max_msg_size_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-msg-size-action"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectVoipProfileMsrpStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
