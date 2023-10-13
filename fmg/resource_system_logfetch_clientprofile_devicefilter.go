// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: List of device filter.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemLogFetchClientProfileDeviceFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLogFetchClientProfileDeviceFilterCreate,
		Read:   resourceSystemLogFetchClientProfileDeviceFilterRead,
		Update: resourceSystemLogFetchClientProfileDeviceFilterUpdate,
		Delete: resourceSystemLogFetchClientProfileDeviceFilterDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"client_profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fmgadom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"device": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemLogFetchClientProfileDeviceFilterCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	client_profile := d.Get("client_profile").(string)
	paradict["client_profile"] = client_profile

	obj, err := getObjectSystemLogFetchClientProfileDeviceFilter(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemLogFetchClientProfileDeviceFilter resource while getting object: %v", err)
	}

	_, err = c.CreateSystemLogFetchClientProfileDeviceFilter(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating SystemLogFetchClientProfileDeviceFilter resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemLogFetchClientProfileDeviceFilterRead(d, m)
}

func resourceSystemLogFetchClientProfileDeviceFilterUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	client_profile := d.Get("client_profile").(string)
	paradict["client_profile"] = client_profile

	obj, err := getObjectSystemLogFetchClientProfileDeviceFilter(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemLogFetchClientProfileDeviceFilter resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemLogFetchClientProfileDeviceFilter(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating SystemLogFetchClientProfileDeviceFilter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemLogFetchClientProfileDeviceFilterRead(d, m)
}

func resourceSystemLogFetchClientProfileDeviceFilterDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	client_profile := d.Get("client_profile").(string)
	paradict["client_profile"] = client_profile

	err = c.DeleteSystemLogFetchClientProfileDeviceFilter(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting SystemLogFetchClientProfileDeviceFilter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLogFetchClientProfileDeviceFilterRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	client_profile := d.Get("client_profile").(string)
	if client_profile == "" {
		client_profile = importOptionChecking(m.(*FortiClient).Cfg, "client_profile")
		if client_profile == "" {
			return fmt.Errorf("Parameter client_profile is missing")
		}
		if err = d.Set("client_profile", client_profile); err != nil {
			return fmt.Errorf("Error set params client_profile: %v", err)
		}
	}
	paradict["client_profile"] = client_profile

	o, err := c.ReadSystemLogFetchClientProfileDeviceFilter(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading SystemLogFetchClientProfileDeviceFilter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemLogFetchClientProfileDeviceFilter(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemLogFetchClientProfileDeviceFilter resource from API: %v", err)
	}
	return nil
}

func flattenSystemLogFetchClientProfileDeviceFilterAdom2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogFetchClientProfileDeviceFilterDevice2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogFetchClientProfileDeviceFilterId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogFetchClientProfileDeviceFilterVdom2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemLogFetchClientProfileDeviceFilter(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fmgadom", flattenSystemLogFetchClientProfileDeviceFilterAdom2edl(o["adom"], d, "fmgadom")); err != nil {
		if vv, ok := fortiAPIPatch(o["adom"], "SystemLogFetchClientProfileDeviceFilter-Adom"); ok {
			if err = d.Set("fmgadom", vv); err != nil {
				return fmt.Errorf("Error reading fmgadom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmgadom: %v", err)
		}
	}

	if err = d.Set("device", flattenSystemLogFetchClientProfileDeviceFilterDevice2edl(o["device"], d, "device")); err != nil {
		if vv, ok := fortiAPIPatch(o["device"], "SystemLogFetchClientProfileDeviceFilter-Device"); ok {
			if err = d.Set("device", vv); err != nil {
				return fmt.Errorf("Error reading device: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading device: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystemLogFetchClientProfileDeviceFilterId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemLogFetchClientProfileDeviceFilter-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("vdom", flattenSystemLogFetchClientProfileDeviceFilterVdom2edl(o["vdom"], d, "vdom")); err != nil {
		if vv, ok := fortiAPIPatch(o["vdom"], "SystemLogFetchClientProfileDeviceFilter-Vdom"); ok {
			if err = d.Set("vdom", vv); err != nil {
				return fmt.Errorf("Error reading vdom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	return nil
}

func flattenSystemLogFetchClientProfileDeviceFilterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemLogFetchClientProfileDeviceFilterAdom2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogFetchClientProfileDeviceFilterDevice2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogFetchClientProfileDeviceFilterId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogFetchClientProfileDeviceFilterVdom2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemLogFetchClientProfileDeviceFilter(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fmgadom"); ok || d.HasChange("fmgadom") {
		t, err := expandSystemLogFetchClientProfileDeviceFilterAdom2edl(d, v, "fmgadom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adom"] = t
		}
	}

	if v, ok := d.GetOk("device"); ok || d.HasChange("device") {
		t, err := expandSystemLogFetchClientProfileDeviceFilterDevice2edl(d, v, "device")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemLogFetchClientProfileDeviceFilterId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("vdom"); ok || d.HasChange("vdom") {
		t, err := expandSystemLogFetchClientProfileDeviceFilterVdom2edl(d, v, "vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	return &obj, nil
}
