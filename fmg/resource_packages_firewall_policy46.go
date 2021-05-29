// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure IPv4 to IPv6 policies.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourcePackagesFirewallPolicy46() *schema.Resource {
	return &schema.Resource{
		Create: resourcePackagesFirewallPolicy46Create,
		Read:   resourcePackagesFirewallPolicy46Read,
		Update: resourcePackagesFirewallPolicy46Update,
		Delete: resourcePackagesFirewallPolicy46Delete,

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
			"dstintf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fixedport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ippool": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logtraffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logtraffic_start": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"per_ip_shaper": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"permit_any_host": &schema.Schema{
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
			"poolname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"schedule": &schema.Schema{
				Type:     schema.TypeString,
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
			"srcintf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_mss_receiver": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tcp_mss_sender": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"traffic_shaper": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"traffic_shaper_reverse": &schema.Schema{
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

func resourcePackagesFirewallPolicy46Create(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectPackagesFirewallPolicy46(d)
	if err != nil {
		return fmt.Errorf("Error creating PackagesFirewallPolicy46 resource while getting object: %v", err)
	}

	_, err = c.CreatePackagesFirewallPolicy46(obj, adomv, paralist)

	if err != nil {
		return fmt.Errorf("Error creating PackagesFirewallPolicy46 resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "policyid")))

	return resourcePackagesFirewallPolicy46Read(d, m)
}

func resourcePackagesFirewallPolicy46Update(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectPackagesFirewallPolicy46(d)
	if err != nil {
		return fmt.Errorf("Error updating PackagesFirewallPolicy46 resource while getting object: %v", err)
	}

	_, err = c.UpdatePackagesFirewallPolicy46(obj, adomv, mkey, paralist)
	if err != nil {
		return fmt.Errorf("Error updating PackagesFirewallPolicy46 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "policyid")))

	return resourcePackagesFirewallPolicy46Read(d, m)
}

func resourcePackagesFirewallPolicy46Delete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeletePackagesFirewallPolicy46(adomv, mkey, paralist)
	if err != nil {
		return fmt.Errorf("Error deleting PackagesFirewallPolicy46 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourcePackagesFirewallPolicy46Read(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadPackagesFirewallPolicy46(adomv, mkey, paralist)
	if err != nil {
		return fmt.Errorf("Error reading PackagesFirewallPolicy46 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectPackagesFirewallPolicy46(d, o)
	if err != nil {
		return fmt.Errorf("Error reading PackagesFirewallPolicy46 resource from API: %v", err)
	}
	return nil
}

func flattenPackagesFirewallPolicy46Action(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenPackagesFirewallPolicy46Comments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicy46Dstaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicy46Dstintf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicy46Fixedport(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenPackagesFirewallPolicy46Ippool(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenPackagesFirewallPolicy46Logtraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenPackagesFirewallPolicy46LogtrafficStart(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenPackagesFirewallPolicy46Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicy46PerIpShaper(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicy46PermitAnyHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenPackagesFirewallPolicy46Policyid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicy46Poolname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicy46Schedule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicy46Service(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicy46Srcaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicy46Srcintf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicy46Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenPackagesFirewallPolicy46TcpMssReceiver(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicy46TcpMssSender(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicy46TrafficShaper(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicy46TrafficShaperReverse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallPolicy46Uuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectPackagesFirewallPolicy46(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("action", flattenPackagesFirewallPolicy46Action(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "PackagesFirewallPolicy46-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("comments", flattenPackagesFirewallPolicy46Comments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "PackagesFirewallPolicy46-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("dstaddr", flattenPackagesFirewallPolicy46Dstaddr(o["dstaddr"], d, "dstaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr"], "PackagesFirewallPolicy46-Dstaddr"); ok {
			if err = d.Set("dstaddr", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr: %v", err)
		}
	}

	if err = d.Set("dstintf", flattenPackagesFirewallPolicy46Dstintf(o["dstintf"], d, "dstintf")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstintf"], "PackagesFirewallPolicy46-Dstintf"); ok {
			if err = d.Set("dstintf", vv); err != nil {
				return fmt.Errorf("Error reading dstintf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstintf: %v", err)
		}
	}

	if err = d.Set("fixedport", flattenPackagesFirewallPolicy46Fixedport(o["fixedport"], d, "fixedport")); err != nil {
		if vv, ok := fortiAPIPatch(o["fixedport"], "PackagesFirewallPolicy46-Fixedport"); ok {
			if err = d.Set("fixedport", vv); err != nil {
				return fmt.Errorf("Error reading fixedport: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fixedport: %v", err)
		}
	}

	if err = d.Set("ippool", flattenPackagesFirewallPolicy46Ippool(o["ippool"], d, "ippool")); err != nil {
		if vv, ok := fortiAPIPatch(o["ippool"], "PackagesFirewallPolicy46-Ippool"); ok {
			if err = d.Set("ippool", vv); err != nil {
				return fmt.Errorf("Error reading ippool: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ippool: %v", err)
		}
	}

	if err = d.Set("logtraffic", flattenPackagesFirewallPolicy46Logtraffic(o["logtraffic"], d, "logtraffic")); err != nil {
		if vv, ok := fortiAPIPatch(o["logtraffic"], "PackagesFirewallPolicy46-Logtraffic"); ok {
			if err = d.Set("logtraffic", vv); err != nil {
				return fmt.Errorf("Error reading logtraffic: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logtraffic: %v", err)
		}
	}

	if err = d.Set("logtraffic_start", flattenPackagesFirewallPolicy46LogtrafficStart(o["logtraffic-start"], d, "logtraffic_start")); err != nil {
		if vv, ok := fortiAPIPatch(o["logtraffic-start"], "PackagesFirewallPolicy46-LogtrafficStart"); ok {
			if err = d.Set("logtraffic_start", vv); err != nil {
				return fmt.Errorf("Error reading logtraffic_start: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading logtraffic_start: %v", err)
		}
	}

	if err = d.Set("name", flattenPackagesFirewallPolicy46Name(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "PackagesFirewallPolicy46-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("per_ip_shaper", flattenPackagesFirewallPolicy46PerIpShaper(o["per-ip-shaper"], d, "per_ip_shaper")); err != nil {
		if vv, ok := fortiAPIPatch(o["per-ip-shaper"], "PackagesFirewallPolicy46-PerIpShaper"); ok {
			if err = d.Set("per_ip_shaper", vv); err != nil {
				return fmt.Errorf("Error reading per_ip_shaper: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading per_ip_shaper: %v", err)
		}
	}

	if err = d.Set("permit_any_host", flattenPackagesFirewallPolicy46PermitAnyHost(o["permit-any-host"], d, "permit_any_host")); err != nil {
		if vv, ok := fortiAPIPatch(o["permit-any-host"], "PackagesFirewallPolicy46-PermitAnyHost"); ok {
			if err = d.Set("permit_any_host", vv); err != nil {
				return fmt.Errorf("Error reading permit_any_host: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading permit_any_host: %v", err)
		}
	}

	if err = d.Set("policyid", flattenPackagesFirewallPolicy46Policyid(o["policyid"], d, "policyid")); err != nil {
		if vv, ok := fortiAPIPatch(o["policyid"], "PackagesFirewallPolicy46-Policyid"); ok {
			if err = d.Set("policyid", vv); err != nil {
				return fmt.Errorf("Error reading policyid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policyid: %v", err)
		}
	}

	if err = d.Set("poolname", flattenPackagesFirewallPolicy46Poolname(o["poolname"], d, "poolname")); err != nil {
		if vv, ok := fortiAPIPatch(o["poolname"], "PackagesFirewallPolicy46-Poolname"); ok {
			if err = d.Set("poolname", vv); err != nil {
				return fmt.Errorf("Error reading poolname: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading poolname: %v", err)
		}
	}

	if err = d.Set("schedule", flattenPackagesFirewallPolicy46Schedule(o["schedule"], d, "schedule")); err != nil {
		if vv, ok := fortiAPIPatch(o["schedule"], "PackagesFirewallPolicy46-Schedule"); ok {
			if err = d.Set("schedule", vv); err != nil {
				return fmt.Errorf("Error reading schedule: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading schedule: %v", err)
		}
	}

	if err = d.Set("service", flattenPackagesFirewallPolicy46Service(o["service"], d, "service")); err != nil {
		if vv, ok := fortiAPIPatch(o["service"], "PackagesFirewallPolicy46-Service"); ok {
			if err = d.Set("service", vv); err != nil {
				return fmt.Errorf("Error reading service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service: %v", err)
		}
	}

	if err = d.Set("srcaddr", flattenPackagesFirewallPolicy46Srcaddr(o["srcaddr"], d, "srcaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr"], "PackagesFirewallPolicy46-Srcaddr"); ok {
			if err = d.Set("srcaddr", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr: %v", err)
		}
	}

	if err = d.Set("srcintf", flattenPackagesFirewallPolicy46Srcintf(o["srcintf"], d, "srcintf")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcintf"], "PackagesFirewallPolicy46-Srcintf"); ok {
			if err = d.Set("srcintf", vv); err != nil {
				return fmt.Errorf("Error reading srcintf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcintf: %v", err)
		}
	}

	if err = d.Set("status", flattenPackagesFirewallPolicy46Status(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "PackagesFirewallPolicy46-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("tcp_mss_receiver", flattenPackagesFirewallPolicy46TcpMssReceiver(o["tcp-mss-receiver"], d, "tcp_mss_receiver")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-mss-receiver"], "PackagesFirewallPolicy46-TcpMssReceiver"); ok {
			if err = d.Set("tcp_mss_receiver", vv); err != nil {
				return fmt.Errorf("Error reading tcp_mss_receiver: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_mss_receiver: %v", err)
		}
	}

	if err = d.Set("tcp_mss_sender", flattenPackagesFirewallPolicy46TcpMssSender(o["tcp-mss-sender"], d, "tcp_mss_sender")); err != nil {
		if vv, ok := fortiAPIPatch(o["tcp-mss-sender"], "PackagesFirewallPolicy46-TcpMssSender"); ok {
			if err = d.Set("tcp_mss_sender", vv); err != nil {
				return fmt.Errorf("Error reading tcp_mss_sender: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading tcp_mss_sender: %v", err)
		}
	}

	if err = d.Set("traffic_shaper", flattenPackagesFirewallPolicy46TrafficShaper(o["traffic-shaper"], d, "traffic_shaper")); err != nil {
		if vv, ok := fortiAPIPatch(o["traffic-shaper"], "PackagesFirewallPolicy46-TrafficShaper"); ok {
			if err = d.Set("traffic_shaper", vv); err != nil {
				return fmt.Errorf("Error reading traffic_shaper: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading traffic_shaper: %v", err)
		}
	}

	if err = d.Set("traffic_shaper_reverse", flattenPackagesFirewallPolicy46TrafficShaperReverse(o["traffic-shaper-reverse"], d, "traffic_shaper_reverse")); err != nil {
		if vv, ok := fortiAPIPatch(o["traffic-shaper-reverse"], "PackagesFirewallPolicy46-TrafficShaperReverse"); ok {
			if err = d.Set("traffic_shaper_reverse", vv); err != nil {
				return fmt.Errorf("Error reading traffic_shaper_reverse: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading traffic_shaper_reverse: %v", err)
		}
	}

	if err = d.Set("uuid", flattenPackagesFirewallPolicy46Uuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "PackagesFirewallPolicy46-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	return nil
}

func flattenPackagesFirewallPolicy46FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandPackagesFirewallPolicy46Action(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicy46Comments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicy46Dstaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicy46Dstintf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicy46Fixedport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicy46Ippool(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicy46Logtraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicy46LogtrafficStart(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicy46Name(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicy46PerIpShaper(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicy46PermitAnyHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicy46Policyid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicy46Poolname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicy46Schedule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicy46Service(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicy46Srcaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicy46Srcintf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicy46Status(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicy46TcpMssReceiver(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicy46TcpMssSender(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicy46TrafficShaper(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicy46TrafficShaperReverse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallPolicy46Uuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectPackagesFirewallPolicy46(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok {
		t, err := expandPackagesFirewallPolicy46Action(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {
		t, err := expandPackagesFirewallPolicy46Comments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr"); ok {
		t, err := expandPackagesFirewallPolicy46Dstaddr(d, v, "dstaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr"] = t
		}
	}

	if v, ok := d.GetOk("dstintf"); ok {
		t, err := expandPackagesFirewallPolicy46Dstintf(d, v, "dstintf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstintf"] = t
		}
	}

	if v, ok := d.GetOk("fixedport"); ok {
		t, err := expandPackagesFirewallPolicy46Fixedport(d, v, "fixedport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fixedport"] = t
		}
	}

	if v, ok := d.GetOk("ippool"); ok {
		t, err := expandPackagesFirewallPolicy46Ippool(d, v, "ippool")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ippool"] = t
		}
	}

	if v, ok := d.GetOk("logtraffic"); ok {
		t, err := expandPackagesFirewallPolicy46Logtraffic(d, v, "logtraffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logtraffic"] = t
		}
	}

	if v, ok := d.GetOk("logtraffic_start"); ok {
		t, err := expandPackagesFirewallPolicy46LogtrafficStart(d, v, "logtraffic_start")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logtraffic-start"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandPackagesFirewallPolicy46Name(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("per_ip_shaper"); ok {
		t, err := expandPackagesFirewallPolicy46PerIpShaper(d, v, "per_ip_shaper")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["per-ip-shaper"] = t
		}
	}

	if v, ok := d.GetOk("permit_any_host"); ok {
		t, err := expandPackagesFirewallPolicy46PermitAnyHost(d, v, "permit_any_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["permit-any-host"] = t
		}
	}

	if v, ok := d.GetOk("policyid"); ok {
		t, err := expandPackagesFirewallPolicy46Policyid(d, v, "policyid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policyid"] = t
		}
	}

	if v, ok := d.GetOk("poolname"); ok {
		t, err := expandPackagesFirewallPolicy46Poolname(d, v, "poolname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poolname"] = t
		}
	}

	if v, ok := d.GetOk("schedule"); ok {
		t, err := expandPackagesFirewallPolicy46Schedule(d, v, "schedule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule"] = t
		}
	}

	if v, ok := d.GetOk("service"); ok {
		t, err := expandPackagesFirewallPolicy46Service(d, v, "service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr"); ok {
		t, err := expandPackagesFirewallPolicy46Srcaddr(d, v, "srcaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr"] = t
		}
	}

	if v, ok := d.GetOk("srcintf"); ok {
		t, err := expandPackagesFirewallPolicy46Srcintf(d, v, "srcintf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcintf"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandPackagesFirewallPolicy46Status(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("tcp_mss_receiver"); ok {
		t, err := expandPackagesFirewallPolicy46TcpMssReceiver(d, v, "tcp_mss_receiver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-mss-receiver"] = t
		}
	}

	if v, ok := d.GetOk("tcp_mss_sender"); ok {
		t, err := expandPackagesFirewallPolicy46TcpMssSender(d, v, "tcp_mss_sender")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-mss-sender"] = t
		}
	}

	if v, ok := d.GetOk("traffic_shaper"); ok {
		t, err := expandPackagesFirewallPolicy46TrafficShaper(d, v, "traffic_shaper")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-shaper"] = t
		}
	}

	if v, ok := d.GetOk("traffic_shaper_reverse"); ok {
		t, err := expandPackagesFirewallPolicy46TrafficShaperReverse(d, v, "traffic_shaper_reverse")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-shaper-reverse"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok {
		t, err := expandPackagesFirewallPolicy46Uuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	return &obj, nil
}
