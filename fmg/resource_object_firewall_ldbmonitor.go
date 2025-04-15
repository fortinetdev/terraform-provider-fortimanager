// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01)
// Documentation:
// Hongbin Lu (@fgtdev-hblu), Frank Shen (@frankshen01),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt)

// Description: Configure server load balancing health monitors.

package fortimanager

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceObjectFirewallLdbMonitor() *schema.Resource {
	return &schema.Resource{
		Create: resourceObjectFirewallLdbMonitorCreate,
		Read:   resourceObjectFirewallLdbMonitorRead,
		Update: resourceObjectFirewallLdbMonitorUpdate,
		Delete: resourceObjectFirewallLdbMonitorDelete,

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
			"dns_match_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_request_domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_get": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_max_redirects": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"retry": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"src_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceObjectFirewallLdbMonitorCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectFirewallLdbMonitor(d)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallLdbMonitor resource while getting object: %v", err)
	}
	wsParams["adom"] = adomv

	_, err = c.CreateObjectFirewallLdbMonitor(obj, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error creating ObjectFirewallLdbMonitor resource: %v", err)
	}

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallLdbMonitorRead(d, m)
}

func resourceObjectFirewallLdbMonitorUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectObjectFirewallLdbMonitor(d)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallLdbMonitor resource while getting object: %v", err)
	}

	wsParams["adom"] = adomv

	_, err = c.UpdateObjectFirewallLdbMonitor(obj, mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error updating ObjectFirewallLdbMonitor resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))

	d.SetId(getStringKey(d, "name"))

	return resourceObjectFirewallLdbMonitorRead(d, m)
}

func resourceObjectFirewallLdbMonitorDelete(d *schema.ResourceData, m interface{}) error {
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

	err = c.DeleteObjectFirewallLdbMonitor(mkey, paradict, wsParams)
	if err != nil {
		return fmt.Errorf("Error deleting ObjectFirewallLdbMonitor resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceObjectFirewallLdbMonitorRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadObjectFirewallLdbMonitor(mkey, paradict)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallLdbMonitor resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectObjectFirewallLdbMonitor(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ObjectFirewallLdbMonitor resource from API: %v", err)
	}
	return nil
}

func flattenObjectFirewallLdbMonitorDnsMatchIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallLdbMonitorDnsProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallLdbMonitorDnsRequestDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallLdbMonitorHttpGet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallLdbMonitorHttpMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallLdbMonitorHttpMaxRedirects(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallLdbMonitorInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallLdbMonitorName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallLdbMonitorPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallLdbMonitorRetry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallLdbMonitorSrcIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallLdbMonitorTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenObjectFirewallLdbMonitorType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectObjectFirewallLdbMonitor(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if stValue := d.Get("scopetype"); stValue == "" {
		d.Set("scopetype", "inherit")
	}

	if err = d.Set("dns_match_ip", flattenObjectFirewallLdbMonitorDnsMatchIp(o["dns-match-ip"], d, "dns_match_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-match-ip"], "ObjectFirewallLdbMonitor-DnsMatchIp"); ok {
			if err = d.Set("dns_match_ip", vv); err != nil {
				return fmt.Errorf("Error reading dns_match_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_match_ip: %v", err)
		}
	}

	if err = d.Set("dns_protocol", flattenObjectFirewallLdbMonitorDnsProtocol(o["dns-protocol"], d, "dns_protocol")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-protocol"], "ObjectFirewallLdbMonitor-DnsProtocol"); ok {
			if err = d.Set("dns_protocol", vv); err != nil {
				return fmt.Errorf("Error reading dns_protocol: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_protocol: %v", err)
		}
	}

	if err = d.Set("dns_request_domain", flattenObjectFirewallLdbMonitorDnsRequestDomain(o["dns-request-domain"], d, "dns_request_domain")); err != nil {
		if vv, ok := fortiAPIPatch(o["dns-request-domain"], "ObjectFirewallLdbMonitor-DnsRequestDomain"); ok {
			if err = d.Set("dns_request_domain", vv); err != nil {
				return fmt.Errorf("Error reading dns_request_domain: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading dns_request_domain: %v", err)
		}
	}

	if err = d.Set("http_get", flattenObjectFirewallLdbMonitorHttpGet(o["http-get"], d, "http_get")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-get"], "ObjectFirewallLdbMonitor-HttpGet"); ok {
			if err = d.Set("http_get", vv); err != nil {
				return fmt.Errorf("Error reading http_get: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_get: %v", err)
		}
	}

	if err = d.Set("http_match", flattenObjectFirewallLdbMonitorHttpMatch(o["http-match"], d, "http_match")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-match"], "ObjectFirewallLdbMonitor-HttpMatch"); ok {
			if err = d.Set("http_match", vv); err != nil {
				return fmt.Errorf("Error reading http_match: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_match: %v", err)
		}
	}

	if err = d.Set("http_max_redirects", flattenObjectFirewallLdbMonitorHttpMaxRedirects(o["http-max-redirects"], d, "http_max_redirects")); err != nil {
		if vv, ok := fortiAPIPatch(o["http-max-redirects"], "ObjectFirewallLdbMonitor-HttpMaxRedirects"); ok {
			if err = d.Set("http_max_redirects", vv); err != nil {
				return fmt.Errorf("Error reading http_max_redirects: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading http_max_redirects: %v", err)
		}
	}

	if err = d.Set("interval", flattenObjectFirewallLdbMonitorInterval(o["interval"], d, "interval")); err != nil {
		if vv, ok := fortiAPIPatch(o["interval"], "ObjectFirewallLdbMonitor-Interval"); ok {
			if err = d.Set("interval", vv); err != nil {
				return fmt.Errorf("Error reading interval: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading interval: %v", err)
		}
	}

	if err = d.Set("name", flattenObjectFirewallLdbMonitorName(o["name"], d, "name")); err != nil {
		if vv, ok := fortiAPIPatch(o["name"], "ObjectFirewallLdbMonitor-Name"); ok {
			if err = d.Set("name", vv); err != nil {
				return fmt.Errorf("Error reading name: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("port", flattenObjectFirewallLdbMonitorPort(o["port"], d, "port")); err != nil {
		if vv, ok := fortiAPIPatch(o["port"], "ObjectFirewallLdbMonitor-Port"); ok {
			if err = d.Set("port", vv); err != nil {
				return fmt.Errorf("Error reading port: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("retry", flattenObjectFirewallLdbMonitorRetry(o["retry"], d, "retry")); err != nil {
		if vv, ok := fortiAPIPatch(o["retry"], "ObjectFirewallLdbMonitor-Retry"); ok {
			if err = d.Set("retry", vv); err != nil {
				return fmt.Errorf("Error reading retry: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading retry: %v", err)
		}
	}

	if err = d.Set("src_ip", flattenObjectFirewallLdbMonitorSrcIp(o["src-ip"], d, "src_ip")); err != nil {
		if vv, ok := fortiAPIPatch(o["src-ip"], "ObjectFirewallLdbMonitor-SrcIp"); ok {
			if err = d.Set("src_ip", vv); err != nil {
				return fmt.Errorf("Error reading src_ip: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading src_ip: %v", err)
		}
	}

	if err = d.Set("timeout", flattenObjectFirewallLdbMonitorTimeout(o["timeout"], d, "timeout")); err != nil {
		if vv, ok := fortiAPIPatch(o["timeout"], "ObjectFirewallLdbMonitor-Timeout"); ok {
			if err = d.Set("timeout", vv); err != nil {
				return fmt.Errorf("Error reading timeout: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading timeout: %v", err)
		}
	}

	if err = d.Set("type", flattenObjectFirewallLdbMonitorType(o["type"], d, "type")); err != nil {
		if vv, ok := fortiAPIPatch(o["type"], "ObjectFirewallLdbMonitor-Type"); ok {
			if err = d.Set("type", vv); err != nil {
				return fmt.Errorf("Error reading type: %v", err)
			}
		} else {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	return nil
}

func flattenObjectFirewallLdbMonitorFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandObjectFirewallLdbMonitorDnsMatchIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallLdbMonitorDnsProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallLdbMonitorDnsRequestDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallLdbMonitorHttpGet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallLdbMonitorHttpMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallLdbMonitorHttpMaxRedirects(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallLdbMonitorInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallLdbMonitorName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallLdbMonitorPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallLdbMonitorRetry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallLdbMonitorSrcIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallLdbMonitorTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandObjectFirewallLdbMonitorType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectObjectFirewallLdbMonitor(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dns_match_ip"); ok || d.HasChange("dns_match_ip") {
		t, err := expandObjectFirewallLdbMonitorDnsMatchIp(d, v, "dns_match_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-match-ip"] = t
		}
	}

	if v, ok := d.GetOk("dns_protocol"); ok || d.HasChange("dns_protocol") {
		t, err := expandObjectFirewallLdbMonitorDnsProtocol(d, v, "dns_protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-protocol"] = t
		}
	}

	if v, ok := d.GetOk("dns_request_domain"); ok || d.HasChange("dns_request_domain") {
		t, err := expandObjectFirewallLdbMonitorDnsRequestDomain(d, v, "dns_request_domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-request-domain"] = t
		}
	}

	if v, ok := d.GetOk("http_get"); ok || d.HasChange("http_get") {
		t, err := expandObjectFirewallLdbMonitorHttpGet(d, v, "http_get")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-get"] = t
		}
	}

	if v, ok := d.GetOk("http_match"); ok || d.HasChange("http_match") {
		t, err := expandObjectFirewallLdbMonitorHttpMatch(d, v, "http_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-match"] = t
		}
	}

	if v, ok := d.GetOk("http_max_redirects"); ok || d.HasChange("http_max_redirects") {
		t, err := expandObjectFirewallLdbMonitorHttpMaxRedirects(d, v, "http_max_redirects")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-max-redirects"] = t
		}
	}

	if v, ok := d.GetOk("interval"); ok || d.HasChange("interval") {
		t, err := expandObjectFirewallLdbMonitorInterval(d, v, "interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interval"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok || d.HasChange("name") {
		t, err := expandObjectFirewallLdbMonitorName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok || d.HasChange("port") {
		t, err := expandObjectFirewallLdbMonitorPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("retry"); ok || d.HasChange("retry") {
		t, err := expandObjectFirewallLdbMonitorRetry(d, v, "retry")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["retry"] = t
		}
	}

	if v, ok := d.GetOk("src_ip"); ok || d.HasChange("src_ip") {
		t, err := expandObjectFirewallLdbMonitorSrcIp(d, v, "src_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-ip"] = t
		}
	}

	if v, ok := d.GetOk("timeout"); ok || d.HasChange("timeout") {
		t, err := expandObjectFirewallLdbMonitorTimeout(d, v, "timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok || d.HasChange("type") {
		t, err := expandObjectFirewallLdbMonitorType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	return &obj, nil
}
