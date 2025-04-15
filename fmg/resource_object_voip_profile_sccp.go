// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: SCCP.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectVoipProfileSccp() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectVoipProfileSccpUpdate,
		Read:   resourceObjectVoipProfileSccpRead,
		Update: resourceObjectVoipProfileSccpUpdate,
		Delete: resourceObjectVoipProfileSccpDelete,

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
			"block_mcast": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_call_summary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_violations": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_calls": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"verify_header": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectVoipProfileSccpUpdate(d *schema.ResourceData, m interface{}) error {
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

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	obj, err := getObjectObjectVoipProfileSccp(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVoipProfileSccp resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectVoipProfileSccp(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectVoipProfileSccp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectVoipProfileSccp")

	return resourceObjectVoipProfileSccpRead(d, m)
}

func resourceObjectVoipProfileSccpDelete(d *schema.ResourceData, m interface{}) error {
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

	profile := d.Get("profile").(string)
	paradict["profile"] = profile

	wsParams["adom"] = adomv

	err = c.DeleteObjectVoipProfileSccp(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectVoipProfileSccp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectVoipProfileSccpRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectVoipProfileSccp(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVoipProfileSccp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectVoipProfileSccp(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectVoipProfileSccp resource from API: %v", err)
	}
	return nil
}

func flattenObjectVoipProfileSccpBlockMcast2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSccpLogCallSummary2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSccpLogViolations2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSccpMaxCalls2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSccpStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectVoipProfileSccpVerifyHeader2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectVoipProfileSccp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("block_mcast", flattenObjectVoipProfileSccpBlockMcast2edl(o["block-mcast"], d, "block_mcast")); err != nil {
		if vv, ok := fortiAPIPatch(o["block-mcast"], "ObjectVoipProfileSccp-BlockMcast"); ok {
			if err = d.Set("block_mcast", vv); err != nil {
				return fmt.Errorf("Error reading block_mcast: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading block_mcast: %v", err)
		}
	}

	if err = d.Set("log_call_summary", flattenObjectVoipProfileSccpLogCallSummary2edl(o["log-call-summary"], d, "log_call_summary")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-call-summary"], "ObjectVoipProfileSccp-LogCallSummary"); ok {
			if err = d.Set("log_call_summary", vv); err != nil {
				return fmt.Errorf("Error reading log_call_summary: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_call_summary: %v", err)
		}
	}

	if err = d.Set("log_violations", flattenObjectVoipProfileSccpLogViolations2edl(o["log-violations"], d, "log_violations")); err != nil {
		if vv, ok := fortiAPIPatch(o["log-violations"], "ObjectVoipProfileSccp-LogViolations"); ok {
			if err = d.Set("log_violations", vv); err != nil {
				return fmt.Errorf("Error reading log_violations: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log_violations: %v", err)
		}
	}

	if err = d.Set("max_calls", flattenObjectVoipProfileSccpMaxCalls2edl(o["max-calls"], d, "max_calls")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-calls"], "ObjectVoipProfileSccp-MaxCalls"); ok {
			if err = d.Set("max_calls", vv); err != nil {
				return fmt.Errorf("Error reading max_calls: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_calls: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectVoipProfileSccpStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectVoipProfileSccp-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("verify_header", flattenObjectVoipProfileSccpVerifyHeader2edl(o["verify-header"], d, "verify_header")); err != nil {
		if vv, ok := fortiAPIPatch(o["verify-header"], "ObjectVoipProfileSccp-VerifyHeader"); ok {
			if err = d.Set("verify_header", vv); err != nil {
				return fmt.Errorf("Error reading verify_header: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading verify_header: %v", err)
		}
	}

	return nil
}

func flattenObjectVoipProfileSccpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectVoipProfileSccpBlockMcast2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSccpLogCallSummary2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSccpLogViolations2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSccpMaxCalls2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSccpStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectVoipProfileSccpVerifyHeader2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectVoipProfileSccp(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("block_mcast"); ok || d.HasChange("block_mcast") {
		t, err := expandObjectVoipProfileSccpBlockMcast2edl(d, v, "block_mcast")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-mcast"] = t
		}
	}

	if v, ok := d.GetOk("log_call_summary"); ok || d.HasChange("log_call_summary") {
		t, err := expandObjectVoipProfileSccpLogCallSummary2edl(d, v, "log_call_summary")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-call-summary"] = t
		}
	}

	if v, ok := d.GetOk("log_violations"); ok || d.HasChange("log_violations") {
		t, err := expandObjectVoipProfileSccpLogViolations2edl(d, v, "log_violations")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-violations"] = t
		}
	}

	if v, ok := d.GetOk("max_calls"); ok || d.HasChange("max_calls") {
		t, err := expandObjectVoipProfileSccpMaxCalls2edl(d, v, "max_calls")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-calls"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectVoipProfileSccpStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("verify_header"); ok || d.HasChange("verify_header") {
		t, err := expandObjectVoipProfileSccpVerifyHeader2edl(d, v, "verify_header")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["verify-header"] = t
		}
	}

	return &obj, nil
}
