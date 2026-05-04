// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: ObjectFirewall ProfileProtocolOptionsProxyRedirect

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallProfileProtocolOptionsProxyRedirect() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallProfileProtocolOptionsProxyRedirectUpdate,
		Read:   resourceObjectFirewallProfileProtocolOptionsProxyRedirectRead,
		Update: resourceObjectFirewallProfileProtocolOptionsProxyRedirectUpdate,
		Delete: resourceObjectFirewallProfileProtocolOptionsProxyRedirectDelete,

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
			"profile_protocol_options": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ports": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeInt},
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

func resourceObjectFirewallProfileProtocolOptionsProxyRedirectUpdate(d *schema.ResourceData, m interface{}) error {
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

	profile_protocol_options := d.Get("profile_protocol_options").(string)
	paradict["profile_protocol_options"] = profile_protocol_options

	obj, err := getObjectObjectFirewallProfileProtocolOptionsProxyRedirect(d, false)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallProfileProtocolOptionsProxyRedirect resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectFirewallProfileProtocolOptionsProxyRedirect(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallProfileProtocolOptionsProxyRedirect resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectFirewallProfileProtocolOptionsProxyRedirect")

	return resourceObjectFirewallProfileProtocolOptionsProxyRedirectRead(d, m)
}

func resourceObjectFirewallProfileProtocolOptionsProxyRedirectDelete(d *schema.ResourceData, m interface{}) error {
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

	profile_protocol_options := d.Get("profile_protocol_options").(string)
	paradict["profile_protocol_options"] = profile_protocol_options

	obj, err := getObjectObjectFirewallProfileProtocolOptionsProxyRedirect(d, true)

	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallProfileProtocolOptionsProxyRedirect resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectFirewallProfileProtocolOptionsProxyRedirect(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error clearing ObjectFirewallProfileProtocolOptionsProxyRedirect resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallProfileProtocolOptionsProxyRedirectRead(d *schema.ResourceData, m interface{}) error {
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

	profile_protocol_options := d.Get("profile_protocol_options").(string)
	if profile_protocol_options == "" {
		profile_protocol_options = importOptionChecking(m.(*FortiClient).Cfg, "profile_protocol_options")
		if profile_protocol_options == "" {
			return fmt.Errorf("Parameter profile_protocol_options is missing")
		}
		if err = d.Set("profile_protocol_options", profile_protocol_options); err != nil {
			return fmt.Errorf("Error set params profile_protocol_options: %v", err)
		}
	}
	paradict["profile_protocol_options"] = profile_protocol_options

	o, err := c.ReadObjectFirewallProfileProtocolOptionsProxyRedirect(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading ObjectFirewallProfileProtocolOptionsProxyRedirect resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallProfileProtocolOptionsProxyRedirect(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallProfileProtocolOptionsProxyRedirect resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallProfileProtocolOptionsProxyRedirectPorts2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenIntegerList(v)
}

func flattenObjectFirewallProfileProtocolOptionsProxyRedirectStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallProfileProtocolOptionsProxyRedirect(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("ports", flattenObjectFirewallProfileProtocolOptionsProxyRedirectPorts2edl(o["ports"], d, "ports")); err != nil {
		if vv, ok := fortiAPIPatch(o["ports"], "ObjectFirewallProfileProtocolOptionsProxyRedirect-Ports"); ok {
			if err = d.Set("ports", vv); err != nil {
				return fmt.Errorf("Error reading ports: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ports: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectFirewallProfileProtocolOptionsProxyRedirectStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectFirewallProfileProtocolOptionsProxyRedirect-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallProfileProtocolOptionsProxyRedirectFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallProfileProtocolOptionsProxyRedirectPorts2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandIntegerList(v.(*schema.Set).List()), nil
}

func expandObjectFirewallProfileProtocolOptionsProxyRedirectStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallProfileProtocolOptionsProxyRedirect(d *schema.ResourceData, bemptysontable bool) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ports"); ok || d.HasChange("ports") {
		t, err := expandObjectFirewallProfileProtocolOptionsProxyRedirectPorts2edl(d, v, "ports")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ports"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectFirewallProfileProtocolOptionsProxyRedirectStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
