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

func resourceWantempSystemSdwanHealthCheckSla() *schema.Resource {
	return &schema.Resource{
		Create: resourceWantempSystemSdwanHealthCheckSlaCreate,
		Read:   resourceWantempSystemSdwanHealthCheckSlaRead,
		Update: resourceWantempSystemSdwanHealthCheckSlaUpdate,
		Delete: resourceWantempSystemSdwanHealthCheckSlaDelete,

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
				Computed: true,
			},
			"latency_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"link_cost_factor": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"mos_threshold": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"packetloss_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"priority_in_sla": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"priority_out_sla": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceWantempSystemSdwanHealthCheckSlaCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
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

	obj, err := getObjectWantempSystemSdwanHealthCheckSla(d)
	if err != nil {
		return fmt.Errorf("Error creating WantempSystemSdwanHealthCheckSla resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateWantempSystemSdwanHealthCheckSla(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating WantempSystemSdwanHealthCheckSla resource: %v", err)
	}

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceWantempSystemSdwanHealthCheckSlaRead(d, m)
}

func resourceWantempSystemSdwanHealthCheckSlaUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
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

	obj, err := getObjectWantempSystemSdwanHealthCheckSla(d)
	if err != nil {
		return fmt.Errorf("Error updating WantempSystemSdwanHealthCheckSla resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateWantempSystemSdwanHealthCheckSla(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating WantempSystemSdwanHealthCheckSla resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(strconv.Itoa(getIntKey(d, "fosid")))

	return resourceWantempSystemSdwanHealthCheckSlaRead(d, m)
}

func resourceWantempSystemSdwanHealthCheckSlaDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
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

	wsParams["adom"] = adomv

	err = c.DeleteWantempSystemSdwanHealthCheckSla(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting WantempSystemSdwanHealthCheckSla resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWantempSystemSdwanHealthCheckSlaRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWantempSystemSdwanHealthCheckSla(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading WantempSystemSdwanHealthCheckSla resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWantempSystemSdwanHealthCheckSla(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WantempSystemSdwanHealthCheckSla resource from API: %v", err)
	}
	return nil
}

func flattenWantempSystemSdwanHealthCheckSlaId3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckSlaJitterThreshold3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckSlaLatencyThreshold3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckSlaLinkCostFactor3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return flattenStringList(v)
}

func flattenWantempSystemSdwanHealthCheckSlaMosThreshold3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckSlaPacketlossThreshold3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckSlaPriorityInSla3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWantempSystemSdwanHealthCheckSlaPriorityOutSla3rdl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWantempSystemSdwanHealthCheckSla(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("fosid", flattenWantempSystemSdwanHealthCheckSlaId3rdl(o["id"], d, "fosid")); err != nil {
		if vv, ok := fortiAPIPatch(o["id"], "WantempSystemSdwanHealthCheckSla-Id"); ok {
			if err = d.Set("fosid", vv); err != nil {
				return fmt.Errorf("Error reading fosid: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("jitter_threshold", flattenWantempSystemSdwanHealthCheckSlaJitterThreshold3rdl(o["jitter-threshold"], d, "jitter_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["jitter-threshold"], "WantempSystemSdwanHealthCheckSla-JitterThreshold"); ok {
			if err = d.Set("jitter_threshold", vv); err != nil {
				return fmt.Errorf("Error reading jitter_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading jitter_threshold: %v", err)
		}
	}

	if err = d.Set("latency_threshold", flattenWantempSystemSdwanHealthCheckSlaLatencyThreshold3rdl(o["latency-threshold"], d, "latency_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["latency-threshold"], "WantempSystemSdwanHealthCheckSla-LatencyThreshold"); ok {
			if err = d.Set("latency_threshold", vv); err != nil {
				return fmt.Errorf("Error reading latency_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading latency_threshold: %v", err)
		}
	}

	if err = d.Set("link_cost_factor", flattenWantempSystemSdwanHealthCheckSlaLinkCostFactor3rdl(o["link-cost-factor"], d, "link_cost_factor")); err != nil {
		if vv, ok := fortiAPIPatch(o["link-cost-factor"], "WantempSystemSdwanHealthCheckSla-LinkCostFactor"); ok {
			if err = d.Set("link_cost_factor", vv); err != nil {
				return fmt.Errorf("Error reading link_cost_factor: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading link_cost_factor: %v", err)
		}
	}

	if err = d.Set("mos_threshold", flattenWantempSystemSdwanHealthCheckSlaMosThreshold3rdl(o["mos-threshold"], d, "mos_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["mos-threshold"], "WantempSystemSdwanHealthCheckSla-MosThreshold"); ok {
			if err = d.Set("mos_threshold", vv); err != nil {
				return fmt.Errorf("Error reading mos_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading mos_threshold: %v", err)
		}
	}

	if err = d.Set("packetloss_threshold", flattenWantempSystemSdwanHealthCheckSlaPacketlossThreshold3rdl(o["packetloss-threshold"], d, "packetloss_threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["packetloss-threshold"], "WantempSystemSdwanHealthCheckSla-PacketlossThreshold"); ok {
			if err = d.Set("packetloss_threshold", vv); err != nil {
				return fmt.Errorf("Error reading packetloss_threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading packetloss_threshold: %v", err)
		}
	}

	if err = d.Set("priority_in_sla", flattenWantempSystemSdwanHealthCheckSlaPriorityInSla3rdl(o["priority-in-sla"], d, "priority_in_sla")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority-in-sla"], "WantempSystemSdwanHealthCheckSla-PriorityInSla"); ok {
			if err = d.Set("priority_in_sla", vv); err != nil {
				return fmt.Errorf("Error reading priority_in_sla: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority_in_sla: %v", err)
		}
	}

	if err = d.Set("priority_out_sla", flattenWantempSystemSdwanHealthCheckSlaPriorityOutSla3rdl(o["priority-out-sla"], d, "priority_out_sla")); err != nil {
		if vv, ok := fortiAPIPatch(o["priority-out-sla"], "WantempSystemSdwanHealthCheckSla-PriorityOutSla"); ok {
			if err = d.Set("priority_out_sla", vv); err != nil {
				return fmt.Errorf("Error reading priority_out_sla: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading priority_out_sla: %v", err)
		}
	}

	return nil
}

func flattenWantempSystemSdwanHealthCheckSlaFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWantempSystemSdwanHealthCheckSlaId3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckSlaJitterThreshold3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckSlaLatencyThreshold3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckSlaLinkCostFactor3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return expandStringList(v.(*schema.Set).List()), nil
}

func expandWantempSystemSdwanHealthCheckSlaMosThreshold3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckSlaPacketlossThreshold3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckSlaPriorityInSla3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWantempSystemSdwanHealthCheckSlaPriorityOutSla3rdl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWantempSystemSdwanHealthCheckSla(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok || d.HasChange("fosid") {
		t, err := expandWantempSystemSdwanHealthCheckSlaId3rdl(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("jitter_threshold"); ok || d.HasChange("jitter_threshold") {
		t, err := expandWantempSystemSdwanHealthCheckSlaJitterThreshold3rdl(d, v, "jitter_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["jitter-threshold"] = t
		}
	}

	if v, ok := d.GetOk("latency_threshold"); ok || d.HasChange("latency_threshold") {
		t, err := expandWantempSystemSdwanHealthCheckSlaLatencyThreshold3rdl(d, v, "latency_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["latency-threshold"] = t
		}
	}

	if v, ok := d.GetOk("link_cost_factor"); ok || d.HasChange("link_cost_factor") {
		t, err := expandWantempSystemSdwanHealthCheckSlaLinkCostFactor3rdl(d, v, "link_cost_factor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["link-cost-factor"] = t
		}
	}

	if v, ok := d.GetOk("mos_threshold"); ok || d.HasChange("mos_threshold") {
		t, err := expandWantempSystemSdwanHealthCheckSlaMosThreshold3rdl(d, v, "mos_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mos-threshold"] = t
		}
	}

	if v, ok := d.GetOk("packetloss_threshold"); ok || d.HasChange("packetloss_threshold") {
		t, err := expandWantempSystemSdwanHealthCheckSlaPacketlossThreshold3rdl(d, v, "packetloss_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["packetloss-threshold"] = t
		}
	}

	if v, ok := d.GetOk("priority_in_sla"); ok || d.HasChange("priority_in_sla") {
		t, err := expandWantempSystemSdwanHealthCheckSlaPriorityInSla3rdl(d, v, "priority_in_sla")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-in-sla"] = t
		}
	}

	if v, ok := d.GetOk("priority_out_sla"); ok || d.HasChange("priority_out_sla") {
		t, err := expandWantempSystemSdwanHealthCheckSlaPriorityOutSla3rdl(d, v, "priority_out_sla")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-out-sla"] = t
		}
	}

	return &obj, nil
}
