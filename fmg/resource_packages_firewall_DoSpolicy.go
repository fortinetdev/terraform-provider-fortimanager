// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IPv4 DoS policies.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourcePackagesFirewallDosPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourcePackagesFirewallDosPolicyCreate,
		Read:   resourcePackagesFirewallDosPolicyRead,
		Update: resourcePackagesFirewallDosPolicyUpdate,
		Delete: resourcePackagesFirewallDosPolicyDelete,

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

func resourcePackagesFirewallDosPolicyCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectPackagesFirewallDosPolicy(d)
	if err != nil {
		return fmt.Errorf("Error creating PackagesFirewallDosPolicy resource while getting object: %v", err)
	}

	_, err = c.CreatePackagesFirewallDosPolicy(obj, adomv, paralist)

	if err != nil {
		return fmt.Errorf("Error creating PackagesFirewallDosPolicy resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "policyid")))

	return resourcePackagesFirewallDosPolicyRead(d, m)
}

func resourcePackagesFirewallDosPolicyUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectPackagesFirewallDosPolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating PackagesFirewallDosPolicy resource while getting object: %v", err)
	}

	_, err = c.UpdatePackagesFirewallDosPolicy(obj, adomv, mkey, paralist)
	if err != nil {
		return fmt.Errorf("Error updating PackagesFirewallDosPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "policyid")))

	return resourcePackagesFirewallDosPolicyRead(d, m)
}

func resourcePackagesFirewallDosPolicyDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeletePackagesFirewallDosPolicy(adomv, mkey, paralist)
	if err != nil {
		return fmt.Errorf("Error deleting PackagesFirewallDosPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourcePackagesFirewallDosPolicyRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadPackagesFirewallDosPolicy(adomv, mkey, paralist)
	if err != nil {
		return fmt.Errorf("Error reading PackagesFirewallDosPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectPackagesFirewallDosPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading PackagesFirewallDosPolicy resource from API: %v", err)
	}
	return nil
}

func flattenPackagesFirewallDosPolicyAnomaly(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			v := flattenPackagesFirewallDosPolicyAnomalyAction(i["action"], d, pre_append)
			tmp["action"] = fortiAPISubPartPatch(v, "PackagesFirewallDosPolicy-Anomaly-Action")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := i["log"]; ok {
			v := flattenPackagesFirewallDosPolicyAnomalyLog(i["log"], d, pre_append)
			tmp["log"] = fortiAPISubPartPatch(v, "PackagesFirewallDosPolicy-Anomaly-Log")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			v := flattenPackagesFirewallDosPolicyAnomalyName(i["name"], d, pre_append)
			tmp["name"] = fortiAPISubPartPatch(v, "PackagesFirewallDosPolicy-Anomaly-Name")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine"
		if _, ok := i["quarantine"]; ok {
			v := flattenPackagesFirewallDosPolicyAnomalyQuarantine(i["quarantine"], d, pre_append)
			tmp["quarantine"] = fortiAPISubPartPatch(v, "PackagesFirewallDosPolicy-Anomaly-Quarantine")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_expiry"
		if _, ok := i["quarantine-expiry"]; ok {
			v := flattenPackagesFirewallDosPolicyAnomalyQuarantineExpiry(i["quarantine-expiry"], d, pre_append)
			tmp["quarantine_expiry"] = fortiAPISubPartPatch(v, "PackagesFirewallDosPolicy-Anomaly-QuarantineExpiry")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_log"
		if _, ok := i["quarantine-log"]; ok {
			v := flattenPackagesFirewallDosPolicyAnomalyQuarantineLog(i["quarantine-log"], d, pre_append)
			tmp["quarantine_log"] = fortiAPISubPartPatch(v, "PackagesFirewallDosPolicy-Anomaly-QuarantineLog")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			v := flattenPackagesFirewallDosPolicyAnomalyStatus(i["status"], d, pre_append)
			tmp["status"] = fortiAPISubPartPatch(v, "PackagesFirewallDosPolicy-Anomaly-Status")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold"
		if _, ok := i["threshold"]; ok {
			v := flattenPackagesFirewallDosPolicyAnomalyThreshold(i["threshold"], d, pre_append)
			tmp["threshold"] = fortiAPISubPartPatch(v, "PackagesFirewallDosPolicy-Anomaly-Threshold")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "thresholddefault"
		if _, ok := i["threshold(default)"]; ok {
			v := flattenPackagesFirewallDosPolicyAnomalyThresholdDefault(i["threshold(default)"], d, pre_append)
			tmp["thresholddefault"] = fortiAPISubPartPatch(v, "PackagesFirewallDosPolicy-Anomaly-ThresholdDefault")
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenPackagesFirewallDosPolicyAnomalyAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenPackagesFirewallDosPolicyAnomalyLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenPackagesFirewallDosPolicyAnomalyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicyAnomalyQuarantine(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenPackagesFirewallDosPolicyAnomalyQuarantineExpiry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicyAnomalyQuarantineLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenPackagesFirewallDosPolicyAnomalyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenPackagesFirewallDosPolicyAnomalyThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicyAnomalyThresholdDefault(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicyComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicyDstaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicyInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicyPolicyid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicyService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicySrcaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func refreshObjectPackagesFirewallDosPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if isImportTable() {
		if err = d.Set("anomaly", flattenPackagesFirewallDosPolicyAnomaly(o["anomaly"], d, "anomaly")); err != nil {
			if vv, ok := fortiAPIPatch(o["anomaly"], "PackagesFirewallDosPolicy-Anomaly"); ok {
				if err = d.Set("anomaly", vv); err != nil {
					return fmt.Errorf("Error reading anomaly: %v", err)
				}
			} else {
				return fmt.Errorf("Error reading anomaly: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("anomaly"); ok {
			if err = d.Set("anomaly", flattenPackagesFirewallDosPolicyAnomaly(o["anomaly"], d, "anomaly")); err != nil {
				if vv, ok := fortiAPIPatch(o["anomaly"], "PackagesFirewallDosPolicy-Anomaly"); ok {
					if err = d.Set("anomaly", vv); err != nil {
						return fmt.Errorf("Error reading anomaly: %v", err)
					}
				} else {
					return fmt.Errorf("Error reading anomaly: %v", err)
				}
			}
		}
	}

	if err = d.Set("comments", flattenPackagesFirewallDosPolicyComments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "PackagesFirewallDosPolicy-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("dstaddr", flattenPackagesFirewallDosPolicyDstaddr(o["dstaddr"], d, "dstaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr"], "PackagesFirewallDosPolicy-Dstaddr"); ok {
			if err = d.Set("dstaddr", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr: %v", err)
		}
	}

	if err = d.Set("interface", flattenPackagesFirewallDosPolicyInterface(o["interface"], d, "interface")); err != nil {
		if vv, ok := fortiAPIPatch(o["interface"], "PackagesFirewallDosPolicy-Interface"); ok {
			if err = d.Set("interface", vv); err != nil {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("name", flattenPackagesFirewallDosPolicyName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "PackagesFirewallDosPolicy-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("policyid", flattenPackagesFirewallDosPolicyPolicyid(o["policyid"], d, "policyid")); err != nil {
		if vv, ok := fortiAPIPatch(o["policyid"], "PackagesFirewallDosPolicy-Policyid"); ok {
			if err = d.Set("policyid", vv); err != nil {
				return fmt.Errorf("Error reading policyid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policyid: %v", err)
		}
	}

	if err = d.Set("service", flattenPackagesFirewallDosPolicyService(o["service"], d, "service")); err != nil {
		if vv, ok := fortiAPIPatch(o["service"], "PackagesFirewallDosPolicy-Service"); ok {
			if err = d.Set("service", vv); err != nil {
				return fmt.Errorf("Error reading service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service: %v", err)
		}
	}

	if err = d.Set("srcaddr", flattenPackagesFirewallDosPolicySrcaddr(o["srcaddr"], d, "srcaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr"], "PackagesFirewallDosPolicy-Srcaddr"); ok {
			if err = d.Set("srcaddr", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr: %v", err)
		}
	}

	if err = d.Set("status", flattenPackagesFirewallDosPolicyStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "PackagesFirewallDosPolicy-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenPackagesFirewallDosPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandPackagesFirewallDosPolicyAnomaly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["action"], _ = expandPackagesFirewallDosPolicyAnomalyAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["log"], _ = expandPackagesFirewallDosPolicyAnomalyLog(d, i["log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandPackagesFirewallDosPolicyAnomalyName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["quarantine"], _ = expandPackagesFirewallDosPolicyAnomalyQuarantine(d, i["quarantine"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_expiry"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["quarantine-expiry"], _ = expandPackagesFirewallDosPolicyAnomalyQuarantineExpiry(d, i["quarantine_expiry"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_log"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["quarantine-log"], _ = expandPackagesFirewallDosPolicyAnomalyQuarantineLog(d, i["quarantine_log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandPackagesFirewallDosPolicyAnomalyStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["threshold"], _ = expandPackagesFirewallDosPolicyAnomalyThreshold(d, i["threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "thresholddefault"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["threshold(default)"], _ = expandPackagesFirewallDosPolicyAnomalyThresholdDefault(d, i["thresholddefault"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandPackagesFirewallDosPolicyAnomalyAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicyAnomalyLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicyAnomalyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicyAnomalyQuarantine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicyAnomalyQuarantineExpiry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicyAnomalyQuarantineLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicyAnomalyStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicyAnomalyThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicyAnomalyThresholdDefault(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicyComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicyDstaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicyInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicyPolicyid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicyService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicySrcaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicyStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectPackagesFirewallDosPolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("anomaly"); ok {
		t, err := expandPackagesFirewallDosPolicyAnomaly(d, v, "anomaly")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["anomaly"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {
		t, err := expandPackagesFirewallDosPolicyComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr"); ok {
		t, err := expandPackagesFirewallDosPolicyDstaddr(d, v, "dstaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandPackagesFirewallDosPolicyInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandPackagesFirewallDosPolicyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("policyid"); ok {
		t, err := expandPackagesFirewallDosPolicyPolicyid(d, v, "policyid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policyid"] = t
		}
	}

	if v, ok := d.GetOk("service"); ok {
		t, err := expandPackagesFirewallDosPolicyService(d, v, "service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr"); ok {
		t, err := expandPackagesFirewallDosPolicySrcaddr(d, v, "srcaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandPackagesFirewallDosPolicyStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
