// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Enable/disable push updates, and override the default IP address and port used by FortiGuard to send antivirus and IPS push messages for clients.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFmupdateFdsSettingPushOverride() *schema.Resource {
	return &schema.Resource{
		Create: resourceFmupdateFdsSettingPushOverrideUpdate,
		Read:   resourceFmupdateFdsSettingPushOverrideRead,
		Update: resourceFmupdateFdsSettingPushOverrideUpdate,
		Delete: resourceFmupdateFdsSettingPushOverrideDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
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
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFmupdateFdsSettingPushOverrideUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectFmupdateFdsSettingPushOverride(d)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateFdsSettingPushOverride resource while getting object: %v", err)
	}

	_, err = c.UpdateFmupdateFdsSettingPushOverride(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating FmupdateFdsSettingPushOverride resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("FmupdateFdsSettingPushOverride")

	return resourceFmupdateFdsSettingPushOverrideRead(d, m)
}

func resourceFmupdateFdsSettingPushOverrideDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	err = c.DeleteFmupdateFdsSettingPushOverride(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting FmupdateFdsSettingPushOverride resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFmupdateFdsSettingPushOverrideRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadFmupdateFdsSettingPushOverride(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateFdsSettingPushOverride resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFmupdateFdsSettingPushOverride(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FmupdateFdsSettingPushOverride resource from API: %v", err)
	}
	return nil
}

func flattenFmupdateFdsSettingPushOverrideIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingPushOverridePort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFmupdateFdsSettingPushOverrideStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFmupdateFdsSettingPushOverride(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ip", flattenFmupdateFdsSettingPushOverrideIp(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "FmupdateFdsSettingPushOverride-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("port", flattenFmupdateFdsSettingPushOverridePort(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "FmupdateFdsSettingPushOverride-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("status", flattenFmupdateFdsSettingPushOverrideStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "FmupdateFdsSettingPushOverride-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenFmupdateFdsSettingPushOverrideFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFmupdateFdsSettingPushOverrideIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingPushOverridePort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFmupdateFdsSettingPushOverrideStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFmupdateFdsSettingPushOverride(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandFmupdateFdsSettingPushOverrideIp(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandFmupdateFdsSettingPushOverridePort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandFmupdateFdsSettingPushOverrideStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
