// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure driver background scan for SSE.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectSystemNpuBackgroundSseScan() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectSystemNpuBackgroundSseScanUpdate,
		Read:   resourceObjectSystemNpuBackgroundSseScanRead,
		Update: resourceObjectSystemNpuBackgroundSseScanUpdate,
		Delete: resourceObjectSystemNpuBackgroundSseScanDelete,

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
			"scan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"stats_update_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"udp_keepalive_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceObjectSystemNpuBackgroundSseScanUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	obj, err := getObjectObjectSystemNpuBackgroundSseScan(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuBackgroundSseScan resource while getting object: %v", err)
	}

	_, err = c.UpdateObjectSystemNpuBackgroundSseScan(obj, adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuBackgroundSseScan resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId("ObjectSystemNpuBackgroundSseScan")

	return resourceObjectSystemNpuBackgroundSseScanRead(d, m)
}

func resourceObjectSystemNpuBackgroundSseScanDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	err = c.DeleteObjectSystemNpuBackgroundSseScan(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectSystemNpuBackgroundSseScan resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectSystemNpuBackgroundSseScanRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}

	o, err := c.ReadObjectSystemNpuBackgroundSseScan(adomv, mkey, nil)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuBackgroundSseScan resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectSystemNpuBackgroundSseScan(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectSystemNpuBackgroundSseScan resource from API: %v", err)
	}
	return nil
}

func flattenObjectSystemNpuBackgroundSseScanScan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuBackgroundSseScanStatsUpdateInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuBackgroundSseScanUdpKeepaliveInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemNpuBackgroundSseScan(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("scan", flattenObjectSystemNpuBackgroundSseScanScan(o["scan"], d, "scan")); err != nil {
		if vv, ok := fortiAPIPatch(o["scan"], "ObjectSystemNpuBackgroundSseScan-Scan"); ok {
			if err = d.Set("scan", vv); err != nil {
				return fmt.Errorf("Error reading scan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scan: %v", err)
		}
	}

	if err = d.Set("stats_update_interval", flattenObjectSystemNpuBackgroundSseScanStatsUpdateInterval(o["stats-update-interval"], d, "stats_update_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["stats-update-interval"], "ObjectSystemNpuBackgroundSseScan-StatsUpdateInterval"); ok {
			if err = d.Set("stats_update_interval", vv); err != nil {
				return fmt.Errorf("Error reading stats_update_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading stats_update_interval: %v", err)
		}
	}

	if err = d.Set("udp_keepalive_interval", flattenObjectSystemNpuBackgroundSseScanUdpKeepaliveInterval(o["udp-keepalive-interval"], d, "udp_keepalive_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["udp-keepalive-interval"], "ObjectSystemNpuBackgroundSseScan-UdpKeepaliveInterval"); ok {
			if err = d.Set("udp_keepalive_interval", vv); err != nil {
				return fmt.Errorf("Error reading udp_keepalive_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading udp_keepalive_interval: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemNpuBackgroundSseScanFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemNpuBackgroundSseScanScan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuBackgroundSseScanStatsUpdateInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuBackgroundSseScanUdpKeepaliveInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemNpuBackgroundSseScan(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("scan"); ok || d.HasChange("scan") {
		t, err := expandObjectSystemNpuBackgroundSseScanScan(d, v, "scan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan"] = t
		}
	}

	if v, ok := d.GetOk("stats_update_interval"); ok || d.HasChange("stats_update_interval") {
		t, err := expandObjectSystemNpuBackgroundSseScanStatsUpdateInterval(d, v, "stats_update_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stats-update-interval"] = t
		}
	}

	if v, ok := d.GetOk("udp_keepalive_interval"); ok || d.HasChange("udp_keepalive_interval") {
		t, err := expandObjectSystemNpuBackgroundSseScanUdpKeepaliveInterval(d, v, "udp_keepalive_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["udp-keepalive-interval"] = t
		}
	}

	return &obj, nil
}
