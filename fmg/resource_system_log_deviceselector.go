// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Accept/reject devices matching specified filter types.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemLogDeviceSelector() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLogDeviceSelectorCreate,
		Read:   resourceSystemLogDeviceSelectorRead,
		Update: resourceSystemLogDeviceSelectorUpdate,
		Delete: resourceSystemLogDeviceSelectorDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"devid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"expire": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"srcip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"srcip_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemLogDeviceSelectorCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemLogDeviceSelector(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemLogDeviceSelector resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSystemLogDeviceSelector(mkey, paradict)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSystemLogDeviceSelector(obj, mkey, paradict, wsParams)
			if err != nil {
				return fmt.Errorf("Error updating SystemLogDeviceSelector resource: %v", err)
			}
		}
	}

	if !existing {
		v, err := c.CreateSystemLogDeviceSelector(obj, paradict, wsParams)
		if err != nil {
			return fmt.Errorf("Error creating SystemLogDeviceSelector resource: %v", err)
		}

		if v != nil && v["id"] != nil {
			if vidn, ok := v["id"].(float64); ok {
				d.SetId(strconv.Itoa(int(vidn)))
				return resourceSystemLogDeviceSelectorRead(d, m)
			} else {
				return fmt.Errorf("Error creating SystemLogDeviceSelector resource: %v", err)
			}
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemLogDeviceSelectorRead(d, m)
}

func resourceSystemLogDeviceSelectorUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	obj, err := getObjectSystemLogDeviceSelector(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemLogDeviceSelector resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateSystemLogDeviceSelector(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating SystemLogDeviceSelector resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceSystemLogDeviceSelectorRead(d, m)
}

func resourceSystemLogDeviceSelectorDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	wsParams["adom"] = adomv

	err = c.DeleteSystemLogDeviceSelector(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting SystemLogDeviceSelector resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLogDeviceSelectorRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)

	adomv, err := "global", fmt.Errorf("")
	paradict["adom"] = adomv

	o, err := c.ReadSystemLogDeviceSelector(mkey, paradict)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading SystemLogDeviceSelector resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemLogDeviceSelector(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemLogDeviceSelector resource from API: %v", err)
	}
	return nil
}

func flattenSystemLogDeviceSelectorAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogDeviceSelectorComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogDeviceSelectorDevid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogDeviceSelectorExpire(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogDeviceSelectorId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogDeviceSelectorSrcip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogDeviceSelectorSrcipMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLogDeviceSelectorType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemLogDeviceSelector(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("action", flattenSystemLogDeviceSelectorAction(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "SystemLogDeviceSelector-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("comment", flattenSystemLogDeviceSelectorComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "SystemLogDeviceSelector-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("devid", flattenSystemLogDeviceSelectorDevid(o["devid"], d, "devid")); err != nil {
		if vv, ok := fortiAPIPatch(o["devid"], "SystemLogDeviceSelector-Devid"); ok {
			if err = d.Set("devid", vv); err != nil {
				return fmt.Errorf("Error reading devid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading devid: %v", err)
		}
	}

	if err = d.Set("expire", flattenSystemLogDeviceSelectorExpire(o["expire"], d, "expire")); err != nil {
		if vv, ok := fortiAPIPatch(o["expire"], "SystemLogDeviceSelector-Expire"); ok {
			if err = d.Set("expire", vv); err != nil {
				return fmt.Errorf("Error reading expire: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading expire: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystemLogDeviceSelectorId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "SystemLogDeviceSelector-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("srcip", flattenSystemLogDeviceSelectorSrcip(o["srcip"], d, "srcip")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcip"], "SystemLogDeviceSelector-Srcip"); ok {
			if err = d.Set("srcip", vv); err != nil {
				return fmt.Errorf("Error reading srcip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcip: %v", err)
		}
	}

	if err = d.Set("srcip_mode", flattenSystemLogDeviceSelectorSrcipMode(o["srcip-mode"], d, "srcip_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcip-mode"], "SystemLogDeviceSelector-SrcipMode"); ok {
			if err = d.Set("srcip_mode", vv); err != nil {
				return fmt.Errorf("Error reading srcip_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcip_mode: %v", err)
		}
	}

	if err = d.Set("type", flattenSystemLogDeviceSelectorType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "SystemLogDeviceSelector-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenSystemLogDeviceSelectorFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemLogDeviceSelectorAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogDeviceSelectorComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogDeviceSelectorDevid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogDeviceSelectorExpire(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogDeviceSelectorId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogDeviceSelectorSrcip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogDeviceSelectorSrcipMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLogDeviceSelectorType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemLogDeviceSelector(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandSystemLogDeviceSelectorAction(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok || d.HasChange("comment") {
		t, err := expandSystemLogDeviceSelectorComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("devid"); ok || d.HasChange("devid") {
		t, err := expandSystemLogDeviceSelectorDevid(d, v, "devid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["devid"] = t
		}
	}

	if v, ok := d.GetOk("expire"); ok || d.HasChange("expire") {
		t, err := expandSystemLogDeviceSelectorExpire(d, v, "expire")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expire"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandSystemLogDeviceSelectorId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("srcip"); ok || d.HasChange("srcip") {
		t, err := expandSystemLogDeviceSelectorSrcip(d, v, "srcip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcip"] = t
		}
	}

	if v, ok := d.GetOk("srcip_mode"); ok || d.HasChange("srcip_mode") {
		t, err := expandSystemLogDeviceSelectorSrcipMode(d, v, "srcip_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcip-mode"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandSystemLogDeviceSelectorType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
