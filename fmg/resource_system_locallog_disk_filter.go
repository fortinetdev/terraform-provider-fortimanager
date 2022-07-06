// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Filter for disk logging.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemLocallogDiskFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLocallogDiskFilterUpdate,
		Read:   resourceSystemLocallogDiskFilterRead,
		Update: resourceSystemLocallogDiskFilterUpdate,
		Delete: resourceSystemLocallogDiskFilterDelete,

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

func resourceSystemLocallogDiskFilterUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemLocallogDiskFilter(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemLocallogDiskFilter resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemLocallogDiskFilter(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemLocallogDiskFilter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemLocallogDiskFilter")

	return resourceSystemLocallogDiskFilterRead(d, m)
}

func resourceSystemLocallogDiskFilterDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemLocallogDiskFilter(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemLocallogDiskFilter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLocallogDiskFilterRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemLocallogDiskFilter(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemLocallogDiskFilter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemLocallogDiskFilter(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemLocallogDiskFilter resource from API: %v", err)
	}
	return nil
}

func flattenSystemLocallogDiskFilterAid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskFilterDevcfg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskFilterDevops(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskFilterDiskquota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskFilterDm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskFilterDocker(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskFilterDvm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskFilterEdiscovery(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskFilterEpmgr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskFilterEvent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskFilterEventmgmt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskFilterFaz(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskFilterFazha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskFilterFazsys(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskFilterFgd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskFilterFgfm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskFilterFips(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskFilterFmgws(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskFilterFmlmgr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskFilterFmwmgr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskFilterFortiview(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskFilterGlbcfg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskFilterHa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskFilterHcache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskFilterIncident(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskFilterIolog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskFilterLogd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskFilterLogdb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskFilterLogdev(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskFilterLogfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskFilterLogging(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskFilterLrmgr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskFilterObjcfg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskFilterReport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskFilterRev(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskFilterRtmon(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskFilterScfw(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskFilterScply(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskFilterScrmgr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskFilterScvpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskFilterSystem(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogDiskFilterWebport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemLocallogDiskFilter(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("aid", flattenSystemLocallogDiskFilterAid(o["aid"], d, "aid")); err != nil {
		if vv, ok := fortiAPIPatch(o["aid"], "SystemLocallogDiskFilter-Aid"); ok {
			if err = d.Set("aid", vv); err != nil {
				return fmt.Errorf("Error reading aid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading aid: %v", err)
		}
	}

	if err = d.Set("devcfg", flattenSystemLocallogDiskFilterDevcfg(o["devcfg"], d, "devcfg")); err != nil {
		if vv, ok := fortiAPIPatch(o["devcfg"], "SystemLocallogDiskFilter-Devcfg"); ok {
			if err = d.Set("devcfg", vv); err != nil {
				return fmt.Errorf("Error reading devcfg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading devcfg: %v", err)
		}
	}

	if err = d.Set("devops", flattenSystemLocallogDiskFilterDevops(o["devops"], d, "devops")); err != nil {
		if vv, ok := fortiAPIPatch(o["devops"], "SystemLocallogDiskFilter-Devops"); ok {
			if err = d.Set("devops", vv); err != nil {
				return fmt.Errorf("Error reading devops: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading devops: %v", err)
		}
	}

	if err = d.Set("diskquota", flattenSystemLocallogDiskFilterDiskquota(o["diskquota"], d, "diskquota")); err != nil {
		if vv, ok := fortiAPIPatch(o["diskquota"], "SystemLocallogDiskFilter-Diskquota"); ok {
			if err = d.Set("diskquota", vv); err != nil {
				return fmt.Errorf("Error reading diskquota: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diskquota: %v", err)
		}
	}

	if err = d.Set("dm", flattenSystemLocallogDiskFilterDm(o["dm"], d, "dm")); err != nil {
		if vv, ok := fortiAPIPatch(o["dm"], "SystemLocallogDiskFilter-Dm"); ok {
			if err = d.Set("dm", vv); err != nil {
				return fmt.Errorf("Error reading dm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dm: %v", err)
		}
	}

	if err = d.Set("docker", flattenSystemLocallogDiskFilterDocker(o["docker"], d, "docker")); err != nil {
		if vv, ok := fortiAPIPatch(o["docker"], "SystemLocallogDiskFilter-Docker"); ok {
			if err = d.Set("docker", vv); err != nil {
				return fmt.Errorf("Error reading docker: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading docker: %v", err)
		}
	}

	if err = d.Set("dvm", flattenSystemLocallogDiskFilterDvm(o["dvm"], d, "dvm")); err != nil {
		if vv, ok := fortiAPIPatch(o["dvm"], "SystemLocallogDiskFilter-Dvm"); ok {
			if err = d.Set("dvm", vv); err != nil {
				return fmt.Errorf("Error reading dvm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dvm: %v", err)
		}
	}

	if err = d.Set("ediscovery", flattenSystemLocallogDiskFilterEdiscovery(o["ediscovery"], d, "ediscovery")); err != nil {
		if vv, ok := fortiAPIPatch(o["ediscovery"], "SystemLocallogDiskFilter-Ediscovery"); ok {
			if err = d.Set("ediscovery", vv); err != nil {
				return fmt.Errorf("Error reading ediscovery: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ediscovery: %v", err)
		}
	}

	if err = d.Set("epmgr", flattenSystemLocallogDiskFilterEpmgr(o["epmgr"], d, "epmgr")); err != nil {
		if vv, ok := fortiAPIPatch(o["epmgr"], "SystemLocallogDiskFilter-Epmgr"); ok {
			if err = d.Set("epmgr", vv); err != nil {
				return fmt.Errorf("Error reading epmgr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading epmgr: %v", err)
		}
	}

	if err = d.Set("event", flattenSystemLocallogDiskFilterEvent(o["event"], d, "event")); err != nil {
		if vv, ok := fortiAPIPatch(o["event"], "SystemLocallogDiskFilter-Event"); ok {
			if err = d.Set("event", vv); err != nil {
				return fmt.Errorf("Error reading event: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading event: %v", err)
		}
	}

	if err = d.Set("eventmgmt", flattenSystemLocallogDiskFilterEventmgmt(o["eventmgmt"], d, "eventmgmt")); err != nil {
		if vv, ok := fortiAPIPatch(o["eventmgmt"], "SystemLocallogDiskFilter-Eventmgmt"); ok {
			if err = d.Set("eventmgmt", vv); err != nil {
				return fmt.Errorf("Error reading eventmgmt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eventmgmt: %v", err)
		}
	}

	if err = d.Set("faz", flattenSystemLocallogDiskFilterFaz(o["faz"], d, "faz")); err != nil {
		if vv, ok := fortiAPIPatch(o["faz"], "SystemLocallogDiskFilter-Faz"); ok {
			if err = d.Set("faz", vv); err != nil {
				return fmt.Errorf("Error reading faz: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading faz: %v", err)
		}
	}

	if err = d.Set("fazha", flattenSystemLocallogDiskFilterFazha(o["fazha"], d, "fazha")); err != nil {
		if vv, ok := fortiAPIPatch(o["fazha"], "SystemLocallogDiskFilter-Fazha"); ok {
			if err = d.Set("fazha", vv); err != nil {
				return fmt.Errorf("Error reading fazha: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fazha: %v", err)
		}
	}

	if err = d.Set("fazsys", flattenSystemLocallogDiskFilterFazsys(o["fazsys"], d, "fazsys")); err != nil {
		if vv, ok := fortiAPIPatch(o["fazsys"], "SystemLocallogDiskFilter-Fazsys"); ok {
			if err = d.Set("fazsys", vv); err != nil {
				return fmt.Errorf("Error reading fazsys: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fazsys: %v", err)
		}
	}

	if err = d.Set("fgd", flattenSystemLocallogDiskFilterFgd(o["fgd"], d, "fgd")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgd"], "SystemLocallogDiskFilter-Fgd"); ok {
			if err = d.Set("fgd", vv); err != nil {
				return fmt.Errorf("Error reading fgd: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgd: %v", err)
		}
	}

	if err = d.Set("fgfm", flattenSystemLocallogDiskFilterFgfm(o["fgfm"], d, "fgfm")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgfm"], "SystemLocallogDiskFilter-Fgfm"); ok {
			if err = d.Set("fgfm", vv); err != nil {
				return fmt.Errorf("Error reading fgfm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgfm: %v", err)
		}
	}

	if err = d.Set("fips", flattenSystemLocallogDiskFilterFips(o["fips"], d, "fips")); err != nil {
		if vv, ok := fortiAPIPatch(o["fips"], "SystemLocallogDiskFilter-Fips"); ok {
			if err = d.Set("fips", vv); err != nil {
				return fmt.Errorf("Error reading fips: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fips: %v", err)
		}
	}

	if err = d.Set("fmgws", flattenSystemLocallogDiskFilterFmgws(o["fmgws"], d, "fmgws")); err != nil {
		if vv, ok := fortiAPIPatch(o["fmgws"], "SystemLocallogDiskFilter-Fmgws"); ok {
			if err = d.Set("fmgws", vv); err != nil {
				return fmt.Errorf("Error reading fmgws: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmgws: %v", err)
		}
	}

	if err = d.Set("fmlmgr", flattenSystemLocallogDiskFilterFmlmgr(o["fmlmgr"], d, "fmlmgr")); err != nil {
		if vv, ok := fortiAPIPatch(o["fmlmgr"], "SystemLocallogDiskFilter-Fmlmgr"); ok {
			if err = d.Set("fmlmgr", vv); err != nil {
				return fmt.Errorf("Error reading fmlmgr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmlmgr: %v", err)
		}
	}

	if err = d.Set("fmwmgr", flattenSystemLocallogDiskFilterFmwmgr(o["fmwmgr"], d, "fmwmgr")); err != nil {
		if vv, ok := fortiAPIPatch(o["fmwmgr"], "SystemLocallogDiskFilter-Fmwmgr"); ok {
			if err = d.Set("fmwmgr", vv); err != nil {
				return fmt.Errorf("Error reading fmwmgr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmwmgr: %v", err)
		}
	}

	if err = d.Set("fortiview", flattenSystemLocallogDiskFilterFortiview(o["fortiview"], d, "fortiview")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortiview"], "SystemLocallogDiskFilter-Fortiview"); ok {
			if err = d.Set("fortiview", vv); err != nil {
				return fmt.Errorf("Error reading fortiview: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortiview: %v", err)
		}
	}

	if err = d.Set("glbcfg", flattenSystemLocallogDiskFilterGlbcfg(o["glbcfg"], d, "glbcfg")); err != nil {
		if vv, ok := fortiAPIPatch(o["glbcfg"], "SystemLocallogDiskFilter-Glbcfg"); ok {
			if err = d.Set("glbcfg", vv); err != nil {
				return fmt.Errorf("Error reading glbcfg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading glbcfg: %v", err)
		}
	}

	if err = d.Set("ha", flattenSystemLocallogDiskFilterHa(o["ha"], d, "ha")); err != nil {
		if vv, ok := fortiAPIPatch(o["ha"], "SystemLocallogDiskFilter-Ha"); ok {
			if err = d.Set("ha", vv); err != nil {
				return fmt.Errorf("Error reading ha: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ha: %v", err)
		}
	}

	if err = d.Set("hcache", flattenSystemLocallogDiskFilterHcache(o["hcache"], d, "hcache")); err != nil {
		if vv, ok := fortiAPIPatch(o["hcache"], "SystemLocallogDiskFilter-Hcache"); ok {
			if err = d.Set("hcache", vv); err != nil {
				return fmt.Errorf("Error reading hcache: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hcache: %v", err)
		}
	}

	if err = d.Set("incident", flattenSystemLocallogDiskFilterIncident(o["incident"], d, "incident")); err != nil {
		if vv, ok := fortiAPIPatch(o["incident"], "SystemLocallogDiskFilter-Incident"); ok {
			if err = d.Set("incident", vv); err != nil {
				return fmt.Errorf("Error reading incident: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading incident: %v", err)
		}
	}

	if err = d.Set("iolog", flattenSystemLocallogDiskFilterIolog(o["iolog"], d, "iolog")); err != nil {
		if vv, ok := fortiAPIPatch(o["iolog"], "SystemLocallogDiskFilter-Iolog"); ok {
			if err = d.Set("iolog", vv); err != nil {
				return fmt.Errorf("Error reading iolog: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading iolog: %v", err)
		}
	}

	if err = d.Set("logd", flattenSystemLocallogDiskFilterLogd(o["logd"], d, "logd")); err != nil {
		if vv, ok := fortiAPIPatch(o["logd"], "SystemLocallogDiskFilter-Logd"); ok {
			if err = d.Set("logd", vv); err != nil {
				return fmt.Errorf("Error reading logd: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logd: %v", err)
		}
	}

	if err = d.Set("logdb", flattenSystemLocallogDiskFilterLogdb(o["logdb"], d, "logdb")); err != nil {
		if vv, ok := fortiAPIPatch(o["logdb"], "SystemLocallogDiskFilter-Logdb"); ok {
			if err = d.Set("logdb", vv); err != nil {
				return fmt.Errorf("Error reading logdb: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logdb: %v", err)
		}
	}

	if err = d.Set("logdev", flattenSystemLocallogDiskFilterLogdev(o["logdev"], d, "logdev")); err != nil {
		if vv, ok := fortiAPIPatch(o["logdev"], "SystemLocallogDiskFilter-Logdev"); ok {
			if err = d.Set("logdev", vv); err != nil {
				return fmt.Errorf("Error reading logdev: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logdev: %v", err)
		}
	}

	if err = d.Set("logfile", flattenSystemLocallogDiskFilterLogfile(o["logfile"], d, "logfile")); err != nil {
		if vv, ok := fortiAPIPatch(o["logfile"], "SystemLocallogDiskFilter-Logfile"); ok {
			if err = d.Set("logfile", vv); err != nil {
				return fmt.Errorf("Error reading logfile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logfile: %v", err)
		}
	}

	if err = d.Set("logging", flattenSystemLocallogDiskFilterLogging(o["logging"], d, "logging")); err != nil {
		if vv, ok := fortiAPIPatch(o["logging"], "SystemLocallogDiskFilter-Logging"); ok {
			if err = d.Set("logging", vv); err != nil {
				return fmt.Errorf("Error reading logging: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logging: %v", err)
		}
	}

	if err = d.Set("lrmgr", flattenSystemLocallogDiskFilterLrmgr(o["lrmgr"], d, "lrmgr")); err != nil {
		if vv, ok := fortiAPIPatch(o["lrmgr"], "SystemLocallogDiskFilter-Lrmgr"); ok {
			if err = d.Set("lrmgr", vv); err != nil {
				return fmt.Errorf("Error reading lrmgr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lrmgr: %v", err)
		}
	}

	if err = d.Set("objcfg", flattenSystemLocallogDiskFilterObjcfg(o["objcfg"], d, "objcfg")); err != nil {
		if vv, ok := fortiAPIPatch(o["objcfg"], "SystemLocallogDiskFilter-Objcfg"); ok {
			if err = d.Set("objcfg", vv); err != nil {
				return fmt.Errorf("Error reading objcfg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading objcfg: %v", err)
		}
	}

	if err = d.Set("report", flattenSystemLocallogDiskFilterReport(o["report"], d, "report")); err != nil {
		if vv, ok := fortiAPIPatch(o["report"], "SystemLocallogDiskFilter-Report"); ok {
			if err = d.Set("report", vv); err != nil {
				return fmt.Errorf("Error reading report: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading report: %v", err)
		}
	}

	if err = d.Set("rev", flattenSystemLocallogDiskFilterRev(o["rev"], d, "rev")); err != nil {
		if vv, ok := fortiAPIPatch(o["rev"], "SystemLocallogDiskFilter-Rev"); ok {
			if err = d.Set("rev", vv); err != nil {
				return fmt.Errorf("Error reading rev: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rev: %v", err)
		}
	}

	if err = d.Set("rtmon", flattenSystemLocallogDiskFilterRtmon(o["rtmon"], d, "rtmon")); err != nil {
		if vv, ok := fortiAPIPatch(o["rtmon"], "SystemLocallogDiskFilter-Rtmon"); ok {
			if err = d.Set("rtmon", vv); err != nil {
				return fmt.Errorf("Error reading rtmon: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rtmon: %v", err)
		}
	}

	if err = d.Set("scfw", flattenSystemLocallogDiskFilterScfw(o["scfw"], d, "scfw")); err != nil {
		if vv, ok := fortiAPIPatch(o["scfw"], "SystemLocallogDiskFilter-Scfw"); ok {
			if err = d.Set("scfw", vv); err != nil {
				return fmt.Errorf("Error reading scfw: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scfw: %v", err)
		}
	}

	if err = d.Set("scply", flattenSystemLocallogDiskFilterScply(o["scply"], d, "scply")); err != nil {
		if vv, ok := fortiAPIPatch(o["scply"], "SystemLocallogDiskFilter-Scply"); ok {
			if err = d.Set("scply", vv); err != nil {
				return fmt.Errorf("Error reading scply: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scply: %v", err)
		}
	}

	if err = d.Set("scrmgr", flattenSystemLocallogDiskFilterScrmgr(o["scrmgr"], d, "scrmgr")); err != nil {
		if vv, ok := fortiAPIPatch(o["scrmgr"], "SystemLocallogDiskFilter-Scrmgr"); ok {
			if err = d.Set("scrmgr", vv); err != nil {
				return fmt.Errorf("Error reading scrmgr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scrmgr: %v", err)
		}
	}

	if err = d.Set("scvpn", flattenSystemLocallogDiskFilterScvpn(o["scvpn"], d, "scvpn")); err != nil {
		if vv, ok := fortiAPIPatch(o["scvpn"], "SystemLocallogDiskFilter-Scvpn"); ok {
			if err = d.Set("scvpn", vv); err != nil {
				return fmt.Errorf("Error reading scvpn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scvpn: %v", err)
		}
	}

	if err = d.Set("system", flattenSystemLocallogDiskFilterSystem(o["system"], d, "system")); err != nil {
		if vv, ok := fortiAPIPatch(o["system"], "SystemLocallogDiskFilter-System"); ok {
			if err = d.Set("system", vv); err != nil {
				return fmt.Errorf("Error reading system: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading system: %v", err)
		}
	}

	if err = d.Set("webport", flattenSystemLocallogDiskFilterWebport(o["webport"], d, "webport")); err != nil {
		if vv, ok := fortiAPIPatch(o["webport"], "SystemLocallogDiskFilter-Webport"); ok {
			if err = d.Set("webport", vv); err != nil {
				return fmt.Errorf("Error reading webport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webport: %v", err)
		}
	}

	return nil
}

func flattenSystemLocallogDiskFilterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemLocallogDiskFilterAid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskFilterDevcfg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskFilterDevops(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskFilterDiskquota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskFilterDm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskFilterDocker(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskFilterDvm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskFilterEdiscovery(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskFilterEpmgr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskFilterEvent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskFilterEventmgmt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskFilterFaz(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskFilterFazha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskFilterFazsys(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskFilterFgd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskFilterFgfm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskFilterFips(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskFilterFmgws(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskFilterFmlmgr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskFilterFmwmgr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskFilterFortiview(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskFilterGlbcfg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskFilterHa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskFilterHcache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskFilterIncident(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskFilterIolog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskFilterLogd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskFilterLogdb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskFilterLogdev(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskFilterLogfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskFilterLogging(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskFilterLrmgr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskFilterObjcfg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskFilterReport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskFilterRev(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskFilterRtmon(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskFilterScfw(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskFilterScply(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskFilterScrmgr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskFilterScvpn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskFilterSystem(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogDiskFilterWebport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemLocallogDiskFilter(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("aid"); ok || d.HasChange("aid") {
		t, err := expandSystemLocallogDiskFilterAid(d, v, "aid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aid"] = t
		}
	}

	if v, ok := d.GetOk("devcfg"); ok || d.HasChange("devcfg") {
		t, err := expandSystemLocallogDiskFilterDevcfg(d, v, "devcfg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["devcfg"] = t
		}
	}

	if v, ok := d.GetOk("devops"); ok || d.HasChange("devops") {
		t, err := expandSystemLocallogDiskFilterDevops(d, v, "devops")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["devops"] = t
		}
	}

	if v, ok := d.GetOk("diskquota"); ok || d.HasChange("diskquota") {
		t, err := expandSystemLocallogDiskFilterDiskquota(d, v, "diskquota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diskquota"] = t
		}
	}

	if v, ok := d.GetOk("dm"); ok || d.HasChange("dm") {
		t, err := expandSystemLocallogDiskFilterDm(d, v, "dm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dm"] = t
		}
	}

	if v, ok := d.GetOk("docker"); ok || d.HasChange("docker") {
		t, err := expandSystemLocallogDiskFilterDocker(d, v, "docker")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["docker"] = t
		}
	}

	if v, ok := d.GetOk("dvm"); ok || d.HasChange("dvm") {
		t, err := expandSystemLocallogDiskFilterDvm(d, v, "dvm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dvm"] = t
		}
	}

	if v, ok := d.GetOk("ediscovery"); ok || d.HasChange("ediscovery") {
		t, err := expandSystemLocallogDiskFilterEdiscovery(d, v, "ediscovery")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ediscovery"] = t
		}
	}

	if v, ok := d.GetOk("epmgr"); ok || d.HasChange("epmgr") {
		t, err := expandSystemLocallogDiskFilterEpmgr(d, v, "epmgr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["epmgr"] = t
		}
	}

	if v, ok := d.GetOk("event"); ok || d.HasChange("event") {
		t, err := expandSystemLocallogDiskFilterEvent(d, v, "event")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["event"] = t
		}
	}

	if v, ok := d.GetOk("eventmgmt"); ok || d.HasChange("eventmgmt") {
		t, err := expandSystemLocallogDiskFilterEventmgmt(d, v, "eventmgmt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eventmgmt"] = t
		}
	}

	if v, ok := d.GetOk("faz"); ok || d.HasChange("faz") {
		t, err := expandSystemLocallogDiskFilterFaz(d, v, "faz")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["faz"] = t
		}
	}

	if v, ok := d.GetOk("fazha"); ok || d.HasChange("fazha") {
		t, err := expandSystemLocallogDiskFilterFazha(d, v, "fazha")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fazha"] = t
		}
	}

	if v, ok := d.GetOk("fazsys"); ok || d.HasChange("fazsys") {
		t, err := expandSystemLocallogDiskFilterFazsys(d, v, "fazsys")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fazsys"] = t
		}
	}

	if v, ok := d.GetOk("fgd"); ok || d.HasChange("fgd") {
		t, err := expandSystemLocallogDiskFilterFgd(d, v, "fgd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgd"] = t
		}
	}

	if v, ok := d.GetOk("fgfm"); ok || d.HasChange("fgfm") {
		t, err := expandSystemLocallogDiskFilterFgfm(d, v, "fgfm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgfm"] = t
		}
	}

	if v, ok := d.GetOk("fips"); ok || d.HasChange("fips") {
		t, err := expandSystemLocallogDiskFilterFips(d, v, "fips")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fips"] = t
		}
	}

	if v, ok := d.GetOk("fmgws"); ok || d.HasChange("fmgws") {
		t, err := expandSystemLocallogDiskFilterFmgws(d, v, "fmgws")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fmgws"] = t
		}
	}

	if v, ok := d.GetOk("fmlmgr"); ok || d.HasChange("fmlmgr") {
		t, err := expandSystemLocallogDiskFilterFmlmgr(d, v, "fmlmgr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fmlmgr"] = t
		}
	}

	if v, ok := d.GetOk("fmwmgr"); ok || d.HasChange("fmwmgr") {
		t, err := expandSystemLocallogDiskFilterFmwmgr(d, v, "fmwmgr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fmwmgr"] = t
		}
	}

	if v, ok := d.GetOk("fortiview"); ok || d.HasChange("fortiview") {
		t, err := expandSystemLocallogDiskFilterFortiview(d, v, "fortiview")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiview"] = t
		}
	}

	if v, ok := d.GetOk("glbcfg"); ok || d.HasChange("glbcfg") {
		t, err := expandSystemLocallogDiskFilterGlbcfg(d, v, "glbcfg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["glbcfg"] = t
		}
	}

	if v, ok := d.GetOk("ha"); ok || d.HasChange("ha") {
		t, err := expandSystemLocallogDiskFilterHa(d, v, "ha")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha"] = t
		}
	}

	if v, ok := d.GetOk("hcache"); ok || d.HasChange("hcache") {
		t, err := expandSystemLocallogDiskFilterHcache(d, v, "hcache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hcache"] = t
		}
	}

	if v, ok := d.GetOk("incident"); ok || d.HasChange("incident") {
		t, err := expandSystemLocallogDiskFilterIncident(d, v, "incident")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["incident"] = t
		}
	}

	if v, ok := d.GetOk("iolog"); ok || d.HasChange("iolog") {
		t, err := expandSystemLocallogDiskFilterIolog(d, v, "iolog")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["iolog"] = t
		}
	}

	if v, ok := d.GetOk("logd"); ok || d.HasChange("logd") {
		t, err := expandSystemLocallogDiskFilterLogd(d, v, "logd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logd"] = t
		}
	}

	if v, ok := d.GetOk("logdb"); ok || d.HasChange("logdb") {
		t, err := expandSystemLocallogDiskFilterLogdb(d, v, "logdb")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logdb"] = t
		}
	}

	if v, ok := d.GetOk("logdev"); ok || d.HasChange("logdev") {
		t, err := expandSystemLocallogDiskFilterLogdev(d, v, "logdev")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logdev"] = t
		}
	}

	if v, ok := d.GetOk("logfile"); ok || d.HasChange("logfile") {
		t, err := expandSystemLocallogDiskFilterLogfile(d, v, "logfile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logfile"] = t
		}
	}

	if v, ok := d.GetOk("logging"); ok || d.HasChange("logging") {
		t, err := expandSystemLocallogDiskFilterLogging(d, v, "logging")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logging"] = t
		}
	}

	if v, ok := d.GetOk("lrmgr"); ok || d.HasChange("lrmgr") {
		t, err := expandSystemLocallogDiskFilterLrmgr(d, v, "lrmgr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lrmgr"] = t
		}
	}

	if v, ok := d.GetOk("objcfg"); ok || d.HasChange("objcfg") {
		t, err := expandSystemLocallogDiskFilterObjcfg(d, v, "objcfg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["objcfg"] = t
		}
	}

	if v, ok := d.GetOk("report"); ok || d.HasChange("report") {
		t, err := expandSystemLocallogDiskFilterReport(d, v, "report")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["report"] = t
		}
	}

	if v, ok := d.GetOk("rev"); ok || d.HasChange("rev") {
		t, err := expandSystemLocallogDiskFilterRev(d, v, "rev")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rev"] = t
		}
	}

	if v, ok := d.GetOk("rtmon"); ok || d.HasChange("rtmon") {
		t, err := expandSystemLocallogDiskFilterRtmon(d, v, "rtmon")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rtmon"] = t
		}
	}

	if v, ok := d.GetOk("scfw"); ok || d.HasChange("scfw") {
		t, err := expandSystemLocallogDiskFilterScfw(d, v, "scfw")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scfw"] = t
		}
	}

	if v, ok := d.GetOk("scply"); ok || d.HasChange("scply") {
		t, err := expandSystemLocallogDiskFilterScply(d, v, "scply")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scply"] = t
		}
	}

	if v, ok := d.GetOk("scrmgr"); ok || d.HasChange("scrmgr") {
		t, err := expandSystemLocallogDiskFilterScrmgr(d, v, "scrmgr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scrmgr"] = t
		}
	}

	if v, ok := d.GetOk("scvpn"); ok || d.HasChange("scvpn") {
		t, err := expandSystemLocallogDiskFilterScvpn(d, v, "scvpn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scvpn"] = t
		}
	}

	if v, ok := d.GetOk("system"); ok || d.HasChange("system") {
		t, err := expandSystemLocallogDiskFilterSystem(d, v, "system")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["system"] = t
		}
	}

	if v, ok := d.GetOk("webport"); ok || d.HasChange("webport") {
		t, err := expandSystemLocallogDiskFilterWebport(d, v, "webport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webport"] = t
		}
	}

	return &obj, nil
}
