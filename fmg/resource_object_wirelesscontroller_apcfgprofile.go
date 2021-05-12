// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure AP local configuration profiles.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceObjectWirelessControllerApcfgProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectWirelessControllerApcfgProfileCreate,
		Read:   resourceObjectWirelessControllerApcfgProfileRead,
		Update: resourceObjectWirelessControllerApcfgProfileUpdate,
		Delete: resourceObjectWirelessControllerApcfgProfileDelete,

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
			"ac_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ac_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ac_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ac_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"command_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
							Computed: true,
						},
					},
				},
			},
			"comment": &schema.Schema{
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
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceObjectWirelessControllerApcfgProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectWirelessControllerApcfgProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerApcfgProfile resource while getting object: %v", err)
	}

	_, err = c.CreateObjectWirelessControllerApcfgProfile(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating ObjectWirelessControllerApcfgProfile resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerApcfgProfileRead(d, m)
}

func resourceObjectWirelessControllerApcfgProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectWirelessControllerApcfgProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerApcfgProfile resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectWirelessControllerApcfgProfile(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectWirelessControllerApcfgProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectWirelessControllerApcfgProfileRead(d, m)
}

func resourceObjectWirelessControllerApcfgProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectWirelessControllerApcfgProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectWirelessControllerApcfgProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectWirelessControllerApcfgProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectWirelessControllerApcfgProfile(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerApcfgProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectWirelessControllerApcfgProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectWirelessControllerApcfgProfile resource from API: %v", err)
	}
	return nil
}

func flattenObjectWirelessControllerApcfgProfileAcIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerApcfgProfileAcPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerApcfgProfileAcTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerApcfgProfileAcType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "default",
			1: "specify",
			2: "apcfg",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectWirelessControllerApcfgProfileCommandList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			v := flattenObjectWirelessControllerApcfgProfileCommandListId(i["id"], d, pre_append)
			tmp["id"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerApcfgProfile-CommandList-Id")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenObjectWirelessControllerApcfgProfileCommandListName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerApcfgProfile-CommandList-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passwd_value"
		if _, ok := i["passwd-value"]; ok {
			v := flattenObjectWirelessControllerApcfgProfileCommandListPasswdValue(i["passwd-value"], d, pre_append)
			tmp["passwd_value"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerApcfgProfile-CommandList-PasswdValue")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			v := flattenObjectWirelessControllerApcfgProfileCommandListType(i["type"], d, pre_append)
			tmp["type"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerApcfgProfile-CommandList-Type")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			v := flattenObjectWirelessControllerApcfgProfileCommandListValue(i["value"], d, pre_append)
			tmp["value"] = fortiAPISubPartPatch(v, "ObjectWirelessControllerApcfgProfile-CommandList-Value")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenObjectWirelessControllerApcfgProfileCommandListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerApcfgProfileCommandListName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerApcfgProfileCommandListPasswdValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenObjectWirelessControllerApcfgProfileCommandListType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "non-password",
			1: "password",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenObjectWirelessControllerApcfgProfileCommandListValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerApcfgProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectWirelessControllerApcfgProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectWirelessControllerApcfgProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ac_ip", flattenObjectWirelessControllerApcfgProfileAcIp(o["ac-ip"], d, "ac_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ac-ip"], "ObjectWirelessControllerApcfgProfile-AcIp"); ok {
			if err = d.Set("ac_ip", vv); err != nil {
				return fmt.Errorf("Error reading ac_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ac_ip: %v", err)
		}
	}

	if err = d.Set("ac_port", flattenObjectWirelessControllerApcfgProfileAcPort(o["ac-port"], d, "ac_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["ac-port"], "ObjectWirelessControllerApcfgProfile-AcPort"); ok {
			if err = d.Set("ac_port", vv); err != nil {
				return fmt.Errorf("Error reading ac_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ac_port: %v", err)
		}
	}

	if err = d.Set("ac_timer", flattenObjectWirelessControllerApcfgProfileAcTimer(o["ac-timer"], d, "ac_timer")); err != nil {
		if vv, ok := fortiAPIPatch(o["ac-timer"], "ObjectWirelessControllerApcfgProfile-AcTimer"); ok {
			if err = d.Set("ac_timer", vv); err != nil {
				return fmt.Errorf("Error reading ac_timer: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ac_timer: %v", err)
		}
	}

	if err = d.Set("ac_type", flattenObjectWirelessControllerApcfgProfileAcType(o["ac-type"], d, "ac_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["ac-type"], "ObjectWirelessControllerApcfgProfile-AcType"); ok {
			if err = d.Set("ac_type", vv); err != nil {
				return fmt.Errorf("Error reading ac_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ac_type: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("command_list", flattenObjectWirelessControllerApcfgProfileCommandList(o["command-list"], d, "command_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["command-list"], "ObjectWirelessControllerApcfgProfile-CommandList"); ok {
				if err = d.Set("command_list", vv); err != nil {
					return fmt.Errorf("Error reading command_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading command_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("command_list"); ok {
			if err = d.Set("command_list", flattenObjectWirelessControllerApcfgProfileCommandList(o["command-list"], d, "command_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["command-list"], "ObjectWirelessControllerApcfgProfile-CommandList"); ok {
					if err = d.Set("command_list", vv); err != nil {
						return fmt.Errorf("Error reading command_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading command_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("comment", flattenObjectWirelessControllerApcfgProfileComment(o["comment"], d, "comment")); err != nil {
		if vv, ok := fortiAPIPatch(o["comment"], "ObjectWirelessControllerApcfgProfile-Comment"); ok {
			if err = d.Set("comment", vv); err != nil {
				return fmt.Errorf("Error reading comment: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectWirelessControllerApcfgProfileName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectWirelessControllerApcfgProfile-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenObjectWirelessControllerApcfgProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectWirelessControllerApcfgProfileAcIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerApcfgProfileAcPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerApcfgProfileAcTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerApcfgProfileAcType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerApcfgProfileCommandList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandObjectWirelessControllerApcfgProfileCommandListId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandObjectWirelessControllerApcfgProfileCommandListName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passwd_value"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["passwd-value"], _ = expandObjectWirelessControllerApcfgProfileCommandListPasswdValue(d, i["passwd_value"], pre_append)
		} else {
			tmp["passwd-value"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandObjectWirelessControllerApcfgProfileCommandListType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["value"], _ = expandObjectWirelessControllerApcfgProfileCommandListValue(d, i["value"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandObjectWirelessControllerApcfgProfileCommandListId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerApcfgProfileCommandListName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerApcfgProfileCommandListPasswdValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandObjectWirelessControllerApcfgProfileCommandListType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerApcfgProfileCommandListValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerApcfgProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectWirelessControllerApcfgProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectWirelessControllerApcfgProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ac_ip"); ok {
		t, err := expandObjectWirelessControllerApcfgProfileAcIp(d, v, "ac_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ac-ip"] = t
		}
	}

	if v, ok := d.GetOk("ac_port"); ok {
		t, err := expandObjectWirelessControllerApcfgProfileAcPort(d, v, "ac_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ac-port"] = t
		}
	}

	if v, ok := d.GetOk("ac_timer"); ok {
		t, err := expandObjectWirelessControllerApcfgProfileAcTimer(d, v, "ac_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ac-timer"] = t
		}
	}

	if v, ok := d.GetOk("ac_type"); ok {
		t, err := expandObjectWirelessControllerApcfgProfileAcType(d, v, "ac_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ac-type"] = t
		}
	}

	if v, ok := d.GetOk("command_list"); ok {
		t, err := expandObjectWirelessControllerApcfgProfileCommandList(d, v, "command_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["command-list"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandObjectWirelessControllerApcfgProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandObjectWirelessControllerApcfgProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
