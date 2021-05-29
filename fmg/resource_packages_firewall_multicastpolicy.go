// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure multicast NAT policies.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourcePackagesFirewallMulticastPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourcePackagesFirewallMulticastPolicyCreate,
		Read:   resourcePackagesFirewallMulticastPolicyRead,
		Update: resourcePackagesFirewallMulticastPolicyUpdate,
		Delete: resourcePackagesFirewallMulticastPolicyDelete,

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
				Computed: true,
			},
			"auto_asic_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dnat": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dstaddr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dstintf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"end_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"logtraffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"snat": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"snat_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcaddr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcintf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"start_port": &schema.Schema{
				Type:     schema.TypeInt,
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
		},
	}
}

func resourcePackagesFirewallMulticastPolicyCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectPackagesFirewallMulticastPolicy(d)
	if err != nil {
		return fmt.Errorf("Error creating PackagesFirewallMulticastPolicy resource while getting object: %v", err)
	}

	_, err = c.CreatePackagesFirewallMulticastPolicy(obj, adomv, paralist)

	if err != nil {
		return fmt.Errorf("Error creating PackagesFirewallMulticastPolicy resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourcePackagesFirewallMulticastPolicyRead(d, m)
}

func resourcePackagesFirewallMulticastPolicyUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectPackagesFirewallMulticastPolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating PackagesFirewallMulticastPolicy resource while getting object: %v", err)
	}

	_, err = c.UpdatePackagesFirewallMulticastPolicy(obj, adomv, mkey, paralist)
	if err != nil {
		return fmt.Errorf("Error updating PackagesFirewallMulticastPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourcePackagesFirewallMulticastPolicyRead(d, m)
}

func resourcePackagesFirewallMulticastPolicyDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeletePackagesFirewallMulticastPolicy(adomv, mkey, paralist)
	if err != nil {
		return fmt.Errorf("Error deleting PackagesFirewallMulticastPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourcePackagesFirewallMulticastPolicyRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadPackagesFirewallMulticastPolicy(adomv, mkey, paralist)
	if err != nil {
		return fmt.Errorf("Error reading PackagesFirewallMulticastPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectPackagesFirewallMulticastPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading PackagesFirewallMulticastPolicy resource from API: %v", err)
	}
	return nil
}

func flattenPackagesFirewallMulticastPolicyAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v != nil {
		emap := map[int]string{
			0: "deny",
			1: "accept",
		}
		res := getEnumVal(v, emap)
		return res
	}
	return v
}

func flattenPackagesFirewallMulticastPolicyAutoAsicOffload(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenPackagesFirewallMulticastPolicyComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallMulticastPolicyDnat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallMulticastPolicyDstaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallMulticastPolicyDstintf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallMulticastPolicyEndPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallMulticastPolicyId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallMulticastPolicyLogtraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenPackagesFirewallMulticastPolicyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallMulticastPolicyProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallMulticastPolicySnat(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenPackagesFirewallMulticastPolicySnatIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallMulticastPolicySrcaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallMulticastPolicySrcintf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallMulticastPolicyStartPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallMulticastPolicyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenPackagesFirewallMulticastPolicyUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectPackagesFirewallMulticastPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("action", flattenPackagesFirewallMulticastPolicyAction(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "PackagesFirewallMulticastPolicy-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("auto_asic_offload", flattenPackagesFirewallMulticastPolicyAutoAsicOffload(o["auto-asic-offload"], d, "auto_asic_offload")); err != nil {
		if vv, ok := fortiAPIPatch(o["auto-asic-offload"], "PackagesFirewallMulticastPolicy-AutoAsicOffload"); ok {
			if err = d.Set("auto_asic_offload", vv); err != nil {
				return fmt.Errorf("Error reading auto_asic_offload: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading auto_asic_offload: %v", err)
		}
	}

	if err = d.Set("comments", flattenPackagesFirewallMulticastPolicyComments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "PackagesFirewallMulticastPolicy-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("dnat", flattenPackagesFirewallMulticastPolicyDnat(o["dnat"], d, "dnat")); err != nil {
		if vv, ok := fortiAPIPatch(o["dnat"], "PackagesFirewallMulticastPolicy-Dnat"); ok {
			if err = d.Set("dnat", vv); err != nil {
				return fmt.Errorf("Error reading dnat: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dnat: %v", err)
		}
	}

	if err = d.Set("dstaddr", flattenPackagesFirewallMulticastPolicyDstaddr(o["dstaddr"], d, "dstaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr"], "PackagesFirewallMulticastPolicy-Dstaddr"); ok {
			if err = d.Set("dstaddr", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr: %v", err)
		}
	}

	if err = d.Set("dstintf", flattenPackagesFirewallMulticastPolicyDstintf(o["dstintf"], d, "dstintf")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstintf"], "PackagesFirewallMulticastPolicy-Dstintf"); ok {
			if err = d.Set("dstintf", vv); err != nil {
				return fmt.Errorf("Error reading dstintf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstintf: %v", err)
		}
	}

	if err = d.Set("end_port", flattenPackagesFirewallMulticastPolicyEndPort(o["end-port"], d, "end_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["end-port"], "PackagesFirewallMulticastPolicy-EndPort"); ok {
			if err = d.Set("end_port", vv); err != nil {
				return fmt.Errorf("Error reading end_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading end_port: %v", err)
		}
	}

	if err = d.Set("fosid", flattenPackagesFirewallMulticastPolicyId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "PackagesFirewallMulticastPolicy-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("logtraffic", flattenPackagesFirewallMulticastPolicyLogtraffic(o["logtraffic"], d, "logtraffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["logtraffic"], "PackagesFirewallMulticastPolicy-Logtraffic"); ok {
			if err = d.Set("logtraffic", vv); err != nil {
				return fmt.Errorf("Error reading logtraffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logtraffic: %v", err)
		}
	}

	if err = d.Set("name", flattenPackagesFirewallMulticastPolicyName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "PackagesFirewallMulticastPolicy-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("protocol", flattenPackagesFirewallMulticastPolicyProtocol(o["protocol"], d, "protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["protocol"], "PackagesFirewallMulticastPolicy-Protocol"); ok {
			if err = d.Set("protocol", vv); err != nil {
				return fmt.Errorf("Error reading protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("snat", flattenPackagesFirewallMulticastPolicySnat(o["snat"], d, "snat")); err != nil {
		if vv, ok := fortiAPIPatch(o["snat"], "PackagesFirewallMulticastPolicy-Snat"); ok {
			if err = d.Set("snat", vv); err != nil {
				return fmt.Errorf("Error reading snat: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading snat: %v", err)
		}
	}

	if err = d.Set("snat_ip", flattenPackagesFirewallMulticastPolicySnatIp(o["snat-ip"], d, "snat_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["snat-ip"], "PackagesFirewallMulticastPolicy-SnatIp"); ok {
			if err = d.Set("snat_ip", vv); err != nil {
				return fmt.Errorf("Error reading snat_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading snat_ip: %v", err)
		}
	}

	if err = d.Set("srcaddr", flattenPackagesFirewallMulticastPolicySrcaddr(o["srcaddr"], d, "srcaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr"], "PackagesFirewallMulticastPolicy-Srcaddr"); ok {
			if err = d.Set("srcaddr", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr: %v", err)
		}
	}

	if err = d.Set("srcintf", flattenPackagesFirewallMulticastPolicySrcintf(o["srcintf"], d, "srcintf")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcintf"], "PackagesFirewallMulticastPolicy-Srcintf"); ok {
			if err = d.Set("srcintf", vv); err != nil {
				return fmt.Errorf("Error reading srcintf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcintf: %v", err)
		}
	}

	if err = d.Set("start_port", flattenPackagesFirewallMulticastPolicyStartPort(o["start-port"], d, "start_port")); err != nil {
		if vv, ok := fortiAPIPatch(o["start-port"], "PackagesFirewallMulticastPolicy-StartPort"); ok {
			if err = d.Set("start_port", vv); err != nil {
				return fmt.Errorf("Error reading start_port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading start_port: %v", err)
		}
	}

	if err = d.Set("status", flattenPackagesFirewallMulticastPolicyStatus(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "PackagesFirewallMulticastPolicy-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("uuid", flattenPackagesFirewallMulticastPolicyUuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "PackagesFirewallMulticastPolicy-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	return nil
}

func flattenPackagesFirewallMulticastPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandPackagesFirewallMulticastPolicyAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallMulticastPolicyAutoAsicOffload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallMulticastPolicyComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallMulticastPolicyDnat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallMulticastPolicyDstaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallMulticastPolicyDstintf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallMulticastPolicyEndPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallMulticastPolicyId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallMulticastPolicyLogtraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallMulticastPolicyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallMulticastPolicyProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallMulticastPolicySnat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallMulticastPolicySnatIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallMulticastPolicySrcaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallMulticastPolicySrcintf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallMulticastPolicyStartPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallMulticastPolicyStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallMulticastPolicyUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectPackagesFirewallMulticastPolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok {
		t, err := expandPackagesFirewallMulticastPolicyAction(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("auto_asic_offload"); ok {
		t, err := expandPackagesFirewallMulticastPolicyAutoAsicOffload(d, v, "auto_asic_offload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-asic-offload"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {
		t, err := expandPackagesFirewallMulticastPolicyComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("dnat"); ok {
		t, err := expandPackagesFirewallMulticastPolicyDnat(d, v, "dnat")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dnat"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr"); ok {
		t, err := expandPackagesFirewallMulticastPolicyDstaddr(d, v, "dstaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr"] = t
		}
	}

	if v, ok := d.GetOk("dstintf"); ok {
		t, err := expandPackagesFirewallMulticastPolicyDstintf(d, v, "dstintf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstintf"] = t
		}
	}

	if v, ok := d.GetOk("end_port"); ok {
		t, err := expandPackagesFirewallMulticastPolicyEndPort(d, v, "end_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-port"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok {
		t, err := expandPackagesFirewallMulticastPolicyId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("logtraffic"); ok {
		t, err := expandPackagesFirewallMulticastPolicyLogtraffic(d, v, "logtraffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logtraffic"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandPackagesFirewallMulticastPolicyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok {
		t, err := expandPackagesFirewallMulticastPolicyProtocol(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("snat"); ok {
		t, err := expandPackagesFirewallMulticastPolicySnat(d, v, "snat")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["snat"] = t
		}
	}

	if v, ok := d.GetOk("snat_ip"); ok {
		t, err := expandPackagesFirewallMulticastPolicySnatIp(d, v, "snat_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["snat-ip"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr"); ok {
		t, err := expandPackagesFirewallMulticastPolicySrcaddr(d, v, "srcaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr"] = t
		}
	}

	if v, ok := d.GetOk("srcintf"); ok {
		t, err := expandPackagesFirewallMulticastPolicySrcintf(d, v, "srcintf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcintf"] = t
		}
	}

	if v, ok := d.GetOk("start_port"); ok {
		t, err := expandPackagesFirewallMulticastPolicyStartPort(d, v, "start_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-port"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandPackagesFirewallMulticastPolicyStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok {
		t, err := expandPackagesFirewallMulticastPolicyUuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	return &obj, nil
}
