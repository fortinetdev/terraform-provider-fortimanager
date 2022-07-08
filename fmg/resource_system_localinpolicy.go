// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: IPv4 local in policy configuration.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemLocalInPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLocalInPolicyCreate,
		Read:   resourceSystemLocalInPolicyRead,
		Update: resourceSystemLocalInPolicyUpdate,
		Delete: resourceSystemLocalInPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dst": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"intf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"src": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemLocalInPolicyCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemLocalInPolicy(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemLocalInPolicy resource while getting object: %v", err)
	}

	_, err = c.CreateSystemLocalInPolicy(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating SystemLocalInPolicy resource: %v", err)
	}

	d.SetId(getStringKey(d, ""))

	return resourceSystemLocalInPolicyRead(d, m)
}

func resourceSystemLocalInPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemLocalInPolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemLocalInPolicy resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemLocalInPolicy(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemLocalInPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, ""))

	return resourceSystemLocalInPolicyRead(d, m)
}

func resourceSystemLocalInPolicyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemLocalInPolicy(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemLocalInPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLocalInPolicyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemLocalInPolicy(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemLocalInPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemLocalInPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemLocalInPolicy resource from API: %v", err)
	}
	return nil
}

func flattenSystemLocalInPolicyAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocalInPolicyDport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocalInPolicyDst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocalInPolicyId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocalInPolicyIntf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocalInPolicyProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocalInPolicySrc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemLocalInPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("action", flattenSystemLocalInPolicyAction(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "SystemLocalInPolicy-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("dport", flattenSystemLocalInPolicyDport(o["dport"], d, "dport")); err != nil {
		if vv, ok := fortiAPIPatch(o["dport"], "SystemLocalInPolicy-Dport"); ok {
			if err = d.Set("dport", vv); err != nil {
				return fmt.Errorf("Error reading dport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dport: %v", err)
		}
	}

	if err = d.Set("dst", flattenSystemLocalInPolicyDst(o["dst"], d, "dst")); err != nil {
		if vv, ok := fortiAPIPatch(o["dst"], "SystemLocalInPolicy-Dst"); ok {
			if err = d.Set("dst", vv); err != nil {
				return fmt.Errorf("Error reading dst: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dst: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystemLocalInPolicyId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemLocalInPolicy-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("intf", flattenSystemLocalInPolicyIntf(o["intf"], d, "intf")); err != nil {
		if vv, ok := fortiAPIPatch(o["intf"], "SystemLocalInPolicy-Intf"); ok {
			if err = d.Set("intf", vv); err != nil {
				return fmt.Errorf("Error reading intf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading intf: %v", err)
		}
	}

	if err = d.Set("protocol", flattenSystemLocalInPolicyProtocol(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "SystemLocalInPolicy-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("src", flattenSystemLocalInPolicySrc(o["src"], d, "src")); err != nil {
		if vv, ok := fortiAPIPatch(o["src"], "SystemLocalInPolicy-Src"); ok {
			if err = d.Set("src", vv); err != nil {
				return fmt.Errorf("Error reading src: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src: %v", err)
		}
	}

	return nil
}

func flattenSystemLocalInPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemLocalInPolicyAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocalInPolicyDport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocalInPolicyDst(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocalInPolicyId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocalInPolicyIntf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocalInPolicyProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocalInPolicySrc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemLocalInPolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandSystemLocalInPolicyAction(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("dport"); ok || d.HasChange("dport") {
		t, err := expandSystemLocalInPolicyDport(d, v, "dport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dport"] = t
		}
	}

	if v, ok := d.GetOk("dst"); ok || d.HasChange("dst") {
		t, err := expandSystemLocalInPolicyDst(d, v, "dst")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("id") {
		t, err := expandSystemLocalInPolicyId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("intf"); ok || d.HasChange("intf") {
		t, err := expandSystemLocalInPolicyIntf(d, v, "intf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["intf"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok || d.HasChange("protocol") {
		t, err := expandSystemLocalInPolicyProtocol(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("src"); ok || d.HasChange("src") {
		t, err := expandSystemLocalInPolicySrc(d, v, "src")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src"] = t
		}
	}

	return &obj, nil
}
