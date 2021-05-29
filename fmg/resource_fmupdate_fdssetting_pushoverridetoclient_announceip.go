// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Announce IP addresses for the device.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFmupdateFdsSettingPushOverrideToClientAnnounceIp() *schema.Resource {
	return &schema.Resource{
		Create: resourceFmupdateFdsSettingPushOverrideToClientAnnounceIpCreate,
		Read:   resourceFmupdateFdsSettingPushOverrideToClientAnnounceIpRead,
		Update: resourceFmupdateFdsSettingPushOverrideToClientAnnounceIpUpdate,
		Delete: resourceFmupdateFdsSettingPushOverrideToClientAnnounceIpDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFmupdateFdsSettingPushOverrideToClientAnnounceIpCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectFmupdateFdsSettingPushOverrideToClientAnnounceIp(d)
	if err != nil {
		return fmt.Errorf("Error creating FmupdateFdsSettingPushOverrideToClientAnnounceIp resource while getting object: %v", err)
	}

	_, err = c.CreateFmupdateFdsSettingPushOverrideToClientAnnounceIp(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating FmupdateFdsSettingPushOverrideToClientAnnounceIp resource: %v", err)
	}

	d.SetId(getStringKey(d, ""))

	return resourceFmupdateFdsSettingPushOverrideToClientAnnounceIpRead(d, m)
}

func resourceFmupdateFdsSettingPushOverrideToClientAnnounceIpUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectFmupdateFdsSettingPushOverrideToClientAnnounceIp(d)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateFdsSettingPushOverrideToClientAnnounceIp resource while getting object: %v", err)
	}

	_, err = c.UpdateFmupdateFdsSettingPushOverrideToClientAnnounceIp(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateFdsSettingPushOverrideToClientAnnounceIp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, ""))

	return resourceFmupdateFdsSettingPushOverrideToClientAnnounceIpRead(d, m)
}

func resourceFmupdateFdsSettingPushOverrideToClientAnnounceIpDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteFmupdateFdsSettingPushOverrideToClientAnnounceIp(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting FmupdateFdsSettingPushOverrideToClientAnnounceIp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFmupdateFdsSettingPushOverrideToClientAnnounceIpRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadFmupdateFdsSettingPushOverrideToClientAnnounceIp(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateFdsSettingPushOverrideToClientAnnounceIp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFmupdateFdsSettingPushOverrideToClientAnnounceIp(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateFdsSettingPushOverrideToClientAnnounceIp resource from API: %v", err)
	}
	return nil
}

func flattenFmupdateFdsSettingPushOverrideToClientAnnounceIpId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingPushOverrideToClientAnnounceIpIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingPushOverrideToClientAnnounceIpPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFmupdateFdsSettingPushOverrideToClientAnnounceIp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenFmupdateFdsSettingPushOverrideToClientAnnounceIpId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "FmupdateFdsSettingPushOverrideToClientAnnounceIp-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ip", flattenFmupdateFdsSettingPushOverrideToClientAnnounceIpIp(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "FmupdateFdsSettingPushOverrideToClientAnnounceIp-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("port", flattenFmupdateFdsSettingPushOverrideToClientAnnounceIpPort(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "FmupdateFdsSettingPushOverrideToClientAnnounceIp-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	return nil
}

func flattenFmupdateFdsSettingPushOverrideToClientAnnounceIpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFmupdateFdsSettingPushOverrideToClientAnnounceIpId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingPushOverrideToClientAnnounceIpIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingPushOverrideToClientAnnounceIpPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFmupdateFdsSettingPushOverrideToClientAnnounceIp(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok {
		t, err := expandFmupdateFdsSettingPushOverrideToClientAnnounceIpId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok {
		t, err := expandFmupdateFdsSettingPushOverrideToClientAnnounceIpIp(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok {
		t, err := expandFmupdateFdsSettingPushOverrideToClientAnnounceIpPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	return &obj, nil
}
