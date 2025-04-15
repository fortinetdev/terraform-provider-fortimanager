// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Add a device to the Device Manager database.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceDvmCmdAddDevice() *schema.Resource {
	return &schema.Resource{
		Create: resourceDvmCmdAddDeviceUpdate,
		Read:   resourceDvmCmdAddDeviceRead,
		Update: resourceDvmCmdAddDeviceUpdate,
		Delete: resourceDvmCmdAddDeviceDelete,

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
			"device": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"adm_pass": &schema.Schema{
							Type:      schema.TypeString,
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
						"metafields_map": &schema.Schema{
							Type:     schema.TypeMap,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
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
			"flags": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"groups": &schema.Schema{
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

func resourceDvmCmdAddDeviceUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	deviceVersion, err := c.GetDeviceVersion()
	if err != nil {
		log.Printf("Could not get device version: %v", err)
	}

	paradict := make(map[string]string)
	wsParams := make(map[string]string)

	obj, err := getObjectDvmCmdAddDevice(d, deviceVersion)
	if err != nil {
		return fmt.Errorf("Error updating DvmCmdAddDevice resource while getting object: %v", err)
	}

	adomv := "adom/" + d.Get("fmgadom").(string)
	wsParams["adom"] = adomv

	_, err = c.UpdateDvmCmdAddDevice(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating DvmCmdAddDevice resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("DvmCmdAddDevice")

	return resourceDvmCmdAddDeviceRead(d, m)
}

func resourceDvmCmdAddDeviceDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")

	return nil
}

func resourceDvmCmdAddDeviceRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func flattenDvmCmdAddDeviceAdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmCmdAddDeviceDevice(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex

	pre_append = pre + ".0." + "adm_usr"
	if _, ok := i["adm_usr"]; ok {
		result["adm_usr"] = flattenDvmCmdAddDeviceDeviceAdmUsr(i["adm_usr"], d, pre_append)
	}

	pre_append = pre + ".0." + "authorizationtemplate"
	if _, ok := i["authorization template"]; ok {
		result["authorizationtemplate"] = flattenDvmCmdAddDeviceDeviceAuthorizationTemplate(i["authorization template"], d, pre_append)
	}

	pre_append = pre + ".0." + "desc"
	if _, ok := i["desc"]; ok {
		result["desc"] = flattenDvmCmdAddDeviceDeviceDesc(i["desc"], d, pre_append)
	}

	pre_append = pre + ".0." + "deviceaction"
	if _, ok := i["device action"]; ok {
		result["deviceaction"] = flattenDvmCmdAddDeviceDeviceDeviceAction(i["device action"], d, pre_append)
	}

	pre_append = pre + ".0." + "deviceblueprint"
	if _, ok := i["device blueprint"]; ok {
		result["deviceblueprint"] = flattenDvmCmdAddDeviceDeviceDeviceBlueprint(i["device blueprint"], d, pre_append)
	}

	pre_append = pre + ".0." + "fazquota"
	if _, ok := i["faz.quota"]; ok {
		result["fazquota"] = flattenDvmCmdAddDeviceDeviceFazQuota(i["faz.quota"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip"
	if _, ok := i["ip"]; ok {
		result["ip"] = flattenDvmCmdAddDeviceDeviceIp(i["ip"], d, pre_append)
	}

	pre_append = pre + ".0." + "metafields"
	if _, ok := i["meta fields"]; ok {
		result["metafields"] = flattenDvmCmdAddDeviceDeviceMetaFields(i["meta fields"], d, pre_append)
	}

	pre_append = pre + ".0." + "metafields_map"
	if _, ok := i["meta fields"]; ok {
		result["metafields_map"] = flattenDvmCmdAddDeviceDeviceMetaFieldsMap(i["meta fields"], d, pre_append)
	}

	pre_append = pre + ".0." + "mgmt_mode"
	if _, ok := i["mgmt_mode"]; ok {
		result["mgmt_mode"] = flattenDvmCmdAddDeviceDeviceMgmtMode(i["mgmt_mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "mr"
	if _, ok := i["mr"]; ok {
		result["mr"] = flattenDvmCmdAddDeviceDeviceMr(i["mr"], d, pre_append)
	}

	pre_append = pre + ".0." + "name"
	if _, ok := i["name"]; ok {
		result["name"] = flattenDvmCmdAddDeviceDeviceName(i["name"], d, pre_append)
	}

	pre_append = pre + ".0." + "os_type"
	if _, ok := i["os_type"]; ok {
		result["os_type"] = flattenDvmCmdAddDeviceDeviceOsType(i["os_type"], d, pre_append)
	}

	pre_append = pre + ".0." + "os_ver"
	if _, ok := i["os_ver"]; ok {
		result["os_ver"] = flattenDvmCmdAddDeviceDeviceOsVer(i["os_ver"], d, pre_append)
	}

	pre_append = pre + ".0." + "patch"
	if _, ok := i["patch"]; ok {
		result["patch"] = flattenDvmCmdAddDeviceDevicePatch(i["patch"], d, pre_append)
	}

	pre_append = pre + ".0." + "platform_str"
	if _, ok := i["platform_str"]; ok {
		result["platform_str"] = flattenDvmCmdAddDeviceDevicePlatformStr(i["platform_str"], d, pre_append)
	}

	pre_append = pre + ".0." + "sn"
	if _, ok := i["sn"]; ok {
		result["sn"] = flattenDvmCmdAddDeviceDeviceSn(i["sn"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenDvmCmdAddDeviceDeviceAdmPass(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmCmdAddDeviceDeviceAdmUsr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmCmdAddDeviceDeviceAuthorizationTemplate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenDvmCmdAddDeviceDeviceDesc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmCmdAddDeviceDeviceDeviceAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmCmdAddDeviceDeviceDeviceBlueprint(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenDvmCmdAddDeviceDeviceFazQuota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmCmdAddDeviceDeviceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmCmdAddDeviceDeviceMetaFields(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmCmdAddDeviceDeviceMetaFieldsMap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmCmdAddDeviceDeviceMgmtMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmCmdAddDeviceDeviceMr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmCmdAddDeviceDeviceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmCmdAddDeviceDeviceOsType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmCmdAddDeviceDeviceOsVer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmCmdAddDeviceDevicePatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmCmdAddDeviceDevicePlatformStr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmCmdAddDeviceDeviceSn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmCmdAddDeviceFlags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenDvmCmdAddDeviceGroups(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenDvmCmdAddDeviceGroupsName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "DvmCmdAddDevice-Groups-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenDvmCmdAddDeviceGroupsVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "DvmCmdAddDevice-Groups-Vdom")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenDvmCmdAddDeviceGroupsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmCmdAddDeviceGroupsVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectDvmCmdAddDevice(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if err = d.Set("fmgadom", flattenDvmCmdAddDeviceAdom(o["adom"], d, "fmgadom")); err != nil {
		if vv, ok := fortiAPIPatch(o["adom"], "DvmCmdAddDevice-Adom"); ok {
			if err = d.Set("fmgadom", vv); err != nil {
				return fmt.Errorf("Error reading fmgadom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmgadom: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("device", flattenDvmCmdAddDeviceDevice(o["device"], d, "device")); err != nil {
			if vv, ok := fortiAPIPatch(o["device"], "DvmCmdAddDevice-Device"); ok {
				if err = d.Set("device", vv); err != nil {
					return fmt.Errorf("Error reading device: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading device: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("device"); ok {
			if err = d.Set("device", flattenDvmCmdAddDeviceDevice(o["device"], d, "device")); err != nil {
				if vv, ok := fortiAPIPatch(o["device"], "DvmCmdAddDevice-Device"); ok {
					if err = d.Set("device", vv); err != nil {
						return fmt.Errorf("Error reading device: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading device: %v", err)
				}
			}
		}
	}

	if err = d.Set("flags", flattenDvmCmdAddDeviceFlags(o["flags"], d, "flags")); err != nil {
		if vv, ok := fortiAPIPatch(o["flags"], "DvmCmdAddDevice-Flags"); ok {
			if err = d.Set("flags", vv); err != nil {
				return fmt.Errorf("Error reading flags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading flags: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("groups", flattenDvmCmdAddDeviceGroups(o["groups"], d, "groups")); err != nil {
			if vv, ok := fortiAPIPatch(o["groups"], "DvmCmdAddDevice-Groups"); ok {
				if err = d.Set("groups", vv); err != nil {
					return fmt.Errorf("Error reading groups: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading groups: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("groups"); ok {
			if err = d.Set("groups", flattenDvmCmdAddDeviceGroups(o["groups"], d, "groups")); err != nil {
				if vv, ok := fortiAPIPatch(o["groups"], "DvmCmdAddDevice-Groups"); ok {
					if err = d.Set("groups", vv); err != nil {
						return fmt.Errorf("Error reading groups: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading groups: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenDvmCmdAddDeviceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandDvmCmdAddDeviceAdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmCmdAddDeviceDevice(d *schema.ResourceData, v interface{}, pre string, deviceVersion string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "adm_pass"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["adm_pass"], _ = expandDvmCmdAddDeviceDeviceAdmPass(d, i["adm_pass"], pre_append)
	}
	pre_append = pre + ".0." + "adm_usr"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["adm_usr"], _ = expandDvmCmdAddDeviceDeviceAdmUsr(d, i["adm_usr"], pre_append)
	}
	pre_append = pre + ".0." + "authorizationtemplate"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["authorization template"], _ = expandDvmCmdAddDeviceDeviceAuthorizationTemplate(d, i["authorizationtemplate"], pre_append)
	}
	pre_append = pre + ".0." + "desc"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["desc"], _ = expandDvmCmdAddDeviceDeviceDesc(d, i["desc"], pre_append)
	}
	pre_append = pre + ".0." + "deviceaction"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["device action"], _ = expandDvmCmdAddDeviceDeviceDeviceAction(d, i["deviceaction"], pre_append)
	}
	pre_append = pre + ".0." + "deviceblueprint"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["device blueprint"], _ = expandDvmCmdAddDeviceDeviceDeviceBlueprint(d, i["deviceblueprint"], pre_append)
	}
	pre_append = pre + ".0." + "fazquota"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["faz.quota"], _ = expandDvmCmdAddDeviceDeviceFazQuota(d, i["fazquota"], pre_append)
	}
	pre_append = pre + ".0." + "ip"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["ip"], _ = expandDvmCmdAddDeviceDeviceIp(d, i["ip"], pre_append)
	}
	pre_append = pre + ".0." + "metafields"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		requiredVersion := map[string][]string{
			"operations": []string{"<"},
			"versions":   []string{"6.4.7"},
		}
		if versionMatch, err := checkVersionMatch(deviceVersion, requiredVersion); !versionMatch {
			err := fmt.Errorf("Argument 'metafields' %s.", err)
			return nil, err
		}
		result["meta fields"], _ = expandDvmCmdAddDeviceDeviceMetaFields(d, i["metafields"], pre_append)
	}
	pre_append = pre + ".0." + "metafields_map"
	if _, ok := d.GetOk(pre_append); ok {
		requiredVersion := map[string][]string{
			"operations": []string{">", "="},
			"versions":   []string{"6.4.7"},
		}
		if versionMatch, err := checkVersionMatch(deviceVersion, requiredVersion); !versionMatch {
			err := fmt.Errorf("Argument 'metafields_map' %s.", err)
			return nil, err
		}
		result["meta fields"], _ = expandDvmCmdAddDeviceDeviceMetaFieldsMap(d, i["metafields_map"], pre_append)
	}
	pre_append = pre + ".0." + "mgmt_mode"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["mgmt_mode"], _ = expandDvmCmdAddDeviceDeviceMgmtMode(d, i["mgmt_mode"], pre_append)
	}
	pre_append = pre + ".0." + "mr"
	if _, ok := d.GetOkExists(pre_append); ok || d.HasChange(pre_append) {
		result["mr"], _ = expandDvmCmdAddDeviceDeviceMr(d, i["mr"], pre_append)
	}
	pre_append = pre + ".0." + "name"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["name"], _ = expandDvmCmdAddDeviceDeviceName(d, i["name"], pre_append)
	}
	pre_append = pre + ".0." + "os_type"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["os_type"], _ = expandDvmCmdAddDeviceDeviceOsType(d, i["os_type"], pre_append)
	}
	pre_append = pre + ".0." + "os_ver"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["os_ver"], _ = expandDvmCmdAddDeviceDeviceOsVer(d, i["os_ver"], pre_append)
	}
	pre_append = pre + ".0." + "patch"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["patch"], _ = expandDvmCmdAddDeviceDevicePatch(d, i["patch"], pre_append)
	}
	pre_append = pre + ".0." + "platform_str"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["platform_str"], _ = expandDvmCmdAddDeviceDevicePlatformStr(d, i["platform_str"], pre_append)
	}
	pre_append = pre + ".0." + "sn"
	if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
		result["sn"], _ = expandDvmCmdAddDeviceDeviceSn(d, i["sn"], pre_append)
	}

	return result, nil
}

func expandDvmCmdAddDeviceDeviceAdmPass(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmCmdAddDeviceDeviceAdmUsr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmCmdAddDeviceDeviceAuthorizationTemplate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandDvmCmdAddDeviceDeviceDesc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmCmdAddDeviceDeviceDeviceAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmCmdAddDeviceDeviceDeviceBlueprint(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandDvmCmdAddDeviceDeviceFazQuota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmCmdAddDeviceDeviceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmCmdAddDeviceDeviceMetaFields(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmCmdAddDeviceDeviceMetaFieldsMap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmCmdAddDeviceDeviceMgmtMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmCmdAddDeviceDeviceMr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmCmdAddDeviceDeviceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmCmdAddDeviceDeviceOsType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmCmdAddDeviceDeviceOsVer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmCmdAddDeviceDevicePatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmCmdAddDeviceDevicePlatformStr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmCmdAddDeviceDeviceSn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmCmdAddDeviceFlags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandDvmCmdAddDeviceGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandDvmCmdAddDeviceGroupsName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandDvmCmdAddDeviceGroupsVdom(d, i["vdom"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandDvmCmdAddDeviceGroupsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmCmdAddDeviceGroupsVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectDvmCmdAddDevice(d *schema.ResourceData, deviceVersion string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fmgadom"); ok || d.HasChange("fmgadom") {
		t, err := expandDvmCmdAddDeviceAdom(d, v, "fmgadom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adom"] = t
		}
	}

	if v, ok := d.GetOk("device"); ok || d.HasChange("device") {
		t, err := expandDvmCmdAddDeviceDevice(d, v, "device", deviceVersion)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device"] = t
		}
	}

	if v, ok := d.GetOk("flags"); ok || d.HasChange("flags") {
		t, err := expandDvmCmdAddDeviceFlags(d, v, "flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flags"] = t
		}
	}

	if v, ok := d.GetOk("groups"); ok || d.HasChange("groups") {
		t, err := expandDvmCmdAddDeviceGroups(d, v, "groups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["groups"] = t
		}
	}

	return &obj, nil
}
