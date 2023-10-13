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

func resourcePackagesFirewallDosPolicyAnomaly() *schema.Resource {
	return &schema.Resource{
		Create: resourcePackagesFirewallDosPolicyAnomalyCreate,
		Read:   resourcePackagesFirewallDosPolicyAnomalyRead,
		Update: resourcePackagesFirewallDosPolicyAnomalyUpdate,
		Delete: resourcePackagesFirewallDosPolicyAnomalyDelete,

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
			"dos_policy": &schema.Schema{
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

func resourcePackagesFirewallDosPolicyAnomalyCreate(d *schema.ResourceData, m interface{}) error {
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
	dos_policy := d.Get("dos_policy").(string)
	paradict["pkg_folder_path"] = formatPath(pkg_folder_path)
	paradict["pkg"] = pkg
	paradict["dos_policy"] = dos_policy

	obj, err := getObjectPackagesFirewallDosPolicyAnomaly(d)
	if err != nil {
		return fmt.Errorf("Error creating PackagesFirewallDosPolicyAnomaly resource while getting object: %v", err)
	}

	_, err = c.CreatePackagesFirewallDosPolicyAnomaly(obj, paradict)

	if err != nil {
		return fmt.Errorf("Error creating PackagesFirewallDosPolicyAnomaly resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourcePackagesFirewallDosPolicyAnomalyRead(d, m)
}

func resourcePackagesFirewallDosPolicyAnomalyUpdate(d *schema.ResourceData, m interface{}) error {
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
	dos_policy := d.Get("dos_policy").(string)
	paradict["pkg_folder_path"] = formatPath(pkg_folder_path)
	paradict["pkg"] = pkg
	paradict["dos_policy"] = dos_policy

	obj, err := getObjectPackagesFirewallDosPolicyAnomaly(d)
	if err != nil {
		return fmt.Errorf("Error updating PackagesFirewallDosPolicyAnomaly resource while getting object: %v", err)
	}

	_, err = c.UpdatePackagesFirewallDosPolicyAnomaly(obj, mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error updating PackagesFirewallDosPolicyAnomaly resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourcePackagesFirewallDosPolicyAnomalyRead(d, m)
}

func resourcePackagesFirewallDosPolicyAnomalyDelete(d *schema.ResourceData, m interface{}) error {
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
	dos_policy := d.Get("dos_policy").(string)
	paradict["pkg_folder_path"] = formatPath(pkg_folder_path)
	paradict["pkg"] = pkg
	paradict["dos_policy"] = dos_policy

	err = c.DeletePackagesFirewallDosPolicyAnomaly(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error deleting PackagesFirewallDosPolicyAnomaly resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourcePackagesFirewallDosPolicyAnomalyRead(d *schema.ResourceData, m interface{}) error {
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
	dos_policy := d.Get("dos_policy").(string)
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
	if dos_policy == "" {
		dos_policy = importOptionChecking(m.(*FortiClient).Cfg, "dos_policy")
		if dos_policy == "" {
			return fmt.Errorf("Parameter dos_policy is missing")
		}
		if err = d.Set("dos_policy", dos_policy); err != nil {
			return fmt.Errorf("Error set params dos_policy: %v", err)
		}
	}
	paradict["pkg_folder_path"] = formatPath(pkg_folder_path)
	paradict["pkg"] = pkg
	paradict["dos_policy"] = dos_policy

	o, err := c.ReadPackagesFirewallDosPolicyAnomaly(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading PackagesFirewallDosPolicyAnomaly resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectPackagesFirewallDosPolicyAnomaly(d, o)
	if err != nil {
		return fmt.Errorf("Error reading PackagesFirewallDosPolicyAnomaly resource from API: %v", err)
	}
	return nil
}

func flattenPackagesFirewallDosPolicyAnomalyAction2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicyAnomalyLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicyAnomalyName2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicyAnomalyQuarantine2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicyAnomalyQuarantineExpiry2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicyAnomalyQuarantineLog2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicyAnomalyStatus2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicyAnomalySynproxyTcpMss2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicyAnomalySynproxyTcpSack2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicyAnomalySynproxyTcpTimestamp2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicyAnomalySynproxyTcpWindow2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicyAnomalySynproxyTcpWindowscale2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicyAnomalySynproxyTos2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicyAnomalySynproxyTtl2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicyAnomalyThreshold2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenPackagesFirewallDosPolicyAnomalyThresholdDefault2edl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectPackagesFirewallDosPolicyAnomaly(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("action", flattenPackagesFirewallDosPolicyAnomalyAction2edl(o["action"], d, "action")); err != nil {
		if vv, ok := fortiAPIPatch(o["action"], "PackagesFirewallDosPolicyAnomaly-Action"); ok {
			if err = d.Set("action", vv); err != nil {
				return fmt.Errorf("Error reading action: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("log", flattenPackagesFirewallDosPolicyAnomalyLog2edl(o["log"], d, "log")); err != nil {
		if vv, ok := fortiAPIPatch(o["log"], "PackagesFirewallDosPolicyAnomaly-Log"); ok {
			if err = d.Set("log", vv); err != nil {
				return fmt.Errorf("Error reading log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	if err = d.Set("name", flattenPackagesFirewallDosPolicyAnomalyName2edl(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "PackagesFirewallDosPolicyAnomaly-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("quarantine", flattenPackagesFirewallDosPolicyAnomalyQuarantine2edl(o["quarantine"], d, "quarantine")); err != nil {
		if vv, ok := fortiAPIPatch(o["quarantine"], "PackagesFirewallDosPolicyAnomaly-Quarantine"); ok {
			if err = d.Set("quarantine", vv); err != nil {
				return fmt.Errorf("Error reading quarantine: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading quarantine: %v", err)
		}
	}

	if err = d.Set("quarantine_expiry", flattenPackagesFirewallDosPolicyAnomalyQuarantineExpiry2edl(o["quarantine-expiry"], d, "quarantine_expiry")); err != nil {
		if vv, ok := fortiAPIPatch(o["quarantine-expiry"], "PackagesFirewallDosPolicyAnomaly-QuarantineExpiry"); ok {
			if err = d.Set("quarantine_expiry", vv); err != nil {
				return fmt.Errorf("Error reading quarantine_expiry: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading quarantine_expiry: %v", err)
		}
	}

	if err = d.Set("quarantine_log", flattenPackagesFirewallDosPolicyAnomalyQuarantineLog2edl(o["quarantine-log"], d, "quarantine_log")); err != nil {
		if vv, ok := fortiAPIPatch(o["quarantine-log"], "PackagesFirewallDosPolicyAnomaly-QuarantineLog"); ok {
			if err = d.Set("quarantine_log", vv); err != nil {
				return fmt.Errorf("Error reading quarantine_log: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading quarantine_log: %v", err)
		}
	}

	if err = d.Set("status", flattenPackagesFirewallDosPolicyAnomalyStatus2edl(o["status"], d, "status")); err != nil {
		if vv, ok := fortiAPIPatch(o["status"], "PackagesFirewallDosPolicyAnomaly-Status"); ok {
			if err = d.Set("status", vv); err != nil {
				return fmt.Errorf("Error reading status: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("synproxy_tcp_mss", flattenPackagesFirewallDosPolicyAnomalySynproxyTcpMss2edl(o["synproxy-tcp-mss"], d, "synproxy_tcp_mss")); err != nil {
		if vv, ok := fortiAPIPatch(o["synproxy-tcp-mss"], "PackagesFirewallDosPolicyAnomaly-SynproxyTcpMss"); ok {
			if err = d.Set("synproxy_tcp_mss", vv); err != nil {
				return fmt.Errorf("Error reading synproxy_tcp_mss: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading synproxy_tcp_mss: %v", err)
		}
	}

	if err = d.Set("synproxy_tcp_sack", flattenPackagesFirewallDosPolicyAnomalySynproxyTcpSack2edl(o["synproxy-tcp-sack"], d, "synproxy_tcp_sack")); err != nil {
		if vv, ok := fortiAPIPatch(o["synproxy-tcp-sack"], "PackagesFirewallDosPolicyAnomaly-SynproxyTcpSack"); ok {
			if err = d.Set("synproxy_tcp_sack", vv); err != nil {
				return fmt.Errorf("Error reading synproxy_tcp_sack: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading synproxy_tcp_sack: %v", err)
		}
	}

	if err = d.Set("synproxy_tcp_timestamp", flattenPackagesFirewallDosPolicyAnomalySynproxyTcpTimestamp2edl(o["synproxy-tcp-timestamp"], d, "synproxy_tcp_timestamp")); err != nil {
		if vv, ok := fortiAPIPatch(o["synproxy-tcp-timestamp"], "PackagesFirewallDosPolicyAnomaly-SynproxyTcpTimestamp"); ok {
			if err = d.Set("synproxy_tcp_timestamp", vv); err != nil {
				return fmt.Errorf("Error reading synproxy_tcp_timestamp: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading synproxy_tcp_timestamp: %v", err)
		}
	}

	if err = d.Set("synproxy_tcp_window", flattenPackagesFirewallDosPolicyAnomalySynproxyTcpWindow2edl(o["synproxy-tcp-window"], d, "synproxy_tcp_window")); err != nil {
		if vv, ok := fortiAPIPatch(o["synproxy-tcp-window"], "PackagesFirewallDosPolicyAnomaly-SynproxyTcpWindow"); ok {
			if err = d.Set("synproxy_tcp_window", vv); err != nil {
				return fmt.Errorf("Error reading synproxy_tcp_window: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading synproxy_tcp_window: %v", err)
		}
	}

	if err = d.Set("synproxy_tcp_windowscale", flattenPackagesFirewallDosPolicyAnomalySynproxyTcpWindowscale2edl(o["synproxy-tcp-windowscale"], d, "synproxy_tcp_windowscale")); err != nil {
		if vv, ok := fortiAPIPatch(o["synproxy-tcp-windowscale"], "PackagesFirewallDosPolicyAnomaly-SynproxyTcpWindowscale"); ok {
			if err = d.Set("synproxy_tcp_windowscale", vv); err != nil {
				return fmt.Errorf("Error reading synproxy_tcp_windowscale: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading synproxy_tcp_windowscale: %v", err)
		}
	}

	if err = d.Set("synproxy_tos", flattenPackagesFirewallDosPolicyAnomalySynproxyTos2edl(o["synproxy-tos"], d, "synproxy_tos")); err != nil {
		if vv, ok := fortiAPIPatch(o["synproxy-tos"], "PackagesFirewallDosPolicyAnomaly-SynproxyTos"); ok {
			if err = d.Set("synproxy_tos", vv); err != nil {
				return fmt.Errorf("Error reading synproxy_tos: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading synproxy_tos: %v", err)
		}
	}

	if err = d.Set("synproxy_ttl", flattenPackagesFirewallDosPolicyAnomalySynproxyTtl2edl(o["synproxy-ttl"], d, "synproxy_ttl")); err != nil {
		if vv, ok := fortiAPIPatch(o["synproxy-ttl"], "PackagesFirewallDosPolicyAnomaly-SynproxyTtl"); ok {
			if err = d.Set("synproxy_ttl", vv); err != nil {
				return fmt.Errorf("Error reading synproxy_ttl: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading synproxy_ttl: %v", err)
		}
	}

	if err = d.Set("threshold", flattenPackagesFirewallDosPolicyAnomalyThreshold2edl(o["threshold"], d, "threshold")); err != nil {
		if vv, ok := fortiAPIPatch(o["threshold"], "PackagesFirewallDosPolicyAnomaly-Threshold"); ok {
			if err = d.Set("threshold", vv); err != nil {
				return fmt.Errorf("Error reading threshold: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading threshold: %v", err)
		}
	}

	if err = d.Set("thresholddefault", flattenPackagesFirewallDosPolicyAnomalyThresholdDefault2edl(o["threshold(default)"], d, "thresholddefault")); err != nil {
		if vv, ok := fortiAPIPatch(o["threshold(default)"], "PackagesFirewallDosPolicyAnomaly-ThresholdDefault"); ok {
			if err = d.Set("thresholddefault", vv); err != nil {
				return fmt.Errorf("Error reading thresholddefault: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading thresholddefault: %v", err)
		}
	}

	return nil
}

func flattenPackagesFirewallDosPolicyAnomalyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandPackagesFirewallDosPolicyAnomalyAction2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicyAnomalyLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicyAnomalyName2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicyAnomalyQuarantine2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicyAnomalyQuarantineExpiry2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicyAnomalyQuarantineLog2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicyAnomalyStatus2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicyAnomalySynproxyTcpMss2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicyAnomalySynproxyTcpSack2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicyAnomalySynproxyTcpTimestamp2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicyAnomalySynproxyTcpWindow2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicyAnomalySynproxyTcpWindowscale2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicyAnomalySynproxyTos2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicyAnomalySynproxyTtl2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicyAnomalyThreshold2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandPackagesFirewallDosPolicyAnomalyThresholdDefault2edl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectPackagesFirewallDosPolicyAnomaly(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok || d.HasChange("action") {
		t, err := expandPackagesFirewallDosPolicyAnomalyAction2edl(d, v, "action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok || d.HasChange("log") {
		t, err := expandPackagesFirewallDosPolicyAnomalyLog2edl(d, v, "log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandPackagesFirewallDosPolicyAnomalyName2edl(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("quarantine"); ok || d.HasChange("quarantine") {
		t, err := expandPackagesFirewallDosPolicyAnomalyQuarantine2edl(d, v, "quarantine")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quarantine"] = t
		}
	}

	if v, ok := d.GetOk("quarantine_expiry"); ok || d.HasChange("quarantine_expiry") {
		t, err := expandPackagesFirewallDosPolicyAnomalyQuarantineExpiry2edl(d, v, "quarantine_expiry")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quarantine-expiry"] = t
		}
	}

	if v, ok := d.GetOk("quarantine_log"); ok || d.HasChange("quarantine_log") {
		t, err := expandPackagesFirewallDosPolicyAnomalyQuarantineLog2edl(d, v, "quarantine_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quarantine-log"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok || d.HasChange("status") {
		t, err := expandPackagesFirewallDosPolicyAnomalyStatus2edl(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("synproxy_tcp_mss"); ok || d.HasChange("synproxy_tcp_mss") {
		t, err := expandPackagesFirewallDosPolicyAnomalySynproxyTcpMss2edl(d, v, "synproxy_tcp_mss")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["synproxy-tcp-mss"] = t
		}
	}

	if v, ok := d.GetOk("synproxy_tcp_sack"); ok || d.HasChange("synproxy_tcp_sack") {
		t, err := expandPackagesFirewallDosPolicyAnomalySynproxyTcpSack2edl(d, v, "synproxy_tcp_sack")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["synproxy-tcp-sack"] = t
		}
	}

	if v, ok := d.GetOk("synproxy_tcp_timestamp"); ok || d.HasChange("synproxy_tcp_timestamp") {
		t, err := expandPackagesFirewallDosPolicyAnomalySynproxyTcpTimestamp2edl(d, v, "synproxy_tcp_timestamp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["synproxy-tcp-timestamp"] = t
		}
	}

	if v, ok := d.GetOk("synproxy_tcp_window"); ok || d.HasChange("synproxy_tcp_window") {
		t, err := expandPackagesFirewallDosPolicyAnomalySynproxyTcpWindow2edl(d, v, "synproxy_tcp_window")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["synproxy-tcp-window"] = t
		}
	}

	if v, ok := d.GetOk("synproxy_tcp_windowscale"); ok || d.HasChange("synproxy_tcp_windowscale") {
		t, err := expandPackagesFirewallDosPolicyAnomalySynproxyTcpWindowscale2edl(d, v, "synproxy_tcp_windowscale")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["synproxy-tcp-windowscale"] = t
		}
	}

	if v, ok := d.GetOk("synproxy_tos"); ok || d.HasChange("synproxy_tos") {
		t, err := expandPackagesFirewallDosPolicyAnomalySynproxyTos2edl(d, v, "synproxy_tos")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["synproxy-tos"] = t
		}
	}

	if v, ok := d.GetOk("synproxy_ttl"); ok || d.HasChange("synproxy_ttl") {
		t, err := expandPackagesFirewallDosPolicyAnomalySynproxyTtl2edl(d, v, "synproxy_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["synproxy-ttl"] = t
		}
	}

	if v, ok := d.GetOk("threshold"); ok || d.HasChange("threshold") {
		t, err := expandPackagesFirewallDosPolicyAnomalyThreshold2edl(d, v, "threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold"] = t
		}
	}

	if v, ok := d.GetOk("thresholddefault"); ok || d.HasChange("thresholddefault") {
		t, err := expandPackagesFirewallDosPolicyAnomalyThresholdDefault2edl(d, v, "thresholddefault")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["threshold(default)"] = t
		}
	}

	return &obj, nil
}
