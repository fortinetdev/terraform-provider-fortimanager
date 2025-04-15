// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure Mail signature.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallProfileProtocolOptionsMailSignature() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallProfileProtocolOptionsMailSignatureUpdate,
		Read:   resourceObjectFirewallProfileProtocolOptionsMailSignatureRead,
		Update: resourceObjectFirewallProfileProtocolOptionsMailSignatureUpdate,
		Delete: resourceObjectFirewallProfileProtocolOptionsMailSignatureDelete,

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
			"signature": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectFirewallProfileProtocolOptionsMailSignatureUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectFirewallProfileProtocolOptionsMailSignature(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallProfileProtocolOptionsMailSignature resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectFirewallProfileProtocolOptionsMailSignature(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallProfileProtocolOptionsMailSignature resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectFirewallProfileProtocolOptionsMailSignature")

	return resourceObjectFirewallProfileProtocolOptionsMailSignatureRead(d, m)
}

func resourceObjectFirewallProfileProtocolOptionsMailSignatureDelete(d *schema.ResourceData, m interface{}) error {
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

	wsParams["adom"] = adomv

	err = c.DeleteObjectFirewallProfileProtocolOptionsMailSignature(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallProfileProtocolOptionsMailSignature resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallProfileProtocolOptionsMailSignatureRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectFirewallProfileProtocolOptionsMailSignature(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallProfileProtocolOptionsMailSignature resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallProfileProtocolOptionsMailSignature(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallProfileProtocolOptionsMailSignature resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallProfileProtocolOptionsMailSignatureSignature2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallProfileProtocolOptionsMailSignatureStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallProfileProtocolOptionsMailSignature(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("signature", flattenObjectFirewallProfileProtocolOptionsMailSignatureSignature2edl(o["signature"], d, "signature")); err != nil {
		if vv, ok := fortiAPIPatch(o["signature"], "ObjectFirewallProfileProtocolOptionsMailSignature-Signature"); ok {
			if err = d.Set("signature", vv); err != nil {
				return fmt.Errorf("Error reading signature: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading signature: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectFirewallProfileProtocolOptionsMailSignatureStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectFirewallProfileProtocolOptionsMailSignature-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallProfileProtocolOptionsMailSignatureFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallProfileProtocolOptionsMailSignatureSignature2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallProfileProtocolOptionsMailSignatureStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallProfileProtocolOptionsMailSignature(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("signature"); ok || d.HasChange("signature") {
		t, err := expandObjectFirewallProfileProtocolOptionsMailSignatureSignature2edl(d, v, "signature")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["signature"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectFirewallProfileProtocolOptionsMailSignatureStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
