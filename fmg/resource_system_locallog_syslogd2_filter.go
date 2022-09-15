// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Filter for syslog logging.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemLocallogSyslogd2Filter() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLocallogSyslogd2FilterUpdate,
		Read:   resourceSystemLocallogSyslogd2FilterRead,
		Update: resourceSystemLocallogSyslogd2FilterUpdate,
		Delete: resourceSystemLocallogSyslogd2FilterDelete,

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

func resourceSystemLocallogSyslogd2FilterUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemLocallogSyslogd2Filter(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemLocallogSyslogd2Filter resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemLocallogSyslogd2Filter(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemLocallogSyslogd2Filter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemLocallogSyslogd2Filter")

	return resourceSystemLocallogSyslogd2FilterRead(d, m)
}

func resourceSystemLocallogSyslogd2FilterDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemLocallogSyslogd2Filter(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemLocallogSyslogd2Filter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLocallogSyslogd2FilterRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemLocallogSyslogd2Filter(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemLocallogSyslogd2Filter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemLocallogSyslogd2Filter(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemLocallogSyslogd2Filter resource from API: %v", err)
	}
	return nil
}

func flattenSystemLocallogSyslogd2FilterAid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogd2FilterDevcfg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogd2FilterDevops(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogd2FilterDiskquota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogd2FilterDm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogd2FilterDocker(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogd2FilterDvm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogd2FilterEdiscovery(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogd2FilterEpmgr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogd2FilterEvent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogd2FilterEventmgmt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogd2FilterFaz(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogd2FilterFazha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogd2FilterFazsys(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogd2FilterFgd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogd2FilterFgfm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogd2FilterFips(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogd2FilterFmgws(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogd2FilterFmlmgr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogd2FilterFmwmgr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogd2FilterFortiview(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogd2FilterGlbcfg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogd2FilterHa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogd2FilterHcache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogd2FilterIncident(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogd2FilterIolog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogd2FilterLogd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogd2FilterLogdb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogd2FilterLogdev(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogd2FilterLogfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogd2FilterLogging(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogd2FilterLrmgr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogd2FilterObjcfg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogd2FilterReport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogd2FilterRev(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogd2FilterRtmon(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogd2FilterScfw(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogd2FilterScply(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogd2FilterScrmgr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogd2FilterScvpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogd2FilterSystem(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogSyslogd2FilterWebport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemLocallogSyslogd2Filter(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("aid", flattenSystemLocallogSyslogd2FilterAid(o["aid"], d, "aid")); err != nil {
		if vv, ok := fortiAPIPatch(o["aid"], "SystemLocallogSyslogd2Filter-Aid"); ok {
			if err = d.Set("aid", vv); err != nil {
				return fmt.Errorf("Error reading aid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading aid: %v", err)
		}
	}

	if err = d.Set("devcfg", flattenSystemLocallogSyslogd2FilterDevcfg(o["devcfg"], d, "devcfg")); err != nil {
		if vv, ok := fortiAPIPatch(o["devcfg"], "SystemLocallogSyslogd2Filter-Devcfg"); ok {
			if err = d.Set("devcfg", vv); err != nil {
				return fmt.Errorf("Error reading devcfg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading devcfg: %v", err)
		}
	}

	if err = d.Set("devops", flattenSystemLocallogSyslogd2FilterDevops(o["devops"], d, "devops")); err != nil {
		if vv, ok := fortiAPIPatch(o["devops"], "SystemLocallogSyslogd2Filter-Devops"); ok {
			if err = d.Set("devops", vv); err != nil {
				return fmt.Errorf("Error reading devops: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading devops: %v", err)
		}
	}

	if err = d.Set("diskquota", flattenSystemLocallogSyslogd2FilterDiskquota(o["diskquota"], d, "diskquota")); err != nil {
		if vv, ok := fortiAPIPatch(o["diskquota"], "SystemLocallogSyslogd2Filter-Diskquota"); ok {
			if err = d.Set("diskquota", vv); err != nil {
				return fmt.Errorf("Error reading diskquota: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diskquota: %v", err)
		}
	}

	if err = d.Set("dm", flattenSystemLocallogSyslogd2FilterDm(o["dm"], d, "dm")); err != nil {
		if vv, ok := fortiAPIPatch(o["dm"], "SystemLocallogSyslogd2Filter-Dm"); ok {
			if err = d.Set("dm", vv); err != nil {
				return fmt.Errorf("Error reading dm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dm: %v", err)
		}
	}

	if err = d.Set("docker", flattenSystemLocallogSyslogd2FilterDocker(o["docker"], d, "docker")); err != nil {
		if vv, ok := fortiAPIPatch(o["docker"], "SystemLocallogSyslogd2Filter-Docker"); ok {
			if err = d.Set("docker", vv); err != nil {
				return fmt.Errorf("Error reading docker: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading docker: %v", err)
		}
	}

	if err = d.Set("dvm", flattenSystemLocallogSyslogd2FilterDvm(o["dvm"], d, "dvm")); err != nil {
		if vv, ok := fortiAPIPatch(o["dvm"], "SystemLocallogSyslogd2Filter-Dvm"); ok {
			if err = d.Set("dvm", vv); err != nil {
				return fmt.Errorf("Error reading dvm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dvm: %v", err)
		}
	}

	if err = d.Set("ediscovery", flattenSystemLocallogSyslogd2FilterEdiscovery(o["ediscovery"], d, "ediscovery")); err != nil {
		if vv, ok := fortiAPIPatch(o["ediscovery"], "SystemLocallogSyslogd2Filter-Ediscovery"); ok {
			if err = d.Set("ediscovery", vv); err != nil {
				return fmt.Errorf("Error reading ediscovery: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ediscovery: %v", err)
		}
	}

	if err = d.Set("epmgr", flattenSystemLocallogSyslogd2FilterEpmgr(o["epmgr"], d, "epmgr")); err != nil {
		if vv, ok := fortiAPIPatch(o["epmgr"], "SystemLocallogSyslogd2Filter-Epmgr"); ok {
			if err = d.Set("epmgr", vv); err != nil {
				return fmt.Errorf("Error reading epmgr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading epmgr: %v", err)
		}
	}

	if err = d.Set("event", flattenSystemLocallogSyslogd2FilterEvent(o["event"], d, "event")); err != nil {
		if vv, ok := fortiAPIPatch(o["event"], "SystemLocallogSyslogd2Filter-Event"); ok {
			if err = d.Set("event", vv); err != nil {
				return fmt.Errorf("Error reading event: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading event: %v", err)
		}
	}

	if err = d.Set("eventmgmt", flattenSystemLocallogSyslogd2FilterEventmgmt(o["eventmgmt"], d, "eventmgmt")); err != nil {
		if vv, ok := fortiAPIPatch(o["eventmgmt"], "SystemLocallogSyslogd2Filter-Eventmgmt"); ok {
			if err = d.Set("eventmgmt", vv); err != nil {
				return fmt.Errorf("Error reading eventmgmt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eventmgmt: %v", err)
		}
	}

	if err = d.Set("faz", flattenSystemLocallogSyslogd2FilterFaz(o["faz"], d, "faz")); err != nil {
		if vv, ok := fortiAPIPatch(o["faz"], "SystemLocallogSyslogd2Filter-Faz"); ok {
			if err = d.Set("faz", vv); err != nil {
				return fmt.Errorf("Error reading faz: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading faz: %v", err)
		}
	}

	if err = d.Set("fazha", flattenSystemLocallogSyslogd2FilterFazha(o["fazha"], d, "fazha")); err != nil {
		if vv, ok := fortiAPIPatch(o["fazha"], "SystemLocallogSyslogd2Filter-Fazha"); ok {
			if err = d.Set("fazha", vv); err != nil {
				return fmt.Errorf("Error reading fazha: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fazha: %v", err)
		}
	}

	if err = d.Set("fazsys", flattenSystemLocallogSyslogd2FilterFazsys(o["fazsys"], d, "fazsys")); err != nil {
		if vv, ok := fortiAPIPatch(o["fazsys"], "SystemLocallogSyslogd2Filter-Fazsys"); ok {
			if err = d.Set("fazsys", vv); err != nil {
				return fmt.Errorf("Error reading fazsys: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fazsys: %v", err)
		}
	}

	if err = d.Set("fgd", flattenSystemLocallogSyslogd2FilterFgd(o["fgd"], d, "fgd")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgd"], "SystemLocallogSyslogd2Filter-Fgd"); ok {
			if err = d.Set("fgd", vv); err != nil {
				return fmt.Errorf("Error reading fgd: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgd: %v", err)
		}
	}

	if err = d.Set("fgfm", flattenSystemLocallogSyslogd2FilterFgfm(o["fgfm"], d, "fgfm")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgfm"], "SystemLocallogSyslogd2Filter-Fgfm"); ok {
			if err = d.Set("fgfm", vv); err != nil {
				return fmt.Errorf("Error reading fgfm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgfm: %v", err)
		}
	}

	if err = d.Set("fips", flattenSystemLocallogSyslogd2FilterFips(o["fips"], d, "fips")); err != nil {
		if vv, ok := fortiAPIPatch(o["fips"], "SystemLocallogSyslogd2Filter-Fips"); ok {
			if err = d.Set("fips", vv); err != nil {
				return fmt.Errorf("Error reading fips: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fips: %v", err)
		}
	}

	if err = d.Set("fmgws", flattenSystemLocallogSyslogd2FilterFmgws(o["fmgws"], d, "fmgws")); err != nil {
		if vv, ok := fortiAPIPatch(o["fmgws"], "SystemLocallogSyslogd2Filter-Fmgws"); ok {
			if err = d.Set("fmgws", vv); err != nil {
				return fmt.Errorf("Error reading fmgws: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmgws: %v", err)
		}
	}

	if err = d.Set("fmlmgr", flattenSystemLocallogSyslogd2FilterFmlmgr(o["fmlmgr"], d, "fmlmgr")); err != nil {
		if vv, ok := fortiAPIPatch(o["fmlmgr"], "SystemLocallogSyslogd2Filter-Fmlmgr"); ok {
			if err = d.Set("fmlmgr", vv); err != nil {
				return fmt.Errorf("Error reading fmlmgr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmlmgr: %v", err)
		}
	}

	if err = d.Set("fmwmgr", flattenSystemLocallogSyslogd2FilterFmwmgr(o["fmwmgr"], d, "fmwmgr")); err != nil {
		if vv, ok := fortiAPIPatch(o["fmwmgr"], "SystemLocallogSyslogd2Filter-Fmwmgr"); ok {
			if err = d.Set("fmwmgr", vv); err != nil {
				return fmt.Errorf("Error reading fmwmgr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmwmgr: %v", err)
		}
	}

	if err = d.Set("fortiview", flattenSystemLocallogSyslogd2FilterFortiview(o["fortiview"], d, "fortiview")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortiview"], "SystemLocallogSyslogd2Filter-Fortiview"); ok {
			if err = d.Set("fortiview", vv); err != nil {
				return fmt.Errorf("Error reading fortiview: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortiview: %v", err)
		}
	}

	if err = d.Set("glbcfg", flattenSystemLocallogSyslogd2FilterGlbcfg(o["glbcfg"], d, "glbcfg")); err != nil {
		if vv, ok := fortiAPIPatch(o["glbcfg"], "SystemLocallogSyslogd2Filter-Glbcfg"); ok {
			if err = d.Set("glbcfg", vv); err != nil {
				return fmt.Errorf("Error reading glbcfg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading glbcfg: %v", err)
		}
	}

	if err = d.Set("ha", flattenSystemLocallogSyslogd2FilterHa(o["ha"], d, "ha")); err != nil {
		if vv, ok := fortiAPIPatch(o["ha"], "SystemLocallogSyslogd2Filter-Ha"); ok {
			if err = d.Set("ha", vv); err != nil {
				return fmt.Errorf("Error reading ha: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ha: %v", err)
		}
	}

	if err = d.Set("hcache", flattenSystemLocallogSyslogd2FilterHcache(o["hcache"], d, "hcache")); err != nil {
		if vv, ok := fortiAPIPatch(o["hcache"], "SystemLocallogSyslogd2Filter-Hcache"); ok {
			if err = d.Set("hcache", vv); err != nil {
				return fmt.Errorf("Error reading hcache: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hcache: %v", err)
		}
	}

	if err = d.Set("incident", flattenSystemLocallogSyslogd2FilterIncident(o["incident"], d, "incident")); err != nil {
		if vv, ok := fortiAPIPatch(o["incident"], "SystemLocallogSyslogd2Filter-Incident"); ok {
			if err = d.Set("incident", vv); err != nil {
				return fmt.Errorf("Error reading incident: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading incident: %v", err)
		}
	}

	if err = d.Set("iolog", flattenSystemLocallogSyslogd2FilterIolog(o["iolog"], d, "iolog")); err != nil {
		if vv, ok := fortiAPIPatch(o["iolog"], "SystemLocallogSyslogd2Filter-Iolog"); ok {
			if err = d.Set("iolog", vv); err != nil {
				return fmt.Errorf("Error reading iolog: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading iolog: %v", err)
		}
	}

	if err = d.Set("logd", flattenSystemLocallogSyslogd2FilterLogd(o["logd"], d, "logd")); err != nil {
		if vv, ok := fortiAPIPatch(o["logd"], "SystemLocallogSyslogd2Filter-Logd"); ok {
			if err = d.Set("logd", vv); err != nil {
				return fmt.Errorf("Error reading logd: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logd: %v", err)
		}
	}

	if err = d.Set("logdb", flattenSystemLocallogSyslogd2FilterLogdb(o["logdb"], d, "logdb")); err != nil {
		if vv, ok := fortiAPIPatch(o["logdb"], "SystemLocallogSyslogd2Filter-Logdb"); ok {
			if err = d.Set("logdb", vv); err != nil {
				return fmt.Errorf("Error reading logdb: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logdb: %v", err)
		}
	}

	if err = d.Set("logdev", flattenSystemLocallogSyslogd2FilterLogdev(o["logdev"], d, "logdev")); err != nil {
		if vv, ok := fortiAPIPatch(o["logdev"], "SystemLocallogSyslogd2Filter-Logdev"); ok {
			if err = d.Set("logdev", vv); err != nil {
				return fmt.Errorf("Error reading logdev: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logdev: %v", err)
		}
	}

	if err = d.Set("logfile", flattenSystemLocallogSyslogd2FilterLogfile(o["logfile"], d, "logfile")); err != nil {
		if vv, ok := fortiAPIPatch(o["logfile"], "SystemLocallogSyslogd2Filter-Logfile"); ok {
			if err = d.Set("logfile", vv); err != nil {
				return fmt.Errorf("Error reading logfile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logfile: %v", err)
		}
	}

	if err = d.Set("logging", flattenSystemLocallogSyslogd2FilterLogging(o["logging"], d, "logging")); err != nil {
		if vv, ok := fortiAPIPatch(o["logging"], "SystemLocallogSyslogd2Filter-Logging"); ok {
			if err = d.Set("logging", vv); err != nil {
				return fmt.Errorf("Error reading logging: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logging: %v", err)
		}
	}

	if err = d.Set("lrmgr", flattenSystemLocallogSyslogd2FilterLrmgr(o["lrmgr"], d, "lrmgr")); err != nil {
		if vv, ok := fortiAPIPatch(o["lrmgr"], "SystemLocallogSyslogd2Filter-Lrmgr"); ok {
			if err = d.Set("lrmgr", vv); err != nil {
				return fmt.Errorf("Error reading lrmgr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lrmgr: %v", err)
		}
	}

	if err = d.Set("objcfg", flattenSystemLocallogSyslogd2FilterObjcfg(o["objcfg"], d, "objcfg")); err != nil {
		if vv, ok := fortiAPIPatch(o["objcfg"], "SystemLocallogSyslogd2Filter-Objcfg"); ok {
			if err = d.Set("objcfg", vv); err != nil {
				return fmt.Errorf("Error reading objcfg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading objcfg: %v", err)
		}
	}

	if err = d.Set("report", flattenSystemLocallogSyslogd2FilterReport(o["report"], d, "report")); err != nil {
		if vv, ok := fortiAPIPatch(o["report"], "SystemLocallogSyslogd2Filter-Report"); ok {
			if err = d.Set("report", vv); err != nil {
				return fmt.Errorf("Error reading report: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading report: %v", err)
		}
	}

	if err = d.Set("rev", flattenSystemLocallogSyslogd2FilterRev(o["rev"], d, "rev")); err != nil {
		if vv, ok := fortiAPIPatch(o["rev"], "SystemLocallogSyslogd2Filter-Rev"); ok {
			if err = d.Set("rev", vv); err != nil {
				return fmt.Errorf("Error reading rev: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rev: %v", err)
		}
	}

	if err = d.Set("rtmon", flattenSystemLocallogSyslogd2FilterRtmon(o["rtmon"], d, "rtmon")); err != nil {
		if vv, ok := fortiAPIPatch(o["rtmon"], "SystemLocallogSyslogd2Filter-Rtmon"); ok {
			if err = d.Set("rtmon", vv); err != nil {
				return fmt.Errorf("Error reading rtmon: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rtmon: %v", err)
		}
	}

	if err = d.Set("scfw", flattenSystemLocallogSyslogd2FilterScfw(o["scfw"], d, "scfw")); err != nil {
		if vv, ok := fortiAPIPatch(o["scfw"], "SystemLocallogSyslogd2Filter-Scfw"); ok {
			if err = d.Set("scfw", vv); err != nil {
				return fmt.Errorf("Error reading scfw: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scfw: %v", err)
		}
	}

	if err = d.Set("scply", flattenSystemLocallogSyslogd2FilterScply(o["scply"], d, "scply")); err != nil {
		if vv, ok := fortiAPIPatch(o["scply"], "SystemLocallogSyslogd2Filter-Scply"); ok {
			if err = d.Set("scply", vv); err != nil {
				return fmt.Errorf("Error reading scply: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scply: %v", err)
		}
	}

	if err = d.Set("scrmgr", flattenSystemLocallogSyslogd2FilterScrmgr(o["scrmgr"], d, "scrmgr")); err != nil {
		if vv, ok := fortiAPIPatch(o["scrmgr"], "SystemLocallogSyslogd2Filter-Scrmgr"); ok {
			if err = d.Set("scrmgr", vv); err != nil {
				return fmt.Errorf("Error reading scrmgr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scrmgr: %v", err)
		}
	}

	if err = d.Set("scvpn", flattenSystemLocallogSyslogd2FilterScvpn(o["scvpn"], d, "scvpn")); err != nil {
		if vv, ok := fortiAPIPatch(o["scvpn"], "SystemLocallogSyslogd2Filter-Scvpn"); ok {
			if err = d.Set("scvpn", vv); err != nil {
				return fmt.Errorf("Error reading scvpn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scvpn: %v", err)
		}
	}

	if err = d.Set("system", flattenSystemLocallogSyslogd2FilterSystem(o["system"], d, "system")); err != nil {
		if vv, ok := fortiAPIPatch(o["system"], "SystemLocallogSyslogd2Filter-System"); ok {
			if err = d.Set("system", vv); err != nil {
				return fmt.Errorf("Error reading system: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading system: %v", err)
		}
	}

	if err = d.Set("webport", flattenSystemLocallogSyslogd2FilterWebport(o["webport"], d, "webport")); err != nil {
		if vv, ok := fortiAPIPatch(o["webport"], "SystemLocallogSyslogd2Filter-Webport"); ok {
			if err = d.Set("webport", vv); err != nil {
				return fmt.Errorf("Error reading webport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webport: %v", err)
		}
	}

	return nil
}

func flattenSystemLocallogSyslogd2FilterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemLocallogSyslogd2FilterAid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd2FilterDevcfg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd2FilterDevops(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd2FilterDiskquota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd2FilterDm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd2FilterDocker(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd2FilterDvm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd2FilterEdiscovery(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd2FilterEpmgr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd2FilterEvent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd2FilterEventmgmt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd2FilterFaz(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd2FilterFazha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd2FilterFazsys(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd2FilterFgd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd2FilterFgfm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd2FilterFips(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd2FilterFmgws(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd2FilterFmlmgr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd2FilterFmwmgr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd2FilterFortiview(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd2FilterGlbcfg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd2FilterHa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd2FilterHcache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd2FilterIncident(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd2FilterIolog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd2FilterLogd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd2FilterLogdb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd2FilterLogdev(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd2FilterLogfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd2FilterLogging(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd2FilterLrmgr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd2FilterObjcfg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd2FilterReport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd2FilterRev(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd2FilterRtmon(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd2FilterScfw(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd2FilterScply(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd2FilterScrmgr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd2FilterScvpn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd2FilterSystem(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogSyslogd2FilterWebport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemLocallogSyslogd2Filter(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("aid"); ok || d.HasChange("aid") {
		t, err := expandSystemLocallogSyslogd2FilterAid(d, v, "aid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aid"] = t
		}
	}

	if v, ok := d.GetOk("devcfg"); ok || d.HasChange("devcfg") {
		t, err := expandSystemLocallogSyslogd2FilterDevcfg(d, v, "devcfg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["devcfg"] = t
		}
	}

	if v, ok := d.GetOk("devops"); ok || d.HasChange("devops") {
		t, err := expandSystemLocallogSyslogd2FilterDevops(d, v, "devops")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["devops"] = t
		}
	}

	if v, ok := d.GetOk("diskquota"); ok || d.HasChange("diskquota") {
		t, err := expandSystemLocallogSyslogd2FilterDiskquota(d, v, "diskquota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diskquota"] = t
		}
	}

	if v, ok := d.GetOk("dm"); ok || d.HasChange("dm") {
		t, err := expandSystemLocallogSyslogd2FilterDm(d, v, "dm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dm"] = t
		}
	}

	if v, ok := d.GetOk("docker"); ok || d.HasChange("docker") {
		t, err := expandSystemLocallogSyslogd2FilterDocker(d, v, "docker")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["docker"] = t
		}
	}

	if v, ok := d.GetOk("dvm"); ok || d.HasChange("dvm") {
		t, err := expandSystemLocallogSyslogd2FilterDvm(d, v, "dvm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dvm"] = t
		}
	}

	if v, ok := d.GetOk("ediscovery"); ok || d.HasChange("ediscovery") {
		t, err := expandSystemLocallogSyslogd2FilterEdiscovery(d, v, "ediscovery")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ediscovery"] = t
		}
	}

	if v, ok := d.GetOk("epmgr"); ok || d.HasChange("epmgr") {
		t, err := expandSystemLocallogSyslogd2FilterEpmgr(d, v, "epmgr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["epmgr"] = t
		}
	}

	if v, ok := d.GetOk("event"); ok || d.HasChange("event") {
		t, err := expandSystemLocallogSyslogd2FilterEvent(d, v, "event")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["event"] = t
		}
	}

	if v, ok := d.GetOk("eventmgmt"); ok || d.HasChange("eventmgmt") {
		t, err := expandSystemLocallogSyslogd2FilterEventmgmt(d, v, "eventmgmt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eventmgmt"] = t
		}
	}

	if v, ok := d.GetOk("faz"); ok || d.HasChange("faz") {
		t, err := expandSystemLocallogSyslogd2FilterFaz(d, v, "faz")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["faz"] = t
		}
	}

	if v, ok := d.GetOk("fazha"); ok || d.HasChange("fazha") {
		t, err := expandSystemLocallogSyslogd2FilterFazha(d, v, "fazha")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fazha"] = t
		}
	}

	if v, ok := d.GetOk("fazsys"); ok || d.HasChange("fazsys") {
		t, err := expandSystemLocallogSyslogd2FilterFazsys(d, v, "fazsys")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fazsys"] = t
		}
	}

	if v, ok := d.GetOk("fgd"); ok || d.HasChange("fgd") {
		t, err := expandSystemLocallogSyslogd2FilterFgd(d, v, "fgd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgd"] = t
		}
	}

	if v, ok := d.GetOk("fgfm"); ok || d.HasChange("fgfm") {
		t, err := expandSystemLocallogSyslogd2FilterFgfm(d, v, "fgfm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgfm"] = t
		}
	}

	if v, ok := d.GetOk("fips"); ok || d.HasChange("fips") {
		t, err := expandSystemLocallogSyslogd2FilterFips(d, v, "fips")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fips"] = t
		}
	}

	if v, ok := d.GetOk("fmgws"); ok || d.HasChange("fmgws") {
		t, err := expandSystemLocallogSyslogd2FilterFmgws(d, v, "fmgws")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fmgws"] = t
		}
	}

	if v, ok := d.GetOk("fmlmgr"); ok || d.HasChange("fmlmgr") {
		t, err := expandSystemLocallogSyslogd2FilterFmlmgr(d, v, "fmlmgr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fmlmgr"] = t
		}
	}

	if v, ok := d.GetOk("fmwmgr"); ok || d.HasChange("fmwmgr") {
		t, err := expandSystemLocallogSyslogd2FilterFmwmgr(d, v, "fmwmgr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fmwmgr"] = t
		}
	}

	if v, ok := d.GetOk("fortiview"); ok || d.HasChange("fortiview") {
		t, err := expandSystemLocallogSyslogd2FilterFortiview(d, v, "fortiview")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiview"] = t
		}
	}

	if v, ok := d.GetOk("glbcfg"); ok || d.HasChange("glbcfg") {
		t, err := expandSystemLocallogSyslogd2FilterGlbcfg(d, v, "glbcfg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["glbcfg"] = t
		}
	}

	if v, ok := d.GetOk("ha"); ok || d.HasChange("ha") {
		t, err := expandSystemLocallogSyslogd2FilterHa(d, v, "ha")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha"] = t
		}
	}

	if v, ok := d.GetOk("hcache"); ok || d.HasChange("hcache") {
		t, err := expandSystemLocallogSyslogd2FilterHcache(d, v, "hcache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hcache"] = t
		}
	}

	if v, ok := d.GetOk("incident"); ok || d.HasChange("incident") {
		t, err := expandSystemLocallogSyslogd2FilterIncident(d, v, "incident")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["incident"] = t
		}
	}

	if v, ok := d.GetOk("iolog"); ok || d.HasChange("iolog") {
		t, err := expandSystemLocallogSyslogd2FilterIolog(d, v, "iolog")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["iolog"] = t
		}
	}

	if v, ok := d.GetOk("logd"); ok || d.HasChange("logd") {
		t, err := expandSystemLocallogSyslogd2FilterLogd(d, v, "logd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logd"] = t
		}
	}

	if v, ok := d.GetOk("logdb"); ok || d.HasChange("logdb") {
		t, err := expandSystemLocallogSyslogd2FilterLogdb(d, v, "logdb")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logdb"] = t
		}
	}

	if v, ok := d.GetOk("logdev"); ok || d.HasChange("logdev") {
		t, err := expandSystemLocallogSyslogd2FilterLogdev(d, v, "logdev")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logdev"] = t
		}
	}

	if v, ok := d.GetOk("logfile"); ok || d.HasChange("logfile") {
		t, err := expandSystemLocallogSyslogd2FilterLogfile(d, v, "logfile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logfile"] = t
		}
	}

	if v, ok := d.GetOk("logging"); ok || d.HasChange("logging") {
		t, err := expandSystemLocallogSyslogd2FilterLogging(d, v, "logging")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logging"] = t
		}
	}

	if v, ok := d.GetOk("lrmgr"); ok || d.HasChange("lrmgr") {
		t, err := expandSystemLocallogSyslogd2FilterLrmgr(d, v, "lrmgr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lrmgr"] = t
		}
	}

	if v, ok := d.GetOk("objcfg"); ok || d.HasChange("objcfg") {
		t, err := expandSystemLocallogSyslogd2FilterObjcfg(d, v, "objcfg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["objcfg"] = t
		}
	}

	if v, ok := d.GetOk("report"); ok || d.HasChange("report") {
		t, err := expandSystemLocallogSyslogd2FilterReport(d, v, "report")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["report"] = t
		}
	}

	if v, ok := d.GetOk("rev"); ok || d.HasChange("rev") {
		t, err := expandSystemLocallogSyslogd2FilterRev(d, v, "rev")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rev"] = t
		}
	}

	if v, ok := d.GetOk("rtmon"); ok || d.HasChange("rtmon") {
		t, err := expandSystemLocallogSyslogd2FilterRtmon(d, v, "rtmon")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rtmon"] = t
		}
	}

	if v, ok := d.GetOk("scfw"); ok || d.HasChange("scfw") {
		t, err := expandSystemLocallogSyslogd2FilterScfw(d, v, "scfw")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scfw"] = t
		}
	}

	if v, ok := d.GetOk("scply"); ok || d.HasChange("scply") {
		t, err := expandSystemLocallogSyslogd2FilterScply(d, v, "scply")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scply"] = t
		}
	}

	if v, ok := d.GetOk("scrmgr"); ok || d.HasChange("scrmgr") {
		t, err := expandSystemLocallogSyslogd2FilterScrmgr(d, v, "scrmgr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scrmgr"] = t
		}
	}

	if v, ok := d.GetOk("scvpn"); ok || d.HasChange("scvpn") {
		t, err := expandSystemLocallogSyslogd2FilterScvpn(d, v, "scvpn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scvpn"] = t
		}
	}

	if v, ok := d.GetOk("system"); ok || d.HasChange("system") {
		t, err := expandSystemLocallogSyslogd2FilterSystem(d, v, "system")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["system"] = t
		}
	}

	if v, ok := d.GetOk("webport"); ok || d.HasChange("webport") {
		t, err := expandSystemLocallogSyslogd2FilterWebport(d, v, "webport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webport"] = t
		}
	}

	return &obj, nil
}
