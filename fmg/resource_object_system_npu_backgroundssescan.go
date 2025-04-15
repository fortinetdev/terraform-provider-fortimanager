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
			"scan_stale": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"scan_vt": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"stats_qual_access": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"stats_qual_duration": &schema.Schema{
				Type:     schema.TypeInt,
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
			"udp_qual_access": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"udp_qual_duration": &schema.Schema{
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

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	obj, err := getObjectObjectSystemNpuBackgroundSseScan(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectSystemNpuBackgroundSseScan resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectSystemNpuBackgroundSseScan(obj, mkey, paradict, wsParams)
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

	paradict := make(map[string]string)
	wsParams := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	wsParams["adom"] = adomv

	err = c.DeleteObjectSystemNpuBackgroundSseScan(mkey, paradict, wsParams)
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

	paradict := make(map[string]string)
	cfg := m.(*FortiClient).Cfg
	adomv, err := adomChecking(cfg, d)
	if err != nil {
		return fmt.Errorf("Error adom configuration: %v", err)
	}
	paradict["adom"] = adomv

	o, err := c.ReadObjectSystemNpuBackgroundSseScan(mkey, paradict)
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

func flattenObjectSystemNpuBackgroundSseScanScan2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuBackgroundSseScanScanStale2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuBackgroundSseScanScanVt2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuBackgroundSseScanStatsQualAccess2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuBackgroundSseScanStatsQualDuration2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuBackgroundSseScanStatsUpdateInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuBackgroundSseScanUdpKeepaliveInterval2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuBackgroundSseScanUdpQualAccess2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectSystemNpuBackgroundSseScanUdpQualDuration2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectSystemNpuBackgroundSseScan(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("scan", flattenObjectSystemNpuBackgroundSseScanScan2edl(o["scan"], d, "scan")); err != nil {
		if vv, ok := fortiAPIPatch(o["scan"], "ObjectSystemNpuBackgroundSseScan-Scan"); ok {
			if err = d.Set("scan", vv); err != nil {
				return fmt.Errorf("Error reading scan: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scan: %v", err)
		}
	}

	if err = d.Set("scan_stale", flattenObjectSystemNpuBackgroundSseScanScanStale2edl(o["scan-stale"], d, "scan_stale")); err != nil {
		if vv, ok := fortiAPIPatch(o["scan-stale"], "ObjectSystemNpuBackgroundSseScan-ScanStale"); ok {
			if err = d.Set("scan_stale", vv); err != nil {
				return fmt.Errorf("Error reading scan_stale: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scan_stale: %v", err)
		}
	}

	if err = d.Set("scan_vt", flattenObjectSystemNpuBackgroundSseScanScanVt2edl(o["scan-vt"], d, "scan_vt")); err != nil {
		if vv, ok := fortiAPIPatch(o["scan-vt"], "ObjectSystemNpuBackgroundSseScan-ScanVt"); ok {
			if err = d.Set("scan_vt", vv); err != nil {
				return fmt.Errorf("Error reading scan_vt: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading scan_vt: %v", err)
		}
	}

	if err = d.Set("stats_qual_access", flattenObjectSystemNpuBackgroundSseScanStatsQualAccess2edl(o["stats-qual-access"], d, "stats_qual_access")); err != nil {
		if vv, ok := fortiAPIPatch(o["stats-qual-access"], "ObjectSystemNpuBackgroundSseScan-StatsQualAccess"); ok {
			if err = d.Set("stats_qual_access", vv); err != nil {
				return fmt.Errorf("Error reading stats_qual_access: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading stats_qual_access: %v", err)
		}
	}

	if err = d.Set("stats_qual_duration", flattenObjectSystemNpuBackgroundSseScanStatsQualDuration2edl(o["stats-qual-duration"], d, "stats_qual_duration")); err != nil {
		if vv, ok := fortiAPIPatch(o["stats-qual-duration"], "ObjectSystemNpuBackgroundSseScan-StatsQualDuration"); ok {
			if err = d.Set("stats_qual_duration", vv); err != nil {
				return fmt.Errorf("Error reading stats_qual_duration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading stats_qual_duration: %v", err)
		}
	}

	if err = d.Set("stats_update_interval", flattenObjectSystemNpuBackgroundSseScanStatsUpdateInterval2edl(o["stats-update-interval"], d, "stats_update_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["stats-update-interval"], "ObjectSystemNpuBackgroundSseScan-StatsUpdateInterval"); ok {
			if err = d.Set("stats_update_interval", vv); err != nil {
				return fmt.Errorf("Error reading stats_update_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading stats_update_interval: %v", err)
		}
	}

	if err = d.Set("udp_keepalive_interval", flattenObjectSystemNpuBackgroundSseScanUdpKeepaliveInterval2edl(o["udp-keepalive-interval"], d, "udp_keepalive_interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["udp-keepalive-interval"], "ObjectSystemNpuBackgroundSseScan-UdpKeepaliveInterval"); ok {
			if err = d.Set("udp_keepalive_interval", vv); err != nil {
				return fmt.Errorf("Error reading udp_keepalive_interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading udp_keepalive_interval: %v", err)
		}
	}

	if err = d.Set("udp_qual_access", flattenObjectSystemNpuBackgroundSseScanUdpQualAccess2edl(o["udp-qual-access"], d, "udp_qual_access")); err != nil {
		if vv, ok := fortiAPIPatch(o["udp-qual-access"], "ObjectSystemNpuBackgroundSseScan-UdpQualAccess"); ok {
			if err = d.Set("udp_qual_access", vv); err != nil {
				return fmt.Errorf("Error reading udp_qual_access: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading udp_qual_access: %v", err)
		}
	}

	if err = d.Set("udp_qual_duration", flattenObjectSystemNpuBackgroundSseScanUdpQualDuration2edl(o["udp-qual-duration"], d, "udp_qual_duration")); err != nil {
		if vv, ok := fortiAPIPatch(o["udp-qual-duration"], "ObjectSystemNpuBackgroundSseScan-UdpQualDuration"); ok {
			if err = d.Set("udp_qual_duration", vv); err != nil {
				return fmt.Errorf("Error reading udp_qual_duration: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading udp_qual_duration: %v", err)
		}
	}

	return nil
}

func flattenObjectSystemNpuBackgroundSseScanFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectSystemNpuBackgroundSseScanScan2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuBackgroundSseScanScanStale2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuBackgroundSseScanScanVt2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuBackgroundSseScanStatsQualAccess2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuBackgroundSseScanStatsQualDuration2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuBackgroundSseScanStatsUpdateInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuBackgroundSseScanUdpKeepaliveInterval2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuBackgroundSseScanUdpQualAccess2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectSystemNpuBackgroundSseScanUdpQualDuration2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectSystemNpuBackgroundSseScan(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("scan"); ok || d.HasChange("scan") {
		t, err := expandObjectSystemNpuBackgroundSseScanScan2edl(d, v, "scan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan"] = t
		}
	}

	if v, ok := d.GetOk("scan_stale"); ok || d.HasChange("scan_stale") {
		t, err := expandObjectSystemNpuBackgroundSseScanScanStale2edl(d, v, "scan_stale")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-stale"] = t
		}
	}

	if v, ok := d.GetOk("scan_vt"); ok || d.HasChange("scan_vt") {
		t, err := expandObjectSystemNpuBackgroundSseScanScanVt2edl(d, v, "scan_vt")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-vt"] = t
		}
	}

	if v, ok := d.GetOk("stats_qual_access"); ok || d.HasChange("stats_qual_access") {
		t, err := expandObjectSystemNpuBackgroundSseScanStatsQualAccess2edl(d, v, "stats_qual_access")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stats-qual-access"] = t
		}
	}

	if v, ok := d.GetOk("stats_qual_duration"); ok || d.HasChange("stats_qual_duration") {
		t, err := expandObjectSystemNpuBackgroundSseScanStatsQualDuration2edl(d, v, "stats_qual_duration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stats-qual-duration"] = t
		}
	}

	if v, ok := d.GetOk("stats_update_interval"); ok || d.HasChange("stats_update_interval") {
		t, err := expandObjectSystemNpuBackgroundSseScanStatsUpdateInterval2edl(d, v, "stats_update_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stats-update-interval"] = t
		}
	}

	if v, ok := d.GetOk("udp_keepalive_interval"); ok || d.HasChange("udp_keepalive_interval") {
		t, err := expandObjectSystemNpuBackgroundSseScanUdpKeepaliveInterval2edl(d, v, "udp_keepalive_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["udp-keepalive-interval"] = t
		}
	}

	if v, ok := d.GetOk("udp_qual_access"); ok || d.HasChange("udp_qual_access") {
		t, err := expandObjectSystemNpuBackgroundSseScanUdpQualAccess2edl(d, v, "udp_qual_access")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["udp-qual-access"] = t
		}
	}

	if v, ok := d.GetOk("udp_qual_duration"); ok || d.HasChange("udp_qual_duration") {
		t, err := expandObjectSystemNpuBackgroundSseScanUdpQualDuration2edl(d, v, "udp_qual_duration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["udp-qual-duration"] = t
		}
	}

	return &obj, nil
}
