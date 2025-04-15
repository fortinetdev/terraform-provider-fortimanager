// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Add multiple devices to the Device Manager database.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceDvmCmdAddDevList() *schema.Resource {
	return &schema.Resource{
		Create: resourceDvmCmdAddDevListUpdate,
		Read:   resourceDvmCmdAddDevListRead,
		Update: resourceDvmCmdAddDevListUpdate,
		Delete: resourceDvmCmdAddDevListDelete,

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
			"add_dev_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"adm_pass": &schema.Schema{
							Type:      schema.TypeSet,
							Elem:      &schema.Schema{Type: schema.TypeString},
							Optional:  true,
							Sensitive: true,
							Computed:  true,
						},
						"adm_usr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"authorizationtemplate": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"desc": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"deviceaction": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"deviceblueprint": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"fazquota": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"metafields": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mgmt_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mr": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"os_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"os_ver": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"patch": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"platform_str": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"sn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
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
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceDvmCmdAddDevListUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	obj, err := getObjectDvmCmdAddDevList(d)
	if err != nil {
		return fmt.Errorf("Error updating DvmCmdAddDevList resource while getting object: %v", err)
	}

	adomv := "adom/" + d.Get("fmgadom").(string)
	wsParams["adom"] = adomv

	_, err = c.UpdateDvmCmdAddDevList(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating DvmCmdAddDevList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("DvmCmdAddDevList")

	return resourceDvmCmdAddDevListRead(d, m)
}

func resourceDvmCmdAddDevListDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")

	return nil
}

func resourceDvmCmdAddDevListRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func flattenDvmCmdAddDevListAddDevList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adm_usr"
		if _, ok := i["adm_usr"]; ok {
			v := flattenDvmCmdAddDevListAddDevListAdmUsr(i["adm_usr"], d, pre_append)
			tmp["adm_usr"] = fortiAPISubPartPatch(v, "DvmCmdAddDevList-AddDevList-AdmUsr")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authorizationtemplate"
		if _, ok := i["authorization template"]; ok {
			v := flattenDvmCmdAddDevListAddDevListAuthorizationTemplate(i["authorization template"], d, pre_append)
			tmp["authorizationtemplate"] = fortiAPISubPartPatch(v, "DvmCmdAddDevList-AddDevList-AuthorizationTemplate")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "desc"
		if _, ok := i["desc"]; ok {
			v := flattenDvmCmdAddDevListAddDevListDesc(i["desc"], d, pre_append)
			tmp["desc"] = fortiAPISubPartPatch(v, "DvmCmdAddDevList-AddDevList-Desc")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "deviceaction"
		if _, ok := i["device action"]; ok {
			v := flattenDvmCmdAddDevListAddDevListDeviceAction(i["device action"], d, pre_append)
			tmp["deviceaction"] = fortiAPISubPartPatch(v, "DvmCmdAddDevList-AddDevList-DeviceAction")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "deviceblueprint"
		if _, ok := i["device blueprint"]; ok {
			v := flattenDvmCmdAddDevListAddDevListDeviceBlueprint(i["device blueprint"], d, pre_append)
			tmp["deviceblueprint"] = fortiAPISubPartPatch(v, "DvmCmdAddDevList-AddDevList-DeviceBlueprint")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fazquota"
		if _, ok := i["faz.quota"]; ok {
			v := flattenDvmCmdAddDevListAddDevListFazQuota(i["faz.quota"], d, pre_append)
			tmp["fazquota"] = fortiAPISubPartPatch(v, "DvmCmdAddDevList-AddDevList-FazQuota")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			v := flattenDvmCmdAddDevListAddDevListIp(i["ip"], d, pre_append)
			tmp["ip"] = fortiAPISubPartPatch(v, "DvmCmdAddDevList-AddDevList-Ip")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metafields"
		if _, ok := i["meta fields"]; ok {
			v := flattenDvmCmdAddDevListAddDevListMetaFields(i["meta fields"], d, pre_append)
			tmp["metafields"] = fortiAPISubPartPatch(v, "DvmCmdAddDevList-AddDevList-MetaFields")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mgmt_mode"
		if _, ok := i["mgmt_mode"]; ok {
			v := flattenDvmCmdAddDevListAddDevListMgmtMode(i["mgmt_mode"], d, pre_append)
			tmp["mgmt_mode"] = fortiAPISubPartPatch(v, "DvmCmdAddDevList-AddDevList-MgmtMode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mr"
		if _, ok := i["mr"]; ok {
			v := flattenDvmCmdAddDevListAddDevListMr(i["mr"], d, pre_append)
			tmp["mr"] = fortiAPISubPartPatch(v, "DvmCmdAddDevList-AddDevList-Mr")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenDvmCmdAddDevListAddDevListName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "DvmCmdAddDevList-AddDevList-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "os_type"
		if _, ok := i["os_type"]; ok {
			v := flattenDvmCmdAddDevListAddDevListOsType(i["os_type"], d, pre_append)
			tmp["os_type"] = fortiAPISubPartPatch(v, "DvmCmdAddDevList-AddDevList-OsType")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "os_ver"
		if _, ok := i["os_ver"]; ok {
			v := flattenDvmCmdAddDevListAddDevListOsVer(i["os_ver"], d, pre_append)
			tmp["os_ver"] = fortiAPISubPartPatch(v, "DvmCmdAddDevList-AddDevList-OsVer")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "patch"
		if _, ok := i["patch"]; ok {
			v := flattenDvmCmdAddDevListAddDevListPatch(i["patch"], d, pre_append)
			tmp["patch"] = fortiAPISubPartPatch(v, "DvmCmdAddDevList-AddDevList-Patch")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "platform_str"
		if _, ok := i["platform_str"]; ok {
			v := flattenDvmCmdAddDevListAddDevListPlatformStr(i["platform_str"], d, pre_append)
			tmp["platform_str"] = fortiAPISubPartPatch(v, "DvmCmdAddDevList-AddDevList-PlatformStr")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sn"
		if _, ok := i["sn"]; ok {
			v := flattenDvmCmdAddDevListAddDevListSn(i["sn"], d, pre_append)
			tmp["sn"] = fortiAPISubPartPatch(v, "DvmCmdAddDevList-AddDevList-Sn")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenDvmCmdAddDevListAddDevListAdmUsr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmCmdAddDevListAddDevListAuthorizationTemplate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenDvmCmdAddDevListAddDevListDesc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmCmdAddDevListAddDevListDeviceAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmCmdAddDevListAddDevListDeviceBlueprint(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenDvmCmdAddDevListAddDevListFazQuota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmCmdAddDevListAddDevListIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmCmdAddDevListAddDevListMetaFields(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmCmdAddDevListAddDevListMgmtMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmCmdAddDevListAddDevListMr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmCmdAddDevListAddDevListName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmCmdAddDevListAddDevListOsType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmCmdAddDevListAddDevListOsVer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmCmdAddDevListAddDevListPatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmCmdAddDevListAddDevListPlatformStr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmCmdAddDevListAddDevListSn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmCmdAddDevListAdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmCmdAddDevListFlags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func refreshObjectDvmCmdAddDevList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("add_dev_list", flattenDvmCmdAddDevListAddDevList(o["add-dev-list"], d, "add_dev_list")); err != nil {
			if vv, ok := fortiAPIPatch(o["add-dev-list"], "DvmCmdAddDevList-AddDevList"); ok {
				if err = d.Set("add_dev_list", vv); err != nil {
					return fmt.Errorf("Error reading add_dev_list: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading add_dev_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("add_dev_list"); ok {
			if err = d.Set("add_dev_list", flattenDvmCmdAddDevListAddDevList(o["add-dev-list"], d, "add_dev_list")); err != nil {
				if vv, ok := fortiAPIPatch(o["add-dev-list"], "DvmCmdAddDevList-AddDevList"); ok {
					if err = d.Set("add_dev_list", vv); err != nil {
						return fmt.Errorf("Error reading add_dev_list: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading add_dev_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("fmgadom", flattenDvmCmdAddDevListAdom(o["adom"], d, "fmgadom")); err != nil {
		if vv, ok := fortiAPIPatch(o["adom"], "DvmCmdAddDevList-Adom"); ok {
			if err = d.Set("fmgadom", vv); err != nil {
				return fmt.Errorf("Error reading fmgadom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmgadom: %v", err)
		}
	}

	if err = d.Set("flags", flattenDvmCmdAddDevListFlags(o["flags"], d, "flags")); err != nil {
		if vv, ok := fortiAPIPatch(o["flags"], "DvmCmdAddDevList-Flags"); ok {
			if err = d.Set("flags", vv); err != nil {
				return fmt.Errorf("Error reading flags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading flags: %v", err)
		}
	}

	return nil
}

func flattenDvmCmdAddDevListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandDvmCmdAddDevListAddDevList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adm_pass"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["adm_pass"], _ = expandDvmCmdAddDevListAddDevListAdmPass(d, i["adm_pass"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adm_usr"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["adm_usr"], _ = expandDvmCmdAddDevListAddDevListAdmUsr(d, i["adm_usr"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authorizationtemplate"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["authorization template"], _ = expandDvmCmdAddDevListAddDevListAuthorizationTemplate(d, i["authorizationtemplate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "desc"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["desc"], _ = expandDvmCmdAddDevListAddDevListDesc(d, i["desc"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "deviceaction"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["device action"], _ = expandDvmCmdAddDevListAddDevListDeviceAction(d, i["deviceaction"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "deviceblueprint"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["device blueprint"], _ = expandDvmCmdAddDevListAddDevListDeviceBlueprint(d, i["deviceblueprint"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fazquota"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["faz.quota"], _ = expandDvmCmdAddDevListAddDevListFazQuota(d, i["fazquota"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["ip"], _ = expandDvmCmdAddDevListAddDevListIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metafields"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["meta fields"], _ = expandDvmCmdAddDevListAddDevListMetaFields(d, i["metafields"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mgmt_mode"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mgmt_mode"], _ = expandDvmCmdAddDevListAddDevListMgmtMode(d, i["mgmt_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mr"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["mr"], _ = expandDvmCmdAddDevListAddDevListMr(d, i["mr"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandDvmCmdAddDevListAddDevListName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "os_type"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["os_type"], _ = expandDvmCmdAddDevListAddDevListOsType(d, i["os_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "os_ver"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["os_ver"], _ = expandDvmCmdAddDevListAddDevListOsVer(d, i["os_ver"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "patch"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["patch"], _ = expandDvmCmdAddDevListAddDevListPatch(d, i["patch"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "platform_str"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["platform_str"], _ = expandDvmCmdAddDevListAddDevListPlatformStr(d, i["platform_str"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sn"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["sn"], _ = expandDvmCmdAddDevListAddDevListSn(d, i["sn"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandDvmCmdAddDevListAddDevListAdmPass(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandDvmCmdAddDevListAddDevListAdmUsr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmCmdAddDevListAddDevListAuthorizationTemplate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandDvmCmdAddDevListAddDevListDesc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmCmdAddDevListAddDevListDeviceAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmCmdAddDevListAddDevListDeviceBlueprint(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandDvmCmdAddDevListAddDevListFazQuota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmCmdAddDevListAddDevListIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmCmdAddDevListAddDevListMetaFields(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmCmdAddDevListAddDevListMgmtMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmCmdAddDevListAddDevListMr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmCmdAddDevListAddDevListName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmCmdAddDevListAddDevListOsType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmCmdAddDevListAddDevListOsVer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmCmdAddDevListAddDevListPatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmCmdAddDevListAddDevListPlatformStr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmCmdAddDevListAddDevListSn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmCmdAddDevListAdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmCmdAddDevListFlags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func getObjectDvmCmdAddDevList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("add_dev_list"); ok || d.HasChange("add_dev_list") {
		t, err := expandDvmCmdAddDevListAddDevList(d, v, "add_dev_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["add-dev-list"] = t
		}
	}

	if v, ok := d.GetOk("fmgadom"); ok || d.HasChange("fmgadom") {
		t, err := expandDvmCmdAddDevListAdom(d, v, "fmgadom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adom"] = t
		}
	}

	if v, ok := d.GetOk("flags"); ok || d.HasChange("flags") {
		t, err := expandDvmCmdAddDevListFlags(d, v, "flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flags"] = t
		}
	}

	return &obj, nil
}
