// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Filter for syslog logging.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemLocallogSyslogd3Filter() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLocallogSyslogd3FilterUpdate,
		Read:   resourceSystemLocallogSyslogd3FilterRead,
		Update: resourceSystemLocallogSyslogd3FilterUpdate,
		Delete: resourceSystemLocallogSyslogd3FilterDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"aid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"devcfg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"devops": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diskquota": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"docker": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dvm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ediscovery": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"epmgr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"event": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eventmgmt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"faz": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fazha": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fazsys": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fgd": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fgfm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fips": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fmgws": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fmlmgr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fmwmgr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortiview": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"glbcfg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ha": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hcache": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"incident": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"iolog": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logd": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logdb": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logdev": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logfile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logging": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lrmgr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"objcfg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"report": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rev": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rtmon": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scfw": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scply": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scrmgr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scvpn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"system": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"webport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemLocallogSyslogd3FilterUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemLocallogSyslogd3Filter(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemLocallogSyslogd3Filter resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemLocallogSyslogd3Filter(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemLocallogSyslogd3Filter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemLocallogSyslogd3Filter")

	return resourceSystemLocallogSyslogd3FilterRead(d, m)
}

func resourceSystemLocallogSyslogd3FilterDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemLocallogSyslogd3Filter(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemLocallogSyslogd3Filter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLocallogSyslogd3FilterRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemLocallogSyslogd3Filter(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemLocallogSyslogd3Filter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemLocallogSyslogd3Filter(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemLocallogSyslogd3Filter resource from API: %v", err)
	}
	return nil
}

func flattenSystemLocallogSyslogd3FilterAid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0:   "disable",
			256: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemLocallogSyslogd3FilterDevcfg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			8: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemLocallogSyslogd3FilterDevops(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0:         "disable",
			134217728: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemLocallogSyslogd3FilterDiskquota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0:  "disable",
			16: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemLocallogSyslogd3FilterDm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0:    "disable",
			4096: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemLocallogSyslogd3FilterDocker(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0:    "disable",
			1024: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemLocallogSyslogd3FilterDvm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0:       "disable",
			4194304: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemLocallogSyslogd3FilterEdiscovery(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0:  "disable",
			64: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemLocallogSyslogd3FilterEpmgr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0:    "disable",
			1024: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemLocallogSyslogd3FilterEvent(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenSystemLocallogSyslogd3FilterEventmgmt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			2: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemLocallogSyslogd3FilterFaz(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0:        "disable",
			16777216: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemLocallogSyslogd3FilterFazha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0:   "disable",
			128: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemLocallogSyslogd3FilterFazsys(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0:         "disable",
			268435456: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemLocallogSyslogd3FilterFgd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0:      "disable",
			131072: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemLocallogSyslogd3FilterFgfm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			4: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemLocallogSyslogd3FilterFips(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0:        "disable",
			67108864: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemLocallogSyslogd3FilterFmgws(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0:       "disable",
			8388608: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemLocallogSyslogd3FilterFmlmgr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0:      "disable",
			524288: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemLocallogSyslogd3FilterFmwmgr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0:     "disable",
			65536: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemLocallogSyslogd3FilterFortiview(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0:  "disable",
			32: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemLocallogSyslogd3FilterGlbcfg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0:  "disable",
			16: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemLocallogSyslogd3FilterHa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0:     "disable",
			32768: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemLocallogSyslogd3FilterHcache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			8: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemLocallogSyslogd3FilterIncident(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0:   "disable",
			512: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemLocallogSyslogd3FilterIolog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0:       "disable",
			1048576: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemLocallogSyslogd3FilterLogd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0:        "disable",
			33554432: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemLocallogSyslogd3FilterLogdb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			4: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemLocallogSyslogd3FilterLogdev(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0:         "disable",
			536870912: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemLocallogSyslogd3FilterLogfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			-2147483648: "enable",
			0:           "disable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemLocallogSyslogd3FilterLogging(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0:          "disable",
			1073741824: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemLocallogSyslogd3FilterLrmgr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0:     "disable",
			16384: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemLocallogSyslogd3FilterObjcfg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0:       "disable",
			2097152: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemLocallogSyslogd3FilterReport(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenSystemLocallogSyslogd3FilterRev(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0:    "disable",
			2048: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemLocallogSyslogd3FilterRtmon(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0:    "disable",
			8192: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemLocallogSyslogd3FilterScfw(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0:   "disable",
			128: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemLocallogSyslogd3FilterScply(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0:   "disable",
			256: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemLocallogSyslogd3FilterScrmgr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0:  "disable",
			32: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemLocallogSyslogd3FilterScvpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0:   "disable",
			512: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemLocallogSyslogd3FilterSystem(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "disable",
			2: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenSystemLocallogSyslogd3FilterWebport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0:  "disable",
			64: "enable",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func refreshObjectSystemLocallogSyslogd3Filter(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("aid", flattenSystemLocallogSyslogd3FilterAid(o["aid"], d, "aid")); err != nil {
		if vv, ok := fortiAPIPatch(o["aid"], "SystemLocallogSyslogd3Filter-Aid"); ok {
			if err = d.Set("aid", vv); err != nil {
				return fmt.Errorf("Error reading aid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading aid: %v", err)
		}
	}

	if err = d.Set("devcfg", flattenSystemLocallogSyslogd3FilterDevcfg(o["devcfg"], d, "devcfg")); err != nil {
		if vv, ok := fortiAPIPatch(o["devcfg"], "SystemLocallogSyslogd3Filter-Devcfg"); ok {
			if err = d.Set("devcfg", vv); err != nil {
				return fmt.Errorf("Error reading devcfg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading devcfg: %v", err)
		}
	}

	if err = d.Set("devops", flattenSystemLocallogSyslogd3FilterDevops(o["devops"], d, "devops")); err != nil {
		if vv, ok := fortiAPIPatch(o["devops"], "SystemLocallogSyslogd3Filter-Devops"); ok {
			if err = d.Set("devops", vv); err != nil {
				return fmt.Errorf("Error reading devops: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading devops: %v", err)
		}
	}

	if err = d.Set("diskquota", flattenSystemLocallogSyslogd3FilterDiskquota(o["diskquota"], d, "diskquota")); err != nil {
		if vv, ok := fortiAPIPatch(o["diskquota"], "SystemLocallogSyslogd3Filter-Diskquota"); ok {
			if err = d.Set("diskquota", vv); err != nil {
				return fmt.Errorf("Error reading diskquota: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diskquota: %v", err)
		}
	}

	if err = d.Set("dm", flattenSystemLocallogSyslogd3FilterDm(o["dm"], d, "dm")); err != nil {
		if vv, ok := fortiAPIPatch(o["dm"], "SystemLocallogSyslogd3Filter-Dm"); ok {
			if err = d.Set("dm", vv); err != nil {
				return fmt.Errorf("Error reading dm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dm: %v", err)
		}
	}

	if err = d.Set("docker", flattenSystemLocallogSyslogd3FilterDocker(o["docker"], d, "docker")); err != nil {
		if vv, ok := fortiAPIPatch(o["docker"], "SystemLocallogSyslogd3Filter-Docker"); ok {
			if err = d.Set("docker", vv); err != nil {
				return fmt.Errorf("Error reading docker: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading docker: %v", err)
		}
	}

	if err = d.Set("dvm", flattenSystemLocallogSyslogd3FilterDvm(o["dvm"], d, "dvm")); err != nil {
		if vv, ok := fortiAPIPatch(o["dvm"], "SystemLocallogSyslogd3Filter-Dvm"); ok {
			if err = d.Set("dvm", vv); err != nil {
				return fmt.Errorf("Error reading dvm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dvm: %v", err)
		}
	}

	if err = d.Set("ediscovery", flattenSystemLocallogSyslogd3FilterEdiscovery(o["ediscovery"], d, "ediscovery")); err != nil {
		if vv, ok := fortiAPIPatch(o["ediscovery"], "SystemLocallogSyslogd3Filter-Ediscovery"); ok {
			if err = d.Set("ediscovery", vv); err != nil {
				return fmt.Errorf("Error reading ediscovery: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ediscovery: %v", err)
		}
	}

	if err = d.Set("epmgr", flattenSystemLocallogSyslogd3FilterEpmgr(o["epmgr"], d, "epmgr")); err != nil {
		if vv, ok := fortiAPIPatch(o["epmgr"], "SystemLocallogSyslogd3Filter-Epmgr"); ok {
			if err = d.Set("epmgr", vv); err != nil {
				return fmt.Errorf("Error reading epmgr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading epmgr: %v", err)
		}
	}

	if err = d.Set("event", flattenSystemLocallogSyslogd3FilterEvent(o["event"], d, "event")); err != nil {
		if vv, ok := fortiAPIPatch(o["event"], "SystemLocallogSyslogd3Filter-Event"); ok {
			if err = d.Set("event", vv); err != nil {
				return fmt.Errorf("Error reading event: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading event: %v", err)
		}
	}

	if err = d.Set("eventmgmt", flattenSystemLocallogSyslogd3FilterEventmgmt(o["eventmgmt"], d, "eventmgmt")); err != nil {
		if vv, ok := fortiAPIPatch(o["eventmgmt"], "SystemLocallogSyslogd3Filter-Eventmgmt"); ok {
			if err = d.Set("eventmgmt", vv); err != nil {
				return fmt.Errorf("Error reading eventmgmt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eventmgmt: %v", err)
		}
	}

	if err = d.Set("faz", flattenSystemLocallogSyslogd3FilterFaz(o["faz"], d, "faz")); err != nil {
		if vv, ok := fortiAPIPatch(o["faz"], "SystemLocallogSyslogd3Filter-Faz"); ok {
			if err = d.Set("faz", vv); err != nil {
				return fmt.Errorf("Error reading faz: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading faz: %v", err)
		}
	}

	if err = d.Set("fazha", flattenSystemLocallogSyslogd3FilterFazha(o["fazha"], d, "fazha")); err != nil {
		if vv, ok := fortiAPIPatch(o["fazha"], "SystemLocallogSyslogd3Filter-Fazha"); ok {
			if err = d.Set("fazha", vv); err != nil {
				return fmt.Errorf("Error reading fazha: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fazha: %v", err)
		}
	}

	if err = d.Set("fazsys", flattenSystemLocallogSyslogd3FilterFazsys(o["fazsys"], d, "fazsys")); err != nil {
		if vv, ok := fortiAPIPatch(o["fazsys"], "SystemLocallogSyslogd3Filter-Fazsys"); ok {
			if err = d.Set("fazsys", vv); err != nil {
				return fmt.Errorf("Error reading fazsys: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fazsys: %v", err)
		}
	}

	if err = d.Set("fgd", flattenSystemLocallogSyslogd3FilterFgd(o["fgd"], d, "fgd")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgd"], "SystemLocallogSyslogd3Filter-Fgd"); ok {
			if err = d.Set("fgd", vv); err != nil {
				return fmt.Errorf("Error reading fgd: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgd: %v", err)
		}
	}

	if err = d.Set("fgfm", flattenSystemLocallogSyslogd3FilterFgfm(o["fgfm"], d, "fgfm")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgfm"], "SystemLocallogSyslogd3Filter-Fgfm"); ok {
			if err = d.Set("fgfm", vv); err != nil {
				return fmt.Errorf("Error reading fgfm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgfm: %v", err)
		}
	}

	if err = d.Set("fips", flattenSystemLocallogSyslogd3FilterFips(o["fips"], d, "fips")); err != nil {
		if vv, ok := fortiAPIPatch(o["fips"], "SystemLocallogSyslogd3Filter-Fips"); ok {
			if err = d.Set("fips", vv); err != nil {
				return fmt.Errorf("Error reading fips: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fips: %v", err)
		}
	}

	if err = d.Set("fmgws", flattenSystemLocallogSyslogd3FilterFmgws(o["fmgws"], d, "fmgws")); err != nil {
		if vv, ok := fortiAPIPatch(o["fmgws"], "SystemLocallogSyslogd3Filter-Fmgws"); ok {
			if err = d.Set("fmgws", vv); err != nil {
				return fmt.Errorf("Error reading fmgws: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmgws: %v", err)
		}
	}

	if err = d.Set("fmlmgr", flattenSystemLocallogSyslogd3FilterFmlmgr(o["fmlmgr"], d, "fmlmgr")); err != nil {
		if vv, ok := fortiAPIPatch(o["fmlmgr"], "SystemLocallogSyslogd3Filter-Fmlmgr"); ok {
			if err = d.Set("fmlmgr", vv); err != nil {
				return fmt.Errorf("Error reading fmlmgr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmlmgr: %v", err)
		}
	}

	if err = d.Set("fmwmgr", flattenSystemLocallogSyslogd3FilterFmwmgr(o["fmwmgr"], d, "fmwmgr")); err != nil {
		if vv, ok := fortiAPIPatch(o["fmwmgr"], "SystemLocallogSyslogd3Filter-Fmwmgr"); ok {
			if err = d.Set("fmwmgr", vv); err != nil {
				return fmt.Errorf("Error reading fmwmgr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmwmgr: %v", err)
		}
	}

	if err = d.Set("fortiview", flattenSystemLocallogSyslogd3FilterFortiview(o["fortiview"], d, "fortiview")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortiview"], "SystemLocallogSyslogd3Filter-Fortiview"); ok {
			if err = d.Set("fortiview", vv); err != nil {
				return fmt.Errorf("Error reading fortiview: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortiview: %v", err)
		}
	}

	if err = d.Set("glbcfg", flattenSystemLocallogSyslogd3FilterGlbcfg(o["glbcfg"], d, "glbcfg")); err != nil {
		if vv, ok := fortiAPIPatch(o["glbcfg"], "SystemLocallogSyslogd3Filter-Glbcfg"); ok {
			if err = d.Set("glbcfg", vv); err != nil {
				return fmt.Errorf("Error reading glbcfg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading glbcfg: %v", err)
		}
	}

	if err = d.Set("ha", flattenSystemLocallogSyslogd3FilterHa(o["ha"], d, "ha")); err != nil {
		if vv, ok := fortiAPIPatch(o["ha"], "SystemLocallogSyslogd3Filter-Ha"); ok {
			if err = d.Set("ha", vv); err != nil {
				return fmt.Errorf("Error reading ha: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ha: %v", err)
		}
	}

	if err = d.Set("hcache", flattenSystemLocallogSyslogd3FilterHcache(o["hcache"], d, "hcache")); err != nil {
		if vv, ok := fortiAPIPatch(o["hcache"], "SystemLocallogSyslogd3Filter-Hcache"); ok {
			if err = d.Set("hcache", vv); err != nil {
				return fmt.Errorf("Error reading hcache: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hcache: %v", err)
		}
	}

	if err = d.Set("incident", flattenSystemLocallogSyslogd3FilterIncident(o["incident"], d, "incident")); err != nil {
		if vv, ok := fortiAPIPatch(o["incident"], "SystemLocallogSyslogd3Filter-Incident"); ok {
			if err = d.Set("incident", vv); err != nil {
				return fmt.Errorf("Error reading incident: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading incident: %v", err)
		}
	}

	if err = d.Set("iolog", flattenSystemLocallogSyslogd3FilterIolog(o["iolog"], d, "iolog")); err != nil {
		if vv, ok := fortiAPIPatch(o["iolog"], "SystemLocallogSyslogd3Filter-Iolog"); ok {
			if err = d.Set("iolog", vv); err != nil {
				return fmt.Errorf("Error reading iolog: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading iolog: %v", err)
		}
	}

	if err = d.Set("logd", flattenSystemLocallogSyslogd3FilterLogd(o["logd"], d, "logd")); err != nil {
		if vv, ok := fortiAPIPatch(o["logd"], "SystemLocallogSyslogd3Filter-Logd"); ok {
			if err = d.Set("logd", vv); err != nil {
				return fmt.Errorf("Error reading logd: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logd: %v", err)
		}
	}

	if err = d.Set("logdb", flattenSystemLocallogSyslogd3FilterLogdb(o["logdb"], d, "logdb")); err != nil {
		if vv, ok := fortiAPIPatch(o["logdb"], "SystemLocallogSyslogd3Filter-Logdb"); ok {
			if err = d.Set("logdb", vv); err != nil {
				return fmt.Errorf("Error reading logdb: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logdb: %v", err)
		}
	}

	if err = d.Set("logdev", flattenSystemLocallogSyslogd3FilterLogdev(o["logdev"], d, "logdev")); err != nil {
		if vv, ok := fortiAPIPatch(o["logdev"], "SystemLocallogSyslogd3Filter-Logdev"); ok {
			if err = d.Set("logdev", vv); err != nil {
				return fmt.Errorf("Error reading logdev: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logdev: %v", err)
		}
	}

	if err = d.Set("logfile", flattenSystemLocallogSyslogd3FilterLogfile(o["logfile"], d, "logfile")); err != nil {
		if vv, ok := fortiAPIPatch(o["logfile"], "SystemLocallogSyslogd3Filter-Logfile"); ok {
			if err = d.Set("logfile", vv); err != nil {
				return fmt.Errorf("Error reading logfile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logfile: %v", err)
		}
	}

	if err = d.Set("logging", flattenSystemLocallogSyslogd3FilterLogging(o["logging"], d, "logging")); err != nil {
		if vv, ok := fortiAPIPatch(o["logging"], "SystemLocallogSyslogd3Filter-Logging"); ok {
			if err = d.Set("logging", vv); err != nil {
				return fmt.Errorf("Error reading logging: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logging: %v", err)
		}
	}

	if err = d.Set("lrmgr", flattenSystemLocallogSyslogd3FilterLrmgr(o["lrmgr"], d, "lrmgr")); err != nil {
		if vv, ok := fortiAPIPatch(o["lrmgr"], "SystemLocallogSyslogd3Filter-Lrmgr"); ok {
			if err = d.Set("lrmgr", vv); err != nil {
				return fmt.Errorf("Error reading lrmgr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lrmgr: %v", err)
		}
	}

	if err = d.Set("objcfg", flattenSystemLocallogSyslogd3FilterObjcfg(o["objcfg"], d, "objcfg")); err != nil {
		if vv, ok := fortiAPIPatch(o["objcfg"], "SystemLocallogSyslogd3Filter-Objcfg"); ok {
			if err = d.Set("objcfg", vv); err != nil {
				return fmt.Errorf("Error reading objcfg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading objcfg: %v", err)
		}
	}

	if err = d.Set("report", flattenSystemLocallogSyslogd3FilterReport(o["report"], d, "report")); err != nil {
		if vv, ok := fortiAPIPatch(o["report"], "SystemLocallogSyslogd3Filter-Report"); ok {
			if err = d.Set("report", vv); err != nil {
				return fmt.Errorf("Error reading report: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading report: %v", err)
		}
	}

	if err = d.Set("rev", flattenSystemLocallogSyslogd3FilterRev(o["rev"], d, "rev")); err != nil {
		if vv, ok := fortiAPIPatch(o["rev"], "SystemLocallogSyslogd3Filter-Rev"); ok {
			if err = d.Set("rev", vv); err != nil {
				return fmt.Errorf("Error reading rev: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rev: %v", err)
		}
	}

	if err = d.Set("rtmon", flattenSystemLocallogSyslogd3FilterRtmon(o["rtmon"], d, "rtmon")); err != nil {
		if vv, ok := fortiAPIPatch(o["rtmon"], "SystemLocallogSyslogd3Filter-Rtmon"); ok {
			if err = d.Set("rtmon", vv); err != nil {
				return fmt.Errorf("Error reading rtmon: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rtmon: %v", err)
		}
	}

	if err = d.Set("scfw", flattenSystemLocallogSyslogd3FilterScfw(o["scfw"], d, "scfw")); err != nil {
		if vv, ok := fortiAPIPatch(o["scfw"], "SystemLocallogSyslogd3Filter-Scfw"); ok {
			if err = d.Set("scfw", vv); err != nil {
				return fmt.Errorf("Error reading scfw: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scfw: %v", err)
		}
	}

	if err = d.Set("scply", flattenSystemLocallogSyslogd3FilterScply(o["scply"], d, "scply")); err != nil {
		if vv, ok := fortiAPIPatch(o["scply"], "SystemLocallogSyslogd3Filter-Scply"); ok {
			if err = d.Set("scply", vv); err != nil {
				return fmt.Errorf("Error reading scply: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scply: %v", err)
		}
	}

	if err = d.Set("scrmgr", flattenSystemLocallogSyslogd3FilterScrmgr(o["scrmgr"], d, "scrmgr")); err != nil {
		if vv, ok := fortiAPIPatch(o["scrmgr"], "SystemLocallogSyslogd3Filter-Scrmgr"); ok {
			if err = d.Set("scrmgr", vv); err != nil {
				return fmt.Errorf("Error reading scrmgr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scrmgr: %v", err)
		}
	}

	if err = d.Set("scvpn", flattenSystemLocallogSyslogd3FilterScvpn(o["scvpn"], d, "scvpn")); err != nil {
		if vv, ok := fortiAPIPatch(o["scvpn"], "SystemLocallogSyslogd3Filter-Scvpn"); ok {
			if err = d.Set("scvpn", vv); err != nil {
				return fmt.Errorf("Error reading scvpn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scvpn: %v", err)
		}
	}

	if err = d.Set("system", flattenSystemLocallogSyslogd3FilterSystem(o["system"], d, "system")); err != nil {
		if vv, ok := fortiAPIPatch(o["system"], "SystemLocallogSyslogd3Filter-System"); ok {
			if err = d.Set("system", vv); err != nil {
				return fmt.Errorf("Error reading system: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading system: %v", err)
		}
	}

	if err = d.Set("webport", flattenSystemLocallogSyslogd3FilterWebport(o["webport"], d, "webport")); err != nil {
		if vv, ok := fortiAPIPatch(o["webport"], "SystemLocallogSyslogd3Filter-Webport"); ok {
			if err = d.Set("webport", vv); err != nil {
				return fmt.Errorf("Error reading webport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webport: %v", err)
		}
	}

	return nil
}

func flattenSystemLocallogSyslogd3FilterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemLocallogSyslogd3FilterAid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd3FilterDevcfg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd3FilterDevops(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd3FilterDiskquota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd3FilterDm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd3FilterDocker(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd3FilterDvm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd3FilterEdiscovery(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd3FilterEpmgr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd3FilterEvent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd3FilterEventmgmt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd3FilterFaz(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd3FilterFazha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd3FilterFazsys(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd3FilterFgd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd3FilterFgfm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd3FilterFips(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd3FilterFmgws(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd3FilterFmlmgr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd3FilterFmwmgr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd3FilterFortiview(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd3FilterGlbcfg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd3FilterHa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd3FilterHcache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd3FilterIncident(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd3FilterIolog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd3FilterLogd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd3FilterLogdb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd3FilterLogdev(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd3FilterLogfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd3FilterLogging(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd3FilterLrmgr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd3FilterObjcfg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd3FilterReport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd3FilterRev(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd3FilterRtmon(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd3FilterScfw(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd3FilterScply(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd3FilterScrmgr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd3FilterScvpn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd3FilterSystem(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd3FilterWebport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemLocallogSyslogd3Filter(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("aid"); ok {
		t, err := expandSystemLocallogSyslogd3FilterAid(d, v, "aid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aid"] = t
		}
	}

	if v, ok := d.GetOk("devcfg"); ok {
		t, err := expandSystemLocallogSyslogd3FilterDevcfg(d, v, "devcfg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["devcfg"] = t
		}
	}

	if v, ok := d.GetOk("devops"); ok {
		t, err := expandSystemLocallogSyslogd3FilterDevops(d, v, "devops")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["devops"] = t
		}
	}

	if v, ok := d.GetOk("diskquota"); ok {
		t, err := expandSystemLocallogSyslogd3FilterDiskquota(d, v, "diskquota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diskquota"] = t
		}
	}

	if v, ok := d.GetOk("dm"); ok {
		t, err := expandSystemLocallogSyslogd3FilterDm(d, v, "dm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dm"] = t
		}
	}

	if v, ok := d.GetOk("docker"); ok {
		t, err := expandSystemLocallogSyslogd3FilterDocker(d, v, "docker")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["docker"] = t
		}
	}

	if v, ok := d.GetOk("dvm"); ok {
		t, err := expandSystemLocallogSyslogd3FilterDvm(d, v, "dvm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dvm"] = t
		}
	}

	if v, ok := d.GetOk("ediscovery"); ok {
		t, err := expandSystemLocallogSyslogd3FilterEdiscovery(d, v, "ediscovery")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ediscovery"] = t
		}
	}

	if v, ok := d.GetOk("epmgr"); ok {
		t, err := expandSystemLocallogSyslogd3FilterEpmgr(d, v, "epmgr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["epmgr"] = t
		}
	}

	if v, ok := d.GetOk("event"); ok {
		t, err := expandSystemLocallogSyslogd3FilterEvent(d, v, "event")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["event"] = t
		}
	}

	if v, ok := d.GetOk("eventmgmt"); ok {
		t, err := expandSystemLocallogSyslogd3FilterEventmgmt(d, v, "eventmgmt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eventmgmt"] = t
		}
	}

	if v, ok := d.GetOk("faz"); ok {
		t, err := expandSystemLocallogSyslogd3FilterFaz(d, v, "faz")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["faz"] = t
		}
	}

	if v, ok := d.GetOk("fazha"); ok {
		t, err := expandSystemLocallogSyslogd3FilterFazha(d, v, "fazha")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fazha"] = t
		}
	}

	if v, ok := d.GetOk("fazsys"); ok {
		t, err := expandSystemLocallogSyslogd3FilterFazsys(d, v, "fazsys")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fazsys"] = t
		}
	}

	if v, ok := d.GetOk("fgd"); ok {
		t, err := expandSystemLocallogSyslogd3FilterFgd(d, v, "fgd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgd"] = t
		}
	}

	if v, ok := d.GetOk("fgfm"); ok {
		t, err := expandSystemLocallogSyslogd3FilterFgfm(d, v, "fgfm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgfm"] = t
		}
	}

	if v, ok := d.GetOk("fips"); ok {
		t, err := expandSystemLocallogSyslogd3FilterFips(d, v, "fips")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fips"] = t
		}
	}

	if v, ok := d.GetOk("fmgws"); ok {
		t, err := expandSystemLocallogSyslogd3FilterFmgws(d, v, "fmgws")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fmgws"] = t
		}
	}

	if v, ok := d.GetOk("fmlmgr"); ok {
		t, err := expandSystemLocallogSyslogd3FilterFmlmgr(d, v, "fmlmgr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fmlmgr"] = t
		}
	}

	if v, ok := d.GetOk("fmwmgr"); ok {
		t, err := expandSystemLocallogSyslogd3FilterFmwmgr(d, v, "fmwmgr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fmwmgr"] = t
		}
	}

	if v, ok := d.GetOk("fortiview"); ok {
		t, err := expandSystemLocallogSyslogd3FilterFortiview(d, v, "fortiview")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiview"] = t
		}
	}

	if v, ok := d.GetOk("glbcfg"); ok {
		t, err := expandSystemLocallogSyslogd3FilterGlbcfg(d, v, "glbcfg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["glbcfg"] = t
		}
	}

	if v, ok := d.GetOk("ha"); ok {
		t, err := expandSystemLocallogSyslogd3FilterHa(d, v, "ha")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha"] = t
		}
	}

	if v, ok := d.GetOk("hcache"); ok {
		t, err := expandSystemLocallogSyslogd3FilterHcache(d, v, "hcache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hcache"] = t
		}
	}

	if v, ok := d.GetOk("incident"); ok {
		t, err := expandSystemLocallogSyslogd3FilterIncident(d, v, "incident")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["incident"] = t
		}
	}

	if v, ok := d.GetOk("iolog"); ok {
		t, err := expandSystemLocallogSyslogd3FilterIolog(d, v, "iolog")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["iolog"] = t
		}
	}

	if v, ok := d.GetOk("logd"); ok {
		t, err := expandSystemLocallogSyslogd3FilterLogd(d, v, "logd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logd"] = t
		}
	}

	if v, ok := d.GetOk("logdb"); ok {
		t, err := expandSystemLocallogSyslogd3FilterLogdb(d, v, "logdb")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logdb"] = t
		}
	}

	if v, ok := d.GetOk("logdev"); ok {
		t, err := expandSystemLocallogSyslogd3FilterLogdev(d, v, "logdev")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logdev"] = t
		}
	}

	if v, ok := d.GetOk("logfile"); ok {
		t, err := expandSystemLocallogSyslogd3FilterLogfile(d, v, "logfile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logfile"] = t
		}
	}

	if v, ok := d.GetOk("logging"); ok {
		t, err := expandSystemLocallogSyslogd3FilterLogging(d, v, "logging")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logging"] = t
		}
	}

	if v, ok := d.GetOk("lrmgr"); ok {
		t, err := expandSystemLocallogSyslogd3FilterLrmgr(d, v, "lrmgr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lrmgr"] = t
		}
	}

	if v, ok := d.GetOk("objcfg"); ok {
		t, err := expandSystemLocallogSyslogd3FilterObjcfg(d, v, "objcfg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["objcfg"] = t
		}
	}

	if v, ok := d.GetOk("report"); ok {
		t, err := expandSystemLocallogSyslogd3FilterReport(d, v, "report")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["report"] = t
		}
	}

	if v, ok := d.GetOk("rev"); ok {
		t, err := expandSystemLocallogSyslogd3FilterRev(d, v, "rev")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rev"] = t
		}
	}

	if v, ok := d.GetOk("rtmon"); ok {
		t, err := expandSystemLocallogSyslogd3FilterRtmon(d, v, "rtmon")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rtmon"] = t
		}
	}

	if v, ok := d.GetOk("scfw"); ok {
		t, err := expandSystemLocallogSyslogd3FilterScfw(d, v, "scfw")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scfw"] = t
		}
	}

	if v, ok := d.GetOk("scply"); ok {
		t, err := expandSystemLocallogSyslogd3FilterScply(d, v, "scply")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scply"] = t
		}
	}

	if v, ok := d.GetOk("scrmgr"); ok {
		t, err := expandSystemLocallogSyslogd3FilterScrmgr(d, v, "scrmgr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scrmgr"] = t
		}
	}

	if v, ok := d.GetOk("scvpn"); ok {
		t, err := expandSystemLocallogSyslogd3FilterScvpn(d, v, "scvpn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scvpn"] = t
		}
	}

	if v, ok := d.GetOk("system"); ok {
		t, err := expandSystemLocallogSyslogd3FilterSystem(d, v, "system")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["system"] = t
		}
	}

	if v, ok := d.GetOk("webport"); ok {
		t, err := expandSystemLocallogSyslogd3FilterWebport(d, v, "webport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webport"] = t
		}
	}

	return &obj, nil
}
