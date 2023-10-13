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

func resourceWantempSystemSdwanNeighbor() *schema.Resource {
	return &schema.Resource{
		Create: resourceWantempSystemSdwanNeighborCreate,
		Read:   resourceWantempSystemSdwanNeighborRead,
		Update: resourceWantempSystemSdwanNeighborUpdate,
		Delete: resourceWantempSystemSdwanNeighborDelete,

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
				ForceNew: true,
				Required: true,
			},
			"member": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"minimum_sla_meet_members": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"role": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"sla_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceWantempSystemSdwanNeighborCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWantempSystemSdwanNeighbor(d)
	if err != nil {
		return fmt.Errorf("Error creating WantempSystemSdwanNeighbor resource while getting object: %v", err)
	}

	_, err = c.CreateWantempSystemSdwanNeighbor(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating WantempSystemSdwanNeighbor resource: %v", err)
	}

	d.SetId(getStringKey(d, "ip"))

	return resourceWantempSystemSdwanNeighborRead(d, m)
}

func resourceWantempSystemSdwanNeighborUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWantempSystemSdwanNeighbor(d)
	if err != nil {
		return fmt.Errorf("Error updating WantempSystemSdwanNeighbor resource while getting object: %v", err)
	}

	_, err = c.UpdateWantempSystemSdwanNeighbor(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WantempSystemSdwanNeighbor resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "ip"))

	return resourceWantempSystemSdwanNeighborRead(d, m)
}

func resourceWantempSystemSdwanNeighborDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteWantempSystemSdwanNeighbor(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WantempSystemSdwanNeighbor resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWantempSystemSdwanNeighborRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWantempSystemSdwanNeighbor(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WantempSystemSdwanNeighbor resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWantempSystemSdwanNeighbor(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WantempSystemSdwanNeighbor resource from API: %v", err)
	}
	return nil
}

func flattenWantempSystemSdwanNeighborHealthCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanNeighborIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanNeighborMember(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanNeighborMinimumSlaMeetMembers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanNeighborMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanNeighborRole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanNeighborServiceId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanNeighborSlaId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWantempSystemSdwanNeighbor(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("health_check", flattenWantempSystemSdwanNeighborHealthCheck(o["health-check"], d, "health_check")); err != nil {
		if vv, ok := fortiAPIPatch(o["health-check"], "WantempSystemSdwanNeighbor-HealthCheck"); ok {
			if err = d.Set("health_check", vv); err != nil {
				return fmt.Errorf("Error reading health_check: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading health_check: %v", err)
		}
	}

	if err = d.Set("ip", flattenWantempSystemSdwanNeighborIp(o["ip"], d, "ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["ip"], "WantempSystemSdwanNeighbor-Ip"); ok {
			if err = d.Set("ip", vv); err != nil {
				return fmt.Errorf("Error reading ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("member", flattenWantempSystemSdwanNeighborMember(o["member"], d, "member")); err != nil {
		if vv, ok := fortiAPIPatch(o["member"], "WantempSystemSdwanNeighbor-Member"); ok {
			if err = d.Set("member", vv); err != nil {
				return fmt.Errorf("Error reading member: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading member: %v", err)
		}
	}

	if err = d.Set("minimum_sla_meet_members", flattenWantempSystemSdwanNeighborMinimumSlaMeetMembers(o["minimum-sla-meet-members"], d, "minimum_sla_meet_members")); err != nil {
		if vv, ok := fortiAPIPatch(o["minimum-sla-meet-members"], "WantempSystemSdwanNeighbor-MinimumSlaMeetMembers"); ok {
			if err = d.Set("minimum_sla_meet_members", vv); err != nil {
				return fmt.Errorf("Error reading minimum_sla_meet_members: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading minimum_sla_meet_members: %v", err)
		}
	}

	if err = d.Set("mode", flattenWantempSystemSdwanNeighborMode(o["mode"], d, "mode")); err != nil {
		if vv, ok := fortiAPIPatch(o["mode"], "WantempSystemSdwanNeighbor-Mode"); ok {
			if err = d.Set("mode", vv); err != nil {
				return fmt.Errorf("Error reading mode: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("role", flattenWantempSystemSdwanNeighborRole(o["role"], d, "role")); err != nil {
		if vv, ok := fortiAPIPatch(o["role"], "WantempSystemSdwanNeighbor-Role"); ok {
			if err = d.Set("role", vv); err != nil {
				return fmt.Errorf("Error reading role: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading role: %v", err)
		}
	}

	if err = d.Set("service_id", flattenWantempSystemSdwanNeighborServiceId(o["service-id"], d, "service_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["service-id"], "WantempSystemSdwanNeighbor-ServiceId"); ok {
			if err = d.Set("service_id", vv); err != nil {
				return fmt.Errorf("Error reading service_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading service_id: %v", err)
		}
	}

	if err = d.Set("sla_id", flattenWantempSystemSdwanNeighborSlaId(o["sla-id"], d, "sla_id")); err != nil {
		if vv, ok := fortiAPIPatch(o["sla-id"], "WantempSystemSdwanNeighbor-SlaId"); ok {
			if err = d.Set("sla_id", vv); err != nil {
				return fmt.Errorf("Error reading sla_id: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading sla_id: %v", err)
		}
	}

	return nil
}

func flattenWantempSystemSdwanNeighborFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWantempSystemSdwanNeighborHealthCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanNeighborIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanNeighborMember(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanNeighborMinimumSlaMeetMembers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanNeighborMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanNeighborRole(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanNeighborServiceId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanNeighborSlaId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWantempSystemSdwanNeighbor(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("health_check"); ok || d.HasChange("health_check") {
		t, err := expandWantempSystemSdwanNeighborHealthCheck(d, v, "health_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health-check"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok || d.HasChange("ip") {
		t, err := expandWantempSystemSdwanNeighborIp(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("member"); ok || d.HasChange("member") {
		t, err := expandWantempSystemSdwanNeighborMember(d, v, "member")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["member"] = t
		}
	}

	if v, ok := d.GetOk("minimum_sla_meet_members"); ok || d.HasChange("minimum_sla_meet_members") {
		t, err := expandWantempSystemSdwanNeighborMinimumSlaMeetMembers(d, v, "minimum_sla_meet_members")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["minimum-sla-meet-members"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok || d.HasChange("mode") {
		t, err := expandWantempSystemSdwanNeighborMode(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("role"); ok || d.HasChange("role") {
		t, err := expandWantempSystemSdwanNeighborRole(d, v, "role")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["role"] = t
		}
	}

	if v, ok := d.GetOk("service_id"); ok || d.HasChange("service_id") {
		t, err := expandWantempSystemSdwanNeighborServiceId(d, v, "service_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-id"] = t
		}
	}

	if v, ok := d.GetOk("sla_id"); ok || d.HasChange("sla_id") {
		t, err := expandWantempSystemSdwanNeighborSlaId(d, v, "sla_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sla-id"] = t
		}
	}

	return &obj, nil
}
