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
						"synproxy_tcp_mss": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"synproxy_tcp_sack": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"synproxy_tcp_timestamp": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"synproxy_tcp_window": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"synproxy_tcp_windowscale": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"synproxy_tos": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"synproxy_ttl": &schema.Schema{
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
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"srcaddr": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uuid": &schema.Schema{
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

	v, err := c.CreatePackagesFirewallDosPolicy(obj, adomv, paralist)

	if err != nil {
		return fmt.Errorf("Error creating PackagesFirewallDosPolicy resource: %v", err)
	}

	if v != nil && v["policyid"] != nil {
		if vidn, ok := v["policyid"].(float64); ok {
			d.SetId(strconv.Itoa(int(vidn)))
			return resourcePackagesFirewallDosPolicyRead(d, m)
		} else {
			return fmt.Errorf("Error creating PackagesFirewallDosPolicy resource: %v", err)
		}
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_tcp_mss"
		if _, ok := i["synproxy-tcp-mss"]; ok {
			v := flattenPackagesFirewallDosPolicyAnomalySynproxyTcpMss(i["synproxy-tcp-mss"], d, pre_append)
			tmp["synproxy_tcp_mss"] = fortiAPISubPartPatch(v, "PackagesFirewallDosPolicy-Anomaly-SynproxyTcpMss")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_tcp_sack"
		if _, ok := i["synproxy-tcp-sack"]; ok {
			v := flattenPackagesFirewallDosPolicyAnomalySynproxyTcpSack(i["synproxy-tcp-sack"], d, pre_append)
			tmp["synproxy_tcp_sack"] = fortiAPISubPartPatch(v, "PackagesFirewallDosPolicy-Anomaly-SynproxyTcpSack")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_tcp_timestamp"
		if _, ok := i["synproxy-tcp-timestamp"]; ok {
			v := flattenPackagesFirewallDosPolicyAnomalySynproxyTcpTimestamp(i["synproxy-tcp-timestamp"], d, pre_append)
			tmp["synproxy_tcp_timestamp"] = fortiAPISubPartPatch(v, "PackagesFirewallDosPolicy-Anomaly-SynproxyTcpTimestamp")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_tcp_window"
		if _, ok := i["synproxy-tcp-window"]; ok {
			v := flattenPackagesFirewallDosPolicyAnomalySynproxyTcpWindow(i["synproxy-tcp-window"], d, pre_append)
			tmp["synproxy_tcp_window"] = fortiAPISubPartPatch(v, "PackagesFirewallDosPolicy-Anomaly-SynproxyTcpWindow")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_tcp_windowscale"
		if _, ok := i["synproxy-tcp-windowscale"]; ok {
			v := flattenPackagesFirewallDosPolicyAnomalySynproxyTcpWindowscale(i["synproxy-tcp-windowscale"], d, pre_append)
			tmp["synproxy_tcp_windowscale"] = fortiAPISubPartPatch(v, "PackagesFirewallDosPolicy-Anomaly-SynproxyTcpWindowscale")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_tos"
		if _, ok := i["synproxy-tos"]; ok {
			v := flattenPackagesFirewallDosPolicyAnomalySynproxyTos(i["synproxy-tos"], d, pre_append)
			tmp["synproxy_tos"] = fortiAPISubPartPatch(v, "PackagesFirewallDosPolicy-Anomaly-SynproxyTos")
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_ttl"
		if _, ok := i["synproxy-ttl"]; ok {
			v := flattenPackagesFirewallDosPolicyAnomalySynproxyTtl(i["synproxy-ttl"], d, pre_append)
			tmp["synproxy_ttl"] = fortiAPISubPartPatch(v, "PackagesFirewallDosPolicy-Anomaly-SynproxyTtl")
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
	return v
}

func flattenPackagesFirewallDosPolicyAnomalyLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicyAnomalyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicyAnomalyQuarantine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicyAnomalyQuarantineExpiry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicyAnomalyQuarantineLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicyAnomalyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicyAnomalySynproxyTcpMss(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicyAnomalySynproxyTcpSack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicyAnomalySynproxyTcpTimestamp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicyAnomalySynproxyTcpWindow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicyAnomalySynproxyTcpWindowscale(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicyAnomalySynproxyTos(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicyAnomalySynproxyTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
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
	return flattenStringList(v)
}

func flattenPackagesFirewallDosPolicyInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesFirewallDosPolicyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicyPolicyid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicyService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesFirewallDosPolicySrcaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenPackagesFirewallDosPolicyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicyUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectPackagesFirewallDosPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if dssValue := d.Get("dynamic_sort_subtable"); dssValue == "" {
		d.Set("dynamic_sort_subtable", "false")
	}

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

	if err = d.Set("uuid", flattenPackagesFirewallDosPolicyUuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "PackagesFirewallDosPolicy-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
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
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["action"], _ = expandPackagesFirewallDosPolicyAnomalyAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["log"], _ = expandPackagesFirewallDosPolicyAnomalyLog(d, i["log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["name"], _ = expandPackagesFirewallDosPolicyAnomalyName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["quarantine"], _ = expandPackagesFirewallDosPolicyAnomalyQuarantine(d, i["quarantine"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_expiry"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["quarantine-expiry"], _ = expandPackagesFirewallDosPolicyAnomalyQuarantineExpiry(d, i["quarantine_expiry"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_log"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["quarantine-log"], _ = expandPackagesFirewallDosPolicyAnomalyQuarantineLog(d, i["quarantine_log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["status"], _ = expandPackagesFirewallDosPolicyAnomalyStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_tcp_mss"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["synproxy-tcp-mss"], _ = expandPackagesFirewallDosPolicyAnomalySynproxyTcpMss(d, i["synproxy_tcp_mss"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_tcp_sack"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["synproxy-tcp-sack"], _ = expandPackagesFirewallDosPolicyAnomalySynproxyTcpSack(d, i["synproxy_tcp_sack"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_tcp_timestamp"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["synproxy-tcp-timestamp"], _ = expandPackagesFirewallDosPolicyAnomalySynproxyTcpTimestamp(d, i["synproxy_tcp_timestamp"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_tcp_window"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["synproxy-tcp-window"], _ = expandPackagesFirewallDosPolicyAnomalySynproxyTcpWindow(d, i["synproxy_tcp_window"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_tcp_windowscale"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["synproxy-tcp-windowscale"], _ = expandPackagesFirewallDosPolicyAnomalySynproxyTcpWindowscale(d, i["synproxy_tcp_windowscale"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_tos"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["synproxy-tos"], _ = expandPackagesFirewallDosPolicyAnomalySynproxyTos(d, i["synproxy_tos"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "synproxy_ttl"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["synproxy-ttl"], _ = expandPackagesFirewallDosPolicyAnomalySynproxyTtl(d, i["synproxy_ttl"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["threshold"], _ = expandPackagesFirewallDosPolicyAnomalyThreshold(d, i["threshold"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "thresholddefault"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
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

func expandPackagesFirewallDosPolicyAnomalySynproxyTcpMss(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicyAnomalySynproxyTcpSack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicyAnomalySynproxyTcpTimestamp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicyAnomalySynproxyTcpWindow(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicyAnomalySynproxyTcpWindowscale(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicyAnomalySynproxyTos(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicyAnomalySynproxyTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
	return expandStringList(v.([]interface{})), nil
}

func expandPackagesFirewallDosPolicyInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandPackagesFirewallDosPolicyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicyPolicyid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicyService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandPackagesFirewallDosPolicySrcaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.([]interface{})), nil
}

func expandPackagesFirewallDosPolicyStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicyUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectPackagesFirewallDosPolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("anomaly"); ok || d.HasChange("anomaly") {
		t, err := expandPackagesFirewallDosPolicyAnomaly(d, v, "anomaly")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["anomaly"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandPackagesFirewallDosPolicyComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr"); ok || d.HasChange("dstaddr") {
		t, err := expandPackagesFirewallDosPolicyDstaddr(d, v, "dstaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandPackagesFirewallDosPolicyInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandPackagesFirewallDosPolicyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("policyid"); ok || d.HasChange("policyid") {
		t, err := expandPackagesFirewallDosPolicyPolicyid(d, v, "policyid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policyid"] = t
		}
	}

	if v, ok := d.GetOk("service"); ok || d.HasChange("service") {
		t, err := expandPackagesFirewallDosPolicyService(d, v, "service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr"); ok || d.HasChange("srcaddr") {
		t, err := expandPackagesFirewallDosPolicySrcaddr(d, v, "srcaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandPackagesFirewallDosPolicyStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandPackagesFirewallDosPolicyUuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	return &obj, nil
}
