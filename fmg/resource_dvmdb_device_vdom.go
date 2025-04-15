// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Device VDOM table.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceDvmdbDeviceVdom() *schema.Resource {
	return &schema.Resource{
		Create: resourceDvmdbDeviceVdomCreate,
		Read:   resourceDvmdbDeviceVdomRead,
		Update: resourceDvmdbDeviceVdomUpdate,
		Delete: resourceDvmdbDeviceVdomDelete,

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
			"device_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"metafields": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"opmode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rtm_prof_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vdom_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vpn_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceDvmdbDeviceVdomCreate(d *schema.ResourceData, m interface{}) error {
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

	device_name := d.Get("device_name").(string)
	paradict["device"] = device_name

	obj, err := getObjectDvmdbDeviceVdom(d)
	if err != nil {
		return fmt.Errorf("Error creating DvmdbDeviceVdom resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateDvmdbDeviceVdom(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating DvmdbDeviceVdom resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceDvmdbDeviceVdomRead(d, m)
}

func resourceDvmdbDeviceVdomUpdate(d *schema.ResourceData, m interface{}) error {
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

	device_name := d.Get("device_name").(string)
	paradict["device"] = device_name

	obj, err := getObjectDvmdbDeviceVdom(d)
	if err != nil {
		return fmt.Errorf("Error updating DvmdbDeviceVdom resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateDvmdbDeviceVdom(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating DvmdbDeviceVdom resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceDvmdbDeviceVdomRead(d, m)
}

func resourceDvmdbDeviceVdomDelete(d *schema.ResourceData, m interface{}) error {
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

	device_name := d.Get("device_name").(string)
	paradict["device"] = device_name

	wsParams["adom"] = adomv

	err = c.DeleteDvmdbDeviceVdom(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting DvmdbDeviceVdom resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceDvmdbDeviceVdomRead(d *schema.ResourceData, m interface{}) error {
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

	device_name := d.Get("device_name").(string)
	if device_name == "" {
		device_name = importOptionChecking(m.(*FortiClient).Cfg, "device_name")
		if device_name == "" {
			return fmt.Errorf("Parameter device_name is missing")
		}
		if err = d.Set("device_name", device_name); err != nil {
			return fmt.Errorf("Error set params device_name: %v", err)
		}
	}
	paradict["device"] = device_name

	o, err := c.ReadDvmdbDeviceVdom(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading DvmdbDeviceVdom resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectDvmdbDeviceVdom(d, o)
	if err != nil {
		return fmt.Errorf("Error reading DvmdbDeviceVdom resource from API: %v", err)
	}
	return nil
}

func flattenDvmdbDeviceVdomComments2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceVdomMetaFields2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceVdomName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceVdomOpmode2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceVdomRtmProfId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceVdomStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceVdomVdomType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceVdomVpnId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectDvmdbDeviceVdom(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("comments", flattenDvmdbDeviceVdomComments2edl(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "DvmdbDeviceVdom-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("metafields", flattenDvmdbDeviceVdomMetaFields2edl(o["meta fields"], d, "metafields")); err != nil {
		if vv, ok := fortiAPIPatch(o["meta fields"], "DvmdbDeviceVdom-MetaFields"); ok {
			if err = d.Set("metafields", vv); err != nil {
				return fmt.Errorf("Error reading metafields: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading metafields: %v", err)
		}
	}

	if err = d.Set("name", flattenDvmdbDeviceVdomName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "DvmdbDeviceVdom-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("opmode", flattenDvmdbDeviceVdomOpmode2edl(o["opmode"], d, "opmode")); err != nil {
		if vv, ok := fortiAPIPatch(o["opmode"], "DvmdbDeviceVdom-Opmode"); ok {
			if err = d.Set("opmode", vv); err != nil {
				return fmt.Errorf("Error reading opmode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading opmode: %v", err)
		}
	}

	if err = d.Set("rtm_prof_id", flattenDvmdbDeviceVdomRtmProfId2edl(o["rtm_prof_id"], d, "rtm_prof_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["rtm_prof_id"], "DvmdbDeviceVdom-RtmProfId"); ok {
			if err = d.Set("rtm_prof_id", vv); err != nil {
				return fmt.Errorf("Error reading rtm_prof_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rtm_prof_id: %v", err)
		}
	}

	if err = d.Set("status", flattenDvmdbDeviceVdomStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "DvmdbDeviceVdom-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("vdom_type", flattenDvmdbDeviceVdomVdomType2edl(o["vdom_type"], d, "vdom_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["vdom_type"], "DvmdbDeviceVdom-VdomType"); ok {
			if err = d.Set("vdom_type", vv); err != nil {
				return fmt.Errorf("Error reading vdom_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vdom_type: %v", err)
		}
	}

	if err = d.Set("vpn_id", flattenDvmdbDeviceVdomVpnId2edl(o["vpn_id"], d, "vpn_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["vpn_id"], "DvmdbDeviceVdom-VpnId"); ok {
			if err = d.Set("vpn_id", vv); err != nil {
				return fmt.Errorf("Error reading vpn_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vpn_id: %v", err)
		}
	}

	return nil
}

func flattenDvmdbDeviceVdomFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandDvmdbDeviceVdomComments2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceVdomMetaFields2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceVdomName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceVdomOpmode2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceVdomRtmProfId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceVdomStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceVdomVdomType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceVdomVpnId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectDvmdbDeviceVdom(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandDvmdbDeviceVdomComments2edl(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("metafields"); ok || d.HasChange("metafields") {
		t, err := expandDvmdbDeviceVdomMetaFields2edl(d, v, "metafields")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["meta fields"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandDvmdbDeviceVdomName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("opmode"); ok || d.HasChange("opmode") {
		t, err := expandDvmdbDeviceVdomOpmode2edl(d, v, "opmode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["opmode"] = t
		}
	}

	if v, ok := d.GetOk("rtm_prof_id"); ok || d.HasChange("rtm_prof_id") {
		t, err := expandDvmdbDeviceVdomRtmProfId2edl(d, v, "rtm_prof_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rtm_prof_id"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandDvmdbDeviceVdomStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("vdom_type"); ok || d.HasChange("vdom_type") {
		t, err := expandDvmdbDeviceVdomVdomType2edl(d, v, "vdom_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom_type"] = t
		}
	}

	if v, ok := d.GetOk("vpn_id"); ok || d.HasChange("vpn_id") {
		t, err := expandDvmdbDeviceVdomVpnId2edl(d, v, "vpn_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vpn_id"] = t
		}
	}

	return &obj, nil
}
