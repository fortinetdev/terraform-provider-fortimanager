// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Device table, most attributes are read-only and can only be changed internally. Refer to Device Manager Command module for API to add, delete, and manage devices.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceDvmdbDevice() *schema.Resource {
	return &schema.Resource{
		Create: resourceDvmdbDeviceCreate,
		Read:   resourceDvmdbDeviceRead,
		Update: resourceDvmdbDeviceUpdate,
		Delete: resourceDvmdbDeviceDelete,

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
			"adm_pass": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"adm_usr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
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
			"desc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dev_status": &schema.Schema{
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
			"fazquota": &schema.Schema{
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
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"fsw_cnt": &schema.Schema{
				Type:     schema.TypeInt,
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
			"ip": &schema.Schema{
				Type:     schema.TypeString,
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
			"metafields": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mgmt_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mgmt_if": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mgmt_mode": &schema.Schema{
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
			"mr": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"nsxt_service_name": &schema.Schema{
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
			"role": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"comments": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"metafields": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
			"vm_status": &schema.Schema{
				Type:     schema.TypeInt,
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

func resourceDvmdbDeviceCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectDvmdbDevice(d)
	if err != nil {
		return fmt.Errorf("Error creating DvmdbDevice resource while getting object: %v", err)
	}

	_, err = c.CreateDvmdbDevice(obj, adomv, nil)

	if err != nil {
		return fmt.Errorf("Error creating DvmdbDevice resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceDvmdbDeviceRead(d, m)
}

func resourceDvmdbDeviceUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectDvmdbDevice(d)
	if err != nil {
		return fmt.Errorf("Error updating DvmdbDevice resource while getting object: %v", err)
	}

	_, err = c.UpdateDvmdbDevice(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating DvmdbDevice resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceDvmdbDeviceRead(d, m)
}

func resourceDvmdbDeviceDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteDvmdbDevice(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting DvmdbDevice resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceDvmdbDeviceRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadDvmdbDevice(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading DvmdbDevice resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectDvmdbDevice(d, o)
	if err != nil {
		return fmt.Errorf("Error reading DvmdbDevice resource from API: %v", err)
	}
	return nil
}

func flattenDvmdbDeviceAdmPass(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenDvmdbDeviceAdmUsr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceAppVer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceAvVer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceBeta(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceBranchPt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceBuild(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceChecksum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceConfStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "unknown",
			1: "insync",
			2: "outofsync",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenDvmdbDeviceConnMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "active",
			1: "passive",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenDvmdbDeviceConnStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "UNKNOWN",
			1: "up",
			2: "down",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenDvmdbDeviceDbStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "unknown",
			1: "nomod",
			2: "mod",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenDvmdbDeviceDesc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceDevStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0:  "none",
			1:  "unknown",
			2:  "checkedin",
			3:  "inprogress",
			4:  "installed",
			5:  "aborted",
			6:  "sched",
			7:  "retry",
			8:  "canceled",
			9:  "pending",
			10: "retrieved",
			11: "changed_conf",
			12: "sync_fail",
			13: "timeout",
			14: "rev_revert",
			15: "auto_updated",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenDvmdbDeviceFapCnt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceFazFullAct(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceFazPerm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceFazQuota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceFazUsed(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceFexCnt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceFlags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1:         "has_hdd",
			2:         "vdom_enabled",
			8:         "discover",
			16:        "reload",
			1024:      "interim_build",
			4096:      "offline_mode",
			262144:    "is_model",
			524288:    "fips_mode",
			67108864:  "linked_to_model",
			268435456: "ip-conflict",
			536870912: "faz-autosync",
		}
		res := getEnumValbyBit(v, emap)
		return res
	}
	return v
}

func flattenDvmdbDeviceFoslicCpu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceFoslicDrSite(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			1: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenDvmdbDeviceFoslicInstTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceFoslicLastSync(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceFoslicRam(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceFoslicType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "temporary",
			1: "trial",
			2: "regular",
			3: "trial_expired",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenDvmdbDeviceFoslicUtm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1:  "fw",
			2:  "av",
			4:  "ips",
			8:  "app",
			16: "url",
			32: "utm",
			64: "fwb",
		}
		res := getEnumValbyBit(v, emap)
		return res
	}
	return v
}

func flattenDvmdbDeviceFswCnt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceHaGroupId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceHaGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceHaMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "standalone",
			1: "AP",
			2: "AA",
			3: "ELBC",
			4: "DUAL",
			5: "fmg-enabled",
			6: "autoscale",
			7: "unknown",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenDvmdbDeviceHdiskSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceHostname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceHwRevMajor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceHwRevMinor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceHyperscale(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceIpsExt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceIpsVer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceLastChecked(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceLastResync(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceLatitude(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceLicFlags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceLicRegion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceLocationFrom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceLogdiskSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceLongitude(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceMaxvdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceMetaFields(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceMgmtId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceMgmtIf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceMgmtMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "unreg",
			1: "fmg",
			2: "faz",
			3: "fmgfaz",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenDvmdbDeviceMgtVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceModuleSn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceMr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceNsxtServiceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceOsType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			-1: "unknown",
			0:  "fos",
			1:  "fsw",
			2:  "foc",
			3:  "fml",
			4:  "faz",
			5:  "fwb",
			6:  "fch",
			7:  "fct",
			8:  "log",
			9:  "fmg",
			10: "fsa",
			11: "fdd",
			12: "fac",
			13: "fpx",
			14: "fna",
			15: "ffw",
			16: "fsr",
			17: "fad",
			18: "fdc",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenDvmdbDeviceOsVer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			-1: "unknown",
			0:  "0.0",
			1:  "1.0",
			2:  "2.0",
			3:  "3.0",
			4:  "4.0",
			5:  "5.0",
			6:  "6.0",
			7:  "7.0",
			8:  "8.0",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenDvmdbDevicePatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDevicePlatformStr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDevicePreferImgVer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDevicePrio(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDevicePrivateKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDevicePrivateKeyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDevicePsk(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceRole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "master",
			1: "ha-slave",
			2: "autoscale-slave",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenDvmdbDeviceSn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceVdom(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comments"
		if _, ok := i["comments"]; ok {
			v := flattenDvmdbDeviceVdomComments(i["comments"], d, pre_append)
			tmp["comments"] = fortiAPISubPartPatch(v, "DvmdbDevice-Vdom-Comments")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metafields"
		if _, ok := i["meta fields"]; ok {
			v := flattenDvmdbDeviceVdomMetaFields(i["meta fields"], d, pre_append)
			tmp["metafields"] = fortiAPISubPartPatch(v, "DvmdbDevice-Vdom-MetaFields")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenDvmdbDeviceVdomName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "DvmdbDevice-Vdom-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "opmode"
		if _, ok := i["opmode"]; ok {
			v := flattenDvmdbDeviceVdomOpmode(i["opmode"], d, pre_append)
			tmp["opmode"] = fortiAPISubPartPatch(v, "DvmdbDevice-Vdom-Opmode")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rtm_prof_id"
		if _, ok := i["rtm_prof_id"]; ok {
			v := flattenDvmdbDeviceVdomRtmProfId(i["rtm_prof_id"], d, pre_append)
			tmp["rtm_prof_id"] = fortiAPISubPartPatch(v, "DvmdbDevice-Vdom-RtmProfId")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenDvmdbDeviceVdomStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "DvmdbDevice-Vdom-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vpn_id"
		if _, ok := i["vpn_id"]; ok {
			v := flattenDvmdbDeviceVdomVpnId(i["vpn_id"], d, pre_append)
			tmp["vpn_id"] = fortiAPISubPartPatch(v, "DvmdbDevice-Vdom-VpnId")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenDvmdbDeviceVdomComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceVdomMetaFields(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceVdomName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceVdomOpmode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			1: "nat",
			2: "transparent",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenDvmdbDeviceVdomRtmProfId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceVdomStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceVdomVpnId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceVmCpu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceVmCpuLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceVmLicExpire(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceVmMem(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceVmMemLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDvmdbDeviceVmStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectDvmdbDevice(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("adm_pass", flattenDvmdbDeviceAdmPass(o["adm_pass"], d, "adm_pass")); err != nil {
		if vv, ok := fortiAPIPatch(o["adm_pass"], "DvmdbDevice-AdmPass"); ok {
			if err = d.Set("adm_pass", vv); err != nil {
				return fmt.Errorf("Error reading adm_pass: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adm_pass: %v", err)
		}
	}

	if err = d.Set("adm_usr", flattenDvmdbDeviceAdmUsr(o["adm_usr"], d, "adm_usr")); err != nil {
		if vv, ok := fortiAPIPatch(o["adm_usr"], "DvmdbDevice-AdmUsr"); ok {
			if err = d.Set("adm_usr", vv); err != nil {
				return fmt.Errorf("Error reading adm_usr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading adm_usr: %v", err)
		}
	}

	if err = d.Set("app_ver", flattenDvmdbDeviceAppVer(o["app_ver"], d, "app_ver")); err != nil {
		if vv, ok := fortiAPIPatch(o["app_ver"], "DvmdbDevice-AppVer"); ok {
			if err = d.Set("app_ver", vv); err != nil {
				return fmt.Errorf("Error reading app_ver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading app_ver: %v", err)
		}
	}

	if err = d.Set("av_ver", flattenDvmdbDeviceAvVer(o["av_ver"], d, "av_ver")); err != nil {
		if vv, ok := fortiAPIPatch(o["av_ver"], "DvmdbDevice-AvVer"); ok {
			if err = d.Set("av_ver", vv); err != nil {
				return fmt.Errorf("Error reading av_ver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading av_ver: %v", err)
		}
	}

	if err = d.Set("beta", flattenDvmdbDeviceBeta(o["beta"], d, "beta")); err != nil {
		if vv, ok := fortiAPIPatch(o["beta"], "DvmdbDevice-Beta"); ok {
			if err = d.Set("beta", vv); err != nil {
				return fmt.Errorf("Error reading beta: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading beta: %v", err)
		}
	}

	if err = d.Set("branch_pt", flattenDvmdbDeviceBranchPt(o["branch_pt"], d, "branch_pt")); err != nil {
		if vv, ok := fortiAPIPatch(o["branch_pt"], "DvmdbDevice-BranchPt"); ok {
			if err = d.Set("branch_pt", vv); err != nil {
				return fmt.Errorf("Error reading branch_pt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading branch_pt: %v", err)
		}
	}

	if err = d.Set("build", flattenDvmdbDeviceBuild(o["build"], d, "build")); err != nil {
		if vv, ok := fortiAPIPatch(o["build"], "DvmdbDevice-Build"); ok {
			if err = d.Set("build", vv); err != nil {
				return fmt.Errorf("Error reading build: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading build: %v", err)
		}
	}

	if err = d.Set("checksum", flattenDvmdbDeviceChecksum(o["checksum"], d, "checksum")); err != nil {
		if vv, ok := fortiAPIPatch(o["checksum"], "DvmdbDevice-Checksum"); ok {
			if err = d.Set("checksum", vv); err != nil {
				return fmt.Errorf("Error reading checksum: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading checksum: %v", err)
		}
	}

	if err = d.Set("conf_status", flattenDvmdbDeviceConfStatus(o["conf_status"], d, "conf_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["conf_status"], "DvmdbDevice-ConfStatus"); ok {
			if err = d.Set("conf_status", vv); err != nil {
				return fmt.Errorf("Error reading conf_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading conf_status: %v", err)
		}
	}

	if err = d.Set("conn_mode", flattenDvmdbDeviceConnMode(o["conn_mode"], d, "conn_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["conn_mode"], "DvmdbDevice-ConnMode"); ok {
			if err = d.Set("conn_mode", vv); err != nil {
				return fmt.Errorf("Error reading conn_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading conn_mode: %v", err)
		}
	}

	if err = d.Set("conn_status", flattenDvmdbDeviceConnStatus(o["conn_status"], d, "conn_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["conn_status"], "DvmdbDevice-ConnStatus"); ok {
			if err = d.Set("conn_status", vv); err != nil {
				return fmt.Errorf("Error reading conn_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading conn_status: %v", err)
		}
	}

	if err = d.Set("db_status", flattenDvmdbDeviceDbStatus(o["db_status"], d, "db_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["db_status"], "DvmdbDevice-DbStatus"); ok {
			if err = d.Set("db_status", vv); err != nil {
				return fmt.Errorf("Error reading db_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading db_status: %v", err)
		}
	}

	if err = d.Set("desc", flattenDvmdbDeviceDesc(o["desc"], d, "desc")); err != nil {
		if vv, ok := fortiAPIPatch(o["desc"], "DvmdbDevice-Desc"); ok {
			if err = d.Set("desc", vv); err != nil {
				return fmt.Errorf("Error reading desc: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading desc: %v", err)
		}
	}

	if err = d.Set("dev_status", flattenDvmdbDeviceDevStatus(o["dev_status"], d, "dev_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["dev_status"], "DvmdbDevice-DevStatus"); ok {
			if err = d.Set("dev_status", vv); err != nil {
				return fmt.Errorf("Error reading dev_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dev_status: %v", err)
		}
	}

	if err = d.Set("fap_cnt", flattenDvmdbDeviceFapCnt(o["fap_cnt"], d, "fap_cnt")); err != nil {
		if vv, ok := fortiAPIPatch(o["fap_cnt"], "DvmdbDevice-FapCnt"); ok {
			if err = d.Set("fap_cnt", vv); err != nil {
				return fmt.Errorf("Error reading fap_cnt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fap_cnt: %v", err)
		}
	}

	if err = d.Set("fazfull_act", flattenDvmdbDeviceFazFullAct(o["faz.full_act"], d, "fazfull_act")); err != nil {
		if vv, ok := fortiAPIPatch(o["faz.full_act"], "DvmdbDevice-FazFullAct"); ok {
			if err = d.Set("fazfull_act", vv); err != nil {
				return fmt.Errorf("Error reading fazfull_act: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fazfull_act: %v", err)
		}
	}

	if err = d.Set("fazperm", flattenDvmdbDeviceFazPerm(o["faz.perm"], d, "fazperm")); err != nil {
		if vv, ok := fortiAPIPatch(o["faz.perm"], "DvmdbDevice-FazPerm"); ok {
			if err = d.Set("fazperm", vv); err != nil {
				return fmt.Errorf("Error reading fazperm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fazperm: %v", err)
		}
	}

	if err = d.Set("fazquota", flattenDvmdbDeviceFazQuota(o["faz.quota"], d, "fazquota")); err != nil {
		if vv, ok := fortiAPIPatch(o["faz.quota"], "DvmdbDevice-FazQuota"); ok {
			if err = d.Set("fazquota", vv); err != nil {
				return fmt.Errorf("Error reading fazquota: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fazquota: %v", err)
		}
	}

	if err = d.Set("fazused", flattenDvmdbDeviceFazUsed(o["faz.used"], d, "fazused")); err != nil {
		if vv, ok := fortiAPIPatch(o["faz.used"], "DvmdbDevice-FazUsed"); ok {
			if err = d.Set("fazused", vv); err != nil {
				return fmt.Errorf("Error reading fazused: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fazused: %v", err)
		}
	}

	if err = d.Set("fex_cnt", flattenDvmdbDeviceFexCnt(o["fex_cnt"], d, "fex_cnt")); err != nil {
		if vv, ok := fortiAPIPatch(o["fex_cnt"], "DvmdbDevice-FexCnt"); ok {
			if err = d.Set("fex_cnt", vv); err != nil {
				return fmt.Errorf("Error reading fex_cnt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fex_cnt: %v", err)
		}
	}

	if err = d.Set("flags", flattenDvmdbDeviceFlags(o["flags"], d, "flags")); err != nil {
		if vv, ok := fortiAPIPatch(o["flags"], "DvmdbDevice-Flags"); ok {
			if err = d.Set("flags", vv); err != nil {
				return fmt.Errorf("Error reading flags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading flags: %v", err)
		}
	}

	if err = d.Set("foslic_cpu", flattenDvmdbDeviceFoslicCpu(o["foslic_cpu"], d, "foslic_cpu")); err != nil {
		if vv, ok := fortiAPIPatch(o["foslic_cpu"], "DvmdbDevice-FoslicCpu"); ok {
			if err = d.Set("foslic_cpu", vv); err != nil {
				return fmt.Errorf("Error reading foslic_cpu: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading foslic_cpu: %v", err)
		}
	}

	if err = d.Set("foslic_dr_site", flattenDvmdbDeviceFoslicDrSite(o["foslic_dr_site"], d, "foslic_dr_site")); err != nil {
		if vv, ok := fortiAPIPatch(o["foslic_dr_site"], "DvmdbDevice-FoslicDrSite"); ok {
			if err = d.Set("foslic_dr_site", vv); err != nil {
				return fmt.Errorf("Error reading foslic_dr_site: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading foslic_dr_site: %v", err)
		}
	}

	if err = d.Set("foslic_inst_time", flattenDvmdbDeviceFoslicInstTime(o["foslic_inst_time"], d, "foslic_inst_time")); err != nil {
		if vv, ok := fortiAPIPatch(o["foslic_inst_time"], "DvmdbDevice-FoslicInstTime"); ok {
			if err = d.Set("foslic_inst_time", vv); err != nil {
				return fmt.Errorf("Error reading foslic_inst_time: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading foslic_inst_time: %v", err)
		}
	}

	if err = d.Set("foslic_last_sync", flattenDvmdbDeviceFoslicLastSync(o["foslic_last_sync"], d, "foslic_last_sync")); err != nil {
		if vv, ok := fortiAPIPatch(o["foslic_last_sync"], "DvmdbDevice-FoslicLastSync"); ok {
			if err = d.Set("foslic_last_sync", vv); err != nil {
				return fmt.Errorf("Error reading foslic_last_sync: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading foslic_last_sync: %v", err)
		}
	}

	if err = d.Set("foslic_ram", flattenDvmdbDeviceFoslicRam(o["foslic_ram"], d, "foslic_ram")); err != nil {
		if vv, ok := fortiAPIPatch(o["foslic_ram"], "DvmdbDevice-FoslicRam"); ok {
			if err = d.Set("foslic_ram", vv); err != nil {
				return fmt.Errorf("Error reading foslic_ram: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading foslic_ram: %v", err)
		}
	}

	if err = d.Set("foslic_type", flattenDvmdbDeviceFoslicType(o["foslic_type"], d, "foslic_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["foslic_type"], "DvmdbDevice-FoslicType"); ok {
			if err = d.Set("foslic_type", vv); err != nil {
				return fmt.Errorf("Error reading foslic_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading foslic_type: %v", err)
		}
	}

	if err = d.Set("foslic_utm", flattenDvmdbDeviceFoslicUtm(o["foslic_utm"], d, "foslic_utm")); err != nil {
		if vv, ok := fortiAPIPatch(o["foslic_utm"], "DvmdbDevice-FoslicUtm"); ok {
			if err = d.Set("foslic_utm", vv); err != nil {
				return fmt.Errorf("Error reading foslic_utm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading foslic_utm: %v", err)
		}
	}

	if err = d.Set("fsw_cnt", flattenDvmdbDeviceFswCnt(o["fsw_cnt"], d, "fsw_cnt")); err != nil {
		if vv, ok := fortiAPIPatch(o["fsw_cnt"], "DvmdbDevice-FswCnt"); ok {
			if err = d.Set("fsw_cnt", vv); err != nil {
				return fmt.Errorf("Error reading fsw_cnt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fsw_cnt: %v", err)
		}
	}

	if err = d.Set("ha_group_id", flattenDvmdbDeviceHaGroupId(o["ha_group_id"], d, "ha_group_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["ha_group_id"], "DvmdbDevice-HaGroupId"); ok {
			if err = d.Set("ha_group_id", vv); err != nil {
				return fmt.Errorf("Error reading ha_group_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ha_group_id: %v", err)
		}
	}

	if err = d.Set("ha_group_name", flattenDvmdbDeviceHaGroupName(o["ha_group_name"], d, "ha_group_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["ha_group_name"], "DvmdbDevice-HaGroupName"); ok {
			if err = d.Set("ha_group_name", vv); err != nil {
				return fmt.Errorf("Error reading ha_group_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ha_group_name: %v", err)
		}
	}

	if err = d.Set("ha_mode", flattenDvmdbDeviceHaMode(o["ha_mode"], d, "ha_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["ha_mode"], "DvmdbDevice-HaMode"); ok {
			if err = d.Set("ha_mode", vv); err != nil {
				return fmt.Errorf("Error reading ha_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ha_mode: %v", err)
		}
	}

	if err = d.Set("hdisk_size", flattenDvmdbDeviceHdiskSize(o["hdisk_size"], d, "hdisk_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["hdisk_size"], "DvmdbDevice-HdiskSize"); ok {
			if err = d.Set("hdisk_size", vv); err != nil {
				return fmt.Errorf("Error reading hdisk_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hdisk_size: %v", err)
		}
	}

	if err = d.Set("hostname", flattenDvmdbDeviceHostname(o["hostname"], d, "hostname")); err != nil {
		if vv, ok := fortiAPIPatch(o["hostname"], "DvmdbDevice-Hostname"); ok {
			if err = d.Set("hostname", vv); err != nil {
				return fmt.Errorf("Error reading hostname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hostname: %v", err)
		}
	}

	if err = d.Set("hw_rev_major", flattenDvmdbDeviceHwRevMajor(o["hw_rev_major"], d, "hw_rev_major")); err != nil {
		if vv, ok := fortiAPIPatch(o["hw_rev_major"], "DvmdbDevice-HwRevMajor"); ok {
			if err = d.Set("hw_rev_major", vv); err != nil {
				return fmt.Errorf("Error reading hw_rev_major: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hw_rev_major: %v", err)
		}
	}

	if err = d.Set("hw_rev_minor", flattenDvmdbDeviceHwRevMinor(o["hw_rev_minor"], d, "hw_rev_minor")); err != nil {
		if vv, ok := fortiAPIPatch(o["hw_rev_minor"], "DvmdbDevice-HwRevMinor"); ok {
			if err = d.Set("hw_rev_minor", vv); err != nil {
				return fmt.Errorf("Error reading hw_rev_minor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hw_rev_minor: %v", err)
		}
	}

	if err = d.Set("hyperscale", flattenDvmdbDeviceHyperscale(o["hyperscale"], d, "hyperscale")); err != nil {
		if vv, ok := fortiAPIPatch(o["hyperscale"], "DvmdbDevice-Hyperscale"); ok {
			if err = d.Set("hyperscale", vv); err != nil {
				return fmt.Errorf("Error reading hyperscale: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hyperscale: %v", err)
		}
	}

	if err = d.Set("ip", flattenDvmdbDeviceIp(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "DvmdbDevice-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("ips_ext", flattenDvmdbDeviceIpsExt(o["ips_ext"], d, "ips_ext")); err != nil {
		if vv, ok := fortiAPIPatch(o["ips_ext"], "DvmdbDevice-IpsExt"); ok {
			if err = d.Set("ips_ext", vv); err != nil {
				return fmt.Errorf("Error reading ips_ext: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ips_ext: %v", err)
		}
	}

	if err = d.Set("ips_ver", flattenDvmdbDeviceIpsVer(o["ips_ver"], d, "ips_ver")); err != nil {
		if vv, ok := fortiAPIPatch(o["ips_ver"], "DvmdbDevice-IpsVer"); ok {
			if err = d.Set("ips_ver", vv); err != nil {
				return fmt.Errorf("Error reading ips_ver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ips_ver: %v", err)
		}
	}

	if err = d.Set("last_checked", flattenDvmdbDeviceLastChecked(o["last_checked"], d, "last_checked")); err != nil {
		if vv, ok := fortiAPIPatch(o["last_checked"], "DvmdbDevice-LastChecked"); ok {
			if err = d.Set("last_checked", vv); err != nil {
				return fmt.Errorf("Error reading last_checked: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading last_checked: %v", err)
		}
	}

	if err = d.Set("last_resync", flattenDvmdbDeviceLastResync(o["last_resync"], d, "last_resync")); err != nil {
		if vv, ok := fortiAPIPatch(o["last_resync"], "DvmdbDevice-LastResync"); ok {
			if err = d.Set("last_resync", vv); err != nil {
				return fmt.Errorf("Error reading last_resync: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading last_resync: %v", err)
		}
	}

	if err = d.Set("latitude", flattenDvmdbDeviceLatitude(o["latitude"], d, "latitude")); err != nil {
		if vv, ok := fortiAPIPatch(o["latitude"], "DvmdbDevice-Latitude"); ok {
			if err = d.Set("latitude", vv); err != nil {
				return fmt.Errorf("Error reading latitude: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading latitude: %v", err)
		}
	}

	if err = d.Set("lic_flags", flattenDvmdbDeviceLicFlags(o["lic_flags"], d, "lic_flags")); err != nil {
		if vv, ok := fortiAPIPatch(o["lic_flags"], "DvmdbDevice-LicFlags"); ok {
			if err = d.Set("lic_flags", vv); err != nil {
				return fmt.Errorf("Error reading lic_flags: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lic_flags: %v", err)
		}
	}

	if err = d.Set("lic_region", flattenDvmdbDeviceLicRegion(o["lic_region"], d, "lic_region")); err != nil {
		if vv, ok := fortiAPIPatch(o["lic_region"], "DvmdbDevice-LicRegion"); ok {
			if err = d.Set("lic_region", vv); err != nil {
				return fmt.Errorf("Error reading lic_region: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lic_region: %v", err)
		}
	}

	if err = d.Set("location_from", flattenDvmdbDeviceLocationFrom(o["location_from"], d, "location_from")); err != nil {
		if vv, ok := fortiAPIPatch(o["location_from"], "DvmdbDevice-LocationFrom"); ok {
			if err = d.Set("location_from", vv); err != nil {
				return fmt.Errorf("Error reading location_from: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading location_from: %v", err)
		}
	}

	if err = d.Set("logdisk_size", flattenDvmdbDeviceLogdiskSize(o["logdisk_size"], d, "logdisk_size")); err != nil {
		if vv, ok := fortiAPIPatch(o["logdisk_size"], "DvmdbDevice-LogdiskSize"); ok {
			if err = d.Set("logdisk_size", vv); err != nil {
				return fmt.Errorf("Error reading logdisk_size: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logdisk_size: %v", err)
		}
	}

	if err = d.Set("longitude", flattenDvmdbDeviceLongitude(o["longitude"], d, "longitude")); err != nil {
		if vv, ok := fortiAPIPatch(o["longitude"], "DvmdbDevice-Longitude"); ok {
			if err = d.Set("longitude", vv); err != nil {
				return fmt.Errorf("Error reading longitude: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading longitude: %v", err)
		}
	}

	if err = d.Set("maxvdom", flattenDvmdbDeviceMaxvdom(o["maxvdom"], d, "maxvdom")); err != nil {
		if vv, ok := fortiAPIPatch(o["maxvdom"], "DvmdbDevice-Maxvdom"); ok {
			if err = d.Set("maxvdom", vv); err != nil {
				return fmt.Errorf("Error reading maxvdom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading maxvdom: %v", err)
		}
	}

	if err = d.Set("metafields", flattenDvmdbDeviceMetaFields(o["meta fields"], d, "metafields")); err != nil {
		if vv, ok := fortiAPIPatch(o["meta fields"], "DvmdbDevice-MetaFields"); ok {
			if err = d.Set("metafields", vv); err != nil {
				return fmt.Errorf("Error reading metafields: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading metafields: %v", err)
		}
	}

	if err = d.Set("mgmt_id", flattenDvmdbDeviceMgmtId(o["mgmt_id"], d, "mgmt_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["mgmt_id"], "DvmdbDevice-MgmtId"); ok {
			if err = d.Set("mgmt_id", vv); err != nil {
				return fmt.Errorf("Error reading mgmt_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mgmt_id: %v", err)
		}
	}

	if err = d.Set("mgmt_if", flattenDvmdbDeviceMgmtIf(o["mgmt_if"], d, "mgmt_if")); err != nil {
		if vv, ok := fortiAPIPatch(o["mgmt_if"], "DvmdbDevice-MgmtIf"); ok {
			if err = d.Set("mgmt_if", vv); err != nil {
				return fmt.Errorf("Error reading mgmt_if: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mgmt_if: %v", err)
		}
	}

	if err = d.Set("mgmt_mode", flattenDvmdbDeviceMgmtMode(o["mgmt_mode"], d, "mgmt_mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["mgmt_mode"], "DvmdbDevice-MgmtMode"); ok {
			if err = d.Set("mgmt_mode", vv); err != nil {
				return fmt.Errorf("Error reading mgmt_mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mgmt_mode: %v", err)
		}
	}

	if err = d.Set("mgt_vdom", flattenDvmdbDeviceMgtVdom(o["mgt_vdom"], d, "mgt_vdom")); err != nil {
		if vv, ok := fortiAPIPatch(o["mgt_vdom"], "DvmdbDevice-MgtVdom"); ok {
			if err = d.Set("mgt_vdom", vv); err != nil {
				return fmt.Errorf("Error reading mgt_vdom: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mgt_vdom: %v", err)
		}
	}

	if err = d.Set("module_sn", flattenDvmdbDeviceModuleSn(o["module_sn"], d, "module_sn")); err != nil {
		if vv, ok := fortiAPIPatch(o["module_sn"], "DvmdbDevice-ModuleSn"); ok {
			if err = d.Set("module_sn", vv); err != nil {
				return fmt.Errorf("Error reading module_sn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading module_sn: %v", err)
		}
	}

	if err = d.Set("mr", flattenDvmdbDeviceMr(o["mr"], d, "mr")); err != nil {
		if vv, ok := fortiAPIPatch(o["mr"], "DvmdbDevice-Mr"); ok {
			if err = d.Set("mr", vv); err != nil {
				return fmt.Errorf("Error reading mr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mr: %v", err)
		}
	}

	if err = d.Set("name", flattenDvmdbDeviceName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "DvmdbDevice-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("nsxt_service_name", flattenDvmdbDeviceNsxtServiceName(o["nsxt_service_name"], d, "nsxt_service_name")); err != nil {
		if vv, ok := fortiAPIPatch(o["nsxt_service_name"], "DvmdbDevice-NsxtServiceName"); ok {
			if err = d.Set("nsxt_service_name", vv); err != nil {
				return fmt.Errorf("Error reading nsxt_service_name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nsxt_service_name: %v", err)
		}
	}

	if err = d.Set("os_type", flattenDvmdbDeviceOsType(o["os_type"], d, "os_type")); err != nil {
		if vv, ok := fortiAPIPatch(o["os_type"], "DvmdbDevice-OsType"); ok {
			if err = d.Set("os_type", vv); err != nil {
				return fmt.Errorf("Error reading os_type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading os_type: %v", err)
		}
	}

	if err = d.Set("os_ver", flattenDvmdbDeviceOsVer(o["os_ver"], d, "os_ver")); err != nil {
		if vv, ok := fortiAPIPatch(o["os_ver"], "DvmdbDevice-OsVer"); ok {
			if err = d.Set("os_ver", vv); err != nil {
				return fmt.Errorf("Error reading os_ver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading os_ver: %v", err)
		}
	}

	if err = d.Set("patch", flattenDvmdbDevicePatch(o["patch"], d, "patch")); err != nil {
		if vv, ok := fortiAPIPatch(o["patch"], "DvmdbDevice-Patch"); ok {
			if err = d.Set("patch", vv); err != nil {
				return fmt.Errorf("Error reading patch: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading patch: %v", err)
		}
	}

	if err = d.Set("platform_str", flattenDvmdbDevicePlatformStr(o["platform_str"], d, "platform_str")); err != nil {
		if vv, ok := fortiAPIPatch(o["platform_str"], "DvmdbDevice-PlatformStr"); ok {
			if err = d.Set("platform_str", vv); err != nil {
				return fmt.Errorf("Error reading platform_str: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading platform_str: %v", err)
		}
	}

	if err = d.Set("prefer_img_ver", flattenDvmdbDevicePreferImgVer(o["prefer_img_ver"], d, "prefer_img_ver")); err != nil {
		if vv, ok := fortiAPIPatch(o["prefer_img_ver"], "DvmdbDevice-PreferImgVer"); ok {
			if err = d.Set("prefer_img_ver", vv); err != nil {
				return fmt.Errorf("Error reading prefer_img_ver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prefer_img_ver: %v", err)
		}
	}

	if err = d.Set("prio", flattenDvmdbDevicePrio(o["prio"], d, "prio")); err != nil {
		if vv, ok := fortiAPIPatch(o["prio"], "DvmdbDevice-Prio"); ok {
			if err = d.Set("prio", vv); err != nil {
				return fmt.Errorf("Error reading prio: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading prio: %v", err)
		}
	}

	if err = d.Set("private_key", flattenDvmdbDevicePrivateKey(o["private_key"], d, "private_key")); err != nil {
		if vv, ok := fortiAPIPatch(o["private_key"], "DvmdbDevice-PrivateKey"); ok {
			if err = d.Set("private_key", vv); err != nil {
				return fmt.Errorf("Error reading private_key: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading private_key: %v", err)
		}
	}

	if err = d.Set("private_key_status", flattenDvmdbDevicePrivateKeyStatus(o["private_key_status"], d, "private_key_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["private_key_status"], "DvmdbDevice-PrivateKeyStatus"); ok {
			if err = d.Set("private_key_status", vv); err != nil {
				return fmt.Errorf("Error reading private_key_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading private_key_status: %v", err)
		}
	}

	if err = d.Set("psk", flattenDvmdbDevicePsk(o["psk"], d, "psk")); err != nil {
		if vv, ok := fortiAPIPatch(o["psk"], "DvmdbDevice-Psk"); ok {
			if err = d.Set("psk", vv); err != nil {
				return fmt.Errorf("Error reading psk: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading psk: %v", err)
		}
	}

	if err = d.Set("role", flattenDvmdbDeviceRole(o["role"], d, "role")); err != nil {
		if vv, ok := fortiAPIPatch(o["role"], "DvmdbDevice-Role"); ok {
			if err = d.Set("role", vv); err != nil {
				return fmt.Errorf("Error reading role: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading role: %v", err)
		}
	}

	if err = d.Set("sn", flattenDvmdbDeviceSn(o["sn"], d, "sn")); err != nil {
		if vv, ok := fortiAPIPatch(o["sn"], "DvmdbDevice-Sn"); ok {
			if err = d.Set("sn", vv); err != nil {
				return fmt.Errorf("Error reading sn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sn: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("vdom", flattenDvmdbDeviceVdom(o["vdom"], d, "vdom")); err != nil {
			if vv, ok := fortiAPIPatch(o["vdom"], "DvmdbDevice-Vdom"); ok {
				if err = d.Set("vdom", vv); err != nil {
					return fmt.Errorf("Error reading vdom: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading vdom: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("vdom"); ok {
			if err = d.Set("vdom", flattenDvmdbDeviceVdom(o["vdom"], d, "vdom")); err != nil {
				if vv, ok := fortiAPIPatch(o["vdom"], "DvmdbDevice-Vdom"); ok {
					if err = d.Set("vdom", vv); err != nil {
						return fmt.Errorf("Error reading vdom: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading vdom: %v", err)
				}
			}
		}
	}

	if err = d.Set("version", flattenDvmdbDeviceVersion(o["version"], d, "version")); err != nil {
		if vv, ok := fortiAPIPatch(o["version"], "DvmdbDevice-Version"); ok {
			if err = d.Set("version", vv); err != nil {
				return fmt.Errorf("Error reading version: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading version: %v", err)
		}
	}

	if err = d.Set("vm_cpu", flattenDvmdbDeviceVmCpu(o["vm_cpu"], d, "vm_cpu")); err != nil {
		if vv, ok := fortiAPIPatch(o["vm_cpu"], "DvmdbDevice-VmCpu"); ok {
			if err = d.Set("vm_cpu", vv); err != nil {
				return fmt.Errorf("Error reading vm_cpu: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vm_cpu: %v", err)
		}
	}

	if err = d.Set("vm_cpu_limit", flattenDvmdbDeviceVmCpuLimit(o["vm_cpu_limit"], d, "vm_cpu_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["vm_cpu_limit"], "DvmdbDevice-VmCpuLimit"); ok {
			if err = d.Set("vm_cpu_limit", vv); err != nil {
				return fmt.Errorf("Error reading vm_cpu_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vm_cpu_limit: %v", err)
		}
	}

	if err = d.Set("vm_lic_expire", flattenDvmdbDeviceVmLicExpire(o["vm_lic_expire"], d, "vm_lic_expire")); err != nil {
		if vv, ok := fortiAPIPatch(o["vm_lic_expire"], "DvmdbDevice-VmLicExpire"); ok {
			if err = d.Set("vm_lic_expire", vv); err != nil {
				return fmt.Errorf("Error reading vm_lic_expire: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vm_lic_expire: %v", err)
		}
	}

	if err = d.Set("vm_mem", flattenDvmdbDeviceVmMem(o["vm_mem"], d, "vm_mem")); err != nil {
		if vv, ok := fortiAPIPatch(o["vm_mem"], "DvmdbDevice-VmMem"); ok {
			if err = d.Set("vm_mem", vv); err != nil {
				return fmt.Errorf("Error reading vm_mem: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vm_mem: %v", err)
		}
	}

	if err = d.Set("vm_mem_limit", flattenDvmdbDeviceVmMemLimit(o["vm_mem_limit"], d, "vm_mem_limit")); err != nil {
		if vv, ok := fortiAPIPatch(o["vm_mem_limit"], "DvmdbDevice-VmMemLimit"); ok {
			if err = d.Set("vm_mem_limit", vv); err != nil {
				return fmt.Errorf("Error reading vm_mem_limit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vm_mem_limit: %v", err)
		}
	}

	if err = d.Set("vm_status", flattenDvmdbDeviceVmStatus(o["vm_status"], d, "vm_status")); err != nil {
		if vv, ok := fortiAPIPatch(o["vm_status"], "DvmdbDevice-VmStatus"); ok {
			if err = d.Set("vm_status", vv); err != nil {
				return fmt.Errorf("Error reading vm_status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading vm_status: %v", err)
		}
	}

	return nil
}

func flattenDvmdbDeviceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandDvmdbDeviceAdmPass(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandDvmdbDeviceAdmUsr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceAppVer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceAvVer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceBeta(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceBranchPt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceBuild(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceChecksum(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceConfStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceConnMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceConnStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceDbStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceDesc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceDevStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceFapCnt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceFazFullAct(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceFazPerm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceFazQuota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceFazUsed(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceFexCnt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceFlags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandDvmdbDeviceFoslicCpu(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceFoslicDrSite(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceFoslicInstTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceFoslicLastSync(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceFoslicRam(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceFoslicType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceFoslicUtm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandDvmdbDeviceFswCnt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceHaGroupId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceHaGroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceHaMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceHdiskSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceHostname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceHwRevMajor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceHwRevMinor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceHyperscale(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceIpsExt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceIpsVer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceLastChecked(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceLastResync(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceLatitude(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceLicFlags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceLicRegion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceLocationFrom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceLogdiskSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceLongitude(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceMaxvdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceMetaFields(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceMgmtId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceMgmtIf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceMgmtMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceMgtVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceModuleSn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceMr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceNsxtServiceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceOsType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceOsVer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDevicePatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDevicePlatformStr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDevicePreferImgVer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDevicePrio(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDevicePrivateKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDevicePrivateKeyStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDevicePsk(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceRole(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceSn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comments"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["comments"], _ = expandDvmdbDeviceVdomComments(d, i["comments"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metafields"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["meta fields"], _ = expandDvmdbDeviceVdomMetaFields(d, i["metafields"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandDvmdbDeviceVdomName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "opmode"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["opmode"], _ = expandDvmdbDeviceVdomOpmode(d, i["opmode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rtm_prof_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rtm_prof_id"], _ = expandDvmdbDeviceVdomRtmProfId(d, i["rtm_prof_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandDvmdbDeviceVdomStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vpn_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vpn_id"], _ = expandDvmdbDeviceVdomVpnId(d, i["vpn_id"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandDvmdbDeviceVdomComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceVdomMetaFields(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceVdomName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceVdomOpmode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceVdomRtmProfId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceVdomStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceVdomVpnId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceVmCpu(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceVmCpuLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceVmLicExpire(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceVmMem(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceVmMemLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDvmdbDeviceVmStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectDvmdbDevice(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("adm_pass"); ok {
		t, err := expandDvmdbDeviceAdmPass(d, v, "adm_pass")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adm_pass"] = t
		}
	}

	if v, ok := d.GetOk("adm_usr"); ok {
		t, err := expandDvmdbDeviceAdmUsr(d, v, "adm_usr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adm_usr"] = t
		}
	}

	if v, ok := d.GetOk("app_ver"); ok {
		t, err := expandDvmdbDeviceAppVer(d, v, "app_ver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["app_ver"] = t
		}
	}

	if v, ok := d.GetOk("av_ver"); ok {
		t, err := expandDvmdbDeviceAvVer(d, v, "av_ver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av_ver"] = t
		}
	}

	if v, ok := d.GetOk("beta"); ok {
		t, err := expandDvmdbDeviceBeta(d, v, "beta")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["beta"] = t
		}
	}

	if v, ok := d.GetOk("branch_pt"); ok {
		t, err := expandDvmdbDeviceBranchPt(d, v, "branch_pt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["branch_pt"] = t
		}
	}

	if v, ok := d.GetOk("build"); ok {
		t, err := expandDvmdbDeviceBuild(d, v, "build")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["build"] = t
		}
	}

	if v, ok := d.GetOk("checksum"); ok {
		t, err := expandDvmdbDeviceChecksum(d, v, "checksum")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["checksum"] = t
		}
	}

	if v, ok := d.GetOk("conf_status"); ok {
		t, err := expandDvmdbDeviceConfStatus(d, v, "conf_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["conf_status"] = t
		}
	}

	if v, ok := d.GetOk("conn_mode"); ok {
		t, err := expandDvmdbDeviceConnMode(d, v, "conn_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["conn_mode"] = t
		}
	}

	if v, ok := d.GetOk("conn_status"); ok {
		t, err := expandDvmdbDeviceConnStatus(d, v, "conn_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["conn_status"] = t
		}
	}

	if v, ok := d.GetOk("db_status"); ok {
		t, err := expandDvmdbDeviceDbStatus(d, v, "db_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["db_status"] = t
		}
	}

	if v, ok := d.GetOk("desc"); ok {
		t, err := expandDvmdbDeviceDesc(d, v, "desc")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["desc"] = t
		}
	}

	if v, ok := d.GetOk("dev_status"); ok {
		t, err := expandDvmdbDeviceDevStatus(d, v, "dev_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dev_status"] = t
		}
	}

	if v, ok := d.GetOk("fap_cnt"); ok {
		t, err := expandDvmdbDeviceFapCnt(d, v, "fap_cnt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fap_cnt"] = t
		}
	}

	if v, ok := d.GetOk("fazfull_act"); ok {
		t, err := expandDvmdbDeviceFazFullAct(d, v, "fazfull_act")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["faz.full_act"] = t
		}
	}

	if v, ok := d.GetOk("fazperm"); ok {
		t, err := expandDvmdbDeviceFazPerm(d, v, "fazperm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["faz.perm"] = t
		}
	}

	if v, ok := d.GetOk("fazquota"); ok {
		t, err := expandDvmdbDeviceFazQuota(d, v, "fazquota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["faz.quota"] = t
		}
	}

	if v, ok := d.GetOk("fazused"); ok {
		t, err := expandDvmdbDeviceFazUsed(d, v, "fazused")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["faz.used"] = t
		}
	}

	if v, ok := d.GetOk("fex_cnt"); ok {
		t, err := expandDvmdbDeviceFexCnt(d, v, "fex_cnt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fex_cnt"] = t
		}
	}

	if v, ok := d.GetOk("flags"); ok {
		t, err := expandDvmdbDeviceFlags(d, v, "flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flags"] = t
		}
	}

	if v, ok := d.GetOk("foslic_cpu"); ok {
		t, err := expandDvmdbDeviceFoslicCpu(d, v, "foslic_cpu")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["foslic_cpu"] = t
		}
	}

	if v, ok := d.GetOk("foslic_dr_site"); ok {
		t, err := expandDvmdbDeviceFoslicDrSite(d, v, "foslic_dr_site")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["foslic_dr_site"] = t
		}
	}

	if v, ok := d.GetOk("foslic_inst_time"); ok {
		t, err := expandDvmdbDeviceFoslicInstTime(d, v, "foslic_inst_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["foslic_inst_time"] = t
		}
	}

	if v, ok := d.GetOk("foslic_last_sync"); ok {
		t, err := expandDvmdbDeviceFoslicLastSync(d, v, "foslic_last_sync")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["foslic_last_sync"] = t
		}
	}

	if v, ok := d.GetOk("foslic_ram"); ok {
		t, err := expandDvmdbDeviceFoslicRam(d, v, "foslic_ram")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["foslic_ram"] = t
		}
	}

	if v, ok := d.GetOk("foslic_type"); ok {
		t, err := expandDvmdbDeviceFoslicType(d, v, "foslic_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["foslic_type"] = t
		}
	}

	if v, ok := d.GetOk("foslic_utm"); ok {
		t, err := expandDvmdbDeviceFoslicUtm(d, v, "foslic_utm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["foslic_utm"] = t
		}
	}

	if v, ok := d.GetOk("fsw_cnt"); ok {
		t, err := expandDvmdbDeviceFswCnt(d, v, "fsw_cnt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fsw_cnt"] = t
		}
	}

	if v, ok := d.GetOk("ha_group_id"); ok {
		t, err := expandDvmdbDeviceHaGroupId(d, v, "ha_group_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha_group_id"] = t
		}
	}

	if v, ok := d.GetOk("ha_group_name"); ok {
		t, err := expandDvmdbDeviceHaGroupName(d, v, "ha_group_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha_group_name"] = t
		}
	}

	if v, ok := d.GetOk("ha_mode"); ok {
		t, err := expandDvmdbDeviceHaMode(d, v, "ha_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha_mode"] = t
		}
	}

	if v, ok := d.GetOk("hdisk_size"); ok {
		t, err := expandDvmdbDeviceHdiskSize(d, v, "hdisk_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hdisk_size"] = t
		}
	}

	if v, ok := d.GetOk("hostname"); ok {
		t, err := expandDvmdbDeviceHostname(d, v, "hostname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hostname"] = t
		}
	}

	if v, ok := d.GetOk("hw_rev_major"); ok {
		t, err := expandDvmdbDeviceHwRevMajor(d, v, "hw_rev_major")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hw_rev_major"] = t
		}
	}

	if v, ok := d.GetOk("hw_rev_minor"); ok {
		t, err := expandDvmdbDeviceHwRevMinor(d, v, "hw_rev_minor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hw_rev_minor"] = t
		}
	}

	if v, ok := d.GetOk("hyperscale"); ok {
		t, err := expandDvmdbDeviceHyperscale(d, v, "hyperscale")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hyperscale"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok {
		t, err := expandDvmdbDeviceIp(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("ips_ext"); ok {
		t, err := expandDvmdbDeviceIpsExt(d, v, "ips_ext")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips_ext"] = t
		}
	}

	if v, ok := d.GetOk("ips_ver"); ok {
		t, err := expandDvmdbDeviceIpsVer(d, v, "ips_ver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips_ver"] = t
		}
	}

	if v, ok := d.GetOk("last_checked"); ok {
		t, err := expandDvmdbDeviceLastChecked(d, v, "last_checked")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["last_checked"] = t
		}
	}

	if v, ok := d.GetOk("last_resync"); ok {
		t, err := expandDvmdbDeviceLastResync(d, v, "last_resync")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["last_resync"] = t
		}
	}

	if v, ok := d.GetOk("latitude"); ok {
		t, err := expandDvmdbDeviceLatitude(d, v, "latitude")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["latitude"] = t
		}
	}

	if v, ok := d.GetOk("lic_flags"); ok {
		t, err := expandDvmdbDeviceLicFlags(d, v, "lic_flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lic_flags"] = t
		}
	}

	if v, ok := d.GetOk("lic_region"); ok {
		t, err := expandDvmdbDeviceLicRegion(d, v, "lic_region")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lic_region"] = t
		}
	}

	if v, ok := d.GetOk("location_from"); ok {
		t, err := expandDvmdbDeviceLocationFrom(d, v, "location_from")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["location_from"] = t
		}
	}

	if v, ok := d.GetOk("logdisk_size"); ok {
		t, err := expandDvmdbDeviceLogdiskSize(d, v, "logdisk_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logdisk_size"] = t
		}
	}

	if v, ok := d.GetOk("longitude"); ok {
		t, err := expandDvmdbDeviceLongitude(d, v, "longitude")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["longitude"] = t
		}
	}

	if v, ok := d.GetOk("maxvdom"); ok {
		t, err := expandDvmdbDeviceMaxvdom(d, v, "maxvdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maxvdom"] = t
		}
	}

	if v, ok := d.GetOk("metafields"); ok {
		t, err := expandDvmdbDeviceMetaFields(d, v, "metafields")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["meta fields"] = t
		}
	}

	if v, ok := d.GetOk("mgmt_id"); ok {
		t, err := expandDvmdbDeviceMgmtId(d, v, "mgmt_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mgmt_id"] = t
		}
	}

	if v, ok := d.GetOk("mgmt_if"); ok {
		t, err := expandDvmdbDeviceMgmtIf(d, v, "mgmt_if")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mgmt_if"] = t
		}
	}

	if v, ok := d.GetOk("mgmt_mode"); ok {
		t, err := expandDvmdbDeviceMgmtMode(d, v, "mgmt_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mgmt_mode"] = t
		}
	}

	if v, ok := d.GetOk("mgt_vdom"); ok {
		t, err := expandDvmdbDeviceMgtVdom(d, v, "mgt_vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mgt_vdom"] = t
		}
	}

	if v, ok := d.GetOk("module_sn"); ok {
		t, err := expandDvmdbDeviceModuleSn(d, v, "module_sn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["module_sn"] = t
		}
	}

	if v, ok := d.GetOk("mr"); ok {
		t, err := expandDvmdbDeviceMr(d, v, "mr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mr"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandDvmdbDeviceName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("nsxt_service_name"); ok {
		t, err := expandDvmdbDeviceNsxtServiceName(d, v, "nsxt_service_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nsxt_service_name"] = t
		}
	}

	if v, ok := d.GetOk("os_type"); ok {
		t, err := expandDvmdbDeviceOsType(d, v, "os_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["os_type"] = t
		}
	}

	if v, ok := d.GetOk("os_ver"); ok {
		t, err := expandDvmdbDeviceOsVer(d, v, "os_ver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["os_ver"] = t
		}
	}

	if v, ok := d.GetOk("patch"); ok {
		t, err := expandDvmdbDevicePatch(d, v, "patch")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["patch"] = t
		}
	}

	if v, ok := d.GetOk("platform_str"); ok {
		t, err := expandDvmdbDevicePlatformStr(d, v, "platform_str")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["platform_str"] = t
		}
	}

	if v, ok := d.GetOk("prefer_img_ver"); ok {
		t, err := expandDvmdbDevicePreferImgVer(d, v, "prefer_img_ver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefer_img_ver"] = t
		}
	}

	if v, ok := d.GetOk("prio"); ok {
		t, err := expandDvmdbDevicePrio(d, v, "prio")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prio"] = t
		}
	}

	if v, ok := d.GetOk("private_key"); ok {
		t, err := expandDvmdbDevicePrivateKey(d, v, "private_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["private_key"] = t
		}
	}

	if v, ok := d.GetOk("private_key_status"); ok {
		t, err := expandDvmdbDevicePrivateKeyStatus(d, v, "private_key_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["private_key_status"] = t
		}
	}

	if v, ok := d.GetOk("psk"); ok {
		t, err := expandDvmdbDevicePsk(d, v, "psk")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["psk"] = t
		}
	}

	if v, ok := d.GetOk("role"); ok {
		t, err := expandDvmdbDeviceRole(d, v, "role")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["role"] = t
		}
	}

	if v, ok := d.GetOk("sn"); ok {
		t, err := expandDvmdbDeviceSn(d, v, "sn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sn"] = t
		}
	}

	if v, ok := d.GetOk("vdom"); ok {
		t, err := expandDvmdbDeviceVdom(d, v, "vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	if v, ok := d.GetOk("version"); ok {
		t, err := expandDvmdbDeviceVersion(d, v, "version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["version"] = t
		}
	}

	if v, ok := d.GetOk("vm_cpu"); ok {
		t, err := expandDvmdbDeviceVmCpu(d, v, "vm_cpu")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vm_cpu"] = t
		}
	}

	if v, ok := d.GetOk("vm_cpu_limit"); ok {
		t, err := expandDvmdbDeviceVmCpuLimit(d, v, "vm_cpu_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vm_cpu_limit"] = t
		}
	}

	if v, ok := d.GetOk("vm_lic_expire"); ok {
		t, err := expandDvmdbDeviceVmLicExpire(d, v, "vm_lic_expire")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vm_lic_expire"] = t
		}
	}

	if v, ok := d.GetOk("vm_mem"); ok {
		t, err := expandDvmdbDeviceVmMem(d, v, "vm_mem")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vm_mem"] = t
		}
	}

	if v, ok := d.GetOk("vm_mem_limit"); ok {
		t, err := expandDvmdbDeviceVmMemLimit(d, v, "vm_mem_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vm_mem_limit"] = t
		}
	}

	if v, ok := d.GetOk("vm_status"); ok {
		t, err := expandDvmdbDeviceVmStatus(d, v, "vm_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vm_status"] = t
		}
	}

	return &obj, nil
}
