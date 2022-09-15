// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Filter for FortiAnalyzer logging.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemLocallogFortianalyzerFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLocallogFortianalyzerFilterUpdate,
		Read:   resourceSystemLocallogFortianalyzerFilterRead,
		Update: resourceSystemLocallogFortianalyzerFilterUpdate,
		Delete: resourceSystemLocallogFortianalyzerFilterDelete,

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

func resourceSystemLocallogFortianalyzerFilterUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemLocallogFortianalyzerFilter(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemLocallogFortianalyzerFilter resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemLocallogFortianalyzerFilter(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemLocallogFortianalyzerFilter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemLocallogFortianalyzerFilter")

	return resourceSystemLocallogFortianalyzerFilterRead(d, m)
}

func resourceSystemLocallogFortianalyzerFilterDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemLocallogFortianalyzerFilter(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemLocallogFortianalyzerFilter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLocallogFortianalyzerFilterRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemLocallogFortianalyzerFilter(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemLocallogFortianalyzerFilter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemLocallogFortianalyzerFilter(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemLocallogFortianalyzerFilter resource from API: %v", err)
	}
	return nil
}

func flattenSystemLocallogFortianalyzerFilterAid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzerFilterDevcfg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzerFilterDevops(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzerFilterDiskquota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzerFilterDm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzerFilterDocker(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzerFilterDvm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzerFilterEdiscovery(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzerFilterEpmgr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzerFilterEvent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzerFilterEventmgmt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzerFilterFaz(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzerFilterFazha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzerFilterFazsys(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzerFilterFgd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzerFilterFgfm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzerFilterFips(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzerFilterFmgws(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzerFilterFmlmgr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzerFilterFmwmgr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzerFilterFortiview(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzerFilterGlbcfg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzerFilterHa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzerFilterHcache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzerFilterIncident(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzerFilterIolog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzerFilterLogd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzerFilterLogdb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzerFilterLogdev(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzerFilterLogfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzerFilterLogging(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzerFilterLrmgr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzerFilterObjcfg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzerFilterReport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzerFilterRev(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzerFilterRtmon(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzerFilterScfw(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzerFilterScply(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzerFilterScrmgr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzerFilterScvpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzerFilterSystem(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzerFilterWebport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemLocallogFortianalyzerFilter(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("aid", flattenSystemLocallogFortianalyzerFilterAid(o["aid"], d, "aid")); err != nil {
		if vv, ok := fortiAPIPatch(o["aid"], "SystemLocallogFortianalyzerFilter-Aid"); ok {
			if err = d.Set("aid", vv); err != nil {
				return fmt.Errorf("Error reading aid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading aid: %v", err)
		}
	}

	if err = d.Set("devcfg", flattenSystemLocallogFortianalyzerFilterDevcfg(o["devcfg"], d, "devcfg")); err != nil {
		if vv, ok := fortiAPIPatch(o["devcfg"], "SystemLocallogFortianalyzerFilter-Devcfg"); ok {
			if err = d.Set("devcfg", vv); err != nil {
				return fmt.Errorf("Error reading devcfg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading devcfg: %v", err)
		}
	}

	if err = d.Set("devops", flattenSystemLocallogFortianalyzerFilterDevops(o["devops"], d, "devops")); err != nil {
		if vv, ok := fortiAPIPatch(o["devops"], "SystemLocallogFortianalyzerFilter-Devops"); ok {
			if err = d.Set("devops", vv); err != nil {
				return fmt.Errorf("Error reading devops: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading devops: %v", err)
		}
	}

	if err = d.Set("diskquota", flattenSystemLocallogFortianalyzerFilterDiskquota(o["diskquota"], d, "diskquota")); err != nil {
		if vv, ok := fortiAPIPatch(o["diskquota"], "SystemLocallogFortianalyzerFilter-Diskquota"); ok {
			if err = d.Set("diskquota", vv); err != nil {
				return fmt.Errorf("Error reading diskquota: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diskquota: %v", err)
		}
	}

	if err = d.Set("dm", flattenSystemLocallogFortianalyzerFilterDm(o["dm"], d, "dm")); err != nil {
		if vv, ok := fortiAPIPatch(o["dm"], "SystemLocallogFortianalyzerFilter-Dm"); ok {
			if err = d.Set("dm", vv); err != nil {
				return fmt.Errorf("Error reading dm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dm: %v", err)
		}
	}

	if err = d.Set("docker", flattenSystemLocallogFortianalyzerFilterDocker(o["docker"], d, "docker")); err != nil {
		if vv, ok := fortiAPIPatch(o["docker"], "SystemLocallogFortianalyzerFilter-Docker"); ok {
			if err = d.Set("docker", vv); err != nil {
				return fmt.Errorf("Error reading docker: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading docker: %v", err)
		}
	}

	if err = d.Set("dvm", flattenSystemLocallogFortianalyzerFilterDvm(o["dvm"], d, "dvm")); err != nil {
		if vv, ok := fortiAPIPatch(o["dvm"], "SystemLocallogFortianalyzerFilter-Dvm"); ok {
			if err = d.Set("dvm", vv); err != nil {
				return fmt.Errorf("Error reading dvm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dvm: %v", err)
		}
	}

	if err = d.Set("ediscovery", flattenSystemLocallogFortianalyzerFilterEdiscovery(o["ediscovery"], d, "ediscovery")); err != nil {
		if vv, ok := fortiAPIPatch(o["ediscovery"], "SystemLocallogFortianalyzerFilter-Ediscovery"); ok {
			if err = d.Set("ediscovery", vv); err != nil {
				return fmt.Errorf("Error reading ediscovery: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ediscovery: %v", err)
		}
	}

	if err = d.Set("epmgr", flattenSystemLocallogFortianalyzerFilterEpmgr(o["epmgr"], d, "epmgr")); err != nil {
		if vv, ok := fortiAPIPatch(o["epmgr"], "SystemLocallogFortianalyzerFilter-Epmgr"); ok {
			if err = d.Set("epmgr", vv); err != nil {
				return fmt.Errorf("Error reading epmgr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading epmgr: %v", err)
		}
	}

	if err = d.Set("event", flattenSystemLocallogFortianalyzerFilterEvent(o["event"], d, "event")); err != nil {
		if vv, ok := fortiAPIPatch(o["event"], "SystemLocallogFortianalyzerFilter-Event"); ok {
			if err = d.Set("event", vv); err != nil {
				return fmt.Errorf("Error reading event: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading event: %v", err)
		}
	}

	if err = d.Set("eventmgmt", flattenSystemLocallogFortianalyzerFilterEventmgmt(o["eventmgmt"], d, "eventmgmt")); err != nil {
		if vv, ok := fortiAPIPatch(o["eventmgmt"], "SystemLocallogFortianalyzerFilter-Eventmgmt"); ok {
			if err = d.Set("eventmgmt", vv); err != nil {
				return fmt.Errorf("Error reading eventmgmt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eventmgmt: %v", err)
		}
	}

	if err = d.Set("faz", flattenSystemLocallogFortianalyzerFilterFaz(o["faz"], d, "faz")); err != nil {
		if vv, ok := fortiAPIPatch(o["faz"], "SystemLocallogFortianalyzerFilter-Faz"); ok {
			if err = d.Set("faz", vv); err != nil {
				return fmt.Errorf("Error reading faz: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading faz: %v", err)
		}
	}

	if err = d.Set("fazha", flattenSystemLocallogFortianalyzerFilterFazha(o["fazha"], d, "fazha")); err != nil {
		if vv, ok := fortiAPIPatch(o["fazha"], "SystemLocallogFortianalyzerFilter-Fazha"); ok {
			if err = d.Set("fazha", vv); err != nil {
				return fmt.Errorf("Error reading fazha: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fazha: %v", err)
		}
	}

	if err = d.Set("fazsys", flattenSystemLocallogFortianalyzerFilterFazsys(o["fazsys"], d, "fazsys")); err != nil {
		if vv, ok := fortiAPIPatch(o["fazsys"], "SystemLocallogFortianalyzerFilter-Fazsys"); ok {
			if err = d.Set("fazsys", vv); err != nil {
				return fmt.Errorf("Error reading fazsys: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fazsys: %v", err)
		}
	}

	if err = d.Set("fgd", flattenSystemLocallogFortianalyzerFilterFgd(o["fgd"], d, "fgd")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgd"], "SystemLocallogFortianalyzerFilter-Fgd"); ok {
			if err = d.Set("fgd", vv); err != nil {
				return fmt.Errorf("Error reading fgd: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgd: %v", err)
		}
	}

	if err = d.Set("fgfm", flattenSystemLocallogFortianalyzerFilterFgfm(o["fgfm"], d, "fgfm")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgfm"], "SystemLocallogFortianalyzerFilter-Fgfm"); ok {
			if err = d.Set("fgfm", vv); err != nil {
				return fmt.Errorf("Error reading fgfm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgfm: %v", err)
		}
	}

	if err = d.Set("fips", flattenSystemLocallogFortianalyzerFilterFips(o["fips"], d, "fips")); err != nil {
		if vv, ok := fortiAPIPatch(o["fips"], "SystemLocallogFortianalyzerFilter-Fips"); ok {
			if err = d.Set("fips", vv); err != nil {
				return fmt.Errorf("Error reading fips: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fips: %v", err)
		}
	}

	if err = d.Set("fmgws", flattenSystemLocallogFortianalyzerFilterFmgws(o["fmgws"], d, "fmgws")); err != nil {
		if vv, ok := fortiAPIPatch(o["fmgws"], "SystemLocallogFortianalyzerFilter-Fmgws"); ok {
			if err = d.Set("fmgws", vv); err != nil {
				return fmt.Errorf("Error reading fmgws: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmgws: %v", err)
		}
	}

	if err = d.Set("fmlmgr", flattenSystemLocallogFortianalyzerFilterFmlmgr(o["fmlmgr"], d, "fmlmgr")); err != nil {
		if vv, ok := fortiAPIPatch(o["fmlmgr"], "SystemLocallogFortianalyzerFilter-Fmlmgr"); ok {
			if err = d.Set("fmlmgr", vv); err != nil {
				return fmt.Errorf("Error reading fmlmgr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmlmgr: %v", err)
		}
	}

	if err = d.Set("fmwmgr", flattenSystemLocallogFortianalyzerFilterFmwmgr(o["fmwmgr"], d, "fmwmgr")); err != nil {
		if vv, ok := fortiAPIPatch(o["fmwmgr"], "SystemLocallogFortianalyzerFilter-Fmwmgr"); ok {
			if err = d.Set("fmwmgr", vv); err != nil {
				return fmt.Errorf("Error reading fmwmgr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmwmgr: %v", err)
		}
	}

	if err = d.Set("fortiview", flattenSystemLocallogFortianalyzerFilterFortiview(o["fortiview"], d, "fortiview")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortiview"], "SystemLocallogFortianalyzerFilter-Fortiview"); ok {
			if err = d.Set("fortiview", vv); err != nil {
				return fmt.Errorf("Error reading fortiview: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortiview: %v", err)
		}
	}

	if err = d.Set("glbcfg", flattenSystemLocallogFortianalyzerFilterGlbcfg(o["glbcfg"], d, "glbcfg")); err != nil {
		if vv, ok := fortiAPIPatch(o["glbcfg"], "SystemLocallogFortianalyzerFilter-Glbcfg"); ok {
			if err = d.Set("glbcfg", vv); err != nil {
				return fmt.Errorf("Error reading glbcfg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading glbcfg: %v", err)
		}
	}

	if err = d.Set("ha", flattenSystemLocallogFortianalyzerFilterHa(o["ha"], d, "ha")); err != nil {
		if vv, ok := fortiAPIPatch(o["ha"], "SystemLocallogFortianalyzerFilter-Ha"); ok {
			if err = d.Set("ha", vv); err != nil {
				return fmt.Errorf("Error reading ha: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ha: %v", err)
		}
	}

	if err = d.Set("hcache", flattenSystemLocallogFortianalyzerFilterHcache(o["hcache"], d, "hcache")); err != nil {
		if vv, ok := fortiAPIPatch(o["hcache"], "SystemLocallogFortianalyzerFilter-Hcache"); ok {
			if err = d.Set("hcache", vv); err != nil {
				return fmt.Errorf("Error reading hcache: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hcache: %v", err)
		}
	}

	if err = d.Set("incident", flattenSystemLocallogFortianalyzerFilterIncident(o["incident"], d, "incident")); err != nil {
		if vv, ok := fortiAPIPatch(o["incident"], "SystemLocallogFortianalyzerFilter-Incident"); ok {
			if err = d.Set("incident", vv); err != nil {
				return fmt.Errorf("Error reading incident: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading incident: %v", err)
		}
	}

	if err = d.Set("iolog", flattenSystemLocallogFortianalyzerFilterIolog(o["iolog"], d, "iolog")); err != nil {
		if vv, ok := fortiAPIPatch(o["iolog"], "SystemLocallogFortianalyzerFilter-Iolog"); ok {
			if err = d.Set("iolog", vv); err != nil {
				return fmt.Errorf("Error reading iolog: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading iolog: %v", err)
		}
	}

	if err = d.Set("logd", flattenSystemLocallogFortianalyzerFilterLogd(o["logd"], d, "logd")); err != nil {
		if vv, ok := fortiAPIPatch(o["logd"], "SystemLocallogFortianalyzerFilter-Logd"); ok {
			if err = d.Set("logd", vv); err != nil {
				return fmt.Errorf("Error reading logd: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logd: %v", err)
		}
	}

	if err = d.Set("logdb", flattenSystemLocallogFortianalyzerFilterLogdb(o["logdb"], d, "logdb")); err != nil {
		if vv, ok := fortiAPIPatch(o["logdb"], "SystemLocallogFortianalyzerFilter-Logdb"); ok {
			if err = d.Set("logdb", vv); err != nil {
				return fmt.Errorf("Error reading logdb: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logdb: %v", err)
		}
	}

	if err = d.Set("logdev", flattenSystemLocallogFortianalyzerFilterLogdev(o["logdev"], d, "logdev")); err != nil {
		if vv, ok := fortiAPIPatch(o["logdev"], "SystemLocallogFortianalyzerFilter-Logdev"); ok {
			if err = d.Set("logdev", vv); err != nil {
				return fmt.Errorf("Error reading logdev: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logdev: %v", err)
		}
	}

	if err = d.Set("logfile", flattenSystemLocallogFortianalyzerFilterLogfile(o["logfile"], d, "logfile")); err != nil {
		if vv, ok := fortiAPIPatch(o["logfile"], "SystemLocallogFortianalyzerFilter-Logfile"); ok {
			if err = d.Set("logfile", vv); err != nil {
				return fmt.Errorf("Error reading logfile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logfile: %v", err)
		}
	}

	if err = d.Set("logging", flattenSystemLocallogFortianalyzerFilterLogging(o["logging"], d, "logging")); err != nil {
		if vv, ok := fortiAPIPatch(o["logging"], "SystemLocallogFortianalyzerFilter-Logging"); ok {
			if err = d.Set("logging", vv); err != nil {
				return fmt.Errorf("Error reading logging: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logging: %v", err)
		}
	}

	if err = d.Set("lrmgr", flattenSystemLocallogFortianalyzerFilterLrmgr(o["lrmgr"], d, "lrmgr")); err != nil {
		if vv, ok := fortiAPIPatch(o["lrmgr"], "SystemLocallogFortianalyzerFilter-Lrmgr"); ok {
			if err = d.Set("lrmgr", vv); err != nil {
				return fmt.Errorf("Error reading lrmgr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lrmgr: %v", err)
		}
	}

	if err = d.Set("objcfg", flattenSystemLocallogFortianalyzerFilterObjcfg(o["objcfg"], d, "objcfg")); err != nil {
		if vv, ok := fortiAPIPatch(o["objcfg"], "SystemLocallogFortianalyzerFilter-Objcfg"); ok {
			if err = d.Set("objcfg", vv); err != nil {
				return fmt.Errorf("Error reading objcfg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading objcfg: %v", err)
		}
	}

	if err = d.Set("report", flattenSystemLocallogFortianalyzerFilterReport(o["report"], d, "report")); err != nil {
		if vv, ok := fortiAPIPatch(o["report"], "SystemLocallogFortianalyzerFilter-Report"); ok {
			if err = d.Set("report", vv); err != nil {
				return fmt.Errorf("Error reading report: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading report: %v", err)
		}
	}

	if err = d.Set("rev", flattenSystemLocallogFortianalyzerFilterRev(o["rev"], d, "rev")); err != nil {
		if vv, ok := fortiAPIPatch(o["rev"], "SystemLocallogFortianalyzerFilter-Rev"); ok {
			if err = d.Set("rev", vv); err != nil {
				return fmt.Errorf("Error reading rev: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rev: %v", err)
		}
	}

	if err = d.Set("rtmon", flattenSystemLocallogFortianalyzerFilterRtmon(o["rtmon"], d, "rtmon")); err != nil {
		if vv, ok := fortiAPIPatch(o["rtmon"], "SystemLocallogFortianalyzerFilter-Rtmon"); ok {
			if err = d.Set("rtmon", vv); err != nil {
				return fmt.Errorf("Error reading rtmon: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rtmon: %v", err)
		}
	}

	if err = d.Set("scfw", flattenSystemLocallogFortianalyzerFilterScfw(o["scfw"], d, "scfw")); err != nil {
		if vv, ok := fortiAPIPatch(o["scfw"], "SystemLocallogFortianalyzerFilter-Scfw"); ok {
			if err = d.Set("scfw", vv); err != nil {
				return fmt.Errorf("Error reading scfw: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scfw: %v", err)
		}
	}

	if err = d.Set("scply", flattenSystemLocallogFortianalyzerFilterScply(o["scply"], d, "scply")); err != nil {
		if vv, ok := fortiAPIPatch(o["scply"], "SystemLocallogFortianalyzerFilter-Scply"); ok {
			if err = d.Set("scply", vv); err != nil {
				return fmt.Errorf("Error reading scply: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scply: %v", err)
		}
	}

	if err = d.Set("scrmgr", flattenSystemLocallogFortianalyzerFilterScrmgr(o["scrmgr"], d, "scrmgr")); err != nil {
		if vv, ok := fortiAPIPatch(o["scrmgr"], "SystemLocallogFortianalyzerFilter-Scrmgr"); ok {
			if err = d.Set("scrmgr", vv); err != nil {
				return fmt.Errorf("Error reading scrmgr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scrmgr: %v", err)
		}
	}

	if err = d.Set("scvpn", flattenSystemLocallogFortianalyzerFilterScvpn(o["scvpn"], d, "scvpn")); err != nil {
		if vv, ok := fortiAPIPatch(o["scvpn"], "SystemLocallogFortianalyzerFilter-Scvpn"); ok {
			if err = d.Set("scvpn", vv); err != nil {
				return fmt.Errorf("Error reading scvpn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scvpn: %v", err)
		}
	}

	if err = d.Set("system", flattenSystemLocallogFortianalyzerFilterSystem(o["system"], d, "system")); err != nil {
		if vv, ok := fortiAPIPatch(o["system"], "SystemLocallogFortianalyzerFilter-System"); ok {
			if err = d.Set("system", vv); err != nil {
				return fmt.Errorf("Error reading system: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading system: %v", err)
		}
	}

	if err = d.Set("webport", flattenSystemLocallogFortianalyzerFilterWebport(o["webport"], d, "webport")); err != nil {
		if vv, ok := fortiAPIPatch(o["webport"], "SystemLocallogFortianalyzerFilter-Webport"); ok {
			if err = d.Set("webport", vv); err != nil {
				return fmt.Errorf("Error reading webport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webport: %v", err)
		}
	}

	return nil
}

func flattenSystemLocallogFortianalyzerFilterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemLocallogFortianalyzerFilterAid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzerFilterDevcfg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzerFilterDevops(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzerFilterDiskquota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzerFilterDm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzerFilterDocker(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzerFilterDvm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzerFilterEdiscovery(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzerFilterEpmgr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzerFilterEvent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzerFilterEventmgmt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzerFilterFaz(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzerFilterFazha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzerFilterFazsys(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzerFilterFgd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzerFilterFgfm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzerFilterFips(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzerFilterFmgws(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzerFilterFmlmgr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzerFilterFmwmgr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzerFilterFortiview(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzerFilterGlbcfg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzerFilterHa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzerFilterHcache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzerFilterIncident(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzerFilterIolog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzerFilterLogd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzerFilterLogdb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzerFilterLogdev(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzerFilterLogfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzerFilterLogging(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzerFilterLrmgr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzerFilterObjcfg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzerFilterReport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzerFilterRev(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzerFilterRtmon(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzerFilterScfw(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzerFilterScply(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzerFilterScrmgr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzerFilterScvpn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzerFilterSystem(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzerFilterWebport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemLocallogFortianalyzerFilter(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("aid"); ok || d.HasChange("aid") {
		t, err := expandSystemLocallogFortianalyzerFilterAid(d, v, "aid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aid"] = t
		}
	}

	if v, ok := d.GetOk("devcfg"); ok || d.HasChange("devcfg") {
		t, err := expandSystemLocallogFortianalyzerFilterDevcfg(d, v, "devcfg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["devcfg"] = t
		}
	}

	if v, ok := d.GetOk("devops"); ok || d.HasChange("devops") {
		t, err := expandSystemLocallogFortianalyzerFilterDevops(d, v, "devops")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["devops"] = t
		}
	}

	if v, ok := d.GetOk("diskquota"); ok || d.HasChange("diskquota") {
		t, err := expandSystemLocallogFortianalyzerFilterDiskquota(d, v, "diskquota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diskquota"] = t
		}
	}

	if v, ok := d.GetOk("dm"); ok || d.HasChange("dm") {
		t, err := expandSystemLocallogFortianalyzerFilterDm(d, v, "dm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dm"] = t
		}
	}

	if v, ok := d.GetOk("docker"); ok || d.HasChange("docker") {
		t, err := expandSystemLocallogFortianalyzerFilterDocker(d, v, "docker")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["docker"] = t
		}
	}

	if v, ok := d.GetOk("dvm"); ok || d.HasChange("dvm") {
		t, err := expandSystemLocallogFortianalyzerFilterDvm(d, v, "dvm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dvm"] = t
		}
	}

	if v, ok := d.GetOk("ediscovery"); ok || d.HasChange("ediscovery") {
		t, err := expandSystemLocallogFortianalyzerFilterEdiscovery(d, v, "ediscovery")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ediscovery"] = t
		}
	}

	if v, ok := d.GetOk("epmgr"); ok || d.HasChange("epmgr") {
		t, err := expandSystemLocallogFortianalyzerFilterEpmgr(d, v, "epmgr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["epmgr"] = t
		}
	}

	if v, ok := d.GetOk("event"); ok || d.HasChange("event") {
		t, err := expandSystemLocallogFortianalyzerFilterEvent(d, v, "event")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["event"] = t
		}
	}

	if v, ok := d.GetOk("eventmgmt"); ok || d.HasChange("eventmgmt") {
		t, err := expandSystemLocallogFortianalyzerFilterEventmgmt(d, v, "eventmgmt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eventmgmt"] = t
		}
	}

	if v, ok := d.GetOk("faz"); ok || d.HasChange("faz") {
		t, err := expandSystemLocallogFortianalyzerFilterFaz(d, v, "faz")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["faz"] = t
		}
	}

	if v, ok := d.GetOk("fazha"); ok || d.HasChange("fazha") {
		t, err := expandSystemLocallogFortianalyzerFilterFazha(d, v, "fazha")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fazha"] = t
		}
	}

	if v, ok := d.GetOk("fazsys"); ok || d.HasChange("fazsys") {
		t, err := expandSystemLocallogFortianalyzerFilterFazsys(d, v, "fazsys")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fazsys"] = t
		}
	}

	if v, ok := d.GetOk("fgd"); ok || d.HasChange("fgd") {
		t, err := expandSystemLocallogFortianalyzerFilterFgd(d, v, "fgd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgd"] = t
		}
	}

	if v, ok := d.GetOk("fgfm"); ok || d.HasChange("fgfm") {
		t, err := expandSystemLocallogFortianalyzerFilterFgfm(d, v, "fgfm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgfm"] = t
		}
	}

	if v, ok := d.GetOk("fips"); ok || d.HasChange("fips") {
		t, err := expandSystemLocallogFortianalyzerFilterFips(d, v, "fips")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fips"] = t
		}
	}

	if v, ok := d.GetOk("fmgws"); ok || d.HasChange("fmgws") {
		t, err := expandSystemLocallogFortianalyzerFilterFmgws(d, v, "fmgws")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fmgws"] = t
		}
	}

	if v, ok := d.GetOk("fmlmgr"); ok || d.HasChange("fmlmgr") {
		t, err := expandSystemLocallogFortianalyzerFilterFmlmgr(d, v, "fmlmgr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fmlmgr"] = t
		}
	}

	if v, ok := d.GetOk("fmwmgr"); ok || d.HasChange("fmwmgr") {
		t, err := expandSystemLocallogFortianalyzerFilterFmwmgr(d, v, "fmwmgr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fmwmgr"] = t
		}
	}

	if v, ok := d.GetOk("fortiview"); ok || d.HasChange("fortiview") {
		t, err := expandSystemLocallogFortianalyzerFilterFortiview(d, v, "fortiview")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiview"] = t
		}
	}

	if v, ok := d.GetOk("glbcfg"); ok || d.HasChange("glbcfg") {
		t, err := expandSystemLocallogFortianalyzerFilterGlbcfg(d, v, "glbcfg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["glbcfg"] = t
		}
	}

	if v, ok := d.GetOk("ha"); ok || d.HasChange("ha") {
		t, err := expandSystemLocallogFortianalyzerFilterHa(d, v, "ha")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha"] = t
		}
	}

	if v, ok := d.GetOk("hcache"); ok || d.HasChange("hcache") {
		t, err := expandSystemLocallogFortianalyzerFilterHcache(d, v, "hcache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hcache"] = t
		}
	}

	if v, ok := d.GetOk("incident"); ok || d.HasChange("incident") {
		t, err := expandSystemLocallogFortianalyzerFilterIncident(d, v, "incident")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["incident"] = t
		}
	}

	if v, ok := d.GetOk("iolog"); ok || d.HasChange("iolog") {
		t, err := expandSystemLocallogFortianalyzerFilterIolog(d, v, "iolog")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["iolog"] = t
		}
	}

	if v, ok := d.GetOk("logd"); ok || d.HasChange("logd") {
		t, err := expandSystemLocallogFortianalyzerFilterLogd(d, v, "logd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logd"] = t
		}
	}

	if v, ok := d.GetOk("logdb"); ok || d.HasChange("logdb") {
		t, err := expandSystemLocallogFortianalyzerFilterLogdb(d, v, "logdb")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logdb"] = t
		}
	}

	if v, ok := d.GetOk("logdev"); ok || d.HasChange("logdev") {
		t, err := expandSystemLocallogFortianalyzerFilterLogdev(d, v, "logdev")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logdev"] = t
		}
	}

	if v, ok := d.GetOk("logfile"); ok || d.HasChange("logfile") {
		t, err := expandSystemLocallogFortianalyzerFilterLogfile(d, v, "logfile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logfile"] = t
		}
	}

	if v, ok := d.GetOk("logging"); ok || d.HasChange("logging") {
		t, err := expandSystemLocallogFortianalyzerFilterLogging(d, v, "logging")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logging"] = t
		}
	}

	if v, ok := d.GetOk("lrmgr"); ok || d.HasChange("lrmgr") {
		t, err := expandSystemLocallogFortianalyzerFilterLrmgr(d, v, "lrmgr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lrmgr"] = t
		}
	}

	if v, ok := d.GetOk("objcfg"); ok || d.HasChange("objcfg") {
		t, err := expandSystemLocallogFortianalyzerFilterObjcfg(d, v, "objcfg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["objcfg"] = t
		}
	}

	if v, ok := d.GetOk("report"); ok || d.HasChange("report") {
		t, err := expandSystemLocallogFortianalyzerFilterReport(d, v, "report")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["report"] = t
		}
	}

	if v, ok := d.GetOk("rev"); ok || d.HasChange("rev") {
		t, err := expandSystemLocallogFortianalyzerFilterRev(d, v, "rev")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rev"] = t
		}
	}

	if v, ok := d.GetOk("rtmon"); ok || d.HasChange("rtmon") {
		t, err := expandSystemLocallogFortianalyzerFilterRtmon(d, v, "rtmon")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rtmon"] = t
		}
	}

	if v, ok := d.GetOk("scfw"); ok || d.HasChange("scfw") {
		t, err := expandSystemLocallogFortianalyzerFilterScfw(d, v, "scfw")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scfw"] = t
		}
	}

	if v, ok := d.GetOk("scply"); ok || d.HasChange("scply") {
		t, err := expandSystemLocallogFortianalyzerFilterScply(d, v, "scply")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scply"] = t
		}
	}

	if v, ok := d.GetOk("scrmgr"); ok || d.HasChange("scrmgr") {
		t, err := expandSystemLocallogFortianalyzerFilterScrmgr(d, v, "scrmgr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scrmgr"] = t
		}
	}

	if v, ok := d.GetOk("scvpn"); ok || d.HasChange("scvpn") {
		t, err := expandSystemLocallogFortianalyzerFilterScvpn(d, v, "scvpn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scvpn"] = t
		}
	}

	if v, ok := d.GetOk("system"); ok || d.HasChange("system") {
		t, err := expandSystemLocallogFortianalyzerFilterSystem(d, v, "system")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["system"] = t
		}
	}

	if v, ok := d.GetOk("webport"); ok || d.HasChange("webport") {
		t, err := expandSystemLocallogFortianalyzerFilterWebport(d, v, "webport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webport"] = t
		}
	}

	return &obj, nil
}
