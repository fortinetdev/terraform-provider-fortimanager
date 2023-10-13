// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Service level agreement (SLA).

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWantempSystemVirtualWanLinkHealthCheckSla() *schema.Resource {
	return &schema.Resource{
		Create: resourceWantempSystemVirtualWanLinkHealthCheckSlaCreate,
		Read:   resourceWantempSystemVirtualWanLinkHealthCheckSlaRead,
		Update: resourceWantempSystemVirtualWanLinkHealthCheckSlaUpdate,
		Delete: resourceWantempSystemVirtualWanLinkHealthCheckSlaDelete,

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
				Required: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
			},
			"jitter_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"latency_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"link_cost_factor": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"packetloss_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceWantempSystemVirtualWanLinkHealthCheckSlaCreate(d *schema.ResourceData, m interface{}) error {
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
	health_check := d.Get("health_check").(string)
	paradict["wanprof"] = wanprof
	paradict["health_check"] = health_check

	obj, err := getObjectWantempSystemVirtualWanLinkHealthCheckSla(d)
	if err != nil {
		return fmt.Errorf("Error creating WantempSystemVirtualWanLinkHealthCheckSla resource while getting object: %v", err)
	}

	_, err = c.CreateWantempSystemVirtualWanLinkHealthCheckSla(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating WantempSystemVirtualWanLinkHealthCheckSla resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceWantempSystemVirtualWanLinkHealthCheckSlaRead(d, m)
}

func resourceWantempSystemVirtualWanLinkHealthCheckSlaUpdate(d *schema.ResourceData, m interface{}) error {
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
	health_check := d.Get("health_check").(string)
	paradict["wanprof"] = wanprof
	paradict["health_check"] = health_check

	obj, err := getObjectWantempSystemVirtualWanLinkHealthCheckSla(d)
	if err != nil {
		return fmt.Errorf("Error updating WantempSystemVirtualWanLinkHealthCheckSla resource while getting object: %v", err)
	}

	_, err = c.UpdateWantempSystemVirtualWanLinkHealthCheckSla(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating WantempSystemVirtualWanLinkHealthCheckSla resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceWantempSystemVirtualWanLinkHealthCheckSlaRead(d, m)
}

func resourceWantempSystemVirtualWanLinkHealthCheckSlaDelete(d *schema.ResourceData, m interface{}) error {
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
	health_check := d.Get("health_check").(string)
	paradict["wanprof"] = wanprof
	paradict["health_check"] = health_check

	err = c.DeleteWantempSystemVirtualWanLinkHealthCheckSla(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting WantempSystemVirtualWanLinkHealthCheckSla resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWantempSystemVirtualWanLinkHealthCheckSlaRead(d *schema.ResourceData, m interface{}) error {
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
	health_check := d.Get("health_check").(string)
	if wanprof == "" {
		wanprof = importOptionChecking(m.(*FortiClient).Cfg, "wanprof")
		if wanprof == "" {
			return fmt.Errorf("Parameter wanprof is missing")
		}
		if err = d.Set("wanprof", wanprof); err != nil {
			return fmt.Errorf("Error set params wanprof: %v", err)
		}
	}
	if health_check == "" {
		health_check = importOptionChecking(m.(*FortiClient).Cfg, "health_check")
		if health_check == "" {
			return fmt.Errorf("Parameter health_check is missing")
		}
		if err = d.Set("health_check", health_check); err != nil {
			return fmt.Errorf("Error set params health_check: %v", err)
		}
	}
	paradict["wanprof"] = wanprof
	paradict["health_check"] = health_check

	o, err := c.ReadWantempSystemVirtualWanLinkHealthCheckSla(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WantempSystemVirtualWanLinkHealthCheckSla resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWantempSystemVirtualWanLinkHealthCheckSla(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WantempSystemVirtualWanLinkHealthCheckSla resource from API: %v", err)
	}
	return nil
}

func flattenWantempSystemVirtualWanLinkHealthCheckSlaId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckSlaJitterThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckSlaLatencyThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemVirtualWanLinkHealthCheckSlaLinkCostFactor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWantempSystemVirtualWanLinkHealthCheckSlaPacketlossThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWantempSystemVirtualWanLinkHealthCheckSla(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("fosid", flattenWantempSystemVirtualWanLinkHealthCheckSlaId(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "WantempSystemVirtualWanLinkHealthCheckSla-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("jitter_threshold", flattenWantempSystemVirtualWanLinkHealthCheckSlaJitterThreshold(o["jitter-threshold"], d, "jitter_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["jitter-threshold"], "WantempSystemVirtualWanLinkHealthCheckSla-JitterThreshold"); ok {
			if err = d.Set("jitter_threshold", vv); err != nil {
				return fmt.Errorf("Error reading jitter_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading jitter_threshold: %v", err)
		}
	}

	if err = d.Set("latency_threshold", flattenWantempSystemVirtualWanLinkHealthCheckSlaLatencyThreshold(o["latency-threshold"], d, "latency_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["latency-threshold"], "WantempSystemVirtualWanLinkHealthCheckSla-LatencyThreshold"); ok {
			if err = d.Set("latency_threshold", vv); err != nil {
				return fmt.Errorf("Error reading latency_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading latency_threshold: %v", err)
		}
	}

	if err = d.Set("link_cost_factor", flattenWantempSystemVirtualWanLinkHealthCheckSlaLinkCostFactor(o["link-cost-factor"], d, "link_cost_factor")); err != nil {
		if vv, ok := fortiAPIPatch(o["link-cost-factor"], "WantempSystemVirtualWanLinkHealthCheckSla-LinkCostFactor"); ok {
			if err = d.Set("link_cost_factor", vv); err != nil {
				return fmt.Errorf("Error reading link_cost_factor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading link_cost_factor: %v", err)
		}
	}

	if err = d.Set("packetloss_threshold", flattenWantempSystemVirtualWanLinkHealthCheckSlaPacketlossThreshold(o["packetloss-threshold"], d, "packetloss_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["packetloss-threshold"], "WantempSystemVirtualWanLinkHealthCheckSla-PacketlossThreshold"); ok {
			if err = d.Set("packetloss_threshold", vv); err != nil {
				return fmt.Errorf("Error reading packetloss_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading packetloss_threshold: %v", err)
		}
	}

	return nil
}

func flattenWantempSystemVirtualWanLinkHealthCheckSlaFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWantempSystemVirtualWanLinkHealthCheckSlaId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckSlaJitterThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckSlaLatencyThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemVirtualWanLinkHealthCheckSlaLinkCostFactor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWantempSystemVirtualWanLinkHealthCheckSlaPacketlossThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWantempSystemVirtualWanLinkHealthCheckSla(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandWantempSystemVirtualWanLinkHealthCheckSlaId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("jitter_threshold"); ok || d.HasChange("jitter_threshold") {
		t, err := expandWantempSystemVirtualWanLinkHealthCheckSlaJitterThreshold(d, v, "jitter_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["jitter-threshold"] = t
		}
	}

	if v, ok := d.GetOk("latency_threshold"); ok || d.HasChange("latency_threshold") {
		t, err := expandWantempSystemVirtualWanLinkHealthCheckSlaLatencyThreshold(d, v, "latency_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["latency-threshold"] = t
		}
	}

	if v, ok := d.GetOk("link_cost_factor"); ok || d.HasChange("link_cost_factor") {
		t, err := expandWantempSystemVirtualWanLinkHealthCheckSlaLinkCostFactor(d, v, "link_cost_factor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["link-cost-factor"] = t
		}
	}

	if v, ok := d.GetOk("packetloss_threshold"); ok || d.HasChange("packetloss_threshold") {
		t, err := expandWantempSystemVirtualWanLinkHealthCheckSlaPacketlossThreshold(d, v, "packetloss_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["packetloss-threshold"] = t
		}
	}

	return &obj, nil
}
