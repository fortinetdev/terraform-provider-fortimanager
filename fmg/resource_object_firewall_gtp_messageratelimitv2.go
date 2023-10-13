// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Message rate limiting for GTP version 2.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallGtpMessageRateLimitV2() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallGtpMessageRateLimitV2Update,
		Read:   resourceObjectFirewallGtpMessageRateLimitV2Read,
		Update: resourceObjectFirewallGtpMessageRateLimitV2Update,
		Delete: resourceObjectFirewallGtpMessageRateLimitV2Delete,

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
			"gtp": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"create_session_request": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"delete_session_request": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"echo_request": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceObjectFirewallGtpMessageRateLimitV2Update(d *schema.ResourceData, m interface{}) error {
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

	gtp := d.Get("gtp").(string)
	paradict["gtp"] = gtp

	obj, err := getObjectObjectFirewallGtpMessageRateLimitV2(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallGtpMessageRateLimitV2 resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFirewallGtpMessageRateLimitV2(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallGtpMessageRateLimitV2 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectFirewallGtpMessageRateLimitV2")

	return resourceObjectFirewallGtpMessageRateLimitV2Read(d, m)
}

func resourceObjectFirewallGtpMessageRateLimitV2Delete(d *schema.ResourceData, m interface{}) error {
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

	gtp := d.Get("gtp").(string)
	paradict["gtp"] = gtp

	err = c.DeleteObjectFirewallGtpMessageRateLimitV2(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallGtpMessageRateLimitV2 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallGtpMessageRateLimitV2Read(d *schema.ResourceData, m interface{}) error {
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

	gtp := d.Get("gtp").(string)
	if gtp == "" {
		gtp = importOptionChecking(m.(*FortiClient).Cfg, "gtp")
		if gtp == "" {
			return fmt.Errorf("Parameter gtp is missing")
		}
		if err = d.Set("gtp", gtp); err != nil {
			return fmt.Errorf("Error set params gtp: %v", err)
		}
	}
	paradict["gtp"] = gtp

	o, err := c.ReadObjectFirewallGtpMessageRateLimitV2(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallGtpMessageRateLimitV2 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallGtpMessageRateLimitV2(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallGtpMessageRateLimitV2 resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallGtpMessageRateLimitV2CreateSessionRequest2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitV2DeleteSessionRequest2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitV2EchoRequest2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallGtpMessageRateLimitV2(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("create_session_request", flattenObjectFirewallGtpMessageRateLimitV2CreateSessionRequest2edl(o["create-session-request"], d, "create_session_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["create-session-request"], "ObjectFirewallGtpMessageRateLimitV2-CreateSessionRequest"); ok {
			if err = d.Set("create_session_request", vv); err != nil {
				return fmt.Errorf("Error reading create_session_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading create_session_request: %v", err)
		}
	}

	if err = d.Set("delete_session_request", flattenObjectFirewallGtpMessageRateLimitV2DeleteSessionRequest2edl(o["delete-session-request"], d, "delete_session_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["delete-session-request"], "ObjectFirewallGtpMessageRateLimitV2-DeleteSessionRequest"); ok {
			if err = d.Set("delete_session_request", vv); err != nil {
				return fmt.Errorf("Error reading delete_session_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading delete_session_request: %v", err)
		}
	}

	if err = d.Set("echo_request", flattenObjectFirewallGtpMessageRateLimitV2EchoRequest2edl(o["echo-request"], d, "echo_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["echo-request"], "ObjectFirewallGtpMessageRateLimitV2-EchoRequest"); ok {
			if err = d.Set("echo_request", vv); err != nil {
				return fmt.Errorf("Error reading echo_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading echo_request: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallGtpMessageRateLimitV2FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallGtpMessageRateLimitV2CreateSessionRequest2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitV2DeleteSessionRequest2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitV2EchoRequest2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallGtpMessageRateLimitV2(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("create_session_request"); ok || d.HasChange("create_session_request") {
		t, err := expandObjectFirewallGtpMessageRateLimitV2CreateSessionRequest2edl(d, v, "create_session_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["create-session-request"] = t
		}
	}

	if v, ok := d.GetOk("delete_session_request"); ok || d.HasChange("delete_session_request") {
		t, err := expandObjectFirewallGtpMessageRateLimitV2DeleteSessionRequest2edl(d, v, "delete_session_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["delete-session-request"] = t
		}
	}

	if v, ok := d.GetOk("echo_request"); ok || d.HasChange("echo_request") {
		t, err := expandObjectFirewallGtpMessageRateLimitV2EchoRequest2edl(d, v, "echo_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["echo-request"] = t
		}
	}

	return &obj, nil
}
