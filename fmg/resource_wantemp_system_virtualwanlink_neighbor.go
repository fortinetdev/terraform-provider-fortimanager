// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Create SD-WAN neighbor from BGP neighbor table to control route advertisements according to SLA status.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWantempSystemVirtualWanLinkNeighbor() *schema.Resource {
	return &schema.Resource{
		Create: resourceWantempSystemVirtualWanLinkNeighborCreate,
		Read:   resourceWantempSystemVirtualWanLinkNeighborRead,
		Update: resourceWantempSystemVirtualWanLinkNeighborUpdate,
		Delete: resourceWantempSystemVirtualWanLinkNeighborDelete,

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
			"wanprof": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"health_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"member": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"role": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sla_id": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceWantempSystemVirtualWanLinkNeighborCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	wanprof := d.Get("wanprof").(string)
	paradict["wanprof"] = wanprof

	obj, err := getObjectWantempSystemVirtualWanLinkNeighbor(d)
	if err != nil {
		return fmt.Errorf("Error creating WantempSystemVirtualWanLinkNeighbor resource while getting object: %v", err)
	}

	_, err = c.CreateWantempSystemVirtualWanLinkNeighbor(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating WantempSystemVirtualWanLinkNeighbor resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "sla_id")))

	return resourceWantempSystemVirtualWanLinkNeighborRead(d, m)
}

func resourceWantempSystemVirtualWanLinkNeighborUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	wanprof := d.Get("wanprof").(string)
	paradict["wanprof"] = wanprof

	obj, err := getObjectWantempSystemVirtualWanLinkNeighbor(d)
	if err != nil {
		return fmt.Errorf("Error updating WantempSystemVirtualWanLinkNeighbor resource while getting object: %v", err)
	}

	_, err = c.UpdateWantempSystemVirtualWanLinkNeighbor(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WantempSystemVirtualWanLinkNeighbor resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "sla_id")))

	return resourceWantempSystemVirtualWanLinkNeighborRead(d, m)
}

func resourceWantempSystemVirtualWanLinkNeighborDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	wanprof := d.Get("wanprof").(string)
	paradict["wanprof"] = wanprof

	err = c.DeleteWantempSystemVirtualWanLinkNeighbor(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WantempSystemVirtualWanLinkNeighbor resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWantempSystemVirtualWanLinkNeighborRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	wanprof := d.Get("wanprof").(string)
	if wanprof == "" {
		wanprof = importOptionChecking(m.(*FortiClient).Cfg, "wanprof")
		if wanprof == "" {
			return fmt.Errorf("Parameter wanprof is missing")
		}
		if err = d.Set("wanprof", wanprof); err != nil {
			return fmt.Errorf("Error set params wanprof: %v", err)
		}
	}
	paradict["wanprof"] = wanprof

	o, err := c.ReadWantempSystemVirtualWanLinkNeighbor(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WantempSystemVirtualWanLinkNeighbor resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWantempSystemVirtualWanLinkNeighbor(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WantempSystemVirtualWanLinkNeighbor resource from API: %v", err)
	}
	return nil
}

func flattenWantempSystemVirtualWanLinkNeighborHealthCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkNeighborIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkNeighborMember(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkNeighborRole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkNeighborSlaId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWantempSystemVirtualWanLinkNeighbor(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("health_check", flattenWantempSystemVirtualWanLinkNeighborHealthCheck(o["health-check"], d, "health_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["health-check"], "WantempSystemVirtualWanLinkNeighbor-HealthCheck"); ok {
			if err = d.Set("health_check", vv); err != nil {
				return fmt.Errorf("Error reading health_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading health_check: %v", err)
		}
	}

	if err = d.Set("ip", flattenWantempSystemVirtualWanLinkNeighborIp(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "WantempSystemVirtualWanLinkNeighbor-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("member", flattenWantempSystemVirtualWanLinkNeighborMember(o["member"], d, "member")); err != nil {
		if vv, ok := fortiAPIPatch(o["member"], "WantempSystemVirtualWanLinkNeighbor-Member"); ok {
			if err = d.Set("member", vv); err != nil {
				return fmt.Errorf("Error reading member: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading member: %v", err)
		}
	}

	if err = d.Set("role", flattenWantempSystemVirtualWanLinkNeighborRole(o["role"], d, "role")); err != nil {
		if vv, ok := fortiAPIPatch(o["role"], "WantempSystemVirtualWanLinkNeighbor-Role"); ok {
			if err = d.Set("role", vv); err != nil {
				return fmt.Errorf("Error reading role: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading role: %v", err)
		}
	}

	if err = d.Set("sla_id", flattenWantempSystemVirtualWanLinkNeighborSlaId(o["sla-id"], d, "sla_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["sla-id"], "WantempSystemVirtualWanLinkNeighbor-SlaId"); ok {
			if err = d.Set("sla_id", vv); err != nil {
				return fmt.Errorf("Error reading sla_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sla_id: %v", err)
		}
	}

	return nil
}

func flattenWantempSystemVirtualWanLinkNeighborFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWantempSystemVirtualWanLinkNeighborHealthCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkNeighborIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkNeighborMember(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkNeighborRole(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkNeighborSlaId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWantempSystemVirtualWanLinkNeighbor(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("health_check"); ok || d.HasChange("health_check") {
		t, err := expandWantempSystemVirtualWanLinkNeighborHealthCheck(d, v, "health_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health-check"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandWantempSystemVirtualWanLinkNeighborIp(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("member"); ok || d.HasChange("member") {
		t, err := expandWantempSystemVirtualWanLinkNeighborMember(d, v, "member")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["member"] = t
		}
	}

	if v, ok := d.GetOk("role"); ok || d.HasChange("role") {
		t, err := expandWantempSystemVirtualWanLinkNeighborRole(d, v, "role")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["role"] = t
		}
	}

	if v, ok := d.GetOk("sla_id"); ok || d.HasChange("sla_id") {
		t, err := expandWantempSystemVirtualWanLinkNeighborSlaId(d, v, "sla_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sla-id"] = t
		}
	}

	return &obj, nil
}
