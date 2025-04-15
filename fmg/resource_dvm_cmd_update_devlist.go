// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Refresh FGFM connection and system information for a list of devices.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceDvmCmdUpdateDevList() *schema.Resource {
	return &schema.Resource{
		Create: resourceDvmCmdUpdateDevListUpdate,
		Read:   resourceDvmCmdUpdateDevListRead,
		Update: resourceDvmCmdUpdateDevListUpdate,
		Delete: resourceDvmCmdUpdateDevListDelete,

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
			"update_dev_member_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"vdom": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceDvmCmdUpdateDevListUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	obj, err := getObjectDvmCmdUpdateDevList(d)
	if err != nil {
		return fmt.Errorf("Error updating DvmCmdUpdateDevList resource while getting object: %v", err)
	}

	adomv := "adom/" + d.Get("fmgadom").(string)
	wsParams["adom"] = adomv

	_, err = c.UpdateDvmCmdUpdateDevList(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating DvmCmdUpdateDevList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("DvmCmdUpdateDevList")

	return resourceDvmCmdUpdateDevListRead(d, m)
}

func resourceDvmCmdUpdateDevListDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")

	return nil
}

func resourceDvmCmdUpdateDevListRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func flattenDvmCmdUpdateDevListAdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmCmdUpdateDevListFlags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenDvmCmdUpdateDevListUpdateDevMemberList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenDvmCmdUpdateDevListUpdateDevMemberListName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "DvmCmdUpdateDevList-UpdateDevMemberList-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenDvmCmdUpdateDevListUpdateDevMemberListVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "DvmCmdUpdateDevList-UpdateDevMemberList-Vdom")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenDvmCmdUpdateDevListUpdateDevMemberListName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmCmdUpdateDevListUpdateDevMemberListVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectDvmCmdUpdateDevList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("fmgadom", flattenDvmCmdUpdateDevListAdom(o["adom"], d, "fmgadom")); err != nil {
		if vv, ok := fortiAPIPatch(o["adom"], "DvmCmdUpdateDevList-Adom"); ok {
			if err = d.Set("fmgadom", vv); err != nil {
				return fmt.Errorf("Error reading fmgadom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmgadom: %v", err)
		}
	}

	if err = d.Set("flags", flattenDvmCmdUpdateDevListFlags(o["flags"], d, "flags")); err != nil {
		if vv, ok := fortiAPIPatch(o["flags"], "DvmCmdUpdateDevList-Flags"); ok {
			if err = d.Set("flags", vv); err != nil {
				return fmt.Errorf("Error reading flags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading flags: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("update_dev_member_list", flattenDvmCmdUpdateDevListUpdateDevMemberList(o["update-dev-member-list"], d, "update_dev_member_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["update-dev-member-list"], "DvmCmdUpdateDevList-UpdateDevMemberList"); ok {
				if err = d.Set("update_dev_member_list", vv); err != nil {
					return fmt.Errorf("Error reading update_dev_member_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading update_dev_member_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("update_dev_member_list"); ok {
			if err = d.Set("update_dev_member_list", flattenDvmCmdUpdateDevListUpdateDevMemberList(o["update-dev-member-list"], d, "update_dev_member_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["update-dev-member-list"], "DvmCmdUpdateDevList-UpdateDevMemberList"); ok {
					if err = d.Set("update_dev_member_list", vv); err != nil {
						return fmt.Errorf("Error reading update_dev_member_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading update_dev_member_list: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenDvmCmdUpdateDevListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandDvmCmdUpdateDevListAdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmCmdUpdateDevListFlags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandDvmCmdUpdateDevListUpdateDevMemberList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandDvmCmdUpdateDevListUpdateDevMemberListName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandDvmCmdUpdateDevListUpdateDevMemberListVdom(d, i["vdom"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandDvmCmdUpdateDevListUpdateDevMemberListName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmCmdUpdateDevListUpdateDevMemberListVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectDvmCmdUpdateDevList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fmgadom"); ok || d.HasChange("fmgadom") {
		t, err := expandDvmCmdUpdateDevListAdom(d, v, "fmgadom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adom"] = t
		}
	}

	if v, ok := d.GetOk("flags"); ok || d.HasChange("flags") {
		t, err := expandDvmCmdUpdateDevListFlags(d, v, "flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flags"] = t
		}
	}

	if v, ok := d.GetOk("update_dev_member_list"); ok || d.HasChange("update_dev_member_list") {
		t, err := expandDvmCmdUpdateDevListUpdateDevMemberList(d, v, "update_dev_member_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-dev-member-list"] = t
		}
	}

	return &obj, nil
}
