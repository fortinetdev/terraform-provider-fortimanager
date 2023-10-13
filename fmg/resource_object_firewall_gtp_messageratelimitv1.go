// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Message rate limiting for GTP version 1.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallGtpMessageRateLimitV1() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallGtpMessageRateLimitV1Update,
		Read:   resourceObjectFirewallGtpMessageRateLimitV1Read,
		Update: resourceObjectFirewallGtpMessageRateLimitV1Update,
		Delete: resourceObjectFirewallGtpMessageRateLimitV1Delete,

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
			"create_pdp_request": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"delete_pdp_request": &schema.Schema{
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

func resourceObjectFirewallGtpMessageRateLimitV1Update(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectFirewallGtpMessageRateLimitV1(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallGtpMessageRateLimitV1 resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFirewallGtpMessageRateLimitV1(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallGtpMessageRateLimitV1 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectFirewallGtpMessageRateLimitV1")

	return resourceObjectFirewallGtpMessageRateLimitV1Read(d, m)
}

func resourceObjectFirewallGtpMessageRateLimitV1Delete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectFirewallGtpMessageRateLimitV1(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallGtpMessageRateLimitV1 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallGtpMessageRateLimitV1Read(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectFirewallGtpMessageRateLimitV1(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallGtpMessageRateLimitV1 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallGtpMessageRateLimitV1(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallGtpMessageRateLimitV1 resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallGtpMessageRateLimitV1CreatePdpRequest2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitV1DeletePdpRequest2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallGtpMessageRateLimitV1EchoRequest2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallGtpMessageRateLimitV1(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("create_pdp_request", flattenObjectFirewallGtpMessageRateLimitV1CreatePdpRequest2edl(o["create-pdp-request"], d, "create_pdp_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["create-pdp-request"], "ObjectFirewallGtpMessageRateLimitV1-CreatePdpRequest"); ok {
			if err = d.Set("create_pdp_request", vv); err != nil {
				return fmt.Errorf("Error reading create_pdp_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading create_pdp_request: %v", err)
		}
	}

	if err = d.Set("delete_pdp_request", flattenObjectFirewallGtpMessageRateLimitV1DeletePdpRequest2edl(o["delete-pdp-request"], d, "delete_pdp_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["delete-pdp-request"], "ObjectFirewallGtpMessageRateLimitV1-DeletePdpRequest"); ok {
			if err = d.Set("delete_pdp_request", vv); err != nil {
				return fmt.Errorf("Error reading delete_pdp_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading delete_pdp_request: %v", err)
		}
	}

	if err = d.Set("echo_request", flattenObjectFirewallGtpMessageRateLimitV1EchoRequest2edl(o["echo-request"], d, "echo_request")); err != nil {
		if vv, ok := fortiAPIPatch(o["echo-request"], "ObjectFirewallGtpMessageRateLimitV1-EchoRequest"); ok {
			if err = d.Set("echo_request", vv); err != nil {
				return fmt.Errorf("Error reading echo_request: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading echo_request: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallGtpMessageRateLimitV1FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallGtpMessageRateLimitV1CreatePdpRequest2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitV1DeletePdpRequest2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallGtpMessageRateLimitV1EchoRequest2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallGtpMessageRateLimitV1(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("create_pdp_request"); ok || d.HasChange("create_pdp_request") {
		t, err := expandObjectFirewallGtpMessageRateLimitV1CreatePdpRequest2edl(d, v, "create_pdp_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["create-pdp-request"] = t
		}
	}

	if v, ok := d.GetOk("delete_pdp_request"); ok || d.HasChange("delete_pdp_request") {
		t, err := expandObjectFirewallGtpMessageRateLimitV1DeletePdpRequest2edl(d, v, "delete_pdp_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["delete-pdp-request"] = t
		}
	}

	if v, ok := d.GetOk("echo_request"); ok || d.HasChange("echo_request") {
		t, err := expandObjectFirewallGtpMessageRateLimitV1EchoRequest2edl(d, v, "echo_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["echo-request"] = t
		}
	}

	return &obj, nil
}
