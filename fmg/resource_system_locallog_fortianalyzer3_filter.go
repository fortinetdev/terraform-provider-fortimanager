// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Filter for FortiAnalyzer3 logging.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemLocallogFortianalyzer3Filter() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLocallogFortianalyzer3FilterUpdate,
		Read:   resourceSystemLocallogFortianalyzer3FilterRead,
		Update: resourceSystemLocallogFortianalyzer3FilterUpdate,
		Delete: resourceSystemLocallogFortianalyzer3FilterDelete,

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

func resourceSystemLocallogFortianalyzer3FilterUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	obj, err := getObjectSystemLocallogFortianalyzer3Filter(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemLocallogFortianalyzer3Filter resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemLocallogFortianalyzer3Filter(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating SystemLocallogFortianalyzer3Filter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("SystemLocallogFortianalyzer3Filter")

	return resourceSystemLocallogFortianalyzer3FilterRead(d, m)
}

func resourceSystemLocallogFortianalyzer3FilterDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	err = c.DeleteSystemLocallogFortianalyzer3Filter(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting SystemLocallogFortianalyzer3Filter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLocallogFortianalyzer3FilterRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	adomv, err := "global", fmt.Errorf("")

	o, err := c.ReadSystemLocallogFortianalyzer3Filter(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading SystemLocallogFortianalyzer3Filter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemLocallogFortianalyzer3Filter(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemLocallogFortianalyzer3Filter resource from API: %v", err)
	}
	return nil
}

func flattenSystemLocallogFortianalyzer3FilterAid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzer3FilterDevcfg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzer3FilterDevops(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzer3FilterDiskquota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzer3FilterDm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzer3FilterDocker(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzer3FilterDvm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzer3FilterEdiscovery(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzer3FilterEpmgr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzer3FilterEvent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzer3FilterEventmgmt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzer3FilterFaz(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzer3FilterFazha(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzer3FilterFazsys(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzer3FilterFgd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzer3FilterFgfm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzer3FilterFips(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzer3FilterFmgws(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzer3FilterFmlmgr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzer3FilterFmwmgr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzer3FilterFortiview(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzer3FilterGlbcfg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzer3FilterHa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzer3FilterHcache(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzer3FilterIncident(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzer3FilterIolog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzer3FilterLogd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzer3FilterLogdb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzer3FilterLogdev(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzer3FilterLogfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzer3FilterLogging(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzer3FilterLrmgr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzer3FilterObjcfg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzer3FilterReport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzer3FilterRev(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzer3FilterRtmon(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzer3FilterScfw(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzer3FilterScply(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzer3FilterScrmgr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzer3FilterScvpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzer3FilterSystem(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLocallogFortianalyzer3FilterWebport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemLocallogFortianalyzer3Filter(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("aid", flattenSystemLocallogFortianalyzer3FilterAid(o["aid"], d, "aid")); err != nil {
		if vv, ok := fortiAPIPatch(o["aid"], "SystemLocallogFortianalyzer3Filter-Aid"); ok {
			if err = d.Set("aid", vv); err != nil {
				return fmt.Errorf("Error reading aid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading aid: %v", err)
		}
	}

	if err = d.Set("devcfg", flattenSystemLocallogFortianalyzer3FilterDevcfg(o["devcfg"], d, "devcfg")); err != nil {
		if vv, ok := fortiAPIPatch(o["devcfg"], "SystemLocallogFortianalyzer3Filter-Devcfg"); ok {
			if err = d.Set("devcfg", vv); err != nil {
				return fmt.Errorf("Error reading devcfg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading devcfg: %v", err)
		}
	}

	if err = d.Set("devops", flattenSystemLocallogFortianalyzer3FilterDevops(o["devops"], d, "devops")); err != nil {
		if vv, ok := fortiAPIPatch(o["devops"], "SystemLocallogFortianalyzer3Filter-Devops"); ok {
			if err = d.Set("devops", vv); err != nil {
				return fmt.Errorf("Error reading devops: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading devops: %v", err)
		}
	}

	if err = d.Set("diskquota", flattenSystemLocallogFortianalyzer3FilterDiskquota(o["diskquota"], d, "diskquota")); err != nil {
		if vv, ok := fortiAPIPatch(o["diskquota"], "SystemLocallogFortianalyzer3Filter-Diskquota"); ok {
			if err = d.Set("diskquota", vv); err != nil {
				return fmt.Errorf("Error reading diskquota: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diskquota: %v", err)
		}
	}

	if err = d.Set("dm", flattenSystemLocallogFortianalyzer3FilterDm(o["dm"], d, "dm")); err != nil {
		if vv, ok := fortiAPIPatch(o["dm"], "SystemLocallogFortianalyzer3Filter-Dm"); ok {
			if err = d.Set("dm", vv); err != nil {
				return fmt.Errorf("Error reading dm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dm: %v", err)
		}
	}

	if err = d.Set("docker", flattenSystemLocallogFortianalyzer3FilterDocker(o["docker"], d, "docker")); err != nil {
		if vv, ok := fortiAPIPatch(o["docker"], "SystemLocallogFortianalyzer3Filter-Docker"); ok {
			if err = d.Set("docker", vv); err != nil {
				return fmt.Errorf("Error reading docker: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading docker: %v", err)
		}
	}

	if err = d.Set("dvm", flattenSystemLocallogFortianalyzer3FilterDvm(o["dvm"], d, "dvm")); err != nil {
		if vv, ok := fortiAPIPatch(o["dvm"], "SystemLocallogFortianalyzer3Filter-Dvm"); ok {
			if err = d.Set("dvm", vv); err != nil {
				return fmt.Errorf("Error reading dvm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dvm: %v", err)
		}
	}

	if err = d.Set("ediscovery", flattenSystemLocallogFortianalyzer3FilterEdiscovery(o["ediscovery"], d, "ediscovery")); err != nil {
		if vv, ok := fortiAPIPatch(o["ediscovery"], "SystemLocallogFortianalyzer3Filter-Ediscovery"); ok {
			if err = d.Set("ediscovery", vv); err != nil {
				return fmt.Errorf("Error reading ediscovery: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ediscovery: %v", err)
		}
	}

	if err = d.Set("epmgr", flattenSystemLocallogFortianalyzer3FilterEpmgr(o["epmgr"], d, "epmgr")); err != nil {
		if vv, ok := fortiAPIPatch(o["epmgr"], "SystemLocallogFortianalyzer3Filter-Epmgr"); ok {
			if err = d.Set("epmgr", vv); err != nil {
				return fmt.Errorf("Error reading epmgr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading epmgr: %v", err)
		}
	}

	if err = d.Set("event", flattenSystemLocallogFortianalyzer3FilterEvent(o["event"], d, "event")); err != nil {
		if vv, ok := fortiAPIPatch(o["event"], "SystemLocallogFortianalyzer3Filter-Event"); ok {
			if err = d.Set("event", vv); err != nil {
				return fmt.Errorf("Error reading event: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading event: %v", err)
		}
	}

	if err = d.Set("eventmgmt", flattenSystemLocallogFortianalyzer3FilterEventmgmt(o["eventmgmt"], d, "eventmgmt")); err != nil {
		if vv, ok := fortiAPIPatch(o["eventmgmt"], "SystemLocallogFortianalyzer3Filter-Eventmgmt"); ok {
			if err = d.Set("eventmgmt", vv); err != nil {
				return fmt.Errorf("Error reading eventmgmt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading eventmgmt: %v", err)
		}
	}

	if err = d.Set("faz", flattenSystemLocallogFortianalyzer3FilterFaz(o["faz"], d, "faz")); err != nil {
		if vv, ok := fortiAPIPatch(o["faz"], "SystemLocallogFortianalyzer3Filter-Faz"); ok {
			if err = d.Set("faz", vv); err != nil {
				return fmt.Errorf("Error reading faz: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading faz: %v", err)
		}
	}

	if err = d.Set("fazha", flattenSystemLocallogFortianalyzer3FilterFazha(o["fazha"], d, "fazha")); err != nil {
		if vv, ok := fortiAPIPatch(o["fazha"], "SystemLocallogFortianalyzer3Filter-Fazha"); ok {
			if err = d.Set("fazha", vv); err != nil {
				return fmt.Errorf("Error reading fazha: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fazha: %v", err)
		}
	}

	if err = d.Set("fazsys", flattenSystemLocallogFortianalyzer3FilterFazsys(o["fazsys"], d, "fazsys")); err != nil {
		if vv, ok := fortiAPIPatch(o["fazsys"], "SystemLocallogFortianalyzer3Filter-Fazsys"); ok {
			if err = d.Set("fazsys", vv); err != nil {
				return fmt.Errorf("Error reading fazsys: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fazsys: %v", err)
		}
	}

	if err = d.Set("fgd", flattenSystemLocallogFortianalyzer3FilterFgd(o["fgd"], d, "fgd")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgd"], "SystemLocallogFortianalyzer3Filter-Fgd"); ok {
			if err = d.Set("fgd", vv); err != nil {
				return fmt.Errorf("Error reading fgd: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgd: %v", err)
		}
	}

	if err = d.Set("fgfm", flattenSystemLocallogFortianalyzer3FilterFgfm(o["fgfm"], d, "fgfm")); err != nil {
		if vv, ok := fortiAPIPatch(o["fgfm"], "SystemLocallogFortianalyzer3Filter-Fgfm"); ok {
			if err = d.Set("fgfm", vv); err != nil {
				return fmt.Errorf("Error reading fgfm: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fgfm: %v", err)
		}
	}

	if err = d.Set("fips", flattenSystemLocallogFortianalyzer3FilterFips(o["fips"], d, "fips")); err != nil {
		if vv, ok := fortiAPIPatch(o["fips"], "SystemLocallogFortianalyzer3Filter-Fips"); ok {
			if err = d.Set("fips", vv); err != nil {
				return fmt.Errorf("Error reading fips: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fips: %v", err)
		}
	}

	if err = d.Set("fmgws", flattenSystemLocallogFortianalyzer3FilterFmgws(o["fmgws"], d, "fmgws")); err != nil {
		if vv, ok := fortiAPIPatch(o["fmgws"], "SystemLocallogFortianalyzer3Filter-Fmgws"); ok {
			if err = d.Set("fmgws", vv); err != nil {
				return fmt.Errorf("Error reading fmgws: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmgws: %v", err)
		}
	}

	if err = d.Set("fmlmgr", flattenSystemLocallogFortianalyzer3FilterFmlmgr(o["fmlmgr"], d, "fmlmgr")); err != nil {
		if vv, ok := fortiAPIPatch(o["fmlmgr"], "SystemLocallogFortianalyzer3Filter-Fmlmgr"); ok {
			if err = d.Set("fmlmgr", vv); err != nil {
				return fmt.Errorf("Error reading fmlmgr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmlmgr: %v", err)
		}
	}

	if err = d.Set("fmwmgr", flattenSystemLocallogFortianalyzer3FilterFmwmgr(o["fmwmgr"], d, "fmwmgr")); err != nil {
		if vv, ok := fortiAPIPatch(o["fmwmgr"], "SystemLocallogFortianalyzer3Filter-Fmwmgr"); ok {
			if err = d.Set("fmwmgr", vv); err != nil {
				return fmt.Errorf("Error reading fmwmgr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fmwmgr: %v", err)
		}
	}

	if err = d.Set("fortiview", flattenSystemLocallogFortianalyzer3FilterFortiview(o["fortiview"], d, "fortiview")); err != nil {
		if vv, ok := fortiAPIPatch(o["fortiview"], "SystemLocallogFortianalyzer3Filter-Fortiview"); ok {
			if err = d.Set("fortiview", vv); err != nil {
				return fmt.Errorf("Error reading fortiview: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fortiview: %v", err)
		}
	}

	if err = d.Set("glbcfg", flattenSystemLocallogFortianalyzer3FilterGlbcfg(o["glbcfg"], d, "glbcfg")); err != nil {
		if vv, ok := fortiAPIPatch(o["glbcfg"], "SystemLocallogFortianalyzer3Filter-Glbcfg"); ok {
			if err = d.Set("glbcfg", vv); err != nil {
				return fmt.Errorf("Error reading glbcfg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading glbcfg: %v", err)
		}
	}

	if err = d.Set("ha", flattenSystemLocallogFortianalyzer3FilterHa(o["ha"], d, "ha")); err != nil {
		if vv, ok := fortiAPIPatch(o["ha"], "SystemLocallogFortianalyzer3Filter-Ha"); ok {
			if err = d.Set("ha", vv); err != nil {
				return fmt.Errorf("Error reading ha: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ha: %v", err)
		}
	}

	if err = d.Set("hcache", flattenSystemLocallogFortianalyzer3FilterHcache(o["hcache"], d, "hcache")); err != nil {
		if vv, ok := fortiAPIPatch(o["hcache"], "SystemLocallogFortianalyzer3Filter-Hcache"); ok {
			if err = d.Set("hcache", vv); err != nil {
				return fmt.Errorf("Error reading hcache: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading hcache: %v", err)
		}
	}

	if err = d.Set("incident", flattenSystemLocallogFortianalyzer3FilterIncident(o["incident"], d, "incident")); err != nil {
		if vv, ok := fortiAPIPatch(o["incident"], "SystemLocallogFortianalyzer3Filter-Incident"); ok {
			if err = d.Set("incident", vv); err != nil {
				return fmt.Errorf("Error reading incident: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading incident: %v", err)
		}
	}

	if err = d.Set("iolog", flattenSystemLocallogFortianalyzer3FilterIolog(o["iolog"], d, "iolog")); err != nil {
		if vv, ok := fortiAPIPatch(o["iolog"], "SystemLocallogFortianalyzer3Filter-Iolog"); ok {
			if err = d.Set("iolog", vv); err != nil {
				return fmt.Errorf("Error reading iolog: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading iolog: %v", err)
		}
	}

	if err = d.Set("logd", flattenSystemLocallogFortianalyzer3FilterLogd(o["logd"], d, "logd")); err != nil {
		if vv, ok := fortiAPIPatch(o["logd"], "SystemLocallogFortianalyzer3Filter-Logd"); ok {
			if err = d.Set("logd", vv); err != nil {
				return fmt.Errorf("Error reading logd: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logd: %v", err)
		}
	}

	if err = d.Set("logdb", flattenSystemLocallogFortianalyzer3FilterLogdb(o["logdb"], d, "logdb")); err != nil {
		if vv, ok := fortiAPIPatch(o["logdb"], "SystemLocallogFortianalyzer3Filter-Logdb"); ok {
			if err = d.Set("logdb", vv); err != nil {
				return fmt.Errorf("Error reading logdb: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logdb: %v", err)
		}
	}

	if err = d.Set("logdev", flattenSystemLocallogFortianalyzer3FilterLogdev(o["logdev"], d, "logdev")); err != nil {
		if vv, ok := fortiAPIPatch(o["logdev"], "SystemLocallogFortianalyzer3Filter-Logdev"); ok {
			if err = d.Set("logdev", vv); err != nil {
				return fmt.Errorf("Error reading logdev: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logdev: %v", err)
		}
	}

	if err = d.Set("logfile", flattenSystemLocallogFortianalyzer3FilterLogfile(o["logfile"], d, "logfile")); err != nil {
		if vv, ok := fortiAPIPatch(o["logfile"], "SystemLocallogFortianalyzer3Filter-Logfile"); ok {
			if err = d.Set("logfile", vv); err != nil {
				return fmt.Errorf("Error reading logfile: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logfile: %v", err)
		}
	}

	if err = d.Set("logging", flattenSystemLocallogFortianalyzer3FilterLogging(o["logging"], d, "logging")); err != nil {
		if vv, ok := fortiAPIPatch(o["logging"], "SystemLocallogFortianalyzer3Filter-Logging"); ok {
			if err = d.Set("logging", vv); err != nil {
				return fmt.Errorf("Error reading logging: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logging: %v", err)
		}
	}

	if err = d.Set("lrmgr", flattenSystemLocallogFortianalyzer3FilterLrmgr(o["lrmgr"], d, "lrmgr")); err != nil {
		if vv, ok := fortiAPIPatch(o["lrmgr"], "SystemLocallogFortianalyzer3Filter-Lrmgr"); ok {
			if err = d.Set("lrmgr", vv); err != nil {
				return fmt.Errorf("Error reading lrmgr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading lrmgr: %v", err)
		}
	}

	if err = d.Set("objcfg", flattenSystemLocallogFortianalyzer3FilterObjcfg(o["objcfg"], d, "objcfg")); err != nil {
		if vv, ok := fortiAPIPatch(o["objcfg"], "SystemLocallogFortianalyzer3Filter-Objcfg"); ok {
			if err = d.Set("objcfg", vv); err != nil {
				return fmt.Errorf("Error reading objcfg: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading objcfg: %v", err)
		}
	}

	if err = d.Set("report", flattenSystemLocallogFortianalyzer3FilterReport(o["report"], d, "report")); err != nil {
		if vv, ok := fortiAPIPatch(o["report"], "SystemLocallogFortianalyzer3Filter-Report"); ok {
			if err = d.Set("report", vv); err != nil {
				return fmt.Errorf("Error reading report: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading report: %v", err)
		}
	}

	if err = d.Set("rev", flattenSystemLocallogFortianalyzer3FilterRev(o["rev"], d, "rev")); err != nil {
		if vv, ok := fortiAPIPatch(o["rev"], "SystemLocallogFortianalyzer3Filter-Rev"); ok {
			if err = d.Set("rev", vv); err != nil {
				return fmt.Errorf("Error reading rev: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rev: %v", err)
		}
	}

	if err = d.Set("rtmon", flattenSystemLocallogFortianalyzer3FilterRtmon(o["rtmon"], d, "rtmon")); err != nil {
		if vv, ok := fortiAPIPatch(o["rtmon"], "SystemLocallogFortianalyzer3Filter-Rtmon"); ok {
			if err = d.Set("rtmon", vv); err != nil {
				return fmt.Errorf("Error reading rtmon: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading rtmon: %v", err)
		}
	}

	if err = d.Set("scfw", flattenSystemLocallogFortianalyzer3FilterScfw(o["scfw"], d, "scfw")); err != nil {
		if vv, ok := fortiAPIPatch(o["scfw"], "SystemLocallogFortianalyzer3Filter-Scfw"); ok {
			if err = d.Set("scfw", vv); err != nil {
				return fmt.Errorf("Error reading scfw: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scfw: %v", err)
		}
	}

	if err = d.Set("scply", flattenSystemLocallogFortianalyzer3FilterScply(o["scply"], d, "scply")); err != nil {
		if vv, ok := fortiAPIPatch(o["scply"], "SystemLocallogFortianalyzer3Filter-Scply"); ok {
			if err = d.Set("scply", vv); err != nil {
				return fmt.Errorf("Error reading scply: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scply: %v", err)
		}
	}

	if err = d.Set("scrmgr", flattenSystemLocallogFortianalyzer3FilterScrmgr(o["scrmgr"], d, "scrmgr")); err != nil {
		if vv, ok := fortiAPIPatch(o["scrmgr"], "SystemLocallogFortianalyzer3Filter-Scrmgr"); ok {
			if err = d.Set("scrmgr", vv); err != nil {
				return fmt.Errorf("Error reading scrmgr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scrmgr: %v", err)
		}
	}

	if err = d.Set("scvpn", flattenSystemLocallogFortianalyzer3FilterScvpn(o["scvpn"], d, "scvpn")); err != nil {
		if vv, ok := fortiAPIPatch(o["scvpn"], "SystemLocallogFortianalyzer3Filter-Scvpn"); ok {
			if err = d.Set("scvpn", vv); err != nil {
				return fmt.Errorf("Error reading scvpn: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scvpn: %v", err)
		}
	}

	if err = d.Set("system", flattenSystemLocallogFortianalyzer3FilterSystem(o["system"], d, "system")); err != nil {
		if vv, ok := fortiAPIPatch(o["system"], "SystemLocallogFortianalyzer3Filter-System"); ok {
			if err = d.Set("system", vv); err != nil {
				return fmt.Errorf("Error reading system: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading system: %v", err)
		}
	}

	if err = d.Set("webport", flattenSystemLocallogFortianalyzer3FilterWebport(o["webport"], d, "webport")); err != nil {
		if vv, ok := fortiAPIPatch(o["webport"], "SystemLocallogFortianalyzer3Filter-Webport"); ok {
			if err = d.Set("webport", vv); err != nil {
				return fmt.Errorf("Error reading webport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading webport: %v", err)
		}
	}

	return nil
}

func flattenSystemLocallogFortianalyzer3FilterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemLocallogFortianalyzer3FilterAid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzer3FilterDevcfg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzer3FilterDevops(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzer3FilterDiskquota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzer3FilterDm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzer3FilterDocker(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzer3FilterDvm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzer3FilterEdiscovery(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzer3FilterEpmgr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzer3FilterEvent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzer3FilterEventmgmt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzer3FilterFaz(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzer3FilterFazha(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzer3FilterFazsys(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzer3FilterFgd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzer3FilterFgfm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzer3FilterFips(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzer3FilterFmgws(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzer3FilterFmlmgr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzer3FilterFmwmgr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzer3FilterFortiview(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzer3FilterGlbcfg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzer3FilterHa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzer3FilterHcache(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzer3FilterIncident(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzer3FilterIolog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzer3FilterLogd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzer3FilterLogdb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzer3FilterLogdev(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzer3FilterLogfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzer3FilterLogging(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzer3FilterLrmgr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzer3FilterObjcfg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzer3FilterReport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzer3FilterRev(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzer3FilterRtmon(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzer3FilterScfw(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzer3FilterScply(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzer3FilterScrmgr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzer3FilterScvpn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzer3FilterSystem(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLocallogFortianalyzer3FilterWebport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemLocallogFortianalyzer3Filter(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("aid"); ok {
		t, err := expandSystemLocallogFortianalyzer3FilterAid(d, v, "aid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aid"] = t
		}
	}

	if v, ok := d.GetOk("devcfg"); ok {
		t, err := expandSystemLocallogFortianalyzer3FilterDevcfg(d, v, "devcfg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["devcfg"] = t
		}
	}

	if v, ok := d.GetOk("devops"); ok {
		t, err := expandSystemLocallogFortianalyzer3FilterDevops(d, v, "devops")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["devops"] = t
		}
	}

	if v, ok := d.GetOk("diskquota"); ok {
		t, err := expandSystemLocallogFortianalyzer3FilterDiskquota(d, v, "diskquota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diskquota"] = t
		}
	}

	if v, ok := d.GetOk("dm"); ok {
		t, err := expandSystemLocallogFortianalyzer3FilterDm(d, v, "dm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dm"] = t
		}
	}

	if v, ok := d.GetOk("docker"); ok {
		t, err := expandSystemLocallogFortianalyzer3FilterDocker(d, v, "docker")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["docker"] = t
		}
	}

	if v, ok := d.GetOk("dvm"); ok {
		t, err := expandSystemLocallogFortianalyzer3FilterDvm(d, v, "dvm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dvm"] = t
		}
	}

	if v, ok := d.GetOk("ediscovery"); ok {
		t, err := expandSystemLocallogFortianalyzer3FilterEdiscovery(d, v, "ediscovery")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ediscovery"] = t
		}
	}

	if v, ok := d.GetOk("epmgr"); ok {
		t, err := expandSystemLocallogFortianalyzer3FilterEpmgr(d, v, "epmgr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["epmgr"] = t
		}
	}

	if v, ok := d.GetOk("event"); ok {
		t, err := expandSystemLocallogFortianalyzer3FilterEvent(d, v, "event")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["event"] = t
		}
	}

	if v, ok := d.GetOk("eventmgmt"); ok {
		t, err := expandSystemLocallogFortianalyzer3FilterEventmgmt(d, v, "eventmgmt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eventmgmt"] = t
		}
	}

	if v, ok := d.GetOk("faz"); ok {
		t, err := expandSystemLocallogFortianalyzer3FilterFaz(d, v, "faz")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["faz"] = t
		}
	}

	if v, ok := d.GetOk("fazha"); ok {
		t, err := expandSystemLocallogFortianalyzer3FilterFazha(d, v, "fazha")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fazha"] = t
		}
	}

	if v, ok := d.GetOk("fazsys"); ok {
		t, err := expandSystemLocallogFortianalyzer3FilterFazsys(d, v, "fazsys")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fazsys"] = t
		}
	}

	if v, ok := d.GetOk("fgd"); ok {
		t, err := expandSystemLocallogFortianalyzer3FilterFgd(d, v, "fgd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgd"] = t
		}
	}

	if v, ok := d.GetOk("fgfm"); ok {
		t, err := expandSystemLocallogFortianalyzer3FilterFgfm(d, v, "fgfm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fgfm"] = t
		}
	}

	if v, ok := d.GetOk("fips"); ok {
		t, err := expandSystemLocallogFortianalyzer3FilterFips(d, v, "fips")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fips"] = t
		}
	}

	if v, ok := d.GetOk("fmgws"); ok {
		t, err := expandSystemLocallogFortianalyzer3FilterFmgws(d, v, "fmgws")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fmgws"] = t
		}
	}

	if v, ok := d.GetOk("fmlmgr"); ok {
		t, err := expandSystemLocallogFortianalyzer3FilterFmlmgr(d, v, "fmlmgr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fmlmgr"] = t
		}
	}

	if v, ok := d.GetOk("fmwmgr"); ok {
		t, err := expandSystemLocallogFortianalyzer3FilterFmwmgr(d, v, "fmwmgr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fmwmgr"] = t
		}
	}

	if v, ok := d.GetOk("fortiview"); ok {
		t, err := expandSystemLocallogFortianalyzer3FilterFortiview(d, v, "fortiview")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiview"] = t
		}
	}

	if v, ok := d.GetOk("glbcfg"); ok {
		t, err := expandSystemLocallogFortianalyzer3FilterGlbcfg(d, v, "glbcfg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["glbcfg"] = t
		}
	}

	if v, ok := d.GetOk("ha"); ok {
		t, err := expandSystemLocallogFortianalyzer3FilterHa(d, v, "ha")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha"] = t
		}
	}

	if v, ok := d.GetOk("hcache"); ok {
		t, err := expandSystemLocallogFortianalyzer3FilterHcache(d, v, "hcache")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hcache"] = t
		}
	}

	if v, ok := d.GetOk("incident"); ok {
		t, err := expandSystemLocallogFortianalyzer3FilterIncident(d, v, "incident")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["incident"] = t
		}
	}

	if v, ok := d.GetOk("iolog"); ok {
		t, err := expandSystemLocallogFortianalyzer3FilterIolog(d, v, "iolog")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["iolog"] = t
		}
	}

	if v, ok := d.GetOk("logd"); ok {
		t, err := expandSystemLocallogFortianalyzer3FilterLogd(d, v, "logd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logd"] = t
		}
	}

	if v, ok := d.GetOk("logdb"); ok {
		t, err := expandSystemLocallogFortianalyzer3FilterLogdb(d, v, "logdb")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logdb"] = t
		}
	}

	if v, ok := d.GetOk("logdev"); ok {
		t, err := expandSystemLocallogFortianalyzer3FilterLogdev(d, v, "logdev")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logdev"] = t
		}
	}

	if v, ok := d.GetOk("logfile"); ok {
		t, err := expandSystemLocallogFortianalyzer3FilterLogfile(d, v, "logfile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logfile"] = t
		}
	}

	if v, ok := d.GetOk("logging"); ok {
		t, err := expandSystemLocallogFortianalyzer3FilterLogging(d, v, "logging")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logging"] = t
		}
	}

	if v, ok := d.GetOk("lrmgr"); ok {
		t, err := expandSystemLocallogFortianalyzer3FilterLrmgr(d, v, "lrmgr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lrmgr"] = t
		}
	}

	if v, ok := d.GetOk("objcfg"); ok {
		t, err := expandSystemLocallogFortianalyzer3FilterObjcfg(d, v, "objcfg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["objcfg"] = t
		}
	}

	if v, ok := d.GetOk("report"); ok {
		t, err := expandSystemLocallogFortianalyzer3FilterReport(d, v, "report")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["report"] = t
		}
	}

	if v, ok := d.GetOk("rev"); ok {
		t, err := expandSystemLocallogFortianalyzer3FilterRev(d, v, "rev")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rev"] = t
		}
	}

	if v, ok := d.GetOk("rtmon"); ok {
		t, err := expandSystemLocallogFortianalyzer3FilterRtmon(d, v, "rtmon")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rtmon"] = t
		}
	}

	if v, ok := d.GetOk("scfw"); ok {
		t, err := expandSystemLocallogFortianalyzer3FilterScfw(d, v, "scfw")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scfw"] = t
		}
	}

	if v, ok := d.GetOk("scply"); ok {
		t, err := expandSystemLocallogFortianalyzer3FilterScply(d, v, "scply")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scply"] = t
		}
	}

	if v, ok := d.GetOk("scrmgr"); ok {
		t, err := expandSystemLocallogFortianalyzer3FilterScrmgr(d, v, "scrmgr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scrmgr"] = t
		}
	}

	if v, ok := d.GetOk("scvpn"); ok {
		t, err := expandSystemLocallogFortianalyzer3FilterScvpn(d, v, "scvpn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scvpn"] = t
		}
	}

	if v, ok := d.GetOk("system"); ok {
		t, err := expandSystemLocallogFortianalyzer3FilterSystem(d, v, "system")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["system"] = t
		}
	}

	if v, ok := d.GetOk("webport"); ok {
		t, err := expandSystemLocallogFortianalyzer3FilterWebport(d, v, "webport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webport"] = t
		}
	}

	return &obj, nil
}
