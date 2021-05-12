// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IPv4 to IPv6 virtual IP groups.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectFirewallVipgrp46() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallVipgrp46Create,
		Read:   resourceObjectFirewallVipgrp46Read,
		Update: resourceObjectFirewallVipgrp46Update,
		Delete: resourceObjectFirewallVipgrp46Delete,

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
			"color": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"member": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceObjectFirewallVipgrp46Create(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectFirewallVipgrp46(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallVipgrp46 resource while getting object: %v", err)
	}

	_, err = c.CreateObjectFirewallVipgrp46(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallVipgrp46 resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallVipgrp46Read(d, m)
}

func resourceObjectFirewallVipgrp46Update(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectFirewallVipgrp46(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallVipgrp46 resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectFirewallVipgrp46(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallVipgrp46 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallVipgrp46Read(d, m)
}

func resourceObjectFirewallVipgrp46Delete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectFirewallVipgrp46(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallVipgrp46 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallVipgrp46Read(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectFirewallVipgrp46(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallVipgrp46 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallVipgrp46(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallVipgrp46 resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallVipgrp46Color(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipgrp46Comments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipgrp46Member(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipgrp46Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallVipgrp46Uuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallVipgrp46(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("color", flattenObjectFirewallVipgrp46Color(o["color"], d, "color")); err != nil {
		if vv, ok := fortiAPIPatch(o["color"], "ObjectFirewallVipgrp46-Color"); ok {
			if err = d.Set("color", vv); err != nil {
				return fmt.Errorf("Error reading color: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}

	if err = d.Set("comments", flattenObjectFirewallVipgrp46Comments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "ObjectFirewallVipgrp46-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("member", flattenObjectFirewallVipgrp46Member(o["member"], d, "member")); err != nil {
		if vv, ok := fortiAPIPatch(o["member"], "ObjectFirewallVipgrp46-Member"); ok {
			if err = d.Set("member", vv); err != nil {
				return fmt.Errorf("Error reading member: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading member: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectFirewallVipgrp46Name(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectFirewallVipgrp46-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("uuid", flattenObjectFirewallVipgrp46Uuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "ObjectFirewallVipgrp46-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallVipgrp46FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallVipgrp46Color(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipgrp46Comments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipgrp46Member(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipgrp46Name(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallVipgrp46Uuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallVipgrp46(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("color"); ok {
		t, err := expandObjectFirewallVipgrp46Color(d, v, "color")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {
		t, err := expandObjectFirewallVipgrp46Comments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("member"); ok {
		t, err := expandObjectFirewallVipgrp46Member(d, v, "member")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["member"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectFirewallVipgrp46Name(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok {
		t, err := expandObjectFirewallVipgrp46Uuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	return &obj, nil
}