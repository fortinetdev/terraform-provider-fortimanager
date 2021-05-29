// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Retrieve a list of devices.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceDvmCmdReloadDevList() *schema.Resource {
	return &schema.Resource{
		Create: resourceDvmCmdReloadDevListUpdate,
		Read:   resourceDvmCmdReloadDevListRead,
		Update: resourceDvmCmdReloadDevListUpdate,
		Delete: resourceDvmCmdReloadDevListDelete,

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
			"flags": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"from": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reload_dev_member_list": &schema.Schema{
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
			"tag": &schema.Schema{
				Type:     schema.TypeString,
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

func resourceDvmCmdReloadDevListUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "", fmt.Errorf("")

	obj, err := getObjectDvmCmdReloadDevList(d)
	if err != nil {
		return fmt.Errorf("Error updating DvmCmdReloadDevList resource while getting object: %v", err)
	}

	_, err = c.UpdateDvmCmdReloadDevList(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating DvmCmdReloadDevList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("DvmCmdReloadDevList")

	return resourceDvmCmdReloadDevListRead(d, m)
}

func resourceDvmCmdReloadDevListDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")

	return nil
}

func resourceDvmCmdReloadDevListRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func flattenDvmCmdReloadDevListAdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmCmdReloadDevListFlags(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenDvmCmdReloadDevListFrom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "um",
			1: "fgfm",
			2: "apache",
			3: "dvm",
			4: "fwm",
			5: "xml",
			6: "json",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenDvmCmdReloadDevListReloadDevMemberList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenDvmCmdReloadDevListReloadDevMemberListName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "DvmCmdReloadDevList-ReloadDevMemberList-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenDvmCmdReloadDevListReloadDevMemberListVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "DvmCmdReloadDevList-ReloadDevMemberList-Vdom")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenDvmCmdReloadDevListReloadDevMemberListName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmCmdReloadDevListReloadDevMemberListVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmCmdReloadDevListTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectDvmCmdReloadDevList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fmgadom", flattenDvmCmdReloadDevListAdom(o["adom"], d, "fmgadom")); err != nil {
		if vv, ok := fortiAPIPatch(o["adom"], "DvmCmdReloadDevList-Adom"); ok {
			if err = d.Set("fmgadom", vv); err != nil {
				return fmt.Errorf("Error reading fmgadom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmgadom: %v", err)
		}
	}

	if err = d.Set("flags", flattenDvmCmdReloadDevListFlags(o["flags"], d, "flags")); err != nil {
		if vv, ok := fortiAPIPatch(o["flags"], "DvmCmdReloadDevList-Flags"); ok {
			if err = d.Set("flags", vv); err != nil {
				return fmt.Errorf("Error reading flags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading flags: %v", err)
		}
	}

	if err = d.Set("from", flattenDvmCmdReloadDevListFrom(o["from"], d, "from")); err != nil {
		if vv, ok := fortiAPIPatch(o["from"], "DvmCmdReloadDevList-From"); ok {
			if err = d.Set("from", vv); err != nil {
				return fmt.Errorf("Error reading from: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading from: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("reload_dev_member_list", flattenDvmCmdReloadDevListReloadDevMemberList(o["reload-dev-member-list"], d, "reload_dev_member_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["reload-dev-member-list"], "DvmCmdReloadDevList-ReloadDevMemberList"); ok {
				if err = d.Set("reload_dev_member_list", vv); err != nil {
					return fmt.Errorf("Error reading reload_dev_member_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading reload_dev_member_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("reload_dev_member_list"); ok {
			if err = d.Set("reload_dev_member_list", flattenDvmCmdReloadDevListReloadDevMemberList(o["reload-dev-member-list"], d, "reload_dev_member_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["reload-dev-member-list"], "DvmCmdReloadDevList-ReloadDevMemberList"); ok {
					if err = d.Set("reload_dev_member_list", vv); err != nil {
						return fmt.Errorf("Error reading reload_dev_member_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading reload_dev_member_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("tag", flattenDvmCmdReloadDevListTag(o["tag"], d, "tag")); err != nil {
		if vv, ok := fortiAPIPatch(o["tag"], "DvmCmdReloadDevList-Tag"); ok {
			if err = d.Set("tag", vv); err != nil {
				return fmt.Errorf("Error reading tag: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tag: %v", err)
		}
	}

	return nil
}

func flattenDvmCmdReloadDevListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandDvmCmdReloadDevListAdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmCmdReloadDevListFlags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandDvmCmdReloadDevListFrom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmCmdReloadDevListReloadDevMemberList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandDvmCmdReloadDevListReloadDevMemberListName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vdom"], _ = expandDvmCmdReloadDevListReloadDevMemberListVdom(d, i["vdom"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandDvmCmdReloadDevListReloadDevMemberListName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmCmdReloadDevListReloadDevMemberListVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmCmdReloadDevListTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectDvmCmdReloadDevList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fmgadom"); ok {
		t, err := expandDvmCmdReloadDevListAdom(d, v, "fmgadom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adom"] = t
		}
	}

	if v, ok := d.GetOk("flags"); ok {
		t, err := expandDvmCmdReloadDevListFlags(d, v, "flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flags"] = t
		}
	}

	if v, ok := d.GetOk("from"); ok {
		t, err := expandDvmCmdReloadDevListFrom(d, v, "from")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["from"] = t
		}
	}

	if v, ok := d.GetOk("reload_dev_member_list"); ok {
		t, err := expandDvmCmdReloadDevListReloadDevMemberList(d, v, "reload_dev_member_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reload-dev-member-list"] = t
		}
	}

	if v, ok := d.GetOk("tag"); ok {
		t, err := expandDvmCmdReloadDevListTag(d, v, "tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tag"] = t
		}
	}

	return &obj, nil
}
