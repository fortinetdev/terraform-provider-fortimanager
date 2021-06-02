// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IPv6 DoS policies.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourcePackagesFirewallDosPolicy6() *schema.Resource {
	return &schema.Resource{
		Create: resourcePackagesFirewallDosPolicy6Create,
		Read:   resourcePackagesFirewallDosPolicy6Read,
		Update: resourcePackagesFirewallDosPolicy6Update,
		Delete: resourcePackagesFirewallDosPolicy6Delete,

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
			"pkg": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"anomaly": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"quarantine": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"quarantine_expiry": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"quarantine_log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"thresholddefault": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dstaddr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"policyid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcaddr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
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

func resourcePackagesFirewallDosPolicy6Create(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	pkg := d.Get("pkg").(string)
	var paralist []string
	paralist = append(paralist, pkg)

	obj, err := getObjectPackagesFirewallDosPolicy6(d)
	if err != nil {
		return fmt.Errorf("Error creating PackagesFirewallDosPolicy6 resource while getting object: %v", err)
	}

	v, err := c.CreatePackagesFirewallDosPolicy6(obj, adomv, paralist)

	if err != nil {
		return fmt.Errorf("Error creating PackagesFirewallDosPolicy6 resource: %v", err)
	}

	if v != nil && v["policyid"] != nil {
		if vidn, ok := v["policyid"].(float64); ok {
			d.SetId(strconv.Itoa(int(vidn)))
			return resourcePackagesFirewallDosPolicy6Read(d, m)
		} else {
			return fmt.Errorf("Error creating PackagesFirewallDosPolicy6 resource: %v", err)
		}
	}

	d.SetId(strconv.Itoa(getIntKey(d, "policyid")))

	return resourcePackagesFirewallDosPolicy6Read(d, m)
}

func resourcePackagesFirewallDosPolicy6Update(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	pkg := d.Get("pkg").(string)
	var paralist []string
	paralist = append(paralist, pkg)

	obj, err := getObjectPackagesFirewallDosPolicy6(d)
	if err != nil {
		return fmt.Errorf("Error updating PackagesFirewallDosPolicy6 resource while getting object: %v", err)
	}

	_, err = c.UpdatePackagesFirewallDosPolicy6(obj, adomv, mkey, paralist)
	if err != nil {
		return fmt.Errorf("Error updating PackagesFirewallDosPolicy6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "policyid")))

	return resourcePackagesFirewallDosPolicy6Read(d, m)
}

func resourcePackagesFirewallDosPolicy6Delete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	pkg := d.Get("pkg").(string)
	var paralist []string
	paralist = append(paralist, pkg)

	err = c.DeletePackagesFirewallDosPolicy6(adomv, mkey, paralist)
	if err != nil {
		return fmt.Errorf("Error deleting PackagesFirewallDosPolicy6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourcePackagesFirewallDosPolicy6Read(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	pkg := d.Get("pkg").(string)
	if pkg == "" {
		pkg = importOptionChecking(m.(*FortiClient).Cfg, "pkg")
		if err = d.Set("pkg", pkg); err != nil {
			return fmt.Errorf("Error set params pkg: %v", err)
		}
	}
	var paralist []string
	paralist = append(paralist, pkg)

	o, err := c.ReadPackagesFirewallDosPolicy6(adomv, mkey, paralist)
	if err != nil {
		return fmt.Errorf("Error reading PackagesFirewallDosPolicy6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectPackagesFirewallDosPolicy6(d, o)
	if err != nil {
		return fmt.Errorf("Error reading PackagesFirewallDosPolicy6 resource from API: %v", err)
	}
	return nil
}

func flattenPackagesFirewallDosPolicy6Anomaly(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			v := flattenPackagesFirewallDosPolicy6AnomalyAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "PackagesFirewallDosPolicy6-Anomaly-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := i["log"]; ok {
			v := flattenPackagesFirewallDosPolicy6AnomalyLog(i["log"], d, pre_append)
			tmp["log"] = fortiAPISubPartPatch(v, "PackagesFirewallDosPolicy6-Anomaly-Log")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenPackagesFirewallDosPolicy6AnomalyName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "PackagesFirewallDosPolicy6-Anomaly-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine"
		if _, ok := i["quarantine"]; ok {
			v := flattenPackagesFirewallDosPolicy6AnomalyQuarantine(i["quarantine"], d, pre_append)
			tmp["quarantine"] = fortiAPISubPartPatch(v, "PackagesFirewallDosPolicy6-Anomaly-Quarantine")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_expiry"
		if _, ok := i["quarantine-expiry"]; ok {
			v := flattenPackagesFirewallDosPolicy6AnomalyQuarantineExpiry(i["quarantine-expiry"], d, pre_append)
			tmp["quarantine_expiry"] = fortiAPISubPartPatch(v, "PackagesFirewallDosPolicy6-Anomaly-QuarantineExpiry")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_log"
		if _, ok := i["quarantine-log"]; ok {
			v := flattenPackagesFirewallDosPolicy6AnomalyQuarantineLog(i["quarantine-log"], d, pre_append)
			tmp["quarantine_log"] = fortiAPISubPartPatch(v, "PackagesFirewallDosPolicy6-Anomaly-QuarantineLog")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenPackagesFirewallDosPolicy6AnomalyStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "PackagesFirewallDosPolicy6-Anomaly-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold"
		if _, ok := i["threshold"]; ok {
			v := flattenPackagesFirewallDosPolicy6AnomalyThreshold(i["threshold"], d, pre_append)
			tmp["threshold"] = fortiAPISubPartPatch(v, "PackagesFirewallDosPolicy6-Anomaly-Threshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "thresholddefault"
		if _, ok := i["threshold(default)"]; ok {
			v := flattenPackagesFirewallDosPolicy6AnomalyThresholdDefault(i["threshold(default)"], d, pre_append)
			tmp["thresholddefault"] = fortiAPISubPartPatch(v, "PackagesFirewallDosPolicy6-Anomaly-ThresholdDefault")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenPackagesFirewallDosPolicy6AnomalyAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "pass",
			1: "block",
			2: "proxy",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenPackagesFirewallDosPolicy6AnomalyLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenPackagesFirewallDosPolicy6AnomalyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicy6AnomalyQuarantine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "none",
			1: "attacker",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenPackagesFirewallDosPolicy6AnomalyQuarantineExpiry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicy6AnomalyQuarantineLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenPackagesFirewallDosPolicy6AnomalyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenPackagesFirewallDosPolicy6AnomalyThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicy6AnomalyThresholdDefault(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicy6Comments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicy6Dstaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicy6Interface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicy6Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicy6Policyid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicy6Service(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicy6Srcaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicy6Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func refreshObjectPackagesFirewallDosPolicy6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if isImportTable() {
		if err = d.Set("anomaly", flattenPackagesFirewallDosPolicy6Anomaly(o["anomaly"], d, "anomaly")); err != nil {
			if vv, ok := fortiAPIPatch(o["anomaly"], "PackagesFirewallDosPolicy6-Anomaly"); ok {
				if err = d.Set("anomaly", vv); err != nil {
					return fmt.Errorf("Error reading anomaly: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading anomaly: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("anomaly"); ok {
			if err = d.Set("anomaly", flattenPackagesFirewallDosPolicy6Anomaly(o["anomaly"], d, "anomaly")); err != nil {
				if vv, ok := fortiAPIPatch(o["anomaly"], "PackagesFirewallDosPolicy6-Anomaly"); ok {
					if err = d.Set("anomaly", vv); err != nil {
						return fmt.Errorf("Error reading anomaly: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading anomaly: %v", err)
				}
			}
		}
	}

	if err = d.Set("comments", flattenPackagesFirewallDosPolicy6Comments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "PackagesFirewallDosPolicy6-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("dstaddr", flattenPackagesFirewallDosPolicy6Dstaddr(o["dstaddr"], d, "dstaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr"], "PackagesFirewallDosPolicy6-Dstaddr"); ok {
			if err = d.Set("dstaddr", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr: %v", err)
		}
	}

	if err = d.Set("interface", flattenPackagesFirewallDosPolicy6Interface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "PackagesFirewallDosPolicy6-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("name", flattenPackagesFirewallDosPolicy6Name(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "PackagesFirewallDosPolicy6-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("policyid", flattenPackagesFirewallDosPolicy6Policyid(o["policyid"], d, "policyid")); err != nil {
		if vv, ok := fortiAPIPatch(o["policyid"], "PackagesFirewallDosPolicy6-Policyid"); ok {
			if err = d.Set("policyid", vv); err != nil {
				return fmt.Errorf("Error reading policyid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policyid: %v", err)
		}
	}

	if err = d.Set("service", flattenPackagesFirewallDosPolicy6Service(o["service"], d, "service")); err != nil {
		if vv, ok := fortiAPIPatch(o["service"], "PackagesFirewallDosPolicy6-Service"); ok {
			if err = d.Set("service", vv); err != nil {
				return fmt.Errorf("Error reading service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service: %v", err)
		}
	}

	if err = d.Set("srcaddr", flattenPackagesFirewallDosPolicy6Srcaddr(o["srcaddr"], d, "srcaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr"], "PackagesFirewallDosPolicy6-Srcaddr"); ok {
			if err = d.Set("srcaddr", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr: %v", err)
		}
	}

	if err = d.Set("status", flattenPackagesFirewallDosPolicy6Status(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "PackagesFirewallDosPolicy6-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenPackagesFirewallDosPolicy6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandPackagesFirewallDosPolicy6Anomaly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandPackagesFirewallDosPolicy6AnomalyAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["log"], _ = expandPackagesFirewallDosPolicy6AnomalyLog(d, i["log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandPackagesFirewallDosPolicy6AnomalyName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["quarantine"], _ = expandPackagesFirewallDosPolicy6AnomalyQuarantine(d, i["quarantine"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_expiry"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["quarantine-expiry"], _ = expandPackagesFirewallDosPolicy6AnomalyQuarantineExpiry(d, i["quarantine_expiry"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_log"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["quarantine-log"], _ = expandPackagesFirewallDosPolicy6AnomalyQuarantineLog(d, i["quarantine_log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandPackagesFirewallDosPolicy6AnomalyStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["threshold"], _ = expandPackagesFirewallDosPolicy6AnomalyThreshold(d, i["threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "thresholddefault"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["threshold(default)"], _ = expandPackagesFirewallDosPolicy6AnomalyThresholdDefault(d, i["thresholddefault"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandPackagesFirewallDosPolicy6AnomalyAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicy6AnomalyLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicy6AnomalyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicy6AnomalyQuarantine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicy6AnomalyQuarantineExpiry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicy6AnomalyQuarantineLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicy6AnomalyStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicy6AnomalyThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicy6AnomalyThresholdDefault(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicy6Comments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicy6Dstaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicy6Interface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicy6Name(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicy6Policyid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicy6Service(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicy6Srcaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicy6Status(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectPackagesFirewallDosPolicy6(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("anomaly"); ok {
		t, err := expandPackagesFirewallDosPolicy6Anomaly(d, v, "anomaly")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["anomaly"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {
		t, err := expandPackagesFirewallDosPolicy6Comments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr"); ok {
		t, err := expandPackagesFirewallDosPolicy6Dstaddr(d, v, "dstaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandPackagesFirewallDosPolicy6Interface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandPackagesFirewallDosPolicy6Name(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("policyid"); ok {
		t, err := expandPackagesFirewallDosPolicy6Policyid(d, v, "policyid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policyid"] = t
		}
	}

	if v, ok := d.GetOk("service"); ok {
		t, err := expandPackagesFirewallDosPolicy6Service(d, v, "service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr"); ok {
		t, err := expandPackagesFirewallDosPolicy6Srcaddr(d, v, "srcaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandPackagesFirewallDosPolicy6Status(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
