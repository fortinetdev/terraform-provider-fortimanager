// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: AP local configuration command list.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectWirelessControllerApcfgProfileCommandList() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerApcfgProfileCommandListCreate,
		Read:   resourceObjectWirelessControllerApcfgProfileCommandListRead,
		Update: resourceObjectWirelessControllerApcfgProfileCommandListUpdate,
		Delete: resourceObjectWirelessControllerApcfgProfileCommandListDelete,

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
			"apcfg_profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"passwd_value": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectWirelessControllerApcfgProfileCommandListCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	apcfg_profile := d.Get("apcfg_profile").(string)
	paradict["apcfg_profile"] = apcfg_profile

	obj, err := getObjectObjectWirelessControllerApcfgProfileCommandList(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerApcfgProfileCommandList resource while getting object: %v", err)
	}

	_, err = c.CreateObjectWirelessControllerApcfgProfileCommandList(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerApcfgProfileCommandList resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectWirelessControllerApcfgProfileCommandListRead(d, m)
}

func resourceObjectWirelessControllerApcfgProfileCommandListUpdate(d *schema.ResourceData, m interface{}) error {
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

	apcfg_profile := d.Get("apcfg_profile").(string)
	paradict["apcfg_profile"] = apcfg_profile

	obj, err := getObjectObjectWirelessControllerApcfgProfileCommandList(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerApcfgProfileCommandList resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWirelessControllerApcfgProfileCommandList(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerApcfgProfileCommandList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceObjectWirelessControllerApcfgProfileCommandListRead(d, m)
}

func resourceObjectWirelessControllerApcfgProfileCommandListDelete(d *schema.ResourceData, m interface{}) error {
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

	apcfg_profile := d.Get("apcfg_profile").(string)
	paradict["apcfg_profile"] = apcfg_profile

	err = c.DeleteObjectWirelessControllerApcfgProfileCommandList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWirelessControllerApcfgProfileCommandList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerApcfgProfileCommandListRead(d *schema.ResourceData, m interface{}) error {
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

	apcfg_profile := d.Get("apcfg_profile").(string)
	if apcfg_profile == "" {
		apcfg_profile = importOptionChecking(m.(*FortiClient).Cfg, "apcfg_profile")
		if apcfg_profile == "" {
			return fmt.Errorf("Parameter apcfg_profile is missing")
		}
		if err = d.Set("apcfg_profile", apcfg_profile); err != nil {
			return fmt.Errorf("Error set params apcfg_profile: %v", err)
		}
	}
	paradict["apcfg_profile"] = apcfg_profile

	o, err := c.ReadObjectWirelessControllerApcfgProfileCommandList(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerApcfgProfileCommandList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWirelessControllerApcfgProfileCommandList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerApcfgProfileCommandList resource from API: %v", err)
	}
	return nil
}

func flattenObjectWirelessControllerApcfgProfileCommandListId2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerApcfgProfileCommandListName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerApcfgProfileCommandListPasswdValue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerApcfgProfileCommandListType2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerApcfgProfileCommandListValue2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWirelessControllerApcfgProfileCommandList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("fosid", flattenObjectWirelessControllerApcfgProfileCommandListId2edl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "ObjectWirelessControllerApcfgProfileCommandList-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectWirelessControllerApcfgProfileCommandListName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectWirelessControllerApcfgProfileCommandList-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("passwd_value", flattenObjectWirelessControllerApcfgProfileCommandListPasswdValue2edl(o["passwd-value"], d, "passwd_value")); err != nil {
		if vv, ok := fortiAPIPatch(o["passwd-value"], "ObjectWirelessControllerApcfgProfileCommandList-PasswdValue"); ok {
			if err = d.Set("passwd_value", vv); err != nil {
				return fmt.Errorf("Error reading passwd_value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading passwd_value: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectWirelessControllerApcfgProfileCommandListType2edl(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectWirelessControllerApcfgProfileCommandList-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("value", flattenObjectWirelessControllerApcfgProfileCommandListValue2edl(o["value"], d, "value")); err != nil {
		if vv, ok := fortiAPIPatch(o["value"], "ObjectWirelessControllerApcfgProfileCommandList-Value"); ok {
			if err = d.Set("value", vv); err != nil {
				return fmt.Errorf("Error reading value: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading value: %v", err)
		}
	}

	return nil
}

func flattenObjectWirelessControllerApcfgProfileCommandListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerApcfgProfileCommandListId2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerApcfgProfileCommandListName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerApcfgProfileCommandListPasswdValue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerApcfgProfileCommandListType2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerApcfgProfileCommandListValue2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWirelessControllerApcfgProfileCommandList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandObjectWirelessControllerApcfgProfileCommandListId2edl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectWirelessControllerApcfgProfileCommandListName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("passwd_value"); ok || d.HasChange("passwd_value") {
		t, err := expandObjectWirelessControllerApcfgProfileCommandListPasswdValue2edl(d, v, "passwd_value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passwd-value"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectWirelessControllerApcfgProfileCommandListType2edl(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("value"); ok || d.HasChange("value") {
		t, err := expandObjectWirelessControllerApcfgProfileCommandListValue2edl(d, v, "value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["value"] = t
		}
	}

	return &obj, nil
}
