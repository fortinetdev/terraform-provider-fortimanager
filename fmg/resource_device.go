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

func resourceDevice() *schema.Resource {
	return &schema.Resource{
		Create: resourceDeviceCreate,
		Read:   resourceDeviceRead,
		Update: resourceDeviceUpdate,
		Delete: resourceDeviceDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			// FMG argument
			"adom": &schema.Schema{
				Type:     schema.TypeString,
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
			// Device arguments
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
				Computed: true,
			},
			"deviceaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"deviceblueprint": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fazquota": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"metafields": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"mgmt_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mr": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"os_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"os_ver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"patch": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"platform_str": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			// Device attributes
			"app_ver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"av_ver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"beta": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"branch_pt": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"build": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"checksum": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cluster_worker": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"conf_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"conn_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"conn_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"db_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dev_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fap_cnt": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fazfull_act": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fazperm": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fazused": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fex_cnt": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"first_tunnel_up": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"flags": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"foslic_cpu": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"foslic_dr_site": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"foslic_inst_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"foslic_last_sync": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"foslic_ram": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"foslic_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"foslic_utm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fsw_cnt": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"havsn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ha_group_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ha_group_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ha_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ha_upgrade_mode": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"hdisk_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hw_generation": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"hw_rev_major": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"hw_rev_minor": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"hyperscale": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ips_ext": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ips_ver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"last_checked": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"last_resync": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"latitude": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lic_flags": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"lic_region": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"location_from": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logdisk_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"longitude": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maxvdom": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mgmt_if": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mgmt_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mgt_vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"module_sn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nsxt_service_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"prefer_img_ver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"prio": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"private_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"private_key_status": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"psk": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"relver_info": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"role": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sov_sase_license": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tunnel_sn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"comments": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"metafields": &schema.Schema{
							Type:     schema.TypeMap,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"opmode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rtm_prof_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vdom_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vpn_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},

			"version": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vm_cpu": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vm_cpu_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vm_lic_expire": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vm_lic_overdue_since": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vm_mem": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vm_mem_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vm_payg_status": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vm_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceDeviceCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client

	wsParams := make(map[string]string)
	adomv := getAdom(d, m)
	if adomv == "" {
		return fmt.Errorf("Adom is missing, ")
	}

	obj, err := getObjectDevice(d, "create", adomv)
	if err != nil {
		return fmt.Errorf("Error creating Device resource while getting object: %v", err)
	}

	wsParams["adom"] = "adom/" + adomv

	o, err := c.CreateDevice(obj, nil, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating Device resource: %v", err)
	}

	taskID, err := getTaskID(o)
	if err != nil {
		return fmt.Errorf("Error get task ID for create device: %v", err)
	}

	err = c.WaitTask(taskID)
	if err != nil {
		return fmt.Errorf("Error wait task finish for create device: %v", err)
	}

	d.SetId(d.Get("name").(string))

	return resourceDeviceRead(d, m)
}

func resourceDeviceUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client

	wsParams := make(map[string]string)
	adomv := getAdom(d, m)
	if adomv == "" {
		return fmt.Errorf("Adom is missing, ")
	}

	obj, err := getObjectDevice(d, "update", adomv)
	if err != nil {
		return fmt.Errorf("Error updating Device resource while getting object: %v", err)
	}

	wsParams["adom"] = "adom/" + adomv

	o, err := c.UpdateDevice(obj, mkey, nil, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating Device resource: %v", err)
	}

	taskID, err := getTaskID(o)
	if err != nil {
		return fmt.Errorf("Error get task ID for update device: %v", err)
	}

	err = c.WaitTask(taskID)
	if err != nil {
		return fmt.Errorf("Error wait task finish for update device: %v", err)
	}

	d.SetId(d.Get("name").(string))

	return resourceDeviceRead(d, m)
}

func resourceDeviceDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	wsParams := make(map[string]string)
	adomv := getAdom(d, m)
	if adomv == "" {
		return fmt.Errorf("Adom is missing, ")
	}

	obj, err := getObjectDevice(d, "delete", adomv)
	if err != nil {
		return fmt.Errorf("Error delete Device resource while getting object: %v", err)
	}

	wsParams["adom"] = "adom/" + adomv

	o, err := c.DeleteDevice(obj, mkey, nil, wsParams)
	if err != nil {
		return fmt.Errorf("Error delete Device resource: %v", err)
	}

	taskID, err := getTaskID(o)
	if err != nil {
		return fmt.Errorf("Error get task ID for delete device: %v", err)
	}

	err = c.WaitTask(taskID)
	if err != nil {
		return fmt.Errorf("Error wait task finish for delete device: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceDeviceRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadDevice(mkey, nil)
	if err != nil {
		d.SetId("")
		return fmt.Errorf("Error reading Device resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshDevice(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectDevice resource from API: %v", err)
	}
	return nil
}

func getAdom(d *schema.ResourceData, m interface{}) (adom string) {
	if d.Get("adom") != nil {
		adom = d.Get("adom").(string)
	}
	if adom == "" {
		cfg := m.(*FortiClient).Cfg
		adom = cfg.Adom
	}
	return
}

func flattenDeviceAdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDeviceDevice(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {

	///
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex

	pre_append = pre + ".0." + "sn"
	if _, ok := i["sn"]; ok {
		result["sn"] = flattenDeviceDeviceSn(i["sn"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenDeviceDeviceAdmPass(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDeviceDeviceAdmUsr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDeviceDeviceAuthorizationTemplate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2str(v, d.Get(pre))
}

func flattenDeviceDeviceDesc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDeviceDeviceDeviceAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDeviceDeviceDeviceBlueprint(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return conv2str(v)
}

func flattenDeviceDeviceFazQuota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDeviceDeviceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDeviceDeviceMetaFields(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDeviceDeviceMgmtMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDeviceDeviceMr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDeviceDeviceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDeviceDeviceOsType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDeviceDeviceOsVer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDeviceDevicePatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDeviceDevicePlatformStr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDeviceDeviceSn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDeviceFlags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenDeviceGroups(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenDeviceGroupsName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "Device-Groups-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			v := flattenDeviceGroupsVdom(i["vdom"], d, pre_append)
			tmp["vdom"] = fortiAPISubPartPatch(v, "Device-Groups-Vdom")
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenDeviceVdom(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		if _, ok := i["comments"]; ok {
			tmp["comments"] = i["comments"]
		}

		if _, ok := i["meta fields"]; ok {
			tmp["metafields"] = i["meta fields"]
		}

		if _, ok := i["name"]; ok {
			tmp["name"] = i["name"]
		}

		if _, ok := i["opmode"]; ok {
			tmp["opmode"] = i["opmode"]
		}

		if _, ok := i["rtm_prof_id"]; ok {
			tmp["rtm_prof_id"] = i["rtm_prof_id"]
		}

		if _, ok := i["status"]; ok {
			tmp["status"] = i["status"]
		}

		if _, ok := i["vdom_type"]; ok {
			tmp["vdom_type"] = i["vdom_type"]
		}

		if _, ok := i["vpn_id"]; ok {
			tmp["vpn_id"] = i["vpn_id"]
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result
}

func flattenDeviceGroupsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDeviceGroupsVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshDevice(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

	if isImportTable() {
		if err = d.Set("groups", flattenDeviceGroups(o["groups"], d, "groups")); err != nil {
			if vv, ok := fortiAPIPatch(o["groups"], "Device-Groups"); ok {
				if err = d.Set("groups", vv); err != nil {
					return fmt.Errorf("Error reading groups: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading groups: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("groups"); ok {
			if err = d.Set("groups", flattenDeviceGroups(o["groups"], d, "groups")); err != nil {
				if vv, ok := fortiAPIPatch(o["groups"], "Device-Groups"); ok {
					if err = d.Set("groups", vv); err != nil {
						return fmt.Errorf("Error reading groups: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading groups: %v", err)
				}
			}
		}
	}

	// Device arguments
	if err = d.Set("adm_usr", o["adm_usr"]); err != nil {
		return fmt.Errorf("Error reading adm_usr: %v", err)
	}

	if err = d.Set("authorizationtemplate", o["authorization template"]); err != nil {
		return fmt.Errorf("Error reading authorizationtemplate: %v", err)
	}

	if err = d.Set("desc", o["desc"]); err != nil {
		return fmt.Errorf("Error reading desc: %v", err)
	}

	if err = d.Set("deviceaction", o["device action"]); err != nil {
		return fmt.Errorf("Error reading deviceaction: %v", err)
	}

	if err = d.Set("deviceblueprint", o["device blueprint"]); err != nil {
		return fmt.Errorf("Error reading deviceblueprint: %v", err)
	}

	if err = d.Set("fazquota", o["faz.quota"]); err != nil {
		return fmt.Errorf("Error reading fazquota: %v", err)
	}

	if err = d.Set("ip", o["ip"]); err != nil {
		return fmt.Errorf("Error reading ip: %v", err)
	}

	if err = d.Set("metafields", flattenDeviceDeviceMetaFields(o["meta fields"], d, "metafields")); err != nil {
		return fmt.Errorf("Error reading metafields: %v", err)
	}

	if err = d.Set("mgmt_mode", o["mgmt_mode"]); err != nil {
		return fmt.Errorf("Error reading mgmt_mode: %v", err)
	}

	if err = d.Set("mr", o["mr"]); err != nil {
		return fmt.Errorf("Error reading mr: %v", err)
	}

	if err = d.Set("name", o["name"]); err != nil {
		return fmt.Errorf("Error reading name: %v", err)
	}

	if err = d.Set("os_type", o["os_type"]); err != nil {
		return fmt.Errorf("Error reading os_type: %v", err)
	}

	if err = d.Set("os_ver", o["os_ver"]); err != nil {
		return fmt.Errorf("Error reading os_ver: %v", err)
	}

	if err = d.Set("patch", o["patch"]); err != nil {
		return fmt.Errorf("Error reading patch: %v", err)
	}

	if err = d.Set("platform_str", o["platform_str"]); err != nil {
		return fmt.Errorf("Error reading platform_str: %v", err)
	}

	if err = d.Set("sn", o["sn"]); err != nil {
		return fmt.Errorf("Error reading sn: %v", err)
	}

	// Device attribute
	if err = d.Set("app_ver", o["app_ver"]); err != nil {
		return fmt.Errorf("Error reading app_ver: %v", err)
	}
	if err = d.Set("av_ver", o["av_ver"]); err != nil {
		return fmt.Errorf("Error reading av_ver: %v", err)
	}
	if err = d.Set("beta", o["beta"]); err != nil {
		return fmt.Errorf("Error reading beta: %v", err)
	}
	if err = d.Set("branch_pt", o["branch_pt"]); err != nil {
		return fmt.Errorf("Error reading branch_pt: %v", err)
	}
	if err = d.Set("build", o["build"]); err != nil {
		return fmt.Errorf("Error reading build: %v", err)
	}
	if err = d.Set("checksum", o["checksum"]); err != nil {
		return fmt.Errorf("Error reading checksum: %v", err)
	}
	if err = d.Set("cluster_worker", o["cluster_worker"]); err != nil {
		return fmt.Errorf("Error reading cluster_worker: %v", err)
	}
	if err = d.Set("conf_status", o["conf_status"]); err != nil {
		return fmt.Errorf("Error reading conf_status: %v", err)
	}
	if err = d.Set("conn_mode", o["conn_mode"]); err != nil {
		return fmt.Errorf("Error reading conn_mode: %v", err)
	}
	if err = d.Set("conn_status", o["conn_status"]); err != nil {
		return fmt.Errorf("Error reading conn_status: %v", err)
	}
	if err = d.Set("db_status", o["db_status"]); err != nil {
		return fmt.Errorf("Error reading db_status: %v", err)
	}
	if err = d.Set("dev_status", o["dev_status"]); err != nil {
		return fmt.Errorf("Error reading dev_status: %v", err)
	}
	if err = d.Set("eip", o["eip"]); err != nil {
		return fmt.Errorf("Error reading eip: %v", err)
	}
	if err = d.Set("fap_cnt", o["fap_cnt"]); err != nil {
		return fmt.Errorf("Error reading fap_cnt: %v", err)
	}
	if err = d.Set("fazfull_act", o["faz.full_act"]); err != nil {
		return fmt.Errorf("Error reading fazfull_act: %v", err)
	}
	if err = d.Set("fazperm", o["faz.perm"]); err != nil {
		return fmt.Errorf("Error reading fazperm: %v", err)
	}
	if err = d.Set("fazused", o["faz.used"]); err != nil {
		return fmt.Errorf("Error reading fazused: %v", err)
	}
	if err = d.Set("fex_cnt", o["fex_cnt"]); err != nil {
		return fmt.Errorf("Error reading fex_cnt: %v", err)
	}
	if err = d.Set("first_tunnel_up", o["first_tunnel_up"]); err != nil {
		return fmt.Errorf("Error reading first_tunnel_up: %v", err)
	}
	if err = d.Set("flags", o["flags"]); err != nil {
		return fmt.Errorf("Error reading flags: %v", err)
	}
	if err = d.Set("foslic_cpu", o["foslic_cpu"]); err != nil {
		return fmt.Errorf("Error reading foslic_cpu: %v", err)
	}
	if err = d.Set("foslic_dr_site", o["foslic_dr_site"]); err != nil {
		return fmt.Errorf("Error reading foslic_dr_site: %v", err)
	}
	if err = d.Set("foslic_inst_time", o["foslic_inst_time"]); err != nil {
		return fmt.Errorf("Error reading foslic_inst_time: %v", err)
	}
	if err = d.Set("foslic_last_sync", o["foslic_last_sync"]); err != nil {
		return fmt.Errorf("Error reading foslic_last_sync: %v", err)
	}
	if err = d.Set("foslic_last_sync", o["foslic_last_sync"]); err != nil {
		return fmt.Errorf("Error reading foslic_last_sync: %v", err)
	}
	if err = d.Set("foslic_ram", o["foslic_ram"]); err != nil {
		return fmt.Errorf("Error reading foslic_ram: %v", err)
	}
	if err = d.Set("foslic_type", o["foslic_type"]); err != nil {
		return fmt.Errorf("Error reading foslic_type: %v", err)
	}
	if err = d.Set("foslic_utm", o["foslic_utm"]); err != nil {
		return fmt.Errorf("Error reading foslic_utm: %v", err)
	}
	if err = d.Set("fsw_cnt", o["fsw_cnt"]); err != nil {
		return fmt.Errorf("Error reading fsw_cnt: %v", err)
	}
	if err = d.Set("havsn", o["ha.vsn"]); err != nil {
		return fmt.Errorf("Error reading havsn: %v", err)
	}
	if err = d.Set("ha_group_id", o["ha_group_id"]); err != nil {
		return fmt.Errorf("Error reading ha_group_id: %v", err)
	}
	if err = d.Set("ha_group_name", o["ha_group_name"]); err != nil {
		return fmt.Errorf("Error reading ha_group_name: %v", err)
	}
	if err = d.Set("ha_mode", o["ha_mode"]); err != nil {
		return fmt.Errorf("Error reading ha_mode: %v", err)
	}
	if err = d.Set("ha_upgrade_mode", o["ha_upgrade_mode"]); err != nil {
		return fmt.Errorf("Error reading ha_upgrade_mode: %v", err)
	}
	if err = d.Set("hdisk_size", o["hdisk_size"]); err != nil {
		return fmt.Errorf("Error reading hdisk_size: %v", err)
	}
	if err = d.Set("hostname", o["hostname"]); err != nil {
		return fmt.Errorf("Error reading hostname: %v", err)
	}
	if err = d.Set("hw_generation", o["hw_generation"]); err != nil {
		return fmt.Errorf("Error reading hw_generation: %v", err)
	}
	if err = d.Set("hw_rev_major", o["hw_rev_major"]); err != nil {
		return fmt.Errorf("Error reading hw_rev_major: %v", err)
	}
	if err = d.Set("hw_rev_minor", o["hw_rev_minor"]); err != nil {
		return fmt.Errorf("Error reading hw_rev_minor: %v", err)
	}
	if err = d.Set("hyperscale", o["hyperscale"]); err != nil {
		return fmt.Errorf("Error reading hyperscale: %v", err)
	}
	if err = d.Set("ips_ext", o["ips_ext"]); err != nil {
		return fmt.Errorf("Error reading ips_ext: %v", err)
	}
	if err = d.Set("ips_ver", o["ips_ver"]); err != nil {
		return fmt.Errorf("Error reading ips_ver: %v", err)
	}
	if err = d.Set("last_checked", o["last_checked"]); err != nil {
		return fmt.Errorf("Error reading sn: %v", err)
	}
	if err = d.Set("last_resync", o["last_resync"]); err != nil {
		return fmt.Errorf("Error reading last_resync: %v", err)
	}
	if err = d.Set("latitude", o["latitude"]); err != nil {
		return fmt.Errorf("Error reading latitude: %v", err)
	}
	if err = d.Set("lic_flags", o["lic_flags"]); err != nil {
		return fmt.Errorf("Error reading lic_flags: %v", err)
	}
	if err = d.Set("lic_region", o["lic_region"]); err != nil {
		return fmt.Errorf("Error reading lic_region: %v", err)
	}
	if err = d.Set("location_from", o["location_from"]); err != nil {
		return fmt.Errorf("Error reading location_from: %v", err)
	}
	if err = d.Set("logdisk_size", o["logdisk_size"]); err != nil {
		return fmt.Errorf("Error reading logdisk_size: %v", err)
	}
	if err = d.Set("longitude", o["longitude"]); err != nil {
		return fmt.Errorf("Error reading longitude: %v", err)
	}
	if err = d.Set("maxvdom", o["maxvdom"]); err != nil {
		return fmt.Errorf("Error reading maxvdom: %v", err)
	}
	if err = d.Set("mgmt_if", o["mgmt_if"]); err != nil {
		return fmt.Errorf("Error reading sn: %v", err)
	}
	if err = d.Set("mgmt_uuid", o["mgmt_uuid"]); err != nil {
		return fmt.Errorf("Error reading mgmt_uuid: %v", err)
	}
	if err = d.Set("mgt_vdom", o["mgt_vdom"]); err != nil {
		return fmt.Errorf("Error reading mgt_vdom: %v", err)
	}
	if err = d.Set("module_sn", o["module_sn"]); err != nil {
		return fmt.Errorf("Error reading module_sn: %v", err)
	}
	if err = d.Set("nsxt_service_name", o["nsxt_service_name"]); err != nil {
		return fmt.Errorf("Error reading nsxt_service_name: %v", err)
	}
	if err = d.Set("prefer_img_ver", o["prefer_img_ver"]); err != nil {
		return fmt.Errorf("Error reading prefer_img_ver: %v", err)
	}
	if err = d.Set("prio", o["prio"]); err != nil {
		return fmt.Errorf("Error reading prio: %v", err)
	}
	if err = d.Set("private_key", o["private_key"]); err != nil {
		return fmt.Errorf("Error reading private_key: %v", err)
	}
	if err = d.Set("private_key_status", o["private_key_status"]); err != nil {
		return fmt.Errorf("Error reading private_key_status: %v", err)
	}
	if err = d.Set("psk", o["psk"]); err != nil {
		return fmt.Errorf("Error reading psk: %v", err)
	}
	if err = d.Set("relver_info", o["relver_info"]); err != nil {
		return fmt.Errorf("Error reading relver_info: %v", err)
	}
	if err = d.Set("role", o["role"]); err != nil {
		return fmt.Errorf("Error reading role: %v", err)
	}
	if err = d.Set("sov_sase_license", o["sov_sase_license"]); err != nil {
		return fmt.Errorf("Error reading sov_sase_license: %v", err)
	}
	if err = d.Set("tunnel_sn", o["tunnel_sn"]); err != nil {
		return fmt.Errorf("Error reading tunnel_sn: %v", err)
	}
	if err = d.Set("vdom", flattenDeviceVdom(o["vdom"], d, "vdom")); err != nil {
		return fmt.Errorf("Error reading sn: %v", err)
	}
	if err = d.Set("version", o["version"]); err != nil {
		return fmt.Errorf("Error reading version: %v", err)
	}
	if err = d.Set("vm_cpu", o["vm_cpu"]); err != nil {
		return fmt.Errorf("Error reading vm_cpu: %v", err)
	}
	if err = d.Set("vm_cpu_limit", o["vm_cpu_limit"]); err != nil {
		return fmt.Errorf("Error reading vm_cpu_limit: %v", err)
	}
	if err = d.Set("vm_lic_expire", o["vm_lic_expire"]); err != nil {
		return fmt.Errorf("Error reading vm_lic_expire: %v", err)
	}
	if err = d.Set("vm_lic_overdue_since", o["vm_lic_overdue_since"]); err != nil {
		return fmt.Errorf("Error reading vm_lic_overdue_since: %v", err)
	}
	if err = d.Set("vm_mem", o["vm_mem"]); err != nil {
		return fmt.Errorf("Error reading vm_mem: %v", err)
	}
	if err = d.Set("vm_mem_limit", o["vm_mem_limit"]); err != nil {
		return fmt.Errorf("Error reading vm_mem_limit: %v", err)
	}
	if err = d.Set("vm_payg_status", o["vm_payg_status"]); err != nil {
		return fmt.Errorf("Error reading vm_payg_status: %v", err)
	}
	if err = d.Set("vm_status", o["vm_status"]); err != nil {
		return fmt.Errorf("Error reading vm_status: %v", err)
	}

	return nil
}

func flattenDeviceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandDeviceAdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDeviceDevice(d *schema.ResourceData) (interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("adm_pass"); ok || d.HasChange("adm_pass") {
		t, err := expandDeviceDeviceAdmPass(d, v, "adm_pass")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adm_pass"] = t
		}
	}

	if v, ok := d.GetOk("adm_usr"); ok || d.HasChange("adm_usr") {
		t, err := expandDeviceDeviceAdmUsr(d, v, "adm_usr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adm_usr"] = t
		}
	}

	if v, ok := d.GetOk("authorizationtemplate"); ok || d.HasChange("authorizationtemplate") {
		t, err := expandDeviceDeviceAuthorizationTemplate(d, v, "authorizationtemplate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authorization template"] = t
		}
	}

	if v, ok := d.GetOk("desc"); ok || d.HasChange("desc") {
		t, err := expandDeviceDeviceDesc(d, v, "desc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["desc"] = t
		}
	}

	if v, ok := d.GetOk("deviceaction"); ok || d.HasChange("deviceaction") {
		t, err := expandDeviceDeviceDeviceAction(d, v, "deviceaction")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device action"] = t
		}
	}

	if v, ok := d.GetOk("deviceblueprint"); ok || d.HasChange("deviceblueprint") {
		t, err := expandDeviceDeviceDeviceBlueprint(d, v, "deviceblueprint")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adom"] = t
		}
	}

	if v, ok := d.GetOk("fazquota"); ok || d.HasChange("fazquota") {
		t, err := expandDeviceDeviceFazQuota(d, v, "fazquota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["faz.quota"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandDeviceDeviceIp(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("metafields"); ok || d.HasChange("metafields") {
		t, err := expandDeviceDeviceMetaFields(d, v, "metafields")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["meta fields"] = t
		}
	}

	if v, ok := d.GetOk("mgmt_mode"); ok || d.HasChange("mgmt_mode") {
		t, err := expandDeviceDeviceMgmtMode(d, v, "mgmt_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mgmt_mode"] = t
		}
	}

	if v, ok := d.GetOk("mr"); ok || d.HasChange("mr") {
		t, err := expandDeviceDeviceMr(d, v, "mr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mr"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandDeviceDeviceName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("os_type"); ok || d.HasChange("os_type") {
		t, err := expandDeviceDeviceOsType(d, v, "os_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["os_type"] = t
		}
	}

	if v, ok := d.GetOk("os_ver"); ok || d.HasChange("os_ver") {
		t, err := expandDeviceDeviceOsVer(d, v, "os_ver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["os_ver"] = t
		}
	}

	if v, ok := d.GetOk("patch"); ok || d.HasChange("patch") {
		t, err := expandDeviceDevicePatch(d, v, "patch")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["patch"] = t
		}
	}

	if v, ok := d.GetOk("platform_str"); ok || d.HasChange("platform_str") {
		t, err := expandDeviceDevicePlatformStr(d, v, "platform_str")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["platform_str"] = t
		}
	}

	if v, ok := d.GetOk("sn"); ok || d.HasChange("sn") {
		t, err := expandDeviceDeviceSn(d, v, "sn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sn"] = t
		}
	}

	return obj, nil
}

func expandDeviceDeviceAdmPass(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDeviceDeviceAdmUsr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDeviceDeviceAuthorizationTemplate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return convstr2list(v, nil), nil
}

func expandDeviceDeviceDesc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDeviceDeviceDeviceAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDeviceDeviceDeviceBlueprint(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDeviceDeviceFazQuota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDeviceDeviceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDeviceDeviceMetaFields(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDeviceDeviceMgmtMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDeviceDeviceMr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDeviceDeviceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDeviceDeviceOsType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDeviceDeviceOsVer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDeviceDevicePatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDeviceDevicePlatformStr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDeviceDeviceSn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDeviceFlags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandDeviceGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandDeviceGroupsName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandDeviceGroupsVdom(d, i["vdom"], pre_append)
		}

		if len(tmp) > 0 {
			result = append(result, tmp)
		}

		con += 1
	}

	return result, nil
}

func expandDeviceGroupsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDeviceGroupsVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectDevice(d *schema.ResourceData, operation, adomv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	obj["adom"] = adomv

	if operation == "create" {
		t, err := expandDeviceDevice(d)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device"] = t
		}

		if v, ok := d.GetOk("groups"); ok || d.HasChange("groups") {
			t, err := expandDeviceGroups(d, v, "groups")
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["groups"] = t
			}
		}
	} else {
		obj["device"] = d.Id()
	}

	// Using task ID to track the process
	obj["flags"] = []string{"create_task", "nonblocking"}

	return &obj, nil
}
