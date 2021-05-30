// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure user defined IPv6 local-in policies.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourcePackagesFirewallLocalInPolicy6() *schema.Resource {
	return &schema.Resource{
		Create: resourcePackagesFirewallLocalInPolicy6Create,
		Read:   resourcePackagesFirewallLocalInPolicy6Read,
		Update: resourcePackagesFirewallLocalInPolicy6Update,
		Delete: resourcePackagesFirewallLocalInPolicy6Delete,

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
			"intf": &schema.Schema{
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

func resourcePackagesFirewallLocalInPolicy6Create(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectPackagesFirewallLocalInPolicy6(d)
	if err != nil {
		return fmt.Errorf("Error creating PackagesFirewallLocalInPolicy6 resource while getting object: %v", err)
	}

	_, err = c.CreatePackagesFirewallLocalInPolicy6(obj, adomv, paralist)

	if err != nil {
		return fmt.Errorf("Error creating PackagesFirewallLocalInPolicy6 resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "policyid")))

	return resourcePackagesFirewallLocalInPolicy6Read(d, m)
}

func resourcePackagesFirewallLocalInPolicy6Update(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectPackagesFirewallLocalInPolicy6(d)
	if err != nil {
		return fmt.Errorf("Error updating PackagesFirewallLocalInPolicy6 resource while getting object: %v", err)
	}

	_, err = c.UpdatePackagesFirewallLocalInPolicy6(obj, adomv, mkey, paralist)
	if err != nil {
		return fmt.Errorf("Error updating PackagesFirewallLocalInPolicy6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "policyid")))

	return resourcePackagesFirewallLocalInPolicy6Read(d, m)
}

func resourcePackagesFirewallLocalInPolicy6Delete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeletePackagesFirewallLocalInPolicy6(adomv, mkey, paralist)
	if err != nil {
		return fmt.Errorf("Error deleting PackagesFirewallLocalInPolicy6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourcePackagesFirewallLocalInPolicy6Read(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadPackagesFirewallLocalInPolicy6(adomv, mkey, paralist)
	if err != nil {
		return fmt.Errorf("Error reading PackagesFirewallLocalInPolicy6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectPackagesFirewallLocalInPolicy6(d, o)
	if err != nil {
		return fmt.Errorf("Error reading PackagesFirewallLocalInPolicy6 resource from API: %v", err)
	}
	return nil
}

func flattenPackagesFirewallLocalInPolicy6Action(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenPackagesFirewallLocalInPolicy6Comments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallLocalInPolicy6Dstaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallLocalInPolicy6Intf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallLocalInPolicy6Policyid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallLocalInPolicy6Schedule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallLocalInPolicy6Service(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallLocalInPolicy6Srcaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallLocalInPolicy6Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func flattenPackagesFirewallLocalInPolicy6Uuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectPackagesFirewallLocalInPolicy6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("action", flattenPackagesFirewallLocalInPolicy6Action(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "PackagesFirewallLocalInPolicy6-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("comments", flattenPackagesFirewallLocalInPolicy6Comments(o["comments"], d, "comments")); err != nil {
		if vv, ok := fortiAPIPatch(o["comments"], "PackagesFirewallLocalInPolicy6-Comments"); ok {
			if err = d.Set("comments", vv); err != nil {
				return fmt.Errorf("Error reading comments: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("dstaddr", flattenPackagesFirewallLocalInPolicy6Dstaddr(o["dstaddr"], d, "dstaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["dstaddr"], "PackagesFirewallLocalInPolicy6-Dstaddr"); ok {
			if err = d.Set("dstaddr", vv); err != nil {
				return fmt.Errorf("Error reading dstaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dstaddr: %v", err)
		}
	}

	if err = d.Set("intf", flattenPackagesFirewallLocalInPolicy6Intf(o["intf"], d, "intf")); err != nil {
		if vv, ok := fortiAPIPatch(o["intf"], "PackagesFirewallLocalInPolicy6-Intf"); ok {
			if err = d.Set("intf", vv); err != nil {
				return fmt.Errorf("Error reading intf: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading intf: %v", err)
		}
	}

	if err = d.Set("policyid", flattenPackagesFirewallLocalInPolicy6Policyid(o["policyid"], d, "policyid")); err != nil {
		if vv, ok := fortiAPIPatch(o["policyid"], "PackagesFirewallLocalInPolicy6-Policyid"); ok {
			if err = d.Set("policyid", vv); err != nil {
				return fmt.Errorf("Error reading policyid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading policyid: %v", err)
		}
	}

	if err = d.Set("schedule", flattenPackagesFirewallLocalInPolicy6Schedule(o["schedule"], d, "schedule")); err != nil {
		if vv, ok := fortiAPIPatch(o["schedule"], "PackagesFirewallLocalInPolicy6-Schedule"); ok {
			if err = d.Set("schedule", vv); err != nil {
				return fmt.Errorf("Error reading schedule: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading schedule: %v", err)
		}
	}

	if err = d.Set("service", flattenPackagesFirewallLocalInPolicy6Service(o["service"], d, "service")); err != nil {
		if vv, ok := fortiAPIPatch(o["service"], "PackagesFirewallLocalInPolicy6-Service"); ok {
			if err = d.Set("service", vv); err != nil {
				return fmt.Errorf("Error reading service: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service: %v", err)
		}
	}

	if err = d.Set("srcaddr", flattenPackagesFirewallLocalInPolicy6Srcaddr(o["srcaddr"], d, "srcaddr")); err != nil {
		if vv, ok := fortiAPIPatch(o["srcaddr"], "PackagesFirewallLocalInPolicy6-Srcaddr"); ok {
			if err = d.Set("srcaddr", vv); err != nil {
				return fmt.Errorf("Error reading srcaddr: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading srcaddr: %v", err)
		}
	}

	if err = d.Set("status", flattenPackagesFirewallLocalInPolicy6Status(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "PackagesFirewallLocalInPolicy6-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("uuid", flattenPackagesFirewallLocalInPolicy6Uuid(o["uuid"], d, "uuid")); err != nil {
		if vv, ok := fortiAPIPatch(o["uuid"], "PackagesFirewallLocalInPolicy6-Uuid"); ok {
			if err = d.Set("uuid", vv); err != nil {
				return fmt.Errorf("Error reading uuid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	return nil
}

func flattenPackagesFirewallLocalInPolicy6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandPackagesFirewallLocalInPolicy6Action(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallLocalInPolicy6Comments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallLocalInPolicy6Dstaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallLocalInPolicy6Intf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallLocalInPolicy6Policyid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallLocalInPolicy6Schedule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallLocalInPolicy6Service(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallLocalInPolicy6Srcaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallLocalInPolicy6Status(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallLocalInPolicy6Uuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectPackagesFirewallLocalInPolicy6(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok {
		t, err := expandPackagesFirewallLocalInPolicy6Action(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {
		t, err := expandPackagesFirewallLocalInPolicy6Comments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr"); ok {
		t, err := expandPackagesFirewallLocalInPolicy6Dstaddr(d, v, "dstaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr"] = t
		}
	}

	if v, ok := d.GetOk("intf"); ok {
		t, err := expandPackagesFirewallLocalInPolicy6Intf(d, v, "intf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["intf"] = t
		}
	}

	if v, ok := d.GetOk("policyid"); ok {
		t, err := expandPackagesFirewallLocalInPolicy6Policyid(d, v, "policyid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policyid"] = t
		}
	}

	if v, ok := d.GetOk("schedule"); ok {
		t, err := expandPackagesFirewallLocalInPolicy6Schedule(d, v, "schedule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule"] = t
		}
	}

	if v, ok := d.GetOk("service"); ok {
		t, err := expandPackagesFirewallLocalInPolicy6Service(d, v, "service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr"); ok {
		t, err := expandPackagesFirewallLocalInPolicy6Srcaddr(d, v, "srcaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandPackagesFirewallLocalInPolicy6Status(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok {
		t, err := expandPackagesFirewallLocalInPolicy6Uuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	return &obj, nil
}
