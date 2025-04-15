// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure per-IP traffic shaper.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallShaperPerIpShaper() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallShaperPerIpShaperCreate,
		Read:   resourceObjectFirewallShaperPerIpShaperRead,
		Update: resourceObjectFirewallShaperPerIpShaperUpdate,
		Delete: resourceObjectFirewallShaperPerIpShaperDelete,

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
			"bandwidth_unit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diffserv_forward": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diffserv_reverse": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diffservcode_forward": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diffservcode_rev": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"max_bandwidth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_concurrent_session": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_concurrent_tcp_session": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_concurrent_udp_session": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

func resourceObjectFirewallShaperPerIpShaperCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectFirewallShaperPerIpShaper(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallShaperPerIpShaper resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectFirewallShaperPerIpShaper(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallShaperPerIpShaper resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallShaperPerIpShaperRead(d, m)
}

func resourceObjectFirewallShaperPerIpShaperUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectFirewallShaperPerIpShaper(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallShaperPerIpShaper resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectFirewallShaperPerIpShaper(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallShaperPerIpShaper resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallShaperPerIpShaperRead(d, m)
}

func resourceObjectFirewallShaperPerIpShaperDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectFirewallShaperPerIpShaper(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallShaperPerIpShaper resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallShaperPerIpShaperRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectFirewallShaperPerIpShaper(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallShaperPerIpShaper resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallShaperPerIpShaper(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallShaperPerIpShaper resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallShaperPerIpShaperBandwidthUnit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallShaperPerIpShaperDiffservForward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallShaperPerIpShaperDiffservReverse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallShaperPerIpShaperDiffservcodeForward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallShaperPerIpShaperDiffservcodeRev(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallShaperPerIpShaperMaxBandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallShaperPerIpShaperMaxConcurrentSession(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallShaperPerIpShaperMaxConcurrentTcpSession(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallShaperPerIpShaperMaxConcurrentUdpSession(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallShaperPerIpShaperName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallShaperPerIpShaper(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("bandwidth_unit", flattenObjectFirewallShaperPerIpShaperBandwidthUnit(o["bandwidth-unit"], d, "bandwidth_unit")); err != nil {
		if vv, ok := fortiAPIPatch(o["bandwidth-unit"], "ObjectFirewallShaperPerIpShaper-BandwidthUnit"); ok {
			if err = d.Set("bandwidth_unit", vv); err != nil {
				return fmt.Errorf("Error reading bandwidth_unit: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading bandwidth_unit: %v", err)
		}
	}

	if err = d.Set("diffserv_forward", flattenObjectFirewallShaperPerIpShaperDiffservForward(o["diffserv-forward"], d, "diffserv_forward")); err != nil {
		if vv, ok := fortiAPIPatch(o["diffserv-forward"], "ObjectFirewallShaperPerIpShaper-DiffservForward"); ok {
			if err = d.Set("diffserv_forward", vv); err != nil {
				return fmt.Errorf("Error reading diffserv_forward: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diffserv_forward: %v", err)
		}
	}

	if err = d.Set("diffserv_reverse", flattenObjectFirewallShaperPerIpShaperDiffservReverse(o["diffserv-reverse"], d, "diffserv_reverse")); err != nil {
		if vv, ok := fortiAPIPatch(o["diffserv-reverse"], "ObjectFirewallShaperPerIpShaper-DiffservReverse"); ok {
			if err = d.Set("diffserv_reverse", vv); err != nil {
				return fmt.Errorf("Error reading diffserv_reverse: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diffserv_reverse: %v", err)
		}
	}

	if err = d.Set("diffservcode_forward", flattenObjectFirewallShaperPerIpShaperDiffservcodeForward(o["diffservcode-forward"], d, "diffservcode_forward")); err != nil {
		if vv, ok := fortiAPIPatch(o["diffservcode-forward"], "ObjectFirewallShaperPerIpShaper-DiffservcodeForward"); ok {
			if err = d.Set("diffservcode_forward", vv); err != nil {
				return fmt.Errorf("Error reading diffservcode_forward: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diffservcode_forward: %v", err)
		}
	}

	if err = d.Set("diffservcode_rev", flattenObjectFirewallShaperPerIpShaperDiffservcodeRev(o["diffservcode-rev"], d, "diffservcode_rev")); err != nil {
		if vv, ok := fortiAPIPatch(o["diffservcode-rev"], "ObjectFirewallShaperPerIpShaper-DiffservcodeRev"); ok {
			if err = d.Set("diffservcode_rev", vv); err != nil {
				return fmt.Errorf("Error reading diffservcode_rev: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading diffservcode_rev: %v", err)
		}
	}

	if err = d.Set("max_bandwidth", flattenObjectFirewallShaperPerIpShaperMaxBandwidth(o["max-bandwidth"], d, "max_bandwidth")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-bandwidth"], "ObjectFirewallShaperPerIpShaper-MaxBandwidth"); ok {
			if err = d.Set("max_bandwidth", vv); err != nil {
				return fmt.Errorf("Error reading max_bandwidth: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_bandwidth: %v", err)
		}
	}

	if err = d.Set("max_concurrent_session", flattenObjectFirewallShaperPerIpShaperMaxConcurrentSession(o["max-concurrent-session"], d, "max_concurrent_session")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-concurrent-session"], "ObjectFirewallShaperPerIpShaper-MaxConcurrentSession"); ok {
			if err = d.Set("max_concurrent_session", vv); err != nil {
				return fmt.Errorf("Error reading max_concurrent_session: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_concurrent_session: %v", err)
		}
	}

	if err = d.Set("max_concurrent_tcp_session", flattenObjectFirewallShaperPerIpShaperMaxConcurrentTcpSession(o["max-concurrent-tcp-session"], d, "max_concurrent_tcp_session")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-concurrent-tcp-session"], "ObjectFirewallShaperPerIpShaper-MaxConcurrentTcpSession"); ok {
			if err = d.Set("max_concurrent_tcp_session", vv); err != nil {
				return fmt.Errorf("Error reading max_concurrent_tcp_session: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_concurrent_tcp_session: %v", err)
		}
	}

	if err = d.Set("max_concurrent_udp_session", flattenObjectFirewallShaperPerIpShaperMaxConcurrentUdpSession(o["max-concurrent-udp-session"], d, "max_concurrent_udp_session")); err != nil {
		if vv, ok := fortiAPIPatch(o["max-concurrent-udp-session"], "ObjectFirewallShaperPerIpShaper-MaxConcurrentUdpSession"); ok {
			if err = d.Set("max_concurrent_udp_session", vv); err != nil {
				return fmt.Errorf("Error reading max_concurrent_udp_session: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading max_concurrent_udp_session: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectFirewallShaperPerIpShaperName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectFirewallShaperPerIpShaper-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallShaperPerIpShaperFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallShaperPerIpShaperBandwidthUnit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShaperPerIpShaperDiffservForward(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShaperPerIpShaperDiffservReverse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShaperPerIpShaperDiffservcodeForward(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShaperPerIpShaperDiffservcodeRev(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShaperPerIpShaperMaxBandwidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShaperPerIpShaperMaxConcurrentSession(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShaperPerIpShaperMaxConcurrentTcpSession(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShaperPerIpShaperMaxConcurrentUdpSession(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallShaperPerIpShaperName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallShaperPerIpShaper(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("bandwidth_unit"); ok || d.HasChange("bandwidth_unit") {
		t, err := expandObjectFirewallShaperPerIpShaperBandwidthUnit(d, v, "bandwidth_unit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bandwidth-unit"] = t
		}
	}

	if v, ok := d.GetOk("diffserv_forward"); ok || d.HasChange("diffserv_forward") {
		t, err := expandObjectFirewallShaperPerIpShaperDiffservForward(d, v, "diffserv_forward")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffserv-forward"] = t
		}
	}

	if v, ok := d.GetOk("diffserv_reverse"); ok || d.HasChange("diffserv_reverse") {
		t, err := expandObjectFirewallShaperPerIpShaperDiffservReverse(d, v, "diffserv_reverse")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffserv-reverse"] = t
		}
	}

	if v, ok := d.GetOk("diffservcode_forward"); ok || d.HasChange("diffservcode_forward") {
		t, err := expandObjectFirewallShaperPerIpShaperDiffservcodeForward(d, v, "diffservcode_forward")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffservcode-forward"] = t
		}
	}

	if v, ok := d.GetOk("diffservcode_rev"); ok || d.HasChange("diffservcode_rev") {
		t, err := expandObjectFirewallShaperPerIpShaperDiffservcodeRev(d, v, "diffservcode_rev")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diffservcode-rev"] = t
		}
	}

	if v, ok := d.GetOk("max_bandwidth"); ok || d.HasChange("max_bandwidth") {
		t, err := expandObjectFirewallShaperPerIpShaperMaxBandwidth(d, v, "max_bandwidth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-bandwidth"] = t
		}
	}

	if v, ok := d.GetOk("max_concurrent_session"); ok || d.HasChange("max_concurrent_session") {
		t, err := expandObjectFirewallShaperPerIpShaperMaxConcurrentSession(d, v, "max_concurrent_session")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-concurrent-session"] = t
		}
	}

	if v, ok := d.GetOk("max_concurrent_tcp_session"); ok || d.HasChange("max_concurrent_tcp_session") {
		t, err := expandObjectFirewallShaperPerIpShaperMaxConcurrentTcpSession(d, v, "max_concurrent_tcp_session")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-concurrent-tcp-session"] = t
		}
	}

	if v, ok := d.GetOk("max_concurrent_udp_session"); ok || d.HasChange("max_concurrent_udp_session") {
		t, err := expandObjectFirewallShaperPerIpShaperMaxConcurrentUdpSession(d, v, "max_concurrent_udp_session")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-concurrent-udp-session"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectFirewallShaperPerIpShaperName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
