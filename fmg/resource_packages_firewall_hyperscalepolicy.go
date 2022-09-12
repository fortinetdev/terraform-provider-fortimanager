// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IPv4 hyperscale-policies.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourcePackagesFirewallHyperscalePolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourcePackagesFirewallHyperscalePolicyCreate,
		Read:   resourcePackagesFirewallHyperscalePolicyRead,
		Update: resourcePackagesFirewallHyperscalePolicyUpdate,
		Delete: resourcePackagesFirewallHyperscalePolicyDelete,

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
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auto_asic_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cgn_eif": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cgn_eim": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cgn_log_server_grp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cgn_resource_quota": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"cgn_session_quota": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"delay_tcp_npu_session": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dstaddr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dstaddr_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dstaddr6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dstintf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"firewall_session_dirty": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"global_label": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ippool": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"label": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nat": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"policy_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"policyid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"poolname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"poolname6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"send_deny_packet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"srcaddr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"srcaddr_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"srcaddr6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"srcintf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tcp_timeout_pid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"traffic_shaper": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"traffic_shaper_reverse": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"udp_timeout_pid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourcePackagesFirewallHyperscalePolicyCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectPackagesFirewallHyperscalePolicy(d)
	if err != nil {
		return fmt.Errorf("Error creating PackagesFirewallHyperscalePolicy resource while getting object: %v", err)
	}

	_, err = c.CreatePackagesFirewallHyperscalePolicy(obj, adomv, paralist)

	if err != nil {
		return fmt.Errorf("Error creating PackagesFirewallHyperscalePolicy resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "policyid")))

	return resourcePackagesFirewallHyperscalePolicyRead(d, m)
}

func resourcePackagesFirewallHyperscalePolicyUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectPackagesFirewallHyperscalePolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating PackagesFirewallHyperscalePolicy resource while getting object: %v", err)
	}

	_, err = c.UpdatePackagesFirewallHyperscalePolicy(obj, adomv, mkey, paralist)
	if err != nil {
		return fmt.Errorf("Error updating PackagesFirewallHyperscalePolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "policyid")))

	return resourcePackagesFirewallHyperscalePolicyRead(d, m)
}

func resourcePackagesFirewallHyperscalePolicyDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeletePackagesFirewallHyperscalePolicy(adomv, mkey, paralist)
	if err != nil {
		return fmt.Errorf("Error deleting PackagesFirewallHyperscalePolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourcePackagesFirewallHyperscalePolicyRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadPackagesFirewallHyperscalePolicy(adomv, mkey, paralist)
	if err != nil {
		return fmt.Errorf("Error reading PackagesFirewallHyperscalePolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectPackagesFirewallHyperscalePolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading PackagesFirewallHyperscalePolicy resource from API: %v", err)
	}
	return nil
}

func flattenPackagesFirewallHyperscalePolicyAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallHyperscalePolicyAutoAsicOffload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallHyperscalePolicyCgnEif(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallHyperscalePolicyCgnEim(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallHyperscalePolicyCgnLogServerGrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallHyperscalePolicyCgnResourceQuota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallHyperscalePolicyCgnSessionQuota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallHyperscalePolicyComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallHyperscalePolicyDelayTcpNpuSession(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallHyperscalePolicyDstaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallHyperscalePolicyDstaddrNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallHyperscalePolicyDstaddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallHyperscalePolicyDstintf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallHyperscalePolicyFirewallSessionDirty(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallHyperscalePolicyGlobalLabel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallHyperscalePolicyIppool(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallHyperscalePolicyLabel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallHyperscalePolicyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallHyperscalePolicyNat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallHyperscalePolicyPolicyOffload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallHyperscalePolicyPolicyid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallHyperscalePolicyPoolname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallHyperscalePolicyPoolname6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallHyperscalePolicySendDenyPacket(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallHyperscalePolicyService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallHyperscalePolicyServiceNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallHyperscalePolicySrcaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallHyperscalePolicySrcaddrNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallHyperscalePolicySrcaddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallHyperscalePolicySrcintf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallHyperscalePolicyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallHyperscalePolicyTcpTimeoutPid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallHyperscalePolicyTrafficShaper(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallHyperscalePolicyTrafficShaperReverse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallHyperscalePolicyUdpTimeoutPid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallHyperscalePolicyUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectPackagesFirewallHyperscalePolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("action", flattenPackagesFirewallHyperscalePolicyAction(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "PackagesFirewallHyperscalePolicy-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("auto_asic_offload", flattenPackagesFirewallHyperscalePolicyAutoAsicOffload(o["auto-asic-offload"], d, "auto_asic_offload")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-asic-offload"], "PackagesFirewallHyperscalePolicy-AutoAsicOffload"); ok {
			if err = d.Set("auto_asic_offload", vv); err != nil {
				return fmt.Errorf("Error reading auto_asic_offload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_asic_offload: %v", err)
		}
	}

	if err = d.Set("cgn_eif", flattenPackagesFirewallHyperscalePolicyCgnEif(o["cgn-eif"], d, "cgn_eif")); err != nil {
		if vv, ok := fortiAPIPatch(o["cgn-eif"], "PackagesFirewallHyperscalePolicy-CgnEif"); ok {
			if err = d.Set("cgn_eif", vv); err != nil {
				return fmt.Errorf("Error reading cgn_eif: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cgn_eif: %v", err)
		}
	}

	if err = d.Set("cgn_eim", flattenPackagesFirewallHyperscalePolicyCgnEim(o["cgn-eim"], d, "cgn_eim")); err != nil {
		if vv, ok := fortiAPIPatch(o["cgn-eim"], "PackagesFirewallHyperscalePolicy-CgnEim"); ok {
			if err = d.Set("cgn_eim", vv); err != nil {
				return fmt.Errorf("Error reading cgn_eim: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cgn_eim: %v", err)
		}
	}

	if err = d.Set("cgn_log_server_grp", flattenPackagesFirewallHyperscalePolicyCgnLogServerGrp(o["cgn-log-server-grp"], d, "cgn_log_server_grp")); err != nil {
		if vv, ok := fortiAPIPatch(o["cgn-log-server-grp"], "PackagesFirewallHyperscalePolicy-CgnLogServerGrp"); ok {
			if err = d.Set("cgn_log_server_grp", vv); err != nil {
				return fmt.Errorf("Error reading cgn_log_server_grp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cgn_log_server_grp: %v", err)
		}
	}

	if err = d.Set("cgn_resource_quota", flattenPackagesFirewallHyperscalePolicyCgnResourceQuota(o["cgn-resource-quota"], d, "cgn_resource_quota")); err != nil {
		if vv, ok := fortiAPIPatch(o["cgn-resource-quota"], "PackagesFirewallHyperscalePolicy-CgnResourceQuota"); ok {
			if err = d.Set("cgn_resource_quota", vv); err != nil {
				return fmt.Errorf("Error reading cgn_resource_quota: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cgn_resource_quota: %v", err)
		}
	}

	if err = d.Set("cgn_session_quota", flattenPackagesFirewallHyperscalePolicyCgnSessionQuota(o["cgn-session-quota"], d, "cgn_session_quota")); err != nil {
		if vv, ok := fortiAPIPatch(o["cgn-session-quota"], "PackagesFirewallHyperscalePolicy-CgnSessionQuota"); ok {
			if err = d.Set("cgn_session_quota", vv); err != nil {
				return fmt.Errorf("Error reading cgn_session_quota: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading cgn_session_quota: %v", err)
		}
	}

	if err = d.Set("comments", flattenPackagesFirewallHyperscalePolicyComments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "PackagesFirewallHyperscalePolicy-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("delay_tcp_npu_session", flattenPackagesFirewallHyperscalePolicyDelayTcpNpuSession(o["delay-tcp-npu-session"], d, "delay_tcp_npu_session")); err != nil {
		if vv, ok := fortiAPIPatch(o["delay-tcp-npu-session"], "PackagesFirewallHyperscalePolicy-DelayTcpNpuSession"); ok {
			if err = d.Set("delay_tcp_npu_session", vv); err != nil {
				return fmt.Errorf("Error reading delay_tcp_npu_session: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading delay_tcp_npu_session: %v", err)
		}
	}

	if err = d.Set("dstaddr", flattenPackagesFirewallHyperscalePolicyDstaddr(o["dstaddr"], d, "dstaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr"], "PackagesFirewallHyperscalePolicy-Dstaddr"); ok {
			if err = d.Set("dstaddr", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr: %v", err)
		}
	}

	if err = d.Set("dstaddr_negate", flattenPackagesFirewallHyperscalePolicyDstaddrNegate(o["dstaddr-negate"], d, "dstaddr_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr-negate"], "PackagesFirewallHyperscalePolicy-DstaddrNegate"); ok {
			if err = d.Set("dstaddr_negate", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr_negate: %v", err)
		}
	}

	if err = d.Set("dstaddr6", flattenPackagesFirewallHyperscalePolicyDstaddr6(o["dstaddr6"], d, "dstaddr6")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr6"], "PackagesFirewallHyperscalePolicy-Dstaddr6"); ok {
			if err = d.Set("dstaddr6", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr6: %v", err)
		}
	}

	if err = d.Set("dstintf", flattenPackagesFirewallHyperscalePolicyDstintf(o["dstintf"], d, "dstintf")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstintf"], "PackagesFirewallHyperscalePolicy-Dstintf"); ok {
			if err = d.Set("dstintf", vv); err != nil {
				return fmt.Errorf("Error reading dstintf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstintf: %v", err)
		}
	}

	if err = d.Set("firewall_session_dirty", flattenPackagesFirewallHyperscalePolicyFirewallSessionDirty(o["firewall-session-dirty"], d, "firewall_session_dirty")); err != nil {
		if vv, ok := fortiAPIPatch(o["firewall-session-dirty"], "PackagesFirewallHyperscalePolicy-FirewallSessionDirty"); ok {
			if err = d.Set("firewall_session_dirty", vv); err != nil {
				return fmt.Errorf("Error reading firewall_session_dirty: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading firewall_session_dirty: %v", err)
		}
	}

	if err = d.Set("global_label", flattenPackagesFirewallHyperscalePolicyGlobalLabel(o["global-label"], d, "global_label")); err != nil {
		if vv, ok := fortiAPIPatch(o["global-label"], "PackagesFirewallHyperscalePolicy-GlobalLabel"); ok {
			if err = d.Set("global_label", vv); err != nil {
				return fmt.Errorf("Error reading global_label: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading global_label: %v", err)
		}
	}

	if err = d.Set("ippool", flattenPackagesFirewallHyperscalePolicyIppool(o["ippool"], d, "ippool")); err != nil {
		if vv, ok := fortiAPIPatch(o["ippool"], "PackagesFirewallHyperscalePolicy-Ippool"); ok {
			if err = d.Set("ippool", vv); err != nil {
				return fmt.Errorf("Error reading ippool: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ippool: %v", err)
		}
	}

	if err = d.Set("label", flattenPackagesFirewallHyperscalePolicyLabel(o["label"], d, "label")); err != nil {
		if vv, ok := fortiAPIPatch(o["label"], "PackagesFirewallHyperscalePolicy-Label"); ok {
			if err = d.Set("label", vv); err != nil {
				return fmt.Errorf("Error reading label: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading label: %v", err)
		}
	}

	if err = d.Set("name", flattenPackagesFirewallHyperscalePolicyName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "PackagesFirewallHyperscalePolicy-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("nat", flattenPackagesFirewallHyperscalePolicyNat(o["nat"], d, "nat")); err != nil {
		if vv, ok := fortiAPIPatch(o["nat"], "PackagesFirewallHyperscalePolicy-Nat"); ok {
			if err = d.Set("nat", vv); err != nil {
				return fmt.Errorf("Error reading nat: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading nat: %v", err)
		}
	}

	if err = d.Set("policy_offload", flattenPackagesFirewallHyperscalePolicyPolicyOffload(o["policy-offload"], d, "policy_offload")); err != nil {
		if vv, ok := fortiAPIPatch(o["policy-offload"], "PackagesFirewallHyperscalePolicy-PolicyOffload"); ok {
			if err = d.Set("policy_offload", vv); err != nil {
				return fmt.Errorf("Error reading policy_offload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policy_offload: %v", err)
		}
	}

	if err = d.Set("policyid", flattenPackagesFirewallHyperscalePolicyPolicyid(o["policyid"], d, "policyid")); err != nil {
		if vv, ok := fortiAPIPatch(o["policyid"], "PackagesFirewallHyperscalePolicy-Policyid"); ok {
			if err = d.Set("policyid", vv); err != nil {
				return fmt.Errorf("Error reading policyid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policyid: %v", err)
		}
	}

	if err = d.Set("poolname", flattenPackagesFirewallHyperscalePolicyPoolname(o["poolname"], d, "poolname")); err != nil {
		if vv, ok := fortiAPIPatch(o["poolname"], "PackagesFirewallHyperscalePolicy-Poolname"); ok {
			if err = d.Set("poolname", vv); err != nil {
				return fmt.Errorf("Error reading poolname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading poolname: %v", err)
		}
	}

	if err = d.Set("poolname6", flattenPackagesFirewallHyperscalePolicyPoolname6(o["poolname6"], d, "poolname6")); err != nil {
		if vv, ok := fortiAPIPatch(o["poolname6"], "PackagesFirewallHyperscalePolicy-Poolname6"); ok {
			if err = d.Set("poolname6", vv); err != nil {
				return fmt.Errorf("Error reading poolname6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading poolname6: %v", err)
		}
	}

	if err = d.Set("send_deny_packet", flattenPackagesFirewallHyperscalePolicySendDenyPacket(o["send-deny-packet"], d, "send_deny_packet")); err != nil {
		if vv, ok := fortiAPIPatch(o["send-deny-packet"], "PackagesFirewallHyperscalePolicy-SendDenyPacket"); ok {
			if err = d.Set("send_deny_packet", vv); err != nil {
				return fmt.Errorf("Error reading send_deny_packet: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading send_deny_packet: %v", err)
		}
	}

	if err = d.Set("service", flattenPackagesFirewallHyperscalePolicyService(o["service"], d, "service")); err != nil {
		if vv, ok := fortiAPIPatch(o["service"], "PackagesFirewallHyperscalePolicy-Service"); ok {
			if err = d.Set("service", vv); err != nil {
				return fmt.Errorf("Error reading service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service: %v", err)
		}
	}

	if err = d.Set("service_negate", flattenPackagesFirewallHyperscalePolicyServiceNegate(o["service-negate"], d, "service_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["service-negate"], "PackagesFirewallHyperscalePolicy-ServiceNegate"); ok {
			if err = d.Set("service_negate", vv); err != nil {
				return fmt.Errorf("Error reading service_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service_negate: %v", err)
		}
	}

	if err = d.Set("srcaddr", flattenPackagesFirewallHyperscalePolicySrcaddr(o["srcaddr"], d, "srcaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr"], "PackagesFirewallHyperscalePolicy-Srcaddr"); ok {
			if err = d.Set("srcaddr", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr: %v", err)
		}
	}

	if err = d.Set("srcaddr_negate", flattenPackagesFirewallHyperscalePolicySrcaddrNegate(o["srcaddr-negate"], d, "srcaddr_negate")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr-negate"], "PackagesFirewallHyperscalePolicy-SrcaddrNegate"); ok {
			if err = d.Set("srcaddr_negate", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr_negate: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr_negate: %v", err)
		}
	}

	if err = d.Set("srcaddr6", flattenPackagesFirewallHyperscalePolicySrcaddr6(o["srcaddr6"], d, "srcaddr6")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr6"], "PackagesFirewallHyperscalePolicy-Srcaddr6"); ok {
			if err = d.Set("srcaddr6", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr6: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr6: %v", err)
		}
	}

	if err = d.Set("srcintf", flattenPackagesFirewallHyperscalePolicySrcintf(o["srcintf"], d, "srcintf")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcintf"], "PackagesFirewallHyperscalePolicy-Srcintf"); ok {
			if err = d.Set("srcintf", vv); err != nil {
				return fmt.Errorf("Error reading srcintf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcintf: %v", err)
		}
	}

	if err = d.Set("status", flattenPackagesFirewallHyperscalePolicyStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "PackagesFirewallHyperscalePolicy-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("tcp_timeout_pid", flattenPackagesFirewallHyperscalePolicyTcpTimeoutPid(o["tcp-timeout-pid"], d, "tcp_timeout_pid")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-timeout-pid"], "PackagesFirewallHyperscalePolicy-TcpTimeoutPid"); ok {
			if err = d.Set("tcp_timeout_pid", vv); err != nil {
				return fmt.Errorf("Error reading tcp_timeout_pid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_timeout_pid: %v", err)
		}
	}

	if err = d.Set("traffic_shaper", flattenPackagesFirewallHyperscalePolicyTrafficShaper(o["traffic-shaper"], d, "traffic_shaper")); err != nil {
		if vv, ok := fortiAPIPatch(o["traffic-shaper"], "PackagesFirewallHyperscalePolicy-TrafficShaper"); ok {
			if err = d.Set("traffic_shaper", vv); err != nil {
				return fmt.Errorf("Error reading traffic_shaper: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading traffic_shaper: %v", err)
		}
	}

	if err = d.Set("traffic_shaper_reverse", flattenPackagesFirewallHyperscalePolicyTrafficShaperReverse(o["traffic-shaper-reverse"], d, "traffic_shaper_reverse")); err != nil {
		if vv, ok := fortiAPIPatch(o["traffic-shaper-reverse"], "PackagesFirewallHyperscalePolicy-TrafficShaperReverse"); ok {
			if err = d.Set("traffic_shaper_reverse", vv); err != nil {
				return fmt.Errorf("Error reading traffic_shaper_reverse: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading traffic_shaper_reverse: %v", err)
		}
	}

	if err = d.Set("udp_timeout_pid", flattenPackagesFirewallHyperscalePolicyUdpTimeoutPid(o["udp-timeout-pid"], d, "udp_timeout_pid")); err != nil {
		if vv, ok := fortiAPIPatch(o["udp-timeout-pid"], "PackagesFirewallHyperscalePolicy-UdpTimeoutPid"); ok {
			if err = d.Set("udp_timeout_pid", vv); err != nil {
				return fmt.Errorf("Error reading udp_timeout_pid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading udp_timeout_pid: %v", err)
		}
	}

	if err = d.Set("uuid", flattenPackagesFirewallHyperscalePolicyUuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "PackagesFirewallHyperscalePolicy-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	return nil
}

func flattenPackagesFirewallHyperscalePolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandPackagesFirewallHyperscalePolicyAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallHyperscalePolicyAutoAsicOffload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallHyperscalePolicyCgnEif(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallHyperscalePolicyCgnEim(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallHyperscalePolicyCgnLogServerGrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallHyperscalePolicyCgnResourceQuota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallHyperscalePolicyCgnSessionQuota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallHyperscalePolicyComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallHyperscalePolicyDelayTcpNpuSession(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallHyperscalePolicyDstaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallHyperscalePolicyDstaddrNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallHyperscalePolicyDstaddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallHyperscalePolicyDstintf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallHyperscalePolicyFirewallSessionDirty(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallHyperscalePolicyGlobalLabel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallHyperscalePolicyIppool(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallHyperscalePolicyLabel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallHyperscalePolicyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallHyperscalePolicyNat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallHyperscalePolicyPolicyOffload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallHyperscalePolicyPolicyid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallHyperscalePolicyPoolname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallHyperscalePolicyPoolname6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallHyperscalePolicySendDenyPacket(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallHyperscalePolicyService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallHyperscalePolicyServiceNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallHyperscalePolicySrcaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallHyperscalePolicySrcaddrNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallHyperscalePolicySrcaddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallHyperscalePolicySrcintf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallHyperscalePolicyStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallHyperscalePolicyTcpTimeoutPid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallHyperscalePolicyTrafficShaper(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallHyperscalePolicyTrafficShaperReverse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallHyperscalePolicyUdpTimeoutPid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallHyperscalePolicyUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectPackagesFirewallHyperscalePolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandPackagesFirewallHyperscalePolicyAction(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("auto_asic_offload"); ok || d.HasChange("auto_asic_offload") {
		t, err := expandPackagesFirewallHyperscalePolicyAutoAsicOffload(d, v, "auto_asic_offload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-asic-offload"] = t
		}
	}

	if v, ok := d.GetOk("cgn_eif"); ok || d.HasChange("cgn_eif") {
		t, err := expandPackagesFirewallHyperscalePolicyCgnEif(d, v, "cgn_eif")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cgn-eif"] = t
		}
	}

	if v, ok := d.GetOk("cgn_eim"); ok || d.HasChange("cgn_eim") {
		t, err := expandPackagesFirewallHyperscalePolicyCgnEim(d, v, "cgn_eim")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cgn-eim"] = t
		}
	}

	if v, ok := d.GetOk("cgn_log_server_grp"); ok || d.HasChange("cgn_log_server_grp") {
		t, err := expandPackagesFirewallHyperscalePolicyCgnLogServerGrp(d, v, "cgn_log_server_grp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cgn-log-server-grp"] = t
		}
	}

	if v, ok := d.GetOk("cgn_resource_quota"); ok || d.HasChange("cgn_resource_quota") {
		t, err := expandPackagesFirewallHyperscalePolicyCgnResourceQuota(d, v, "cgn_resource_quota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cgn-resource-quota"] = t
		}
	}

	if v, ok := d.GetOk("cgn_session_quota"); ok || d.HasChange("cgn_session_quota") {
		t, err := expandPackagesFirewallHyperscalePolicyCgnSessionQuota(d, v, "cgn_session_quota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cgn-session-quota"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok || d.HasChange("comments") {
		t, err := expandPackagesFirewallHyperscalePolicyComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("delay_tcp_npu_session"); ok || d.HasChange("delay_tcp_npu_session") {
		t, err := expandPackagesFirewallHyperscalePolicyDelayTcpNpuSession(d, v, "delay_tcp_npu_session")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["delay-tcp-npu-session"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr"); ok || d.HasChange("dstaddr") {
		t, err := expandPackagesFirewallHyperscalePolicyDstaddr(d, v, "dstaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr_negate"); ok || d.HasChange("dstaddr_negate") {
		t, err := expandPackagesFirewallHyperscalePolicyDstaddrNegate(d, v, "dstaddr_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr-negate"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr6"); ok || d.HasChange("dstaddr6") {
		t, err := expandPackagesFirewallHyperscalePolicyDstaddr6(d, v, "dstaddr6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr6"] = t
		}
	}

	if v, ok := d.GetOk("dstintf"); ok || d.HasChange("dstintf") {
		t, err := expandPackagesFirewallHyperscalePolicyDstintf(d, v, "dstintf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstintf"] = t
		}
	}

	if v, ok := d.GetOk("firewall_session_dirty"); ok || d.HasChange("firewall_session_dirty") {
		t, err := expandPackagesFirewallHyperscalePolicyFirewallSessionDirty(d, v, "firewall_session_dirty")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["firewall-session-dirty"] = t
		}
	}

	if v, ok := d.GetOk("global_label"); ok || d.HasChange("global_label") {
		t, err := expandPackagesFirewallHyperscalePolicyGlobalLabel(d, v, "global_label")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["global-label"] = t
		}
	}

	if v, ok := d.GetOk("ippool"); ok || d.HasChange("ippool") {
		t, err := expandPackagesFirewallHyperscalePolicyIppool(d, v, "ippool")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ippool"] = t
		}
	}

	if v, ok := d.GetOk("label"); ok || d.HasChange("label") {
		t, err := expandPackagesFirewallHyperscalePolicyLabel(d, v, "label")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["label"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandPackagesFirewallHyperscalePolicyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("nat"); ok || d.HasChange("nat") {
		t, err := expandPackagesFirewallHyperscalePolicyNat(d, v, "nat")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat"] = t
		}
	}

	if v, ok := d.GetOk("policy_offload"); ok || d.HasChange("policy_offload") {
		t, err := expandPackagesFirewallHyperscalePolicyPolicyOffload(d, v, "policy_offload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy-offload"] = t
		}
	}

	if v, ok := d.GetOk("policyid"); ok || d.HasChange("policyid") {
		t, err := expandPackagesFirewallHyperscalePolicyPolicyid(d, v, "policyid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policyid"] = t
		}
	}

	if v, ok := d.GetOk("poolname"); ok || d.HasChange("poolname") {
		t, err := expandPackagesFirewallHyperscalePolicyPoolname(d, v, "poolname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poolname"] = t
		}
	}

	if v, ok := d.GetOk("poolname6"); ok || d.HasChange("poolname6") {
		t, err := expandPackagesFirewallHyperscalePolicyPoolname6(d, v, "poolname6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poolname6"] = t
		}
	}

	if v, ok := d.GetOk("send_deny_packet"); ok || d.HasChange("send_deny_packet") {
		t, err := expandPackagesFirewallHyperscalePolicySendDenyPacket(d, v, "send_deny_packet")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send-deny-packet"] = t
		}
	}

	if v, ok := d.GetOk("service"); ok || d.HasChange("service") {
		t, err := expandPackagesFirewallHyperscalePolicyService(d, v, "service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	if v, ok := d.GetOk("service_negate"); ok || d.HasChange("service_negate") {
		t, err := expandPackagesFirewallHyperscalePolicyServiceNegate(d, v, "service_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-negate"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr"); ok || d.HasChange("srcaddr") {
		t, err := expandPackagesFirewallHyperscalePolicySrcaddr(d, v, "srcaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr_negate"); ok || d.HasChange("srcaddr_negate") {
		t, err := expandPackagesFirewallHyperscalePolicySrcaddrNegate(d, v, "srcaddr_negate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr-negate"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr6"); ok || d.HasChange("srcaddr6") {
		t, err := expandPackagesFirewallHyperscalePolicySrcaddr6(d, v, "srcaddr6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr6"] = t
		}
	}

	if v, ok := d.GetOk("srcintf"); ok || d.HasChange("srcintf") {
		t, err := expandPackagesFirewallHyperscalePolicySrcintf(d, v, "srcintf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcintf"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandPackagesFirewallHyperscalePolicyStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("tcp_timeout_pid"); ok || d.HasChange("tcp_timeout_pid") {
		t, err := expandPackagesFirewallHyperscalePolicyTcpTimeoutPid(d, v, "tcp_timeout_pid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-timeout-pid"] = t
		}
	}

	if v, ok := d.GetOk("traffic_shaper"); ok || d.HasChange("traffic_shaper") {
		t, err := expandPackagesFirewallHyperscalePolicyTrafficShaper(d, v, "traffic_shaper")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-shaper"] = t
		}
	}

	if v, ok := d.GetOk("traffic_shaper_reverse"); ok || d.HasChange("traffic_shaper_reverse") {
		t, err := expandPackagesFirewallHyperscalePolicyTrafficShaperReverse(d, v, "traffic_shaper_reverse")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-shaper-reverse"] = t
		}
	}

	if v, ok := d.GetOk("udp_timeout_pid"); ok || d.HasChange("udp_timeout_pid") {
		t, err := expandPackagesFirewallHyperscalePolicyUdpTimeoutPid(d, v, "udp_timeout_pid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["udp-timeout-pid"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok || d.HasChange("uuid") {
		t, err := expandPackagesFirewallHyperscalePolicyUuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	return &obj, nil
}
