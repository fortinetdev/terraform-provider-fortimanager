// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure UDP timeout profile.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSystemNpuUdpTimeoutProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemNpuUdpTimeoutProfileCreate,
		Read:   resourceObjectSystemNpuUdpTimeoutProfileRead,
		Update: resourceObjectSystemNpuUdpTimeoutProfileUpdate,
		Delete: resourceObjectSystemNpuUdpTimeoutProfileDelete,

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
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"udp_idle": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceObjectSystemNpuUdpTimeoutProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSystemNpuUdpTimeoutProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemNpuUdpTimeoutProfile resource while getting object: %v", err)
	}

	_, err = c.CreateObjectSystemNpuUdpTimeoutProfile(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectSystemNpuUdpTimeoutProfile resource: %v", err)
	}

	d.SetId(getStringKey(d, ""))

	return resourceObjectSystemNpuUdpTimeoutProfileRead(d, m)
}

func resourceObjectSystemNpuUdpTimeoutProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSystemNpuUdpTimeoutProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuUdpTimeoutProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSystemNpuUdpTimeoutProfile(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuUdpTimeoutProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, ""))

	return resourceObjectSystemNpuUdpTimeoutProfileRead(d, m)
}

func resourceObjectSystemNpuUdpTimeoutProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectSystemNpuUdpTimeoutProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemNpuUdpTimeoutProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemNpuUdpTimeoutProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectSystemNpuUdpTimeoutProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuUdpTimeoutProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemNpuUdpTimeoutProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuUdpTimeoutProfile resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemNpuUdpTimeoutProfileId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuUdpTimeoutProfileUdpIdle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemNpuUdpTimeoutProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("fosid", flattenObjectSystemNpuUdpTimeoutProfileId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectSystemNpuUdpTimeoutProfile-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("udp_idle", flattenObjectSystemNpuUdpTimeoutProfileUdpIdle(o["udp-idle"], d, "udp_idle")); err != nil {
		if vv, ok := fortiAPIPatch(o["udp-idle"], "ObjectSystemNpuUdpTimeoutProfile-UdpIdle"); ok {
			if err = d.Set("udp_idle", vv); err != nil {
				return fmt.Errorf("Error reading udp_idle: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading udp_idle: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemNpuUdpTimeoutProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemNpuUdpTimeoutProfileId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuUdpTimeoutProfileUdpIdle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemNpuUdpTimeoutProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectSystemNpuUdpTimeoutProfileId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("udp_idle"); ok || d.HasChange("udp_idle") {
		t, err := expandObjectSystemNpuUdpTimeoutProfileUdpIdle(d, v, "udp_idle")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["udp-idle"] = t
		}
	}

	return &obj, nil
}
