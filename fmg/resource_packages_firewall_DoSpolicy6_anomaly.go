// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Anomaly name.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourcePackagesFirewallDosPolicy6Anomaly() *schema.Resource {
	return &schema.Resource{
		Create: resourcePackagesFirewallDosPolicy6AnomalyCreate,
		Read:   resourcePackagesFirewallDosPolicy6AnomalyRead,
		Update: resourcePackagesFirewallDosPolicy6AnomalyUpdate,
		Delete: resourcePackagesFirewallDosPolicy6AnomalyDelete,

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
			"pkg_folder_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"pkg": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"dos_policy6": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"quarantine": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"quarantine_expiry": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"quarantine_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"synproxy_tcp_mss": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"synproxy_tcp_sack": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"synproxy_tcp_timestamp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"synproxy_tcp_window": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"synproxy_tcp_windowscale": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"synproxy_tos": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"synproxy_ttl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"thresholddefault": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourcePackagesFirewallDosPolicy6AnomalyCreate(d *schema.ResourceData, m interface{}) error {
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

	pkg_folder_path := d.Get("pkg_folder_path").(string)
	pkg := d.Get("pkg").(string)
	dos_policy6 := d.Get("dos_policy6").(string)
	paradict["pkg_folder_path"] = formatPath(pkg_folder_path)
	paradict["pkg"] = pkg
	paradict["dos_policy6"] = dos_policy6

	obj, err := getObjectPackagesFirewallDosPolicy6Anomaly(d)
	if err != nil {
		return fmt.Errorf("Error creating PackagesFirewallDosPolicy6Anomaly resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreatePackagesFirewallDosPolicy6Anomaly(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating PackagesFirewallDosPolicy6Anomaly resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourcePackagesFirewallDosPolicy6AnomalyRead(d, m)
}

func resourcePackagesFirewallDosPolicy6AnomalyUpdate(d *schema.ResourceData, m interface{}) error {
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

	pkg_folder_path := d.Get("pkg_folder_path").(string)
	pkg := d.Get("pkg").(string)
	dos_policy6 := d.Get("dos_policy6").(string)
	paradict["pkg_folder_path"] = formatPath(pkg_folder_path)
	paradict["pkg"] = pkg
	paradict["dos_policy6"] = dos_policy6

	obj, err := getObjectPackagesFirewallDosPolicy6Anomaly(d)
	if err != nil {
		return fmt.Errorf("Error updating PackagesFirewallDosPolicy6Anomaly resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdatePackagesFirewallDosPolicy6Anomaly(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating PackagesFirewallDosPolicy6Anomaly resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourcePackagesFirewallDosPolicy6AnomalyRead(d, m)
}

func resourcePackagesFirewallDosPolicy6AnomalyDelete(d *schema.ResourceData, m interface{}) error {
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

	pkg_folder_path := d.Get("pkg_folder_path").(string)
	pkg := d.Get("pkg").(string)
	dos_policy6 := d.Get("dos_policy6").(string)
	paradict["pkg_folder_path"] = formatPath(pkg_folder_path)
	paradict["pkg"] = pkg
	paradict["dos_policy6"] = dos_policy6

	wsParams["adom"] = adomv

	err = c.DeletePackagesFirewallDosPolicy6Anomaly(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting PackagesFirewallDosPolicy6Anomaly resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourcePackagesFirewallDosPolicy6AnomalyRead(d *schema.ResourceData, m interface{}) error {
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

	pkg_folder_path := d.Get("pkg_folder_path").(string)
	pkg := d.Get("pkg").(string)
	dos_policy6 := d.Get("dos_policy6").(string)
	if pkg_folder_path == "" {
		pkg_folder_path = importOptionChecking(m.(*FortiClient).Cfg, "pkg_folder_path")
	}
	if pkg == "" {
		pkg = importOptionChecking(m.(*FortiClient).Cfg, "pkg")
		if pkg == "" {
			return fmt.Errorf("Parameter pkg is missing")
		}
		if err = d.Set("pkg", pkg); err != nil {
			return fmt.Errorf("Error set params pkg: %v", err)
		}
	}
	if dos_policy6 == "" {
		dos_policy6 = importOptionChecking(m.(*FortiClient).Cfg, "dos_policy6")
		if dos_policy6 == "" {
			return fmt.Errorf("Parameter dos_policy6 is missing")
		}
		if err = d.Set("dos_policy6", dos_policy6); err != nil {
			return fmt.Errorf("Error set params dos_policy6: %v", err)
		}
	}
	paradict["pkg_folder_path"] = formatPath(pkg_folder_path)
	paradict["pkg"] = pkg
	paradict["dos_policy6"] = dos_policy6

	o, err := c.ReadPackagesFirewallDosPolicy6Anomaly(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading PackagesFirewallDosPolicy6Anomaly resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectPackagesFirewallDosPolicy6Anomaly(d, o)
	if err != nil {
		return fmt.Errorf("Error reading PackagesFirewallDosPolicy6Anomaly resource from API: %v", err)
	}
	return nil
}

func flattenPackagesFirewallDosPolicy6AnomalyAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicy6AnomalyLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicy6AnomalyName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicy6AnomalyQuarantine2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicy6AnomalyQuarantineExpiry2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicy6AnomalyQuarantineLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicy6AnomalyStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicy6AnomalySynproxyTcpMss2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicy6AnomalySynproxyTcpSack2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicy6AnomalySynproxyTcpTimestamp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicy6AnomalySynproxyTcpWindow2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicy6AnomalySynproxyTcpWindowscale2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicy6AnomalySynproxyTos2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicy6AnomalySynproxyTtl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicy6AnomalyThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicy6AnomalyThresholdDefault2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectPackagesFirewallDosPolicy6Anomaly(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("action", flattenPackagesFirewallDosPolicy6AnomalyAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "PackagesFirewallDosPolicy6Anomaly-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("log", flattenPackagesFirewallDosPolicy6AnomalyLog2edl(o["log"], d, "log")); err != nil {
		if vv, ok := fortiAPIPatch(o["log"], "PackagesFirewallDosPolicy6Anomaly-Log"); ok {
			if err = d.Set("log", vv); err != nil {
				return fmt.Errorf("Error reading log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	if err = d.Set("name", flattenPackagesFirewallDosPolicy6AnomalyName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "PackagesFirewallDosPolicy6Anomaly-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("quarantine", flattenPackagesFirewallDosPolicy6AnomalyQuarantine2edl(o["quarantine"], d, "quarantine")); err != nil {
		if vv, ok := fortiAPIPatch(o["quarantine"], "PackagesFirewallDosPolicy6Anomaly-Quarantine"); ok {
			if err = d.Set("quarantine", vv); err != nil {
				return fmt.Errorf("Error reading quarantine: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading quarantine: %v", err)
		}
	}

	if err = d.Set("quarantine_expiry", flattenPackagesFirewallDosPolicy6AnomalyQuarantineExpiry2edl(o["quarantine-expiry"], d, "quarantine_expiry")); err != nil {
		if vv, ok := fortiAPIPatch(o["quarantine-expiry"], "PackagesFirewallDosPolicy6Anomaly-QuarantineExpiry"); ok {
			if err = d.Set("quarantine_expiry", vv); err != nil {
				return fmt.Errorf("Error reading quarantine_expiry: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading quarantine_expiry: %v", err)
		}
	}

	if err = d.Set("quarantine_log", flattenPackagesFirewallDosPolicy6AnomalyQuarantineLog2edl(o["quarantine-log"], d, "quarantine_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["quarantine-log"], "PackagesFirewallDosPolicy6Anomaly-QuarantineLog"); ok {
			if err = d.Set("quarantine_log", vv); err != nil {
				return fmt.Errorf("Error reading quarantine_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading quarantine_log: %v", err)
		}
	}

	if err = d.Set("status", flattenPackagesFirewallDosPolicy6AnomalyStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "PackagesFirewallDosPolicy6Anomaly-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("synproxy_tcp_mss", flattenPackagesFirewallDosPolicy6AnomalySynproxyTcpMss2edl(o["synproxy-tcp-mss"], d, "synproxy_tcp_mss")); err != nil {
		if vv, ok := fortiAPIPatch(o["synproxy-tcp-mss"], "PackagesFirewallDosPolicy6Anomaly-SynproxyTcpMss"); ok {
			if err = d.Set("synproxy_tcp_mss", vv); err != nil {
				return fmt.Errorf("Error reading synproxy_tcp_mss: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading synproxy_tcp_mss: %v", err)
		}
	}

	if err = d.Set("synproxy_tcp_sack", flattenPackagesFirewallDosPolicy6AnomalySynproxyTcpSack2edl(o["synproxy-tcp-sack"], d, "synproxy_tcp_sack")); err != nil {
		if vv, ok := fortiAPIPatch(o["synproxy-tcp-sack"], "PackagesFirewallDosPolicy6Anomaly-SynproxyTcpSack"); ok {
			if err = d.Set("synproxy_tcp_sack", vv); err != nil {
				return fmt.Errorf("Error reading synproxy_tcp_sack: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading synproxy_tcp_sack: %v", err)
		}
	}

	if err = d.Set("synproxy_tcp_timestamp", flattenPackagesFirewallDosPolicy6AnomalySynproxyTcpTimestamp2edl(o["synproxy-tcp-timestamp"], d, "synproxy_tcp_timestamp")); err != nil {
		if vv, ok := fortiAPIPatch(o["synproxy-tcp-timestamp"], "PackagesFirewallDosPolicy6Anomaly-SynproxyTcpTimestamp"); ok {
			if err = d.Set("synproxy_tcp_timestamp", vv); err != nil {
				return fmt.Errorf("Error reading synproxy_tcp_timestamp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading synproxy_tcp_timestamp: %v", err)
		}
	}

	if err = d.Set("synproxy_tcp_window", flattenPackagesFirewallDosPolicy6AnomalySynproxyTcpWindow2edl(o["synproxy-tcp-window"], d, "synproxy_tcp_window")); err != nil {
		if vv, ok := fortiAPIPatch(o["synproxy-tcp-window"], "PackagesFirewallDosPolicy6Anomaly-SynproxyTcpWindow"); ok {
			if err = d.Set("synproxy_tcp_window", vv); err != nil {
				return fmt.Errorf("Error reading synproxy_tcp_window: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading synproxy_tcp_window: %v", err)
		}
	}

	if err = d.Set("synproxy_tcp_windowscale", flattenPackagesFirewallDosPolicy6AnomalySynproxyTcpWindowscale2edl(o["synproxy-tcp-windowscale"], d, "synproxy_tcp_windowscale")); err != nil {
		if vv, ok := fortiAPIPatch(o["synproxy-tcp-windowscale"], "PackagesFirewallDosPolicy6Anomaly-SynproxyTcpWindowscale"); ok {
			if err = d.Set("synproxy_tcp_windowscale", vv); err != nil {
				return fmt.Errorf("Error reading synproxy_tcp_windowscale: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading synproxy_tcp_windowscale: %v", err)
		}
	}

	if err = d.Set("synproxy_tos", flattenPackagesFirewallDosPolicy6AnomalySynproxyTos2edl(o["synproxy-tos"], d, "synproxy_tos")); err != nil {
		if vv, ok := fortiAPIPatch(o["synproxy-tos"], "PackagesFirewallDosPolicy6Anomaly-SynproxyTos"); ok {
			if err = d.Set("synproxy_tos", vv); err != nil {
				return fmt.Errorf("Error reading synproxy_tos: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading synproxy_tos: %v", err)
		}
	}

	if err = d.Set("synproxy_ttl", flattenPackagesFirewallDosPolicy6AnomalySynproxyTtl2edl(o["synproxy-ttl"], d, "synproxy_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["synproxy-ttl"], "PackagesFirewallDosPolicy6Anomaly-SynproxyTtl"); ok {
			if err = d.Set("synproxy_ttl", vv); err != nil {
				return fmt.Errorf("Error reading synproxy_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading synproxy_ttl: %v", err)
		}
	}

	if err = d.Set("threshold", flattenPackagesFirewallDosPolicy6AnomalyThreshold2edl(o["threshold"], d, "threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["threshold"], "PackagesFirewallDosPolicy6Anomaly-Threshold"); ok {
			if err = d.Set("threshold", vv); err != nil {
				return fmt.Errorf("Error reading threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading threshold: %v", err)
		}
	}

	if err = d.Set("thresholddefault", flattenPackagesFirewallDosPolicy6AnomalyThresholdDefault2edl(o["threshold(default)"], d, "thresholddefault")); err != nil {
		if vv, ok := fortiAPIPatch(o["threshold(default)"], "PackagesFirewallDosPolicy6Anomaly-ThresholdDefault"); ok {
			if err = d.Set("thresholddefault", vv); err != nil {
				return fmt.Errorf("Error reading thresholddefault: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading thresholddefault: %v", err)
		}
	}

	return nil
}

func flattenPackagesFirewallDosPolicy6AnomalyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandPackagesFirewallDosPolicy6AnomalyAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicy6AnomalyLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicy6AnomalyName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicy6AnomalyQuarantine2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicy6AnomalyQuarantineExpiry2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicy6AnomalyQuarantineLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicy6AnomalyStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicy6AnomalySynproxyTcpMss2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicy6AnomalySynproxyTcpSack2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicy6AnomalySynproxyTcpTimestamp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicy6AnomalySynproxyTcpWindow2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicy6AnomalySynproxyTcpWindowscale2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicy6AnomalySynproxyTos2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicy6AnomalySynproxyTtl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicy6AnomalyThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicy6AnomalyThresholdDefault2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectPackagesFirewallDosPolicy6Anomaly(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandPackagesFirewallDosPolicy6AnomalyAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok || d.HasChange("log") {
		t, err := expandPackagesFirewallDosPolicy6AnomalyLog2edl(d, v, "log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandPackagesFirewallDosPolicy6AnomalyName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("quarantine"); ok || d.HasChange("quarantine") {
		t, err := expandPackagesFirewallDosPolicy6AnomalyQuarantine2edl(d, v, "quarantine")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quarantine"] = t
		}
	}

	if v, ok := d.GetOk("quarantine_expiry"); ok || d.HasChange("quarantine_expiry") {
		t, err := expandPackagesFirewallDosPolicy6AnomalyQuarantineExpiry2edl(d, v, "quarantine_expiry")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quarantine-expiry"] = t
		}
	}

	if v, ok := d.GetOk("quarantine_log"); ok || d.HasChange("quarantine_log") {
		t, err := expandPackagesFirewallDosPolicy6AnomalyQuarantineLog2edl(d, v, "quarantine_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quarantine-log"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandPackagesFirewallDosPolicy6AnomalyStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("synproxy_tcp_mss"); ok || d.HasChange("synproxy_tcp_mss") {
		t, err := expandPackagesFirewallDosPolicy6AnomalySynproxyTcpMss2edl(d, v, "synproxy_tcp_mss")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["synproxy-tcp-mss"] = t
		}
	}

	if v, ok := d.GetOk("synproxy_tcp_sack"); ok || d.HasChange("synproxy_tcp_sack") {
		t, err := expandPackagesFirewallDosPolicy6AnomalySynproxyTcpSack2edl(d, v, "synproxy_tcp_sack")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["synproxy-tcp-sack"] = t
		}
	}

	if v, ok := d.GetOk("synproxy_tcp_timestamp"); ok || d.HasChange("synproxy_tcp_timestamp") {
		t, err := expandPackagesFirewallDosPolicy6AnomalySynproxyTcpTimestamp2edl(d, v, "synproxy_tcp_timestamp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["synproxy-tcp-timestamp"] = t
		}
	}

	if v, ok := d.GetOk("synproxy_tcp_window"); ok || d.HasChange("synproxy_tcp_window") {
		t, err := expandPackagesFirewallDosPolicy6AnomalySynproxyTcpWindow2edl(d, v, "synproxy_tcp_window")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["synproxy-tcp-window"] = t
		}
	}

	if v, ok := d.GetOk("synproxy_tcp_windowscale"); ok || d.HasChange("synproxy_tcp_windowscale") {
		t, err := expandPackagesFirewallDosPolicy6AnomalySynproxyTcpWindowscale2edl(d, v, "synproxy_tcp_windowscale")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["synproxy-tcp-windowscale"] = t
		}
	}

	if v, ok := d.GetOk("synproxy_tos"); ok || d.HasChange("synproxy_tos") {
		t, err := expandPackagesFirewallDosPolicy6AnomalySynproxyTos2edl(d, v, "synproxy_tos")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["synproxy-tos"] = t
		}
	}

	if v, ok := d.GetOk("synproxy_ttl"); ok || d.HasChange("synproxy_ttl") {
		t, err := expandPackagesFirewallDosPolicy6AnomalySynproxyTtl2edl(d, v, "synproxy_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["synproxy-ttl"] = t
		}
	}

	if v, ok := d.GetOk("threshold"); ok || d.HasChange("threshold") {
		t, err := expandPackagesFirewallDosPolicy6AnomalyThreshold2edl(d, v, "threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold"] = t
		}
	}

	if v, ok := d.GetOk("thresholddefault"); ok || d.HasChange("thresholddefault") {
		t, err := expandPackagesFirewallDosPolicy6AnomalyThresholdDefault2edl(d, v, "thresholddefault")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold(default)"] = t
		}
	}

	return &obj, nil
}
