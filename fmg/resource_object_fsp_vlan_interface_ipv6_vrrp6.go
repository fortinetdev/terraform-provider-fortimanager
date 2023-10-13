// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: IPv6 VRRP configuration.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFspVlanInterfaceIpv6Vrrp6() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFspVlanInterfaceIpv6Vrrp6Create,
		Read:   resourceObjectFspVlanInterfaceIpv6Vrrp6Read,
		Update: resourceObjectFspVlanInterfaceIpv6Vrrp6Update,
		Delete: resourceObjectFspVlanInterfaceIpv6Vrrp6Delete,

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
			"vlan": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"accept_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"adv_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"preempt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"start_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vrdst6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vrgrp": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vrid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"vrip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectFspVlanInterfaceIpv6Vrrp6Create(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	vlan := d.Get("vlan").(string)
	paradict["vlan"] = vlan

	obj, err := getObjectObjectFspVlanInterfaceIpv6Vrrp6(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFspVlanInterfaceIpv6Vrrp6 resource while getting object: %v", err)
	}

	_, err = c.CreateObjectFspVlanInterfaceIpv6Vrrp6(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectFspVlanInterfaceIpv6Vrrp6 resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "vrid")))

	return resourceObjectFspVlanInterfaceIpv6Vrrp6Read(d, m)
}

func resourceObjectFspVlanInterfaceIpv6Vrrp6Update(d *schema.ResourceData, m interface{}) error {
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

	vlan := d.Get("vlan").(string)
	paradict["vlan"] = vlan

	obj, err := getObjectObjectFspVlanInterfaceIpv6Vrrp6(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFspVlanInterfaceIpv6Vrrp6 resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFspVlanInterfaceIpv6Vrrp6(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFspVlanInterfaceIpv6Vrrp6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "vrid")))

	return resourceObjectFspVlanInterfaceIpv6Vrrp6Read(d, m)
}

func resourceObjectFspVlanInterfaceIpv6Vrrp6Delete(d *schema.ResourceData, m interface{}) error {
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

	vlan := d.Get("vlan").(string)
	paradict["vlan"] = vlan

	err = c.DeleteObjectFspVlanInterfaceIpv6Vrrp6(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFspVlanInterfaceIpv6Vrrp6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFspVlanInterfaceIpv6Vrrp6Read(d *schema.ResourceData, m interface{}) error {
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

	vlan := d.Get("vlan").(string)
	if vlan == "" {
		vlan = importOptionChecking(m.(*FortiClient).Cfg, "vlan")
		if vlan == "" {
			return fmt.Errorf("Parameter vlan is missing")
		}
		if err = d.Set("vlan", vlan); err != nil {
			return fmt.Errorf("Error set params vlan: %v", err)
		}
	}
	paradict["vlan"] = vlan

	o, err := c.ReadObjectFspVlanInterfaceIpv6Vrrp6(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFspVlanInterfaceIpv6Vrrp6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFspVlanInterfaceIpv6Vrrp6(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFspVlanInterfaceIpv6Vrrp6 resource from API: %v", err)
	}
	return nil
}

func flattenObjectFspVlanInterfaceIpv6Vrrp6AcceptMode4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Vrrp6AdvInterval4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Vrrp6Preempt4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Vrrp6Priority4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Vrrp6StartTime4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Vrrp6Status4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Vrrp6Vrdst64thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Vrrp6Vrgrp4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Vrrp6Vrid4thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFspVlanInterfaceIpv6Vrrp6Vrip64thl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFspVlanInterfaceIpv6Vrrp6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("accept_mode", flattenObjectFspVlanInterfaceIpv6Vrrp6AcceptMode4thl(o["accept-mode"], d, "accept_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["accept-mode"], "ObjectFspVlanInterfaceIpv6Vrrp6-AcceptMode"); ok {
			if err = d.Set("accept_mode", vv); err != nil {
				return fmt.Errorf("Error reading accept_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading accept_mode: %v", err)
		}
	}

	if err = d.Set("adv_interval", flattenObjectFspVlanInterfaceIpv6Vrrp6AdvInterval4thl(o["adv-interval"], d, "adv_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["adv-interval"], "ObjectFspVlanInterfaceIpv6Vrrp6-AdvInterval"); ok {
			if err = d.Set("adv_interval", vv); err != nil {
				return fmt.Errorf("Error reading adv_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adv_interval: %v", err)
		}
	}

	if err = d.Set("preempt", flattenObjectFspVlanInterfaceIpv6Vrrp6Preempt4thl(o["preempt"], d, "preempt")); err != nil {
		if vv, ok := fortiAPIPatch(o["preempt"], "ObjectFspVlanInterfaceIpv6Vrrp6-Preempt"); ok {
			if err = d.Set("preempt", vv); err != nil {
				return fmt.Errorf("Error reading preempt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading preempt: %v", err)
		}
	}

	if err = d.Set("priority", flattenObjectFspVlanInterfaceIpv6Vrrp6Priority4thl(o["priority"], d, "priority")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority"], "ObjectFspVlanInterfaceIpv6Vrrp6-Priority"); ok {
			if err = d.Set("priority", vv); err != nil {
				return fmt.Errorf("Error reading priority: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("start_time", flattenObjectFspVlanInterfaceIpv6Vrrp6StartTime4thl(o["start-time"], d, "start_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-time"], "ObjectFspVlanInterfaceIpv6Vrrp6-StartTime"); ok {
			if err = d.Set("start_time", vv); err != nil {
				return fmt.Errorf("Error reading start_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_time: %v", err)
		}
	}

	if err = d.Set("status", flattenObjectFspVlanInterfaceIpv6Vrrp6Status4thl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "ObjectFspVlanInterfaceIpv6Vrrp6-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("vrdst6", flattenObjectFspVlanInterfaceIpv6Vrrp6Vrdst64thl(o["vrdst6"], d, "vrdst6")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrdst6"], "ObjectFspVlanInterfaceIpv6Vrrp6-Vrdst6"); ok {
			if err = d.Set("vrdst6", vv); err != nil {
				return fmt.Errorf("Error reading vrdst6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrdst6: %v", err)
		}
	}

	if err = d.Set("vrgrp", flattenObjectFspVlanInterfaceIpv6Vrrp6Vrgrp4thl(o["vrgrp"], d, "vrgrp")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrgrp"], "ObjectFspVlanInterfaceIpv6Vrrp6-Vrgrp"); ok {
			if err = d.Set("vrgrp", vv); err != nil {
				return fmt.Errorf("Error reading vrgrp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrgrp: %v", err)
		}
	}

	if err = d.Set("vrid", flattenObjectFspVlanInterfaceIpv6Vrrp6Vrid4thl(o["vrid"], d, "vrid")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrid"], "ObjectFspVlanInterfaceIpv6Vrrp6-Vrid"); ok {
			if err = d.Set("vrid", vv); err != nil {
				return fmt.Errorf("Error reading vrid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrid: %v", err)
		}
	}

	if err = d.Set("vrip6", flattenObjectFspVlanInterfaceIpv6Vrrp6Vrip64thl(o["vrip6"], d, "vrip6")); err != nil {
		if vv, ok := fortiAPIPatch(o["vrip6"], "ObjectFspVlanInterfaceIpv6Vrrp6-Vrip6"); ok {
			if err = d.Set("vrip6", vv); err != nil {
				return fmt.Errorf("Error reading vrip6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vrip6: %v", err)
		}
	}

	return nil
}

func flattenObjectFspVlanInterfaceIpv6Vrrp6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFspVlanInterfaceIpv6Vrrp6AcceptMode4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Vrrp6AdvInterval4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Vrrp6Preempt4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Vrrp6Priority4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Vrrp6StartTime4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Vrrp6Status4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Vrrp6Vrdst64thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Vrrp6Vrgrp4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Vrrp6Vrid4thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFspVlanInterfaceIpv6Vrrp6Vrip64thl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFspVlanInterfaceIpv6Vrrp6(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("accept_mode"); ok || d.HasChange("accept_mode") {
		t, err := expandObjectFspVlanInterfaceIpv6Vrrp6AcceptMode4thl(d, v, "accept_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["accept-mode"] = t
		}
	}

	if v, ok := d.GetOk("adv_interval"); ok || d.HasChange("adv_interval") {
		t, err := expandObjectFspVlanInterfaceIpv6Vrrp6AdvInterval4thl(d, v, "adv_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adv-interval"] = t
		}
	}

	if v, ok := d.GetOk("preempt"); ok || d.HasChange("preempt") {
		t, err := expandObjectFspVlanInterfaceIpv6Vrrp6Preempt4thl(d, v, "preempt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preempt"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok || d.HasChange("priority") {
		t, err := expandObjectFspVlanInterfaceIpv6Vrrp6Priority4thl(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("start_time"); ok || d.HasChange("start_time") {
		t, err := expandObjectFspVlanInterfaceIpv6Vrrp6StartTime4thl(d, v, "start_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-time"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandObjectFspVlanInterfaceIpv6Vrrp6Status4thl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("vrdst6"); ok || d.HasChange("vrdst6") {
		t, err := expandObjectFspVlanInterfaceIpv6Vrrp6Vrdst64thl(d, v, "vrdst6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrdst6"] = t
		}
	}

	if v, ok := d.GetOk("vrgrp"); ok || d.HasChange("vrgrp") {
		t, err := expandObjectFspVlanInterfaceIpv6Vrrp6Vrgrp4thl(d, v, "vrgrp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrgrp"] = t
		}
	}

	if v, ok := d.GetOk("vrid"); ok || d.HasChange("vrid") {
		t, err := expandObjectFspVlanInterfaceIpv6Vrrp6Vrid4thl(d, v, "vrid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrid"] = t
		}
	}

	if v, ok := d.GetOk("vrip6"); ok || d.HasChange("vrip6") {
		t, err := expandObjectFspVlanInterfaceIpv6Vrrp6Vrip64thl(d, v, "vrip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrip6"] = t
		}
	}

	return &obj, nil
}
