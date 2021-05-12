// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Delete a list of devices.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceDvmCmdDelDevList() *schema.Resource {
	return &schema.Resource{
		Create: resourceDvmCmdDelDevListUpdate,
		Read:   resourceDvmCmdDelDevListRead,
		Update: resourceDvmCmdDelDevListUpdate,
		Delete: resourceDvmCmdDelDevListDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"force_recreate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"fmgadom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"del_dev_member_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vdom": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"flags": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
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

func resourceDvmCmdDelDevListUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "", fmt.Errorf("")

	obj, err := getObjectDvmCmdDelDevList(d)
	if err != nil {
		return fmt.Errorf("Error updating DvmCmdDelDevList resource while getting object: %v", err)
	}

	_, err = c.UpdateDvmCmdDelDevList(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating DvmCmdDelDevList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("DvmCmdDelDevList")

	return resourceDvmCmdDelDevListRead(d, m)
}

func resourceDvmCmdDelDevListDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")

	return nil
}

func resourceDvmCmdDelDevListRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func flattenDvmCmdDelDevListAdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmCmdDelDevListDelDevMemberList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenDvmCmdDelDevListDelDevMemberListName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "DvmCmdDelDevList-DelDevMemberList-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenDvmCmdDelDevListDelDevMemberListVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "DvmCmdDelDevList-DelDevMemberList-Vdom")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenDvmCmdDelDevListDelDevMemberListName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmCmdDelDevListDelDevMemberListVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmCmdDelDevListFlags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "none",
			1: "create_task",
			2: "nonblocking",
			4: "log_dev",
		}
		res := getEnumValbyBit(v, emap)
		return res
	}
	return v
}

func refreshObjectDvmCmdDelDevList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fmgadom", flattenDvmCmdDelDevListAdom(o["adom"], d, "fmgadom")); err != nil {
		if vv, ok := fortiAPIPatch(o["adom"], "DvmCmdDelDevList-Adom"); ok {
			if err = d.Set("fmgadom", vv); err != nil {
				return fmt.Errorf("Error reading fmgadom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmgadom: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("del_dev_member_list", flattenDvmCmdDelDevListDelDevMemberList(o["del-dev-member-list"], d, "del_dev_member_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["del-dev-member-list"], "DvmCmdDelDevList-DelDevMemberList"); ok {
				if err = d.Set("del_dev_member_list", vv); err != nil {
					return fmt.Errorf("Error reading del_dev_member_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading del_dev_member_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("del_dev_member_list"); ok {
			if err = d.Set("del_dev_member_list", flattenDvmCmdDelDevListDelDevMemberList(o["del-dev-member-list"], d, "del_dev_member_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["del-dev-member-list"], "DvmCmdDelDevList-DelDevMemberList"); ok {
					if err = d.Set("del_dev_member_list", vv); err != nil {
						return fmt.Errorf("Error reading del_dev_member_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading del_dev_member_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("flags", flattenDvmCmdDelDevListFlags(o["flags"], d, "flags")); err != nil {
		if vv, ok := fortiAPIPatch(o["flags"], "DvmCmdDelDevList-Flags"); ok {
			if err = d.Set("flags", vv); err != nil {
				return fmt.Errorf("Error reading flags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading flags: %v", err)
		}
	}

	return nil
}

func flattenDvmCmdDelDevListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandDvmCmdDelDevListAdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmCmdDelDevListDelDevMemberList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandDvmCmdDelDevListDelDevMemberListName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vdom"], _ = expandDvmCmdDelDevListDelDevMemberListVdom(d, i["vdom"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandDvmCmdDelDevListDelDevMemberListName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmCmdDelDevListDelDevMemberListVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmCmdDelDevListFlags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectDvmCmdDelDevList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fmgadom"); ok {
		t, err := expandDvmCmdDelDevListAdom(d, v, "fmgadom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adom"] = t
		}
	}

	if v, ok := d.GetOk("del_dev_member_list"); ok {
		t, err := expandDvmCmdDelDevListDelDevMemberList(d, v, "del_dev_member_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["del-dev-member-list"] = t
		}
	}

	if v, ok := d.GetOk("flags"); ok {
		t, err := expandDvmCmdDelDevListFlags(d, v, "flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flags"] = t
		}
	}

	return &obj, nil
}
